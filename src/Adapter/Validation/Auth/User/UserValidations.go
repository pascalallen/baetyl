package User

type RegisterUserRules struct {
	FirstName       string `form:"first_name" json:"first_name" binding:"required,max=100"`
	LastName        string `form:"last_name" json:"last_name" binding:"required,max=100"`
	EmailAddress    string `form:"email_address" json:"email_address" binding:"required,max=100,email"`
	Password        string `form:"password" json:"password" binding:"required"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password" binding:"required,eqfield=Password"`
}
