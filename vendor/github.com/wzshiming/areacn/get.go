package areacn

import (
	"encoding/json"
	"sort"
	"strings"
)

type Area struct {
	AreaID string `json:"area_id"`
	Name   string `json:"name"`
	Level  int    `json:"level"`
}

type AreaLink struct {
	Area
	Children []*AreaLink `json:"children"`
}

var Areas []*AreaLink

func search(areas []*AreaLink, areaID string) int {
	return sort.Search(len(areas), func(i int) bool {
		return areaID <= areas[i].AreaID
	})
}

func init() {
	data := MustAsset("pcctv.json")
	json.Unmarshal(data, &Areas)
	setLevel(Areas, 1)
}

func setLevel(areas []*AreaLink, level int) {
	for _, area := range areas {
		area.Level = level
		setLevel(area.Children, level+1)
	}
}

func lookup(areas []*AreaLink, areaID string) []*AreaLink {
	switch len(areaID) {
	default:
		return nil
	case 0:
	case 2, 4, 6, 9:
		off := []int{2, 4, 6, 9}
		l := len(areaID) / 2
		for i := 0; i != l; i++ {
			index := search(areas, areaID[:off[i]])
			if index == -1 || index == len(areas) {
				return nil
			}
			if !strings.HasPrefix(areaID, areas[index].AreaID) {
				return nil
			}
			areas = areas[index].Children
		}
	}
	return areas
}

// Get 根据 areaID 获取地址信息
// 传空字符串的获取第一级省
func Get(areaID string) []*Area {
	raw := lookup(Areas, areaID)
	if raw == nil {
		return nil
	}

	areas := make([]*Area, 0, len(raw))
	for _, area := range raw {
		areas = append(areas, &area.Area)
	}
	return areas
}
