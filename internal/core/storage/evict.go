package storage

import (
	"container/list"

	"github.com/knvi/kvne/internal/config"
)

func EvictLRU() {
    element := lru.Back()
    if element != nil {
        Del(element.Value.(*Object).key)
    }
}

func EvictLFU() {
    minCounter := int(^uint(0) >> 1) // max int
    var lfuElement *list.Element

    for _, element := range Storage {
        obj := element.Value.(*Object)
        if obj.accessCounter < minCounter {
            minCounter = obj.accessCounter
            lfuElement = element
        }
    }

    if lfuElement != nil {
        Del(lfuElement.Value.(*Object).key)
    }
}

func AddToMap(v *Object) *list.Element {
	if config.EvictMode == config.EVICT_MODE_LRU {
		return lru.PushFront(v)
	} 
	
	// TODO: don't return an element if evict mode is none
	return lfu.PushFront(v) 
}

func DelFromMap(v *list.Element) {
	if config.EvictMode == config.EVICT_MODE_LRU {
		lru.Remove(v)
	} else if config.EvictMode == config.EVICT_MODE_LFU{
		lfu.Remove(v)
	}
}

func Evict() {
	if config.EvictMode == config.EVICT_MODE_LRU {
		EvictLRU()
	} else if config.EvictMode == config.EVICT_MODE_LFU{
		EvictLFU()
	}
}