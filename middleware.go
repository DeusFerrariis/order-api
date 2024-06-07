package main

import "net/http"

type MiddlewareChain struct {
	middlewares []func(http.HandlerFunc) http.HandlerFunc
}

func NewMiddlewareChain() MiddlewareChain {
	return MiddlewareChain{
		middlewares: make([]func(http.HandlerFunc) http.HandlerFunc, 0),
	}
}

func (mc *MiddlewareChain) Append(middleware func(http.HandlerFunc) http.HandlerFunc) {
	mc.middlewares = append(mc.middlewares, middleware)
}

func (mc *MiddlewareChain) WrapHandler(next http.HandlerFunc) http.HandlerFunc {
	chainedHandler := next
	for _, middleware := range mc.middlewares {
		chainedHandler = middleware(chainedHandler)
	}
	return chainedHandler
}
