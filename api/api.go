package api

import (
	"github.com/louahola/cdis/repository"
	"net/http"
)

func Initialize(repo repository.Repository) {
	http.Handle("/api/monthlyReport", MonthlyReportHandler{repo: repo})
}
