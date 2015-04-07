package protorpc

import (
	"io"
	"net/rpc"
	"sync"
)

type serverCodec struct {
	dec     io.Reader
	enc     io.Writer
	c       io.Closer
	req     serverRequest
	mutex   sync.Mutex
	seq     uint64
	pending map[uint64][]byte
}

type serverRequest struct {
}

func (r *serverRequest) reset() {

}

func NewServerCodec(conn io.ReadWriteCloser) rpc.ServerCodec {
	return nil
}

func (c *serverCodec) ReadRequestHeader(r *rpc.Request) error {
	return nil
}

func (c *serverCodec) ReadRequestBody(x interface{}) error {
	return nil
}

func (c *serverCodec) Close() error {
	return c.c.Close()
}

func (c *serverCodec) WriteResponse(r *rpc.Response, x interface{}) error {
	return nil
}

func ServeConn(conn io.ReadWriteCloser) {
	rpc.ServeCodec(NewServerCodec(conn))
}
