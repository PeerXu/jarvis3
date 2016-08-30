package http

import (
	httptransport "github.com/go-kit/kit/transport/http"
)

type EncodeRequestFuncMiddleware func(httptransport.EncodeRequestFunc) httptransport.EncodeRequestFunc

func EncodeRequestFuncChain(outer EncodeRequestFuncMiddleware, others ...EncodeRequestFuncMiddleware) EncodeRequestFuncMiddleware {
	return func(next httptransport.EncodeRequestFunc) httptransport.EncodeRequestFunc {
		for i := len(others) - 1; i >= 0; i-- {
			next = others[i](next)
		}
		return outer(next)
	}
}

type DecodeResponseFuncMiddleware func(httptransport.DecodeResponseFunc) httptransport.DecodeResponseFunc

func DecodeResponseFuncChain(outer DecodeResponseFuncMiddleware, others ...DecodeResponseFuncMiddleware) DecodeResponseFuncMiddleware {
	return func(next httptransport.DecodeResponseFunc) httptransport.DecodeResponseFunc {
		for i := len(others) - 1; i >= 0; i-- {
			next = others[i](next)
		}
		return outer(next)
	}
}
