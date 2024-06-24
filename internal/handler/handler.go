package handler

import (
	"github.com/chenchi1009/go-starter-kit/pkg/jwt"
	"github.com/chenchi1009/go-starter-kit/pkg/log"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	logger *log.Logger
}

func NewHandler(
	logger *log.Logger,
) *Handler {
	return &Handler{
		logger: logger,
	}
}
func GetUsernameFromCtx(ctx *gin.Context) string {
	v, exists := ctx.Get("claims")
	if !exists {
		return ""
	}
	return v.(*jwt.MyCustomClaims).Username
}
