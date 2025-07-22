package web

import (
	"encoding/json"
	"errors"
	"html/template"
	"time"
)

// templateFuncs defines template functions used across all templates
var templateFuncs = template.FuncMap{
	"toJSON": func(v any) template.JS {
		data, _ := json.Marshal(v)
		return template.JS(data)
	},
	"dict": dict,
	"html": func(s string) template.HTML {
		return template.HTML(s)
	},
	"parseJSON": func(s string) (interface{}, error) {
		var data interface{}
		err := json.Unmarshal([]byte(s), &data)
		return data, err
	},
	"safeURL": func(s string) template.URL {
		return template.URL(s)
	},
	"formatDate": func(dateStr string) string {
		// Handle empty or invalid dates
		if dateStr == "" {
			return ""
		}

		// Parse ISO date format (YYYY-MM-DD)
		parsedDate, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			// If parsing fails, return original string
			return dateStr
		}

		// Format as "Month Day, Year" (e.g., "July 12, 2024")
		return parsedDate.Format("January 2, 2006")
	},
}

// dict creates a map from key-value pairs for use in templates
func dict(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, errors.New("dict requires even number of arguments")
	}
	m := make(map[string]interface{})
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, errors.New("dict keys must be strings")
		}
		m[key] = values[i+1]
	}
	return m, nil
}