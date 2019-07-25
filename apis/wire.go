package apis

import (
	"github.com/bbcyyb/pcrs/apis/misc"
	"github.com/bbcyyb/pcrs/pkg/jwt"
	"github.com/google/wire"
)

var superSet = wire.NewSet(misc.NewMiscController, wire.Bind(new(misc.IService), misc.NewMiscService()), jwt.NewJwt)

func initializeMisc() (*misc.MiscController, error) {
	panic(wire.Build(superSet))
}
