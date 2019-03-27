package bankcn

import (
	"encoding/json"
)

type Bank struct {
	// 银行标识符
	Bank string `json:"bank,omitempty"`
	// 银行名字
	Name string `json:"name,omitempty"`
	// 银行联行号
	BankUnionID string `json:"bank_union_id,omitempty"`
	// 联系地址
	Address string `json:"address,omitempty"`
	// 联系电话
	Phone string `json:"phone,omitempty"`
	// 所在区域代号
	AreaID string `json:"area_id,omitempty"`
}

var Banks []*Bank

var AreaMapBank = map[string]map[string][]*Bank{}

func init() {
	data := MustAsset("banks.json")
	json.Unmarshal(data, &Banks)

	for _, bank := range Banks {

		if AreaMapBank[bank.Bank] == nil {
			AreaMapBank[bank.Bank] = map[string][]*Bank{}
		}
		areaID := bank.AreaID
		switch len(areaID) {
		case 4:
			AreaMapBank[bank.Bank][areaID[:2]] = append(AreaMapBank[bank.Bank][areaID[:2]], bank)
			fallthrough
		case 2:
			AreaMapBank[bank.Bank][areaID] = append(AreaMapBank[bank.Bank][areaID], bank)
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
