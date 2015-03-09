// Package geo_pattern creates beautiful generative background image
// patterns from a string.
package geo_pattern

import (
	"encoding/base64"
	"fmt"
	"github.com/pravj/geo_pattern/pattern"
)

// Generate returns pattern's SVG string
func Generate(args map[string]string) string {
	p := pattern.New(args)

	return p.SvgStr()
}

// Base64String returns pattern's Base64 encoded string
func Base64String(args map[string]string) string {
	svgStr := Generate(args)
	base64Str := base64.StdEncoding.EncodeToString([]byte(svgStr))

	return base64Str
}

// URIimage returns pattern's uri image string
func URIimage(args map[string]string) string {
	base64Str := Base64String(args)

	return fmt.Sprintf("url(data:image/svg+xml;base64,%s);", base64Str)
}
