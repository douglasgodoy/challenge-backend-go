package infra

import "log"

func runMigrations() {
	_, err := DbClient.Query(
		`CREATE TABLE IF NOT EXISTS people (
      id UUID PRIMARY KEY,
      name VARCHAR(100),
      nickname VARCHAR(32),
      stack VARCHAR(32)[],
      birthday VARCHAR(10)
      )`)

	if err != nil {
		log.Fatal(err)
	}
}
