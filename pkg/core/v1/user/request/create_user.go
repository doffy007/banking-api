package request

type CreateUserRequestBody struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (r *CreateUserRequestBody) CustomValidate() error {
	return nil
}
