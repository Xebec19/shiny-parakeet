package auth

type registerRequest struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" binding:"email"`
	Password  string `json:"password" binding:"min=8"`
}
