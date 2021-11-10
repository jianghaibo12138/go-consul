package consul_client

import (
	"github.com/jianghaibo12138/go-consul/consul_client/connect"
	"testing"
)

func TestConsulClient_GetConnectCA(t *testing.T) {
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
	tests := []struct {
		name    string
		fields  fields
		want    bool
		wantErr bool
	}{
		{name: "c1", fields: f, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &ConsulClient{
				Host:  tt.fields.Host,
				Port:  tt.fields.Port,
				Token: tt.fields.Token,
				Ssl:   tt.fields.Ssl,
			}
			got, err := client.GetConnectCARoots()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetConnectCARoots() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			t.Logf("%+v", got)
		})
	}
}

func TestConsulClient_GetConnectCAConfiguration(t *testing.T) {
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
	tests := []struct {
		name    string
		fields  fields
		want    *connect.CAConfiguration
		wantErr bool
	}{
		{name: "c1", fields: f, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &ConsulClient{
				Host:  tt.fields.Host,
				Port:  tt.fields.Port,
				Token: tt.fields.Token,
				Ssl:   tt.fields.Ssl,
			}
			got, err := client.GetConnectCAConfiguration()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetConnectCAConfiguration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("%+v", got)
		})
	}
}

func TestConsulClient_UpdateConnectCAConfiguration(t *testing.T) {
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
		config *connect.CAConfiguration
	}
	c1 := args{
		config: &connect.CAConfiguration{
			Provider: "consul",
			Config: connect.ProviderConfig{
				LeafCertTTL:         "72h",
				PrivateKey:          "-----BEGIN CERTIFICATE-----\nMIICDzCCAbWgAwIBAgIBCDAKBggqhkjOPQQDAjAxMS8wLQYDVQQDEyZwcmktMWNq\nOHphbW0uY29uc3VsLmNhLjA3OTMzYTEzLmNvbnN1bDAeFw0yMDEwMDgxOTQ4MzZa\nFw0zMDEwMDgxOTQ4MzZaMDExLzAtBgNVBAMTJnByaS0xY2o4emFtbS5jb25zdWwu\nY2EuMDc5MzNhMTMuY29uc3VsMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEDCkT\nIIxSDhA3XCKIuDcj4s9IVjf0NQT6QHPAzFBb964/4fTtX/J8x2n6A1lOXowFIWtx\nGvAD/IJF74zn5ZA/wqOBvTCBujAOBgNVHQ8BAf8EBAMCAYYwDwYDVR0TAQH/BAUw\nAwEB/zApBgNVHQ4EIgQguPlAkrIkOnLr9+8DZ4afZWrYZUd2LB6nMJP72jDVxmcw\nKwYDVR0jBCQwIoAguPlAkrIkOnLr9+8DZ4afZWrYZUd2LB6nMJP72jDVxmcwPwYD\nVR0RBDgwNoY0c3BpZmZlOi8vMDc5MzNhMTMtYTYyYi1iZTkwLTQ0ZjEtZGVkOWE2\nNjczNzZlLmNvbnN1bDAKBggqhkjOPQQDAgNIADBFAiEA0ExkvLESG1I1TMFVronr\n2fjoORukgzBgRMbWAEC2DJ0CIACsxeFS6tprHiRv4cEa2Md75h1iIisb2V2U7dvY\nZ7Rr\n-----END CERTIFICATE-----",
				RootCert:            "-----BEGIN CERTIFICATE-----\nMIICEzCCAbigAwIBAgIBCTAKBggqhkjOPQQDAjAxMS8wLQYDVQQDEyZwcmktMWNq\nOHphbW0uY29uc3VsLmNhLjA3OTMzYTEzLmNvbnN1bDAeFw0yMDEwMDgxOTQ3Mzda\nFw0yMTEwMDgxOTQ3MzdaMDExLzAtBgNVBAMTJnNlYy0xbmIxMHZ0by5jb25zdWwu\nY2EuMDc5MzNhMTMuY29uc3VsMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE9zWs\nUxEYvLZUySoflz6e+HqLcaXM8heNRRkAiLiGkmn6nan6olnnrVBLyHAfHaHWJQ9W\nwI8HwSZf0g4Ms16LWKOBwDCBvTAOBgNVHQ8BAf8EBAMCAYYwEgYDVR0TAQH/BAgw\nBgEB/wIBADApBgNVHQ4EIgQg+csK9Sg6odIfLLk3aiRY2OB4O0DiOa1XRTVdOVDE\nt6QwKwYDVR0jBCQwIoAguPlAkrIkOnLr9+8DZ4afZWrYZUd2LB6nMJP72jDVxmcw\nPwYDVR0RBDgwNoY0c3BpZmZlOi8vMDc5MzNhMTMtYTYyYi1iZTkwLTQ0ZjEtZGVk\nOWE2NjczNzZlLmNvbnN1bDAKBggqhkjOPQQDAgNJADBGAiEAqJ60KJepAP4Xe4Ak\n5UYB1huu/B8Lyz5yEYUpUplgdD4CIQCrrkoXoD4SGJ4HaIjy6a5eNf3YkhLpmbXO\n6DL6FXVa1Q==\n-----END CERTIFICATE-----",
				IntermediateCertTTL: "8760h",
			},
			CreateIndex: 5,
			ModifyIndex: 5,
		},
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
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
			if err := client.UpdateConnectCAConfiguration(tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("UpdateConnectCAConfiguration() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConsulClient_UpsertIntentionByName(t *testing.T) {
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
		sourceService      string
		destinationService string
		payload            *connect.IntentionPayload
	}
	c1 := args{
		sourceService:      "web",
		destinationService: "db",
		payload: &connect.IntentionPayload{
			SourceType:  "consul",
			Action:      "allow",
			Description: "",
		},
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
			got, err := client.UpsertIntentionByName(tt.args.sourceService, tt.args.destinationService, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpsertIntentionByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UpsertIntentionByName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConsulClient_CreateIntentionWithId(t *testing.T) {
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
		payload *connect.IntentionPayload
	}
	c1 := args{
		payload: &connect.IntentionPayload{
			SourceName:      "web1",
			DestinationName: "db1",
			SourceType:      "consul1",
			Action:          "allow",
			Description:     "",
		},
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *connect.IntentionResponse
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
			got, err := client.CreateIntentionWithId(tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateIntentionWithId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("%+v", got)
		})
	}
}

func TestConsulClient_UpdateIntentionWithId(t *testing.T) {
	type fields struct {
		Host  string
		Port  int
		Token string
		Ssl   bool
	}
	type args struct {
		id      string
		payload *connect.IntentionPayload
	}
	f := fields{
		Host:  "127.0.0.1",
		Port:  8500,
		Token: "e95b599e-166e-7d80-08ad-aee76e7ddf19",
	}
	c1 := args{
		id: "8f246b77-f3e1-ff88-5b48-8ec93abf3e05",
		payload: &connect.IntentionPayload{
			SourceName:      "web",
			DestinationName: "db",
			SourceType:      "consul",
			Action:          "allow",
			Description:     "",
		},
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *connect.IntentionResponse
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
			got, err := client.UpdateIntentionById(tt.args.id, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateIntentionById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("%+v", got)
		})
	}
}

func TestConsulClient_GetIntentionByName(t *testing.T) {
	type fields struct {
		Host  string
		Port  int
		Token string
		Ssl   bool
	}
	type args struct {
		source      string
		destination string
	}
	f := fields{
		Host:  "127.0.0.1",
		Port:  8500,
		Token: "e95b599e-166e-7d80-08ad-aee76e7ddf19",
	}
	c1 := args{
		source:      "web",
		destination: "db",
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *connect.IntentionPayload
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
			got, err := client.GetIntentionByName(tt.args.source, tt.args.destination)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIntentionByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("%+v", got)
		})
	}
}

func TestConsulClient_GetIntentionById(t *testing.T) {
	type fields struct {
		Host  string
		Port  int
		Token string
		Ssl   bool
	}
	type args struct {
		id string
	}
	f := fields{
		Host:  "127.0.0.1",
		Port:  8500,
		Token: "e95b599e-166e-7d80-08ad-aee76e7ddf19",
	}
	c1 := args{
		id: "8f246b77-f3e1-ff88-5b48-8ec93abf3e05",
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *connect.IntentionPayload
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
			got, err := client.GetIntentionById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIntentionById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("%+v", got)
		})
	}
}

func TestConsulClient_GetIntentionList(t *testing.T) {
	type fields struct {
		Host  string
		Port  int
		Token string
		Ssl   bool
	}
	type args struct {
		filter string
	}
	f := fields{
		Host:  "127.0.0.1",
		Port:  8500,
		Token: "e95b599e-166e-7d80-08ad-aee76e7ddf19",
	}
	c1 := args{
		filter: "",
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *connect.IntentionPayload
		wantErr bool
	}{
		{name: "c1", args: c1, fields: f, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &ConsulClient{
				Host:  tt.fields.Host,
				Port:  tt.fields.Port,
				Token: tt.fields.Token,
				Ssl:   tt.fields.Ssl,
			}
			got, err := client.GetIntentionList(tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIntentionList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("%+v", got)
		})
	}
}

func TestConsulClient_DeleteIntentionByName(t *testing.T) {
	type fields struct {
		Host  string
		Port  int
		Token string
		Ssl   bool
	}
	type args struct {
		source      string
		destination string
	}
	f := fields{
		Host:  "127.0.0.1",
		Port:  8500,
		Token: "e95b599e-166e-7d80-08ad-aee76e7ddf19",
	}
	c1 := args{
		source:      "web1",
		destination: "db1",
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "c1", args: c1, fields: f, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &ConsulClient{
				Host:  tt.fields.Host,
				Port:  tt.fields.Port,
				Token: tt.fields.Token,
				Ssl:   tt.fields.Ssl,
			}
			err := client.DeleteIntentionByName(tt.args.source, tt.args.destination)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteIntentionByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func TestConsulClient_DeleteIntentionById(t *testing.T) {
	type fields struct {
		Host  string
		Port  int
		Token string
		Ssl   bool
	}
	type args struct {
		id string
	}
	f := fields{
		Host:  "127.0.0.1",
		Port:  8500,
		Token: "e95b599e-166e-7d80-08ad-aee76e7ddf19",
	}
	c1 := args{
		id: "web1",
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "c1", args: c1, fields: f, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &ConsulClient{
				Host:  tt.fields.Host,
				Port:  tt.fields.Port,
				Token: tt.fields.Token,
				Ssl:   tt.fields.Ssl,
			}
			if err := client.DeleteIntentionById(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteIntentionById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConsulClient_CheckIntentionResult(t *testing.T) {
	type fields struct {
		Host  string
		Port  int
		Token string
		Ssl   bool
	}
	type args struct {
		source      string
		destination string
	}
	f := fields{
		Host:  "127.0.0.1",
		Port:  8500,
		Token: "e95b599e-166e-7d80-08ad-aee76e7ddf19",
	}
	c1 := args{
		source:      "web",
		destination: "db",
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *connect.IntentionResponse
		wantErr bool
	}{
		{name: "c1", args: c1, fields: f, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &ConsulClient{
				Host:  tt.fields.Host,
				Port:  tt.fields.Port,
				Token: tt.fields.Token,
				Ssl:   tt.fields.Ssl,
			}
			got, err := client.CheckIntentionResult(tt.args.source, tt.args.destination)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckIntentionResult() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("%+v", got)
		})
	}
}

func TestConsulClient_GetMatchingIntentionList(t *testing.T) {
	type fields struct {
		Host  string
		Port  int
		Token string
		Ssl   bool
	}
	type args struct {
		by   string
		name string
	}
	f := fields{
		Host:  "127.0.0.1",
		Port:  8500,
		Token: "e95b599e-166e-7d80-08ad-aee76e7ddf19",
	}
	c1 := args{
		by:      "web",
		name: "db",
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *map[string][]connect.IntentionPayload
		wantErr bool
	}{
		{name: "c1", args: c1, fields: f, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &ConsulClient{
				Host:  tt.fields.Host,
				Port:  tt.fields.Port,
				Token: tt.fields.Token,
				Ssl:   tt.fields.Ssl,
			}
			got, err := client.GetMatchingIntentionList(tt.args.by, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMatchingIntentionList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("%+v", got)
		})
	}
}