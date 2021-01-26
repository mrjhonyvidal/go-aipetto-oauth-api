package cassandra

import (
	"fmt"
	"github.com/gocql/gocql"
	"log"
	"time"
)

var (
	cluster *gocql.ClusterConfig
)

func init() {
	// Connect to Cassandra cluster:
	cluster = gocql.NewCluster("go-oauth-cassandra-db")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
}

func GetSession() (*gocql.Session, error) {
	return cluster.CreateSession()
}

func CheckConnectionOnStartApplication() {
	retryCount := 10
	for {
		if session, dbErr := GetSession(); dbErr != nil {
			if retryCount == 0 {
				log.Fatalf("It was not able to establish connection to db.")
			}
			log.Printf(fmt.Sprintf("Could not connect to database. Wait 5 seconds. %d retries left...", retryCount))
			retryCount--
			time.Sleep(5 * time.Second)
		}else{
			session.Close()
			break;
		}
	}
}