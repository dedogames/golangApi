package entities

type ProductResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type ProductListResponse struct {
	StatusCode  int            `json:"statusCode"`
	Message     string         `json:"message"`
	ProductList []*ProductBody `json:"products"`
}
