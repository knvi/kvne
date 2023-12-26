package storage

func evictFirst() {
	for k := range Storage {
		delete(Storage, k)
		break
	}
}

func Evict() {
	evictFirst()
}