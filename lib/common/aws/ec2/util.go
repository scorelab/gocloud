package ec2


import (
	"net/url"
	"encoding/hex"
	"crypto/rand"
)

func multimap(p map[string]string) url.Values {
	q := make(url.Values, len(p))
	for k, v := range p {
		q[k] = []string{v}
	}
	return q
}


func makeParams(action string) map[string]string {
	params := make(map[string]string)
	params["Action"] = action
	return params
}

func clientToken() (string, error) {
	// Maximum EC2 client token size is 64 bytes.
	buf := make([]byte, 32)
	_, err := rand.Read(buf)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(buf), nil
}

// Create a base set of params for an action
func MakeParams(action string) map[string]string {
	params := make(map[string]string)
	params["Action"] = action
	return params
}
