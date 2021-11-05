package consul_client

import (
	"encoding/json"
	"fmt"
	"github.com/jianghaibo12138/go-consul/consul_client/kv"
	"testing"
)

func TestConsulClient_ReadKey(t *testing.T) {
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
		ns        string
		key       string
		dc        string
		recurse   bool
		raw       bool
		keys      bool
		separator bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]kv.KeyValue
		wantErr bool
	}{
		{name: "ReadKey", fields: f, args: args{
			key: "key",
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
			got, err := client.ReadKey(tt.args.ns, tt.args.key, tt.args.dc, tt.args.recurse, tt.args.raw, tt.args.keys, tt.args.separator)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}

func TestConsulClient_UpsertKey(t *testing.T) {
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
		ns      string
		key     string
		dc      string
		acquire string
		release string
		flags   int
		cas     int
		value   []byte
	}
	jsonByte, _ := json.Marshal(struct {
		Num int
	}{
		Num: 1000,
	})
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{name: "UpsertKey", fields: f, args: args{
			key:   "key",
			value: jsonByte,
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
			got, err := client.UpsertKey(tt.args.ns, tt.args.key, tt.args.dc, tt.args.acquire, tt.args.release, tt.args.flags, tt.args.cas, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpsertKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}

func TestConsulClient_DeleteKey(t *testing.T) {
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
		ns      string
		key     string
		dc      string
		cas     int
		recurse bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{name: "DeleteKey", fields: f, args: args{
			key: "key",
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
			got, err := client.DeleteKey(tt.args.ns, tt.args.key, tt.args.dc, tt.args.cas, tt.args.recurse)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}
