package main

import (
	"crypto/tls"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/marthjod/ansible-one-inventory/config"
	"github.com/marthjod/ansible-one-inventory/discovery"
	"github.com/marthjod/ansible-one-inventory/filter"
	"github.com/marthjod/ansible-one-inventory/hostnameextractor"
	"github.com/marthjod/gocart/api"
	"github.com/marthjod/gocart/ocatypes"
	"github.com/marthjod/gocart/vmpool"
	flag "github.com/ogier/pflag"
	"net/http"
	"os"
	"path"
)

const (
	configFile = "opennebula-inventory.yaml"
)

func main() {
	var (
		host      = flag.String("host", "", "")
		list      = flag.Bool("list", true, "")
		debug     = flag.Bool("debug", false, "Enable debug output.")
		extractor hostnameextractor.HostnameExtractor
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

	log.Debug("Acquiring hostnames")
	extractor = &hostnameextractor.VmNameExtractor{}
	if conf.HostnameFieldInUserTemplate != "" {
		extractor = &hostnameextractor.UserTemplateExtractor{
			Field: conf.HostnameFieldInUserTemplate,
		}
	}

	hostNames := discovery.GetHostnames(vmPool, extractor)

	groupFilters := conf.StaticGroupFilters

	if conf.DynamicGroupFilters.Pattern != "" {
		fqdnExtractor := func(vm *ocatypes.Vm) string {
			h, _ := vm.UserTemplate.Items.GetCustom(conf.HostnameFieldInUserTemplate)
			return h
		}
		distinctPatterns := vmPool.GetDistinctVmNamePatternsExtractHostname(
			conf.DynamicGroupFilters.Pattern, conf.DynamicGroupFilters.Prefix, conf.DynamicGroupFilters.Infix,
			conf.DynamicGroupFilters.Suffix, fqdnExtractor)

		for pattern, _ := range distinctPatterns {
			groupFilters[filter.AdjustPatternName(pattern, conf.DynamicGroupFilters.PatternReplace)] = pattern
		}
	}

	inventory := discovery.GetInventoryGroups(hostNames, groupFilters)

	if *host != "" {
		log.Error("--host option not implemented yet")
		os.Exit(1)
	} else if *list {
		log.Debug("Writing inventory")
		fmt.Println(inventory)
	}

}
