package main

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	"github.com/josephhilby/todoApi/api/todo/app/v1"
	"github.com/josephhilby/todoApi/api/todo/app/v1/v1connect"
	"github.com/steady-bytes/draft/pkg/chassis"
)

func main() {
	x := &v1.TodoItem{}
	fmt.Println(x)
}

type Rpc interface {
	chassis.RPCRegistrar
	v1connect.TodoServiceHandler
}

type controller struct{}

func (h *controller) RegisterRPC(server chassis.Rpcer) {
	pattern, handler := v1connect.NewTodoServiceHandler(h)
	server.AddHandler(pattern, handler, true)
}

func (c *controller) List(ctx context.Context, req *connect.Request[v1.ListRequest]) (*connect.Response[v1.ListResponse], error) {
	return connect.NewResponse(&v1.ListResponse{
		Items: []*v1.TodoItem{},
	}), nil
}
