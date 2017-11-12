package catapult

import (
	"net"
	"time"
)

var (
	// DialTimeout represents the maximum amount of time the network dialer can take.
	DialTimeout = 30 * time.Second

	// DialKeepAlive represents the maximum amount of time too keep alive the socket.
	DialKeepAlive = 30 * time.Second

	// TLSHandshakeTimeout represents the maximum amount of time that
	// TLS handshake can take defined in the default http.Transport.
	TLSHandshakeTimeout = 10 * time.Second

	// RequestTimeout represents the maximum about of time that
	// a request can take, including dial / request / redirect processes.
	RequestTimeout = 60 * time.Second

	// DefaultDialer defines the default network dialer.
	DefaultDialer = &net.Dialer{
		Timeout:   DialTimeout,
		KeepAlive: DialKeepAlive,
	}

	//DefaultTransport stores the default HTTP transport to be used.
	//DefaultTransport = NewDefaultTransport(DefaultDialer)
)
