package bq

import (
	"context"
	"go-training/log"
	"os"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

func getTables() map[int]string {
	projectID := os.Getenv("google_project")
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID)
	checkErrors(err, "connection error to pq")

	defer client.Close()
	q := client.Query("SELECT table_name FROM information_schema.tables WHERE table_schema='public'")
	it, err := q.Read(ctx)
	checkErrors(err, "invalid command to pg")
	tables := make(map[int]string)
	i := 0
	for {
		var row []bigquery.Value
		err := it.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Err("Error")
		}
		tables[i] = row[0].(string)
		i++
	}
	return tables

}

func checkErrors(err error, message string) {
	if err != nil {
		log.Err(message)
		os.Exit(1)
	}
}
