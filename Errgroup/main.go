package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

var urls = []string{"https://example.com", "https://example.org", "https://example.net"}

func main() {
	ctx := context.Background()

	g, qcxt := errgroup.WithContext(ctx)
	g.SetLimit(2) // кол-во используемых горутин одновременно

	for _, url := range urls {
		g.Go(func() error {
			return isAvailable(qcxt, url)
		})
	}

	if err := g.Wait(); err != nil {
		log.Fatalf("Some resourse is not avaiable: %v", err)
	} else {
		log.Infof("All resourses are avaiable")
	}

}

func isAvailable(ctx context.Context, url string) error {
	c := http.Client{}

	request, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}

	response, err := c.Do(request)
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("wrong status code - %d for url - %s", response.StatusCode, url)
	}
	time.Sleep(5 * time.Second)
	return nil
}
