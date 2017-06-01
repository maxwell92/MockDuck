package contract

import (
	"encoding/json"
	"io/ioutil"
)

type Response struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func (r Response) String() string {
	data, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}

	return string(data)
}

type Contract struct {
	URL  string   `json:"url"`
	Get  Response `json:"get"`
	Post Response `json:"post"`
}

func NewContract(file string) *Contract {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	c := new(Contract)
	err = json.Unmarshal([]byte(data), c)
	if err != nil {
		panic(err)
	}

	return c
}
