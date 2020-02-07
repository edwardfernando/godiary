package contract

// Response defines standard response returned by all APIs
type Response struct {
	Data    interface{}   `json:"data"`
	Errors  []ResponseErr `json:"errors"`
	Success bool          `json:"success"`
}

// ResponseErr defines standard response returned when there is error
type ResponseErr struct {
	Code            string `json:"code"`
	MessageTitle    string `json:"message_title"`
	Message         string `json:"message"`
	MessageSeverity string `json:"message_severity"`
}
