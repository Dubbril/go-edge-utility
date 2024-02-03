package models

import "errors"

type AesRequest struct {
	Key  string `json:"key"`
	Data string `json:"data"`
}

func (a *AesRequest) Validate() error {
	if a == nil {
		return errors.New("AesRequest is nil")
	}

	if a.Key == "" {
		return errors.New("key cannot be empty")
	}

	if a.Data == "" {
		return errors.New("data cannot be empty")
	}

	return nil
}
