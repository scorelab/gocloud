package ec2

import (
	"encoding/hex"
	"crypto/rand"
	"strconv"
	"sort"
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


type Filter struct {
	m map[string][]string
}

func NewFilter() *Filter {
	return &Filter{make(map[string][]string)}
}


func (f *Filter) Add(name string, value ...string) {
	f.m[name] = append(f.m[name], value...)
}

func (f *Filter) addParams(params map[string]string) {
	if f != nil {
		a := make([]string, len(f.m))
		i := 0
		for k := range f.m {
			a[i] = k
			i++
		}
		sort.StringSlice(a).Sort()
		for i, k := range a {
			prefix := "Filter." + strconv.Itoa(i+1)
			params[prefix+".Name"] = k
			for j, v := range f.m[k] {
				params[prefix+".Value."+strconv.Itoa(j+1)] = v
			}
		}
	}
}

