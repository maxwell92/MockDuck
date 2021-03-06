package contract

import (
	"encoding/json"
	"io/ioutil"
)

type Response struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
	Name    string `json:"name,omitempty"`
	Comment string `json:"comment,omitempty"`
}

func (r Response) Encode() []byte {
	data, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}

	return data
}

func (r Response) String() string {
	data := r.Encode()
	return string(data)
}

/*
type Contract struct {
	URL  string   `json:"url"`
	Get  Response `json:"get"`
	Post Response `json:"post"`
}
*/

type Contract struct {
	URL  string   `json:"url"`
	Get  Response `json:"get"`
	Post string   `json:"post"`
}

type ContractList struct {
	Contracts []Contract `json:"contracts"`
}

func NewContract(file string) *ContractList {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	c := new(ContractList)
	c.Contracts = make([]Contract, 0)
	err = json.Unmarshal([]byte(data), c)
	if err != nil {
		panic(err)
	}

	return c
}
