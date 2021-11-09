package consul_client

import (
	"encoding/json"
	"github.com/jianghaibo12138/go-consul/consul_client/check"
	"testing"
)

func TestConsulClient_CheckRegister(t *testing.T) {
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
		payload check.RegisterPayload
	}
	var c1 args
	bytes := ReadJsonConf("../examples/gin-consul/json_conf/check_register.json")
	err := json.Unmarshal(bytes, &c1.payload)
	if err != nil {
		t.Error(err)
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]byte
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
			got, err := client.CheckRegister(tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckRegister() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}

func TestConsulClient_CheckDeRegister(t *testing.T) {
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
		checkId string
	}
	c1 := args{
		checkId: "gin-consul",
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
			got, err := client.CheckDeRegister(tt.args.checkId)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckDeRegister() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckDeRegister() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConsulClient_CheckList(t *testing.T) {
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
			got, err := client.CheckList(tt.args.filter, tt.args.ns)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}