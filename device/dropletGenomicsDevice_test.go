package device

// func TestSetupDeviceURL(t *testing.T) {
// 	var dummyDevice = DropletGenomicsDevice{"192.168.1.100", 8764}
// 	correctAnswer := "http://192.168.1.100:8764"

// 	generatedURL := setupDeviceURL(&dummyDevice)
// 	if !strings.EqualFold(generatedURL, correctAnswer) {
// 		t.Errorf("Setting url failed with URL: %s", generatedURL)
// 	}
// }

// func TestAvailable(t *testing.T) {

// 	var devices = []DropletGenomicsDevice{
// 		DropletGenomicsDevice{"google.com", 80},
// 		DropletGenomicsDevice{"jsonplaceholder.typicode.com", 80},
// 	}

// 	for _, device := range devices {
// 		if !device.Available() {
// 			t.Errorf("Error in Available function")
// 		}
// 	}
// }
