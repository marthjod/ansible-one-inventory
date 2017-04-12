package config

import (
	"io/ioutil"
	"encoding/json"
	"github.com/marthjod/ansible-one-inventory/filter"
)

type Config struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	Url           string `json:"url"`
	SslSkipVerify bool `json:"skip_ssl_verify"`
	HostnameField string `json:"hostname_field"`
	GroupFilters  filter.GroupFilters `json:"group_filters"`
}

func FromFile(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	c := Config{}
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, err
	}

	return &c, nil
}