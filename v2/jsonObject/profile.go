package jsonObject

type extra struct {
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

// TopProfileDTO return user's  profile (top)
type TopProfileDTO struct {
	Illusts   jsonMap[string, illustItem] `json:"illusts"`
	Manga     jsonMap[string, mangaItem]  `json:"manga"`
	Novels    jsonMap[string, novelItem]  `json:"novels"`
	ExtraData extra                       `json:"extra_data"`
}

// AllProfile return user's  profile (all)
type AllProfile struct {
	Illusts artWorkIds `json:"illusts"`
	Manga   artWorkIds `json:"mangas"`
	Novel   artWorkIds `json:"novels"`
}
