package aws

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/url"
	"sort"
	"strings"
	"time"
	"errors"
	"os"
)



type Auth struct {
	AccessKey, SecretKey string
	token                string
	expiration           time.Time
}

var unreserved = make([]bool, 128)
var hex = "0123456789ABCDEF"
var b64 = base64.StdEncoding
func init() {
	// RFC3986
	u := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz01234567890-_.~"
	for _, c := range u {
		unreserved[c] = true
	}
}


type credentials struct {
	Code            string
	LastUpdated     string
	Type            string
	AccessKeyId     string
	SecretAccessKey string
	Token           string
	Expiration      string
}

func (a *Auth) Token() string {
	if a.token == "" {
		return ""
	}
	if time.Since(a.expiration) >= -30 * time.Second {
		*a, _ = GetAuth("", "")
	}
	return a.token
}

func (a *Auth) Expiration() time.Time {
	return a.expiration
}

// To be used with other APIs that return auth credentials such as STS
func NewAuth(accessKey, secretKey, token string, expiration time.Time) *Auth {
	return &Auth{
		AccessKey:  accessKey,
		SecretKey:  secretKey,
		token:      token,
		expiration: expiration,
	}
}

// GetAuth creates an Auth based on either passed in credentials, or environment information
func GetAuth(accessKey string, secretKey string) (auth Auth, err error ) {
	if( accessKey!= "" &&  secretKey !="") {
		auth.AccessKey = accessKey
		auth.SecretKey = secretKey

	}
	auth, err = EnvAuth()
	if err == nil {
		return
	}
	err = errors.New("No valid AWS authentication found")
	return auth, err
}

// The AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY environment variables are used.
func EnvAuth() (auth Auth, err error) {
	auth.AccessKey = os.Getenv("AWS_ACCESS_KEY_ID")
	if auth.AccessKey == "" {
		auth.AccessKey = os.Getenv("AWS_ACCESS_KEY")
	}

	auth.SecretKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	if auth.SecretKey == "" {
		auth.SecretKey = os.Getenv("AWS_SECRET_KEY")
	}

	if auth.AccessKey == "" {
		err = errors.New("AWS_ACCESS_KEY_ID or AWS_ACCESS_KEY not found in environment")
	}
	if auth.SecretKey == "" {
		err = errors.New("AWS_SECRET_ACCESS_KEY or AWS_SECRET_KEY not found in environment")
	}
	auth.token = os.Getenv("AWS_SESSION_TOKEN")
	return
}





// Designates a signer interface suitable for signing AWS requests, params
// should be appropriately encoded for the request before signing.
type Signer interface {
	Sign(method, path string, params map[string]string)
}

type V2Signer struct {
	auth    Auth
	service ServiceInfo
	host    string
}

func NewV2Signer(auth Auth, service ServiceInfo) (*V2Signer, error) {
	u, err := url.Parse(service.Endpoint)
	if err != nil {
		return nil, err
	}
	return &V2Signer{auth: auth, service: service, host: u.Host}, nil
}

func (s *V2Signer) Sign(method, path string, params map[string]string) {
	params["AWSAccessKeyId"] = s.auth.AccessKey
	params["SignatureVersion"] = "2"
	params["SignatureMethod"] = "HmacSHA256"
	if s.auth.Token() != "" {
		params["SecurityToken"] = s.auth.Token()
	}
	var keys, sarray []string
	for k, _ := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		sarray = append(sarray, Encode(k) + "=" + Encode(params[k]))
	}
	joined := strings.Join(sarray, "&")
	payload := method + "\n" + s.host + "\n" + path + "\n" + joined
	hash := hmac.New(sha256.New, []byte(s.auth.SecretKey))
	hash.Write([]byte(payload))
	signature := make([]byte, b64.EncodedLen(hash.Size()))
	b64.Encode(signature, hash.Sum(nil))
	params["Signature"] = string(signature)
}




// Encode takes a string and URI-encodes it in a way suitable
func Encode(s string) string {
	encode := false
	for i := 0; i != len(s); i++ {
		c := s[i]
		if c > 127 || !unreserved[c] {
			encode = true
			break
		}
	}
	if !encode {
		return s
	}
	e := make([]byte, len(s) * 3)
	ei := 0
	for i := 0; i != len(s); i++ {
		c := s[i]
		if c > 127 || !unreserved[c] {
			e[ei] = '%'
			e[ei + 1] = hex[c >> 4]
			e[ei + 2] = hex[c & 0xF]
			ei += 3
		} else {
			e[ei] = c
			ei += 1
		}
	}
	return string(e[:ei])
}

