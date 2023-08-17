package xip

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/gogf/gf/v2/text/gstr"
)

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
