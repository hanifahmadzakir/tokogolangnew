package entity


type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Photo       string `json:"photo"`
	Price       int    `json:"price"`
	Category    string `json:"category"`
	Description string `json:"description"`
}

type ProductRequest struct {
	Name        string `json:"name" form:"name"`
	Photo       string `json:"photo" form:"photo"`
	Price       int    `json:"price" form:"price"`
	Category    string `json:"category" form:"category"`
	Description string `json:"description" form:"description"`
}

type ProductResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Photo       string `json:"photo"`
	Price       int    `json:"price"`
	Category    string `json:"category"`
	Description string `json:"description"`
}

type UpdateProductRequest struct {
	Name        string `json:"name"`
	Photo       string `json:"photo"`
	Price       int    `json:"price"`
	Category    string `json:"category"`
	Description string `json:"description"`
}