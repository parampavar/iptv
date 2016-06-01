package main

import (
	"bufio"
	"log" // "github.com/cihub/seelog"
	"net/http"
	"os"
)
import (
	"bytes"
	"fmt"
	"io"
	//	"errors"
	//	"regexp"
	//	"strconv"
	"strings"
	//	"time"
	//	"encoding/json"
	"regexp"
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

const (
	DB_USER     = "zvmfafqf"
	DB_PASSWORD = "3D9nQGb1LUrdJoSWMRcL1KtaFUAvbVkR"
	DB_LOCATION = "pellefant-02.db.elephantsql.com:5432"
	DB_NAME     = "zvmfafqf"
	DB_SSLMODE  = "disable" //verify-full"
)

const (
	DEFAULT_PORT = "9000"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("HomeHandler Starting")
	//fmt.Fprintln(w, "Hello, GO World!n")
	fmt.Fprintln(w, "#EXTM3U")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,About Script")
	fmt.Fprintln(w, "http://on.meshra.com/tiptv/about.wmv")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Angel TV")
	fmt.Fprintln(w, "http://doh57iesfegop.cloudfront.net/angeltv/angeltvindia600/playlist.m3u8")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Madha TV")
	fmt.Fprintln(w, "rtmp://103.49.238.70:1935/madhatv/_definst_playpath=mp4:myStream_source swfUrl=http://p.jwpcdn.com/player/v/7.2.2/jwplayer.flash.swf pageUrl=http://madha.tv/live-player.php")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Anboli TV")
	fmt.Fprintln(w, "rtmp://38.96.148.172:1935/live/ playpath=anboli swfUrl=http://live.akamain.info/player/jw6/jwplayer/jwplayer.flash.swf pageUrl=http://live.akamain.info/player/files/anboli/index.html")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Sri Sankara TV")
	fmt.Fprintln(w, "rtmp://live.wmncdn.net/srisankaratv1/ playpath=986995610b14c8529ce441b18236d4c7.sdp swfUrl=http://player-apac-sing.webmobilive.com/player/liveplayer.swf pageUrl=http://www.srisankaratv.com/watchtv.php")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,News 7 Tamil")
	fmt.Fprintln(w, "http://d2voe16cs5psaw.cloudfront.net/news7new/smil:news7new.smil/playlist.m3u8")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Athavan TV")
	fmt.Fprintln(w, "http://play-fs.amgiptv.com/vxtvathavan/athavantvlive.stream/playlist.m3u8")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,EET TV")
	fmt.Fprintln(w, "http://37.59.17.222:1935/live_ca/eettv/playlist.m3u8")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Arra TV")
	fmt.Fprintln(w, "http://arratvcloud.purplestream.in/arratvorg/smil:arratv.smil/playlist.m3u8")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Imayam TV")
	fmt.Fprintln(w, "rtmp://38.96.148.172:1935/live/ playpath=imayamtv swfUrl=http://live.akamain.info/v1/imayamtv/jwplayer/jwplayer.flash.swf pageUrl=http://live.akamain.info/v1/imayamtv/index.html live=1")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Lotus News")
	fmt.Fprintln(w, "http://118.102.228.195:1935/live/stream/playlist.m3u8")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Puthiya Thalaimurai")
	fmt.Fprintln(w, "http://dzamqgwxt2cln.cloudfront.net/ngmedia/ptm_400/playlist.m3u8")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Puthuyugam")
	fmt.Fprintln(w, "http://d211bpuke56xnv.cloudfront.net/ngmedia/pym_400/playlist.m3u8")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Vasantham TV")
	fmt.Fprintln(w, "http://cdncities.com/vasanthamtv/vasanthamtv/playlist.m3u8")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Dheeran TV")
	fmt.Fprintln(w, "http://edgecastdheerantv.purplestream.in/dheeranorg/ngrp:dheeran_all/playlist.m3u8")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Thanthi TV")
	fmt.Fprintln(w, "http://edgecastthanthitv.purplestream.in/httporg/ngrp:thanthi_all/playlist.m3u8")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Vimbam TV")
	fmt.Fprintln(w, "rtmp://91.219.43.148:1935/vimbamtv/ playpath=vimbamtv swfUrl=http://p.jwpcdn.com/6/12/jwplayer.flash.swf pageUrl=http://www.vimbamtv.com/movie.html")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Vendhar TV")
	fmt.Fprintln(w, "rtmp://45.79.203.234:1935/vendhar/ playpath=myStream swfUrl=http://vendharmedia.in/video-js.swf pageUrl=http://vendharmedia.in/")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,IBC Tamil")
	fmt.Fprintln(w, "http://ibctamil.zecast.net/ibctamil/smil:ibctamil.smil/playlist.m3u8")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Vasanth TV")
	fmt.Fprintln(w, "http://vasanth.live.cdn.bitgravity.com/vasanth/secure/live/feed01?e=0%26h=a9be0836bc39f96d0a9a958a659dfc1d%26bgsecuredir=1&amp;autoPlay=true")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Makkal TV")
	fmt.Fprintln(w, "http://edgecastmakkaltv.purplestream.in/makkalorg/ngrp:makkaltv_all/playlist.m3u8")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Masala TV")
	fmt.Fprintln(w, "rtmp://188.166.127.249/masalatv/ playpath=myStream swfUrl=http://www.mediaworldasia.dk/media/yendifvideoshare/player/player.swf?1462632676628 pageUrl=http://www.mediaworldasia.dk/index.php?option=com_yendifvideoshare&amp;view=video&amp;id=124&amp;tmpl=component")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,TTN Tamil Oli")
	fmt.Fprintln(w, "http://166.62.121.101:1935/ttn/ttn/playlist.m3u8")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Vanavil Plus")
	fmt.Fprintln(w, "rtmp://server.maxwellstreaming.com/vanavil playpath=myStream swfUrl=http://www.maxwellstreaming.com/player.swf pageUrl=http://www.vanaviltv.in/live.php")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,DD Podhigai")
	fmt.Fprintln(w, "rtmp://live-fs.wmncdn.net/ddpodigailive1/ playpath=live1.stream swfUrl=http://p.jwpcdn.com/6/12/jwplayer.flash.swf pageUrl=http://www.ddpodhigai.org.in/live.html")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Kathir TV")
	fmt.Fprintln(w, "rtmp://server.maxwellstreaming.com:1935/kathir playpath=myStream swfUrl=http://www.livewebcast.in/jwplayer/player.swf pageUrl=http://kathirtv.com/livetv/")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Saras TV")
	fmt.Fprintln(w, "http://live.streamjo.com:1935/sarastv/sarastv/playlist.m3u8")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,1Yes TV")
	fmt.Fprintln(w, "http://stream.mslive.in:1967/live/gd416/playlist.m3u8")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Tamil Biz TV")
	fmt.Fprintln(w, "rtmp://rtmp.santhoratv.zecast.net/tamilbusinesstv// playpath=tamilbusinesstv swfUrl=http://p.jwpcdn.com/6/12/jwplayer.flash.swf pageUrl=http://livetvchannelsfree.com/tamilbusinesstv.html")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Santhora TV")
	fmt.Fprintln(w, "http://santhoratv.zecast.net/santhoratv/santhoratv/playlist.m3u8")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Peppers TV")
	fmt.Fprintln(w, "rtmp://live.wmncdn.net/pepperstv/ playpath=51f5cdae2cc7bd5e4e8e76b7357ae0f8.sdp swfUrl=http://p.jwpcdn.com/6/12/jwplayer.flash.swf pageUrl=http://livetvchannelsfree.com/pepperstv.html")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Test 01")
	fmt.Fprintln(w, "http://mobile.amgiptv.com/vxtvlqkalaingertv/kalaingertvlq.stream/playlist.m3u8")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Test 02")
	fmt.Fprintln(w, "http://play-fs.amgiptv.com/vxtvmegatv/megatv.stream/playlist.m3u8")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Test 03")
	fmt.Fprintln(w, "http://50.7.1.178/india_iptv/sunmusicINDIA/playlist.m3u8")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil Radio,Vanavil FM (R)")
	fmt.Fprintln(w, "http://s1.yesstreaming.net:9000/:stream/1")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil Radio,Mukil FM (R)")
	fmt.Fprintln(w, "http://live64.jiljilradio.com/;stream1.mp3")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil Radio,Kalasam FM (R)")
	fmt.Fprintln(w, "http://live.kalasam.com:8084")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil Radio,LankaSri FM (R)")
	fmt.Fprintln(w, "http://media2.lankasri.fm")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil Radio,Oli 96.8 FM (R)")
	fmt.Fprintln(w, "http://mediacorp.rastream.com/968fm?type=.flv")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil Radio,Oli 96.8 FM (R)")
	fmt.Fprintln(w, "http://mediacorp.rastream.com/968fm?type=.flv")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil Radio,THR Raaga (R)")
	fmt.Fprintln(w, "http://astro3.rastream.com:80/raaga?type=.flv")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil Radio,ETR FM (R)")
	fmt.Fprintln(w, "http://cast2.serverhostingcenter.com:8646/stream")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil Radio,Puradsi FM (R)")
	fmt.Fprintln(w, "http://puradsifm.net:9994")
	log.Println("HomeHandler Ending")
}

func V1Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("HomeHandler Starting")
	//fmt.Fprintln(w, "Hello, GO World!n")
	fmt.Fprintln(w, "#EXTM3U")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,News 7 Tamil")
	fmt.Fprintln(w, "http://d2voe16cs5psaw.cloudfront.net/news7new/smil:news7new.smil/playlist.m3u8")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Thanthi TV")
	fmt.Fprintln(w, "http://edgecastthanthitv.purplestream.in/httporg/ngrp:thanthi_all/playlist.m3u8")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Puthiya Thalaimurai")
	fmt.Fprintln(w, "http://dzamqgwxt2cln.cloudfront.net/ngmedia/ptm_400/playlist.m3u8")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Vendhar TV")
	fmt.Fprintln(w, "rtmp://45.79.203.234:1935/vendhar/ playpath=myStream swfUrl=http://vendharmedia.in/video-js.swf pageUrl=http://vendharmedia.in/")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,EET TV")
	fmt.Fprintln(w, "http://37.59.17.222:1935/live_ca/eettv/playlist.m3u8")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Vimbam TV")
	fmt.Fprintln(w, "rtmp://91.219.43.148:1935/vimbamtv/ playpath=vimbamtv swfUrl=http://p.jwpcdn.com/6/12/jwplayer.flash.swf pageUrl=http://www.vimbamtv.com/movie.html")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Peppers TV")
	fmt.Fprintln(w, "rtmp://live.wmncdn.net/pepperstv/ playpath=51f5cdae2cc7bd5e4e8e76b7357ae0f8.sdp swfUrl=http://p.jwpcdn.com/6/12/jwplayer.flash.swf pageUrl=http://livetvchannelsfree.com/pepperstv.html")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Athavan TV")
	fmt.Fprintln(w, "http://play-fs.amgiptv.com/vxtvathavan/athavantvlive.stream/playlist.m3u8")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Arra TV")
	fmt.Fprintln(w, "http://arratvcloud.purplestream.in/arratvorg/smil:arratv.smil/playlist.m3u8")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Imayam TV")
	fmt.Fprintln(w, "rtmp://38.96.148.172:1935/live/ playpath=imayamtv swfUrl=http://live.akamain.info/v1/imayamtv/jwplayer/jwplayer.flash.swf pageUrl=http://live.akamain.info/v1/imayamtv/index.html live=1")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Lotus News")
	fmt.Fprintln(w, "http://118.102.228.195:1935/live/stream/playlist.m3u8")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Puthuyugam")
	fmt.Fprintln(w, "http://d211bpuke56xnv.cloudfront.net/ngmedia/pym_400/playlist.m3u8")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Vasantham TV")
	fmt.Fprintln(w, "http://cdncities.com/vasanthamtv/vasanthamtv/playlist.m3u8")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Dheeran TV")
	fmt.Fprintln(w, "http://edgecastdheerantv.purplestream.in/dheeranorg/ngrp:dheeran_all/playlist.m3u8")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,IBC Tamil")
	fmt.Fprintln(w, "http://ibctamil.zecast.net/ibctamil/smil:ibctamil.smil/playlist.m3u8")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Vasanth TV")
	fmt.Fprintln(w, "http://vasanth.live.cdn.bitgravity.com/vasanth/secure/live/feed01?e=0%26h=a9be0836bc39f96d0a9a958a659dfc1d%26bgsecuredir=1&amp;autoPlay=true")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Makkal TV")
	fmt.Fprintln(w, "http://edgecastmakkaltv.purplestream.in/makkalorg/ngrp:makkaltv_all/playlist.m3u8")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Masala TV")
	fmt.Fprintln(w, "rtmp://188.166.127.249/masalatv/ playpath=myStream swfUrl=http://www.mediaworldasia.dk/media/yendifvideoshare/player/player.swf?1462632676628 pageUrl=http://www.mediaworldasia.dk/index.php?option=com_yendifvideoshare&amp;view=video&amp;id=124&amp;tmpl=component")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,TTN Tamil Oli")
	fmt.Fprintln(w, "http://166.62.121.101:1935/ttn/ttn/playlist.m3u8")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Vanavil Plus")
	fmt.Fprintln(w, "rtmp://server.maxwellstreaming.com/vanavil playpath=myStream swfUrl=http://www.maxwellstreaming.com/player.swf pageUrl=http://www.vanaviltv.in/live.php")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,DD Podhigai")
	fmt.Fprintln(w, "rtmp://live-fs.wmncdn.net/ddpodigailive1/ playpath=live1.stream swfUrl=http://p.jwpcdn.com/6/12/jwplayer.flash.swf pageUrl=http://www.ddpodhigai.org.in/live.html")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Kathir TV")
	fmt.Fprintln(w, "rtmp://server.maxwellstreaming.com:1935/kathir playpath=myStream swfUrl=http://www.livewebcast.in/jwplayer/player.swf pageUrl=http://kathirtv.com/livetv/")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Saras TV")
	fmt.Fprintln(w, "http://live.streamjo.com:1935/sarastv/sarastv/playlist.m3u8")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,1Yes TV")
	fmt.Fprintln(w, "http://stream.mslive.in:1967/live/gd416/playlist.m3u8")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Tamil Biz TV")
	fmt.Fprintln(w, "rtmp://rtmp.santhoratv.zecast.net/tamilbusinesstv// playpath=tamilbusinesstv swfUrl=http://p.jwpcdn.com/6/12/jwplayer.flash.swf pageUrl=http://livetvchannelsfree.com/tamilbusinesstv.html")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Santhora TV")
	fmt.Fprintln(w, "http://santhoratv.zecast.net/santhoratv/santhoratv/playlist.m3u8")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Test 01")
	fmt.Fprintln(w, "http://mobile.amgiptv.com/vxtvlqkalaingertv/kalaingertvlq.stream/playlist.m3u8")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Test 02")
	fmt.Fprintln(w, "http://play-fs.amgiptv.com/vxtvmegatv/megatv.stream/playlist.m3u8")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil TV,Test 03")
	fmt.Fprintln(w, "http://50.7.1.178/india_iptv/sunmusicINDIA/playlist.m3u8")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil Radio,Vanavil FM (R)")
	fmt.Fprintln(w, "http://s1.yesstreaming.net:9000/:stream/1")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil Radio,Mukil FM (R)")
	fmt.Fprintln(w, "http://live64.jiljilradio.com/;stream1.mp3")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil Radio,Kalasam FM (R)")
	fmt.Fprintln(w, "http://live.kalasam.com:8084")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil Radio,LankaSri FM (R)")
	fmt.Fprintln(w, "http://media2.lankasri.fm")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil Radio,Oli 96.8 FM (R)")
	fmt.Fprintln(w, "http://mediacorp.rastream.com/968fm?type=.flv")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil Radio,Oli 96.8 FM (R)")
	fmt.Fprintln(w, "http://mediacorp.rastream.com/968fm?type=.flv")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil Radio,THR Raaga (R)")
	fmt.Fprintln(w, "http://astro3.rastream.com:80/raaga?type=.flv")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil Radio,ETR FM (R)")
	fmt.Fprintln(w, "http://cast2.serverhostingcenter.com:8646/stream")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "#EXTINF:-1 group-title=Tamil Radio,Puradsi FM (R)")
	fmt.Fprintln(w, "http://puradsifm.net:9994")
	log.Println("HomeHandler Ending")
}

func main() {
	// defer log.Flush()
	log.Println("App Started")

	f, err := os.Open("final.m3u")
	if err != nil {
		panic(err)
	}
	//	var mw = MediaPlaylist{SeqNo: 1, Title: "Abc", Urls: map[string]string{"u1": "u100", "u2": "u200"}}
	//	fmt.Println(mw)
	mw1, err := DecodeFrom(bufio.NewReader(f), true)
	if err != nil {
		panic(err)
	}
	fmt.Println("XXXXXXXXXXX")
	fmt.Println(mw1)

	// router := mux.NewRouter()
	// router.HandleFunc("/", HomeHandler)
	// router.HandleFunc("/db", DBHandler)
	// // Bind to a port and pass our router in
	// http.ListenAndServe(":8000", nil)

	//	var port string
	//	if port = os.Getenv("PORT"); len(port) == 0 {
	//		log.Println("Warning, PORT not set. Defaulting to %+vn", DEFAULT_PORT)
	//		port = DEFAULT_PORT
	//	}
	//	//	port = DEFAULT_PORT

	//	http.HandleFunc("/", HomeHandler)
	//	http.HandleFunc("/v1", V1Handler)

	//	err = http.ListenAndServe(":"+port, nil)
	//	if err != nil {
	//		log.Println("ListenAndServe: ", err)
	//	}

}

type TVChannel struct {
	SeqNo        int
	Title        string
	MainUrl      string
	Urls         map[string]string
	Groupdetails map[string]string
}

type MediaPlaylist struct {
	Items []TVChannel
}

func (box *MediaPlaylist) AddItem(item TVChannel) []TVChannel {
	box.Items = append(box.Items, item)
	return box.Items
}

// Detect playlist type and decode it from input stream.
func DecodeFrom(reader io.Reader, strict bool) ([]TVChannel, error) {
	//mw := TVChannels
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(reader)
	if err != nil {
		return nil, err
	}
	return decode(buf, strict)
}

// Detect playlist type and decode it. May be used as decoder for both master and media playlists.
func decode(buf *bytes.Buffer, strict bool) ([]TVChannel, error) {
	var eof bool
	var line string
	var err error

	//	chn1 := TVChannel{SeqNo: 1, Title: "Abc", Urls: map[string]string{"u1": "u100", "u2": "u200"}}
	//	chn2 := TVChannel{SeqNo: 2, Title: "Def", Urls: map[string]string{"u1": "u100", "u2": "u200"}}
	//	items := []TVChannel{}
	//	box := MediaPlaylist{items}
	//	box.AddItem(chn1)
	//	box.AddItem(chn2)
	//	fmt.Println(box.Items)
	//	return box.Items, err

	items := []TVChannel{}
	box := MediaPlaylist{items}

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

		err = decodeLineOfMediaPlaylist(&box, line, strict)
		if strict && err != nil {
			return box.Items, err
		}
	}
	return box.Items, nil
}

func decodeLineOfMediaPlaylist(p *MediaPlaylist, line string, strict bool) error {
	var title string
	var err error
	var m3u bool
	var duration float64
	title = "a"
	m3u = true
	duration = 1

	reGroups, _ := regexp.Compile(`(?P<groupKey>\w+-\w+\s*)=(?P<groupValue>\s*\"*\w+\s*\.*\w+\"*)`)           // matches group-title="Tamil TV" group-img="Tamil.jpg"
	reUrls, _ := regexp.Compile(`(\w+)=((([\w-]+://?)?|www[.])[^\s()<>]+(?:\([\w\d]+\)|([^[:punct:]\s]|/)))`) // matches rtmp://38.96.148.172:1935/live/ playpath=imayamtv swfUrl=http://live.akamain.info/v1/imayamtv/jwplayer/jwplayer.flash.swf pageUrl=http://live.akamain.info/v1/imayamtv/index.html live=1
	line = strings.TrimSpace(line)
	switch {
	// start tag first
	case line == "#EXTM3U":
		m3u = true
		//	case !state.tagInf && strings.HasPrefix(line, "#EXTINF:"):
	case strings.HasPrefix(line, "#EXTINF:"):
		params := strings.Split(line[11:], ",")
		if len(params) > 0 {
			//last is Title
			title = params[len(params)-1]
		}
		rMatchedGroups := reGroups.FindAllStringSubmatch(params[0], -1)

		md := map[string]string{}
		for j, _ := range rMatchedGroups {
			md[rMatchedGroups[j][1]] = rMatchedGroups[j][2]
		}
		//chn1 := TVChannel{SeqNo: len(p.Items) + 1, Title: title, Groupdetails: md, Urls: map[string]string{"u1": "u100", "u2": "u200"}}
		chn1 := TVChannel{SeqNo: len(p.Items) + 1, Title: title, Groupdetails: md}
		p.AddItem(chn1)
		//fmt.Println(md)

	case !strings.HasPrefix(line, "#"):
		if len(line) > 0 {
			// There could be multiple urls like
			// rtmp://38.96.148.172:1935/live/ playpath=imayamtv swfUrl=http://live.akamain.info/v1/imayamtv/jwplayer/jwplayer.flash.swf pageUrl=http://live.akamain.info/v1/imayamtv/index.html live=1
			urls := strings.SplitN(line, " ", 2)
			fmt.Println(len(urls))

			if len(urls) > 0 && len(urls[0]) > 0 {
				chn1 := &p.Items[len(p.Items)-1] // Get the last item
				chn1.MainUrl = urls[0]
				if len(urls) > 1 {
					fmt.Println("more than 1")
					fmt.Println(urls[1])
					rUrls := reUrls.FindAllStringSubmatch(urls[1], -1)
					mU := map[string]string{}
					for j, _ := range rUrls {
						mU[rUrls[j][1]] = rUrls[j][2]
					}
					chn1.Urls = mU
				}
			}
		}
		//		if state.tagInf {
		//			p.Append(line, state.duration, title)
		//			state.tagInf = false
		//		}
		//		if state.tagRange {
		//			if err = p.SetRange(state.limit, state.offset); strict && err != nil {
		//				return err
		//			}
		//			state.tagRange = false
		//		}
		//		if state.tagSCTE35 {
		//			state.tagSCTE35 = false
		//			scte := *state.scte
		//			if err = p.SetSCTE(scte.Cue, scte.ID, scte.Time); strict && err != nil {
		//				return err
		//			}
		//		}
		//		if state.tagDiscontinuity {
		//			state.tagDiscontinuity = false
		//			if err = p.SetDiscontinuity(); strict && err != nil {
		//				return err
		//			}
		//		}
		//		if state.tagProgramDateTime {
		//			state.tagProgramDateTime = false
		//			if err = p.SetProgramDateTime(state.programDateTime); strict && err != nil {
		//				return err
		//			}
		//		}
		//		// If EXT-X-KEY appeared before reference to segment (EXTINF) then it linked to this segment
		//		if state.tagKey {
		//			p.Segments[p.last()].Key = &Key{state.xkey.Method, state.xkey.URI, state.xkey.IV, state.xkey.Keyformat, state.xkey.Keyformatversions}
		//			// First EXT-X-KEY may appeared in the header of the playlist and linked to first segment
		//			// but for convenient playlist generation it also linked as default playlist key
		//			if p.Key == nil {
		//				p.Key = state.xkey
		//			}
		//			state.tagKey = false
		//		}
		//		// If EXT-X-MAP appeared before reference to segment (EXTINF) then it linked to this segment
		//		if state.tagMap {
		//			p.Segments[p.last()].Map = &Map{state.xmap.URI, state.xmap.Limit, state.xmap.Offset}
		//			// First EXT-X-MAP may appeared in the header of the playlist and linked to first segment
		//			// but for convenient playlist generation it also linked as default playlist map
		//			if p.Map == nil {
		//				p.Map = state.xmap
		//			}
		//			state.tagMap = false
		//		}
	case strings.HasPrefix(line, "#"): // unknown tags treated as comments
		return err
	}
	title = title + ""
	m3u = !m3u
	duration = duration + 0
	return err
}
