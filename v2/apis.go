package v2

import (
	"fmt"
	"github.com/imroc/req/v3"
	"net/http"
	"strconv"

	"github.com/YuzuWiki/Pixivlee/v2/jsonObject"
)

type ResponseData[T interface{}] struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Body    T      `json:"body"`
}

func (r *ResponseData[T]) Assert() error {
	if r.Error || len(r.Message) > 0 {
		return fmt.Errorf(r.Message)
	}
	return nil
}

// request API.request
func pixivRequest(p IPixiver) *req.Request {
	r := requestClient.NewRequest()

	// set pixiv cookie
	r.SetCookies(
		&http.Cookie{
			Name:   Phpsessid,
			Value:  p.SessionID(),
			Path:   "/",
			Domain: PixivDomain,
		})

	r.AddQueryParam("lang", "jp")

	r.OnAfterResponse(func(client *req.Client, resp *req.Response) error {
		if p.Pid() == 0 {
			if pid, err := strconv.ParseUint(resp.Header.Get("x-userid"), 10, 64); err == nil {
				p.SetPid(TPid(pid))
			}

		}

		switch resp.StatusCode {

		case 401:
			p.UpdateState(PixiverInvalid)

		case 404:
			// 404 Not Found

		case 409:
			p.UpdateState(PixiverRateLimiting)

		default:
			p.UpdateState(PixiverRateLimiting)
		}

		return nil
	})

	return r
}

type PixivApi struct {
}

func (P PixivApi) GetPid(p IPixiver) (TPid, error) {
	return p.Pid(), nil
}

func (P PixivApi) AccountInfo(p IPixiver) (data *ResponseData[jsonObject.User], err error) {
	r := pixivRequest(p)
	r.SetPathParam("full", "1").SetSuccessResult(&data)

	if _, err = r.Get(fmt.Sprintf("/ajax/user/%d", p.Pid())); err != nil {
		return nil, err
	}

	if err = data.Assert(); err != nil {
		return nil, err
	}
	return data, nil
}

func (P PixivApi) ProfileAll(p IPixiver, pid TPid) (data *ResponseData[jsonObject.ProfileAll], err error) {
	r := pixivRequest(p).SetSuccessResult(&data)

	if _, err = r.Get(fmt.Sprintf("/ajax/user/%d/profile/all", pid)); err != nil {
		return nil, err
	}

	if err = data.Assert(); err != nil {
		return nil, err
	}
	return data, nil
}

func (P PixivApi) ProfileTop(p IPixiver, pid TPid) (data *ResponseData[jsonObject.ProfileTop], err error) {
	r := pixivRequest(p).SetSuccessResult(&data)

	if _, err = r.Get(fmt.Sprintf("/ajax/user/%d/profile/top", pid)); err != nil {
		return nil, err
	}

	if err = data.Assert(); err != nil {
		return nil, err
	}
	return data, err
}

func (P PixivApi) FanBoxUrl(p IPixiver, pid TPid) (_ string, err error) {
	r := pixivRequest(p)

	resp, err := r.Get(fmt.Sprintf("/fanbox/creator/%d", pid))
	if err != nil {
		return "", err
	}

	return resp.Response.Request.URL.String(), nil
}

func (P PixivApi) artWork(p IPixiver, artType string, artIId TArtId) (data *ResponseData[jsonObject.ArtWork], err error) {
	r := pixivRequest(p).SetSuccessResult(&data)

	if _, err = r.Get(fmt.Sprintf("/ajax/%s/%d", artType, artIId)); err != nil {
		return nil, err
	}

	if err = data.Assert(); err != nil {
		return nil, err
	}
	return data, nil
}

func (P PixivApi) ArtWorkIllust(p IPixiver, artId TArtId) (*ResponseData[jsonObject.ArtWork], error) {
	return P.artWork(p, ArtTypeIllust, artId)
}

func (P PixivApi) ArtWorkManga(p IPixiver, artId TArtId) (*ResponseData[jsonObject.ArtWork], error) {
	return P.artWork(p, ArtTypeManga, artId)
}

func (P PixivApi) ArtWorkNovel(p IPixiver, artId TArtId) (*ResponseData[jsonObject.ArtWork], error) {
	return P.artWork(p, ArtTypeNovel, artId)
}

func (P PixivApi) bookmarks(p IPixiver, rest string, pid TPid, tag string, offset, limit int) (data *ResponseData[jsonObject.Bookmark], err error) {
	r := pixivRequest(p).
		AddQueryParam("tag", tag).
		AddQueryParam("limit", fmt.Sprint(limit)).
		AddQueryParams("offset", fmt.Sprint(offset)).
		AddQueryParams("rest", rest).
		SetSuccessResult(&data)

	if _, err = r.Get(fmt.Sprintf("/ajax/user/%d/illusts/bookmarks", pid)); err != nil {
		return nil, err
	}

	if err = data.Assert(); err != nil {
		return nil, err
	}
	return data, nil
}

func (P PixivApi) BookmarkShow(p IPixiver, pid TPid, tag string, offset, limit int) (*ResponseData[jsonObject.Bookmark], error) {
	return P.bookmarks(p, RestTypeShow, pid, tag, offset, limit)
}

func (P PixivApi) BookmarkHide(p IPixiver, pid TPid, tag string, offset, limit int) (*ResponseData[jsonObject.Bookmark], error) {
	return P.bookmarks(p, RestTypeHide, pid, tag, offset, limit)
}

// followedLast (Last By Followed)
// page > 0 && data from https://www.pixiv.net/bookmark_new_illust.php
func (P PixivApi) followedLast(p IPixiver, mode string, page int) (data *ResponseData[jsonObject.FollowLatestDTO], err error) {
	r := pixivRequest(p).
		AddQueryParam("p", fmt.Sprint(page)).
		AddQueryParam("mode", mode).
		SetSuccessResult(&data)

	if _, err = r.Get(fmt.Sprintf("/ajax/follow_latest/%s", mode)); err != nil {
		return nil, err
	}

	if err = data.Assert(); err != nil {
		return nil, err
	}
	return data, nil
}

func (P PixivApi) FollowedLastIllust(p IPixiver, page int) (*ResponseData[jsonObject.FollowLatestDTO], error) {
	return P.followedLast(p, ArtTypeIllust, page)
}

func (P PixivApi) FollowedLastManga(p IPixiver, page int) (*ResponseData[jsonObject.FollowLatestDTO], error) {
	return P.followedLast(p, ArtTypeManga, page)
}

func (P PixivApi) FollowedLastNovel(p IPixiver, page int) (*ResponseData[jsonObject.FollowLatestDTO], error) {
	return P.followedLast(p, ArtTypeNovel, page)
}

func (P PixivApi) FollowingList(p IPixiver, pid TPid, limit int, offset int) (data *ResponseData[jsonObject.Following], err error) {
	r := pixivRequest(p).SetPathParams(map[string]string{
		"offset": fmt.Sprint(offset),
		"limit":  fmt.Sprint(limit),
		"tag":    "",
		"rest":   RestTypeShow,
	}).SetSuccessResult(&data)

	if _, err = r.Get(fmt.Sprintf("/ajax/use/%d/following", pid)); err != nil {
		return nil, err
	}

	if err = data.Assert(); err != nil {
		return nil, err
	}
	return data, nil
}

func (P PixivApi) TagSearch(p IPixiver, jpName string) (data *ResponseData[jsonObject.Tag], err error) {
	r := pixivRequest(p).SetSuccessResult(&data)

	if _, err = r.Get(fmt.Sprintf("/ajax/search/tags/%s", jpName)); err != nil {
		return nil, err
	}

	if err = data.Assert(); err != nil {
		return nil, err
	}
	return data, nil
}

func (P PixivApi) rank(p IPixiver, mode string, content string, page int, date string) (data *jsonObject.RankJson, err error) {
	r := pixivRequest(p).
		SetSuccessResult(&data).
		AddQueryParam("page", fmt.Sprint(page)).
		AddQueryParam("format", "json")
	if len(mode) > 0 {
		r.AddQueryParam("mode", mode)
	}
	if len(content) > 0 {
		r.AddQueryParam("content", content)
	}
	if len(date) > 0 {
		r.AddQueryParam("date", date)
	}

	if _, err = r.Get(fmt.Sprintf("/ranking.php")); err != nil {
		return nil, err
	}
	return data, nil
}

func (P PixivApi) RankALL(p IPixiver, mode string, page int, date string) (*jsonObject.RankJson, error) {
	switch mode {
	case RankDaily, RankDailyR18:
	case RankWeekly, RankWeeklyR18:
	case RankMonthly:
	case RankRookie:
	case RankOriginal:
	case RankAI, RankAIR18:
	case RankMale, RankMaleR18:
	case RankFemale, RankFemaleR18:
	default:
		return nil, fmt.Errorf("invalid mode value")
	}
	return P.rank(p, mode, "", page, date)
}

func (P PixivApi) RankIllust(p IPixiver, mode string, page int, date string) (*jsonObject.RankJson, error) {
	switch mode {
	case RankDaily, RankDailyR18:
	case RankWeekly, RankWeeklyR18:
	case RankMonthly:
	case RankRookie:
	default:
		return nil, fmt.Errorf("invalid mode value")
	}
	return P.rank(p, mode, ArtTypeIllust, page, date)
}

func (P PixivApi) RankUgoira(p IPixiver, mode string, page int, date string) (*jsonObject.RankJson, error) {
	switch mode {
	case RankDaily, RankDailyR18:
	case RankWeekly, RankWeeklyR18:
	default:
		return nil, fmt.Errorf("invalid mode value")
	}
	return P.rank(p, mode, ArtTypeUgoira, page, date)
}

func (P PixivApi) RankManga(p IPixiver, mode string, page int, date string) (*jsonObject.RankJson, error) {
	switch mode {
	case RankDaily, RankDailyR18:
	case RankWeekly, RankWeeklyR18:
	case RankMonthly:
	case RankRookie:
	default:
		return nil, fmt.Errorf("invalid mode value")
	}
	return P.rank(p, mode, ArtTypeManga, page, date)
}
