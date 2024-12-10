package validation

type Login struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"passwpord" binding:"required"`
}
