package http_client

import (
	"testing"
)

func TestHttpClient_RequestStream(t *testing.T) {
	type fields struct {
		Method      string
		Url         string
		ContentType string
		Headers     map[string]string
	}
	f := fields{
		Method: "GET",
		Url: "http://127.0.0.1:8500",
		ContentType: "application/json",
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "RequestStream", fields: f, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &HttpClient{
				Method:      tt.fields.Method,
				Url:         tt.fields.Url,
				ContentType: tt.fields.ContentType,
				Headers:     tt.fields.Headers,
			}
			_, err := client.RequestStream(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("RequestStream() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
