package hostnameextractor

import (
	"fmt"
	"github.com/marthjod/gocart/ocatypes"
)

type HostnameExtractor interface {
	Extract(vm *ocatypes.Vm) (string, error)
}

type VmNameExtractor struct{}

type UserTemplateExtractor struct {
	Field string
}

func (vne *VmNameExtractor) Extract(vm *ocatypes.Vm) (string, error) {
	return vm.Name, nil
}

func (ute *UserTemplateExtractor) Extract(vm *ocatypes.Vm) (string, error) {
	if ute.Field != "" {
		return vm.UserTemplate.Items.GetCustom(ute.Field)
	}
	return "", fmt.Errorf("Field is empty")
}
