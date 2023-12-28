package storage

import (
	"container/list"
	"time"

	"github.com/knvi/kvne/internal/config"
)

type Object struct {
	key           string
    Value         interface{}
    TypeEncoding  uint8
    lastAccessed  time.Time
    accessCounter int
}

var Storage map[string]*list.Element
var expires map[*Object]int64

var (
    lru = list.New()
    lfu = list.New()
)

func init() {
    Storage = make(map[string]*list.Element)
    expires = make(map[*Object]int64)
}

func NewObject(value interface{}, ttl int64, o_type uint8, o_enc uint8) *Object {
    if len(Storage) >= config.KeyNumLimit {
        Evict() // or EvictLFU() depending on your policy
    }

    obj := &Object{
        Value:        value,
        TypeEncoding: o_type | o_enc,
        lastAccessed: time.Now(),
        accessCounter: 1,
    }

    if ttl > 0 {
        SetExpiration(obj, ttl)
    }

    return obj
}

func Get(k string) *Object {
    element, ok := Storage[k]
    if !ok {
        return nil
    }

    obj := element.Value.(*Object)
    if HasExpired(obj) {
        Del(k)
        return nil
    }

    obj.lastAccessed = time.Now()
    obj.accessCounter++
    lru.MoveToFront(element)

    return obj
}

func Put(k string, v *Object) {
    if len(Storage) >= config.KeyNumLimit {
        EvictLRU() // or EvictLFU() depending on your policy
    }

	element := AddToMap(v)
    Storage[k] = element
}

func Del(k string) bool {
    if element, exist := Storage[k]; exist {
        delete(Storage, k)
        delete(expires, element.Value.(*Object))
        DelFromMap(element)
        return true
    }

    return false
}