/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package testutil

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"testing"
)

// TestError holds the different errors for Windows and Linux
type TestError struct {
	WindowsError error
	DefaultError error
}

// GetExpectedError returns the expected error depending on OS
func (t TestError) GetExpectedError() error {
	if isWindows() {
		return t.WindowsError
	}
	return t.DefaultError
}

// AssertError matches the actual and expected errors
func AssertError(actual *TestError, expected error) bool {
	if isWindows() {
		if actual.WindowsError == nil {
			return reflect.DeepEqual(actual.DefaultError, expected)
		}
		return reflect.DeepEqual(actual.WindowsError, expected)
	}
	return reflect.DeepEqual(actual.DefaultError, expected)
}

func IsRunningInProw() bool {
	_, ok := os.LookupEnv("AZURE_CREDENTIALS")
	return ok
}

func isWindows() bool {
	return runtime.GOOS == "windows"
}

// GetWorkDirPath returns the path to the current working directory
func GetWorkDirPath(dir string, t *testing.T) string {
	path, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get working directory: %s", err)
	}
	return fmt.Sprintf("%s%c%s", path, os.PathSeparator, dir)
}
