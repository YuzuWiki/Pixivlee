package Pixivlee

import (
	"encoding/json"
	"sort"
	"strconv"
	"strings"
	"time"
)

// IPixiver pixiver
type IPixiver interface {
	Pid() TPid
	SessionID() string
	State() int

	SetPid(TPid)
	UpdateState(state int)
}

type (
	TPid        uint64 // pid type alise
	TArtId      uint64 // artId type alise
	TIllustType uint8  // illust 0; manga 1; ugoira 2

	TCount     uint32
	TTimestamp time.Time

	artWorkIds               []TArtId
	jsonMap[k string, v any] map[k]v
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

func (t *TTimestamp) UnmarshalJSON(body []byte) error {
	if _t, err := strconv.ParseInt(string(body), 10, 64); err != nil {
		return err
	} else {
		*t = TTimestamp(time.Unix(_t, 0))
	}
	return nil
}

func (s *artWorkIds) UnmarshalJSON(body []byte) error {
	var (
		data map[string]struct{}
		ids  artWorkIds
	)

	if err := json.Unmarshal(body, &data); err == nil {
		for idStr := range data {
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err != nil {
				return err
			}
			ids = append(ids, TArtId(id))
		}
	}
	sort.Slice(ids, func(i, j int) bool {
		return ids[i] < ids[j]
	})
	*s = ids
	return nil
}

func (m *jsonMap[k, v]) UnmarshalJSON(body []byte) error {
	data := map[k]v{}
	if len(body) > 5 {
		if err := json.Unmarshal(body, &data); err != nil {
			return err
		}
	}
	*m = data
	return nil
}
