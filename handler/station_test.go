package handler

import (
	"net/http"
	"testing"

	"github.com/hidromatologia-v2/models/common/random"
	"github.com/hidromatologia-v2/models/tables"
	"github.com/stretchr/testify/assert"
)

func TestStation(t *testing.T) {
	t.Run("Valid", func(tt *testing.T) {
		expect, h, _, closeFunc := defaultHandler(tt)
		defer closeFunc()
		u := tables.RandomUser()
		assert.Nil(tt, h.Controller.DB.Create(u).Error)
		s := tables.RandomStation(u)
		assert.Nil(tt, h.Controller.DB.Create(s).Error)
		expect.GET(StationRoute).
			WithHeader(XAuthKeyHeader, s.APIKey).
			Expect().
			Status(http.StatusOK)
	})
	t.Run("Unauthorized", func(tt *testing.T) {
		expect, _, _, closeFunc := defaultHandler(tt)
		defer closeFunc()
		expect.GET(StationRoute).
			WithHeader(XAuthKeyHeader, random.String()).
			Expect().
			Status(http.StatusUnauthorized)
	})
}
