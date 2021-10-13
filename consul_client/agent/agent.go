package agent

type Agent struct {
	Memory struct {
		Total          int64   `json:"total"`
		Available      int64   `json:"available"`
		Used           int64   `json:"used"`
		UsedPercent    float64 `json:"usedPercent"`
		Free           int64   `json:"free"`
		Active         int     `json:"active"`
		Inactive       int64   `json:"inactive"`
		Wired          int     `json:"wired"`
		Laundry        int     `json:"laundry"`
		Buffers        int     `json:"buffers"`
		Cached         int64   `json:"cached"`
		WriteBack      int     `json:"writeBack"`
		Dirty          int     `json:"dirty"`
		WriteBackTmp   int     `json:"writeBackTmp"`
		Shared         int     `json:"shared"`
		Slab           int     `json:"slab"`
		Sreclaimable   int     `json:"sreclaimable"`
		Sunreclaim     int     `json:"sunreclaim"`
		PageTables     int     `json:"pageTables"`
		SwapCached     int     `json:"swapCached"`
		CommitLimit    int64   `json:"commitLimit"`
		CommittedAS    int64   `json:"committedAS"`
		HighTotal      int     `json:"highTotal"`
		HighFree       int     `json:"highFree"`
		LowTotal       int     `json:"lowTotal"`
		LowFree        int     `json:"lowFree"`
		SwapTotal      int     `json:"swapTotal"`
		SwapFree       int     `json:"swapFree"`
		Mapped         int     `json:"mapped"`
		VmallocTotal   int64   `json:"vmallocTotal"`
		VmallocUsed    int     `json:"vmallocUsed"`
		VmallocChunk   int     `json:"vmallocChunk"`
		HugePagesTotal int     `json:"hugePagesTotal"`
		HugePagesFree  int     `json:"hugePagesFree"`
		HugePageSize   int     `json:"hugePageSize"`
	} `json:"Memory"`
	CPU []struct {
		Cpu        int      `json:"cpu"`
		VendorId   string   `json:"vendorId"`
		Family     string   `json:"family"`
		Model      string   `json:"model"`
		Stepping   int      `json:"stepping"`
		PhysicalId string   `json:"physicalId"`
		CoreId     string   `json:"coreId"`
		Cores      int      `json:"cores"`
		ModelName  string   `json:"modelName"`
		Mhz        int      `json:"mhz"`
		CacheSize  int      `json:"cacheSize"`
		Flags      []string `json:"flags"`
		Microcode  string   `json:"microcode"`
	} `json:"CPU"`
	Host struct {
		Hostname             string `json:"hostname"`
		Uptime               int    `json:"uptime"`
		BootTime             int    `json:"bootTime"`
		Procs                int    `json:"procs"`
		Os                   string `json:"os"`
		Platform             string `json:"platform"`
		PlatformFamily       string `json:"platformFamily"`
		PlatformVersion      string `json:"platformVersion"`
		KernelVersion        string `json:"kernelVersion"`
		KernelArch           string `json:"kernelArch"`
		VirtualizationSystem string `json:"virtualizationSystem"`
		VirtualizationRole   string `json:"virtualizationRole"`
		HostId               string `json:"hostId"`
	} `json:"Host"`
	Disk struct {
		Path              string  `json:"path"`
		Fstype            string  `json:"fstype"`
		Total             int64   `json:"total"`
		Free              int64   `json:"free"`
		Used              int64   `json:"used"`
		UsedPercent       float64 `json:"usedPercent"`
		InodesTotal       int     `json:"inodesTotal"`
		InodesUsed        int     `json:"inodesUsed"`
		InodesFree        int     `json:"inodesFree"`
		InodesUsedPercent float64 `json:"inodesUsedPercent"`
	} `json:"Disk"`
	CollectionTime int64       `json:"CollectionTime"`
	Errors         interface{} `json:"Errors"`
}
