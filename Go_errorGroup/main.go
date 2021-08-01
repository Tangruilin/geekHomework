package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	// 定义一个上游的ctx
	ctx := context.Background()
	//根据上游的ctx确定一个下游存在的ctx
	ctx, cancel := context.WithCancel(ctx)
	//注册errgroup
	g, errCtx := errgroup.WithContext(ctx)
	//初始化一个http.server
	server := &http.Server{
		Addr: ":8080",
	}
	//	起一个http server
	g.Go(func() error {
		http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
			_, err := writer.Write([]byte("Hello World"))
			if err != nil {
				fmt.Printf("error is %+v", err)
			}
			return
		})
		//将服务监听在这里挂起
		return server.ListenAndServe()
	})
	//监控协程
	g.Go(func() error {
		//阻塞在这里
		<-errCtx.Done()
		return server.Shutdown(errCtx)
	})

	g.Go(func() error {
		c := make(chan os.Signal)
		signal.Notify(c)
		//通过发送信号的方式，将函数阻塞在此处
		for {
			select {
			case <-errCtx.Done():
				return errCtx.Err()
			case <-c:
				cancel()
			}
		}
		return nil
	})

	err := g.Wait()
	fmt.Println(err)
	fmt.Println(ctx.Err())
	return
}
