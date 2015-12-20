package yahoofinance
import (
	"fmt"
)

type Quote struct {
	DividendShare float32 `json:"string"`
	DividendYield float32 `json:"string"`
	EarningsShare float32 `json:"string"`
	LastTradePriceOnly float32 `json:"string"`
	LastTradeDate string
	Name string
	PEGRatio float32 `json:"string"`
	PERatio float32 `json:"string"`
	StockExchange string
	Symbol string
	YearHigh float32 `json:"string"`
	YearLow float32 `json:"string"`
}

func QuoteYQL(symbol string) string {
	return fmt.Sprintf(`
		SELECT
		DividendShare,
		DividendYield,
		EarningsShare,
		LastTradePriceOnly,
		LastTradeDate,
		Name,
		PEGRatio,
		PERatio,
		StockExchange,
		Symbol,
		YearHigh,
		YearLow
		FROM yahoo.finance.quotes WHERE symbol="%s"`, symbol)
}

