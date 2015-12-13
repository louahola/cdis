package api

import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/louahola/cdis/repository"
	"github.com/louahola/cdis/monthlyreportfactory"
)

type MonthlyReportHandler struct {
	repo repository.Repository
}

func (this MonthlyReportHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//request := ApiRequest{}
	decoder := json.NewDecoder(r.Body)
	var t map[string]interface{}
	err := decoder.Decode(&t)
	if err != nil {
		fmt.Fprintf(w, "ERROR: %s", err.Error())
		return
	}

	symbol := t["symbol"].(string)

	report, err := MonthlyReportFactory.Generate(symbol)
	err = this.repo.Save(report)
	if err != nil {
		fmt.Fprintf(w, "ERROR: %s", err.Error())
		return
	}

	fmt.Fprintf(w, "SUCCESS: %s", symbol)
}
