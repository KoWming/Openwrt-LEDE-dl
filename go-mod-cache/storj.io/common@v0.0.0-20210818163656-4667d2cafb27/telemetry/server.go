// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package telemetry

import (
	"context"
	"net"
	"syscall"

	"github.com/spacemonkeygo/monkit/v3"
	"github.com/zeebo/admission/v3"
	"github.com/zeebo/admission/v3/admproto"
)

var (
	mon = monkit.Package()
)

// Handler is called every time a new metric comes in.
type Handler interface {
	Metric(application, instance string, key []byte, val float64)
}

// HandlerFunc turns a func into a Handler.
type HandlerFunc func(application, instance string, key []byte, val float64)

// Metric implements the Handler interface.
func (f HandlerFunc) Metric(a, i string, k []byte, v float64) { f(a, i, k, v) }

// Server listens for incoming metrics.
type Server struct {
	conn net.PacketConn
}

// Addr returns the address the server is serving on.
func (s *Server) Addr() string {
	return s.conn.LocalAddr().String()
}

// Listen will start listening on addr for metrics.
func Listen(addr string) (*Server, error) {
	conn, err := net.ListenPacket("udp", addr)
	if err != nil {
		return nil, err
	}
	return &Server{conn: conn}, nil
}

// Close will stop listening.
func (s *Server) Close() error {
	return s.conn.Close()
}

// Serve will wait for metrics and call Handler h as they come in.
func (s *Server) Serve(ctx context.Context, h Handler) error {
	scconn, ok := s.conn.(syscall.Conn)
	if !ok {
		return Error.New("invalid conn: %T", s.conn)
	}
	rc, err := scconn.SyscallConn()
	if err != nil {
		return err
	}
	return admission.Dispatcher{
		Handler: handlerWrapper{h: h},
		Conn:    rc,
	}.Run(ctx)
}

// ListenAndServe combines Listen and Serve.
func ListenAndServe(ctx context.Context, addr string, h Handler) error {
	s, err := Listen(addr)
	if err != nil {
		return err
	}

	defer func() { _ = s.Close() }()
	return s.Serve(ctx, h)
}

type handlerWrapper struct {
	h Handler
}

var (
	handleTask = mon.Task()
)

func (h handlerWrapper) Handle(ctx context.Context, m *admission.Message) {
	finish := handleTask(nil)

	data, err := admproto.CheckChecksum(m.Data)
	if err != nil {
		finish(&err)
		return
	}
	r := admproto.NewReaderWith(m.Scratch[:])
	data, applicationB, instanceB, numHeaders, err := r.Begin(data)
	if err != nil {
		finish(&err)
		return
	}

	// Even though we don't use the headers, if they exist on the buffer we
	// need to read them off.
	for i := 0; i < numHeaders; i++ {
		data, _, _, err = r.NextHeader(data)
		if err != nil {
			finish(&err)
			return
		}
	}

	application, instance := string(applicationB), string(instanceB)
	var key []byte
	var value float64
	for len(data) > 0 {
		data, key, value, err = r.Next(data)
		if err != nil {
			finish(&err)
			return
		}
		h.h.Metric(application, instance, key, value)
	}

	finish(nil)
}
