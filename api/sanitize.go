package api

import (
	"encoding/json"
	"github.com/microcosm-cc/bluemonday"
)

var sanitizer *bluemonday.Policy

func init() {
	sanitizer = bluemonday.UGCPolicy()
}

func SanitizeJSON(data interface{}) ([]byte, error) {
	switch d := data.(type) {
	case map[string]interface{}:
		for k, v := range d {
			switch v.(type) {
			case string:
				d[k] = sanitizer.Sanitize(v.(string))
			case map[string]interface{}:
				SanitizeJSON(v)
			case []interface{}:
				SanitizeJSON(v)
			}
		}
	case []interface{}:
		if len(d) > 0 {
			switch d[0].(type) {
			case string:
				for i, s := range d {
					d[i] = sanitizer.Sanitize(s.(string))
				}
			case map[string]interface{}:
				for _, t := range d {
					SanitizeJSON(t)
				}
			case []interface{}:
				for _, t := range d {
					SanitizeJSON(t)
				}
			}
		}
	}
	return json.MarshalIndent(data, "", "    ")
}