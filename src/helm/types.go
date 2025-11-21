package helm

type Chart struct {
	Name        string `json:"name" yaml:"name"`
	Version     string `json:"version" yaml:"version"`
	AppVersion  string `json:"app_version" yaml:"app_version"`
	Description string `json:"description" yaml:"description"`
}

type OCIChartVersions struct {
	Url  string   `json:"-"`
	Tags []string `json:"tags"`
}

type HelmRepo struct {
	Name string `json:"name" yaml:"name"`
	Url  string `json:"url" yaml:"url"`
	Type string `json:"-" yaml:"-"`
}

type OCIRepos struct {
	Repos []HelmRepo `mapstructure:"repos"`
}
