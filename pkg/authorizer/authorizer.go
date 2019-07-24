package authorizer

import (
	"net/http"

	"github.com/bbcyyb/pcrs/conf"
	"github.com/casbin/casbin"
)

type BasicAuthorizer struct {
	enforcer *casbin.Enforcer
}

type IAuthorizer interface {
	CheckPermission(string, *http.Request) bool
}

func NewBasicAuthorizer() *BasicAuthorizer {
	e := casbin.NewEnforcer(conf.C.Pkg.Authorizer.Model, conf.C.Pkg.Authorizer.Policy)
	return &BasicAuthorizer{enforcer: e}
}

func (a *BasicAuthorizer) CheckPermission(user string, r *http.Request) bool {
	method := r.Method
	path := r.URL.Path
	return a.enforcer.Enforce(user, path, method)
}
