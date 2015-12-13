package MonthlyReportFactory

import (
	"github.com/louahola/cdis/model"
	"github.com/louahola/yahoofinance"
	"time"
	"fmt"
)

func Generate(symbol string) (model.MonthlyReport, error) {
	quote, err := yahoofinance.GetQuote(symbol)
	monthlyReport := model.MonthlyReport{
		Id: fmt.Sprintf("%s-%04d-%02d", symbol, time.Now().Year(), int(time.Now().Month())),
		Price: quote.LastTradePriceOnly,
		Symbol: quote.Symbol,
		Name: quote.Name,
		Date: time.Now(),
		StockExchange: quote.StockExchange,
		LastMonth: 0.0,
		YearHigh: quote.YearHigh,
		YearLow: quote.YearLow,
		PegRatio: quote.PEGRatio,
		PeTtm: quote.PERatio,
		Beta: 0.0,
		YearlyYield: quote.DividendShare,
		YearlyYieldPercent: quote.DividendYield,
		Eps: 0.0,
		DebtEquityRatio: 0.0,
	}
	return monthlyReport, err
}
