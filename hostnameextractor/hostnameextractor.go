package hostnameextractor

import (
	"github.com/marthjod/gocart/ocatypes"
)

type HostnameExtractor interface {
	Extract(vm *ocatypes.Vm) string
}

type VmNameExtractor struct{}

type UserTemplateExtractor struct {
	Field string
}

func (vne *VmNameExtractor) Extract(vm *ocatypes.Vm) string {
	return vm.Name
}

func (ute *UserTemplateExtractor) Extract(vm *ocatypes.Vm) string {
	if ute.Field != "" {
		custom, _ := vm.UserTemplate.Items.GetCustom(ute.Field)
		return custom
	}
	return ""
}
