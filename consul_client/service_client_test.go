package consul_client

import (
	"encoding/json"
	"github.com/jianghaibo12138/go-consul/consul_client/service"
	"testing"
)

func TestConsulClient_RegisterService(t *testing.T) {
	type fields struct {
		Host  string
		Port  int
		Token string
		Ssl   bool
	}
	f := fields{
		Host:  "127.0.0.1",
		Port:  8500,
		Token: "e95b599e-166e-7d80-08ad-aee76e7ddf19",
	}
	type args struct {
		srv service.Payload
	}
	var c1 args
	bytes := ReadJsonConf("../examples/gin-consul/json_conf/service_register.json")
	err := json.Unmarshal(bytes, &c1.srv)
	if err != nil {
		t.Error(err)
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{name: "c1", fields: f, args: c1, want: true, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &ConsulClient{
				Host:  tt.fields.Host,
				Port:  tt.fields.Port,
				Token: tt.fields.Token,
				Ssl:   tt.fields.Ssl,
			}
			got, err := client.RegisterService(tt.args.srv)
			if (err != nil) != tt.wantErr {
				t.Errorf("Payload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}

func TestConsulClient_ServiceList(t *testing.T) {
	type fields struct {
		Host  string
		Port  int
		Token string
		Ssl   bool
	}
	f := fields{
		Host:  "127.0.0.1",
		Port:  8500,
		Token: "e95b599e-166e-7d80-08ad-aee76e7ddf19",
	}
	type args struct {
		filter string
		ns     string
	}
	c1 := args{}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{name: "c1", fields: f, args: c1, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &ConsulClient{
				Host:  tt.fields.Host,
				Port:  tt.fields.Port,
				Token: tt.fields.Token,
				Ssl:   tt.fields.Ssl,
			}
			got, err := client.ServiceList(tt.args.filter, tt.args.ns)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServiceList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}

func TestConsulClient_GetServiceConfiguration(t *testing.T) {
	type fields struct {
		Host  string
		Port  int
		Token string
		Ssl   bool
	}
	f := fields{
		Host:  "127.0.0.1",
		Port:  8500,
		Token: "e95b599e-166e-7d80-08ad-aee76e7ddf19",
	}
	type args struct {
		serviceId string
		ns        string
	}
	c1 := args{
		serviceId: "stp-go",
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *service.Configuration
		wantErr bool
	}{
		{name: "c1", fields: f, args: c1, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &ConsulClient{
				Host:  tt.fields.Host,
				Port:  tt.fields.Port,
				Token: tt.fields.Token,
				Ssl:   tt.fields.Ssl,
			}
			got, err := client.GetServiceConfiguration(tt.args.serviceId, tt.args.ns)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetServiceConfiguration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}

func TestConsulClient_GetServiceHealthByName(t *testing.T) {
	type fields struct {
		Host  string
		Port  int
		Token string
		Ssl   bool
	}
	f := fields{
		Host:  "127.0.0.1",
		Port:  8500,
		Token: "e95b599e-166e-7d80-08ad-aee76e7ddf19",
	}
	type args struct {
		serviceName string
		ns          string
	}
	c1 := args{
		serviceName: "StpGo",
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]service.Health
		wantErr bool
	}{
		{name: "c1", fields: f, args: c1, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &ConsulClient{
				Host:  tt.fields.Host,
				Port:  tt.fields.Port,
				Token: tt.fields.Token,
				Ssl:   tt.fields.Ssl,
			}
			got, err := client.GetServiceHealthByName(tt.args.serviceName, tt.args.ns)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetServiceHealthByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}

func TestConsulClient_DeRegisterService(t *testing.T) {
	type fields struct {
		Host  string
		Port  int
		Token string
		Ssl   bool
	}
	f := fields{
		Host:  "127.0.0.1",
		Port:  8500,
		Token: "e95b599e-166e-7d80-08ad-aee76e7ddf19",
	}
	type args struct {
		serviceId string
	}
	c1 := args{
		serviceId: "stp-go",
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{name: "c1", fields: f, args: c1, want: true, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &ConsulClient{
				Host:  tt.fields.Host,
				Port:  tt.fields.Port,
				Token: tt.fields.Token,
				Ssl:   tt.fields.Ssl,
			}
			got, err := client.DeRegisterService(tt.args.serviceId)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeRegisterService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DeRegisterService() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConsulClient_GetServiceHealthById(t *testing.T) {
	type fields struct {
		Host  string
		Port  int
		Token string
		Ssl   bool
	}
	f := fields{
		Host:  "127.0.0.1",
		Port:  8500,
		Token: "e95b599e-166e-7d80-08ad-aee76e7ddf19",
	}
	type args struct {
		serviceId string
		ns        string
	}
	c1 := args{
		serviceId: "stp-go",
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *service.Health
		wantErr bool
	}{
		{name: "c1", fields: f, args: c1, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &ConsulClient{
				Host:  tt.fields.Host,
				Port:  tt.fields.Port,
				Token: tt.fields.Token,
				Ssl:   tt.fields.Ssl,
			}
			got, err := client.GetServiceHealthById(tt.args.serviceId, tt.args.ns)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetServiceHealthById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}
