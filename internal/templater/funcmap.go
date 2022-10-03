package templater

import "strings"

var (
	FuncMap = map[string]interface{}{
		"ToUpper": toUpper,
	}
)

// toUpper is the local function makes strings.ToUpper() available as a callable function
// within HTML templates via the FuncMap var. It can be called within an HTML template
// like so: {{ ToUpper "foo" }}
func toUpper(s string) string {
	return strings.ToUpper(s)
}
