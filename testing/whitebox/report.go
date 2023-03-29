package report

var getReport = func(id string) string {
	//connect to database
	return "avcdf"
}

func generateReport(id string) string {
	data := getReport(id)
	return data + " : VERIFIED."
}
