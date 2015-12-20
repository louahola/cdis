package yahoofinance

type YqlResponse struct {
	Query struct {
		Results struct {
			Quote Quote
		}
	}
}
