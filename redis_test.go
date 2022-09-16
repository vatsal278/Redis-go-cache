package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

func TestHealth(t *testing.T) {
	tests := []struct {
		name         string
		requestBody  string
		validateFunc func(string, error)
	}{
		{
			name:        "Success::Health",
			requestBody: "localhost:9096",
			validateFunc: func(s string, err error) {
				if err != nil {
					t.Errorf("want %v got %v", nil, err.Error())
				}
				if s != "PONG" {
					t.Errorf("want %v got %v", "PONG", s)
				}
			},
		},
		{
			name:        "Failure:: Health",
			requestBody: "",
			validateFunc: func(s string, err error) {
				if err == nil {
					t.Errorf("want %v got %v", "not nil", nil)
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cacher := NewCacher(Config{
				Addr: tt.requestBody,
			})
			x, err := cacher.Health()
			tt.validateFunc(x, err)
		})
	}

}
func TestSet(t *testing.T) {

	tests := []struct {
		name         string
		requestBody  string
		expiry       time.Duration
		validateFunc func(Cacher, string, error)
	}{
		{
			name:        "Success:: Set",
			requestBody: "localhost:9096",
			validateFunc: func(cacher Cacher, data string, err error) {
				if err != nil {
					t.Errorf("want %v got %v", nil, err.Error())
				}
				x, err := cacher.Get("1")
				if err != nil {
					t.Errorf("want %v got %v", nil, err.Error())
				}
				if fmt.Sprintf("%s", x) != data {
					t.Errorf("want %v got %v", data, fmt.Sprintf("%s", x))
				}
			},
		},
		{
			name:        "Success:: Set:: With Expiry",
			requestBody: "localhost:9096",
			expiry:      2 * time.Second,
			validateFunc: func(cacher Cacher, data string, err error) {
				if err != nil {
					t.Errorf("want %v got %v", nil, err.Error())
				}
				time.Sleep(2 * time.Second)
				_, err = cacher.Get("1")
				if err != redis.Nil {
					t.Errorf("want %v got %v", redis.Nil, err)
				}
			},
		},
		{
			name:        "Failure:: Set",
			requestBody: "",
			validateFunc: func(cacher Cacher, data string, err error) {
				if err == nil {
					t.Errorf("want %v got %v", "not nil", nil)
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cacher := NewCacher(Config{
				Addr: tt.requestBody,
			})
			data := tt.requestBody
			err := cacher.Set("1", data, tt.expiry)
			tt.validateFunc(cacher, data, err)
		})
	}

}

func TestDelete(t *testing.T) {
	cacher := NewCacher(Config{
		Addr: "localhost:9096",
	})
	tests := []struct {
		name         string
		requestBody  string
		setupFunc    func(string)
		validateFunc func(error)
	}{
		{
			name:        "Success:: Delete",
			requestBody: "1",
			setupFunc: func(data string) {
				err := cacher.Set("1", data, 0)
				if err != nil {
					t.Errorf("want %v got %v", nil, err.Error())
				}
			},
			validateFunc: func(err error) {
				if err != nil {
					t.Errorf("want %v got %v", nil, err)
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := tt.requestBody
			tt.setupFunc("Hello")
			err := cacher.Delete(key)
			tt.validateFunc(err)
		})
	}

}

func TestGet(t *testing.T) {
	cacher := NewCacher(Config{
		Addr: "localhost:9096",
	})
	tests := []struct {
		name         string
		requestBody  string
		setupFunc    func(string)
		validateFunc func([]byte, string, error)
	}{
		{
			name:        "Success:: Get",
			requestBody: "1",
			setupFunc: func(data string) {
				err := cacher.Set("1", data, 0)
				if err != nil {
					t.Errorf("want %v got %v", nil, err.Error())
				}
			},
			validateFunc: func(s []byte, request string, err error) {
				if err != nil {
					t.Errorf("want %v got %v", nil, err.Error())
				}
				if fmt.Sprintf("%s", s) != request {
					t.Errorf("want %v got %v", request, fmt.Sprintf("%s", s))
				}
			},
		},
		{
			name:        "Failure:: Get",
			requestBody: "2",
			setupFunc: func(data string) {

			},
			validateFunc: func(s []byte, request string, err error) {
				if err != redis.Nil {
					t.Errorf("want %v got %v", redis.Nil, err)
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := tt.requestBody
			tt.setupFunc("Hello")
			x, err := cacher.Get(key)
			tt.validateFunc(x, "Hello", err)
		})
	}

}
