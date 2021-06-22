package rpc

import (
	"errors"
	"fmt"
	"github.com/shaofeinus/gomq"
	"github.com/shaofeinus/gomq/middleware"
	"log"
	"time"
)

var RPCFUNCS = make(map[string]RPCFunc)

type Fn func(args map[string]interface{}) (interface{}, error)

type RPCFunc struct {
	Name  string
	Queue string
	Fn    Fn
}

func RegisterRPCFunc(name string, queue string, fn Fn) {
	RPCFUNCS[name] = RPCFunc{
		Name:  name,
		Queue: getQueue(queue),
		Fn:    fn,
	}
}

func InvokeRPC(name string, args map[string]interface{}) error {
	rpcFunc, err := findRPCFunc(name)
	if err != nil {
		return err
	}
	start := time.Now()
	err = gomq.SendJSONToQueue(rpcFunc.Queue, map[string]interface{}{
		"name": rpcFunc.Name,
		"args": args,
	})
	middleware.ApplySenderMiddlewares(middleware.Message{
		Protocol: "rpc",
		Name:     name,
		Args:     args,
		Ret:      nil,
		Error:    err,
		Duration: time.Now().Sub(start),
	})
	if err != nil {
		return err
	}
	return nil
}

func WorkOnRPC(queues []string) {
	for _, queue := range queues {
		gomq.ConsumeJSON(getQueue(queue), handleRPCFuncJson)
	}
}

func getQueue(queue string) string {
	return fmt.Sprintf("rpc-%s", queue)
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
	if err != nil {
		log.Println(err.Error())
		return
	}
	start := time.Now()
	args := rpcJson["args"].(map[string]interface{})
	ret, err := rpcFunc.Fn(args)
	middleware.ApplyReceiverMiddlewares(middleware.Message{
		Protocol: "rpc",
		Name:     rpcFunc.Name,
		Args:     args,
		Ret:      ret,
		Error:    err,
		Duration: time.Now().Sub(start),
	})
}
