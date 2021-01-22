package ethspider

type options struct {
	rpcAddr string
}

func defaultOptions() options {
	return options{
		rpcAddr: "",
	}
}

// Option configures how we set up the spider.
type Option interface {
	apply(*options)
}

type funcOption struct {
	f func(*options)
}

func (fdo *funcOption) apply(do *options) {
	fdo.f(do)
}

func newFuncOption(f func(*options)) *funcOption {
	return &funcOption{
		f: f,
	}
}

// WithRPCAddr determines which addr to connect to
func WithRPCAddr(addr string) Option {
	return newFuncOption(func(o *options) {
		o.rpcAddr = addr
	})
}
