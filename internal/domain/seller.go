package domain

type Seller struct {
	ID            string  `json:"id"`
	Nickname      string  `json:"nickname"`
	City          string  `json:"city"`
	Sales         int     `json:"sales"`
	Reputation    string  `json:"reputation"`
	RatingAverage float64 `json:"rating_average"`
}
