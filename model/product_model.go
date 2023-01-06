package model

type ProductRequest struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	Quantity int32  `json:"quantity"`
}

type ProductResponse struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	Quantity int32  `json:"quantity"`
}
