package apiperms

type ApiPermission = string

const (
	PostUser     ApiPermission = "post_user"
	GetUser      ApiPermission = "get_user"
	UserAge      ApiPermission = "user_age"
	Applications ApiPermission = "applications"
)
