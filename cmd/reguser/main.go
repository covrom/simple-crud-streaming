package main

import (
	"context"
	"os"
	"os/signal"
	"sync"

	"github.com/covrom/simple-crud-streaming/api/handler"
	"github.com/covrom/simple-crud-streaming/api/server"
	"github.com/covrom/simple-crud-streaming/app/repos/user"
	"github.com/covrom/simple-crud-streaming/app/starter"
	"github.com/covrom/simple-crud-streaming/db/mem/usermemstore"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	ust := usermemstore.NewUsers()
	a := starter.NewApp(ust)
	us := user.NewUsers(ust)
	h := handler.NewRouter(us)
	srv := server.NewServer(":8000", h)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go a.Serve(ctx, wg, srv)

	<-ctx.Done()
	cancel()
	wg.Wait()
}
