package domain

type PaymentMethod struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Installments []int  `json:"installments"`
}
