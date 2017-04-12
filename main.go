package main

import (
	"crypto/tls"
	flag "github.com/ogier/pflag"
	"fmt"
	"github.com/marthjod/gocart/api"
	"github.com/marthjod/gocart/vmpool"
	"net/http"
	"os"
	"github.com/marthjod/ansible-one-inventory/model"
	"github.com/marthjod/ansible-one-inventory/filter"
	"github.com/marthjod/gocart/ocatypes"
	"github.com/marthjod/ansible-one-inventory/config"
	"sync"
	"github.com/orcaman/concurrent-map"
)

const (
	configFile = "opennebula-inventory.json"
)


func main() {
	var (
		host = flag.String("host", "", "")
		list = flag.Bool("list", true, "")
		w sync.WaitGroup
	)
	flag.Parse()

	conf, err := config.FromFile(configFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	apiClient, err := api.NewClient(conf.Url, conf.Username, conf.Password, &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: conf.SslSkipVerify},
	})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	vmPool := vmpool.NewVmPool()
	if err := apiClient.Call(vmPool); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	tempMap := cmap.New()

	inventory := model.Inventory{}
	for name, pattern := range conf.GroupFilters {
		w.Add(1)
		go func(vmPool *vmpool.VmPool, name, pattern string) {
			tempMap.Set(name, filter.Filter(vmPool, func(vm *ocatypes.Vm) (string, error) {
				return vm.UserTemplate.Items.GetCustom(conf.HostnameField)
			}, pattern))

			w.Done()
		}(vmPool, name, pattern)
	}

	w.Wait()

	if *host != "" {
		fmt.Println("Not implemented yet")
		os.Exit(1)
	} else if *list {
		for k, v := range tempMap.Items() {
			inventory[k] = v.(model.InventoryGroup)
		}
		fmt.Println(inventory)
	}

}
