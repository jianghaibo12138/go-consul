package consul_client

import (
	"fmt"
	"github.com/jianghaibo12138/go-consul/consul_client/config"
	"reflect"
	"testing"
)

func TestConsulClient_ConfigUpsert(t *testing.T) {
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
		dc       string
		ns       string
		kind     string
		name     string
		protocol string
		cas      int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{name: "ConfigUpsert", fields: f, args: args{
			kind:     "service-defaults",
			name:     "web",
			protocol: "http",
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
			got, err := client.ConfigUpsert(tt.args.dc, tt.args.ns, tt.args.kind, tt.args.name, tt.args.protocol, tt.args.cas)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConfigUpsert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}

func TestConsulClient_ConfigGetter(t *testing.T) {
	type fields struct {
		Host  string
		Port  int
		Token string
		Ssl   bool
	}
	type args struct {
		dc   string
		ns   string
		kind string
		name string
	}
	f := fields{
		Host:  "127.0.0.1",
		Port:  8500,
		Token: "",
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *config.ResponseConfig
		wantErr bool
	}{
		{name: "ConfigGetter", fields: f, args: args{
			kind: "service-defaults",
			name: "web",
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
			got, err := client.ConfigGetter(tt.args.dc, tt.args.ns, tt.args.kind, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConfigGetter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}

func TestConsulClient_ConfigListGetter(t *testing.T) {
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
		dc   string
		ns   string
		kind string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]config.ResponseConfig
		wantErr bool
	}{

		{name: "ConfigListGetter", fields: f, args: args{
			kind: "service-defaults",
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
			got, err := client.ConfigListGetter(tt.args.dc, tt.args.ns, tt.args.kind)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConfigListGetter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}

func TestConsulClient_ConfigDelete(t *testing.T) {
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
		dc   string
		ns   string
		kind string
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want1  *config.ResponseConfig
		want2  error
	}{
		{name: "ConfigListGetter", fields: f, args: args{
			kind: "service-defaults",
			name: "web",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &ConsulClient{
				Host:  tt.fields.Host,
				Port:  tt.fields.Port,
				Token: tt.fields.Token,
				Ssl:   tt.fields.Ssl,
			}
			if got, _ := client.ConfigDelete(tt.args.dc, tt.args.ns, tt.args.kind, tt.args.name); !reflect.DeepEqual(got, tt.want1) {
				t.Errorf("ConfigDelete() = %v, want1 %v, want2 %v", got, tt.want1, tt.want2)
			}
		})
	}
}
