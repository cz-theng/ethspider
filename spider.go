package ethspider

import (
	"context"

	"github.com/cz-theng/czkit-go/log"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Spider is the main service
type Spider struct {
	opts   options
	client *ethclient.Client
	ctx    context.Context
}

// Init create and init a zap SugarLogger
func (s *Spider) Init(opts ...Option) (err error) {
	err = nil
	for _, opt := range opts {
		opt.apply(&s.opts)
	}
	s.ctx = context.Background()
	return
}

// Start start the spider
func (s *Spider) Start() (err error) {
	err = nil
	log.Info("spider start")
	s.client, err = ethclient.DialContext(s.ctx, s.opts.rpcAddr)
	if err != nil {
		log.Error("connect rpc error %s", err.Error())
		return err
	}

	latest, err := s.client.BlockNumber(s.ctx)
	if err != nil {
		log.Error("get latest block number error %s", err.Error())
		return err
	}
	log.Info(" get latest block number:%d", latest)
	return
}
