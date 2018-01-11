package dataprovider

import (
	"fmonitor/conf"
)

type DataProvider interface {
	SaveMemoryInfo(server string, used int, peak int) int64
	SaveInfoCommand(server string, info map[string]string) int64
}

func NewProvider(config *conf.Config) DataProvider {
	if config.Datatype == "sqlite" {
		return new(SqliteProvide)
	}
	return nil
}
