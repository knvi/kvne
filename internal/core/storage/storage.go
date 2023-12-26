package storage

import (
	"time"

	"github.com/knvi/kvne/internal/config"
)

type Object struct {
	Value	interface{}
	Expire	int64
}

var Storage map[string]*Object

func init() {
	Storage = make(map[string]*Object)
}

func Add(value interface{}, expire int64) *Object {
	var expireAt int64 = -1
	if expire > 0 {
		expireAt = int64(time.Now().Unix()) + expire
	}

	obj := &Object{
		Value: value,
		Expire: expireAt,
	}

	return obj
}

func Get(k string) *Object {
	v := Storage[k]
	if v != nil {
		if v.Expire != -1 && v.Expire <= time.Now().UnixMilli() {
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
	if _, exist := Storage[k]; exist {
		delete(Storage, k)
		return true
	}

	return false
}