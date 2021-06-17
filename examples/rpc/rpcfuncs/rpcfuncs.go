package rpcfuncs

import (
	"fmt"
	"github.com/shaofeinus/gomq"
)

// RegisterRPC Register RPC functions. Should be called during initialization for both sender and receiver
func RegisterRPC() {
	gomq.RegisterRPCFunc("FuncA", "a", FuncA)
	gomq.RegisterRPCFunc("FuncB", "b", FuncB)
}

func FuncA(args map[string]interface{}) {
	fmt.Printf("Func A: %v\n", args)
}

func FuncB(args map[string]interface{}) {
	fmt.Printf("Func B: %v\n", args)
}
