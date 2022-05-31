package utils

import (
	"strconv"
)

//string转int
func StrToInt(str string) int {
	i, e := strconv.Atoi(str)
	if e != nil {
		return 0
	}
	return i
}

//string转uint
func StrToUInt(str string) uint {
	i, e := strconv.Atoi(str)
	if e != nil {
		return 0
	}
	return uint(i)
}

//string转int64
func StringToint64(a string) (int64, error) {
	return strconv.ParseInt(a, 10, 64)
}

// string转float64
func StringToFloat64(a string) (float64, error) {
	return strconv.ParseFloat(a, 64)
}

//int转string
func IntToString(a int) string {
	return strconv.Itoa(a)
}

//int64转string
func Int64ToString(a int64) string {
	return strconv.FormatInt(a, 10)
}

//float32转string
func Float32ToString(a float32) string {
	return strconv.FormatFloat(3.1415, 'E', -1, 32)
}

//float64转strinng
func Float64ToString(a float64) string {
	return strconv.FormatFloat(3.1415, 'E', -1, 64)
}

//int转int64
func IntToInt64(a int) int64 {
	return int64(a)
}

//int64转int
func Int64ToInt(a int64) int {
	return int(a)
}
