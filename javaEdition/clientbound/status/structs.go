package status

type Version struct {
	Name     string `json:"name"`
	Protocol int    `json:"protocol"`
}

type Players struct {
	Max    uint `json:"max"`
	Online uint `json:"online"`
	Sample []PlayerSample
}

type PlayerSample struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type Description struct {
	Text string `json:"text"`
}
