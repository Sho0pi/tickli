package types

import (
	"encoding/json"
	"github.com/gookit/color"
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
