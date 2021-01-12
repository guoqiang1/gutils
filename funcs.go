package gutils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"hash/crc32"
	"math/big"
	"strings"
	"time"
)

// Md5 md5
func Md5(data string) string {
	m := md5.New()
	m.Write([]byte(data))
	return hex.EncodeToString(m.Sum([]byte(nil)))
}

// Crc32 crc32
func Crc32(data string) uint32 {
	return crc32.ChecksumIEEE([]byte(data))
}

// RandString10 返回数字组成的随机字符串
func RandString10(l int) (string, error) {
	result := make([]byte, l)

	str := "01234567890"
	byteStr := []byte(str)

	for i := 0; i < l; i++ {
		temp, err := rand.Int(rand.Reader, big.NewInt(10))

		if err != nil {
			return "", err
		}

		result[i] = byteStr[temp.Int64()]

	}

	return string(result), nil
}

// RandString62 生成由数字、小写字母、大写字母组成的随机数
func RandString62(l int) (string, error) {
	result := make([]byte, l)

	str := "01234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	byteStr := []byte(str)

	for i := 0; i < l; i++ {
		temp, err := rand.Int(rand.Reader, big.NewInt(62))

		if err != nil {
			return "", err
		}

		result[i] = byteStr[temp.Int64()]

	}

	return string(result), nil
}

// 获取当月第一天凌晨的时间
func FirstDayOfNowMonth() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
}

// DateTimePower 生成格式化的时间字符串
func DateTimePower(fm string, timestamp int64) string {

	if timestamp == 0 {
		timestamp = time.Now().Unix()
	}

	var lctime = time.Unix(timestamp, 0)

	fm = strings.Replace(fm, "Y", "2006", 1)
	fm = strings.Replace(fm, "m", "01", 1)
	fm = strings.Replace(fm, "d", "02", 1)
	fm = strings.Replace(fm, "H", "15", 1)
	fm = strings.Replace(fm, "i", "04", 1)
	fm = strings.Replace(fm, "s", "05", 1)

	return lctime.Format(fm)
}

// 获取本周一凌晨的时间
func FirstDayOfNowWeek() time.Time {
	now := time.Now()

	offset := int(time.Monday - now.Weekday())

	// 如果>0则说明是周天
	if offset > 0 {
		offset = -6
	}

	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).AddDate(0, 0, offset)
}

// ArrayUniqueString 切片去重
func ArrayUniqueString(data []string) []string {
	if len(data) == 0 {
		return data
	}

	newdata := make([]string, 0, 0)

	m1 := make(map[string]bool)

	for _, item := range data {
		if _, ok := m1[item]; ok {
			continue
		}

		m1[item] = true
		newdata = append(newdata, item)
	}

	return newdata
}

// ArrayUniqueInt64  切片去重
func ArrayUniqueInt64(data []int64) []int64 {
	if len(data) == 0 {
		return data
	}

	newdata := make([]int64, 0, 0)

	m1 := make(map[int64]bool)

	for _, item := range data {
		if _, ok := m1[item]; ok {
			continue
		}

		m1[item] = true
		newdata = append(newdata, item)
	}

	return newdata
}
