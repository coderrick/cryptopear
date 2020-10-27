package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"os"

	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/textileio/go-threads/core/thread"
	"github.com/textileio/textile/v2/cmd"
	"github.com/textileio/textile/v2/mail/local"
)

func main() {
	// Setup the mail lib
	mail := local.NewMail(cmd.NewClients("api.textile.io:443", false), local.DefaultConfConfig())

	// Create a libp2p identity (this can be any thread.Identity)
	privKey, _, err := crypto.GenerateEd25519Key(rand.Reader)
	id := thread.NewLibp2pIdentity(privKey)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	// Create a new mailbox with config
	mailbox, err := mail.NewMailbox(context.Background(), local.Config{
		Path:      "/Users/derrick/Desktop/CryptoPear/.textile/mail", // Usually a global location like ~/.textile/mail
		Identity:  id,
		APIKey:    "b25pwri3ftrdzp3pjltji7ft5uu",
		APISecret: "b55efcp3ugi6rfwuwqfwncdx7mydactqd55plmpa",
	})

	fmt.Println(mailbox.Identity().GetPublic().String())
}
