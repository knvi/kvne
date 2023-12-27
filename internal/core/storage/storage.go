package storage

import (
	"github.com/knvi/kvne/internal/config"
)

type Object struct {
	Value	interface{}
	TypeEncoding uint8
}

var Storage map[string]*Object
var expires map[*Object]int64

func init() {
	Storage = make(map[string]*Object)
	expires = make(map[*Object]int64)
}

func Add(value interface{}, ttl int64, o_type uint8, o_enc uint8) *Object {
	obj := &Object{
		Value: value,
		TypeEncoding: o_type | o_enc,
	}

	if ttl > 0 {
		SetExpiration(obj, ttl)
	}

	return obj
}

func Get(k string) *Object {
	v := Storage[k]
	if v != nil {
		if HasExpired(v) {
			Del(k)
			return nil
		}
	}

	return v
}

func Put(k string, v *Object) {
	if len(Storage) >= config.KeyNumLimit {
		Evict()
	}

	Storage[k] = v
}

func Del(k string) bool {
	if obj, exist := Storage[k]; exist {
		delete(Storage, k)
		delete(expires, obj)
		return true
	}

	return false
}