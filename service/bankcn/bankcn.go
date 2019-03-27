package bankcn

import (
	"github.com/wzshiming/bankcn"
)

// BankcnService #path:"/bankcn/"#
type BankcnService struct {
}

// NewBankcnService Create a new BankcnService
func NewBankcnService() (*BankcnService, error) {
	return &BankcnService{}, nil
}

// Banks #route:"GET /banks"# 获取全部的银行
func (s *BankcnService) Banks() (banks map[string]string, err error) {
	return bankcn.BankMap, nil
}

// Banks #route:"GET /banks/{bank}.jpg"# 获取某地区的银行图标
func (s *BankcnService) BanksJPG(bank string /* #name:"bank"# */) (file []byte /* #content:"image/jpeg"# */, err error /* #code:"404"# */) {
	return bankcn.Asset("icon/" + bank + ".jpg")
}

// Get #route:"GET /{bank}/{area_id}"# 获取某地区的银行数据
func (s *BankcnService) Get(bank /* #name:"bank"# */, areaID string /* #name:"area_id"# */) (banks []*bankcn.Bank, err error) {
	return bankcn.Get(bank, areaID), nil
}

// Verify #route:"GET /{bank_id}"# 获取银行卡属性
func (s *BankcnService) Verify(bankID string /* #name:"bank_id"# */) (valid *bankcn.Valid, err error) {
	return bankcn.Verify(bankID)
}
