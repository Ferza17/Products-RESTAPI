package generate

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func GetRandomString(length int, str string) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" + str +
		"0123456789")
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}

// order number generated with format PO-123/IX/2020
//(IX is current month)(2020 is current year),
//(123 reset per month).
func GetOrderNumber() string {
	sb := strings.Builder{}
	tm := time.Now()

	sb.WriteString("PO-")
	// 123 reset per month
	// 1611759012
	code := fmt.Sprint(tm.Year(), int(tm.Month()))
	code = strings.ReplaceAll(code, "20", "")
	code = strings.ReplaceAll(code, " ", "")

	sb.WriteString(fmt.Sprint(code))
	sb.WriteString("/")
	month := tm.Month()
	sb.WriteString(intToRoman(int(month)))
	sb.WriteString("/")
	sb.WriteString(fmt.Sprint(tm.Year()))

	return sb.String()
}

func intToRoman(num int) string {
	values := []int{
		1000, 900, 500, 400,
		100, 90, 50, 40,
		10, 9, 5, 4, 1,
	}

	symbols := []string{
		"M", "CM", "D", "CD",
		"C", "XC", "L", "XL",
		"X", "IX", "V", "IV",
		"I"}
	roman := ""
	i := 0

	for num > 0 {
		k := num / values[i]
		for j := 0; j < k; j++ {
			roman += symbols[i]
			num -= values[i]
		}
		i++
	}
	return roman
}
