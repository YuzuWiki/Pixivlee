package types

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type (
	TArtId   uint64 // artId type alise
	TArtType uint8  // illust 0; manga 1; ugoira 2
)

type TPid uint64 // pid type alise

func (o *TPid) UnmarshalJSON(body []byte) error {
	cnt, err := strconv.ParseUint(strings.ReplaceAll(string(body), "\"", ""), 10, 32)
	if err != nil {
		return err
	}

	*o = TPid(cnt)
	return nil
}

type TPostId uint64 // pid type alise

func (o *TPostId) UnmarshalJSON(body []byte) error {
	cnt, err := strconv.ParseUint(strings.ReplaceAll(string(body), "\"", ""), 10, 32)
	if err != nil {
		return err
	}

	*o = TPostId(cnt)
	return nil
}

type TCount uint32

func (o *TCount) UnmarshalJSON(body []byte) error {
	cnt, err := strconv.ParseUint(strings.ReplaceAll(string(body), "\"", ""), 10, 32)
	if err != nil {
		return err
	}

	*o = TCount(cnt)
	return nil
}

type TTimestamp struct {
	time.Time
}

func (t *TTimestamp) UnmarshalJSON(body []byte) error {
	timestamp, err := strconv.ParseInt(string(body), 10, 64)
	if err != nil {
		return err
	}

	*t = TTimestamp{time.Unix(timestamp, 0)}
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

type TTimeDate struct {
	time.Time
}

func (t *TTimeDate) UnmarshalJSON(body []byte) error {
	td, err := time.ParseInLocation(time.RFC3339, strings.ReplaceAll(string(body), "\"", ""), nil)
	if err != nil {
		return err
	}
	*t = TTimeDate{td}
	return nil
}

type TArtIds []TArtId

func (o *TArtIds) UnmarshalJSON(body []byte) error {
	var artDict map[string]any

	err := json.Unmarshal(body, &artDict)
	if err != nil {
		return err
	}

	artIds := make(TArtIds, 0)
	for artIdStr, _ := range artDict {
		artId, err := strconv.ParseInt(artIdStr, 10, 64)
		if err != nil {
			return err
		}

		artIds = append(artIds, TArtId(artId))
	}

	*o = artIds
	return nil
}
