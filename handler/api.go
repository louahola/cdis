package handler
import (
	"net/http"
	"github.com/louahola/cdis/repository"
)

type Api struct {
	repo repository.Repository
}

func (this *Api) initialize() {
	http.Handle("/monthlyReport", MonthlyReportHandler{repo: repo})
}
