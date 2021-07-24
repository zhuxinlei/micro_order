package common

import (
	"encoding/json"
	"strconv"
)

func ToInt(input interface{}) int {
	jsonByte, _ := json.Marshal(input)
	userId, _ := strconv.Atoi(string(jsonByte))
	return userId
}
