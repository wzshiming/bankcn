package bankcn

import (
	"strings"

	"github.com/wzshiming/bankcn"
)

// BankcnService #path:"/bankcn/"#
type BankcnService struct {
}

// NewBankcnService Create a new BankcnService
func NewBankcnService() (*BankcnService, error) {
	return &BankcnService{}, nil
}

// Banks #route:"GET /banks"# 获取全部的银行类型的列表
func (s *BankcnService) Banks() (bankItems []*bankcn.BankItem, err error) {
	return bankcn.BankItems, nil
}

// Banks #route:"GET /banks/{bank}.jpg"# 获取某银行图标
func (s *BankcnService) BanksJPG(bank string /* #name:"bank"# */) (file []byte /* #content:"image/jpeg"# */, err error /* #code:"404"# */) {
	return bankcn.Asset("icon/" + bank + ".jpg")
}

// Get #route:"GET /banks/{bank}/{area_id}"# 获取某地区的银行数据
func (s *BankcnService) Get(bank /* #name:"bank"# */, areaID string /* #name:"area_id"# */, filter string) (banks []*bankcn.Bank, err error) {
	banks = bankcn.Get(bank, areaID)
	if filter == "" {
		return banks, nil
	}
	banks0 := []*bankcn.Bank{}
	for _, bank := range banks {
		if strings.Index(bank.Address, filter) != -1 {
			banks0 = append(banks0, bank)
		}
	}
	return banks0, nil
}

// Verify #route:"GET /bank_id/{bank_id}"# 获取银行卡号的 所属银行 以及 类型
func (s *BankcnService) Verify(bankID string /* #name:"bank_id"# */) (valid *bankcn.Valid, err error) {
	return bankcn.Verify(bankID)
}
