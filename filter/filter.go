// Package filter handles group assignment via hostname filtering.
package filter

import (
	"fmt"
	"github.com/marthjod/ansible-one-inventory/model"
	"regexp"
)

// A GroupFilter is a key-value map mapping group names to regexps
// which describe hostnames belonging to a respective group.
type GroupFilters map[string]string

// Takes a list of hostnames and returns an InventoryGroup
// consisting of those matching regex.
func Filter(hostNames []string, regex string) (model.InventoryGroup, error) {

	ig := model.InventoryGroup{}
	re, err := regexp.Compile(regex)
	if err != nil {
		return ig, fmt.Errorf("%q: %s", regex, err.Error())
	}

	for _, h := range hostNames {
		if re.MatchString(h) {
			ig = append(ig, h)
		}
	}

	return ig, nil
}

func AdjustPatternName(regex, replacement string) string {
	re := regexp.MustCompile(replacement)
	return re.ReplaceAllString(regex, "")
}
