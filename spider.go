package ethspider

import (
	"context"
	"math/big"
	"reflect"
	"sync/atomic"
	"unsafe"

	"github.com/cz-theng/czkit-go/log"
	"github.com/cz-theng/ethspider/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

	s.pullBlocks(latest)
	return
}

func (s *Spider) pullBlocks(latest uint64) {
	for {
		if latest <= 1 {
			break
		}
		num := latest
		latest--
		blk, err := s.pullBlock(num)
		if err != nil {
			continue
		}
		log.Info("get block[%d]:%s", blk.NumberU64(), blk.Hash().Hex())
		s.dealBlock(blk)
	}
}

func (s *Spider) pullBlock(num uint64) (blk *types.Block, err error) {
	bigNum := big.NewInt(0)
	bigNum.SetUint64(num)
	blk, err = s.client.BlockByNumber(s.ctx, bigNum)
	return
}

func (s *Spider) dealBlock(blk *types.Block) {
	if nil == blk {
		return
	}

	for _, trx := range blk.Transactions() {
		s.dealTransaction(trx)
	}
}

func (s *Spider) dealTransaction(trx *types.Transaction) {
	if nil == trx {
		return
	}
	fromPtr := reflect.ValueOf(trx).Elem().FieldByName("from")
	fromPtr = reflect.NewAt(fromPtr.Type(), unsafe.Pointer(fromPtr.UnsafeAddr())).Elem()
	if v, ok := fromPtr.Interface().(atomic.Value); ok {
		value := v.Load()
		if fromAddr, ok := value.(common.Address); ok {
			log.Info("trx[%s]from addr :%s", trx.Hash().Hex(), fromAddr.Hex())
		}
	}
}
