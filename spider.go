package ethspider

import (
	"github.com/cz-theng/czkit-go/log"
)

// Spider is the main service
type Spider struct {
	opts options
}

// Init create and init a zap SugarLogger
func (s *Spider) Init(opts ...Option) error {

	for _, opt := range opts {
		opt.apply(&s.opts)
	}
	return nil
}

// Start start the spider
func (s *Spider) Start() {
	log.Info("spider start")
}
