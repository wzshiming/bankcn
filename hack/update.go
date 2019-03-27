//+build ignore
package main

import (
	"bytes"
	"encoding/json"
	"image"
	"image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/wzshiming/areacn"
	"github.com/wzshiming/bankcn"
	"github.com/wzshiming/goquery"
	"github.com/wzshiming/requests"
	ffmt "gopkg.in/ffmt.v1"
)

func init() {
	for k, v := range bankcn.BankMap {
		BankRemap[v] = k
	}
}

var BankRemap = map[string]string{}

// http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/

var ua = `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.119 Safari/537.36`

// var host = `http://www.lianhanghao.com/index.php`
var cli = requests.NewClient().
	SetCache(requests.FileCacheDir("./tmp/")).
	//	SetLogLevel(requests.LogMessageAll).
	NewRequest().
	// SetTimeout(time.Second).
	SetUserAgent(ua)

func main() {
	os.MkdirAll("icon", 0755)
	for k, _ := range bankcn.BankMap {
		data, err := getIcon(k)
		if err != nil {
			ffmt.Mark(err)
			return
		}
		err = ioutil.WriteFile("icon/"+k+".jpg", data, 0666)
		if err != nil {
			ffmt.Mark(err)
			return
		}
	}

	banks, err := getBank()
	if err != nil {
		ffmt.Mark(err)
		return
	}

	p, err := getIndex("0")
	if err != nil {
		ffmt.Mark(err)
		return
	}
	bankss := []*Bank{}
	for _, bank := range banks {
		for _, v1 := range p {
			c, err := getIndex(v1.SID)
			if err != nil {
				ffmt.Mark(err)
				return
			}
			// ffmt.Puts(c)

			for _, v2 := range c {
				bk, err := getAll(bank, v1, v2)
				if err != nil {
					ffmt.Mark(err)
					return
				}
				bankss = append(bankss, bk...)
			}
		}
	}

	data, err := json.Marshal(bankss)
	if err != nil {
		ffmt.Mark(err)
		return
	}

	ioutil.WriteFile("banks.json", data, 0666)
}

func getAll(bank *indexBank, province, city *index) (banks []*Bank, err error) {
	for i := 0; ; i++ {
		//time.Sleep(time.Second / 5)
		bank, err := getList(i, bank, province, city)
		if err != nil {
			return nil, err
		}
		if len(bank) == 0 {
			break
		}
		banks = append(banks, bank...)
	}

	sort.Slice(banks, func(i, j int) bool {
		return banks[i].BankUnionID < banks[j].BankUnionID
	})
	return banks, nil
}

func getIndex(id string) (indexs []*index, err error) {
	cli := cli.Clone()
	resp, err := cli.
		SetMethod(requests.MethodGet).
		SetURLByStr(`http://www.lianhanghao.com/index.php/Index/Ajax`).
		SetQuery("id", id).
		Do()
	if err != nil {
		return nil, err
	}

	body := resp.Body()
	body = bytes.TrimPrefix(body, []byte{0xef, 0xbb, 0xbf})

	err = json.Unmarshal(body, &indexs)
	if err != nil {
		return nil, err
	}
	return indexs, nil
}

func getList(page int, bank *indexBank, province, city *index) (banks []*Bank, err error) {
	cli := cli.Clone()
	req := cli.
		SetMethod(requests.MethodGet).
		// /Index/index/p/2/bank/2/province/6/city/72.html
		// /index.php/Index/index/p/1/bank/1/province/19/city/230.html
		SetURLByStr(`http://www.lianhanghao.com/index.php/Index/index/p/{page}/bank/{bank}/province/{province}/city/{city}.html`).
		// SetURLByStr(`http://www.lianhanghao.com/index.php/Index/index/p/{page}.html`).
		SetPath("page", strconv.FormatInt(int64(page+1), 10)).
		SetPath("bank", bank.ID).
		SetPath("province", province.SID).
		SetPath("city", city.SID)
	resp, err := req.
		Do()
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(resp.RawBody())
	if err != nil {
		return nil, err
	}

	doc.Find("tbody tr").EachWithBreak(func(i int, v *goquery.Selection) bool {
		b := &Bank{}
		b.BankUnionID = v.Find("td:nth-child(1)").Text()
		b.Name = v.Find("td:nth-child(2)").Text()
		b.Phone = strings.TrimLeft(v.Find("td:nth-child(3)").Text(), "淘宝网 | 信用卡办理")
		b.Address = strings.TrimLeft(v.Find("td:nth-child(4)").Text(), "手机就是 pos机 无卡支付秒到 0.55%")

		area := areacn.Get(province.AreaID[:2])
		if len(area) != 0 {
			if len(area) == 1 {
				b.AreaID = area[0].AreaID
			} else {
				for _, v := range area {
					if v.Name == city.Name {
						b.AreaID = v.AreaID
					}
				}
			}
		}
		if b.AreaID == "" {
			b.AreaID = province.AreaID[:2]
		}
		// b.Province = province.Name
		// b.City = city.Name
		// b.BankName = bank.Name
		b.BankID = BankRemap[bank.Name]
		if b.BankUnionID != "" {
			banks = append(banks, b)
		}
		return true
	})
	return banks, nil
}

func getBank() (banks []*indexBank, err error) {
	cli := cli.Clone()
	req := cli.
		SetMethod(requests.MethodGet).
		SetURLByStr(`http://www.lianhanghao.com/index.php`)
	resp, err := req.
		Do()
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(resp.RawBody())
	if err != nil {
		return nil, err
	}

	doc.Find("#bank option").EachWithBreak(func(i int, v *goquery.Selection) bool {
		value := v.AttrOr("value", "")
		if value == "" {
			return true
		}
		banks = append(banks, &indexBank{
			ID:   value,
			Name: v.Text(),
		})
		return true
	})
	return banks, nil
}

// Bank
type Bank struct {
	BankID      string `json:"bank_id,omitempty"`
	Name        string `json:"name,omitempty"`
	BankUnionID string `json:"bank_union_id,omitempty"`
	Address     string `json:"address,omitempty"`
	Phone       string `json:"phone,omitempty"`
	AreaID      string `json:"area_id,omitempty"`
}

type indexBank struct {
	ID   string
	Name string
}

type index struct {
	AreaID string `json:"area_id,omitempty"`
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	PID    string `json:"pid,omitempty"`
	SID    string `json:"sid,omitempty"`
}

func getIcon(code string) ([]byte, error) {
	resp, err := cli.Clone().
		// NoCache().
		SetURLByStr(`https://apimg.alipay.com/combo.png?d=cashier`).
		SetQuery("t", code).
		Get("")
	if err != nil {
		return nil, err
	}
	i, _, err := image.Decode(resp.RawBody())
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(nil)
	err = jpeg.Encode(buf, i, nil)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
