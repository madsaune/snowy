package fieldtype

import (
	"bytes"
	"strconv"
)

type SNBool bool

func (s *SNBool) UnmarshalJSON(bs []byte) error {
	st := string(bytes.Trim(bs, "\""))

	b, err := strconv.ParseBool(st)
	if err != nil {
		return err
	}

	*s = SNBool(b)

	return nil
}
