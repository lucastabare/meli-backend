package domain

type Ad struct {
	ID        string `json:"id"`
	Placement string `json:"placement"`
	Image     string `json:"image"`
	Href      string `json:"href"`
	Alt       string `json:"alt"`
}
