package types

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type (
	TPid     uint64 // pid type alise
	TArtId   uint64 // artId type alise
	TArtType uint8  // illust 0; manga 1; ugoira 2
)

type TCount uint32

func (o *TCount) UnmarshalJSON(body []byte) error {
	cnt, err := strconv.ParseUint(strings.ReplaceAll(string(body), "\"", ""), 10, 32)
	if err != nil {
		return err
	}

	*o = TCount(cnt)
	return nil
}

type TTimestamp time.Time

func (t *TTimestamp) UnmarshalJSON(body []byte) error {
	// todo: 兼容时间格式字符串解析
	timestamp, err := strconv.ParseInt(string(body), 10, 64)
	if err != nil {
		return err
	}

	*t = TTimestamp(time.Unix(timestamp, 0))
	return nil
}

type TJsonDict[TK string, TV any] map[TK]TV

func (o *TJsonDict[TK, TV]) UnmarshalJSON(body []byte) error {
	if len(body) < 4 {
		return nil
	}

	Dict := make(map[TK]TV)
	if err := json.Unmarshal(body, &Dict); err != nil {
		return err
	}

	*o = Dict
	return nil
}
