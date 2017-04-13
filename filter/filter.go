package filter

import (
	"fmt"
	"github.com/marthjod/ansible-one-inventory/model"
	"github.com/marthjod/gocart/ocatypes"
	"github.com/marthjod/gocart/vmpool"
	"regexp"
)

type GroupFilters map[string]string

func Filter(
	pool *vmpool.VmPool,
	hostNameExtractor func(*ocatypes.Vm) (string, error),
	regex string) (model.InventoryGroup, error) {

	ig := model.InventoryGroup{}
	re, err := regexp.Compile(regex)
	if err != nil {
		return ig, fmt.Errorf("%q: %s", regex, err.Error())
	}

	for _, vm := range pool.Vms {
		hostName, err := hostNameExtractor(vm)
		if err != nil {
			return ig, err
		}
		if re.MatchString(hostName) {
			ig = append(ig, hostName)
		}
	}

	return ig, nil
}
