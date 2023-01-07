package core

const (
	defaultAddr = ":8080"
)

type Option struct {
	F func(o *Options)
}

type Options struct {
	Addr string
}

func NewOptions(opts ...Option) *Options {
	options := &Options{
		Addr: defaultAddr,
	}
	options.Apply(opts)
	return options
}

func (o *Options) Apply(opts []Option) {
	for _, opt := range opts {
		opt.F(o)
	}
}

func WithHostPort(port string) Option {
	return Option{F: func(o *Options) {
		o.Addr = port
	}}
}
