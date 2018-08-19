package cidl

import "fmt"

type autoErrorserrorsTcidl int

const (
	ErrAddCommunityGroupHasExist autoErrorserrorsTcidl = 5000 //用户信息已经添加，请勿重复提交
)

func (m autoErrorserrorsTcidl) Number() int { return int(m) }
func (m autoErrorserrorsTcidl) Message() string {
	switch m {

	case ErrAddCommunityGroupHasExist:
		return "用户信息已经添加，请勿重复提交"
	default:
		return "UNKNOWN_MESSAGE_autoErrorserrorsTcidl"
	}
}
func (m autoErrorserrorsTcidl) Name() string {
	switch m {

	case ErrAddCommunityGroupHasExist:
		return "ErrAddCommunityGroupHasExist"
	default:
		return "UNKNOWN_Name_autoErrorserrorsTcidl"
	}
}
func (m autoErrorserrorsTcidl) String() string {
	return fmt.Sprintf("[%d:%s]%s", m, m.Name(), m.Message())

}
