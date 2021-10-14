package consul_client

import (
	"amazing2j.com/go-consul/consul_client/catalog"
	"fmt"
	"reflect"
	"testing"
)

func TestConsulClient_CatalogRegister(t *testing.T) {
	type fields struct {
		Host  string
		Port  int
		Token string
		Ssl   bool
	}
	type args struct {
		namespace string
		cal       *catalog.RegisterCatalog
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &ConsulClient{
				Host:  tt.fields.Host,
				Port:  tt.fields.Port,
				Token: tt.fields.Token,
				Ssl:   tt.fields.Ssl,
			}
			got, err := client.CatalogRegister(tt.args.namespace, tt.args.cal)
			if (err != nil) != tt.wantErr {
				t.Errorf("CatalogRegister() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CatalogRegister() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConsulClient_CatalogDatacenters(t *testing.T) {
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
		want    []string
		wantErr bool
	}{
		{name: "CatalogDatacenters", fields: f, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &ConsulClient{
				Host:  tt.fields.Host,
				Port:  tt.fields.Port,
				Token: tt.fields.Token,
				Ssl:   tt.fields.Ssl,
			}
			got, err := client.CatalogDatacenters()
			if (err != nil) != tt.wantErr {
				t.Errorf("CatalogDatacenters() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}

func TestConsulClient_CatalogNodes(t *testing.T) {
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
		dc     string
		near   string
		filter string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []catalog.Node
		wantErr bool
	}{
		{name: "CatalogNodes", fields: f, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &ConsulClient{
				Host:  tt.fields.Host,
				Port:  tt.fields.Port,
				Token: tt.fields.Token,
				Ssl:   tt.fields.Ssl,
			}
			got, err := client.CatalogNodes(tt.args.dc, tt.args.near, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("CatalogNodes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}

func TestConsulClient_CatalogServices(t *testing.T) {
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
		dc        string
		nodeMeta  string
		namespace string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string][]string
		wantErr bool
	}{
		{name: "CatalogServices", fields: f, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &ConsulClient{
				Host:  tt.fields.Host,
				Port:  tt.fields.Port,
				Token: tt.fields.Token,
				Ssl:   tt.fields.Ssl,
			}
			got, err := client.CatalogServices(tt.args.dc, tt.args.nodeMeta, tt.args.namespace)
			if (err != nil) != tt.wantErr {
				t.Errorf("CatalogServices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}

func TestConsulClient_CatalogServiceNodes(t *testing.T) {
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
		service string
		dc      string
		tag     string
		near    string
		filter  string
		ns      string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []catalog.Node
		wantErr bool
	}{
		{name: "CatalogServiceNodes", fields: f, args: args{
			service: "gsio",
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
			got, err := client.CatalogServiceNodes(tt.args.service, tt.args.dc, tt.args.tag, tt.args.near, tt.args.filter, tt.args.ns)
			if (err != nil) != tt.wantErr {
				t.Errorf("CatalogServiceNodes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}

func TestConsulClient_CatalogConnectNodes(t *testing.T) {
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
		service string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []catalog.Node
		wantErr bool
	}{
		{name: "CatalogConnectNodes", fields: f, args: args{
			service: "gsio",
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
			got, err := client.CatalogConnectNodes(tt.args.service)
			if (err != nil) != tt.wantErr {
				t.Errorf("CatalogConnectNodes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}

func TestConsulClient_CatalogNodeNodes(t *testing.T) {
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
		dc     string
		filter string
		ns     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []catalog.Node
		wantErr bool
	}{
		{name: "CatalogNodeNodes", fields: f, args: args{
			node: "gsio",
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
			got, err := client.CatalogNodeNodes(tt.args.node, tt.args.dc, tt.args.filter, tt.args.ns)
			if (err != nil) != tt.wantErr {
				t.Errorf("CatalogNodeNodes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}

func TestConsulClient_CatalogNodeServices(t *testing.T) {
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
		dc     string
		filter string
		ns     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []catalog.Service
		wantErr bool
	}{
		{name: "CatalogNodeServices", fields: f, args: args{
			node: "gsio",
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
			got, err := client.CatalogNodeServices(tt.args.node, tt.args.dc, tt.args.filter, tt.args.ns)
			if (err != nil) != tt.wantErr {
				t.Errorf("CatalogNodeServices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}