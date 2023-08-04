package handler

import (
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/hidromatologia-v2/models"
	"github.com/hidromatologia-v2/models/common/connection"
	"github.com/hidromatologia-v2/models/tables"
	"github.com/memphisdev/memphis.go"
	"github.com/stretchr/testify/assert"
)

func TestRegistry(t *testing.T) {
	t.Run("Valid", func(tt *testing.T) {
		expect, h, stationName, closeFunc := defaultHandler(tt)
		defer h.Close()
		defer closeFunc()
		u := tables.RandomUser()
		assert.Nil(tt, h.Controller.DB.Create(u).Error)
		s := tables.RandomStation(u)
		assert.Nil(tt, h.Controller.DB.Create(s).Error)
		sensorUUID := s.Sensors[0].UUID
		expect.PUT(RegistryRoute).
			WithJSON([]tables.SensorRegistry{
				{
					SensorUUID: sensorUUID,
					Value:      10,
				},
				{
					SensorUUID: sensorUUID,
					Value:      11,
				},
				{
					SensorUUID: sensorUUID,
					Value:      12,
				},
			}).
			WithHeader(XAuthKeyHeader, s.APIKey).
			Expect().
			Status(http.StatusCreated)
		registries, queryErr := h.Controller.Historical(&models.HistoricalFilter{
			SensorUUID: sensorUUID,
		})
		assert.Nil(tt, queryErr)
		assert.Len(tt, registries, 3)
		consumer := connection.NewConsumer(tt, stationName)
		tick := time.NewTicker(time.Millisecond)
		defer tick.Stop()
		var msgs []*memphis.Msg
		for i := 0; i < 100; i++ {
			var fErr error
			msgs, fErr = consumer.Fetch(10, true)
			if fErr != nil && !strings.Contains(fErr.Error(), "time out") {
				tt.Fatal(fErr)
			}
			if len(msgs) == 3 {
				break
			}
			<-tick.C
		}
		assert.Len(tt, msgs, 3)
		for _, msg := range msgs {
			msg.Ack()
		}
	})
	t.Run("Invalid request", func(tt *testing.T) {
		expect, h, _, closeFunc := defaultHandler(tt)
		defer h.Close()
		defer closeFunc()
		u := tables.RandomUser()
		assert.Nil(tt, h.Controller.DB.Create(u).Error)
		s := tables.RandomStation(u)
		assert.Nil(tt, h.Controller.DB.Create(s).Error)
		expect.PUT(RegistryRoute).
			WithHeader(XAuthKeyHeader, s.APIKey).
			WithHeader("Content-Type", "application/json").
			WithBytes([]byte("{")).
			Expect().
			Status(http.StatusBadRequest)
	})
}
