package consul_client

import (
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