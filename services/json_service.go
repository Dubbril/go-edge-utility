package services

import (
	"encoding/base64"
	"encoding/json"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strconv"
	"strings"
	"unicode"
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

func (s JsonService) Base64ToImage(base64String string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return nil, err
	}

	return decoded, nil
}

func (s JsonService) ImageToBase64(imageByteArray []byte) string {
	return base64.StdEncoding.EncodeToString(imageByteArray)
}

func (s JsonService) JSONSnakeToCamel(input interface{}) interface{} {
	switch v := input.(type) {
	case map[string]interface{}:
		result := make(map[string]interface{})
		for key, value := range v {
			result[snakeToCamel(key)] = s.JSONSnakeToCamel(value)
		}
		return result

	case []interface{}:
		result := make([]interface{}, len(v))
		for i, element := range v {
			result[i] = s.JSONSnakeToCamel(element)
		}
		return result

	default:
		return input
	}
}

func snakeToCamel(s string) string {
	caser := cases.Title(language.English)
	parts := strings.Split(s, "_")
	for i := range parts {
		if i > 0 {
			parts[i] = caser.String(parts[i])
		}
	}
	return strings.Join(parts, "")
}

func (s JsonService) JSONCamelToSnake(input interface{}) interface{} {
	switch v := input.(type) {
	case map[string]interface{}:
		result := make(map[string]interface{})
		for key, value := range v {
			result[camelToSnake(key)] = s.JSONCamelToSnake(value)
		}
		return result

	case []interface{}:
		result := make([]interface{}, len(v))
		for i, element := range v {
			result[i] = s.JSONCamelToSnake(element)
		}
		return result

	default:
		return input
	}
}

func camelToSnake(s string) string {
	var result []string
	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) {
			result = append(result, "_")
		}
		result = append(result, strings.ToLower(string(r)))
	}
	return strings.Join(result, "")
}
