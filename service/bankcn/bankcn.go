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

// Banks #route:"GET /banks"# 获取某地区的银行数据
func (s *BankcnService) Banks() (banks map[string]string, err error) {
	return bankcn.BankMap, nil
}

// Banks #route:"GET /banks/{bank_id}.jpg"# 获取某地区的银行数据
func (s *BankcnService) BanksJPG(bankID string /* #name:"bank_id"# */) (file []byte /* #content:"image/jpeg"# */, err error /* #code:"404"# */) {
	return bankcn.Asset("icon/" + bankID + ".jpg")
}

// Get #route:"GET /{bank_id}/{area_id}"# 获取某地区的银行数据
func (s *BankcnService) Get(bankID /* #name:"bank_id"# */, areaID string /* #name:"area_id"# */) (banks []*bankcn.Bank, err error) {
	return bankcn.Get(bankID, areaID), nil
}
