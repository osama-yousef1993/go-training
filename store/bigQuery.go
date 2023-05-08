package store

import (
	"context"
	"fmt"
	"sync"

	"cloud.google.com/go/bigquery"
)

type BQStore struct {
	*bigquery.Client
}

var (
	once    sync.Once
	bqStore *BQStore
)

func NewBQStore() (*BQStore, error) {
	if bqStore == nil {
		once.Do(func() {
			client, err := bigquery.NewClient(context.Background(), "project id")
			if err != nil {
				fmt.Println(" bigquery not connected")
			}
			var bqs BQStore
			bqs.Client = client
			bqStore = &bqs
		})
	}
	return bqStore, nil
}

func BQClose() {
	bqStore.Close()
}