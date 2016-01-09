package actions

type Webhook struct {
	URL      string `json:"url,omitempty" validate:"required"`
	Method   string `json:"method,omitempty" validate:"required"`
	Body     string `json:"body, omitempty"`
	Interval int    `json:"body,omitempty"`
	Count    int    `json:"body,omitempty"`
}
