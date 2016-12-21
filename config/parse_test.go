package config

import (
	"testing"
)

var confstr = `
key1 = 1
name = adf #adfkfd
[userInfo]
name = wuciyou
email =        		898060380
`

func TestParse(t *testing.T) {
	parse("./default.ini")

	confMap := GetAll()
	for k, v := range confMap {
		t.Logf("confMap key[%s], value[%s] ", k, v)
	}

}
