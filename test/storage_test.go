package test

import (
	"DHT/internal/logger"
	"DHT/internal/storage"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"os"
	"testing"
	"time"
)

type StorageTestSuite struct {
	suite.Suite
	storage *storage.Storage
}

func (s *StorageTestSuite) SetupSuite() {
	//os.Remove("data/data.db")
	os.RemoveAll("data")
	os.RemoveAll("logs")
	if err := logger.Init("test_storage.log", zap.DebugLevel); err != nil {
		fmt.Println("logger.Init failed", "err", err)
		panic(err)
	}

	s.storage = storage.NewStorage("data.db")
}

func (s *StorageTestSuite) TearDownSuite() {
	logger.Sync()
}

func (s *StorageTestSuite) Test00() {
	key := []byte("key1")
	value := []byte("value1")
	key2 := []byte("key2")
	s.storage.Put(key, value, time.Second)
	resp, ok := s.storage.Get(key)
	assert.Equal(s.T(), value, resp)
	assert.Equal(s.T(), true, ok)
	time.Sleep(time.Second + time.Millisecond)
	resp, ok = s.storage.Get(key)
	assert.Equal(s.T(), []byte(nil), resp)
	assert.Equal(s.T(), false, ok)
	resp, ok = s.storage.Get(key2)
	assert.Equal(s.T(), []byte(nil), resp)
	assert.Equal(s.T(), false, ok)
}

func TestStorageTestSuit(t *testing.T) {
	suite.Run(t, new(StorageTestSuite))
}
