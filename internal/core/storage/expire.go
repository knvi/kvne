package storage

import "time"

func HasExpired(obj *Object) bool {
	exp, ex := expires[obj]
	if !ex {
		return false
	}

	return exp <= int64(time.Now().UnixMilli())
}

func GetExpiration(obj *Object) (int64, bool) {
	exp, ex := expires[obj]
	return exp, ex
}

func SetExpiration(obj *Object, ttl int64) {
	expires[obj] = int64(time.Now().UnixMilli()) + int64(ttl)
}