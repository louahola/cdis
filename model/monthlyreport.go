package model

import "time"

type MonthlyReport struct {
	Id string `json:"id" bson:"_id,omitempty"`
	Symbol string `json:"symbol" bson:"symbol"`
	Name string `json:"name" bson:"name"`
	StockExchange string `json:"stock_exchange" bson:"stock_exchange"`
	Date time.Time `json:"date" bson:"date"`
	Price float32 `json:"price" bson:"price"`
	LastMonth float32 `json:"last_month" bson:"last_month"`
	YearHigh float32 `json:"year_high" bson:"year_high"`
	YearLow float32 `json:"year_low" bson:"year_low"`
	PegRatio float32 `json:"peg_ratio" bson:"peg_ratio"`
	PeTtm float32 `json:"pe_ttm" bson:"pe_ttm"`
	Beta float32 `json:"beta" bson:"beta"`
	YearlyYield float32 `json:"yearly_yield" bson:"yearly_yield"`
	YearlyYieldPercent float32 `json:"yearly_yield_percent" bson:"yearly_yield_percent"`
	Eps float32 `json:"eps" bson:"eps"`
	DebtEquityRatio float32 `json:"debt_equity_ratio" bson:"debt_equity_ratio"`
}
