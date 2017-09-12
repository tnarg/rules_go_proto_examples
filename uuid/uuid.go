package uuid

import (
	"errors"

	google_uuid "github.com/google/uuid"
)

type UUID [16]byte

func New() UUID {
	return UUID(google_uuid.New())
}

func (u UUID) Size() int {
	return len(u)
}

func (u UUID) Marshal() ([]byte, error) {
	return u[:], nil
}

func (u *UUID) MarshalTo(data []byte) (n int, err error) {
	return copy(data, u[:]), nil
}

func (u *UUID) Unmarshal(data []byte) error {
	n := copy(u[:], data)
	if n != 16 {
		return errors.New("UUID expects 16 bytes")
	}
	return nil
}
