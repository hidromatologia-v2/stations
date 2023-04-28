package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hidromatologia-v2/models"
	"github.com/memphisdev/memphis.go"
)

type Handler struct {
	Controller    *models.Controller
	AlertProducer *memphis.Producer
	*gin.Engine
}

func (h *Handler) Close() {
	h.AlertProducer.Destroy()
	h.Controller.Close()
}

func New(c *models.Controller, p *memphis.Producer) *Handler {
	h := &Handler{Controller: c, AlertProducer: p}
	h.Engine = gin.Default()
	auth := h.Engine.Group(APIRoute, h.Authorize)
	auth.GET(StationRoute, h.Station)
	auth.PUT(RegistryRoute, h.Registry)
	return h
}
