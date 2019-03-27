package bankcn

import (
	"encoding/json"
	"errors"

	"github.com/wzshiming/requests"
	"gopkg.in/ffmt.v1"
)

var req = requests.NewClient().NewRequest().
	SetURLByStr(`https://ccdcapi.alipay.com/validateAndCacheCardInfo.json?_input_charset=utf-8&cardBinCheck=true`)

func Verify(bankID string) (v *Valid, err error) {
	resp, err := req.Clone().
		SetQuery("cardNo", bankID).
		Do()
	if err != nil {
		return nil, err
	}

	vv := valid{}
	ffmt.Mark(string(resp.Body()))
	err = json.Unmarshal(resp.Body(), &vv)
	if err != nil {
		return nil, err
	}
	if !vv.Validated {
		return nil, errors.New("这是无效的银行卡")
	}

	return &Valid{
		BankIDType: vv.CardType,
		Bank:       vv.Bank,
	}, nil
}

type valid struct {
	CardType  string `json:"cardType"`
	Bank      string `json:"bank"`
	Validated bool   `json:"validated"`
}

type Valid struct {
	// 银行卡 类型
	BankIDType string `json:"bank_id_type"`
	// 银行卡所属银行
	Bank string `json:"bank"`
}
