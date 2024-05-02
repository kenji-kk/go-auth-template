package handler

type RootHandlers struct {
	UserHandler UserHandler
}

func NewRootHandlers(userHandler UserHandler) RootHandlers {
	return RootHandlers{UserHandler: userHandler}
}
