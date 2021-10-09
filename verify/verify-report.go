package verify

import (
	"context"

	"github.com/stuartshome/verify-document/model"
)

// type Verify interface {
// 	VerifyReport(ctx context.Context, report model.Report) (*model.AllReports, error)
// }

type verifyImpl struct {
}

func NewVerifyReportService() *verifyImpl {
	return &verifyImpl{}
}

// var _ Verify = &verifyImpl{}

func (v *verifyImpl) VerifyReport(ctx context.Context, report model.Report) (*model.AllReports, error) {
	return nil, nil
}
