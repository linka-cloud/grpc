package service

import (
	"net/http"
	"time"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
)

var defaultWebOptions = []grpcweb.Option{
	grpcweb.WithWebsockets(true),
	grpcweb.WithWebsocketOriginFunc(func(req *http.Request) bool {
		return true
	}),
	grpcweb.WithCorsForRegisteredEndpointsOnly(false),
	grpcweb.WithOriginFunc(func(origin string) bool {
		return true
	}),
	grpcweb.WithWebsocketPingInterval(time.Second),
}

func (s *service) grpcWeb(opts ...grpcweb.Option) error {
	if !s.opts.grpcWeb {
		return nil
	}
	h := grpcweb.WrapServer(s.server, opts...)
	for _, v := range grpcweb.ListGRPCResources(s.server) {
		if s.opts.grpcWebPrefix != "" {
			s.mux.Handle(s.opts.grpcWebPrefix+v, http.StripPrefix(s.opts.grpcWebPrefix, h))
		} else {
			s.mux.Handle(v, h)
		}
	}
	return nil
}