package wallpaper

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/proxy"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

// colly 中文文档
//https://blog.csdn.net/qq_27818541/article/details/111492773

const url = "https://wallhaven.cc/latest?page=%d"

func CrawWallPaper() {
	today := time.Now()
	outputDir := fmt.Sprintf("download/wallpaper_%d_%d_%d/", today.Year(), today.Month(), today.Day())
	if _, oe := os.Stat(outputDir); os.IsNotExist(oe) {
		os.MkdirAll(outputDir, 0777)
	}
	c := colly.NewCollector(
		// Restrict crawling to specific domains
		colly.AllowedDomains("w.wallhaven.cc", "wallhaven.cc"),
		//colly.Async(),
		//colly.CacheDir("./_instagram_cache/"),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.84 Safari/537.36"),
	)

	c.WithTransport(&http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   300 * time.Second,
			KeepAlive: 300 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	})

	// cf_clearance=9N8zB2Zfxx.1nQ4qgVyTkyNfu6ytvkYPY9a1O9XFr4k-1647774662-0-250
	c.SetCookies("w.wallhaven.cc", []*http.Cookie{
		&http.Cookie{
			Name: "cf_clearance", Value: "9N8zB2Zfxx.1nQ4qgVyTkyNfu6ytvkYPY9a1O9XFr4k-1647774662-0-250",
		},
	})

	// clash for windows 翻墙代理
	// 防爬注意：SSH服务器可以用作带有-D标志的socks5代理。
	rp, err := proxy.RoundRobinProxySwitcher("socks5://127.0.0.1:7890")
	if err != nil {
		log.Fatal(err)
	}
	_ = rp
	//c.SetProxyFunc(rp)
	//c.SetProxy("http://127.0.0.1:7890")

	c.Limit(&colly.LimitRule{
		DomainRegexp: `wallhaven\.cc`,
		Parallelism:  1,
		Delay:        5 * time.Second,
		//RandomDelay: 1 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("X-Requested-With", "XMLHttpRequest")
		r.Headers.Set("Referer", url)
		r.Headers.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		r.Headers.Set("cookie", "cf_clearance=9N8zB2Zfxx.1nQ4qgVyTkyNfu6ytvkYPY9a1O9XFr4k-1647774662-0-250")
		r.Headers.Set("if-modified-since", "Fri, 01 Apr 2022 13:56:34 GMT")
		r.Headers.Set("if-none-match", "\"62470492-f06a5\"")
		r.Headers.Set("upgrade-insecure-requests", "1")
		//r.Headers.Set("accept-encoding", "gzip, deflate, br")
		//r.Headers.Set("accept-language", "zh-CN,zh;q=0.9")
		r.Headers.Set("cache-control","max-age=0")
		r.Headers.Set("sec-ch-ua","\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"99\", \"Google Chrome\";v=\"99\"")
		r.Headers.Set("sec-ch-ua-mobile","?0")
		r.Headers.Set("sec-ch-ua-platform","\"Windows\"")
		r.Headers.Set("sec-fetch-dest","document")
		r.Headers.Set("sec-fetch-mode","navigate")
		r.Headers.Set("sec-fetch-site","none")
		r.Headers.Set("sec-fetch-user","?1")
		//r.Headers.Set("","")
	})

	// <section class="thumb-listing-page"><ul><li><figure class="thumb thumb-0qr5m5 thumb-sfw thumb-general thumb-seen" data-wallpaper-id="0qr5m5" style="width:300px;height:200px"><a class="thumb-btn thumb-btn-fav jsAnchor overlay-anchor" href="https://wallhaven.cc/favorites/fav/0qr5m5" original-title="Add to favorites"><i class="fas fa-fw fa-star"></i></a><img alt="loading" class="lazyload loaded" data-src="https://th.wallhaven.cc/small/0q/0qr5m5.jpg" src="https://th.wallhaven.cc/small/0q/0qr5m5.jpg"><a class="preview" href="https://wallhaven.cc/w/0qr5m5" target="_blank"></a><div class="thumb-info"><span class="wall-res">1920 x 1080</span><a class="jsAnchor overlay-anchor wall-favs" data-href="https://wallhaven.cc/wallpaper/fav/0qr5m5">10<i class="fa fa-fw fa-star"></i></a><a class="jsAnchor thumb-tags-toggle tagged" data-href="https://wallhaven.cc/wallpaper/tags/0qr5m5" original-title="Tags"><i class="fas fa-fw fa-tags"></i></a></div></figure></li><li><figure class="thumb thumb-rddrmj thumb-sketchy thumb-people" data-wallpaper-id="rddrmj" style="width:300px;height:200px"><a class="thumb-btn thumb-btn-fav jsAnchor overlay-anchor" href="https://wallhaven.cc/favorites/fav/rddrmj" original-title="Add to favorites"><i class="fas fa-fw fa-star"></i></a><img alt="loading" class="lazyload loaded" data-src="https://th.wallhaven.cc/small/rd/rddrmj.jpg" src="https://th.wallhaven.cc/small/rd/rddrmj.jpg"><a class="preview" href="https://wallhaven.cc/w/rddrmj" target="_blank"></a><div class="thumb-info"><span class="wall-res">2560 x 1707</span><a class="jsAnchor overlay-anchor wall-favs" data-href="https://wallhaven.cc/wallpaper/fav/rddrmj">22<i class="fa fa-fw fa-star"></i></a><a class="jsAnchor thumb-tags-toggle tagged" data-href="https://wallhaven.cc/wallpaper/tags/rddrmj" original-title="Tags"><i class="fas fa-fw fa-tags"></i></a></div></figure></li><li><figure class="thumb thumb-6qgqq6 thumb-sketchy thumb-people" data-wallpaper-id="6qgqq6" style="width:300px;height:200px"><a class="thumb-btn thumb-btn-fav jsAnchor overlay-anchor" href="https://wallhaven.cc/favorites/fav/6qgqq6" original-title="Add to favorites"><i class="fas fa-fw fa-star"></i></a><img alt="loading" class="lazyload loaded" data-src="https://th.wallhaven.cc/small/6q/6qgqq6.jpg" src="https://th.wallhaven.cc/small/6q/6qgqq6.jpg"><a class="preview" href="https://wallhaven.cc/w/6qgqq6" target="_blank"></a><div class="thumb-info"><span class="wall-res">2048 x 1408</span><a class="jsAnchor overlay-anchor wall-favs" data-href="https://wallhaven.cc/wallpaper/fav/6qgqq6">39<i class="fa fa-fw fa-star"></i></a><a class="jsAnchor thumb-tags-toggle tagged" data-href="https://wallhaven.cc/wallpaper/tags/6qgqq6" original-title="Tags"><i class="fas fa-fw fa-tags"></i></a></div></figure></li><li><figure class="thumb thumb-n6pv6w thumb-sfw thumb-general" data-wallpaper-id="n6pv6w" style="width:300px;height:200px"><a class="thumb-btn thumb-btn-fav jsAnchor overlay-anchor" href="https://wallhaven.cc/favorites/fav/n6pv6w" original-title="Add to favorites"><i class="fas fa-fw fa-star"></i></a><img alt="loading" class="lazyload loaded" data-src="https://th.wallhaven.cc/small/n6/n6pv6w.jpg" src="https://th.wallhaven.cc/small/n6/n6pv6w.jpg"><a class="preview" href="https://wallhaven.cc/w/n6pv6w" target="_blank"></a><div class="thumb-info"><span class="wall-res">2560 x 1600</span><a class="jsAnchor overlay-anchor wall-favs" data-href="https://wallhaven.cc/wallpaper/fav/n6pv6w">10<i class="fa fa-fw fa-star"></i></a><a class="jsAnchor thumb-tags-toggle tagged" data-href="https://wallhaven.cc/wallpaper/tags/n6pv6w" original-title="Tags"><i class="fas fa-fw fa-tags"></i></a></div></figure></li><li><figure class="thumb thumb-n6d6wq thumb-sfw thumb-general" data-wallpaper-id="n6d6wq" style="width:300px;height:200px"><a class="thumb-btn thumb-btn-fav jsAnchor overlay-anchor" href="https://wallhaven.cc/favorites/fav/n6d6wq" original-title="Add to favorites"><i class="fas fa-fw fa-star"></i></a><img alt="loading" class="lazyload loaded" data-src="https://th.wallhaven.cc/small/n6/n6d6wq.jpg" src="https://th.wallhaven.cc/small/n6/n6d6wq.jpg"><a class="preview" href="https://wallhaven.cc/w/n6d6wq" target="_blank"></a><div class="thumb-info"><span class="wall-res">2560 x 1600</span><a class="jsAnchor overlay-anchor wall-favs" data-href="https://wallhaven.cc/wallpaper/fav/n6d6wq">1<i class="fa fa-fw fa-star"></i></a><span class="png"><span>PNG</span></span><a class="jsAnchor thumb-tags-toggle tagged" data-href="https://wallhaven.cc/wallpaper/tags/n6d6wq" original-title="Tags"><i class="fas fa-fw fa-tags"></i></a></div></figure></li><li><figure class="thumb thumb-nmy55m thumb-sfw thumb-general" data-wallpaper-id="nmy55m" style="width:300px;height:200px"><a class="thumb-btn thumb-btn-fav jsAnchor overlay-anchor" href="https://wallhaven.cc/favorites/fav/nmy55m" original-title="Add to favorites"><i class="fas fa-fw fa-star"></i></a><img alt="loading" class="lazyload loaded" data-src="https://th.wallhaven.cc/small/nm/nmy55m.jpg" src="https://th.wallhaven.cc/small/nm/nmy55m.jpg"><a class="preview" href="https://wallhaven.cc/w/nmy55m" target="_blank"></a><div class="thumb-info"><span class="wall-res">1792 x 1344</span><a class="jsAnchor overlay-anchor wall-favs" data-href="https://wallhaven.cc/wallpaper/fav/nmy55m">15<i class="fa fa-fw fa-star"></i></a><span class="png"><span>PNG</span></span><a class="jsAnchor thumb-tags-toggle tagged" data-href="https://wallhaven.cc/wallpaper/tags/nmy55m" original-title="Tags"><i class="fas fa-fw fa-tags"></i></a></div></figure></li><li><figure class="thumb thumb-48oxzk thumb-sfw thumb-anime" data-wallpaper-id="48oxzk" style="width:300px;height:200px"><a class="thumb-btn thumb-btn-fav jsAnchor overlay-anchor" href="https://wallhaven.cc/favorites/fav/48oxzk" original-title="Add to favorites"><i class="fas fa-fw fa-star"></i></a><img alt="loading" class="lazyload loaded" data-src="https://th.wallhaven.cc/small/48/48oxzk.jpg" src="https://th.wallhaven.cc/small/48/48oxzk.jpg"><a class="preview" href="https://wallhaven.cc/w/48oxzk" target="_blank"></a><div class="thumb-info"><span class="wall-res">1600 x 1200</span><a class="jsAnchor overlay-anchor wall-favs" data-href="https://wallhaven.cc/wallpaper/fav/48oxzk">7<i class="fa fa-fw fa-star"></i></a><a class="jsAnchor thumb-tags-toggle tagged" data-href="https://wallhaven.cc/wallpaper/tags/48oxzk" original-title="Tags"><i class="fas fa-fw fa-tags"></i></a></div></figure></li><li><figure class="thumb thumb-45pqm7 thumb-sketchy thumb-people" data-wallpaper-id="45pqm7" style="width:300px;height:200px"><a class="thumb-btn thumb-btn-fav jsAnchor overlay-anchor" href="https://wallhaven.cc/favorites/fav/45pqm7" original-title="Add to favorites"><i class="fas fa-fw fa-star"></i></a><img alt="loading" class="lazyload loaded" data-src="https://th.wallhaven.cc/small/45/45pqm7.jpg" src="https://th.wallhaven.cc/small/45/45pqm7.jpg"><a class="preview" href="https://wallhaven.cc/w/45pqm7" target="_blank"></a><div class="thumb-info"><span class="wall-res">1920 x 1197</span><a class="jsAnchor overlay-anchor wall-favs" data-href="https://wallhaven.cc/wallpaper/fav/45pqm7">28<i class="fa fa-fw fa-star"></i></a><a class="jsAnchor thumb-tags-toggle tagged" data-href="https://wallhaven.cc/wallpaper/tags/45pqm7" original-title="Tags"><i class="fas fa-fw fa-tags"></i></a></div></figure></li><li><figure class="thumb thumb-47q3ke thumb-sketchy thumb-people" data-wallpaper-id="47q3ke" style="width:300px;height:200px"><a class="thumb-btn thumb-btn-fav jsAnchor overlay-anchor" href="https://wallhaven.cc/favorites/fav/47q3ke" original-title="Add to favorites"><i class="fas fa-fw fa-star"></i></a><img alt="loading" class="lazyload loaded" data-src="https://th.wallhaven.cc/small/47/47q3ke.jpg" src="https://th.wallhaven.cc/small/47/47q3ke.jpg"><a class="preview" href="https://wallhaven.cc/w/47q3ke" target="_blank"></a><div class="thumb-info"><span class="wall-res">2560 x 1707</span><a class="jsAnchor overlay-anchor wall-favs" data-href="https://wallhaven.cc/wallpaper/fav/47q3ke">26<i class="fa fa-fw fa-star"></i></a><a class="jsAnchor thumb-tags-toggle tagged" data-href="https://wallhaven.cc/wallpaper/tags/47q3ke" original-title="Tags"><i class="fas fa-fw fa-tags"></i></a></div></figure></li><li><figure class="thumb thumb-r2qxew thumb-sketchy thumb-anime" data-wallpaper-id="r2qxew" style="width:300px;height:200px"><a class="thumb-btn thumb-btn-fav jsAnchor overlay-anchor" href="https://wallhaven.cc/favorites/fav/r2qxew" original-title="Add to favorites"><i class="fas fa-fw fa-star"></i></a><img alt="loading" class="lazyload loaded" data-src="https://th.wallhaven.cc/small/r2/r2qxew.jpg" src="https://th.wallhaven.cc/small/r2/r2qxew.jpg"><a class="preview" href="https://wallhaven.cc/w/r2qxew" target="_blank"></a><div class="thumb-info"><span class="wall-res">1980 x 1261</span><a class="jsAnchor overlay-anchor wall-favs" data-href="https://wallhaven.cc/wallpaper/fav/r2qxew">51<i class="fa fa-fw fa-star"></i></a><span class="png"><span>PNG</span></span><a class="jsAnchor thumb-tags-toggle tagged" data-href="https://wallhaven.cc/wallpaper/tags/r2qxew" original-title="Tags"><i class="fas fa-fw fa-tags"></i></a></div></figure></li><li><figure class="thumb thumb-x817go thumb-sfw thumb-general" data-wallpaper-id="x817go" style="width:300px;height:200px"><a class="thumb-btn thumb-btn-fav jsAnchor overlay-anchor" href="https://wallhaven.cc/favorites/fav/x817go" original-title="Add to favorites"><i class="fas fa-fw fa-star"></i></a><img alt="loading" class="lazyload loaded" data-src="https://th.wallhaven.cc/small/x8/x817go.jpg" src="https://th.wallhaven.cc/small/x8/x817go.jpg"><a class="preview" href="https://wallhaven.cc/w/x817go" target="_blank"></a><div class="thumb-info"><span class="wall-res">3840 x 2160</span><a class="jsAnchor overlay-anchor wall-favs" data-href="https://wallhaven.cc/wallpaper/fav/x817go">6<i class="fa fa-fw fa-star"></i></a><span class="png"><span>PNG</span></span><a class="jsAnchor thumb-tags-toggle tagged" data-href="https://wallhaven.cc/wallpaper/tags/x817go" original-title="Tags"><i class="fas fa-fw fa-tags"></i></a></div></figure></li><li><figure class="thumb thumb-ox7vkp thumb-sketchy thumb-people" data-wallpaper-id="ox7vkp" style="width:300px;height:200px"><a class="thumb-btn thumb-btn-fav jsAnchor overlay-anchor" href="https://wallhaven.cc/favorites/fav/ox7vkp" original-title="Add to favorites"><i class="fas fa-fw fa-star"></i></a><img alt="loading" class="lazyload loaded" data-src="https://th.wallhaven.cc/small/ox/ox7vkp.jpg" src="https://th.wallhaven.cc/small/ox/ox7vkp.jpg"><a class="preview" href="https://wallhaven.cc/w/ox7vkp" target="_blank"></a><div class="thumb-info"><span class="wall-res">2048 x 1248</span><a class="jsAnchor overlay-anchor wall-favs" data-href="https://wallhaven.cc/wallpaper/fav/ox7vkp">27<i class="fa fa-fw fa-star"></i></a><a class="jsAnchor thumb-tags-toggle tagged" data-href="https://wallhaven.cc/wallpaper/tags/ox7vkp" original-title="Tags"><i class="fas fa-fw fa-tags"></i></a></div></figure></li><li><figure class="thumb thumb-v9o8j8 thumb-sfw thumb-general" data-wallpaper-id="v9o8j8" style="width:300px;height:200px"><a class="thumb-btn thumb-btn-fav jsAnchor overlay-anchor" href="https://wallhaven.cc/favorites/fav/v9o8j8" original-title="Add to favorites"><i class="fas fa-fw fa-star"></i></a><img alt="loading" class="lazyload loaded" data-src="https://th.wallhaven.cc/small/v9/v9o8j8.jpg" src="https://th.wallhaven.cc/small/v9/v9o8j8.jpg"><a class="preview" href="https://wallhaven.cc/w/v9o8j8" target="_blank"></a><div class="thumb-info"><span class="wall-res">1920 x 1147</span><a class="jsAnchor overlay-anchor wall-favs" data-href="https://wallhaven.cc/wallpaper/fav/v9o8j8">24<i class="fa fa-fw fa-star"></i></a><a class="jsAnchor thumb-tags-toggle tagged" data-href="https://wallhaven.cc/wallpaper/tags/v9o8j8" original-title="Tags"><i class="fas fa-fw fa-tags"></i></a></div></figure></li><li><figure class="thumb thumb-lqqywq thumb-sfw thumb-general" data-wallpaper-id="lqqywq" style="width:300px;height:200px"><a class="thumb-btn thumb-btn-fav jsAnchor overlay-anchor" href="https://wallhaven.cc/favorites/fav/lqqywq" original-title="Add to favorites"><i class="fas fa-fw fa-star"></i></a><img alt="loading" class="lazyload loaded" data-src="https://th.wallhaven.cc/small/lq/lqqywq.jpg" src="https://th.wallhaven.cc/small/lq/lqqywq.jpg"><a class="preview" href="https://wallhaven.cc/w/lqqywq" target="_blank"></a><div class="thumb-info"><span class="wall-res">1920 x 1200</span><a class="jsAnchor overlay-anchor wall-favs" data-href="https://wallhaven.cc/wallpaper/fav/lqqywq">23<i class="fa fa-fw fa-star"></i></a><a class="jsAnchor thumb-tags-toggle tagged" data-href="https://wallhaven.cc/wallpaper/tags/lqqywq" original-title="Tags"><i class="fas fa-fw fa-tags"></i></a></div></figure></li><li><figure class="thumb thumb-ymevgx thumb-sketchy thumb-people" data-wallpaper-id="ymevgx" style="width:300px;height:200px"><a class="thumb-btn thumb-btn-fav jsAnchor overlay-anchor" href="https://wallhaven.cc/favorites/fav/ymevgx" original-title="Add to favorites"><i class="fas fa-fw fa-star"></i></a><img alt="loading" class="lazyload loaded" data-src="https://th.wallhaven.cc/small/ym/ymevgx.jpg" src="https://th.wallhaven.cc/small/ym/ymevgx.jpg"><a class="preview" href="https://wallhaven.cc/w/ymevgx" target="_blank"></a><div class="thumb-info"><span class="wall-res">1600 x 1066</span><a class="jsAnchor overlay-anchor wall-favs" data-href="https://wallhaven.cc/wallpaper/fav/ymevgx">27<i class="fa fa-fw fa-star"></i></a><a class="jsAnchor thumb-tags-toggle tagged" data-href="https://wallhaven.cc/wallpaper/tags/ymevgx" original-title="Tags"><i class="fas fa-fw fa-tags"></i></a></div></figure></li><li><figure class="thumb thumb-e7p9jw thumb-sfw thumb-anime" data-wallpaper-id="e7p9jw" style="width:300px;height:200px"><a class="thumb-btn thumb-btn-fav jsAnchor overlay-anchor" href="https://wallhaven.cc/favorites/fav/e7p9jw" original-title="Add to favorites"><i class="fas fa-fw fa-star"></i></a><img alt="loading" class="lazyload loaded" data-src="https://th.wallhaven.cc/small/e7/e7p9jw.jpg" src="https://th.wallhaven.cc/small/e7/e7p9jw.jpg"><a class="preview" href="https://wallhaven.cc/w/e7p9jw" target="_blank"></a><div class="thumb-info"><span class="wall-res">1085 x 1505</span><a class="jsAnchor overlay-anchor wall-favs" data-href="https://wallhaven.cc/wallpaper/fav/e7p9jw">5<i class="fa fa-fw fa-star"></i></a><a class="jsAnchor thumb-tags-toggle tagged" data-href="https://wallhaven.cc/wallpaper/tags/e7p9jw" original-title="Tags"><i class="fas fa-fw fa-tags"></i></a></div></figure></li><li><figure class="thumb thumb-q651kq thumb-sketchy thumb-people" data-wallpaper-id="q651kq" style="width:300px;height:200px"><a class="thumb-btn thumb-btn-fav jsAnchor overlay-anchor" href="https://wallhaven.cc/favorites/fav/q651kq" original-title="Add to favorites"><i class="fas fa-fw fa-star"></i></a><img alt="loading" class="lazyload loaded" data-src="https://th.wallhaven.cc/small/q6/q651kq.jpg" src="https://th.wallhaven.cc/small/q6/q651kq.jpg"><a class="preview" href="https://wallhaven.cc/w/q651kq" target="_blank"></a><div class="thumb-info"><span class="wall-res">1920 x 1080</span><a class="jsAnchor overlay-anchor wall-favs" data-href="https://wallhaven.cc/wallpaper/fav/q651kq">62<i class="fa fa-fw fa-star"></i></a><a class="jsAnchor thumb-tags-toggle tagged" data-href="https://wallhaven.cc/wallpaper/tags/q651kq" original-title="Tags"><i class="fas fa-fw fa-tags"></i></a></div></figure></li><li><figure class="thumb thumb-kwqe76 thumb-sketchy thumb-people" data-wallpaper-id="kwqe76" style="width:300px;height:200px"><a class="thumb-btn thumb-btn-fav jsAnchor overlay-anchor" href="https://wallhaven.cc/favorites/fav/kwqe76" original-title="Add to favorites"><i class="fas fa-fw fa-star"></i></a><img alt="loading" class="lazyload loaded" data-src="https://th.wallhaven.cc/small/kw/kwqe76.jpg" src="https://th.wallhaven.cc/small/kw/kwqe76.jpg"><a class="preview" href="https://wallhaven.cc/w/kwqe76" target="_blank"></a><div class="thumb-info"><span class="wall-res">2000 x 1333</span><a class="jsAnchor overlay-anchor wall-favs" data-href="https://wallhaven.cc/wallpaper/fav/kwqe76">48<i class="fa fa-fw fa-star"></i></a><a class="jsAnchor thumb-tags-toggle tagged" data-href="https://wallhaven.cc/wallpaper/tags/kwqe76" original-title="Tags"><i class="fas fa-fw fa-tags"></i></a></div></figure></li><li><figure class="thumb thumb-nm85zm thumb-sfw thumb-people" data-wallpaper-id="nm85zm" style="width:300px;height:200px"><a class="thumb-btn thumb-btn-fav jsAnchor overlay-anchor" href="https://wallhaven.cc/favorites/fav/nm85zm" original-title="Add to favorites"><i class="fas fa-fw fa-star"></i></a><img alt="loading" class="lazyload loaded" data-src="https://th.wallhaven.cc/small/nm/nm85zm.jpg" src="https://th.wallhaven.cc/small/nm/nm85zm.jpg"><a class="preview" href="https://wallhaven.cc/w/nm85zm" target="_blank"></a><div class="thumb-info"><span class="wall-res">5760 x 3840</span><a class="jsAnchor overlay-anchor wall-favs" data-href="https://wallhaven.cc/wallpaper/fav/nm85zm">17<i class="fa fa-fw fa-star"></i></a><a class="jsAnchor thumb-tags-toggle tagged" data-href="https://wallhaven.cc/wallpaper/tags/nm85zm" original-title="Tags"><i class="fas fa-fw fa-tags"></i></a></div></figure></li><li><figure class="thumb thumb-mpx119 thumb-sfw thumb-people" data-wallpaper-id="mpx119" style="width:300px;height:200px"><a class="thumb-btn thumb-btn-fav jsAnchor overlay-anchor" href="https://wallhaven.cc/favorites/fav/mpx119" original-title="Add to favorites"><i class="fas fa-fw fa-star"></i></a><img alt="loading" class="lazyload loaded" data-src="https://th.wallhaven.cc/small/mp/mpx119.jpg" src="https://th.wallhaven.cc/small/mp/mpx119.jpg"><a class="preview" href="https://wallhaven.cc/w/mpx119" target="_blank"></a><div class="thumb-info"><span class="wall-res">2560 x 1440</span><a class="jsAnchor overlay-anchor wall-favs" data-href="https://wallhaven.cc/wallpaper/fav/mpx119">37<i class="fa fa-fw fa-star"></i></a><a class="jsAnchor thumb-tags-toggle tagged" data-href="https://wallhaven.cc/wallpaper/tags/mpx119" original-title="Tags"><i class="fas fa-fw fa-tags"></i></a></div></figure></li><li><figure class="thumb thumb-eyo5or thumb-sfw thumb-general" data-wallpaper-id="eyo5or" style="width:300px;height:200px"><a class="thumb-btn thumb-btn-fav jsAnchor overlay-anchor" href="https://wallhaven.cc/favorites/fav/eyo5or" original-title="Add to favorites"><i class="fas fa-fw fa-star"></i></a><img alt="loading" class="lazyload loaded" data-src="https://th.wallhaven.cc/small/ey/eyo5or.jpg" src="https://th.wallhaven.cc/small/ey/eyo5or.jpg"><a class="preview" href="https://wallhaven.cc/w/eyo5or" target="_blank"></a><div class="thumb-info"><span class="wall-res">1920 x 1080</span><a class="jsAnchor overlay-anchor wall-favs" data-href="https://wallhaven.cc/wallpaper/fav/eyo5or">5<i class="fa fa-fw fa-star"></i></a><a class="jsAnchor thumb-tags-toggle tagged" data-href="https://wallhaven.cc/wallpaper/tags/eyo5or" original-title="Tags"><i class="fas fa-fw fa-tags"></i></a></div></figure></li><li><figure class="thumb thumb-4d337o thumb-sfw thumb-anime thumb-seen" data-wallpaper-id="4d337o" style="width:300px;height:200px"><a class="thumb-btn thumb-btn-fav jsAnchor overlay-anchor" href="https://wallhaven.cc/favorites/fav/4d337o" original-title="Add to favorites"><i class="fas fa-fw fa-star"></i></a><img alt="loading" class="lazyload loaded" data-src="https://th.wallhaven.cc/small/4d/4d337o.jpg" src="https://th.wallhaven.cc/small/4d/4d337o.jpg"><a class="preview" href="https://wallhaven.cc/w/4d337o" target="_blank"></a><div class="thumb-info"><span class="wall-res">1158 x 818</span><a class="jsAnchor overlay-anchor wall-favs" data-href="https://wallhaven.cc/wallpaper/fav/4d337o">6<i class="fa fa-fw fa-star"></i></a><a class="jsAnchor thumb-tags-toggle tagged" data-href="https://wallhaven.cc/wallpaper/tags/4d337o" original-title="Tags"><i class="fas fa-fw fa-tags"></i></a></div></figure></li><li><figure class="thumb thumb-4lq132 thumb-sfw thumb-general thumb-seen" data-wallpaper-id="4lq132" style="width:300px;height:200px"><a class="thumb-btn thumb-btn-fav jsAnchor overlay-anchor" href="https://wallhaven.cc/favorites/fav/4lq132" original-title="Add to favorites"><i class="fas fa-fw fa-star"></i></a><img alt="loading" class="lazyload loaded" data-src="https://th.wallhaven.cc/small/4l/4lq132.jpg" src="https://th.wallhaven.cc/small/4l/4lq132.jpg"><a class="preview" href="https://wallhaven.cc/w/4lq132" target="_blank"></a><div class="thumb-info"><span class="wall-res">1280 x 800</span><a class="jsAnchor overlay-anchor wall-favs" data-href="https://wallhaven.cc/wallpaper/fav/4lq132">5<i class="fa fa-fw fa-star"></i></a><a class="jsAnchor thumb-tags-toggle tagged" data-href="https://wallhaven.cc/wallpaper/tags/4lq132" original-title="Tags"><i class="fas fa-fw fa-tags"></i></a></div></figure></li><li><figure class="thumb thumb-1j6ze1 thumb-sfw thumb-anime" data-wallpaper-id="1j6ze1" style="width:300px;height:200px"><a class="thumb-btn thumb-btn-fav jsAnchor overlay-anchor" href="https://wallhaven.cc/favorites/fav/1j6ze1" original-title="Add to favorites"><i class="fas fa-fw fa-star"></i></a><img alt="loading" class="lazyload loaded" data-src="https://th.wallhaven.cc/small/1j/1j6ze1.jpg" src="https://th.wallhaven.cc/small/1j/1j6ze1.jpg"><a class="preview" href="https://wallhaven.cc/w/1j6ze1" target="_blank"></a><div class="thumb-info"><span class="wall-res">1500 x 912</span><a class="jsAnchor overlay-anchor wall-favs" data-href="https://wallhaven.cc/wallpaper/fav/1j6ze1">14<i class="fa fa-fw fa-star"></i></a><a class="jsAnchor thumb-tags-toggle tagged" data-href="https://wallhaven.cc/wallpaper/tags/1j6ze1" original-title="Tags"><i class="fas fa-fw fa-tags"></i></a></div></figure></li></ul><a class="thumb-btn thumb-btn-fav jsAnchor overlay-anchor" href="https://wallhaven.cc/favorites/fav/0qr5m5" original-title="Add to favorites"><i class="fas fa-fw fa-star"></i></a></section>
	c.OnHTML("section.thumb-listing-page", func(e *colly.HTMLElement) {
		//f := e.DOM.First().Children().Children().Siblings()
		//println(f)
		e.ForEach("a[href]", func(i int, element *colly.HTMLElement) {
			d := c.Clone()
			a := element.Attr("href")
			if a == "#top" {
				return
			}
			log.Println("down visit ", a)

			// 直接拼凑img地址
			// https://wallhaven.cc/w/9mwzyk ==>
			// https://w.wallhaven.cc/full/9m/wallhaven-9mwzyk.jpg
			ix := strings.LastIndex(a, "/")

			imgUrl := fmt.Sprintf("https://w.wallhaven.cc/full/%s/wallhaven-%s.jpg", a[ix+1:ix+3], a[ix+1:])
			log.Println("downloading ", imgUrl)
			d.Visit(imgUrl)
			ch := make(chan interface{})

			d.OnResponse(func(r *colly.Response) {
				defer close(ch)
				if strings.Index(r.Headers.Get("Content-Type"), "image") > -1 {
					r.Save(outputDir + r.FileName())
					return
				}
			})

			d.OnError(func(r *colly.Response, err error) {
				log.Println("error:", err, r.Request.URL, string(r.Body))
				close(ch)
			})

			// obsolete
			d.OnHTML("#wallpaper", func(e *colly.HTMLElement) {
				detail := e.Attr("src")
				log.Println("downloading ", detail)

				m := d.Clone()
				m.Visit(detail)
				m.OnResponse(func(r *colly.Response) {
					close(ch)
					if strings.Index(r.Headers.Get("Content-Type"), "image") > -1 {
						r.Save(outputDir + r.FileName())
						return
					}
				})
				m.OnError(func(r *colly.Response, err error) {
					close(ch)
					log.Println("error:", err, r.Request.URL, string(r.Body))
				})
			})

			select {
			case _ = <-ch:
			}
		})
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("error:", err, r.Request.URL, string(r.Body))
	})

	c.OnResponse(func(r *colly.Response) {
		//log.Println("rs:", err, r.Request.URL, string(r.Body))
	})

	//c.Visit("https://w.wallhaven.cc/full/8x/wallhaven-8xxl31.jpg")
	for i := 1; i < 10; i++ {
		a := fmt.Sprintf(url, i)
		log.Println("visit ", a)
		c.Visit(a)
	}
	c.Wait()

	for true {
		time.Sleep(1 * time.Second)
	}
}
