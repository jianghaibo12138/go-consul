package agent

type Members struct {
	Name string `json:"Name"`
	Addr string `json:"Addr"`
	Port int    `json:"Port"`
	Tags struct {
		Acls        string `json:"acls"`
		Build       string `json:"build"`
		Dc          string `json:"dc"`
		Expect      string `json:"expect,omitempty"`
		FtFs        string `json:"ft_fs,omitempty"`
		FtSi        string `json:"ft_si,omitempty"`
		Id          string `json:"id"`
		Port        string `json:"port,omitempty"`
		RaftVsn     string `json:"raft_vsn,omitempty"`
		Role        string `json:"role"`
		Segment     string `json:"segment"`
		UseTls      string `json:"use_tls,omitempty"`
		Vsn         string `json:"vsn"`
		VsnMax      string `json:"vsn_max"`
		VsnMin      string `json:"vsn_min"`
		WanJoinPort string `json:"wan_join_port,omitempty"`
	} `json:"Tags"`
	Status      int `json:"Status"`
	ProtocolMin int `json:"ProtocolMin"`
	ProtocolMax int `json:"ProtocolMax"`
	ProtocolCur int `json:"ProtocolCur"`
	DelegateMin int `json:"DelegateMin"`
	DelegateMax int `json:"DelegateMax"`
	DelegateCur int `json:"DelegateCur"`
}
