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
		{name: "GetHostInfos", fields: f, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &ConsulClient{
				Host:  tt.fields.Host,
				Port:  tt.fields.Port,
				Token: tt.fields.Token,
			}
			_, err := client.GetHostInfos()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHostInfos() error = %v, wantErr %v", err, tt.wantErr)
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
		{name: "GetMembers", fields: f, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &ConsulClient{
				Host:  tt.fields.Host,
				Port:  tt.fields.Port,
				Token: tt.fields.Token,
			}
			_, err := client.GetMembers()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMembers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
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
		{name: "GetSelfConfig", fields: f, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &ConsulClient{
				Host:  tt.fields.Host,
				Port:  tt.fields.Port,
				Token: tt.fields.Token,
				Ssl:   tt.fields.Ssl,
			}
			_, err := client.GetSelfConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSelfConfig() error = %v, wantErr %v", err, tt.wantErr)
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
		{name: "Maintenance", fields: f, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &ConsulClient{
				Host:  tt.fields.Host,
				Port:  tt.fields.Port,
				Token: tt.fields.Token,
				Ssl:   tt.fields.Ssl,
			}
			_, err := client.Maintenance(tt.args.enable, tt.args.reason)
			if (err != nil) != tt.wantErr {
				t.Errorf("Maintenance() error = %v, wantErr %v", err, tt.wantErr)
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
		{name: "Metrics", fields: f, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &ConsulClient{
				Host:  tt.fields.Host,
				Port:  tt.fields.Port,
				Token: tt.fields.Token,
				Ssl:   tt.fields.Ssl,
			}
			_, err := client.Metrics(tt.args.format)
			if (err != nil) != tt.wantErr {
				t.Errorf("Metrics() error = %v, wantErr %v", err, tt.wantErr)
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
		{name: "MetricsWithFormat", fields: f, args: args{
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
			_, err := client.MetricsWithFormat(tt.args.format)
			if (err != nil) != tt.wantErr {
				t.Errorf("MetricsWithFormat() error = %v, wantErr %v", err, tt.wantErr)
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
		{name: "Monitor", fields: f, args: args{
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
			if err := client.Monitor(tt.args.logJson, tt.args.logLevel, tt.args.logChannel); (err != nil) != tt.wantErr {
				t.Errorf("Monitor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			select {
			case log := <-tt.args.logChannel:
				fmt.Println(log)
			}

		})
	}
}
