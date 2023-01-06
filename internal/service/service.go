package service

import (
	"DHT/internal/api"
	"DHT/internal/chord"
	"DHT/internal/logger"
	"DHT/internal/storage"
	"DHT/internal/utils"
	"fmt"
	"go.uber.org/zap"
	"gopkg.in/ini.v1"
	"log"
)

// Params defines the parameters for a server.
type Params struct {
	Bootstrapper           string
	ApiAddress, P2pAddress string
	LogFile, DataFile      string
	CACert                 string
	ServerCert, ServerKey  string
}

// readParams reads the parameters out from a configuration file.
func readParams(configurationFile string) (*Params, error) {
	cfg, err := ini.Load(configurationFile)
	if err != nil {
		return nil, err
	}
	return &Params{
		Bootstrapper: cfg.Section("dht").Key("bootstrapper").String(),
		P2pAddress:   cfg.Section("dht").Key("p2p_address").String(),
		ApiAddress:   cfg.Section("dht").Key("api_address").String(),
		LogFile:      cfg.Section("dht").Key("log_file").String(),
		DataFile:     cfg.Section("dht").Key("data_file").String(),
		CACert:       cfg.Section("dht").Key("ca_cert").String(),
		ServerCert:   cfg.Section("dht").Key("hostcert").String(),
		ServerKey:    cfg.Section("").Key("hostkey").String(),
	}, nil
}

// Server defines a DHT server, consisting of Params, api.ApiServer, chord.P2pServer and storage.Storage.
type Server struct {
	Params    *Params
	ApiServer *api.ApiServer
	P2pServer *chord.P2pServer
	Storage   *storage.Storage
}

// NewServer creates a DHT server from the configuration file.
func NewServer(configurationFile string) *Server {
	params, err := readParams(configurationFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", params)
	if !utils.Exists(params.CACert) || !utils.Exists(params.ServerCert) || !utils.Exists(params.ServerKey) {
		log.Fatal("CACert, ServerCert or ServerKey doesn't exists")
	}

	if err := logger.Init(params.LogFile, zap.InfoLevel); err != nil {
		log.Fatal("logger.Init error", err)
	}

	server := &Server{
		Params: params,
	}
	server.Storage = storage.NewStorage(params.DataFile)
	server.P2pServer = chord.NewP2pServer(server.Storage, params.P2pAddress, params.CACert, params.ServerCert, params.ServerKey)
	server.ApiServer = api.NewApiServer(server.P2pServer, params.ApiAddress)
	return server
}

// Serve runs the DHT server, starting the API server and P2P server.
func (s *Server) Serve() {
	logger.Logger.Infow("Start Server", "params", s.Params)
	go func() {
		fmt.Println("api.Serve")
		if err := s.ApiServer.Serve(); err != nil {
			log.Fatal("api.Serve failed", err)
		}
	}()

	fmt.Println("chord.Serve")
	err := s.P2pServer.Serve(s.Params.Bootstrapper)
	if err != nil {
		log.Fatal("chord.Serve failed", err)
	}
}

// Stop stops the DHT server gracefully.
func (s *Server) Stop() {
	s.ApiServer.Stop()
	s.P2pServer.Stop()
}
