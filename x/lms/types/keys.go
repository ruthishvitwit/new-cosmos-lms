package types

const (
	Modulename   = "lms"
	StoreKey     = Modulename
	RouterKey    = Modulename
	QuerierRoute = Modulename
)

var (
	AKey = []byte{0x01}
	SKey = []byte{0x02}
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
