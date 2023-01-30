package services

import (
	"fmt"
	"strings"
)

func fill(rule Rule, form Form) string {
	var result []string
	// `yzN18KdNSpdc4rGD=3P0Nk9qZ49PDN&PyA8s4ioBxyeNNcL=test&Aj3Z2kKydtydbFei=test`

	for _, t := range form.text {
		result = append(result, fmt.Sprintf("%s=%s", t, rule.actionForText))
	}

	for _, r := range form.radios {
		for k, v := range r {
			var selValue string
			switch rule.actionForRadio {
			case rTagLongest:
				for _, s := range v {
					if len(s) > len(selValue) {
						selValue = s
					}
				}
			}

			result = append(result, fmt.Sprintf("%s=%s", k, selValue))
		}
	}

	for _, r := range form.selects {
		for k, v := range r {
			var selValue string
			switch rule.actionForSelect {
			case rTagLongest:
				for _, s := range v {
					if len(s) > len(selValue) {
						selValue = s
					}
				}
			}

			result = append(result, fmt.Sprintf("%s=%s", k, selValue))
		}
	}

	return strings.Join(result, "&")
}
