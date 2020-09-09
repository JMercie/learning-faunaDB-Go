package db

import (
	"log"

	"github.com/JMercie/faunaDB/config"
	f "github.com/fauna/faunadb-go/faunadb"
)

var (
	secret      = config.Config("SECRET")
	endpoint    = f.Endpoint("https://db.fauna.com")
	adminClient = f.NewFaunaClient(secret, endpoint)
	dbName      = config.Config("DBNAME")
)

// InitDB inits fauna database, checks if our DB exist, if dont, creates a new one.
func InitDB() {

	res, err := adminClient.Query(
		f.If(
			f.Exists(f.Database(dbName)),
			true,
			f.CreateDatabase(f.Obj{"name": dbName})),
	)
	if err != nil {
		panic(err)
	}
	if res != f.BooleanV(true) {
		log.Printf("Database Created: %s \n %s", dbName, res)
	} else {
		log.Printf("Database already Exists: %t\n with name: %s", res, dbName)
	}
}

// GetObjClient retrive dbSecretkey
func GetObjClient() {

	//var dbSecret string

	res, err := adminClient.Query(
		f.CreateKey(f.Obj{
			"database": f.Database(dbName),
			"role":     "server",
		}),
	)
	if err != nil {
		panic(err)
	}
	log.Println(res)
}
