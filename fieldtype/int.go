package fieldtype

import (
	"bytes"
	"strconv"
)

type SNInt int

func (s *SNInt) UnmarshalJSON(bs []byte) error {
	st := string(bytes.Trim(bs, "\""))

	i, err := strconv.Atoi(st)
	if err != nil {
		return err
	}

	*s = SNInt(i)

	return nil
}
