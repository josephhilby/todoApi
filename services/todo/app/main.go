package main

import (
	"connectrpc.com/connect"
	"context"
	"github.com/josephhilby/todoApi/api/todo/app/v1"
	"github.com/josephhilby/todoApi/api/todo/app/v1/v1connect"
	"github.com/steady-bytes/draft/pkg/chassis"
	"github.com/steady-bytes/draft/pkg/loggers/zerolog"
	"github.com/steady-bytes/draft/pkg/repositories/postgres/bun"
)

func main() {
	logger := zerolog.New()
	db := bun.New("")
	rpcHandler := &controller{}
	chassis.New(logger).WithRPCHandler(rpcHandler).Register(chassis.RegistrationOptions{
		Namespace: "todo",
	}).WithRepository(db).Start()
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
		Items: []*v1.TodoItem{
			{
				Title: "test Title",
				Body:  "test Body",
			},
		},
	}), nil
}
