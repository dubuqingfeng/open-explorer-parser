package rpc

import (
	"reflect"
	"testing"
)

func TestNewGrpcClient(t *testing.T) {
	tests := []struct {
		name string
		want *GrpcClient
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGrpcClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGrpcClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
