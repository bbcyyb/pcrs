package common

type RoleType int32

const (
	USERROLE_USER RoleType = 1 << iota
	USERROLE_DESIGNER
	USERROLE_ADMIN
)
