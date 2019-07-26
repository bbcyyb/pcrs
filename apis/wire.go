package apis

import (
	"github.com/bbcyyb/pcrs/apis/misc"
	"github.com/bbcyyb/pcrs/pkg/jwt"
	"github.com/google/wire"
)

func initializeMisc() (*misc.MiscController, error) {
	panic(wire.Build(misc.NewMiscController, misc.NewMiscService, jwt.NewJwtGenerator))
}
