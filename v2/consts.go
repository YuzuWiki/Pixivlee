package v2

const (
	FanboxHost = "www.fanbox.cc"

	PixivDomain = ".pixiv.net"
	PixivHost   = "www.pixiv.net"
	PximgHost   = "i.pximg.net"
)

const (
	Phpsessid = "PHPSESSID"
	UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:84.0) Gecko/20100101 Firefox/84.0"
)

const (
	PixiverUnknown = iota
	PixiverNormal
	PixiverRateLimiting
	PixiverInvalid
)

const (
	ArtTypeIllust = "illust"
	ArtTypeManga  = "manga"
	ArtTypeNovel  = "novel"
	ArtTypeUgoira = "ugoira"

	RestTypeShow = "show"
	RestTypeHide = "hide"

	/*
		See https://www.pixiv.net/ranking.php
	*/
	RankDaily     = "daily" // 日
	RankDailyR18  = "daily_r18"
	RankWeekly    = "weekly" // 周
	RankWeeklyR18 = "weekly_r18"
	RankMonthly   = "monthly"  // 月
	RankRookie    = "rookie"   // 新人
	RankOriginal  = "original" // 原创
	RankAI        = "daily_ai" // AI
	RankAIR18     = "daily_r18_ai"
	RankMale      = "male" //最受到男性欢迎
	RankMaleR18   = "male_r18"
	RankFemale    = "female" // 最收到女性欢迎
	RankFemaleR18 = "female_r18"
)
