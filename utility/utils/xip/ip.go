package xip

import (
	"context"
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type Location struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	Anycast  bool   `json:"anycast"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
}

func GetIpInfo(ip string) (string, error) {
	c := colly.NewCollector()
	var location string
	c.OnHTML("body > div.wrapper > div.container > div:nth-child(1) > div > div.ft > table > tbody > tr > td:nth-child(2) > span", func(e *colly.HTMLElement) {
		location = e.Text
	})
	err := c.Visit(fmt.Sprintf("https://www.ipshudi.com/%s.htm", ip))
	if err != nil {
		return "", err
	}
	c.Wait()
	return gstr.Trim(location), nil
}

func GetIpInfoIo(ctx context.Context, ip string) Location {
	loc := Location{}
	get, err := g.Client().Get(ctx, fmt.Sprintf("https://ipinfo.io/%s?token=752e995a5aa063", ip))
	if err != nil {
		g.Log().Error(ctx, err)
		return loc
	}
	allString := get.ReadAllString()
	gconv.Struct(allString, &loc)

	defer get.Close()

	return loc
}
