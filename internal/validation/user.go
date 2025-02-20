package validation

type RegisterUserStoreAdminInput struct {
	Name            string `json:"name" binding:"required"`
	Phone           string `json:"phone" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=8"`
	ConfirmPassword string `json:"confrimed_password" binding:"required"`
}

type LoginUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
