package main

import (
	"context"
	"fmt"
	"os"

	"github.com/austinorth/flag"
	"github.com/google/go-github/v27/github"
	"golang.org/x/oauth2"
)

func main() {
	token := flag.String("token", "", "Personal/Oauth2 github token")
	query := flag.String("query", "", "Query")
	flag.Parse()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: *token},
	)

	ctx := context.Background()
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	searchOpt := &github.SearchOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}
	reposSet := make(map[string]struct{})
	for {
		code, resp, err := client.Search.Code(ctx, *query, searchOpt)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, cr := range code.CodeResults {
			repoURL := *cr.Repository.HTMLURL
			if _, ok := reposSet[repoURL]; !ok {
				reposSet[repoURL] = struct{}{}
			}
		}
		if resp.NextPage == 0 {
			break
		}
		searchOpt.Page = resp.NextPage
	}

	for r := range reposSet {
		fmt.Println(r)
	}
}

// test for checking other API capabilities, unused
// func listAllRepos(ctx context.Context, client *github.Client) {
// 	opt := &github.RepositoryListByOrgOptions{
// 		Type:        "private",
// 		ListOptions: github.ListOptions{PerPage: 10},
// 	}
// 	log.Printf("opt: %+v", opt)

// 	var allRepos []*github.Repository
// 	for {
// 		repos, resp, err := client.Repositories.ListByOrg(ctx, "Shopify", opt)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		allRepos = append(allRepos, repos...)
// 		if resp.NextPage == 0 {
// 			break
// 		}
// 		opt.Page = resp.NextPage
// 	}

// 	for _, r := range allRepos {
// 		log.Printf("repo: %+v", *r.HTMLURL)
// 	}
// }
