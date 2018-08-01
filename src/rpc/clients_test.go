package rpc

import (
	"reflect"
	"testing"
)

func TestNewClients(t *testing.T) {
	tests := []struct {
		name string
		want []*RpcClient
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClients(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClients() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRpcClient_Call(t *testing.T) {
	type fields struct {
		address string
		user    string
		passwd  string
		ssl     bool
	}
	type args struct {
		rpc RpcClient
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantMessage string
		wantErr     bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &RpcClient{
				address: tt.fields.address,
				user:    tt.fields.user,
				passwd:  tt.fields.passwd,
				ssl:     tt.fields.ssl,
			}
			gotMessage, err := this.Call(tt.args.rpc)
			if (err != nil) != tt.wantErr {
				t.Errorf("RpcClient.Call() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotMessage != tt.wantMessage {
				t.Errorf("RpcClient.Call() = %v, want %v", gotMessage, tt.wantMessage)
			}
		})
	}
}
