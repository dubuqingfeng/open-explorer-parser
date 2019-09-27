package rpc

import (
	"reflect"
	"testing"
)

func Test_newRPCClient(t *testing.T) {
	type args struct {
		address string
		user    string
		passwd  string
		ssl     bool
	}
	tests := []struct {
		name string
		args args
		want *RpcClient
	}{
		// TODO: Add test cases.
		{"name", args{address: "127.0.0.1", user: "user", passwd: "password", ssl: false}, &RpcClient{"127.0.0.1", "user", "password", false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newRPCClient(tt.args.address, tt.args.user, tt.args.passwd, tt.args.ssl); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newRPCClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRpcClient_call(t *testing.T) {
	type fields struct {
		address string
		user    string
		passwd  string
		ssl     bool
	}
	type args struct {
		method string
		params interface{}
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantResponse rpcResponse
		wantErr      bool
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
			gotResponse, err := this.call(tt.args.method, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("RpcClient.call() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("RpcClient.call() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}
