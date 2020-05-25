package metadata

type (
	Infrastructure struct {
		Http      InfraHttp      `json:"http"`
		Messaging InfraMessaging `json:"messaging"`
		Worker    InfraWorker    `json:"worker"`
	}

	InfraMessaging struct {
		Initialize bool            `json:"initialize"`
		Functions  []MessagingFunc `json:"functions"`
	}

	InfraHttp struct {
		Initialize bool       `json:"initialize"`
		Functions  []HttpFunc `json:"functions"`
	}

	InfraWorker struct {
		Initialize bool         `json:"initialize"`
		Functions  []WorkerFunc `json:"functions"`
	}
)
