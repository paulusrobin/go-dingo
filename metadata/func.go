package metadata

type (
	Func struct {
		Name string `json:"name"`
		Dto  string `json:"dto"`
	}

	HttpFunc struct {
		Func
		Method string `json:"method"`
		Routes string `json:"routes"`
	}

	MessagingFunc struct {
		Func
		Topic string `json:"topic"`
	}

	WorkerFunc struct {
		Func
		Duration string `json:"duration"`
	}
)
