package User

type RegisterUserRules struct {
	FirstName       string `form:"first_name" json:"first_name" binding:"required"`
	LastName        string `form:"last_name" json:"last_name" binding:"required"`
	EmailAddress    string `form:"email_address" json:"email_address" binding:"required,email"`
	Password        string `form:"password" json:"password" binding:"required"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password" binding:"required,eqfield=Password"`
}
