package agent

type Metrics struct {
	Timestamp string `json:"Timestamp"`
	Gauges    []struct {
		Name   string `json:"Name"`
		Value  int    `json:"Value"`
		Labels struct {
			Datacenter string `json:"datacenter,omitempty"`
			Segment    string `json:"segment,omitempty"`
		} `json:"Labels"`
	} `json:"Gauges"`
	Points   []interface{} `json:"Points"`
	Counters []struct {
		Name   string  `json:"Name"`
		Count  int     `json:"Count"`
		Rate   float64 `json:"Rate"`
		Sum    int     `json:"Sum"`
		Min    int     `json:"Min"`
		Max    int     `json:"Max"`
		Mean   float64 `json:"Mean"`
		Stddev float64 `json:"Stddev"`
		Labels struct {
		} `json:"Labels"`
	} `json:"Counters"`
	Samples []struct {
		Name   string  `json:"Name"`
		Count  int     `json:"Count"`
		Rate   float64 `json:"Rate"`
		Sum    float64 `json:"Sum"`
		Min    float64 `json:"Min"`
		Max    float64 `json:"Max"`
		Mean   float64 `json:"Mean"`
		Stddev float64 `json:"Stddev"`
		Labels struct {
		} `json:"Labels"`
	} `json:"Samples"`
}
