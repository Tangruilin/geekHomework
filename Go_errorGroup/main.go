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
	g, ctx := errgroup.WithContext(context.Background())

	//	起一个http server
	g.Go(func() error {
		http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
			_, err := writer.Write([]byte("Hello World"))
			if err != nil {
				fmt.Printf("error is %+v", err)
			}
			return
		})
		serve := &http.Server{
			Addr:    "8080",
			Handler: nil,
		}
		err := serve.ListenAndServe()
		defer serve.Close()
		return err
	})

	g.Go(func() error {
		c := make(chan os.Signal)
		signal.Notify(c)
		//通过发送信号的方式，将函数阻塞在此处
		<-c
		return nil
	})
	err := g.Wait()
	fmt.Println(err)
	fmt.Println(ctx.Err())
	return
}
