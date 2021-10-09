package model

type Report struct {
	Title       string
	Author      string
	Language    string
	PublishYear int
}

type AllReports struct {
	Report Report
}

type ReportOutput struct {
	Output string
}
