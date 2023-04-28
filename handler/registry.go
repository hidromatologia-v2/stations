package handler

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hidromatologia-v2/models/tables"
)

func (h *Handler) Registry(ctx *gin.Context) {
	var registries []tables.SensorRegistry
	bErr := ctx.Bind(&registries)
	if bErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, Response{Message: bErr.Error()})
		return
	}
	station := ctx.MustGet(StationVariable).(*tables.Station)
	pErr := h.Controller.PushRegistry(station, registries)
	if pErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, Response{Message: pErr.Error()})
		return
	}
	for index := 0; index < len(registries); index++ {
		var registry bytes.Buffer
		json.NewEncoder(&registry).Encode(registries[index])
		go h.AlertProducer.Produce(registry.Bytes())
	}
	ctx.JSON(http.StatusCreated, RegistriesCreatedResponse)
	ctx.Done()
}
