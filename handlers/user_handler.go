package handlers

type UserHandler struct{
	// UserService service.IUserService
}

func NewUserHandler() *UserHandler{
	return &UserHandler{
		// UserService: service.NewUserService(),
	}
}

func 