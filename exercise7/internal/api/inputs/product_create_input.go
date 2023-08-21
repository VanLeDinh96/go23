package inputs

type CreateProductInput struct {
	Name     string `json:"name" binding:"required"`
	Price    int    `json:"price" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
}