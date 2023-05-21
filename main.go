package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"google.golang.org/grpc/metadata"

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
	ctx = metadata.AppendToOutgoingContext(ctx, "x-goog-api-key", os.Getenv("GOOGLE_GENERATIVE_AI_API_KEY"))
	dc, err := generativelanguage.NewDiscussClient(ctx, option.WithTokenSource(oauth2.StaticTokenSource(&oauth2.Token{})))
	if err != nil {
		return err
	}
	defer dc.Close()

	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	req := &generativelanguagepb.GenerateMessageRequest{
		Model: "models/chat-bison-001",
		Prompt: &generativelanguagepb.MessagePrompt{
			Messages: []*generativelanguagepb.Message{{
				Author:  "user",
				Content: string(in),
			},
			},
		},
	}
	resp, err := dc.GenerateMessage(ctx, req)
	if err != nil {
		return err
	}
	_ = resp
	fmt.Println(resp)
	return nil
}

func listModels(ctx context.Context) error {
	c, err := generativelanguage.NewModelClient(ctx, option.WithTokenSource(oauth2.StaticTokenSource(&oauth2.Token{})))
	if err != nil {
		return err
	}
	mIter := c.ListModels(ctx, &generativelanguagepb.ListModelsRequest{})
	for {
		m, err := mIter.Next()
		fmt.Println(m, err)
		if err != nil {
			break
		}
	}
	return nil
}
