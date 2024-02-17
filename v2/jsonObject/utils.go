package jsonObject

import (
	"encoding/json"
	"sort"
	"strconv"
)

type jsonMap[k string, v any] map[k]v

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

type artWorkIds []TArtId

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
