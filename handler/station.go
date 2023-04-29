package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hidromatologia-v2/models/tables"
)

func (h *Handler) Station(ctx *gin.Context) {
	station := ctx.MustGet(StationVariable).(*tables.Station)
	result, qErr := h.Controller.QueryStationNoAPIKey(station)
	if qErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, Response{Message: qErr.Error()})
		return
	}
	ctx.JSON(http.StatusOK, result)
	ctx.Done()
}
