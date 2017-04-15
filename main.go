package main

import (
	"crypto/tls"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/marthjod/ansible-one-inventory/config"
	"github.com/marthjod/ansible-one-inventory/filter"
	"github.com/marthjod/ansible-one-inventory/model"
	"github.com/marthjod/gocart/api"
	"github.com/marthjod/gocart/ocatypes"
	"github.com/marthjod/gocart/vmpool"
	flag "github.com/ogier/pflag"
	"github.com/orcaman/concurrent-map"
	"net/http"
	"os"
	"sync"
	"path"
)

const (
	configFile = "opennebula-inventory.yaml"
)

func main() {
	var (
		host  = flag.String("host", "", "")
		list  = flag.Bool("list", true, "")
		debug = flag.Bool("debug", false, "Enable debug output.")
		w     sync.WaitGroup
	)
	flag.Parse()

	if *debug {
		log.SetLevel(log.DebugLevel)
	}

	confPath := path.Dir(os.Args[0]) + "/" + configFile
	log.Debug("Loading config ", confPath)
	conf, err := config.FromFile(confPath)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	log.Debug("Accessing API at ", conf.Url)
	apiClient, err := api.NewClient(conf.Url, conf.Username, conf.Password, &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: conf.SslSkipVerify},
	})
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	log.Debug("Fetching VM pool")
	vmPool := vmpool.NewVmPool()
	if err := apiClient.Call(vmPool); err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	hostnameExtractor := func(vm *ocatypes.Vm) (string, error) {
		return vm.Name, nil
	}
	if conf.HostnameFieldInUserTemplate != "" {
		hostnameExtractor = func(vm *ocatypes.Vm) (string, error) {
			return vm.UserTemplate.Items.GetCustom(conf.HostnameFieldInUserTemplate)
		}
	}

	tempMap := cmap.New()

	inventory := model.Inventory{}
	for name, pattern := range conf.GroupFilters {
		w.Add(1)
		go func(vmPool *vmpool.VmPool, name, pattern string) {
			log.Debugf("Populating filter group %q", name)
			inventoryGroup, err := filter.Filter(vmPool, hostnameExtractor, pattern)
			if err != nil {
				log.Error(err.Error())
			} else {
				tempMap.Set(name, inventoryGroup)
			}

			w.Done()
		}(vmPool, name, pattern)
	}

	w.Wait()

	if *host != "" {
		log.Error("--host option not implemented yet")
		os.Exit(1)
	} else if *list {
		log.Debug("Writing inventory")
		for k, v := range tempMap.Items() {
			inventory[k] = v.(model.InventoryGroup)
		}

		fmt.Println(inventory)
	}

}
