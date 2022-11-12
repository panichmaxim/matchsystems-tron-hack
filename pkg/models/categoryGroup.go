package models

type CategoryGroup struct {
	ID     int     `json:"id"`
	NameRu string  `json:"name_ru"`
	NameEn string  `json:"name_en"`
	Risk   float64 `json:"risk"`
}

var categories = []*CategoryGroup{
	// 1: 'Хищения и компьютерные преступления (ущерб 100%)'
	{
		ID:     1,
		NameEn: "Stolen funds and Computer crimes",
		NameRu: "Хищения и компьютерные преступления",
		Risk:   100,
	},
	// 2: 'Наркотики (ущерб 100%)
	{
		ID:     2,
		NameEn: "Drugs",
		NameRu: "Наркотики",
		Risk:   100,
	},
	// 3: 'Персональные данные и документы (ущерб 60%)'
	{
		ID:     3,
		NameEn: "Personal data and documents",
		NameRu: "Персональные данные и документы",
		Risk:   60,
	},
	// Отмыв (ущерб 90%)
	{
		ID:     4,
		NameEn: "Money laundering",
		NameRu: "Легализация (отмывание) средств, полученных преступным путём",
		Risk:   90,
	},
	// Высокорискованный обмен (ущерб 75%)
	{
		ID:     5,
		NameEn: "Exchange with high risk",
		NameRu: "Высокорискованный обмен",
		Risk:   75,
	},
	// Обмен и сделки умеренного риска (ущерб 50%)
	{
		ID:     6,
		NameEn: "Exchange and transactions with medium risk",
		NameRu: "Обмен и сделки умеренного риска",
		Risk:   50,
	},
	// Низкорисокванный обмен (ущерб 25%)
	{
		ID:     7,
		NameEn: "Exchange with low risk",
		NameRu: "Низкорискованный обмен",
		Risk:   25,
	},
	// Child Abuse (ущерб 100%)
	{
		ID:     8,
		NameEn: "Child Abuse",
		NameRu: "Детская порнография",
		Risk:   100,
	},
	// Государства, чиновники (ущерб 75%)
	{
		ID:     9,
		NameEn: "States, officials",
		NameRu: "Государства, чиновники",
		Risk:   75,
	},
	// Злоупотребление властью, санкции, военизированные организации и терроризм (ущерб 100%)
	{
		ID:     10,
		NameEn: "Abuse of authority, sanctions, paramilitary organisations and terrorism",
		NameRu: "Злоупотребление властью, санкции, военизированные организации и терроризм",
		Risk:   100,
	},
	// Медиа (ущерб 25%)
	{
		ID:     11,
		NameEn: "Media",
		NameRu: "Медиа",
		Risk:   25,
	},
	// Иная незаконная деятельность (ущерб 100%)
	{
		ID:     12,
		NameEn: "Other illegal activities",
		NameRu: "Иная незаконная деятельность",
		Risk:   100,
	},
	// DEX (ущерб 60%)
	{
		ID:     13,
		NameEn: "DEX",
		NameRu: "DEX",
		Risk:   60,
	},
	// Кристальная чистота (ущерб 0%)
	{
		ID:     14,
		NameEn: "Clean funds",
		NameRu: "Чистые средства",
		Risk:   0,
	},
	// Ransom (ущерб 100%)
	{
		ID:     15,
		NameEn: "Ransom",
		NameRu: "Вымогательство",
		Risk:   100,
	},
	// Неразмеченные данные (Unmarked 50%)
	{
		ID:     16,
		NameEn: "Unmarked",
		NameRu: "Неразмеченные",
		Risk:   50,
	},
}

func GetCategoryGroups() []*CategoryGroup {
	return categories
}

func FindDirectory(id int) *CategoryGroup {
	for _, c := range GetCategoryGroups() {
		if c.ID == id {
			return c
		}
	}

	return nil
}
