package discovery

import (
	log "github.com/Sirupsen/logrus"
	"github.com/marthjod/ansible-one-inventory/filter"
	"github.com/marthjod/ansible-one-inventory/hostnameextractor"
	"github.com/marthjod/gocart/vmpool"
	"github.com/orcaman/concurrent-map"
	"sync"

	"github.com/marthjod/ansible-one-inventory/model"
)

type AutodiscoveryConfig struct {
	Pattern        string `yaml:"pattern"`
	Prefix         string `yaml:"prefix"`
	Infix          string `yaml:"infix"`
	Suffix         string `yaml:"suffix"`
	PatternReplace string `yaml:"pattern_replace"`
}

func GetHostnames(pool *vmpool.VmPool, extractor hostnameextractor.HostnameExtractor) []string {
	var hostNames = []string{}
	for _, vm := range pool.Vms {
		h := extractor.Extract(vm)
		hostNames = append(hostNames, h)
	}

	return hostNames
}

func GetInventoryGroups(hostNames []string, filters filter.GroupFilters) *model.Inventory {
	var (
		w         sync.WaitGroup
		tempMap   = cmap.New()
		inventory = model.Inventory{}
	)

	for groupName, pattern := range filters {
		w.Add(1)
		go func(hostNames []string, groupName, pattern string) {
			log.Debugf("Populating filter group %q", groupName)
			inventoryGroup, err := filter.Filter(hostNames, pattern)
			if err != nil {
				log.Error(err.Error())
			} else {
				tempMap.Set(groupName, inventoryGroup)
			}

			w.Done()
		}(hostNames, groupName, pattern)
	}

	w.Wait()

	for k, v := range tempMap.Items() {
		inventory[k] = v.(model.InventoryGroup)
	}

	return &inventory
}
