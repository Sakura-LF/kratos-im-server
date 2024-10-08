package server

import (
	"github.com/google/wire"
)

// ServerProviderSet ProviderSet is server providers.
var ServerProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer)
