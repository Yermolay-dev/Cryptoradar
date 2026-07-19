package models

type AlertResponse struct {
	Status  string `json:"status"`
	Message string `json:"messages"`
}
type AlertRequest struct {
	Coin      string  `json:"coin"`
	Price     float64 `json:"price"`
	Direction string  `json:"direction"`
}
