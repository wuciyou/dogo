package config

import (
	"testing"
)

func TestParse(t *testing.T) {
	confMap := map[string]string{
		"LISTEN_PORT":        "8080",
		"AJAX_RETURN_FORMAT": "json",
		"SERVER_NAME":        "DoGoServerv1",
		"LOG.DATA_CHAN_SIZE": "0",
	}
	parse(default_conf)
	parseConfMap := GetAll()
	for k, v := range confMap {
		if cv, ok := parseConfMap[k]; ok {
			if v != cv {
				t.Errorf("expand value error, key:%s, value1:%s, value2:%s ", k, v, cv)
			} else {
				t.Logf("confMap key[%s], value[%s] ", k, v)
			}
		} else {
			t.Errorf("expand value error, no found key:%s ", k)
		}

	}

}
