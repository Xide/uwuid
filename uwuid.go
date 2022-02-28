package uwuid

import (
	"crypto/rand"
	"errors"
)

// C-C-Cutest unyique identifiew evew *sweats* !
// 128 bits of wandomnyess, up t-to 42 UwU's ?-?-?!?1
type UwUID struct {
	bts []byte
}

// Much efficient, such wow !
// so w-wittwe memowy ÚwÚ footpwint
func (id *UwUID) String() string {
	res := ""
	for _, e := range id.bts {
		for i := 0; i < 8; i++ {
			if e&128 > 0 {
				res += "U"
			} else {
				res += "w"
			}
			e <<= 1
		}
	}
	return res
}

func NewUwUID() (*UwUID, error) {
	token := make([]byte, 16)
	n, err := rand.Read(token)
	if err != nil {
		return nil, err
	}
	if n != 16 {
		return nil, errors.New("could not read 16 random bytes")
	}
	return &UwUID{
		bts: token,
	}, nil
}

func FromBytes(b []byte) (*UwUID, error) {
	if len(b) != 16 {
		return nil, errors.New("invalid number of bytes")
	}
	return &UwUID{
		bts: b,
	}, nil
}

func FromString(s string) (*UwUID, error) {
	token := make([]byte, 16)

	if len(s) != 128 {
		return nil, errors.New("invalid size")
	}
	for i, b := range s {
		if b != 'U' && b != 'w' {
			return nil, errors.New("identifier is not an UwUID")
		}
		var bit byte = 0
		if b == 'U' {
			bit = 1
		}
		token[i/8] <<= 0x1
		token[i/8] |= bit
	}
	return FromBytes(token)
}
