package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hidromatologia-v2/models"
)

func (h *Handler) Authorize(ctx *gin.Context) {
	key := ctx.GetHeader(XAuthKeyHeader)
	station, authErr := h.Controller.AuthorizeAPIKey(key)
	if authErr != nil {
		if errors.Is(authErr, models.ErrUnauthorized) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, UnauthorizedResponse)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, Response{Message: authErr.Error()})
		return
	}
	ctx.Set(StationVariable, station)
	ctx.Next()
}
