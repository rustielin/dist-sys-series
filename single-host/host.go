package main

import (
	"context"
	"crypto/rand"
	"fmt"
	peerstore "github.com/libp2p/go-libp2p-peerstore"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"

)

func main() {

	// context coverns the lifetime of the libp2p node
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// construct simple host with default settings
	host, err := libp2p.New(ctx)
	if err != nil {
		panic(err)
	}

	peerInfo := &peerstore.PeerInfo{
		ID: host.ID(),
		Addrs: host.Addrs(),
	}

	fmt.Printf("Hello World, my host ID is %s\n", host.ID())
	fmt.Printf("PeerInfo: %v\n", peerInfo)

	// libp2p.New can take options as well

	// set manual keypair
	priv, _, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		panic(err)
	}

	host2, err := libp2p.New(ctx,
		libp2p.Identity(priv),
		libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/9000"),
	)
	if err != nil {
		panic(err)
	}

	peerInfo2 := &peerstore.PeerInfo{
		ID: host2.ID(),
		Addrs: host2.Addrs(),
	}
	fmt.Printf("Hello World, my second host ID is %s\n", host2.ID())
	fmt.Printf("PeerInfo: %v\n", peerInfo2)
}
