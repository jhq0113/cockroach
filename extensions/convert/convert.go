package convert

import "strconv"

func Int64ToString(val int64) string {
	return strconv.FormatInt(val, 10)
}

func Int2String(val int) string {
	return strconv.Itoa(val)
}

func String2Int(str string) int {
	if val, err := strconv.Atoi(str); err == nil {
		return val
	}
	return 0
}

func String2Int64(str string) int64 {
	if val, err := strconv.ParseInt(str, 10, 64); err == nil {
		return val
	}
	return 0
}
