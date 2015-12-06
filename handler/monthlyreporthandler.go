package handler
import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/louahola/cdis/factory"
	"github.com/louahola/cdis/repository"
)

type MonthlyReportHandler struct {
	repo repository.Repository
}

func (this *MonthlyReportHandler) HandleHTTP(w http.ResponseWriter, r *http.Request) {
	request := ApiRequest{httpRequest: r}
	decoder := json.NewDecoder(r.Body)
	var t map[string]interface{}
	err := decoder.Decode(&t)
	if err != nil {
		fmt.Fprintf(w, "ERROR: %s", err.Error())
		return
	}

	symbol := t["symbol"].(string)

	report, err := factory.GenerateMonthlyReport(symbol)
	err = this.repo.Save(report)
	if err != nil {
		fmt.Fprintf(w, "ERROR: %s", err.Error())
		return
	}

	fmt.Fprintf(w, "SUCCESS: %s", symbol)
}
