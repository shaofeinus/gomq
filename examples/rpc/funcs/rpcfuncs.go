package funcs

import (
	"fmt"
	"github.com/shaofeinus/gomq/rpc"
)

// Setup Register RPC functions. Should be called during initialization for both sender and receiver
func Setup() {
	rpc.RegisterRPCFunc("FuncA", "a", FuncA)
	rpc.RegisterRPCFunc("FuncB", "b", FuncB)
}

func FuncA(args map[string]interface{}) {
	fmt.Printf("Func A: %v\n", args)
}

func FuncB(args map[string]interface{}) {
	fmt.Printf("Func B: %v\n", args)
}
