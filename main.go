package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"gopkg.in/yaml.v2"
)

type BuildConfig struct {
	Goos   string `yaml:"goos"`
	Goarch string `yaml:"goarch"`
}

type Config struct {
	SourceDir string        `yaml:"source_dir"`
	BuildName string        `yaml:"build_name"`
	Builds    []BuildConfig `yaml:"builds"`
}

// loadConfig reads the YAML build configuration.
// It defaults to the current directory if source_dir is not specified.
func loadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	if config.SourceDir == "" {
		config.SourceDir = "."
	}

	return &config, nil
}

// getBuildName creates a platform-specific binary name.
// For example, "my-app" on Windows (amd64) becomes "my-app-windows-amd64.exe".
func getBuildName(buildName, goos, goarch string) string {
	switch goos {
	case "windows":
		return fmt.Sprintf("%s-%s-%s.exe", buildName, "windows", goarch)
	case "darwin":
		return fmt.Sprintf("%s-%s-%s", buildName, "macOS", goarch)
	default:
		return fmt.Sprintf("%s-%s-%s", buildName, goos, goarch)
	}
}

// buildBinary executes 'go build' for a specific platform.
func buildBinary(config *Config, build BuildConfig) error {
	buildName := getBuildName(config.BuildName, build.Goos, build.Goarch)
	cmd := exec.Command("go", "build", "-o", fmt.Sprintf("builds/%s", buildName), config.SourceDir)
	cmd.Env = append(os.Environ(), "GOOS="+build.Goos, "GOARCH="+build.Goarch)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to build for %s/%s: %s %s", build.Goos, build.Goarch, err, output)
	}

	fmt.Printf("Successfully built for %s/%s", build.Goos, build.Goarch)
	return nil
}

// main is the entry point for the application.
// It expects a single argument: the path to a YAML config file.
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: gobuild <config_file.yaml>")
		return
	}

	config, err := loadConfig(os.Args[1])
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for _, build := range config.Builds {
		if err := buildBinary(config, build); err != nil {
			log.Fatalf("Error: %v", err)
		}
	}
}
