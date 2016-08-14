package ec2

import (
	"encoding/hex"
	"crypto/rand"
	"strconv"
)

func clientToken() (string, error) {
	// Maximum EC2 client token size is 64 bytes.
	buf := make([]byte, 32)
	_, err := rand.Read(buf)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(buf), nil
}



func addParamsList(params map[string]string, label string, ids []string) {
	for i, id := range ids {
		params[label+"."+strconv.Itoa(i+1)] = id
	}
}


func makeParams(action string) map[string]string {
	params := make(map[string]string)
	params["Action"] = action
	return params
}
