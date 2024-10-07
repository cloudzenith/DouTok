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

func GetPageInfo(total int64, page, size int32) int32 {
	if total < 0 {
		total = 0
	}

	totalPage := total / int64(size)
	if total%int64(size) != 0 {
		totalPage++
	}

	return int32(totalPage)
}
