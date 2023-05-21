package httputils

import "github.com/gin-gonic/gin"

var InternalServerError = gin.H{
	"error": "internal server error",
}

var ConflictError = gin.H{
	"error": "entry exists",
}

var BadRequestError = gin.H{
	"error": "bad request",
}

var NotFoundError = gin.H{
	"error": "not found",
}
