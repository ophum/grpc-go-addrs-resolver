package grpcgoaddrsresovler

import (
	"strings"

	"google.golang.org/grpc/resolver"
)

type addrsResolverBuilder struct{}

var _ resolver.Builder = (*addrsResolverBuilder)(nil)

func (b *addrsResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, _ resolver.BuildOptions) (resolver.Resolver, error) {
	r := &addrsResolver{
		target: target,
		cc:     cc,
		addrsStore: map[string][]string{
			target.Endpoint(): strings.Split(target.Endpoint(), ","),
		},
	}
	r.start()
	return r, nil
}

func (b *addrsResolverBuilder) Scheme() string {
	return "addrs"
}

type addrsResolver struct {
	target     resolver.Target
	cc         resolver.ClientConn
	addrsStore map[string][]string
}

var _ resolver.Resolver = (*addrsResolver)(nil)

func (r *addrsResolver) start() {
	addrStrs := r.addrsStore[r.target.Endpoint()]
	addrs := make([]resolver.Address, len(addrStrs))
	for i, s := range addrStrs {
		addrs[i] = resolver.Address{
			Addr: s,
		}
	}
	r.cc.UpdateState(resolver.State{Addresses: addrs})
}

func (r *addrsResolver) ResolveNow(resolver.ResolveNowOptions) {}
func (r *addrsResolver) Close()                                {}

func init() {
	resolver.Register(&addrsResolverBuilder{})
}
