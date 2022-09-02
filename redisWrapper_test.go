package redis

import (
	"fmt"
	"testing"
)

func TestHealth(t *testing.T) {
	tests := []struct {
		name         string
		requestBody  string
		validateFunc func(string, error)
	}{
		{
			name:        "Success:: Register Publisher",
			requestBody: "localhost:9096",
			validateFunc: func(s string, err error) {
				if err != nil {
					t.Log(err.Error())
				}
				if s != "PONG" {
					t.Log(s)
					t.Fail()
				}
			},
		},
		{
			name:        "Failure:: Health",
			requestBody: "",
			validateFunc: func(s string, err error) {
				if err == nil {
					t.Log(err.Error())
					t.Fail()
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cacher := NewCacher(Config{
				Addr:     tt.requestBody,
				Password: "",
				DB:       0,
			})
			x, err := cacher.Health()
			tt.validateFunc(x, err)
		})
	}

}
func TestSet(t *testing.T) {
	cacher := NewCacher(Config{
		Addr:     "localhost:9096",
		Password: "",
		DB:       0,
	})
	tests := []struct {
		name         string
		requestBody  string
		validateFunc func(string, error)
	}{
		{
			name:        "Success:: Set",
			requestBody: "Hello",
			validateFunc: func(data string, err error) {
				if err != nil {
					t.Log(err.Error())
					t.Fail()
				}
				x, err := cacher.Get("1")
				if err != nil {
					t.Log(err.Error())
				}
				if fmt.Sprintf("%s", x) != data {
					t.Log(fmt.Sprintf("%s", x))
					t.Fail()
				}
			},
		},
		{
			name:        "Failure:: Set",
			requestBody: "Hello",
			validateFunc: func(data string, err error) {
				if err != nil {
					t.Log(err.Error())
					t.Fail()
				}
				x, err := cacher.Get("1")
				if err != nil {
					t.Log(err.Error())
				}
				if fmt.Sprintf("%s", x) != data {
					t.Log(fmt.Sprintf("%s", x))
					t.Fail()
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := tt.requestBody
			err := cacher.Set("1", data, 0)
			tt.validateFunc(data, err)
		})
	}

}

func TestGet(t *testing.T) {
	cacher := NewCacher(Config{
		Addr:     "localhost:9096",
		Password: "",
		DB:       0,
	})
	tests := []struct {
		name         string
		requestBody  string
		setupFunc    func(string)
		validateFunc func([]byte, string, error)
	}{
		{
			name:        "Success:: Get",
			requestBody: "Hello",
			setupFunc: func(data string) {
				err := cacher.Set("1", data, 0)
				if err != nil {
					t.Log(err.Error())
				}
			},
			validateFunc: func(s []byte, request string, err error) {
				if err != nil {
					t.Log(err.Error())
				}
				if fmt.Sprintf("%s", s) != request {
					t.Log(fmt.Sprintf("%s", s))
					t.Fail()
				}
			},
		},
		{
			name:        "Failure:: Get",
			requestBody: "Hello",
			setupFunc: func(data string) {
				err := cacher.Set("1", data, 0)
				if err != nil {
					t.Log(err.Error())
				}
			},
			validateFunc: func(s []byte, request string, err error) {
				if err != nil {
					t.Log(err.Error())
				}
				if fmt.Sprintf("%s", s) != request {
					t.Log(fmt.Sprintf("%s", s))
					t.Fail()
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := tt.requestBody
			tt.setupFunc(data)
			x, err := cacher.Get("1")
			tt.validateFunc(x, data, err)
		})
	}

}
