package main

import (
	"reflect"
	"testing"
)

func TestMdHasher(t *testing.T) {
	type urls []string
	tests := []struct {
		name    string
		args    urls
		want    *Results
		wantErr bool
	}{
		{"url format is wrong", urls{"goo.goo"}, nil, true},
		{"success case", urls{"http://example.com"},
			&Results{"http://example.com", [16]byte{132, 35, 141, 252, 128, 146, 229, 217, 192, 218, 200, 239, 147, 55, 26, 7}}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MdHasher(tt.args[0])
			if (err != nil) != tt.wantErr {
				t.Errorf("Process() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Process() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMdHasherHandler(t *testing.T) {
	done := make(chan struct{})

	tests := []struct {
		name string
		urls []string
		want *Results
	}{
		{"run goroutine succesful test", []string{"http://example.com"},
			&Results{"http://example.com", [16]byte{132, 35, 141, 252, 128, 146, 229, 217, 192, 218, 200, 239, 147, 55, 26, 7}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := MdHasherHandler(tt.urls, DefaultGoroutines, done)
			defer close(result)
			if got := <-result; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("result are different. = %v, want %v", got, tt.want)
			}
		})
	}
}
