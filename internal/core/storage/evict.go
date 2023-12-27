package storage

func evictFirst() {
	for k := range Storage {
		Del(k)
		break
	}
}

func Evict() {
	evictFirst()
}