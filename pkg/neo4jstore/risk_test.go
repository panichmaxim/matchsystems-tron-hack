package neo4jstore

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalculateRiskFromNode(t *testing.T) {
	node := &Node{
		ID:     1,
		Labels: nil,
		Props: map[string]interface{}{
			"category": "foobar",
			totalKey:   10,
			"cat11":    1,
			"cat07":    1,
			"cat10":    1,
			"cat14":    1,
			"cat15":    1,
			"cat04":    1,
			"cat05":    1,
			"cat12":    1,
			"cat01":    1,
			"cat03":    1,
			"cat13":    1,
			"cat02":    1,
			"cat06":    1,
			"cat08":    1,
			"cat09":    1,
		},
	}
	r, err := CalculateRiskFromNode(node)
	require.NoError(t, err)
	require.Equal(t, float64(10), r.Total)
	require.Equal(t, float64(81), r.Risk)
}

// 0)Всего получено: 9565 (100.00%), ущерб по директории 0.00%
// 1)Хищения и компьютерные преступления (ущерб 100%): 1 (0.01%), ущерб по директории 0.01%
// 5)Высокорискованный обмен (ущерб 75%): 2129 (22.26%), ущерб по директории 16.69%
// 6)Обмен и сделки умеренного риска (ущерб 50%): 175 (1.83%), ущерб по директории 0.91%
// 11)Медиа (ущерб 25%): 297 (3.11%), ущерб по директории 0.78%
// 14)Кристальная чистота (ущерб 0%): 103 (1.08%), ущерб по директории 0.00%
// Риск-скор адреса: 18.40%
func TestCalculate_Example1(t *testing.T) {
	categories := map[int]float64{
		1:  1,
		5:  2129,
		6:  175,
		11: 297,
		14: 103,
	}
	var total float64 = 9565
	r, err := Calculate(categories, total)
	require.NoError(t, err)
	require.Equal(t, 18.4, r.Risk)
	require.Equal(t, total, r.Total)

	type test struct {
		id      int
		percent float64
		risk    float64
	}

	tests := []test{
		{1, 0.01, 0.01},
		{5, 22.26, 16.69},
		{6, 1.83, 0.91},
		{11, 3.11, 0.78},
		{14, 1.08, 0},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("percent_%d", i), func(t *testing.T) {
			cat, ok := r.Items[tc.id]
			require.True(t, ok)
			require.Equal(t, tc.percent, cat.Percent)
			require.Equal(t, tc.risk, cat.Risk)
		})
	}
}

// 0)Всего получено: 3337355877 (100.00%), ущерб по директории 0.00%
// 1)Хищения и компьютерные преступления (ущерб 100%): 634576922 (19.01%), ущерб по директории 19.01%
// 2)Наркотики (ущерб 100%): 83699 (0.00%), ущерб по директории 0.00%
// 4)Отмыв (ущерб 90%): 50 (0.00%), ущерб по директории 0.00%
// 5)Высокорискованный обмен (ущерб 75%): 1741484779 (52.18%), ущерб по директории 39.14%
// 6)Обмен и сделки умеренного риска (ущерб 50%): 3290806 (0.10%), ущерб по директории 0.05%
// 11)Медиа (ущерб 25%): 338105 (0.01%), ущерб по директории 0.00%
// 12)Иная незаконная деятельность (ущерб 100%): 75075 (0.00%), ущерб по директории 0.00%
// 14)Кристальная чистота (ущерб 0%): 20878716 (0.63%), ущерб по директории 0.00%
// Риск-скор адреса: 58.21%
func TestCalculate_Example2(t *testing.T) {
	categories := map[int]float64{
		1:  196166086,
		2:  1111,
		5:  48679998,
		6:  36946,
		11: 3357,
		12: 319,
		14: 185261,
	}
	var total float64 = 245073084
	r, err := Calculate(categories, total)
	require.NoError(t, err)
	require.Equal(t, 94.95, r.Risk)
	require.Equal(t, total, r.Total)

	type test struct {
		id      int
		percent float64
		risk    float64
	}

	tests := []test{
		{1, 80.04, 80.04},
		{2, 0, 0},
		{5, 19.86, 14.90},
		{6, 0.02, 0.01},
		{11, 0, 0},
		{12, 0, 0},
		{14, 0.08, 0},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("percent_%d", i), func(t *testing.T) {
			cat, ok := r.Items[tc.id]
			require.True(t, ok)
			require.Equal(t, tc.percent, cat.Percent)
			require.Equal(t, tc.risk, cat.Risk)
		})
	}
}

// 0)Всего получено: 3337355877 (100.00%), ущерб по директории 0.00%
// 1)Хищения и компьютерные преступления (ущерб 100%): 634576922 (19.01%), ущерб по директории 19.01%
// 2)Наркотики (ущерб 100%): 83699 (0.00%), ущерб по директории 0.00%
// 4)Отмыв (ущерб 90%): 50 (0.00%), ущерб по директории 0.00%
// 5)Высокорискованный обмен (ущерб 75%): 1741484779 (52.18%), ущерб по директории 39.14%
// 6)Обмен и сделки умеренного риска (ущерб 50%): 3290806 (0.10%), ущерб по директории 0.05%
// 11)Медиа (ущерб 25%): 338105 (0.01%), ущерб по директории 0.00%
// 12)Иная незаконная деятельность (ущерб 100%): 75075 (0.00%), ущерб по директории 0.00%
// 14)Кристальная чистота (ущерб 0%): 20878716 (0.63%), ущерб по директории 0.00%
// Риск-скор адреса: 58.21%
func TestCalculate_Example3(t *testing.T) {
	categories := map[int]float64{
		1:  634576922,
		2:  83699,
		4:  50,
		5:  1741484779,
		6:  3290806,
		11: 338105,
		12: 75075,
		14: 20878716,
	}
	var total float64 = 3337355877
	r, err := Calculate(categories, total)
	require.NoError(t, err)
	require.Equal(t, 58.21, r.Risk)
	require.Equal(t, total, r.Total)

	type test struct {
		id      int
		percent float64
		risk    float64
	}

	tests := []test{
		{1, 19.01, 19.01},
		{2, 0, 0},
		{4, 0, 0},
		{5, 52.18, 39.14},
		{6, 0.10, 0.05},
		{11, 0.01, 0},
		{12, 0, 0},
		{14, 0.63, 0},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("percent_%d", i), func(t *testing.T) {
			cat, ok := r.Items[tc.id]
			require.True(t, ok)
			require.Equal(t, tc.percent, cat.Percent)
			require.Equal(t, tc.risk, cat.Risk)
		})
	}
}
