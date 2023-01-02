package geecache

type PeerPicker interface {
	Pick(key string) (peer PeerGetter, ok bool)
}

type PeerGetter interface {
	Get(group string, key string) ([]byte, error)
}
