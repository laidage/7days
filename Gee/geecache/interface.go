package geecache

import pb "geecache/pb"

type PeerPicker interface {
	Pick(key string) (peer PeerGetter, ok bool)
}

type PeerGetter interface {
	Get(in *pb.Request, out *pb.Response) error
}
