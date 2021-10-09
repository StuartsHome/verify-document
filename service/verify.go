package service

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/stuartshome/verify-document/model"
)

type Process interface {
	Process()
}

type ProcessService struct {
	processors []Process
}
type ProcessData struct {
}

func NewProcessData() *ProcessData {
	return &ProcessData{}
}

var _ Process = &ProcessData{}

// func (pd *ProcessData) VerifyHandler(w http.ResponseWriter, r *http.Request) {
func (ps *ProcessService) VerifyHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Minute)
	defer cancel()

	err := r.ParseForm()
	if err != nil {
		return
	}

	report := &model.Report{
		Title:       "Discourses on Livy",
		Author:      "Niccolò Machiavelli",
		Language:    "Latin",
		PublishYear: 1531,
	}

	fmt.Println(report, ctx)

	// NewVerifyDocumentService
	// data, err := v.VerifyReport(ctx, report)

}

func (pd *ProcessData) Process() {

}
