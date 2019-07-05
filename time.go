package tools

import (
	"time"
)

// WeekCount 从0时间戳开始计算，取得当前是第几周（以周一开始算开始，策划要求）
func WeekCount(t time.Time, timeZone int64) int {
	timeT := t.Unix() + timeZone

	// 时间戳的开始是1970年1月1日0时区0点
	// 0时间戳是周四，所以要加回3天
	timeT += 259200

	week := timeT / 86400 / 7
	return int(week)
}
