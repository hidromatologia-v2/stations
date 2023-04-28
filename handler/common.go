package handler

const (
	APIRoute      = "/api"
	StationRoute  = "/station"
	RegistryRoute = "/registry"
)

const (
	XAuthKeyHeader = "X-Auth-Key"
)

const (
	StationVariable = "STATION"
)

type Response struct {
	Message string `json:"message"`
}

var (
	UnauthorizedResponse      = Response{Message: "unauthorized"}
	RegistriesCreatedResponse = Response{Message: "registries created"}
)
