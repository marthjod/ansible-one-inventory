// Package model provides definitions for inventory data structures.
package model

import (
	"encoding/json"
)

// An inventory consists of key-mapped InventoryGroups.
type Inventory map[string]InventoryGroup

// An InventoryGroup consists of hostnames sharing a common attribute or having a similar role.
// Hosts may belong to several groups. A webserver in production may belong to the webservers
// inventory group and the production inventory group, for example.
type InventoryGroup []string

// Renders Inventory structs as JSON.
func (i Inventory) Json() string {
	j, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(j)
}
