package catagoriser

import (
	"net/url"
	"time"

	"github.com/stuartshome/verify-document/http_client"
	"github.com/stuartshome/verify-document/model"
)

type retryConfig struct {
	MaxRequestTime string
	RetryDelayTime string
}
type CategorisationService interface {
	CategoriseDocument(doc model.Report) (model.ReportOutput, error)
}

type CategorisationImpl struct {
	url                url.URL
	client             http_client.Client
	maxRequestDuration time.Duration
	retryDelayDuration time.Duration
}

var _ CategorisationService = &CategorisationImpl{}

func (service *CategorisationImpl) CategoriseDocument(document model.Report) (model.ReportOutput, error) {

	return model.ReportOutput{}, nil
}
