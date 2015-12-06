package main

import (
	"fmt"

	"github.com/louahola/cdis/model"
	"github.com/louahola/cdis/stringutil"
)

func main() {
	report := &model.BasicReport{Symbol: "QCOM", Name: "Qualcomm Incorporated"}
	fmt.Println(report.Name)
	fmt.Printf(stringutil.Reverse("!oG ,olleH"))
}