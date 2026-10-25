package config

type ApplicationArtifact struct {
	// Currently supported: github-release
	Type string `yaml:"type"`
	// Githubgg repository username/repo
	Repository string `yaml:"repository"`
	// What release to target
	Selector string `yaml:"selector"`
	// Environment variable where to pick up API token from
	TokenEnv string `yaml:"token_env"`
}

type Deployment struct {
	// Path where to deploy
	Path string `yaml:"path"`
	// How many releases to retain on the server (1-20)
	Retain int `yaml:"retain"`
}

type Application struct {
	// Application name
	Name string `yaml:"name"`
	// Artifact information
	Artifact ApplicationArtifact `yaml:"artifact"`
}

type Config struct {
	// Configuration version
	Version int `yaml:"version"`
	// List of applications to be deployed
	Applications []Application `yaml:"applications"`
}
