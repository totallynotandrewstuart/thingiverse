package thingiverse

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"golang.org/x/oauth2"
)

type Client struct {
	c *http.Client
}

func NewClient(username, password string) (*Client, error) {
	jar, _ := cookiejar.New(nil)
	hc := http.Client{
		Jar: jar,
	}

	v := url.Values{"username": {username}, "password": {password}}
	// res, err := hc.Post("https://login.makerbot.com/login", "application/json", bytes.NewReader(bs))
	res, err := hc.Post("https://accounts.thingiverse.com/login", "application/x-www-form-urlencoded", strings.NewReader(v.Encode()))
	if err != nil {
		return nil, err
	}
	res.Body.Close()
	fmt.Printf("res = %+v\n", res)

	res, err = hc.Get("https://www.thingiverse.com/ajax/user/exchange_session_for_token")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var tokErr struct{ Token, Error string }
	err = json.NewDecoder(res.Body).Decode(&tokErr)
	if err != nil {
		return nil, err
	}

	if tokErr.Error != "" {
		return nil, fmt.Errorf("%s", tokErr.Error)
	}

	oCli, err := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(&oauth2.Token{
		TokenType:   "Bearer",
		AccessToken: tokErr.Token,
	})), nil
	if err != nil {
		return nil, err
	}

	return &Client{
		c: oCli,
	}, nil
}

type PageOpt struct {
	Page    int
	PerPage int
	Sort    string
}

func (po PageOpt) Query() string {
	return url.Values{
		"page":     {fmt.Sprint(po.Page)},
		"per_page": {fmt.Sprint(po.PerPage)},
		"sort":     {po.Sort},
	}.Encode()
}

type SearchOpts struct {
	PageOpt
	PostedAfter string
	Type        string
}

func (so SearchOpts) Query() string {
	return so.PageOpt.Query() + url.Values{
		"type":         {so.Type},
		"posted_after": {so.PostedAfter},
	}.Encode()
}

func (c Client) Search(ctx context.Context, opts *SearchOpts) ([]Thing, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://api.thingiverse.com/search/?page=1&per_page=20&sort=popular&posted_after=now-30d&type=things", nil)
	if err != nil {
		return nil, err
	}

	res, err := c.c.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var ths ThingRes
	err = json.NewDecoder(res.Body).Decode(&ths)
	if err != nil {
		return nil, err
	}

	return ths.Hits, nil
}
