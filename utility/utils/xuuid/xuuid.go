package xuuid

import (
	"github.com/bwmarrin/snowflake"
)

const base34Digits = "123456789qwertyuipasdfghjklzxcvbnm"

func GetsnowflakeUUID() snowflake.ID {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return 0
	}
	return node.Generate()
}

func NumberToBase34(num int64) string {
	if num == 0 {
		return string(base34Digits[0])
	}
	var result string
	for num > 0 {
		remainder := num % 34
		result = string(base34Digits[remainder]) + result
		num /= 34
	}
	return result

}
