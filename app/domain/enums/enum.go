package enums

import (
	"encoding/json"
	"fmt"
)

type TypeCampaign int

const (
	GoogleBlue TypeCampaign = iota
	GoogleGreen
	AppleStore
)

var typeMap = map[string]TypeCampaign{
	"GoogleBlue":  GoogleBlue,
	"GoogleGreen": GoogleGreen,
	"AppleStore":  AppleStore,
}

func (t *TypeCampaign) UnmarshalJSON(data []byte) error {
	var typeStr string
	if err := json.Unmarshal(data, &typeStr); err != nil {
		return fmt.Errorf("TypeCampaign должно быть строкой, получено: %s", data)
	}

	type_, exists := typeMap[typeStr]
	if !exists {
		return fmt.Errorf("недопустимое значение TypeCampaign: %q", typeStr)
	}
	*t = type_
	return nil
}

func (t TypeCampaign) Types() string {
	return [...]string{"GoogleBlue", "GoogleGreen", "AppleStore"}[t]
}

func (t TypeCampaign) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Types())
}

type Design int

const (
	Dark Design = iota
	Light
)

var designMap = map[string]Design{
	"Dark":  Dark,
	"Light": Light,
}

func (d *Design) UnmarshalJSON(data []byte) error {
	var designStr string
	if err := json.Unmarshal(data, &designStr); err != nil {
		return fmt.Errorf("Design должно быть строкой, получено: %s", data)
	}

	design, exists := designMap[designStr]
	if !exists {
		return fmt.Errorf("недопустимое значение Design: %q", designStr)
	}
	*d = design
	return nil
}

func (d Design) Designs() string {
	return [...]string{"Dark", "Light"}[d]
}
