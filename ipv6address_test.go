package main

import (
	"net"
	"reflect"
	"testing"
)

func TestGetPublicIpv6(t *testing.T) {
	tests := []struct {
		name    string
		want    []*net.IPNet
		wantErr bool
	}{
		{
			name:    "test",
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPublicIpv6()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPublicIpv6() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPublicIpv6() got = %v, want %v", got, tt.want)
			}
		})
	}
}
