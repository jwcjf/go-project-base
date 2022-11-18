package antd

// Silent ...
const (
	Silent       = "0"
	MessageWarn  = "1"
	MessageError = "2"
	Notification = "4"
	Page         = "9"
)

// Response ...
type Response struct {
	Success      bool   `json:"success,omitempty"`      // if request is success
	ErrorCode    string `json:"errorCode,omitempty"`    // code for errorType
	ErrorMessage string `json:"errorMessage,omitempty"` // message display to user
	ShowType     string `json:"showType,omitempty"`     // error display type： 0 silent; 1 message.warn; 2 message.error; 4 notification; 9 page
	TraceId      string `json:"traceId,omitempty"`      // Convenient for back-end Troubleshooting: unique request ID
	Host         string `json:"host,omitempty"`         // onvenient for backend Troubleshooting: host of current access server
}
type response struct {
	Response
	Data interface{} `json:"data,omitempty"` // response data
}

// Pages ...
type Pages struct {
	Total    int `json:"total,omitempty"`
	Current  int `json:"current,omitempty"`
	PageSize int `json:"pageSize,omitempty"`
}

type pages struct {
	Pages
	Data interface{} `json:"data,omitempty"`
}
