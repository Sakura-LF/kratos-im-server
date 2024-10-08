package biz

import "github.com/google/wire"

// BizProviderSet ProviderSet is biz providers.
var BizProviderSet = wire.NewSet(NewUserUsecase)
