package service

func buildLimitOffset(page, limit int) (int, int) {
	if page <= 0 {
		page = 1
	}

	if limit < 1 || limit > 100 {
		limit = 100
	}

	var offset int
	if page > 1 {
		offset = (page - 1) * limit
	}

	return limit, offset
}
