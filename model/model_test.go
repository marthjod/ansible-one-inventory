package model_test

import (
	"github.com/marthjod/ansible-one-inventory/model"
	"testing"
)

func TestJson(t *testing.T) {
	inv := model.Inventory{
		"web": model.InventoryGroup{
			"webserver-01",
			"webserver-02",
		},
	}

	expected := `{
  "web": [
    "webserver-01",
    "webserver-02"
  ]
}`
	actual := inv.Json()
	if actual != expected {
		t.Errorf("String() did no render correctly: Expected:\n%q\nActual:\n%q", expected, actual)
	}
}
