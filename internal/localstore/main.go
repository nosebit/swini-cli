package localstore

import (
	"encoding/json"
	"os"

	"swini-cli/internal/config"

	"go.etcd.io/bbolt"
)

type Store struct {
	Account Account `json:"account"`
}

type Account struct {
	ID     string `json:"id"`
	PvtKey string `json:"pvtkey"`
}

const (
	bucketName = "store"
	storeKey   = "root"
)

var (
	db    *bbolt.DB
	store *Store
)

// Open initializes the DB connection.
func Open() error {
	if db == nil {
		env := os.Getenv("GO_ENV")
		dbFile := "data.db"
		if env == "development" {
			dbFile = "data.dev.db"
		}

		cfg, err := config.Load()
		if err != nil {
			return err
		}

		db, err = bbolt.Open(cfg.DataDir+"/"+dbFile, 0600, nil)
		if err != nil {
			return err
		}
	}

	return nil
}

// Close closes the DB connection.
func Close() error {
	if db == nil {
		return nil
	}
	return db.Close()
}

func Load() (*Store, error) {
	Open()

	if store == nil {
		store = &Store{}

		err := db.View(func(tx *bbolt.Tx) error {
			b := tx.Bucket([]byte(bucketName))
			if b == nil {
				// no bucket yet -> empty store
				return nil
			}
			data := b.Get([]byte(storeKey))
			if data == nil {
				return nil
			}
			return json.Unmarshal(data, store)
		})
		if err != nil {
			return nil, err
		}
	}

	return store, nil
}

func (s *Store) Save() error {
	return db.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}
		data, err := json.Marshal(s)
		if err != nil {
			return err
		}
		return b.Put([]byte(storeKey), data)
	})
}

// Reset wipes the store in memory and on disk.
func (s *Store) Reset() error {
	// reset the in-memory data
	*s = Store{}

	// persist the empty state
	return s.Save()
}
