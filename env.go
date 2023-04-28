package main

import "github.com/hidromatologia-v2/models/common/config"

type (
	Config struct {
		config.Producer `env:",prefix=MEMPHIS_"`  // Memphis
		config.Postgres `env:",prefix=POSTGRES_"` // POSTGRESQL
	}
)
