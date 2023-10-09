package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

var rssURL string

func main() {
	flag.StringVar(&rssURL, "rssurl", "", "rss URL to load")
	flag.Parse()

	officialRSSURL := "https://www.omnycontent.com/d/playlist/77bedd50-a734-42aa-9c08-ad86013ca0f9/4dbfc420-53a4-40c6-bbc7-ad8d012bc602/6ede3615-a245-4eae-9087-ad8d012bc631/podcast.rss"

	patreonRSS, err := http.Get(rssURL)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		err := patreonRSS.Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	patreonData, err := io.ReadAll(patreonRSS.Body)
	if err != nil {
		fmt.Println(err)
	}

	officialRSS, err := http.Get(officialRSSURL)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		err := officialRSS.Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	// use this to crawl for better images
	// as patreon doesn't give them
	_, err = io.ReadAll(officialRSS.Body)
	if err != nil {
		fmt.Println(err)
	}

	var rssResp RSSMain
	err = xml.Unmarshal(patreonData, &rssResp)
	if err != nil {
		fmt.Println(err)
	}
	categoriesd := rssResp.Categorise()
	for k, v := range categoriesd {
		for i := range v {
			fmt.Println(k, v[i].Image.Href, v[i].Title, v[i].Link)
		}
	}
}

var (
	// campaigns
	campaign1         = "c1"
	campaign2         = "c2"
	eldermourne       = "eldermourne"
	campaign3         = "c3"
	hotBoySummer      = "mavrus chronicles"
	trinyvaleCategory = "trinyvale"

	noShortRest   = " no short rest"
	shortRestOnly = " short rest only"

	// meta categories
	campaign1NoShortRest = campaign1 + noShortRest
	campaign1ShortRest   = campaign1 + shortRestOnly

	campaign2NoShortRest = campaign2 + noShortRest
	campaign2ShortRest   = campaign2 + shortRestOnly

	campaign3NoShortRest = campaign3 + noShortRest
	campaign3ShortRest   = campaign3 + shortRestOnly

	mavrusNoShortRest = hotBoySummer + noShortRest
	mavrusShortRest   = hotBoySummer + shortRestOnly

	sessionZero = "session zero"

	// guests
	guest    = "guest"
	guest2   = "w/"
	zacGuest = "zac oyama"
	// zacGuest2    = "mavrus"
	louGuest     = "lou wilson"
	jasperGuest  = "jasper"
	amirGuest    = "amir"
	nathanGuest  = "nathan"
	ifyGuest     = "ify nwadiwe"
	aabriaGuest  = "aabria"
	allyGuest    = "ally"
	jeremyGuest  = "jeremy"
	juliaGuest   = "julia"
	brennanGuest = "brennan"

	// short rest tier
	shortRestCategory = "short rest"

	// mixed bag tier
	mixedBagCategory = "mixed bag"
	hotBoy2          = "blazing babe"
	tortle2          = "owlbear"

	// non-campaign categories
	animorphs             = "animorphs"
	flinstones            = "flinstones"
	musicCategory         = "music"
	donkeyKongCategory    = "donkey kong"
	bookclubCategory      = "book"
	dungeonCourtCategory  = "dungeon court"
	dungeonCourt2Category = "d&d court"
	liveCategory          = "live"
	hearthSideChat        = "hearthside"
	baggingIt             = "baggin"
	tortle                = "tortle"
	oneShot               = "one-shot"
	oneShot2              = "one shot"
	twoShot               = "two-shot"
	twoShot2              = "two shot"
	hexBuds               = "hexblood"
	bonus                 = "bonus"
	dmAdvice              = "behind the screens"
)

func (r *RSSMain) Categorise() map[string][]RSSItem {
	resp := make(map[string][]RSSItem)
	for i := range r.Channel.Item {
		var categoryString string
		if strings.Contains(strings.ToLower(r.Channel.Item[i].Title), shortRestCategory) {
			categoryString = shortRestCategory
			resp[categoryString] = append(resp[categoryString], r.Channel.Item[i])
		}
		if strings.Contains(strings.ToLower(r.Channel.Item[i].Title), trinyvaleCategory) {
			categoryString = trinyvaleCategory
			resp[categoryString] = append(resp[categoryString], r.Channel.Item[i])
		}
		if strings.Contains(strings.ToLower(r.Channel.Item[i].Title), mixedBagCategory) {
			categoryString = mixedBagCategory
			resp[categoryString] = append(resp[categoryString], r.Channel.Item[i])
		}
		if strings.Contains(strings.ToLower(r.Channel.Item[i].Title), musicCategory) {
			categoryString = musicCategory
			resp[categoryString] = append(resp[categoryString], r.Channel.Item[i])
		}
		if strings.Contains(strings.ToLower(r.Channel.Item[i].Title), donkeyKongCategory) {
			categoryString = donkeyKongCategory
			resp[categoryString] = append(resp[categoryString], r.Channel.Item[i])
		}
		if strings.Contains(strings.ToLower(r.Channel.Item[i].Title), bookclubCategory) {
			categoryString = bookclubCategory
			resp[categoryString] = append(resp[categoryString], r.Channel.Item[i])
		}
		if strings.Contains(strings.ToLower(r.Channel.Item[i].Title), dungeonCourtCategory) ||
			strings.Contains(strings.ToLower(r.Channel.Item[i].Title), dungeonCourt2Category) {
			categoryString = dungeonCourtCategory
			resp[categoryString] = append(resp[categoryString], r.Channel.Item[i])
		}
		if strings.Contains(strings.ToLower(r.Channel.Item[i].Title), liveCategory) {
			categoryString = liveCategory
			resp[categoryString] = append(resp[categoryString], r.Channel.Item[i])
		}
		if strings.Contains(strings.ToLower(r.Channel.Item[i].Title), eldermourne) ||
			strings.Contains(strings.ToLower(r.Channel.Item[i].Title), campaign2) {
			categoryString = campaign2
			resp[categoryString] = append(resp[categoryString], r.Channel.Item[i])
			if !strings.Contains(strings.ToLower(r.Channel.Item[i].Title), shortRestCategory) {
				categoryString = campaign2NoShortRest
				resp[categoryString] = append(resp[categoryString], r.Channel.Item[i])
			}
		}
		if strings.Contains(strings.ToLower(r.Channel.Item[i].Title), campaign3) {
			categoryString = campaign3
			resp[categoryString] = append(resp[categoryString], r.Channel.Item[i])
			if !strings.Contains(strings.ToLower(r.Channel.Item[i].Title), shortRestCategory) {
				categoryString = campaign3NoShortRest
				resp[categoryString] = append(resp[categoryString], r.Channel.Item[i])
			}
		}
		if strings.Contains(strings.ToLower(r.Channel.Item[i].Title), hearthSideChat) {
			categoryString = hearthSideChat
			resp[categoryString] = append(resp[categoryString], r.Channel.Item[i])
		}
		if strings.Contains(strings.ToLower(r.Channel.Item[i].Title), baggingIt) {
			categoryString = baggingIt
			resp[categoryString] = append(resp[categoryString], r.Channel.Item[i])
		}
		if strings.Contains(strings.ToLower(r.Channel.Item[i].Title), hotBoySummer) {
			categoryString = hotBoySummer
			resp[categoryString] = append(resp[categoryString], r.Channel.Item[i])
		}
		if strings.Contains(strings.ToLower(r.Channel.Item[i].Title), hotBoy2) {
			categoryString = hotBoy2
			resp[categoryString] = append(resp[categoryString], r.Channel.Item[i])
		}
		if strings.Contains(strings.ToLower(r.Channel.Item[i].Title), guest) {
			categoryString = guest
			resp[categoryString] = append(resp[categoryString], r.Channel.Item[i])
		}
		if strings.Contains(strings.ToLower(r.Channel.Item[i].Title), hexBuds) {
			categoryString = hexBuds
			resp[categoryString] = append(resp[categoryString], r.Channel.Item[i])
		}
		if strings.Contains(strings.ToLower(r.Channel.Item[i].Title), dmAdvice) {
			categoryString = dmAdvice
			resp[categoryString] = append(resp[categoryString], r.Channel.Item[i])
		}
		if strings.Contains(strings.ToLower(r.Channel.Item[i].Title), bonus) {
			categoryString = bonus
			resp[categoryString] = append(resp[categoryString], r.Channel.Item[i])
		}
		if strings.Contains(strings.ToLower(r.Channel.Item[i].Title), twoShot) {
			categoryString = twoShot
			resp[categoryString] = append(resp[categoryString], r.Channel.Item[i])
		}
		if strings.Contains(strings.ToLower(r.Channel.Item[i].Title), tortle) ||
			strings.Contains(strings.ToLower(r.Channel.Item[i].Title), tortle2) {
			categoryString = tortle
			resp[categoryString] = append(resp[categoryString], r.Channel.Item[i])
		}
		tt, err := time.Parse(time.RFC1123, r.Channel.Item[i].PubDate)
		if err != nil {
			panic(err)
		}
		if tt.Before(time.Date(2020, 5, 20, 0, 0, 0, 0, time.UTC)) &&
			!strings.Contains(strings.ToLower(r.Channel.Item[i].Title), trinyvaleCategory) &&
			!strings.Contains(strings.ToLower(r.Channel.Item[i].Title), liveCategory) &&
			!strings.Contains(strings.ToLower(r.Channel.Item[i].Title), hearthSideChat) &&
			!strings.Contains(strings.ToLower(r.Channel.Item[i].Title), mixedBagCategory) &&
			!strings.Contains(strings.ToLower(r.Channel.Item[i].Title), donkeyKongCategory) &&
			!strings.Contains(strings.ToLower(r.Channel.Item[i].Title), baggingIt) &&
			!strings.Contains(strings.ToLower(r.Channel.Item[i].Title), musicCategory) {
			categoryString = campaign1
			resp[categoryString] = append(resp[categoryString], r.Channel.Item[i])
			if !strings.Contains(strings.ToLower(r.Channel.Item[i].Title), shortRestCategory) {
				categoryString = campaign1NoShortRest
				resp[categoryString] = append(resp[categoryString], r.Channel.Item[i])
			}
		}

		if categoryString == "" {
			resp["UKNOWN"] = append(resp["UKNOWN"], r.Channel.Item[i])
		}
	}

	return resp
}

type RSSLink struct {
	Text string `xml:",chardata"`
	Href string `xml:"href,attr"`
	Rel  string `xml:"rel,attr"`
	Type string `xml:"type,attr"`
}

type RSSOwner struct {
	Text  string `xml:",chardata"`
	Name  string `xml:"name"`
	Email string `xml:"email"`
}

type RSSImage struct {
	Text  string `xml:",chardata"`
	Href  string `xml:"href,attr"`
	URL   string `xml:"url"`
	Title string `xml:"title"`
	Link  string `xml:"link"`
}

type RSSEnclosure struct {
	Text   string `xml:",chardata"`
	URL    string `xml:"url,attr"`
	Length string `xml:"length,attr"`
	Type   string `xml:"type,attr"`
}

type RSSGuid struct {
	Text        string `xml:",chardata"`
	IsPermaLink string `xml:"isPermaLink,attr"`
}

type RSSItem struct {
	Text         string       `xml:",chardata"`
	Title        string       `xml:"title"`
	Link         string       `xml:"link"`
	Image        RSSImage     `xml:"image"`
	Description  string       `xml:"description"`
	Enclosure    RSSEnclosure `xml:"enclosure"`
	Guid         RSSGuid      `xml:"guid"`
	PubDate      string       `xml:"pubDate"`
	CategoryTags []string     `xml:"-"`
}

type RSSChannel struct {
	Text          string    `xml:",chardata"`
	Title         string    `xml:"title"`
	Link          RSSLink   `xml:"link"`
	Description   string    `xml:"description"`
	Owner         RSSOwner  `xml:"owner"`
	Author        string    `xml:"author"`
	Image         RSSImage  `xml:"image"`
	Block         string    `xml:"block"`
	Language      string    `xml:"language"`
	PubDate       string    `xml:"pubDate"`
	LastBuildDate string    `xml:"lastBuildDate"`
	Item          []RSSItem `xml:"item"`
}

type RSSMain struct {
	XMLName    xml.Name   `xml:"rss"`
	Text       string     `xml:",chardata"`
	Version    string     `xml:"version,attr"`
	Itunes     string     `xml:"itunes,attr"`
	Atom       string     `xml:"atom,attr"`
	Googleplay string     `xml:"googleplay,attr"`
	Channel    RSSChannel `xml:"channel"`
}
