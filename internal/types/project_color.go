package types

import (
	"encoding/json"
	"fmt"
	"github.com/gookit/color"
	"regexp"
	"strings"
)

var DefaultColor = ProjectColor(color.HEX("#3694FE"))

type ProjectColor color.RGBColor

func (c *ProjectColor) UnmarshalJSON(data []byte) error {

	var colorStr string
	if err := json.Unmarshal(data, &colorStr); err != nil {
		return err
	}

	if colorStr == "" {
		*c = DefaultColor
	} else {
		*c = ProjectColor(color.HEX(colorStr))
	}
	return nil
}

func (c ProjectColor) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

func (c ProjectColor) Sprint(a ...any) string {
	return color.RGBColor(c).Sprint(a...)
}

func (c ProjectColor) String() string {
	return "#" + strings.ToUpper(color.RGBColor(c).Hex())
}

func (c *ProjectColor) Set(s string) error {
	// Validate the hex color format
	s = strings.TrimSpace(s)

	// This pattern matches both 3-digit and 6-digit hex colors, with optional "#" prefix
	hexPattern := regexp.MustCompile(`^#?([0-9A-Fa-f]{3}|[0-9A-Fa-f]{6})$`)
	if !hexPattern.MatchString(s) {
		return fmt.Errorf("invalid hex color format: must be a 3 or 6-digit hex color code (e.g., '#F18' or '#F18181')")
	}

	// Add the "#" back if it was missing
	if !strings.HasPrefix(s, "#") {
		s = "#" + s
	}

	*c = ProjectColor(color.HEX(s))
	return nil
}

func (c *ProjectColor) Type() string {
	return "ProjectColor"
}
