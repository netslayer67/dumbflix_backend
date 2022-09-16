package usersdto

type UpdateUserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	IsAdmin  bool   `json:"isAdmin" form:"isAdmin"`
	FullName string `json:"full_name" form:"fullName"`
	Phone    string `json:"phone" form:"phone"`
	Gender   string `json:"gender" form:"gender"`
	Address  string `json:"address" form:"address"`
}
