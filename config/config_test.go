package config_test

import (
	"github.com/marthjod/ansible-one-inventory/config"
	"github.com/marthjod/ansible-one-inventory/discovery"
	"github.com/marthjod/ansible-one-inventory/filter"
	"reflect"
	"testing"
)

var expected = &config.Config{
	Username:                    "oneadmin",
	Password:                    "opennebula",
	Url:                         "http://one-sandbox:2633/RPC2",
	SslSkipVerify:               false,
	HostnameFieldInUserTemplate: "",
	StaticGroupFilters: filter.GroupFilters{
		"all":      ".",
		"web":      "web",
		"database": "db",
		"app":      "app",
		"linux":    "[Ll]inux",
	},
	DynamicGroupFilters: discovery.AutodiscoveryConfig{
		Pattern:        "^foo-([a-z]{3}).*(example)$",
		Prefix:         "foo-",
		Infix:          ".+",
		Suffix:         "$",
		PatternReplace: "(foo|[-.+$])",
	},
}

func TestFromFile(t *testing.T) {
	path := "testdata/opennebula-inventory.example.yaml"

	actual, err := config.FromFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Config file %q not loaded correctly. Expected:\n%+v\nActual:\n%+v", path, expected, actual)
	}
}

func TestFromFileErr(t *testing.T) {
	var (
		err error
	)

	path := "doesnotexist"
	_, err = config.FromFile(path)
	if err == nil {
		t.Error("No error returned for non-existing config file!")
	}

	path = "testdata/invalid-config.yaml"
	cfg, err := config.FromFile(path)
	if err == nil {
		t.Error("No error returned for invalid config file!")
	}
	if cfg != nil {
		t.Error("Partial config struct returned for invalid config file!")
	}
}
