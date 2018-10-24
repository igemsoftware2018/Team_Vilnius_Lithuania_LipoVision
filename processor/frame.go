package processor

import (
	"image"
	"image/color"
	"sync/atomic"

	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device"
	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/filter"
	log "github.com/sirupsen/logrus"
	"gocv.io/x/gocv"
)

/* These constants define what streams are supported on this Processor.
StreamRegion      - matched region frame
StreamThresholded - full frame that is thresholded
StreamOriginal    - original frame, as it came from device*/
const (
	StreamRegion      string = "region"
	StreamThresholded string = "thresholded"
	StreamOriginal    string = "original"
)

/* There are the settings that are settable for this Processor
SettingAutonomicRun - enables the auto-coating feature*/
const (
	SettingAutonomicRun string = "auto"
	SettingRegionIsSet  string = "regionSet"
	SettingDangerScore  string = "dangerScore"
)

const (
	frameQueueSize   int = 10
	detectorTreshold int = 10
)

// NewFrameProcessor Creates a frame processor with given settings
func NewFrameProcessor(template image.Image) *FrameProcessor {
	log.WithFields(log.Fields{
		"processor": "Frame",
	}).Info("Created")

	prevFrame := gocv.NewMat()
	counter := 0
	return &FrameProcessor{
		frameCounter: &counter,
		previousFame: &prevFrame,
		frameFilters: []filter.Filter{
			filter.CreateNoiseFilter(&prevFrame, &counter),
		},
		regionFilters: []filter.Filter{
			filter.CreateVerticalFilter(),
			filter.CreateLineApplyFilter(),
		},
		transformFilters: []filter.Filter{
			filter.CreateCvtColorFilter(gocv.ColorBGRToGray),
			filter.CreateThesholdFilter(125, 255, gocv.ThresholdBinaryInv),
		},
		template:     templateSetup(template),
		region:       image.Rectangle{},
		subtractor:   gocv.NewBackgroundSubtractorKNN(),
		dangerCount:  0,
		autonomicRun: 0,
	}
}

// FrameProcessor Defines a processor for incoming frames of the stream
type FrameProcessor struct {
	Processor

	// Settings
	// Some of these are ints instead of bools
	// to protect thread safety
	autonomicRun int32
	regionIsSet  int32

	// Frame threshholding and type filter
	transformFilters []filter.Filter

	// Full frame filters
	frameCounter *int
	previousFame *gocv.Mat
	frameFilters []filter.Filter

	// Region based filters
	regionFilters []filter.Filter

	// Misc
	subtractor  gocv.BackgroundSubtractorKNN
	template    gocv.Mat
	region      image.Rectangle
	rectColor   color.RGBA
	dangerCount int32
}

func templateSetup(templateImage image.Image) (filteredTemplate gocv.Mat) {
	template, err := gocv.ImageToMatRGB(templateImage)
	if err != nil {
		panic(err)
	}
	gocv.CvtColor(template, &template, gocv.ColorBGRToGray)
	gocv.AdaptiveThreshold(template, &template, 255, gocv.AdaptiveThresholdMean, gocv.ThresholdBinaryInv, 31, 2)
	gocv.Erode(template, &template, gocv.GetStructuringElement(0, image.Pt(3, 3)))
	return template
}

func (fp *FrameProcessor) extractAndSendThresholded(recvFrame device.Frame, streamHandlers map[string]func(image.Image)) (gocv.Mat, gocv.Mat, error) {
	original, decodeErr := gocv.ImageToMatRGB(recvFrame.Frame())
	if decodeErr != nil {
		return gocv.Mat{}, gocv.Mat{}, decodeErr
	}

	frame := original.Clone()
	if err := filter.ApplyFilters(&frame, fp.transformFilters); err != nil {
		return gocv.Mat{}, gocv.Mat{}, err
	}

	return original, frame, nil
}

// Launch starts a processing goroutine for the stream of frames coming from a device
func (fp *FrameProcessor) Launch(frames <-chan device.Frame, streamHandlers map[string]func(image.Image)) {
	log.WithFields(log.Fields{"processor": "Frame"}).Info("Launched")

	go func() {
		for recvFrame := range frames {
			select {
			case <-recvFrame.Skip():
				log.WithFields(log.Fields{"processor": "Frame"}).Warn("Frame skip")
				continue
			default:
			}

			// Clone so that at least something is shown, and does not crash
			cropped := fp.template.Clone()
			original, frame, extractErr := fp.extractAndSendThresholded(recvFrame, streamHandlers)
			if extractErr != nil {
				log.WithFields(log.Fields{
					"processor": "Frame",
				}).Error("failed to extract: ", extractErr)
			}

			if atomic.LoadInt32(&fp.regionIsSet) == 0 {
				fp.rectColor = color.RGBA{255, 255, 255, 0}
				templateDims := fp.template.Size()
				regionResult := gocv.NewMat()
				gocv.MatchTemplate(frame, fp.template, &regionResult, gocv.TmCcoeff, gocv.NewMat())
				_, _, _, maxLoc := gocv.MinMaxLoc(regionResult)
				fp.region.Min = maxLoc
				fp.region.Max = image.Point{X: maxLoc.X + templateDims[0], Y: maxLoc.Y + templateDims[1]}
				frameRegion := frame.Region(fp.region)
				cropped = frameRegion.Clone()
			} else {
				fp.rectColor = color.RGBA{255, 0, 0, 0}

				if err := filter.ApplyFilters(&frame, fp.frameFilters); err != nil {
					log.WithFields(log.Fields{
						"processor": "Frame",
					}).Error("failed: ", err)
				}
				frameRegion := frame.Region(fp.region)
				cropped = frameRegion.Clone()

				croppedSubtracted := cropped.Clone()
				fp.subtractor.Apply(croppedSubtracted, &croppedSubtracted)

				if err := filter.ApplyFilters(&cropped, fp.regionFilters); err != nil {
					log.WithFields(log.Fields{
						"processor": "Frame",
					}).Error("failed: ", err)
					continue
				}

				furthestLineX := findBiggestXOfWhitePixel(&cropped)

				newCount := countMisbehaviorPixels(croppedSubtracted, furthestLineX)
				if newCount < 500 {
					fp.dangerCount = fp.dangerCount + countMisbehaviorPixels(croppedSubtracted, furthestLineX)
				}
				log.Info("Danger count: ", fp.dangerCount)

				gocv.AddWeighted(cropped, 1, croppedSubtracted, 1, 0, &cropped)
			}

			gocv.Rectangle(&original, fp.region, fp.rectColor, 2)
			invokeIfPresent(streamHandlers, StreamThresholded, &frame)
			invokeIfPresent(streamHandlers, StreamOriginal, &original)
			invokeIfPresent(streamHandlers, StreamRegion, &cropped)
			if fp.dangerCount > 100 {
				fp.dangerCount = fp.dangerCount - 50
			} else {
				fp.dangerCount = 0
			}
			fp.previousFame = &frame
		}
		fp.subtractor.Close()
		log.WithFields(log.Fields{"processor": "Frame"}).Info("Finished")
	}()

}

// Set the configurables on this Processor
func (fp *FrameProcessor) Set(id string, value interface{}) {
	switch id {
	case SettingRegionIsSet:
		val, ok := value.(int32)
		if !ok {
			panic("frameProcessor cast failed, bad value")
		}
		atomic.StoreInt32(&fp.regionIsSet, val)
	case SettingAutonomicRun:
		val, ok := value.(int32)
		if !ok {
			panic("frameProcessor cast failed, bad value")
		}
		atomic.StoreInt32(&fp.autonomicRun, val)
	case SettingDangerScore:
		val, ok := value.(int32)
		if !ok {
			panic("frameProcessor cast failed, bad value")
		}
		atomic.StoreInt32(&fp.dangerCount, val)
	default:
		panic("nonexistant value given")
	}
}

func (fp *FrameProcessor) Get(id string) int32 {
	switch id {
	case SettingRegionIsSet:
		return atomic.LoadInt32(&fp.regionIsSet)
	case SettingAutonomicRun:
		return atomic.LoadInt32(&fp.autonomicRun)
	case SettingDangerScore:
		return atomic.LoadInt32(&fp.dangerCount)
	default:
		panic("nonexistant value given")
	}
}

// Find white pixel with biggest X coordinate
func findBiggestXOfWhitePixel(cropped *gocv.Mat) (biggestX int) {
	biggestX = 0
	for i := 0; i < cropped.Cols(); i++ {
		for j := 0; j < cropped.Rows(); j++ {
			if cropped.GetUCharAt(j, i) == 255 {
				if biggestX < i {
					biggestX = i
				}

			}
		}
	}
	return
}

/* Counts white pixels with bigger X value than outer right vertical line
   returns true if there are some, false otherwise */
func countMisbehaviorPixels(cropped gocv.Mat, biggestX int) int32 {
	dangerPixelsCounter := 0
	for i := 0; i < cropped.Cols(); i++ {
		for j := 0; j < cropped.Rows(); j++ {
			if cropped.GetUCharAt(j, i) == 255 && i > biggestX {
				dangerPixelsCounter++
			}
		}
	}
	return int32(dangerPixelsCounter)
}
