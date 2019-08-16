package middleware

import (
	"context"
	"errors"
	"net"

	"github.com/insomniacslk/dhcp/dhcpv4"
	"github.com/ppacher/dhcp-ng/pkg/lease"
)

// Handler is a middleware that handles incoming DHCPv4 requests
type Handler interface {
	Serve(ctx *Context, request *dhcpv4.DHCPv4)
}

// HandleFunc is a middleware that handles incoming DHCPv4 requests
// Each HandleFunc automatically satisfies the Handler interface
type HandleFunc func(ctx *Context, request *dhcpv4.DHCPv4)

// Serve implements the Handler interface
func (h HandleFunc) Serve(ctx *Context, request *dhcpv4.DHCPv4) {
	h(ctx, request)
}

// Controller is responsible for managing a subnet
type Controller interface {
	// Database returns the lease database used by the subnet controller
	Database() lease.Database

	// Serve the given DHCP request message
	Serve(ctx context.Context, peer net.Addr, request *dhcpv4.DHCPv4) (*dhcpv4.DHCPv4, error)
}

// Option is a controller option
type Option func(c *ctrl)

// WithDatabase configures the lease database to use
func WithDatabase(db lease.Database) Option {
	return func(c *ctrl) {
		c.db = db
	}
}

// NewController creates a new subnet controller
func NewController(options ...Option) Controller {
	c := &ctrl{}

	for _, fn := range options {
		fn(c)
	}

	return c
}

type ctrl struct {
	db lease.Database
}

func (c *ctrl) Database() lease.Database {
	return c.db
}

func (c *ctrl) Serve(ctx context.Context, peer net.Addr, request *dhcpv4.DHCPv4) (*dhcpv4.DHCPv4, error) {
	return nil, errors.New("not yet implemented")
}
