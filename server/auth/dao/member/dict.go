package member

type MemberStatus int8

const (
	STATUS_DISABLE MemberStatus = iota
	STATUS_ENABLE
)

var statusNameMap = map[MemberStatus]string{
	STATUS_DISABLE: "禁用",
	STATUS_ENABLE:  "启用",
}

func (m MemberStatus) String() string {
	name, ok := statusNameMap[m]
	if !ok {
		return "禁用"
	}
	return name
}

type Gender int8

const (
	GENDER_UNKOWN Gender = iota
	GENDER_MAN
	GENDER_WOMAN
)

var genderNameMap = map[Gender]string{
	GENDER_UNKOWN: "未知",
	GENDER_MAN:    "男",
	GENDER_WOMAN:  "女",
}

func (g Gender) String() string {
	name, ok := genderNameMap[g]
	if !ok {
		return "未知"
	}
	return name
}

type SourceType int8
