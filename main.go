package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	generativelanguagepb "google.golang.org/genproto/googleapis/ai/generativelanguage/v1beta2"
	generativelanguage "google.golang.org/google/ai/generativelanguage/v1beta2"
)

func main() {
	flag.Parse()
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	ctx := context.Background()
	c, err := generativelanguage.NewDiscussClient(ctx)
	if err != nil {
		return err
	}
	defer c.Close()

	req := &generativelanguagepb.GenerateMessageRequest{
		Model: "name=chat-bison@001",
		Prompt: &generativelanguagepb.MessagePrompt{
			Messages: []*generativelanguagepb.Message{{
				Author:  "travis",
				Content: "hello!",
			},
			},
		},
	}
	resp, err := c.GenerateMessage(ctx, req)
	if err != nil {
		return err
	}
	_ = resp
	fmt.Println(resp)
	return nil
}
