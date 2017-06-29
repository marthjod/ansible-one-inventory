package model

import (
	"encoding/json"
)

type Inventory map[string]InventoryGroup

type InventoryGroup []string

func (i Inventory) String() string {
	j, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(j)
}
