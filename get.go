package bankcn

import (
	"encoding/json"
)

type Bank struct {
	BankID      string `json:"bank_id,omitempty"`
	Name        string `json:"name,omitempty"`
	BankUnionID string `json:"bank_union_id,omitempty"`
	Address     string `json:"address,omitempty"`
	Phone       string `json:"phone,omitempty"`
	AreaID      string `json:"area_id,omitempty"`
}

var Banks []*Bank

var AreaMapBank = map[string]map[string][]*Bank{}

func init() {
	data := MustAsset("banks.json")
	json.Unmarshal(data, &Banks)

	for _, bank := range Banks {

		if AreaMapBank[bank.BankID] == nil {
			AreaMapBank[bank.BankID] = map[string][]*Bank{}
		}
		areaID := bank.AreaID
		switch len(areaID) {
		case 4:
			AreaMapBank[bank.BankID][areaID[:2]] = append(AreaMapBank[bank.BankID][areaID[:2]], bank)
			fallthrough
		case 2:
			AreaMapBank[bank.BankID][areaID] = append(AreaMapBank[bank.BankID][areaID], bank)
		}

	}
}

// Get 根据 areaID 获取当前区域下所有支行
func Get(bankID string, areaID string) []*Bank {
	b := AreaMapBank[bankID]
	if b == nil {
		return nil
	}
	return b[areaID]
}
