package jsonObject

import (
	"strconv"
	"strings"
)

type (
	TPid        uint64 // pid type alise
	TArtId      uint64 // artId type alise
	TIllustType uint8  // illust 0; manga 1; ugoira 2

	TCount uint32
)

func (t *TPid) UnmarshalJSON(body []byte) error {
	data, err := strconv.ParseUint(strings.ReplaceAll(string(body), "\"", ""), 10, 64)
	if err != nil {
		return err
	}

	*t = TPid(data)
	return nil
}

func (t *TArtId) UnmarshalJSON(body []byte) error {
	data, err := strconv.ParseUint(strings.ReplaceAll(string(body), "\"", ""), 10, 64)
	if err != nil {
		return err
	}

	*t = TArtId(data)
	return nil
}
func (t *TIllustType) UnmarshalJSON(body []byte) error {
	if _t, err := strconv.ParseUint(strings.ReplaceAll(string(body), "\"", ""), 10, 8); err != nil {
		return err
	} else {
		*t = TIllustType(_t)
	}
	return nil
}

func (t *TCount) UnmarshalJSON(body []byte) error {
	if _t, err := strconv.ParseUint(strings.ReplaceAll(string(body), "\"", ""), 10, 32); err != nil {
		return err
	} else {
		*t = TCount(_t)
	}
	return nil
}
