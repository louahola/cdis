package yahoofinance
import (
	"fmt"
	"encoding/json"
	"net/url"
	"io/ioutil"
	"net/http"
)


func GetQuote(symbol string) (*Quote, error) {
	query := QuoteYQL(symbol)

	params := url.Values{}
	params.Set("q", query)
	params.Set("format", "json")
	params.Set("env", "store://datatables.org/alltableswithkeys")

	url := "http://query.yahooapis.com/v1/public/yql?" + params.Encode()
	fmt.Println(query)
	fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, nil
	} else {
		defer resp.Body.Close()

		httpBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		fmt.Println(string(httpBody))

		var yqlResp YqlResponse
		err = json.Unmarshal(httpBody, &yqlResp)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		quote := &yqlResp.Query.Results.Quote
		return quote, nil
	}
}
