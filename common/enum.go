package common

type RoleType int32

const (
	USERROLE_USER RoleType = iota
	USERROLE_DESIGNER
	USERROLE_ADMIN
)
