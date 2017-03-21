package report

// Action logs events that need intervention or resolving.
//
//     report.Action("http.unavailable", report.Data{"path":req.URL.Path, "error":err.Error()})
//
func Action(event string, payload Data) {
	payload["type"] = "action"
	payload["event"] = event

	publishRawEvent(payload)
}
