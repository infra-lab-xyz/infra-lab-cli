package config

type ILCConfig struct {
	// TODO: investigate how to implement configs versioning
	// Version     string `mapstructure:"version" default:"1.0"`
	Apps apps `mapstructure:"apps"`
	// TODO: hardcoded now
	// ProjectsDir string `mapstructure:"projects_dir" default:"~/.infra-lab"`
}

type apps struct {
	Docker   dockerConfig `mapstructure:"docker"`
	Helm     helm         `mapstructure:"helm"`
	Kind     kind         `mapstructure:"kind"`
	Minikube minikube     `mapstructure:"minikube"`
	Podman   podman       `mapstructure:"podman"`
	Skopeo   skopeo       `mapstructure:"skopeo"`
	Webhook  webhook      `mapstructure:"webhook"`
}

type skopeo struct {
	Binary string `mapstructure:"binary" default:"skopeo"`
}

type helm struct {
	Binary string `mapstructure:"binary" default:"helm"`
}

type dockerConfig struct {
	Binary     string `mapstructure:"binary" default:"docker"`
	ConfigPath string `mapstructure:"config_path" default:"~/.docker"`
}

type webhook struct {
	Binary       string `mapstructure:"binary" default:"webhook"`
	Secret       string `mapstructure:"secret" default:""`
	ListenAddr   string `mapstructure:"listen_addr" default:"127.0.0.1"`
	ListenPort   int    `mapstructure:"listen_port" default:"8080"`
	WebhooksPath string `mapstructure:"webhooks_path" default:"./hooks.yaml"`
	Prefix       string `mapstructure:"prefix" default:"hooks"`
}

type minikube struct {
	Binary      string `mapstructure:"binary" default:"minikube"`
	ClusterName string `mapstructure:"cluster_name" default:"mini-local"`
	KubeVersion string `mapstructure:"kube_version" default:"v1.30.0"`
	Driver      string `mapstructure:"driver" default:"podman"`
	CNI         string `mapstructure:"cni" default:"auto"`
	CPUs        string `mapstructure:"cpus" default:"2"`
	Memory      string `mapstructure:"memory" default:"2G"`
	DiskSize    string `mapstructure:"disk_size" default:"15G"`
	CIDR        string `mapstructure:"cidr" default:"172.16.0.0/16"`
	NodesCount  int    `mapstructure:"nodes_count" default:"1"`
}

type kind struct {
	Binary      string `mapstructure:"binary" default:"kind"`
	ClusterName string `mapstructure:"cluster_name" default:"kind-local"`
	ConfigPath  string `mapstructure:"config_path" default:""`
}

type podman struct {
	Binary      string `mapstructure:"binary" default:"podman"`
	MachineName string `mapstructure:"machine_name" default:"podman-machine-default"`
	CPUs        string `mapstructure:"cpus" default:"2"`
	Memory      string `mapstructure:"memory" default:"2G"`
	DiskSize    string `mapstructure:"disk_size" default:"10G"`
}
