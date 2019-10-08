
package main

import (
	//	"bufio"
	"log" // "github.com/cihub/seelog"
	"net/http"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)
import (
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strings"
)

// const (
//     DB_USER     = "postgres"
//     DB_PASSWORD = "postgres"
//     DB_LOCATION = "localhost"
//     DB_NAME     = "postgres"
//     DB_SSLMODE  = "disable"
// )
// http://phppgadmin-cfready-iptv.cfapps.io/redirect.php?server=pellefant-02.db.elephantsql.com%3A5432%3Aallow&subject=database&database=zvmfafqf
// "postgres://zvmfafqf:3D9nQGb1LUrdJoSWMRcL1KtaFUAvbVkR@pellefant-02.db.elephantsql.com:5432/zvmfafqf"

//const (
//	DB_USER     = "zvmfafqf"
//	DB_PASSWORD = "3D9nQGb1LUrdJoSWMRcL1KtaFUAvbVkR"
//	DB_LOCATION = "pellefant-02.db.elephantsql.com:5432"
//	DB_NAME     = "zvmfafqf"
//	DB_SSLMODE  = "disable" //verify-full"
//)

const (
	DEFAULT_PORT = "9000"
	EXTM3U       = "#EXTM3U"
	EXTINF       = "#EXTINF:"
)

var bFF_USEDB = false
var bFF_ENABLEDBLOG = false
var bFF_DBRECREATE = false
var PGdb *gorm.DB

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("HomeHandler Starting")
	fmt.Fprintln(w, "Hello, GO World!n")
	//	fmt.Fprintln(w, "#EXTM3U")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,About Script")
	//	fmt.Fprintln(w, "http://on.meshra.com/tiptv/about.wmv")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Angel TV")
	//	fmt.Fprintln(w, "http://doh57iesfegop.cloudfront.net/angeltv/angeltvindia600/playlist.m3u8")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Madha TV")
	//	fmt.Fprintln(w, "rtmp://103.49.238.70:1935/madhatv/_definst_playpath=mp4:myStream_source swfUrl=http://p.jwpcdn.com/player/v/7.2.2/jwplayer.flash.swf pageUrl=http://madha.tv/live-player.php")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Anboli TV")
	//	fmt.Fprintln(w, "rtmp://38.96.148.172:1935/live/ playpath=anboli swfUrl=http://live.akamain.info/player/jw6/jwplayer/jwplayer.flash.swf pageUrl=http://live.akamain.info/player/files/anboli/index.html")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Sri Sankara TV")
	//	fmt.Fprintln(w, "rtmp://live.wmncdn.net/srisankaratv1/ playpath=986995610b14c8529ce441b18236d4c7.sdp swfUrl=http://player-apac-sing.webmobilive.com/player/liveplayer.swf pageUrl=http://www.srisankaratv.com/watchtv.php")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,News 7 Tamil")
	//	fmt.Fprintln(w, "http://d2voe16cs5psaw.cloudfront.net/news7new/smil:news7new.smil/playlist.m3u8")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Athavan TV")
	//	fmt.Fprintln(w, "http://play-fs.amgiptv.com/vxtvathavan/athavantvlive.stream/playlist.m3u8")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,EET TV")
	//	fmt.Fprintln(w, "http://37.59.17.222:1935/live_ca/eettv/playlist.m3u8")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Arra TV")
	//	fmt.Fprintln(w, "http://arratvcloud.purplestream.in/arratvorg/smil:arratv.smil/playlist.m3u8")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Imayam TV")
	//	fmt.Fprintln(w, "rtmp://38.96.148.172:1935/live/ playpath=imayamtv swfUrl=http://live.akamain.info/v1/imayamtv/jwplayer/jwplayer.flash.swf pageUrl=http://live.akamain.info/v1/imayamtv/index.html live=1")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Lotus News")
	//	fmt.Fprintln(w, "http://118.102.228.195:1935/live/stream/playlist.m3u8")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Puthiya Thalaimurai")
	//	fmt.Fprintln(w, "http://dzamqgwxt2cln.cloudfront.net/ngmedia/ptm_400/playlist.m3u8")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Puthuyugam")
	//	fmt.Fprintln(w, "http://d211bpuke56xnv.cloudfront.net/ngmedia/pym_400/playlist.m3u8")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Vasantham TV")
	//	fmt.Fprintln(w, "http://cdncities.com/vasanthamtv/vasanthamtv/playlist.m3u8")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Dheeran TV")
	//	fmt.Fprintln(w, "http://edgecastdheerantv.purplestream.in/dheeranorg/ngrp:dheeran_all/playlist.m3u8")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Thanthi TV")
	//	fmt.Fprintln(w, "http://edgecastthanthitv.purplestream.in/httporg/ngrp:thanthi_all/playlist.m3u8")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Vimbam TV")
	//	fmt.Fprintln(w, "rtmp://91.219.43.148:1935/vimbamtv/ playpath=vimbamtv swfUrl=http://p.jwpcdn.com/6/12/jwplayer.flash.swf pageUrl=http://www.vimbamtv.com/movie.html")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Vendhar TV")
	//	fmt.Fprintln(w, "rtmp://45.79.203.234:1935/vendhar/ playpath=myStream swfUrl=http://vendharmedia.in/video-js.swf pageUrl=http://vendharmedia.in/")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,IBC Tamil")
	//	fmt.Fprintln(w, "http://ibctamil.zecast.net/ibctamil/smil:ibctamil.smil/playlist.m3u8")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Vasanth TV")
	//	fmt.Fprintln(w, "http://vasanth.live.cdn.bitgravity.com/vasanth/secure/live/feed01?e=0%26h=a9be0836bc39f96d0a9a958a659dfc1d%26bgsecuredir=1&amp;autoPlay=true")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Makkal TV")
	//	fmt.Fprintln(w, "http://edgecastmakkaltv.purplestream.in/makkalorg/ngrp:makkaltv_all/playlist.m3u8")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Masala TV")
	//	fmt.Fprintln(w, "rtmp://188.166.127.249/masalatv/ playpath=myStream swfUrl=http://www.mediaworldasia.dk/media/yendifvideoshare/player/player.swf?1462632676628 pageUrl=http://www.mediaworldasia.dk/index.php?option=com_yendifvideoshare&amp;view=video&amp;id=124&amp;tmpl=component")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,TTN Tamil Oli")
	//	fmt.Fprintln(w, "http://166.62.121.101:1935/ttn/ttn/playlist.m3u8")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Vanavil Plus")
	//	fmt.Fprintln(w, "rtmp://server.maxwellstreaming.com/vanavil playpath=myStream swfUrl=http://www.maxwellstreaming.com/player.swf pageUrl=http://www.vanaviltv.in/live.php")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,DD Podhigai")
	//	fmt.Fprintln(w, "rtmp://live-fs.wmncdn.net/ddpodigailive1/ playpath=live1.stream swfUrl=http://p.jwpcdn.com/6/12/jwplayer.flash.swf pageUrl=http://www.ddpodhigai.org.in/live.html")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Kathir TV")
	//	fmt.Fprintln(w, "rtmp://server.maxwellstreaming.com:1935/kathir playpath=myStream swfUrl=http://www.livewebcast.in/jwplayer/player.swf pageUrl=http://kathirtv.com/livetv/")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Saras TV")
	//	fmt.Fprintln(w, "http://live.streamjo.com:1935/sarastv/sarastv/playlist.m3u8")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,1Yes TV")
	//	fmt.Fprintln(w, "http://stream.mslive.in:1967/live/gd416/playlist.m3u8")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Tamil Biz TV")
	//	fmt.Fprintln(w, "rtmp://rtmp.santhoratv.zecast.net/tamilbusinesstv// playpath=tamilbusinesstv swfUrl=http://p.jwpcdn.com/6/12/jwplayer.flash.swf pageUrl=http://livetvchannelsfree.com/tamilbusinesstv.html")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Santhora TV")
	//	fmt.Fprintln(w, "http://santhoratv.zecast.net/santhoratv/santhoratv/playlist.m3u8")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Peppers TV")
	//	fmt.Fprintln(w, "rtmp://live.wmncdn.net/pepperstv/ playpath=51f5cdae2cc7bd5e4e8e76b7357ae0f8.sdp swfUrl=http://p.jwpcdn.com/6/12/jwplayer.flash.swf pageUrl=http://livetvchannelsfree.com/pepperstv.html")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Test 01")
	//	fmt.Fprintln(w, "http://mobile.amgiptv.com/vxtvlqkalaingertv/kalaingertvlq.stream/playlist.m3u8")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Test 02")
	//	fmt.Fprintln(w, "http://play-fs.amgiptv.com/vxtvmegatv/megatv.stream/playlist.m3u8")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Test 03")
	//	fmt.Fprintln(w, "http://50.7.1.178/india_iptv/sunmusicINDIA/playlist.m3u8")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil Radio,Vanavil FM (R)")
	//	fmt.Fprintln(w, "http://s1.yesstreaming.net:9000/:stream/1")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil Radio,Mukil FM (R)")
	//	fmt.Fprintln(w, "http://live64.jiljilradio.com/;stream1.mp3")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil Radio,Kalasam FM (R)")
	//	fmt.Fprintln(w, "http://live.kalasam.com:8084")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil Radio,LankaSri FM (R)")
	//	fmt.Fprintln(w, "http://media2.lankasri.fm")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil Radio,Oli 96.8 FM (R)")
	//	fmt.Fprintln(w, "http://mediacorp.rastream.com/968fm?type=.flv")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil Radio,Oli 96.8 FM (R)")
	//	fmt.Fprintln(w, "http://mediacorp.rastream.com/968fm?type=.flv")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil Radio,THR Raaga (R)")
	//	fmt.Fprintln(w, "http://astro3.rastream.com:80/raaga?type=.flv")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil Radio,ETR FM (R)")
	//	fmt.Fprintln(w, "http://cast2.serverhostingcenter.com:8646/stream")
	//	fmt.Fprintln(w, "")
	//	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil Radio,Puradsi FM (R)")
	//	fmt.Fprintln(w, "http://puradsifm.net:9994")
	log.Println("HomeHandler Ending")
}

func V1ListHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("V1ListHandler Starting")

	sUrl := r.URL.Query()["type"]
	sType := ""
	if len(sUrl) > 0 {
		sType = strings.ToLower(sUrl[0])
		log.Println("type=" + sType)
	}
	if bFF_USEDB == true {
		channels := []Channel{}
		if len(sType) == 0 {
			PGdb.Preload("Groupdetails").Preload("Urls").Find(&channels)
		} else {
			PGdb.Joins("JOIN channel_groups on channel_groups.channel_id = channels.id").Where("channel_groups.valuelcase like ?", "%"+sType+"%").Preload("Groupdetails").Preload("Urls").Find(&channels)
		}

		fmt.Fprintln(w, EXTM3U)
		for _, v := range channels {
			fmt.Fprintln(w, v.StringM3U())
		}
	}

	log.Println("V1ListHandler Ending")
}

func V1ImportHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("V1ImportHandler Starting")

	//	var f *os.File
	//	var err error

	//	f, err = os.Open("final.m3u")
	//	if err != nil {
	//		panic(err)
	//	}

	//	var mw1 []Channel
	//	mw1, err = DecodeFrom(bufio.NewReader(f), "final.m3u")
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println("XXXXXXXXXXX")
	//	fmt.Println(mw1)

	//mq, _ := url.ParseQuery(r.URL.String())
	//	mq := r.URL.Query()
	//	for k, v := range mq {
	//		fmt.Fprintln(w, "k="+k+",v="+v[0])
	//		log.Println("k=" + k + ",v=" + v[0])
	//	}
	sUrl := r.URL.Query()["fromurl"]
	if len(sUrl) > 0 {
		log.Println("fromUrl=" + sUrl[0])

		res, err := http.Get(sUrl[0])
		if err != nil {
			log.Fatal(err)
		}

		_, err = DecodeFrom(res.Body, sUrl[0])
		res.Body.Close()
		if err != nil {
			panic(err)
		}

	} else {
		log.Println("fromUrl is empty. nothing to do")
	}
	log.Println("V1ImportHandler Ending")
}

func getEnvVar(varname string, defaultVal bool) bool {
	if sEnvVal := os.Getenv(varname); len(sEnvVal) == 0 {
		bEnvVal, err := strconv.ParseBool(sEnvVal)
		if err != nil {
			return defaultVal
		}
		return bEnvVal
	}
	return defaultVal
}

func main() {

	log.Println("App Started")
	var err error

	bFF_USEDB = getEnvVar("FF_USEDB", bFF_USEDB)
	log.Println("FF_USEDB set to %+vn", bFF_USEDB)

	bFF_ENABLEDBLOG = getEnvVar("FF_ENABLEDBLOG", bFF_ENABLEDBLOG)
	log.Println("FF_ENABLEDBLOG set to %+vn", bFF_ENABLEDBLOG)

	bFF_DBRECREATE = getEnvVar("FF_DBRECREATE", bFF_DBRECREATE)
	log.Println("FF_DBRECREATE set to %+vn", bFF_DBRECREATE)

	if bFF_USEDB == true {
		PGdb, err = gorm.Open("postgres", "user=postgres password=postgres DB.name=iptv sslmode=disable")
		if err != nil {
			panic("failed to connect database")
		}
		PGdb.LogMode(bFF_ENABLEDBLOG)

		if bFF_DBRECREATE == true {
			PGdb.DropTableIfExists(&Channel{})
			PGdb.DropTableIfExists(&ChannelUrl{})
			PGdb.DropTableIfExists(&ChannelGroup{})
		}
		// Migrate the schema
		PGdb.AutoMigrate(&Channel{})
		PGdb.AutoMigrate(&ChannelUrl{})
		PGdb.AutoMigrate(&ChannelGroup{})

		PGdb.Model(&Channel{}).Related(&ChannelUrl{})
		PGdb.Model(&Channel{}).Related(&ChannelGroup{})
	}

	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		log.Println("Warning, PORT not set. Defaulting to %+vn", DEFAULT_PORT)
		port = DEFAULT_PORT
	}
	//	port = DEFAULT_PORT

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/v1/list", V1ListHandler)
	http.HandleFunc("/v1/import", V1ImportHandler)

	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Println("ListenAndServe: ", err)
	}

}

type Channel struct {
	gorm.Model
	SeqNo        int
	Title        string
	Titlelcase   string `gorm:"unique_index:idx_title_source"` // Create index with title and source
	MainUrl      string
	Source       string
	Sourcelcase  string `gorm:"unique_index:idx_title_source"` // Create index with title and source
	Urls         []ChannelUrl
	Groupdetails []ChannelGroup
}
type ChannelUrl struct {
	gorm.Model
	Name       string
	Namelcase  string
	Value      string
	Valuelcase string
	ChannelID  uint
}
type ChannelGroup struct {
	gorm.Model
	Name       string
	Namelcase  string
	Value      string
	Valuelcase string
	ChannelID  uint
}
type MediaPlaylist struct {
	ChannelList []Channel
}

func (h *ChannelGroup) StringM3U() string {
	return fmt.Sprintf("%s=%s", h.Name, h.Value)
}
func (h *ChannelUrl) StringM3U() string {
	return fmt.Sprintf("%s=%s", h.Name, h.Value)
}
func (h *Channel) StringM3U() string {
	grpString := ""
	for _, v := range h.Groupdetails {
		grpString += v.StringM3U() + " "
	}
	urlString := ""
	for _, v := range h.Urls {
		urlString += v.StringM3U() + " "
	}

	retString := ""
	retString = fmt.Sprintf("%s-1 %s,%s\n", EXTINF, grpString, h.Title)
	retString = retString + h.MainUrl + " " + urlString
	return retString
}

// Detect playlist type and decode it from input stream.
func DecodeFrom(reader io.Reader, source string) ([]Channel, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(reader)
	if err != nil {
		return nil, err
	}
	return decode(buf, source)
}

// Decode media playlists.
func decode(buf *bytes.Buffer, source string) ([]Channel, error) {
	var eof bool
	var line string
	var err error

	box := MediaPlaylist{[]Channel{}}

	for !eof {
		if line, err = buf.ReadString('\n'); err == io.EOF {
			eof = true
		} else if err != nil {
			break
		}

		// fixes the issues https://github.com/grafov/m3u8/issues/25
		// TODO: the same should be done in decode functions of both Master- and MediaPlaylists
		// so some DRYing would be needed.
		if len(line) < 1 || line == "\r" {
			continue
		}

		err = decodeLineOfMediaPlaylist(&box, line, source)
		if err != nil {
			return box.ChannelList, err
		}
	}
	return box.ChannelList, nil
}

func decodeLineOfMediaPlaylist(p *MediaPlaylist, line string, source string) error {
	var title string
	var err error

	reGroups, _ := regexp.Compile(`(?P<groupKey>\w+-\w+\s*)=(?P<groupValue>\s*\"*\w+\s*\.*\w+\"*)`)           // matches group-title="Tamil TV" group-img="Tamil.jpg"
	reUrls, _ := regexp.Compile(`(\w+)=((([\w-]+://?)?|www[.])[^\s()<>]+(?:\([\w\d]+\)|([^[:punct:]\s]|/)))`) // matches rtmp://38.96.148.172:1935/live/ playpath=imayamtv swfUrl=http://live.akamain.info/v1/imayamtv/jwplayer/jwplayer.flash.swf pageUrl=http://live.akamain.info/v1/imayamtv/index.html live=1
	line = strings.TrimSpace(line)
	switch {
	case strings.HasPrefix(line, "#EXTINF:"):
		params := strings.Split(line[11:], ",")
		if len(params) > 0 {
			//last is Title
			title = params[len(params)-1]
		}
		channel := Channel{SeqNo: len(p.ChannelList) + 1, Title: title, Source: source, Titlelcase: strings.ToLower(title), Sourcelcase: strings.ToLower(source)}
		rMatchedGroups := reGroups.FindAllStringSubmatch(params[0], -1)
		for j, _ := range rMatchedGroups {
			channel.Groupdetails = append(channel.Groupdetails, ChannelGroup{Name: rMatchedGroups[j][1], Value: rMatchedGroups[j][2], Namelcase: strings.ToLower(rMatchedGroups[j][1]), Valuelcase: strings.ToLower(rMatchedGroups[j][2])})
		}

		p.ChannelList = append(p.ChannelList, channel)

	case !strings.HasPrefix(line, "#"):
		if len(line) > 0 {
			// There could be multiple urls like
			// rtmp://38.96.148.172:1935/live/ playpath=imayamtv swfUrl=http://live.akamain.info/v1/imayamtv/jwplayer/jwplayer.flash.swf pageUrl=http://live.akamain.info/v1/imayamtv/index.html live=1
			urls := strings.SplitN(line, " ", 2)

			if len(urls) > 0 && len(urls[0]) > 0 {
				channel := &p.ChannelList[len(p.ChannelList)-1] // Get the last item
				channel.MainUrl = urls[0]

				if len(urls) > 1 {
					rMatchedUrls := reUrls.FindAllStringSubmatch(urls[1], -1)
					for j, _ := range rMatchedUrls {
						channel.Urls = append(channel.Urls, ChannelUrl{Name: rMatchedUrls[j][1], Value: rMatchedUrls[j][2], Namelcase: strings.ToLower(rMatchedUrls[j][1]), Valuelcase: strings.ToLower(rMatchedUrls[j][2])})
					}

				}
				if bFF_USEDB == true {
					PGdb.Create(&channel)
				} else {
					log.Println(channel)
				}
			}
		}

	case strings.HasPrefix(line, "#"): // unknown tags treated as comments
		return err
	}
	title = title + ""

	return err
}
