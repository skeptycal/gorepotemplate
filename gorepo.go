package gorepo

import (
	"context"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func GetClient() github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "... your access token ..."},
	)

	tc := oauth2.NewClient(ctx, ts)
	// get go-github client
	client := github.NewClient(tc)

}
