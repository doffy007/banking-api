package request

type LogoutUserRequestBody struct {
	Email string `json:"email" validate:"required"`
}

func (r *LogoutUserRequestBody) CustomValidate() error {
	return nil
}
