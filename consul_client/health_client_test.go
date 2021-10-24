package consul_client

import (
	"fmt"
	"jianghaibo12138/go-consul/consul_client/health"
	"testing"
)

func TestConsulClient_NodeHealthStatus(t *testing.T) {
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
		node   string
		ns     string
		dc     string
		filter string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []health.NodeStatus
		wantErr bool
	}{
		{name: "NodeHealthStatus", fields: f, args: args{
			node:   "my-node",
			ns:     "",
			dc:     "",
			filter: "",
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
			got, err := client.NodeHealthStatus(tt.args.node, tt.args.ns, tt.args.dc, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("NodeHealthStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(fmt.Sprintf("%+v", got))
		})
	}
}

func TestConsulClient_ServiceHealthStatus(t *testing.T) {
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
		service  string
		ns       string
		dc       string
		filter   string
		near     string
		nodeMeta string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []health.ServiceStatus
		wantErr bool
	}{
		{name: "ServiceHealthStatus", fields: f, args: args{
			service:  "my-service",
			ns:       "",
			dc:       "",
			filter:   "",
			near:     "",
			nodeMeta: "",
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
			got, err := client.ServiceHealthStatus(tt.args.service, tt.args.ns, tt.args.dc, tt.args.filter, tt.args.near, tt.args.nodeMeta)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServiceHealthStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(fmt.Sprintf("%+v", got))
		})
	}
}

func TestConsulClient_ServiceInstances(t *testing.T) {
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
		service  string
		ns       string
		dc       string
		filter   string
		near     string
		tag      string
		nodeMeta string
		passing  bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []health.ServiceInstance
		wantErr bool
	}{
		{name: "ServiceInstances", fields: f, args: args{
			service:  "my-service",
			ns:       "",
			dc:       "",
			filter:   "",
			tag:      "",
			nodeMeta: "",
			passing:  false,
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
			got, err := client.ServiceInstances(tt.args.service, tt.args.ns, tt.args.dc, tt.args.filter, tt.args.near, tt.args.tag, tt.args.nodeMeta, tt.args.passing)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServiceInstances() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(fmt.Sprintf("%+v", got))
		})
	}
}
