package filter

import (
	"regexp"
	"github.com/marthjod/gocart/ocatypes"
	"github.com/marthjod/ansoneinv/model"
	"github.com/marthjod/gocart/vmpool"
)

type GroupFilters map[string]string

func Filter(pool *vmpool.VmPool, hostNameExtractor func(*ocatypes.Vm) (string, error), regex string) model.InventoryGroup {
	ig := model.InventoryGroup{}
	re := regexp.MustCompile(regex)

	for _, vm := range pool.Vms {
		hostName, err := hostNameExtractor(vm)
		if err != nil {
			continue
		}
		if re.MatchString(hostName) {
			ig = append(ig, hostName)
		}
	}

	return ig
}
