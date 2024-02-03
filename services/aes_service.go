package services

import (
	"encoding/hex"
	"encoding/json"
	"github.com/Dubbril/go-edge-utility/models"
	"github.com/Dubbril/go-edge-utility/utils"
)

type AesService struct {
}

func NewAesService() *AesService {
	return &AesService{}
}

func (a *AesService) EncryptData(aesRequest *models.AesRequest) (string, error) {

	err := aesRequest.Validate()
	if err != nil {
		return "", err
	}

	var encrypted string
	if isHexString(aesRequest.Key) {
		encrypted, err = utils.EncryptWithHexKey(aesRequest.Data, aesRequest.Key)
		if err != nil {
			return "", err
		}
	} else {
		encrypted, err = utils.EncryptWithBase64Key(aesRequest.Data, aesRequest.Key)
		if err != nil {
			return "", err
		}
	}

	return encrypted, nil
}

func (a *AesService) DecryptData(aesRequest *models.AesRequest) (string, error) {
	_, err := json.Marshal(aesRequest.Data)
	if err != nil {
		return "", err
	}

	var encrypted string
	if isHexString(aesRequest.Key) {
		encrypted, err = utils.DecryptWithHexKey(aesRequest.Data, aesRequest.Key)
		if err != nil {
			return "", err
		}
	} else {
		encrypted, err = utils.DecryptWithBase64Key(aesRequest.Data, aesRequest.Key)
		if err != nil {
			return "", err
		}
	}

	return encrypted, nil
}

func isHexString(s string) bool {
	_, err := hex.DecodeString(s)
	return err == nil
}
