package inputs

type CreateProductInput struct {
	Name     string `json:"name" binding:"required"`
	Price    float64    `json:"price" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
}