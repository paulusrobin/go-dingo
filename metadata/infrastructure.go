package metadata

type (
	Infrastructure struct {
		Initialize bool           `json:"initialize"`
		Http       InfraHttp      `json:"http"`
		Messaging  InfraMessaging `json:"messaging"`
	}

	InfraMessaging struct {
		Initialize bool            `json:"initialize"`
		Functions  []MessagingFunc `json:"functions"`
	}

	InfraHttp struct {
		Initialize bool       `json:"initialize"`
		Functions  []HttpFunc `json:"functions"`
	}
)
