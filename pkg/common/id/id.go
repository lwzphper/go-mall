package id

// Identifier Type设计模式

// MemberID 会员id
type MemberID uint64

func (m MemberID) Uint64() uint64 {
	return uint64(m)
}

// AddressID 收货地址id
type AddressID uint64

func (a AddressID) Uint64() uint64 {
	return uint64(a)
}
