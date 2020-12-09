package request

type Login struct {
	Phone    string `form:"phone" validate:"required"`
	Password string `form:"password" validate:"required,min=6,max=20"`
	SmsCode  string `form:"smscode" validate:"required,len=6"`
}

type Register struct {

}
