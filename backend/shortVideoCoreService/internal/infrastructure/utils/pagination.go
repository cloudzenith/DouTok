package utils

func GetLimitOffset(page, size int) (limit int, offset int) {
	if page < 1 {
		page = 1
	}

	if size < 1 {
		size = 10
	}

	return size, (page - 1) * size
}
