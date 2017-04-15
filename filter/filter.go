package filter

import (
	"fmt"
	"github.com/marthjod/ansible-one-inventory/model"
	"regexp"
)

type GroupFilters map[string]string

func Filter(
	hostNames []string,
	regex string) (model.InventoryGroup, error) {

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
