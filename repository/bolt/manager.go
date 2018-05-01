package bolt

import (
	"io"

	bolt "github.com/coreos/bbolt"
	"github.com/pkg/errors"
)

var openedDb *bolt.DB

// GetConnection returns an opened bolt instance
func GetConnection() (*bolt.DB, error) {
	if openedDb == nil {
		db, err := bolt.Open("syshealth.db", 0600, nil)
		if err != nil {
			return nil, err
		}
		openedDb = db
	}
	return openedDb, nil
}

// CloseConnection closes the opened connection, if any
func CloseConnection() error {
	if openedDb != nil {
		return openedDb.Close()
	}
	return nil
}

// Backup open the connection and write DB content to the specified writer
func Backup(w io.Writer) error {
	db, err := GetConnection()
	if err != nil {
		return errors.Wrap(err, "unable to get DB")
	}

	err = db.View(func(tx *bolt.Tx) error {
		_, err := tx.WriteTo(w)
		return err
	})

	if err != nil {
		return errors.Wrap(err, "unable to write DB content to writer")
	}

	return nil
}
