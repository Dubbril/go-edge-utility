package services

import (
	"encoding/json"
	"strconv"
)

type JsonService struct {
}

func NewJsonService() *JsonService {
	return &JsonService{}
}

func (s JsonService) EscapedJSON(strJson string) (string, error) {
	var jsonObject interface{}
	err := json.Unmarshal([]byte(strJson), &jsonObject)
	if err != nil {
		return "", err
	}

	escapedJSON, err := json.Marshal(jsonObject)
	if err != nil {
		return "", err
	}

	return strconv.Quote(string(escapedJSON)), nil
}
