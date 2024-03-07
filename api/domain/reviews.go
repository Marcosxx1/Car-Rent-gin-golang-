package domain

type Reviews struct {
	ID      string `json:"id"`
	UserId  string `json:"user_id"`
	CarId   string `json:"car_id"`
	Rating  *int   `json:"rating"`
	Content string `json:"content"`
}
