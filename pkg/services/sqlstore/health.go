package sqlstore

import (
	"github.com/wviveiro/grafana/pkg/bus"
	m "github.com/wviveiro/grafana/pkg/models"
)

func init() {
	bus.AddHandler("sql", GetDBHealthQuery)
}

func GetDBHealthQuery(query *m.GetDBHealthQuery) error {
	return x.Ping()
}
