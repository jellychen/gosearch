package dominate

import (
	"github.com/seefan/gossdb"
	"github.com/seefan/gossdb/conf"
)

type Center struct {
	config         *Conf
	ssdbConnectors *gossdb.Connectors
}

func NewCenter() *Center {
	return &Center{nil, nil}
}

func (self *Center) LoadConf(file string) bool {
	config, succes := ConfLoad(file)
	if succes {
		self.config = config
		pool, err := gossdb.NewPool(&conf.Config{
			Host:             self.config.SSDB_Host,
			Port:             self.config.SSDB_Port,
			MinPoolSize:      self.config.SSDB_MinPoolSize,
			MaxPoolSize:      self.config.SSDB_MaxPoolSize,
			AcquireIncrement: self.config.SSDB_AcquireIncrement,
		})

		if nil == err {
			self.ssdbConnectors = pool
		}
	}
	return succes
}

func (self *Center) SSDBPool() *gossdb.Connectors {
	return self.ssdbConnectors
}
