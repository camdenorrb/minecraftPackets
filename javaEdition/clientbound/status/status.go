package status

type Response struct {
	Version struct {
		Name     string `json:"name"`
		Protocol int    `json:"protocol"`
	}
	Players struct {
		Max    uint `json:"max"`
		Online uint `json:"online"`
		Sample []struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		}
	}
	Description struct {
		Text string `json:"text"`
	}
	Favicon           string `json:"favicon"`
	EnforceSecureChat bool   `json:"enforceSecureChat"`
}

type PingResponse struct {
	Payload int64
}
