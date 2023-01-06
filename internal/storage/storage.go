package storage

import (
	"DHT/internal/logger"
	"DHT/internal/utils"
	"encoding/base64"
	"github.com/tidwall/buntdb"
	"log"
	"time"
)

// Storage defines a persistent K/V storage.
type Storage struct {
	db *buntdb.DB
}

// NewStorage creates a K/V storage, persisting to the given data file.
func NewStorage(dataFile string) *Storage {
	const DATA_FOLDER string = "./data"
	utils.CheckAndMakeDir(DATA_FOLDER)
	if db, err := buntdb.Open(DATA_FOLDER + "/" + dataFile); err != nil {
		log.Fatal(err)
		return nil
	} else {
		return &Storage{db: db}
	}
}

// Put the key/value pair into the storage expiring in `ttl` seconds.
func (s *Storage) Put(key []byte, value []byte, ttl time.Duration) {
	logger.Logger.Infow("storage.Put", "key", string(key), "value", string(value), "ttl", ttl.Seconds())
	err := s.db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(encodeBytes(key), encodeBytes(value), &buntdb.SetOptions{Expires: true, TTL: ttl})
		return err
	})
	if err != nil {
		logger.Logger.Warnw("storage.Put error", "err", err)
	}
}

// Get finds the value for the given key, if any.
func (s *Storage) Get(key []byte) (valBytes []byte, ok bool) {
	defer logger.Logger.Infow("storage.Get", "key", string(key), "value", string(valBytes), "ok", ok)
	var val string
	err := s.db.View(func(tx *buntdb.Tx) (err error) {
		val, err = tx.Get(encodeBytes(key))
		if err != nil {
			return err
		}
		return nil
	})
	if err == nil {
		if data, err := decodeBytes(val); err != nil {
			logger.Logger.Warnw("storage.Get error", "err", err)
			return nil, false
		} else {
			return data, true
		}
	}
	//if err == buntdb.ErrNotFound {
	//	return nil, false
	//}
	//logger.Logger.Warnw("storage.Get error", "err", err)
	return nil, false
}

// encodeBytes encodes the data of []byte to a string.
func encodeBytes(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// encodeBytes decode the data of string to a []byte.
func decodeBytes(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}
