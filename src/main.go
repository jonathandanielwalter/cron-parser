package main

import (
	"cron-parser/ingestor"
	"fmt"
)

func main() {

	cron, err := ingestor.Ingest()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(cron.CreateCronOutput())
	}
}
