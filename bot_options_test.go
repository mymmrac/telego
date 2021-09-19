package telego

import (
	"testing"
)

func TestFastHTTPClient(t *testing.T) {
	// type args struct {
	// 	client *fasthttp.Client
	// }
	// tests := []struct {
	// 	name string
	// 	args args
	// 	want BotOption
	// }{
	//
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		if got := FastHTTPClient(tt.args.client); !reflect.DeepEqual(got, tt.want) {
	// 			t.Errorf("FastHTTPClient() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
}

// func TestBot_Logger(t *testing.T) {
// 	bot := getBot(t)
//
// 	t.Run("default-logger", func(t *testing.T) {
// 		assert.NotPanics(t, func() {
// 			bot.DefaultLogger(true, true)
// 		})
// 	})
//
// 	t.Run("set-logger", func(t *testing.T) {
// 		assert.NotPanics(t, func() {
// 			var l Logger
// 			bot.SetLogger(l)
// 		})
// 	})
// }

// func TestBot_SetAPIServer(t *testing.T) {
// 	bot := getBot(t)
//
// 	tests := []struct {
// 		name  string
// 		url   string
// 		isErr bool
// 	}{
// 		{
// 			name:  "success",
// 			url:   defaultBotAPIServer,
// 			isErr: false,
// 		},
// 		{
// 			name:  "empty",
// 			url:   "",
// 			isErr: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			actual := bot.SetAPIServer(tt.url)
// 			if tt.isErr {
// 				assert.Error(t, actual)
// 				return
// 			}
// 			assert.NoError(t, actual)
// 		})
// 	}
// }

// func TestBot_SetClient(t *testing.T) {
// 	bot := getBot(t)
//
// 	tests := []struct {
// 		name   string
// 		client *fasthttp.Client
// 		isErr  bool
// 	}{
// 		{
// 			name:   "success",
// 			client: &fasthttp.Client{},
// 	 		isErr:  false,
// 		},
// 		{
// 			name:   "error",
// 			client: nil,
// 			isErr:  true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			actual := bot.SetClient(tt.client)
// 			if tt.isErr {
// 				assert.Error(t, actual)
// 				return
// 			}
// 			assert.NoError(t, actual)
// 		})
// 	}
// }
