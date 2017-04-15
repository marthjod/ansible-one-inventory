package config

import (
	"github.com/marthjod/ansible-one-inventory/discovery"
	"github.com/marthjod/ansible-one-inventory/filter"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Username                    string                        `yaml:"username"`
	Password                    string                        `yaml:"password"`
	Url                         string                        `yaml:"url"`
	SslSkipVerify               bool                          `yaml:"skip_ssl_verify"`
	HostnameFieldInUserTemplate string                        `yaml:"hostname_field_in_user_template"`
	StaticGroupFilters          filter.GroupFilters           `yaml:"static_group_filters"`
	DynamicGroupFilters         discovery.AutodiscoveryConfig `yaml:"dynamic_group_filters"`
}

func FromFile(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	c := Config{}
	if err := yaml.Unmarshal(data, &c); err != nil {
		return nil, err
	}

	return &c, nil
}
