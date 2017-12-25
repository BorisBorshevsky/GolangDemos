package cache

import "errors"

const keyPrefix = "$cache$"

type cacheMaker struct {
	Provider
	key string
}

var NotExist = errors.New("not exist")

type ctxDataKey int

var ctxKey = ctxDataKey(1)
var ctxValue = ctxDataKey(2)
var fbCtxKey = ctxDataKey(3)
var fbCtxValue = ctxDataKey(4)

type Provider interface {
	Set(k, v string) error //todo set byte array
	Get(k string) ([]byte, error)
}
