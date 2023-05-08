package store

import (
	bq "cloud.google.com/go/bigquery"
	"context"
	"fmt"
	"github.com/osama-yousef1993/go-training/log"
	"google.golang.org/api/iterator"
	"os"
)

func BqInfoSchema() ([]bq.Value, error) {
	querystr := fmt.Sprintf("SELECT * FROM `%s`.%s.INFORMATION_SCHEMA.TABLES;", os.Getenv("MY_PROJECT"), os.Getenv("MY_DATASET"))
	ctx := context.Background()

	client, err := bq.NewClient(ctx, os.Getenv("GCP_PROJECT_ID"))

	if err != nil {
		log.LogErr(err.Error())
		return nil, err
	}
	defer client.Close()

	q := client.Query(querystr)

	it, err := q.Read(ctx)
	if err != nil {
		log.LogErr(err.Error())
		return nil, err
	}
	var values []bq.Value
	for {
		err := it.Next(&values)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.LogInfo(err.Error())
		}
		fmt.Println(values)
	}
	return values, nil
}
