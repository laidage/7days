package geecache

import (
	"log"
	"net/http"
	"strings"
)

const defaultPath = "/_geecache/"

type HttpPool struct {
	self     string
	basePath string
}

func NewPool(addr string) *HttpPool {
	return &HttpPool{
		self:     addr,
		basePath: defaultPath,
	}
}

func (p *HttpPool) log(format string, v ...interface{}) {

}
func (p *HttpPool) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Printf(req.URL.Path)
	if !strings.HasPrefix(req.URL.Path, p.basePath) {
		panic("HTTPPool serving unexpected path:")
	}
	parts := strings.SplitN(req.URL.Path[len(p.basePath):], "/", 2)
	if len(parts) < 2 {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	groupName := parts[0]
	key := parts[1]

	group := GetGroup(groupName)
	if group == nil {
		http.Error(w, "no such group: "+groupName, http.StatusNotFound)
	}
	v, err := group.Get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	//w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(v.CloneBytes())

}
