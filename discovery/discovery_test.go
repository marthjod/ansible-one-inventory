package discovery_test

import (
	"encoding/xml"
	"github.com/marthjod/ansible-one-inventory/discovery"
	"github.com/marthjod/ansible-one-inventory/filter"
	"github.com/marthjod/ansible-one-inventory/hostnameextractor"
	"github.com/marthjod/ansible-one-inventory/model"
	"github.com/marthjod/gocart/ocatypes"
	"github.com/marthjod/gocart/vmpool"
	"reflect"
	"testing"
)

var (
	vmPool = &vmpool.VmPool{
		Vms: []*ocatypes.Vm{
			{
				Name: "first-vm",
				UserTemplate: ocatypes.UserTemplate{
					Items: ocatypes.Tags{
						ocatypes.Tag{
							XMLName: xml.Name{
								Local: "CUSTOM_FQDN",
							},
							Content: "vm-01.example.com",
						},
					},
				},
			},
			{
				Name: "second-vm",
				UserTemplate: ocatypes.UserTemplate{
					Items: ocatypes.Tags{
						ocatypes.Tag{
							XMLName: xml.Name{
								Local: "CUSTOM_FQDN",
							},
							Content: "vm-02.example.com",
						},
					},
				},
			},
			{
				Name: "web-staging-west",
			},
		},
	}
	vmNameExtractor       = hostnameextractor.VmNameExtractor{}
	userTemplateExtractor = hostnameextractor.UserTemplateExtractor{
		Field: "CUSTOM_FQDN",
	}
)

func TestGetHostnamesVmNameExtractor(t *testing.T) {
	expected := []string{
		"first-vm",
		"second-vm",
		"web-staging-west",
	}

	actual := discovery.GetHostnames(vmPool, &vmNameExtractor)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Unexpected return value. Expected:\n%+v\nActual:\n%+v", expected, actual)
	}
}

func TestGetHostnamesUserTemplateExtractor(t *testing.T) {
	expected := []string{
		"vm-01.example.com",
		"vm-02.example.com",
		"",
	}

	actual := discovery.GetHostnames(vmPool, &userTemplateExtractor)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Unexpected return value. Expected:\n%+v\nActual:\n%+v", expected, actual)
	}
}

func TestGetInventoryGroups(t *testing.T) {
	expected := &model.Inventory{
		"webservers": []string{
			"web-staging-east",
		},
		"westcoast": []string{
			"db-production-west",
		},
	}
	hostNames := []string{
		"web-staging-east",
		"db-production-west",
	}
	filters := filter.GroupFilters{
		"webservers": "^web",
		"westcoast":  "west$",
	}

	actual := discovery.GetInventoryGroups(hostNames, filters)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Unexpected return value. Expected:\n%s\nActual:\n%s", expected, actual)
	}
}

func TestGetInventoryGroupsErr(t *testing.T) {
	expected := &model.Inventory{}
	hostNames := []string{
		"web-staging-east",
	}
	filters := filter.GroupFilters{
		"foo": "(invalid regexp",
	}

	actual := discovery.GetInventoryGroups(hostNames, filters)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Unexpected return value. Expected:\n%s\nActual:\n%s", expected, actual)
	}
}
