package consul_client

import (
	"fmt"
	"testing"
)

func TestConsulClient_GetHostInfos(t *testing.T) {
	type fields struct {
		Host  string
		Port  int
		Token string
	}
	f := fields{
		Host:  "127.0.0.1",
		Port:  8500,
		Token: "",
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{name: "AgentGetHostInfos", fields: f, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &ConsulClient{
				Host:  tt.fields.Host,
				Port:  tt.fields.Port,
				Token: tt.fields.Token,
			}
			_, err := client.AgentGetHostInfos()
			if (err != nil) != tt.wantErr {
				t.Errorf("AgentGetHostInfos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestConsulClient_GetMembers(t *testing.T) {
	type fields struct {
		Host  string
		Port  int
		Token string
	}
	f := fields{
		Host:  "127.0.0.1",
		Port:  8500,
		Token: "",
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{name: "AgentGetMembers", fields: f, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &ConsulClient{
				Host:  tt.fields.Host,
				Port:  tt.fields.Port,
				Token: tt.fields.Token,
			}
			got, err := client.AgentGetMembers()
			if (err != nil) != tt.wantErr {
				t.Errorf("AgentGetMembers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}

func TestConsulClient_GetSelfConfig(t *testing.T) {
	type fields struct {
		Host  string
		Port  int
		Token string
		Ssl   bool
	}
	f := fields{
		Host:  "127.0.0.1",
		Port:  8500,
		Token: "",
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{name: "AgentGetSelfConfig", fields: f, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &ConsulClient{
				Host:  tt.fields.Host,
				Port:  tt.fields.Port,
				Token: tt.fields.Token,
				Ssl:   tt.fields.Ssl,
			}
			_, err := client.AgentGetSelfConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("AgentGetSelfConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestConsulClient_AgentReload(t *testing.T) {
	type fields struct {
		Host  string
		Port  int
		Token string
		Ssl   bool
	}
	f := fields{
		Host:  "127.0.0.1",
		Port:  8500,
		Token: "",
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		{name: "AgentReload", fields: f, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &ConsulClient{
				Host:  tt.fields.Host,
				Port:  tt.fields.Port,
				Token: tt.fields.Token,
				Ssl:   tt.fields.Ssl,
			}
			_, err := client.AgentReload()
			if (err != nil) != tt.wantErr {
				t.Errorf("AgentReload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestConsulClient_Maintenance(t *testing.T) {
	type fields struct {
		Host  string
		Port  int
		Token string
		Ssl   bool
	}
	f := fields{
		Host:  "127.0.0.1",
		Port:  8500,
		Token: "",
	}
	type args struct {
		enable bool
		reason string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{name: "AgentMaintenance", fields: f, args: args{
			enable: false,
			reason: "not+necessary",
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &ConsulClient{
				Host:  tt.fields.Host,
				Port:  tt.fields.Port,
				Token: tt.fields.Token,
				Ssl:   tt.fields.Ssl,
			}
			_, err := client.AgentMaintenance(tt.args.enable, tt.args.reason)
			if (err != nil) != tt.wantErr {
				t.Errorf("AgentMaintenance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestConsulClient_Metrics(t *testing.T) {
	type fields struct {
		Host  string
		Port  int
		Token string
		Ssl   bool
	}
	f := fields{
		Host:  "127.0.0.1",
		Port:  8500,
		Token: "",
	}
	type args struct {
		format string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "AgentMetrics", fields: f, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &ConsulClient{
				Host:  tt.fields.Host,
				Port:  tt.fields.Port,
				Token: tt.fields.Token,
				Ssl:   tt.fields.Ssl,
			}
			_, err := client.AgentMetrics(tt.args.format)
			if (err != nil) != tt.wantErr {
				t.Errorf("AgentMetrics() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestConsulClient_MetricsWithFormat(t *testing.T) {
	type fields struct {
		Host  string
		Port  int
		Token string
		Ssl   bool
	}
	f := fields{
		Host:  "127.0.0.1",
		Port:  8500,
		Token: "",
	}
	type args struct {
		format string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "AgentMetricsWithFormat", fields: f, args: args{
			format: "prometheus",
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &ConsulClient{
				Host:  tt.fields.Host,
				Port:  tt.fields.Port,
				Token: tt.fields.Token,
				Ssl:   tt.fields.Ssl,
			}
			_, err := client.AgentMetricsWithFormat(tt.args.format)
			if (err != nil) != tt.wantErr {
				t.Errorf("AgentMetricsWithFormat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestConsulClient_Monitor(t *testing.T) {
	type fields struct {
		Host  string
		Port  int
		Token string
		Ssl   bool
	}
	f := fields{
		Host:  "127.0.0.1",
		Port:  8500,
		Token: "",
	}
	type args struct {
		logJson    bool
		logLevel   string
		logChannel chan []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "AgentMonitor", fields: f, args: args{
			logJson:    true,
			logLevel:   "debug",
			logChannel: make(chan []byte),
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &ConsulClient{
				Host:  tt.fields.Host,
				Port:  tt.fields.Port,
				Token: tt.fields.Token,
				Ssl:   tt.fields.Ssl,
			}
			if err := client.AgentMonitor(tt.args.logJson, tt.args.logLevel, tt.args.logChannel); (err != nil) != tt.wantErr {
				t.Errorf("AgentMonitor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			select {
			case log := <-tt.args.logChannel:
				fmt.Println(log)
			}

		})
	}
}

func TestConsulClient_Join(t *testing.T) {
	type fields struct {
		Host  string
		Port  int
		Token string
		Ssl   bool
	}
	f := fields{
		Host:  "127.0.0.1",
		Port:  8500,
		Token: "",
	}
	type args struct {
		address string
		wan     bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "AgentJoin", fields: f, args: args{
			address: "127.0.0.1:5000",
			wan:     true,
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &ConsulClient{
				Host:  tt.fields.Host,
				Port:  tt.fields.Port,
				Token: tt.fields.Token,
				Ssl:   tt.fields.Ssl,
			}
			_, err := client.AgentJoin(tt.args.address, tt.args.wan)
			if (err != nil) != tt.wantErr {
				t.Errorf("AgentJoin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestConsulClient_Leave(t *testing.T)  {
	type fields struct {
		Host  string
		Port  int
		Token string
		Ssl   bool
	}
	f := fields{
		Host:  "127.0.0.1",
		Port:  8500,
		Token: "",
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{name: "AgentLeave", fields: f, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &ConsulClient{
				Host:  tt.fields.Host,
				Port:  tt.fields.Port,
				Token: tt.fields.Token,
				Ssl:   tt.fields.Ssl,
			}
			_, err := client.AgentLeave()
			if (err != nil) != tt.wantErr {
				t.Errorf("AgentLeave() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestConsulClient_ForceLeave(t *testing.T) {
	type fields struct {
		Host  string
		Port  int
		Token string
		Ssl   bool
	}
	f := fields{
		Host:  "127.0.0.1",
		Port:  8500,
		Token: "",
	}
	type args struct {
		node  string
		prune bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]byte
		wantErr bool
	}{
		{name: "AgentForceLeave", fields: f, args: args{
			node: "gsio",
			prune: true,
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &ConsulClient{
				Host:  tt.fields.Host,
				Port:  tt.fields.Port,
				Token: tt.fields.Token,
				Ssl:   tt.fields.Ssl,
			}
			got, err := client.AgentForceLeave(tt.args.node, tt.args.prune)
			if (err != nil) != tt.wantErr {
				t.Errorf("AgentForceLeave() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}