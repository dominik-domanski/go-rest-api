package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABSE_URL")
	if databaseURL == "" {
		databaseURL = "user=postgres password=password dbname=restapi_test sslmode=disable"
	}
	os.Exit(m.Run())
}
