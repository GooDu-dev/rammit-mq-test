package types

import (
	"encoding/json"
	"log"

	"github.com/GooDu-Dev/function-parser-go/pkg/inventory"
)

type Process struct {
	fname       string                 `json:"fname"`
	params      map[string]interface{} `json:"params"`
	returnValue interface{}            `json:"return_value"`
}

func NewProcess(fname string, params map[string]interface{}) inventory.Process {
	return &Process{
		fname:  fname,
		params: params,
	}
}

func (p *Process) Function() string {
	return p.fname
}

func (p *Process) Params() map[string]interface{} {
	return p.params
}

func (p *Process) Return() interface{} {
	return p.returnValue
}

func (p *Process) JSON() []byte {
	jsonBytes, err := json.Marshal(map[string]interface{}{
		"function_name": p.Function(),
		"params":        p.Params(),
	})
	if err != nil {
		log.Fatal("errror mashalling process to json :", err.Error())
		return nil
	}
	log.Printf("Marshal JSON successfully : %v\n", string(jsonBytes))
	return jsonBytes
}
