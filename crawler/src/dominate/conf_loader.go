package dominate

import (
	"encoding/json"
	"io/ioutil"
)

type Conf struct {
	MaxRoutineNum         int
	SSDB_Host             string
	SSDB_Port             int
	SSDB_MinPoolSize      int
	SSDB_MaxPoolSize      int
	SSDB_AcquireIncrement int
}

const (
	_max_routune_num_k_        string = "max_routune_num"
	_ssdb_host_k_              string = "ssdb_host"
	_ssdb_port_k_              string = "ssdb_port"
	_ssdb_min_pool_size_k_     string = "ssdb_min_pool_size"
	_ssdb_max_pool_size_k_     string = "ssdb_max_pool_size"
	_ssdb_acquire_increment_k_ string = "ssdb_acquire_increment"
)

func ConfLoad(file string) (*Conf, bool) {
	data, err := ioutil.ReadFile(file)
	if nil != err {
		return nil, false
	}

	var json_object interface{}
	err = json.Unmarshal(data, &json_object)
	if nil != err {
		return nil, false
	}

	map_object, success := json_object.(map[string]interface{})
	if !success || nil == map_object {
		return nil, false
	}

	var value interface{} = nil
	conf := &Conf{100, "", 0, 5, 50, 0}

	// max routune num
	value, success = map_object[_max_routune_num_k_]
	if success {
		val, success := Interface2Number(value)
		if success {
			conf.MaxRoutineNum = val
		}
	}

	// ssdb host
	value, success = map_object[_ssdb_host_k_]
	if success {
		val, success := Interface2String(value)
		if success {
			conf.SSDB_Host = val
		}
	}

	// ssdb port
	value, success = map_object[_ssdb_port_k_]
	if success {
		val, success := Interface2Number(value)
		if success {
			conf.SSDB_Port = val
		}
	}

	// ssdb min pool size
	value, success = map_object[_ssdb_min_pool_size_k_]
	if success {
		val, success := Interface2Number(value)
		if success {
			conf.SSDB_MinPoolSize = val
		}
	}

	// ssdb max pool size
	value, success = map_object[_ssdb_max_pool_size_k_]
	if success {
		val, success := Interface2Number(value)
		if success {
			conf.SSDB_MaxPoolSize = val
		}
	}

	// ssdb acquire increment
	value, success = map_object[_ssdb_acquire_increment_k_]
	if success {
		val, success := Interface2Number(value)
		if success {
			conf.SSDB_AcquireIncrement = val
		}
	}

	return conf, true
}
