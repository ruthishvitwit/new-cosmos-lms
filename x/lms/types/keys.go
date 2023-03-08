package types

import "strconv"

const (
	Modulename   = "lms"
	StoreKey     = Modulename
	RouterKey    = Modulename
	QuerierRoute = Modulename
)

var (
	AKey  = []byte{0x01}
	SKey  = []byte{0x02}
	LKey  = []byte{0x03}
	LcKey = []byte{0x04}
	AcKey = []byte{0x05}
)

func AdminKey(admin string) []byte {
	key := make([]byte, len(AKey)+len(admin))
	copy(key, AKey)
	copy(key[len(AKey):], []byte(admin))
	return key
}
func StudentKey(student string) []byte {
	key := make([]byte, len(SKey)+len(student))
	copy(key, SKey)
	copy(key[len(SKey):], []byte(student))
	return key
}
func LeaveKey(student string, leave int) []byte {
	leaveId := strconv.Itoa(leave)
	key := make([]byte, len(LKey)+len(student)+len(leaveId))
	copy(key, LKey)
	copy(key[len(LKey):], []byte(student))
	copy(key[len(LKey)+len(student):], []byte(leaveId))
	return key
}

func LeaveCounterKey(id string) []byte {
	key := make([]byte, len(LcKey)+len(id))
	copy(key, LcKey)
	copy(key[len(LcKey):], id)
	return key
}
func AcceptLeaveKey(id string) []byte {
	key := make([]byte, len(AcKey)+len(id))
	copy(key, AcKey)
	copy(key[len(AcKey):], id)
	return key
}
