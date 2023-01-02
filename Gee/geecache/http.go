package geecache

import (
	"fmt"
	"geecache/consistenthash"
	"io/ioutil"
	"log"
	"net/http"
	url2 "net/url"
	"strings"
	"sync"
)

const (
	defaultPath     = "/_geecache/"
	defaultReplicas = 50
)

type HttpPool struct {
	self        string
	basePath    string
	mu          sync.Mutex
	peers       *consistenthash.CHash
	httpGetters map[string]*httpGetter
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

func (p *HttpPool) Set(peers ...string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.peers = consistenthash.NewCHash(defaultReplicas, nil)
	p.peers.Add(peers...)
	p.httpGetters = make(map[string]*httpGetter)
	for _, peer := range peers {
		p.httpGetters[peer] = &httpGetter{baseUrl: peer + defaultPath}
	}
}

func (p *HttpPool) Pick(key string) (PeerGetter, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if peer := p.peers.Get(key); peer != "" && peer != p.self {
		log.Printf("select peer: %v", peer)
		return p.httpGetters[peer], true
	}
	return nil, false
}

type httpGetter struct {
	baseUrl string
}

func (h *httpGetter) Get(group string, key string) ([]byte, error) {
	url := fmt.Sprintf(
		"%v%v/%v",
		h.baseUrl,
		url2.QueryEscape(group),
		url2.QueryEscape(key),
	)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned: %v", res.StatusCode)
	}
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

var _ PeerPicker = (*HttpPool)(nil)
var _ PeerGetter = (*httpGetter)(nil)
