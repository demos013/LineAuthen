package encoding

import (
	"encoding/base64"
	"encoding/json"
)

func Encode(input interface{}) (string, error) {
	encode, err := json.Marshal(input)
	if err != nil {
		return "", err
	}
	sEnc := base64.StdEncoding.EncodeToString([]byte(encode))
	return sEnc, err
}

func Decode(input string) (interface{}, error) {
	sDec, _ := base64.StdEncoding.DecodeString(input)

	var dump interface{}
	err := json.Unmarshal([]byte(sDec), &dump)
	if err != nil {
		return nil, err
	}
	return dump, nil
}
