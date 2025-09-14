package domain

type Picture struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

type Shipping struct {
	FreeShipping bool   `json:"free_shipping"`
	Mode         string `json:"mode"`
}

type Stock struct {
	Available int `json:"available"`
	Sold      int `json:"sold"`
}

type Rating struct {
	User    string `json:"user"`
	Stars   int    `json:"stars"`
	Comment string `json:"comment"`
}

type Product struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Price       float64   `json:"price"`
	Currency    string    `json:"currency"`
	Condition   string    `json:"condition"`
	Pictures    []Picture `json:"pictures"`
	Thumbnail   string    `json:"thumbnail"`
	Permalink   string    `json:"permalink"`
	SellerID    string    `json:"seller_id"`
	Category    string   `json:"category"`               
	Brand       string   `json:"brand"`                  
	Tags        []string `json:"tags,omitempty"`         
	RelatedIDs  []string `json:"related_ids,omitempty"`  
	Shipping    Shipping  `json:"shipping"`
	Stock       Stock     `json:"stock"`
	Attributes  []string  `json:"attributes"`
	Description string    `json:"description"`
	RatingAvg   float64   `json:"rating_avg"`
	Ratings     []Rating  `json:"ratings"`
}
