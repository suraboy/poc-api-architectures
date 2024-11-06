package rest

type UserResponse struct {
	UserID      string `json:"userID"`
	UserName    string `json:"userName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}
