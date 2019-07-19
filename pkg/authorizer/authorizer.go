package authorizer

import (
	"net/http"

	"github.com/bbcyyb/pcrs/infra"
	"github.com/casbin/casbin"
)

type BasicAuthorizer struct {
	enforcer *casbin.Enforcer
}

func NewBasicAuthorizer() *BasicAuthorizer {
	e := casbin.NewEnforcer(infra.C.AuthModelFile, infra.C.AuthPolicyFile)
	return &BasicAuthorizer{enforcer: e}
}

func (a *BasicAuthorizer) CheckPermission(user string, r *http.Request) bool {
	method := r.Method
	path := r.URL.Path
	return a.enforcer.Enforce(user, path, method)
}
