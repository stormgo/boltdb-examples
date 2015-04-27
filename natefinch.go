// http://npf.io/2014/07/intro-to-boltdb-painless-performant-persistence/

package main

import (
    "fmt"
    "log"

    "github.com/boltdb/bolt"
)

var mynames = []byte("lastnames")

func main() {
    db, err := bolt.Open("bolt.db", 0644, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    key := []byte("michael")
    value := []byte("angerman")

    // store some data
    err = db.Update(func(tx *bolt.Tx) error {
        bucket, err := tx.CreateBucketIfNotExists(mynames)
        if err != nil {
            return err
        }

        err = bucket.Put(key, value)
        if err != nil {
            return err
        }
        return nil
    })

    if err != nil {
        log.Fatal(err)
    }

    // retrieve the data
    err = db.View(func(tx *bolt.Tx) error {
        bucket := tx.Bucket(mynames)
        if bucket == nil {
            return fmt.Errorf("Bucket %q not found!", mynames)
        }

        val := bucket.Get(key)
        fmt.Println(string(val))

        return nil
    })

    if err != nil {
        log.Fatal(err)
    }
}

// output:
// angerman
