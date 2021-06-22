package middleware

import "time"

type Message struct {
	Protocol string
	Name     string
	Args     map[string]interface{}
	Ret      interface{}
	Error    error
	Duration time.Duration
}

type Processor struct {
	message        Message
	currMiddleWare []Middleware
}

type Middleware func(message Message)

var SENDER_MIDDLEWARES []Middleware
var RECEIVER_MIDDLEWARES []Middleware

func AddSenderMiddleware(middleware Middleware) {
	SENDER_MIDDLEWARES = append(SENDER_MIDDLEWARES, middleware)
}

func AddReceiverMiddleware(middleware Middleware) {
	RECEIVER_MIDDLEWARES = append(RECEIVER_MIDDLEWARES, middleware)
}

func ApplySenderMiddlewares(message Message) {
	applyMiddlewares(message, SENDER_MIDDLEWARES)
}

func ApplyReceiverMiddlewares(message Message) {
	applyMiddlewares(message, RECEIVER_MIDDLEWARES)
}

func applyMiddlewares(message Message, middlewares []Middleware) {
	for _, middleware := range middlewares {
		middleware(message)
	}
}
