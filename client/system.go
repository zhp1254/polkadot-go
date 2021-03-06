package client

import (
	"github.com/zhex/polkadot-go/jsonrpc"
	"github.com/zhex/polkadot-go/types/rpccall"
	"github.com/zhex/polkadot-go/utils/decoder"
)

func createSystem(p *jsonrpc.WsProvider) *system {
	s := system{}
	s.provider = p
	s.section = "system"
	return &s
}

type system struct {
	rpcBase
}

func (s *system) Name() (string, error) {
	result, err := s.call("name", nil)
	return result.(string), err
}

func (s *system) Version() (string, error) {
	result, err := s.call("version", nil)
	return result.(string), err
}

func (s *system) Chain() (string, error) {
	result, err := s.call("chain", nil)
	return result.(string), err
}

func (s *system) Health() (*rpccall.Health, error) {
	result, err := s.call("health", nil)
	if err != nil {
		return nil, err
	}
	health := rpccall.Health{}
	err = decoder.MapDecode(result, &health)
	return &health, err
}

func (s *system) Peers() ([]rpccall.PeerInfo, error) {
	result, err := s.call("peers", nil)
	var peers []rpccall.PeerInfo
	err = decoder.MapDecode(result, &peers)
	return peers, err
}

func (s *system) NetworkState() (interface{}, error) {
	result, err := s.call("networkState", nil)
	return result, err
}

func (s *system) Properties() (*rpccall.ChainProperties, error) {
	result, err := s.call("properties", nil)
	if err != nil {
		return nil, err
	}
	properties := rpccall.ChainProperties{}
	err = decoder.MapDecode(result, &properties)
	return &properties, err
}
