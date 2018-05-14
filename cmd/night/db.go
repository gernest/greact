package main

import (
	"encoding/json"
	"path/filepath"
	"time"

	"github.com/dgraph-io/badger"
	"github.com/gernest/prom"
	"github.com/gernest/prom/api"
)

const (
	pendingTest    byte = 10
	runningTest    byte = 11
	completedTest  byte = 11
	maxPendingTime      = time.Hour
)

func openDatabase() (*badger.DB, error) {
	h, err := homePath()
	if err != nil {
		return nil, err
	}
	dir := filepath.Join(h, "data")
	return badger.Open(badger.Options{
		Dir:      dir,
		ValueDir: dir,
	})
}

func savePendingTest(db *badger.DB, ts *api.TestRequest) error {
	data, err := json.Marshal(ts)
	if err != nil {
		return err
	}
	return db.Update(func(tx *badger.Txn) error {
		err := tx.SetEntry(&badger.Entry{
			Key:       []byte(ts.ID),
			Value:     data,
			UserMeta:  pendingTest,
			ExpiresAt: uint64(time.Now().Add(maxPendingTime).Unix()),
		})
		if err != nil {
			return err
		}
		return tx.Commit(nil)
	})
}

func saveCompletedTest(db *badger.DB, id string, rs *prom.SpecResult) error {
	data, err := json.Marshal(rs)
	if err != nil {
		return err
	}
	return db.Update(func(tx *badger.Txn) error {
		err := tx.SetEntry(&badger.Entry{
			Key:      []byte(id),
			Value:    data,
			UserMeta: completedTest,
		})
		if err != nil {
			return err
		}
		return tx.Commit(nil)
	})
}
