package entities

type InputOrder struct {
	No                int     `json:"no" validate:"required"`
	PlatformProductId string  `json:"platformProductId" validate:"required"`
	Qty               int     `json:"qty" validate:"required"`
	UnitPrice         float64 `json:"unitPrice" validate:"required"`
	TotalPrice        float64 `json:"totalPrice" validate:"required"`
}

type CleanedOrder struct {
	No         int     `json:"no"`
	ProductId  string  `json:"productId"`
	MaterialId string  `json:"materialId ,omitempty"`
	ModelId    string  `json:"modelId,omitempty"`
	Qty        int     `json:"qty"`
	UnitPrice  float64 `json:"unitPrice"`
	TotalPrice float64 `json:"totalPrice"`
}
