// Package lease provides utility methods for IP address lease management
package lease

import (
	"fmt"
	"net"
	"time"
)

type Client struct {
	// HwAddr is the hardware address of the client that received the lease
	HwAddr net.HardwareAddr

	// Hostname is the hostname of the client that received the lease. This field
	// may be empty
	Hostname string
}

// Lease describes an IPv4 address that has been leased to a client
type Lease struct {
	// Client is the client that received the lease
	Client

	// Expires holds the timestamp in seconds when the lease is going to
	// expire
	Expires int64

	// Address holds the address that has been leased to the client
	Address net.IPAddr
}

// ExpiredAt returns true if the lease was or will be expired at t
func (l *Lease) ExpiredAt(t time.Time) bool {
	return t.After(time.Unix(l.Expires, 0))
}

// Expired returns true if the lease has already been expired
func (l *Lease) Expired() bool {
	return l.ExpiredAt(time.Now())
}

// String implements fmt.Stringer
func (l *Lease) String() string {
	suffix := ""
	if l.Expired() {
		suffix = "; expired"
	}
	return fmt.Sprintf("%s (client=%s%s)", l.Address.String(), l.HwAddr, suffix)
}
