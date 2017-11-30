package provisioning

import (
	"github.com/wviveiro/grafana/pkg/log"
	"github.com/wviveiro/grafana/pkg/services/provisioning/datasources"
)

var (
	logger log.Logger = log.New("services.provisioning")
)

func StartUp(datasourcePath string) error {
	return datasources.Provision(datasourcePath)
}
