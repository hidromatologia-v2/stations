package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/hidromatologia-v2/models/common/postgres"
	"github.com/hidromatologia-v2/models/tables"
	"github.com/hidromatologia-v2/stations/handler"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "stations",
		Usage: "stations microservice for the ResupplyOrg project",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "url",
				Value: "http://127.0.0.1:5000/api",
				Usage: "Stations URL",
			},
		},
		Action: func(ctx *cli.Context) error {
			db := postgres.NewDefault()
			u := tables.RandomUser()
			db.Create(u)
			s := tables.RandomStation(u)
			db.Create(s)
			sensor := s.Sensors[0]
			a := tables.RandomAlert(u, &sensor)
			*a.Condition = tables.Ge
			*a.Value = 10
			a.Enabled = &tables.True
			db.Create(a)
			var registriesBuffer bytes.Buffer
			json.NewEncoder(&registriesBuffer).Encode([]tables.SensorRegistry{
				{
					SensorUUID: sensor.UUID,
					Value:      10,
				},
			})
			req, _ := http.NewRequest(http.MethodGet, ctx.String("url")+handler.StationRoute, nil)
			req.Header.Add(handler.XAuthKeyHeader, s.APIKey)
			res, _ := http.DefaultClient.Do(req)
			defer res.Body.Close()
			fmt.Println(res.Status)
			contents, _ := io.ReadAll(res.Body)
			fmt.Println(string(contents))
			req, _ = http.NewRequest(http.MethodPut, ctx.String("url")+handler.RegistryRoute, bytes.NewReader(registriesBuffer.Bytes()))
			req.Header.Add(handler.XAuthKeyHeader, s.APIKey)
			req.Header.Add("Content-Type", "application/json")
			res, _ = http.DefaultClient.Do(req)
			defer res.Body.Close()
			fmt.Println(res.Status)
			contents, _ = io.ReadAll(res.Body)
			fmt.Println(string(contents))
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
