package cart

type UserID string
type GoodID string

type Cart struct {
	UserID UserID         `json:"user_id,omitempty"`
	Goods  map[GoodID]int `json:"goods,omitempty"`
}

type Order struct {
	UserID UserID `json:"user_id,omitempty"`
	Good   GoodID `json:"good,omitempty"`
	Count  int    `json:"count,omitempty"`
}

func NewUserIdFromString(id string) UserID {
	return UserID(id)
}
