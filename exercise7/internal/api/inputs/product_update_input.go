package inputs

type UpdateProductInput struct {
	Name     string  `json:"name" binding:"required,min=3,max=50"`
	Price    float64 `json:"price" binding:"required,gt=0"`
	Quantity int     `json:"quantity" binding:"required,gt=0"`
}
