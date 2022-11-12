package neo4jstore

import (
	"fmt"
	"github.com/spf13/cast"
	"gitlab.com/rubin-dev/api/pkg/models"
	"math"
	"strings"
)

const unmarked = 16
const totalKey = "totalin"
const categoryPrefix = "cat"

type DirectoryI18n struct {
	Ru string `json:"ru"`
	En string `json:"en"`
}

type CalculateItem struct {
	ID         int           `json:"id"`
	Total      float64       `json:"total"`
	Percent    float64       `json:"percent"`
	PercentRaw float64       `json:"-"`
	Risk       float64       `json:"risk"`
	RiskRaw    float64       `json:"-"`
	I18n       DirectoryI18n `json:"i18n"`
}

type CalculatedRisk struct {
	Total   float64                `json:"total"`
	Risk    float64                `json:"risk"`
	RiskRaw float64                `json:"-"`
	Items   map[int]*CalculateItem `json:"items"`
}

func round(v float64) float64 {
	return math.Round(v*100) / 100
}

func Calculate(categories map[int]float64, total float64) (*CalculatedRisk, error) {
	var totalRisk float64

	items := map[int]*CalculateItem{}
	for id, value := range categories {
		dir := models.FindDirectory(id)
		if dir == nil {
			return nil, fmt.Errorf("directory %d not found", id)
		}

		percent := value / total
		risk := (percent * (dir.Risk / 100)) * 100
		totalRisk += risk

		items[id] = &CalculateItem{
			ID:         id,
			Percent:    round(percent * 100),
			PercentRaw: percent * 100,
			Total:      value,
			Risk:       round(risk),
			RiskRaw:    risk,
			I18n: DirectoryI18n{
				Ru: dir.NameRu,
				En: dir.NameEn,
			},
		}
	}

	return &CalculatedRisk{
		Total:   total,
		RiskRaw: totalRisk,
		Risk:    round(totalRisk),
		Items:   items,
	}, nil
}

func CalculateRiskFromNode(node *Node) (*CalculatedRisk, error) {
	var total float64
	var all float64

	categories := map[int]float64{}
	for k, v := range node.Props {
		if k == totalKey {
			total = cast.ToFloat64(v)
			continue
		}

		if !strings.HasPrefix(k, categoryPrefix) {
			continue
		}

		if len(k) != 5 {
			continue
		}

		id := strings.TrimLeft(strings.TrimLeft(k, "cat"), "0")
		vv := cast.ToFloat64(v)
		categories[cast.ToInt(id)] = vv
		all += vv
	}

	if total == 0 {
		return nil, nil
	}

	if _, ok := categories[unmarked]; !ok {
		categories[unmarked] = total - all
	}

	return Calculate(categories, total)
}
