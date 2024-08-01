package main

import "encoding/json"

type CurrencyList []Currency

func UnmarshalCurrencyList(data []byte) (CurrencyList, error) {
	var r CurrencyList
	err := json.Unmarshal(data, &r)
	return r, err
}

type Currency struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Ticker struct {
	BtcUsd BtcUsd `json:"BTC_USD"`
}

type BtcUsd struct {
	BuyPrice  string `json:"buy_price"`
	SellPrice string `json:"sell_price"`
	LastTrade string `json:"last_trade"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Avg       string `json:"avg"`
	Vol       string `json:"vol"`
	VolCurr   string `json:"vol_curr"`
	Updated   int64  `json:"updated"`
}

type TickerStat map[string]TickerStatValue

func UnmarshalTickerStat(data []byte) (TickerStat, error) {
	var r TickerStat
	err := json.Unmarshal(data, &r)
	return r, err
}

type TickerStatValue struct {
	BuyPrice  string `json:"buy_price"`
	SellPrice string `json:"sell_price"`
	LastTrade string `json:"last_trade"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Avg       string `json:"avg"`
	Vol       string `json:"vol"`
	VolCurr   string `json:"vol_curr"`
	Updated   int64  `json:"updated"`
}

func UnmarshalTrades(data []byte) (Trades, error) {
	var r Trades
	err := json.Unmarshal(data, &r)
	return r, err
}

type Trades map[string][]Trade

type Trade struct {
	TradeID  int64  `json:"trade_id"`
	Date     int64  `json:"date"`
	Type     Type   `json:"type"`
	Quantity string `json:"quantity"`
	Price    string `json:"price"`
	Amount   string `json:"amount"`
}

type Type string

const (
	Buy  Type = "buy"
	Sell Type = "sell"
)

func UnmarshalOrderBook(data []byte) (OrderBook, error) {
	var r OrderBook
	err := json.Unmarshal(data, &r)
	return r, err
}

type OrderBook map[string]Order

type Order struct {
	AskQuantity string     `json:"ask_quantity"`
	AskAmount   string     `json:"ask_amount"`
	AskTop      string     `json:"ask_top"`
	BidQuantity string     `json:"bid_quantity"`
	BidAmount   string     `json:"bid_amount"`
	BidTop      string     `json:"bid_top"`
	Ask         [][]string `json:"ask"`
	Bid         [][]string `json:"bid"`
}

func UnmarshalCandles(data []byte) (Candles, error) {
	var r Candles
	err := json.Unmarshal(data, &r)
	return r, err
}

type Candles map[string][]Candle

type Candle struct {
	T int64   `json:"t"`
	O float64 `json:"o"`
	C float64 `json:"c"`
	H float64 `json:"h"`
	L float64 `json:"l"`
	V float64 `json:"v"`
}
