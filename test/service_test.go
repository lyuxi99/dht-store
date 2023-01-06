package test

import (
	"DHT/internal/chord"
	"DHT/internal/logger"
	"DHT/internal/service"
	"DHT/pkg/client"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"os"
	"strings"
	"testing"
	"time"
)

type ServiceTestSuite struct {
	suite.Suite
	servers []*service.Server
	ring    []int
}

func (s *ServiceTestSuite) SetupSuite() {
	if wd, err := os.Getwd(); err == nil {
		if strings.HasSuffix(wd, "test") {
			os.Chdir("..")
		}
	}

	os.RemoveAll("data")
	os.RemoveAll("logs")
	if err := logger.Init("test_service.log", zap.InfoLevel); err != nil {
		fmt.Println("logger.Init failed", "err", err)
		panic(err)
	}
}

func (s *ServiceTestSuite) TearDownSuite() {
	for _, server := range s.servers {
		server.Stop()
	}
	logger.Sync()
}
func (s *ServiceTestSuite) CreateServer(configurationFile string) *service.Server {
	server := service.NewServer(configurationFile)
	assert.NotEqual(s.T(), nil, server)
	s.servers = append(s.servers, server)
	return server
}

func (s *ServiceTestSuite) Test00_StartFirstServer() {
	go s.CreateServer("./config/node0/config0.ini").Serve()
	time.Sleep(time.Second * 2)
}

func (s *ServiceTestSuite) Test01_JoinMoreServer() {
	go s.CreateServer("./config/node1/config1.ini").Serve()
	time.Sleep(time.Second * 2)
	go s.CreateServer("./config/node2/config2.ini").Serve()
	time.Sleep(time.Second * 2)
	go s.CreateServer("./config/node3/config3.ini").Serve()
	time.Sleep(time.Second * 2)
	s.ring = []int{0, 3, 2, 1}
}

func (s *ServiceTestSuite) Test02_CheckPredecessorAndSuccessor() {
	for i := 0; i < 4; i++ {
		a, b := s.ring[i], s.ring[(i+1)%4]
		assert.Equal(s.T(), s.servers[b].P2pServer.RpcServer.Self.Addr, s.servers[a].P2pServer.RpcServer.Finger[0].Addr)
		assert.Equal(s.T(), s.servers[a].P2pServer.RpcServer.Self.Addr, s.servers[b].P2pServer.RpcServer.Predecessor.Addr)
	}
}
func (s *ServiceTestSuite) Test03_CheckStabilizedFingers() {
	time.Sleep(10 * time.Second)
	assert.Equal(s.T(), "127.0.0.1:7422", s.servers[0].P2pServer.RpcServer.Finger[chord.M-2].Addr)
	assert.Equal(s.T(), "127.0.0.1:7412", s.servers[0].P2pServer.RpcServer.Finger[chord.M-1].Addr)
	assert.Equal(s.T(), "127.0.0.1:7402", s.servers[1].P2pServer.RpcServer.Finger[chord.M-2].Addr)
	assert.Equal(s.T(), "127.0.0.1:7432", s.servers[1].P2pServer.RpcServer.Finger[chord.M-1].Addr)
}
func (s *ServiceTestSuite) Test10_ApiPutGet() {
	c := client.NewClient(s.servers[0].Params.ApiAddress)
	assert.NotNil(s.T(), c)
	key := []byte("keyyyy")
	value := []byte("valueee")
	_, ok, _ := c.Get(key)
	assert.False(s.T(), ok)

	c.Put(key, value, 3, 2)
	time.Sleep(time.Second)

	v, ok, _ := c.Get(key)
	assert.Equal(s.T(), value, v)
	assert.True(s.T(), ok)

	// test whether expired
	time.Sleep(time.Second * 3)
	_, ok, _ = c.Get(key)
	assert.False(s.T(), ok)

	c.Close()
}

func TestServiceTestSuit(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}
