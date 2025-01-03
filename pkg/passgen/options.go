package passgen

import (
	"golang.org/x/crypto/bcrypt"
	"strings"
)

const (
	defaultPasswordLength = 10
	defaultPasswordCount  = 10
	defaultCharset        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-=_+[]{}|;:,.<>?/`~"
)

type Options struct {
	passLen   int
	passCount int
	charset   string
	cost      int

	password string
}

type Option func(*Options)

func NewOptions(opts ...Option) *Options {
	options := &Options{
		passLen:   defaultPasswordLength,
		passCount: defaultPasswordCount,
		charset:   defaultCharset,

		cost: bcrypt.DefaultCost,
	}
	for _, o := range opts {
		o(options)
	}

	if strings.TrimSpace(options.password) != "" {
		options.passLen = len(options.password)
	}
	return options
}

func WithPassLen(l int) Option {
	return func(options *Options) {
		options.passLen = l
	}
}

func WithPassCount(c int) Option {
	return func(options *Options) {
		options.passCount = c
	}
}

func WithCharset(s string) Option {
	return func(options *Options) {
		options.charset = s
	}
}

func WithDefaultPassword(s string) Option {
	return func(options *Options) {
		options.password = s
	}
}

func WithCost(c int) Option {
	return func(options *Options) {
		switch {
		case c < bcrypt.MinCost:
			options.cost = bcrypt.MinCost
		case c > bcrypt.MaxCost:
			options.cost = bcrypt.MaxCost
		default:
			options.cost = c
		}
	}
}
