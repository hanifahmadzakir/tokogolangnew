package entity

type Cart struct {
	ID        int  `json:"id"`
	Quantity  int  `json:"quantity"`
	Checkout  bool `json:"checkout"`
	UserID    int
	User      User
	ProductID int
	Product   Product
}

type CartRequest struct {
	Quantity  int  `json:"quantity"`
	Checkout  bool `json:"checkout"`
	UserID    int
	User      User
	ProductID int
	Product   Product
}

type CartResponse struct {
	ID        int  `json:"id"`
	Quantity  int  `json:"quantity"`
	Checkout  bool `json:"checkout"`
	UserID    int
	User      User
	ProductID int
	Product   Product
}

type UpdateCartRequest struct {
	Quantity  int  `json:"quantity"`
	Checkout  bool `json:"checkout"`
	UserID    int
	User      User
	ProductID int
	Product   Product
}
