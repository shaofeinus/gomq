package gomq

import (
	"errors"
	"fmt"
	"log"
)

var RPCFUNCS = make(map[string]RPCFunc)

type RPCFunc struct {
	Name  string
	Queue string
	Fn    func(args map[string]interface{})
}

func RegisterRPCFunc(name string, queue string, fn func(args map[string]interface{})) {
	RPCFUNCS[name] = RPCFunc{
		Name:  name,
		Queue: queue,
		Fn:    fn,
	}
}

func InvokeRPC(name string, args map[string]interface{}) error {
	rpcFunc, err := findRPCFunc(name)
	if err != nil {
		return err
	}
	err = SendJSONToQueue(rpcFunc.Queue, map[string]interface{}{
		"name": rpcFunc.Name,
		"args": args,
	})
	if err != nil {
		return err
	}
	return nil
}

func WorkOnRPC(queues []string) {
	for _, queue := range queues {
		ConsumeJSON(queue, handleRPCFuncJson)
	}
}

func findRPCFunc(name string) (*RPCFunc, error) {
	rpcFunc, ok := RPCFUNCS[name]
	if !ok {
		return nil, errors.New(fmt.Sprintf("function not found \"%s\"", name))
	}
	return &rpcFunc, nil
}

func handleRPCFuncJson(rpcJson map[string]interface{}) {
	rpcFunc, err := findRPCFunc(rpcJson["name"].(string))
	args := rpcJson["args"].(map[string]interface{})
	if err != nil {
		log.Println(err.Error())
		return
	}
	rpcFunc.Fn(args)
}
