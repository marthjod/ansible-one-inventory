package hostnameextractor_test

import (
	"encoding/xml"
	"github.com/marthjod/ansible-one-inventory/hostnameextractor"
	"github.com/marthjod/gocart/ocatypes"
	"testing"
)

var (
	vm = ocatypes.Vm{
		Name: "vm-database-production",
		UserTemplate: ocatypes.UserTemplate{
			Items: ocatypes.Tags{
				ocatypes.Tag{
					XMLName: xml.Name{
						Local: "CUSTOM_FQDN",
					},
					Content: "db-live.example.com",
				},
			},
		},
	}
	anonymousVm           = ocatypes.Vm{}
	vmNameExtractor       = hostnameextractor.VmNameExtractor{}
	userTemplateExtractor = hostnameextractor.UserTemplateExtractor{
		Field: "CUSTOM_FQDN",
	}
)

func TestVmNameExtractor_Extract(t *testing.T) {
	expected := vm.Name
	actual := vmNameExtractor.Extract(&vm)
	if actual != expected {
		t.Errorf("Extracting VM name failed. Expected:\n%+v\nActual:\n%+v", expected, actual)
	}
}

func TestVmNameExtractor_ExtractErr(t *testing.T) {
	expected := ""
	actual := vmNameExtractor.Extract(&anonymousVm)
	if actual != expected {
		t.Error("No empty string for anonymous VM returned!")
	}
}

func TestUserTemplateExtractor_Extract(t *testing.T) {
	expected := "db-live.example.com"
	actual := userTemplateExtractor.Extract(&vm)
	if actual != expected {
		t.Errorf("Extracting VM name failed. Expected:\n%+v\nActual:\n%+v", expected, actual)
	}
}

func TestUserTemplateExtractor_ExtractEmpty(t *testing.T) {
	expected := ""
	noFieldGiven := hostnameextractor.UserTemplateExtractor{}
	actual := noFieldGiven.Extract(&vm)
	if actual != expected {
		t.Errorf("Extracting VM name failed. Expected:\n%+v\nActual:\n%+v", expected, actual)
	}
}
