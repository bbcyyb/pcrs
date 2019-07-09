package common

type RoleEnum int32

const (
	USERROLE_USER RoleEnum = iota
	USERROLE_DESIGNER
	USERROLE_ADMIN
)

var roleEnumFlags = map[RoleEnum]string{
	USERROLE_USER:     "user",
	USERROLE_DESIGNER: "designer",
	USERROLE_ADMIN:    "admin",
}

func GetRoleEnumMessage(role RoleEnum) string {
	if msg, ok := roleEnumFlags[role]; ok {
		return msg
	}

	return roleEnumFlags[USERROLE_USER]
}
