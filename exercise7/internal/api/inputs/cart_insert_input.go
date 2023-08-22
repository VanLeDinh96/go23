package inputs

type CartInsertInput struct {
	ProductID int `json:"productId" binding:"required"`
	Quantity  int `json:"quantity" binding:"required"`
	Name	  string `json:"name" binding:"required"`
	Price	  int `json:"price" binding:"required"`
}
