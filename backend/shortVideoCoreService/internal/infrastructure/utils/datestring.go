package utils

import (
	"fmt"
	"time"
)

const (
	JustNow    = "刚刚"
	MinutesAgo = "%d分钟前"
	HoursAgo   = "%d小时前"
	DaysAgo    = "%d天前"
)

// ParseToDateString 将time.Time转换为日期字符串
func ParseToDateString(t time.Time) string {
	d := time.Since(t)

	// 一分钟内 -> 刚刚
	if d.Minutes() < 1 {
		return JustNow
	}

	// 一小时内 -> xx分钟前
	if d.Hours() < 1 {
		return fmt.Sprintf(MinutesAgo, int(d.Minutes()))
	}

	// 一天内 -> xx小时前
	if d.Hours() < 24 {
		return fmt.Sprintf(HoursAgo, int(d.Hours()))
	}

	// 一个月内 -> xx天前
	if d.Hours() < 24*30 {
		return fmt.Sprintf(DaysAgo, int(d.Hours()/24))
	}

	// 指定日期
	return t.Format("2006-01-02")
}
