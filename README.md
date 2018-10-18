[![Maintainability](https://api.codeclimate.com/v1/badges/22563fda48d0b0e85cab/maintainability)](https://codeclimate.com/github/Vilnius-Lithuania-iGEM-2018/lipovision/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/22563fda48d0b0e85cab/test_coverage)](https://codeclimate.com/github/Vilnius-Lithuania-iGEM-2018/lipovision/test_coverage)
[![Go Report Card](https://goreportcard.com/badge/github.com/Vilnius-Lithuania-iGEM-2018/lipovision)](https://goreportcard.com/report/github.com/Vilnius-Lithuania-iGEM-2018/lipovision)
[![License: AGPL v3](https://img.shields.io/badge/License-AGPL%20v3-blue.svg)](https://www.gnu.org/licenses/agpl-3.0)

# Lipovision
Automate liposome 


|Platform|Build Status|
|--------|------------|
| Linux x64 |[![Build Status](https://travis-ci.org/Vilnius-Lithuania-iGEM-2018/lipovision.svg?branch=master)](https://travis-ci.org/Vilnius-Lithuania-iGEM-2018/lipovision)|
| Windows |[![Build status](https://ci.appveyor.com/api/projects/status/cbq6uq3iwqhkywwt/branch/master?svg=true)](https://ci.appveyor.com/project/devblok/lipovision/branch/master)|


# Setting up build enviroment

- Set up Git (https://git-scm.com)
- Set up Go (https://golang.org)
- Set up gocv (https://gocv.io)
- Optional: Set up VS Code and Go plugin (https://code.visualstudio.com)

# Contributing/Developing

- Set up mocking tools as described [here](https://github.com/golang/mock)
- Before running tests you need to run `go generate ./...` (do not push generated code - this is not a library)
- Scientific devices can be rather custom and scarce. Meaning, there might be only one person in the world that can maintain parts of it's code 
  (the one that has access to the device)! If it will never see any serial production, it might not be a good idea to add it here! 
  Although, integration with standard tools are always welcome. Let's keep this tidy!
- Someone who's reading your device PR probably can't test it, mock it as much as possible! 
- Avoid any big dependencies (OpenCV is already a big one, and is already annoying to distribute). Best Go dependencies are written in pure Go!
- If you just have to have some cgo dependency, try to have it statically linked!
- Be a good Gopher: Be tidy! Write tests! Use interfaces! Mock interfaces!
