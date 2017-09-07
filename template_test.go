package dogo

import (
	"testing"
)

func TestParser(t *testing.T) {

	template := NewTemplate("./test_data/base.html")
	t.Logf("template str \n %s \n", string(template.data))
}
