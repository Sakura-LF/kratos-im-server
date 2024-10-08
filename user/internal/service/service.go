package service

import "github.com/google/wire"

// ServicesProviderSet ProviderSet is service providers.
var ServicesProviderSet = wire.NewSet(NewUsersService)
