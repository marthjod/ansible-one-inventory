package filter_test

import (
	"github.com/marthjod/ansible-one-inventory/filter"
	"github.com/marthjod/ansible-one-inventory/model"
	"reflect"
	"testing"
)

var expected = []struct {
	pattern     string
	replacement string
	result      string
}{
	{"web-staging-east", "-(we|ea)st", "web-staging"},
	{"db-production-west", "-(we|ea)st", "db-production"},
}

var hostnames = []string{
	"web-staging-east",
	"db-staging-east",
	"db-production-west",
}

func TestAdjustPatternName(t *testing.T) {
	for _, e := range expected {
		actual := filter.AdjustPatternName(e.pattern, e.replacement)
		if actual != e.result {
			t.Errorf("Unexpected adjustment result %q for %q, %q.", actual, e.pattern, e.replacement)
		}
	}
}

func TestFilterErr(t *testing.T) {
	var err error

	regex := "(invalid"
	_, err = filter.Filter(hostnames, regex)
	if err == nil {
		t.Error("No error returned for invalid regexp!")
	}
}

func TestFilter(t *testing.T) {
	expected := model.InventoryGroup{
		"web-staging-east",
	}
	regex := "^[a-z]{3}"

	ig, err := filter.Filter(hostnames, regex)
	if err != nil {
		t.Fatal(err.Error())
	}
	if !reflect.DeepEqual(ig, expected) {
		t.Errorf("Got unexpected result while filtering. Expected:\n%+v\nActual:\n%+v", expected, ig)
	}
}
