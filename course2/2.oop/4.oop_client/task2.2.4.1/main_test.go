package main

import (
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestExmo_Methods(t *testing.T) {
	tests := []struct {
		name    string
		method  func(*Exmo) (interface{}, error)
		wantErr bool
	}{
		{
			name:    "Regular GetCurrencies() test",
			method:  func(e *Exmo) (interface{}, error) { return e.GetCurrencies() },
			wantErr: false,
		},
		{
			name:    "Regular GetTicker() test",
			method:  func(e *Exmo) (interface{}, error) { return e.GetTicker() },
			wantErr: false,
		},
		{
			name:    "Regular GetTrades() test",
			method:  func(e *Exmo) (interface{}, error) { return e.GetTrades("BTC_USD") },
			wantErr: false,
		},
		{
			name:    "Regular GetOrderBook() test",
			method:  func(e *Exmo) (interface{}, error) { return e.GetOrderBook(100, "BTC_USD") },
			wantErr: false,
		},
		{
			name:    "Regular GetCandlesHistory() test",
			method:  func(e *Exmo) (interface{}, error) { return e.GetCandlesHistory("BTC_USD", 30, 1585556979, 1585557979) },
			wantErr: false,
		},
		{
			name:    "Incorrect GetCandlesHistory() test",
			method:  func(e *Exmo) (interface{}, error) { return e.GetCandlesHistory("BTC_USD", 0, 321, -342132) },
			wantErr: true,
		},
		{
			name:    "Regular GetClosePrice() test",
			method:  func(e *Exmo) (interface{}, error) { return e.GetClosePrice("BTC_USD", 30, 1585556979, 1585557979) },
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := NewExmo(WithClient(&http.Client{}), WithUrl(exmoBaseUrl))
			got, err := tt.method(e)
			if (err != nil) != tt.wantErr {
				t.Errorf("%s error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}

			val := reflect.ValueOf(got)
			kind := val.Kind()

			// Проверяем, является ли результат срезом, массивом или мапой
			if kind != reflect.Slice && kind != reflect.Array && kind != reflect.Map {
				t.Errorf("%s got = %T, want a slice, array or map", tt.name, got)
				return
			}

			// Если результат является срезом или мапой, проверяем, что он не пустой
			if kind == reflect.Slice || kind == reflect.Array || kind == reflect.Map {
				if val.Len() == 0 {
					t.Errorf("%s got = empty, want not empty", tt.name)
				}
			}
		})
	}
}

func TestExmo_doRequest(t *testing.T) {
	type args struct {
		method     string
		endpoint   string
		body       *strings.Reader
		urlencoded bool
		params     map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Incorrect query test",
			args: args{
				method:     "@fefewf231322----",
				endpoint:   "@fe32r32//f'/f.",
				body:       nil,
				urlencoded: false,
				params:     nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := NewExmo(WithClient(&http.Client{}), WithUrl(exmoBaseUrl))
			_, err := e.doRequest(tt.args.method, tt.args.endpoint, tt.args.body, tt.args.urlencoded, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMockClient(t *testing.T) {
	client := NewExmo(WithClient(&http.Client{}), WithUrl("fewfewfwefw efwqfe"))
	_, err := client.GetCurrencies()

	if err == nil {
		t.Errorf("GetCurrencies() error = %v, wantErr %v", err, true)
	}

	_, err = client.GetClosePrice("BTC_USD", 30, 2314321, 31432421)

	if err == nil {
		t.Errorf("GetClosePrice() error = %v, wantErr %v", err, true)
	}

	_, err = client.GetCandlesHistory("BTC_USD", 30, 2314321, 31432421)

	if err == nil {
		t.Errorf("GetCandlesHistory() error = %v, wantErr %v", err, true)
	}

	_, err = client.GetOrderBook(30, []string{"BTC_USD"}...)

	if err == nil {
		t.Errorf("GetOrderBook() error = %v, wantErr %v", err, true)
	}

	_, err = client.GetTrades([]string{"BTC_USD"}...)

	if err == nil {
		t.Errorf("GetTrades() error = %v, wantErr %v", err, true)
	}
}

//func MarshalTest(t *testing.T) {
//	invalidJson := []byte("")
//
//	_, err := UnmarshalTrades(invalidJson)
//	if err == nil {
//		t.Errorf("UnmarshalTrades() error = %v, expected not nil", err)
//	}
//
//	_, err = UnmarshalCandles(invalidJson)
//	if err == nil {
//		t.Errorf("UnmarshalCandles() error = %v, expected not nil", err)
//	}
//
//	_, err = UnmarshalCurrencyList(invalidJson)
//	if err == nil {
//		t.Errorf("UnmarshalCurrencyList() error = %v, expected not nil", err)
//	}
//
//	_, err = UnmarshalCurrencyList(invalidJson)
//	if err == nil {
//		t.Errorf("UnmarshalCurrencyList() error = %v, expected not nil", err)
//	}
//}
