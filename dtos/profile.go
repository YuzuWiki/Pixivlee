package dtos

import (
	"encoding/json"
	"sort"
	"strconv"
)

// TopProfileDTO return user's  profile (top)
type TopProfileDTO struct {
	Illusts   illustMapDTO `json:"illusts"`
	Manga     mangaMapDTO  `json:"manga"`
	Novels    novelMapDTO  `json:"novels"`
	ExtraData extraDataDTO `json:"extra_data"`
}

type extraDataDTO struct {
	Meta struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Canonical   string `json:"canonical"`
		Ogp         struct {
			Description string `json:"description"`
			Image       string `json:"image"`
			Title       string `json:"title"`
			Type        string `json:"type"`
		} `json:"ogp"`
		Twitter struct {
			Description string `json:"description"`
			Image       string `json:"image"`
			Title       string `json:"title"`
			Card        string `json:"card"`
		} `json:"twitter"`
		DescriptionHeader string `json:"descriptionHeader"`
	} `json:"meta"`
}

// AllProfileDTO return user's  profile (all)
type AllProfileDTO struct {
	Illusts ArtWorkIdsDTO `json:"illusts"`
	Manga   ArtWorkIdsDTO `json:"manga"`
	Novel   ArtWorkIdsDTO `json:"novels"`
}

type ArtWorkIdsDTO []int64

func (dto *ArtWorkIdsDTO) UnmarshalJSON(body []byte) error {
	var (
		data map[string]struct{}
		ids  []int64
	)

	if err := json.Unmarshal(body, &data); err == nil {
		for idStr := range data {
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err != nil {
				return err
			}
			ids = append(ids, id)
		}
	}
	sort.Slice(ids, func(i, j int) bool {
		return ids[i] < ids[j]
	})
	*dto = ids
	return nil
}
