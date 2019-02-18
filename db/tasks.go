package db

import (
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks") // create the name key for our bucket
var db *bolt.DB

//Task is a todo struct
type Task struct {
	Key   int
	Value string
}

//Init sets up our bolt database
func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	//	db.Update updates the data base with the callback
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err //	returns any errors that occur while bucket being created
	})
}
