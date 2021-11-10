package consul_client

import (
	"testing"
)

func TestReadJsonConf(t *testing.T) {
	type args struct {
		filePath string
	}
	c1 := args{
		filePath: "../examples/gin-consul/json_conf/service_register.json",
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "c1", args: c1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReadJsonConf(tt.args.filePath)
			
t.Logf("%+v", got)
		})
	}
}
