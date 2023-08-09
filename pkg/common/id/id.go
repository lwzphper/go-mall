package id

// Identifier Type设计模式

type MemberID uint64

func (m MemberID) Uint64() uint64 {
	return uint64(m)
}
