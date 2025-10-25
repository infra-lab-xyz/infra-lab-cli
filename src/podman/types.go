package podman

type ListedMachine struct {
	Name               string `json:"Name"`
	Default            bool   `json:"Default"`
	Created            string `json:"Created"`
	Running            bool   `json:"Running"`
	Starting           bool   `json:"Starting"`
	LastUp             string `json:"LastUp"`
	Stream             string `json:"Stream"`
	VMType             string `json:"VMType"`
	CPUs               int    `json:"CPUs"`
	Memory             string `json:"Memory"`
	DiskSize           string `json:"DiskSize"`
	Port               int    `json:"Port"`
	RemoteUsername     string `json:"RemoteUsername"`
	IdentityPath       string `json:"IdentityPath"`
	UserModeNetworking bool   `json:"UserModeNetworking"`
}

type InspectedMachine struct {
	ConfigDir struct {
		Path string `json:"Path"`
	} `json:"ConfigDir"`
	ConnectionInfo struct {
		PodmanSocket struct {
			Path string `json:"Path"`
		} `json:"PodmanSocket"`
		PodmanPipe interface{} `json:"PodmanPipe"`
	} `json:"ConnectionInfo"`
	Created   string `json:"Created"`
	LastUp    string `json:"LastUp"`
	Name      string `json:"Name"`
	Resources struct {
		CPUs     int           `json:"CPUs"`
		DiskSize int           `json:"DiskSize"`
		Memory   int           `json:"Memory"`
		USBs     []interface{} `json:"USBs"`
	} `json:"Resources"`
	SSHConfig struct {
		IdentityPath   string `json:"IdentityPath"`
		Port           int    `json:"Port"`
		RemoteUsername string `json:"RemoteUsername"`
	} `json:"SSHConfig"`
	State              string `json:"State"`
	UserModeNetworking bool   `json:"UserModeNetworking"`
	Rootful            bool   `json:"Rootful"`
	Rosetta            bool   `json:"Rosetta"`
}

type MachineField struct {
	Name   string
	Header string
	Width  int
}

type ConfigParam struct {
	ValueFlag  int
	Value      int
	IsProvided bool
	IsChanged  bool
}

type ConfigParams struct {
	CPUs     ConfigParam
	Memory   ConfigParam
	DiskSize ConfigParam
}

type Connection struct {
	Name      string `json:"Name"`
	Uri       string `json:"URI"`
	Identity  string `json:"Identity"`
	IsMachine bool   `json:"IsMachine"`
	IsDefault bool   `json:"Default"`
	ReadWrite bool   `json:"ReadWrite"`
}
