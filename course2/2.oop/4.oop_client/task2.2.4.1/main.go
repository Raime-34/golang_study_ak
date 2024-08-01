package main

import (
	"internal/itoa"
	"io"
	"net/http"
	"strconv"
	"strings"
)

const (
	exmoBaseUrl = "https://api.exmo.com/v1.1/"

	currencies     = "currency/list/extended"
	ticker         = "ticker"
	trades         = "trades"
	orderBooks     = "order_book"
	candlesHistory = "candles_history"
)

type Exchanger interface {
	GetTicker() (TickerStat, error)
	GetTrades(pairs ...string) (Trades, error)
	GetOrderBook(limit int, pairs ...string) (OrderBook, error)
	GetCurrencies() (CurrencyList, error)
	GetCandlesHistory(pair string, limit int, start, end int64) (Candles, error)
	GetClosePrice(pair string, limit int, start, end int64) ([]float64, error)
}

type Exmo struct {
	client *http.Client
	url    string
}

func (e *Exmo) GetClosePrice(pair string, limit int, start, end int64) ([]float64, error) {
	candles, err := e.GetCandlesHistory(pair, limit, start, end)
	if err != nil {
		return nil, err
	}

	var result []float64
	for _, candle := range candles["candles"] {
		result = append(result, candle.C)
	}

	return result, nil
}

func (e *Exmo) GetCandlesHistory(pair string, limit int, start, end int64) (Candles, error) {
	res, err := e.doRequest(
		"GET",
		candlesHistory,
		nil,
		false,
		map[string]string{
			"symbol":     pair,
			"resolution": itoa.Itoa(limit),
			"from":       strconv.FormatInt(start, 10),
			"to":         strconv.FormatInt(end, 10),
		})
	if err != nil {
		return nil, err
	}

	return UnmarshalCandles(res)
}

func (e *Exmo) GetOrderBook(limit int, pairs ...string) (OrderBook, error) {
	res, err := e.doRequest("POST", orderBooks, composePairs(append(pairs, "limit="+itoa.Itoa(limit))...), true, nil)
	if err != nil {
		return nil, err
	}

	return UnmarshalOrderBook(res)
}

func (e *Exmo) GetTrades(pairs ...string) (Trades, error) {
	res, err := e.doRequest("POST", trades, composePairs(pairs...), true, nil)
	if err != nil {
		return nil, err
	}

	return UnmarshalTrades(res)
}

func (e *Exmo) GetTicker() (TickerStat, error) {
	reqBody := strings.NewReader(`{
    "BTC_USD": {
        "buy_price": "589.06",
        "sell_price": "592",
        "last_trade": "591.221",
        "high": "602.082",
        "low": "584.51011695",
        "avg": "591.14698808",
        "vol": "167.59763535",
        "vol_curr": "99095.17162071",
        "updated": 1470250973
    }
}`)

	res, _ := e.doRequest("POST", ticker, reqBody, false, nil)

	return UnmarshalTickerStat(res)
}

func (e *Exmo) GetCurrencies() (CurrencyList, error) {
	res, _ := e.doRequest("GET", currencies, nil, false, nil)

	return UnmarshalCurrencyList(res)
}

func NewExmo(opts ...func(exmo *Exmo)) *Exmo {
	client := &Exmo{
		client: &http.Client{},
		url:    exmoBaseUrl,
	}

	for _, f := range opts {
		f(client)
	}

	return client
}

func WithClient(client *http.Client) func(exmo *Exmo) {
	return func(exmo *Exmo) {
		exmo.client = client
	}
}

func WithUrl(url string) func(exmo *Exmo) {
	return func(exmo *Exmo) {
		exmo.url = url
	}
}

func (e *Exmo) doRequest(method string, endpoint string, body *strings.Reader, urlencoded bool, params map[string]string) ([]byte, error) {
	if body == nil {
		body = strings.NewReader("")
	}

	req, err := http.NewRequest(method, e.url+endpoint, body)
	if err != nil {
		return nil, err
	}

	if urlencoded {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	q := req.URL.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	res, err := e.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bytes, _ := io.ReadAll(res.Body)
	return bytes, nil
}
