package behavior

import (
	"github.com/gorilla/mux"
	"net"
	"net/http"
	"strconv"
)

// 职责链模式

// 处理器接口
// MiddlewareHandler is the plugin middleware handler which wraps an existing http.MiddlewareHandler passed in.
// Its the responsibility of the MiddlewareHandler to call the next http.MiddlewareHandler in the chain.
type MiddlewareHandler func(http.Handler) http.Handler

// 管理 handler 的
type MiddlewareHandlers struct {
	mux *http.ServeMux

	hs []MiddlewareHandler
}

func NewMiddlewareHandlers() *MiddlewareHandlers {
	return &MiddlewareHandlers{
		hs:  make([]MiddlewareHandler, 0),
		mux: http.NewServeMux(),
	}
}

// 添加 handler
func (hm *MiddlewareHandlers) AddHandler(h MiddlewareHandler) {
	hm.hs = append(hm.hs, h)
}

// 运行 http 服务
func (hm *MiddlewareHandlers) Run(address string) error {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rspData := "not found"
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Length", strconv.Itoa(len(rspData)))
		_, _ = w.Write([]byte(rspData))
	})

	var h http.Handler
	h = r

	for i := len(hm.hs); i > 0; i-- {
		h = hm.hs[i-1](h)
	}

	hm.mux.Handle("/", h)

	l, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	if err := http.Serve(l, hm.mux); err != nil {
		return err
	}

	return nil
}
