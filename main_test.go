package main

import "testing"

func TestGetBuildName(t *testing.T) {
	tests := []struct {
		buildName string
		goos      string
		goarch    string
		expected  string
	}{
		{"my-app", "windows", "amd64", "my-app-windows-amd64.exe"},
		{"my-app", "darwin", "amd64", "my-app-macOS-amd64"},
		{"my-app", "linux", "amd64", "my-app-linux-amd64"},
	}

	for _, tt := range tests {
		actual := getBuildName(tt.buildName, tt.goos, tt.goarch)
		if actual != tt.expected {
			t.Errorf("getBuildName(%q, %q, %q): expected %q, actual %q", tt.buildName, tt.goos, tt.goarch, tt.expected, actual)
		}
	}
}
