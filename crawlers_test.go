package crawlerdetector

import (
	"log"
	"os"
	"regexp"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestShortCrawlersList(t *testing.T) {
	tester := regexp.MustCompile(ShortCrawlersList())
	testCrawlers := []string{
		"007ac9 Crawler",
		"192\\.comAgent",
		"360Spider",
		"360Spider(-Image|-Video)?",
		"A6-Indexer",
		"(A6-Indexer|nuhk|TsolCrawler|Yammybot|Openbot|Gulper Web Bot|grub-client|Download Demon|SearchExpress|Microsoft URL Control|borg|altavista|dataminr.com|tweetedtimes.com|TrendsmapResolver|teoma|blitzbot|oegp|furlbot|http%20client|polybot|htdig|mogimogi|larbin|scrubby|searchsight|seekbot|semanticdiscovery|snappy|vortex(Build)?|zeal|fast-webcrawler|converacrawler|dataparksearch|findlinks|BrowserMob|HttpMonitor|ThumbShotsBot|URL2PNG|ZooShot|GomezA|Google SketchUp|Read%20Later|Minimo|RackspaceBot)",
		"acapbot",
		"Accoona-AI-Agent",
		"acoon",
		"acoonbot",
		"AcoonBot",
		"acrylicapps\\.com\\/pulp",
		"Acunetix",
		"AdAuth\\/",
		"adbeat",
		"adbeat_bot",
		"AddThis",
		"AddThis\\.com",
		"Adidxbot",
		"ADmantX",
		"AdMantX.*admantx\\.com",
		"adressendeutschland",
		"AdsBot-Google([^-]|$)",
		"AdsBot-Google-Mobile",
		"adscanner",
		"adscanner\\/",
		"AdsTxtCrawler",
		"Advanced Email Extractor v",
		"AdvBot",
		"agentslug",
		"AHC",
		"AHC\\/",
		"AhrefsBot",
		"Ahrefs(Bot|SiteAudit)",
		"aihit",
		"aiHitBot",
		"aiohttp\\/",
		"Airmail",
		"AISearchBot",
		"akka-http\\/",
		"akula\\/",
		"alertra",
		"alexa site audit",
		"Alibaba\\.Security\\.Heimdall",
		"Alligator",
		"allloadin\\.com",
		"AllSubmitter",
		"AlphaBot",
		"alyze\\.info",
		"amagit",
		"Amazon CloudFront",
		"AmorankSpider",
		"Anarchie",
		"AndersPinkBot",
		"AndroidDownloadManager",
		"Anemone",
		"AngleSharp\\/",
		"Ant\\.com",
		"antibot",
		"Anturis Agent",
		"AnyEvent-HTTP\\/",
		"ApacheBench",
		"ApacheBench\\/",
		"Apache Droid",
		"Apache-HttpAsyncClient\\/",
		"Apache-HttpClient",
		"Apache-HttpClient\\/",
		"Apercite",
		"Apexoo",
		"APIs-Google",
		"AportWorm\\/[0-9]",
		"AppBeat\\/[0-9]",
		"AppEngine-Google",
		"Applebot",
		"AppStoreScraperZ",
		"a\\.pr-cy\\.ru",
		"arabot",
		"Arachmo",
		"Arachni",
		"arachnode",
		"Arachnophilia",
		"ArchiveBot",
		"archive.org_bot",
		"archive\\.org_bot|special_archiver",
		"aria2",
		"Arukereso",
		"asafaweb.com",
		"Ask Jeeves/Teoma",
		"AskQuickly",
		"ASPSeek",
		"Asterias",
		"Astute",
		"asynchttp",
		"Attach",
		"autocite",
		"Autonomy",
		"axios",
		"axios\\/",
		"[a-z0-9\\-_]*(bot|crawler|archiver|transcoder|spider|uptime|validator|fetcher)",
		"([A-z0-9]*)-Lighthouse",
		"^b0t$",
		"Backlink-Ceck",
		"backlink-check",
		"Backlink-Check\\.de",
		"backlinkcrawler",
		"BacklinkCrawler",
		"BackStreet",
		"BackWeb",
		"Badass",
		"Bad-Neighborhood",
		"baidu\\.com",
		"Baiduspider",
		"baiduspider(-image)?|baidu Transcoder|baidu.*spider",
		"Baidu-YunGuanCe",
		"Bandit",
		"Barkrowler",
		"BatchFTP",
		"Battleztar\\ Bazinga",
		"baypup\\/[0-9]",
		"baypup\\/colbert",
		"BazQux",
		"BBBike",
		"BCKLINKS",
		"BDCbot",
		"BDFetch",
		"BegunAdvertising\\/",
		"BehloolBot",
		"betaBot",
		"bibnum.bnf",
		"BigBozz",
		"Bigfoot",
		"biglotron",
		"BIGLOTRON",
		"bingbot",
		"BingLocalSearch",
		"BingPreview",
		"BingPreview\\/",
		"binlar",
		"biNu image cacher",
		"Bitacle",
		"bitlybot",
		"biz_Directory",
		"Blackboard Safeassign",
		"Black\\ Hole",
		"BlackWidow",
		"blekkobot",
		"Blekkobot",
		"BLEXBot\\/",
		"BLEXBot(Test)?",
		"B-l-i-t-z-B-O-T",
		"Bloglovin",
		"blogmuraBot",
		"BlogPulseLive",
		"BlogSearch",
		"Blogtrottr",
		"BlowFish",
		"BLP_bbot",
		"^bluefish ",
		"bnf.fr_bot",
		"Boardreader",
		"boitho\\.com-dc",
		"BomboraBot",
		"Bot.AraTurka.com",
		"botify",
		"bot-pge.chlooe.com",
		"BountiiBot",
		"BoxcarBot",
		"BPImageWalker",
		"brainobot",
		"Braintree-Webhooks",
		"Branch Metrics API",
		"Branch-Passthrough",
		"Brandprotect",
		"BrandVerity",
		"Brandwatch",
		"Brodie\\/",
		"Browsershots",
		"BTWebClient",
		"BUbiNG",
		"Buck\\/",
		"Buddy",
		"BuiltWith",
		"Bullseye",
		"BunnySlippers",
		"Burf Search",
		"Butterfly\\/",
		"buzzbot",
		"BuzzSumo",
		"CAAM\\/[0-9]",
		"CakePHP",
		"Calculon",
		"^Calypso v\\/",
		"CapsuleChecker",
		"careerbot",
		"CareerBot",
		"CaretNail",
		"Castro 2, Episode Duration Lookup",
		"Catchpoint( bot)?",
		"catexplorador",
		"cb crawl",
		"CCBot",
		"CC Metadata Scaper",
		"Cegbfeieh",
		"centurybot9",
		"Cerberian Drtrs",
		"CERT\\.at-Statistics-Survey",
		"cg-eye",
		"changedetection",
		"ChangesMeter\\/",
		"charlotte",
		"Charlotte",
		"CheckHost",
		"check_http",
		"check_http/v",
		"checkprivacy",
		"CherryPicker",
		"ChinaClaw",
		"Chirp\\/[0-9]",
		"chkme\\.com",
		"Chlooe",
		"CirrusExplorer\\/",
		"CISPA Vulnerability Notification",
		"citeseerxbot",
		"Citoid",
		"CJNetworkQuality",
		"Clarsentia",
		"clips\\.ua\\.ac\\.be",
		"Cliqzbot",
		"Cliqzbot\\/",
		"CloudEndure",
		"CloudFlare-AlwaysOnline",
		"Cloudflare-AMP",
		"Cloudinary\\/[0-9]",
		"Cloud\\ mapping",
		"cmcm\\.com",
		"coccoc/",
		"coccoc",
		"cognitiveseo",
		"collectd",
		"collection@infegy.com",
		"colly -",
		"CommaFeed",
		"Commons-HttpClient",
		"^COMODO DCV",
		"Comodo SSL Checker",
		"Companybook-Crawler",
		"contactbigdatafr",
		"content crawler spider",
		"contentkingapp",
		"ContextAd Bot",
		"contxbot",
		"convera",
		"CookieReports\\.com",
		"CopyRightCheck",
		"copyright sheriff",
		"Copyscape",
		"Cosmos4j\\.feedback",
		"Covario-IDS",
		"crawler4j",
		"crawler|crawl|checker|archiver|transcoder|spider",
		"CrawlForMe\\/[0-9]",
		"Crescent",
		"cron-job\\.org",
		"Crowsnest",
		"CrunchBot",
		"CrystalSemanticsBot",
		"CSHttp",
		"CSS Certificate Spider",
		"curb",
		"Curious George",
		"curl",
		"cuwhois\\/[0-9]",
		"cXensebot",
		"CyberPatrol",
		"cybo\\.com",
		"^DangDang",
		"DareBoost",
		"DatabaseDriverMysqli",
		"DataCha0s",
		"Datadog Agent",
		"Datafeedwatch",
		"datagnionbot",
		"DataparkSearch",
		"dataprovider",
		"Dataprovider",
		"DataXu",
		"Daum\\/",
		"Daum(oa)?[ /][0-9]",
		"Daum(oa)?[ \\/][0-9]",
		"^DavClnt",
		"Dazoobot",
		"dcrawl",
		"deadlinkchecker",
		"Demon",
		"DeuSu",
		"DeuSu\\/",
		"developers\\.google\\.com\\/\\+\\/web\\/snippet\\/",
		"Devil",
		"Diffbot\\/",
		"Digg",
		"Digg Deeper",
		"Digincore",
		"Digincore bot",
		"DigitalPebble",
		"Dirbuster",
		"discobot",
		"discobot(-news)?",
		"Discordbot",
		"Dispatch\\/",
		"Disqus",
		"DittoSpyder",
		"dlvr",
		"DMBrowser",
		"DNSPod-reporting",
		"DNS-Tools Header-Analyzer",
		"DnyzBot",
		"docoloc",
		"Dolphin http client\\/",
		"DomainAppender",
		"domaincrawler",
		"Domain Re-Animator Bot",
		"Domain Re-Animator Bot|support@domainreanimator.com",
		"DomainStatsBot",
		"Donuts Content Explorer",
		"dotbot",
		"DotBot",
		"dotMailer content retrieval",
		"dotSemantic",
		"downforeveryoneorjustme",
		"Download\\ Wonder",
		"downnotifier\\.com",
		"DowntimeDetector",
		"Dragonfly File Reader",
		"Drip",
		"drupact",
		"Drupal \\(\\+http:\\/\\/drupal\\.org\\/\\)",
		"DTS\\ Agent",
		"dubaiindex",
		"DuckDuck",
		"DuckDuckBot",
		"DuckDuckGo-Favicons-Bot",
		"EARTHCOM",
		"EasouSpider",
		"EasyDL",
		"Easy-Thumb",
		"Ebingbong",
		"ec2linkfinder",
		"eCairn-Grabber",
		"eCatch",
		"ECCP",
		"echocrawl",
		"eContext\\/",
		"Ecxi",
		"edisterbot",
		"EirGrabber",
		"electricmonk",
		"ElectricMonk",
		"elefent",
		"elisabot",
		"Email%20Extractor%20Lite",
		"EMail Exractor",
		"EmailWolf",
		"EMail\\ Wolf",
		"Embedly",
		"Embed PHP Library",
		"epicbot",
		"eright",
		"europarchive.org",
		"europarchive\\.org",
		"evc-batch",
		"EventMachine HttpClient",
		"Everwall Link Expander",
		"EveryoneSocialBot",
		"Evidon",
		"Evrinid",
		"exabot",
		"Exabot(-Thumbnails|-Images)?|ExaleadCloudview",
		"ExactSearch",
		"ExactSeek Crawler",
		"ExaleadCloudview",
		"Excel\\/",
		"Exif Viewer",
		"ExperianCrawlUK",
		"Exploratodo",
		"Express WebPictures",
		"ExtLinksBot",
		"ExtractorPro",
		"Extreme\\ Picture\\ Finder",
		"EyeNetIE",
		"EZID",
		"ezooms",
		"Ezooms",
		"facebookexternalhit",
		"facebookexternalhit|facebookplatform",
		"facebookplatform",
		"Facebot",
		"fairshare",
		"Faraday v",
		"FAST Enterprise Crawler",
		"fasthttp",
		"FAST-WebCrawler",
		"Faveeo",
		"Favicon downloader",
		"FavOrg",
		"^FDM ",
		"Feedbin",
		"FeedBooster",
		"FeedBucket",
		"FeedBunch\\/[0-9]",
		"FeedBurner",
		"FeedChecker",
		"Feedfetcher-Google",
		"Feedly",
		"Feedspot",
		"Feedspotbot\\/",
		"Feedwind\\/[0-9]",
		"Feed Wrangler",
		"feeltiptop",
		"FemtosearchBot",
		"Fetch\\/",
		"Fetch\\/[0-9]",
		"Fetch API",
		"Fever",
		"Fever/[0-9]",
		"Fever\\/[0-9]",
		"FHscan",
		"filterdb.iss.net\\/crawler",
		"Fimap",
		"findlink",
		"findthatfile",
		"findxbot",
		"Findxbot",
		"Flamingo_SearchEngine",
		"FlashGet",
		"FlipboardBrowserProxy",
		"FlipboardProxy",
		"FlipboardProxy|FlipboardRSS",
		"FlipboardRSS",
		"Flock\\/",
		"fluffy",
		"Flunky",
		"flynxapp",
		"forensiq",
		"FoundSeoTool\\/[0-9]",
		"fr-crawler",
		"free thumbnails",
		"Freeuploader",
		"FreeWebMonitoring SiteChecker",
		"fuelbot",
		"Funnelback",
		"Fyrebot",
		"g00g1e.net",
		"g00g1e\\.net",
		"g2reader-bot",
		"GAChecker",
		"ganarvisitas\\/[0-9]",
		"GarlikCrawler",
		"geek-tools",
		"Genderanalyzer",
		"Genieo",
		"GentleSource",
		"Getintent",
		"GetLinkInfo",
		"getprismatic\\.com",
		"GetRight",
		"GetURLInfo\\/[0-9]",
		"GetWeb",
		"Ghost Inspector",
		"Gigablast",
		"GigablastOpenSource",
		"gigabot",
		"G-i-g-a-b-o-t",
		"Gigabot",
		"GingerCrawler",
		"GIS-LABS",
		"^git\\/",
		"github-camo",
		"github\\.com\\/",
		"Gluten Free Crawler",
		"Gluten Free Crawler\\/",
		"gnam gnam spider",
		"GnowitNewsbot",
		"Go-Ahead-Got-It",
		"gobyus",
		"Go [\\d\\.]* package http",
		"gofetch",
		"Go-http-client",
		"Go http package",
		"GomezAgent",
		"gooblog",
		"Goodzer\\/[0-9]",
		"Google-Adwords",
		"Google-Adwords-Instant",
		"Google-Apps-Script",
		"Googlebot\\/",
		"Googlebot-Image",
		"Googlebot-Mobile",
		"Googlebot-News",
		"Googlebot-Video",
		"Google-Calendar-Importer",
		"GoogleCloudMonitoring",
		"GoogleDocs",
		"Google favicon",
		"GoogleHC\\/",
		"Google-HotelAdsVerifier",
		"Google-HTTP-Java-Client",
		"Google Keyword Suggestion",
		"Google Keyword Tool",
		"Google Page Speed Insights",
		"google_partner_monitoring",
		"Google PP Default",
		"GoogleProducer",
		"Google-Publisher-Plugin",
		"Google-SearchByImage",
		"Google Search Console",
		"Google-Site-Verification",
		"Google-Structured-Data-Testing-Tool",
		"Google Web Preview",
		"google-xrawler",
		"Google-Youtube-Links",
		"Gookey",
		"^Goose\\/",
		"GoScraper",
		"GoSpotCheck",
		"GoSquared-Status-Checker",
		"gosquared-thumbnailer",
		"Gotit",
		"Gowikibot",
		"Go!Zilla",
		"GoZilla",
		"^Grabber",
		"grabify",
		"GrabNet",
		"Grafula",
		"Grammarly",
		"GrapeFX",
		"GrapeshotCrawler",
		"Grobbot",
		"grokkit",
		"grouphigh",
		"grub-client",
		"grub.org",
		"gslfbot",
		"gSOAP\\/",
		"GTmetrix",
		"GT::WWW",
		"GuzzleHttp",
		"gvfs\\/",
		"HAA(A)?RTLAND http client",
		"Haansoft",
		"hackney\\/",
		"Hatena",
		"Havij",
		"hawkReader",
		"HEADMasterSEO",
		"HeartRails_Capture",
		"help@dataminr\\.com",
		"heritrix",
		"Heurekabot-Feed",
		"historious\\/",
		"hledejLevne\\.cz\\/[0-9]",
		"Hloader",
		"HMView",
		"Holmes",
		"HonesoSearchEngine\\/",
		"HootSuite Image proxy",
		"Hootsuite-WebFeed\\/[0-9]",
		"hosterstats",
		"HostTracker",
		"(<?HTC)[ _]Butterfly/",
		"ht:\\/\\/check",
		"htdig",
		"HTMLparser",
		"^HTTPClient\\/",
		"HttpComponents",
		"HTTP_Compression_Test",
		"http_get",
		"http-get",
		"HTTP-Header-Abfrage",
		"httphr",
		"http-kit",
		"HTTP::Lite",
		"HTTPMon",
		"http\\.rb\\/",
		"http-request\\/",
		"http_request2",
		"http_requester",
		"httpscheck",
		"httpssites_power",
		"HTTP-Tiny",
		"httpunit",
		"HttpUrlConnection",
		"httrack",
		"HTTrack",
		"huaweisymantec",
		"HubPages.*crawlingpolicy",
		"HubSpot ",
		"Humanlinks",
		"HyperZbozi.cz Feeder",
		"i2kconnect\\/",
		"ia_archiver",
		"ia_archiver|alexabot|verifybot",
		"IAS crawler",
		"Iblog",
		"ICC-Crawler",
		"ichiro",
		"ichiro/mobile goo",
		"IdeelaborPlagiaat",
		"IDG Twitter Links Resolver",
		"Id-search",
		"IDwhois\\/[0-9]",
		"Iframely",
		"igdeSpyder",
		"iisbot",
		"IlTrovatore",
		"ImageEngine\\/",
		"Image\\ Fetch",
		"Image\\ Sucker",
		"Imagga",
		"imgsizer",
		"imrbot",
		"InAGist",
		"inbound\\.li parser",
		"IndeedBot",
		"InDesign%20CC",
		"Indy\\ Library",
		"infegy",
		"infohelfer",
		"InfoTekies",
		"InfoWizards Reciprocal Link System PRO",
		"inpwrd\\.com",
		"instabid",
		"Instapaper",
		"Integrity",
		"integromedb",
		"intelium_bot",
		"Intelliseek",
		"InterfaxScanBot",
		"InterGET",
		"internet_archive",
		"Internet\\ Ninja",
		"InternetSeer",
		"internetVista monitor",
		"intraVnews",
		"IODC",
		"IOI",
		"IP-Guide\\.com",
		"iplabel",
		"IPS\\/[0-9]",
		"ips-agent",
		"ip-web-crawler.com",
		"IPWorks HTTP\\/S Component",
		"iqdb\\/",
		"Iria",
		"Irokez",
		"isitup\\.org",
		"iskanie",
		"IstellaBot",
		"it2media-domain-crawler",
		"iZSearch",
		"James BOT",
		"Jamie's Spider",
		"janforman",
		"Jaunt\\/",
		"^Java\\/",
		".*Java.*outbrain",
		"Jbrofuzz",
		"^Jeode\\/",
		"Jersey\\/",
		"JetCar",
		"Jetslide",
		"^Jetty\\/",
		"Jetty",
		"Jigsaw",
		"Jobboerse",
		"JobFeed discovery",
		"Jobg8 URL Monitor",
		"jobo",
		"Jobrapido",
		"Jobsearch1\\.5",
		"JoinVision Generic",
		"JolokiaPwn",
		"Joomla",
		"Jorgee",
		"JS-Kit",
		"Jugendschutzprogramm-Crawler",
		"JustView",
		"jyxobot",
		"K7MLWCBot",
		"Kaspersky Lab CFR link resolver",
		"KeepRight OpenStreetMap Checker",
		"Kelny\\/",
		"Kemvibot",
		"Kerrigan\\/",
		"KeyCDN",
		"Keyword\\ Density",
		"Keyword Extractor",
		"Keywords Research",
		"KickFire",
		"KimonoLabs\\/",
		"Kml-Google",
		"knows\\.is",
		"KOCMOHABT",
		"KosmioBot",
		"kouio",
		"kube-probe",
		"kulturarw3",
		"KumKie",
		"Landau-Media-Spider",
		"larbin",
		"Larbin",
		"Lavf\\/",
		"LayeredExtractor",
		"lb-spider",
		"^LCC ",
		"LeechFTP",
		"LeechGet",
		"Leikibot",
		"letsencrypt",
		"Lftp",
		"LibVLC",
		"LibWeb",
		"Libwhisker",
		"libwww",
		"Licorne Image Snapshot",
		"Liferea\\/",
		"Lightspeedsystems",
		"Likse",
		"Linguee Bot",
		"LinkAlarm\\/",
		"linkapediabot",
		"LinkArchiver",
		"linkCheck",
		"link checker",
		"linkdex",
		"linkdexbot(-mobile)?|linkdex\\.com",
		"LinkedInBot",
		"LinkExaminer",
		"linkfluence",
		"linkpeek",
		"LinkPreviewGenerator",
		"LinkScan",
		"LinksManager",
		"link_thumbnailer",
		"LinkTiger",
		"Link Valet",
		"LinkWalker",
		"lipperhey",
		"Lipperhey",
		"Litemage_walker",
		"livedoor ScreenShot",
		"Livelap[bB]ot",
		"LoadImpactRload",
		"LongURL API",
		"looksystems\\.net",
		"lssbot",
		"lssrocketcrawler",
		"ltx71",
		"lua-resty-http",
		"Luminator-robots",
		"L\\.webis",
		"lwp-request",
		"LWP::Simple",
		"lwp-trivial",
		"lycos",
		"LYT\\.SR",
		"mabontland",
		"Mag-Net",
		"magpie-crawler",
		"MagpieRSS",
		"^Mail\\/",
		"MailChimp",
		"Mail.Ru",
		"Mail.RU_Bot",
		"Mail\\.RU(_Bot)?",
		"Majestic12",
		"makecontact\\/",
		"Mandrill",
		"MapperCmd",
		"mappydata",
		"marketinggrader",
		"MarkMonitor",
		"MarkWatch",
		"masscan\\/[0-9]",
		"Mass\\ Downloader",
		"Mastodon/",
		"Mastodon",
		"Mata\\ Hari",
		"MauiBot",
		"meanpathbot",
		"Mediapartners-Google",
		"Mediapartners \\(Googlebot\\)",
		"Mediatoolkitbot",
		"mediawords",
		"MegaIndex",
		"MegaIndex\\.ru",
		"MeltwaterNews",
		"Melvil Rawi\\/",
		"memorybot",
		"MergeFlow-PageReader",
		"(Meta)?Feedly(Bot|App)?",
		"MetaInspector",
		"MetaJobBot",
		"Metaspinner",
		"MetaURI",
		"MFC_Tear_Sample",
		"^Mget",
		"Microsearch",
		"Microsoft\\ Data\\ Access",
		"Microsoft Office ",
		"Microsoft Outlook",
		"^Microsoft URL Control",
		"Microsoft-WebDAV-MiniRedir",
		"Microsoft Windows Network Diagnostics",
		"MIDown\\ tool",
		"MIIxpc",
		"Mindjet",
		"mindUpBot",
		"Miniature.io\\/",
		"Miniflux",
		"Miniflux\\/",
		"Mister\\ PiX",
		"mixdata dot com",
		"mixed-content-scan",
		"mixnode",
		"MixrankBot",
		"MJ12bot",
		"mlbot",
		"Mnogosearch",
		"moatbot",
		"mogimogi",
		"Mojeek",
		"MojeekBot",
		"MojeekBot\\/",
		"Mojolicious \\(Perl\\)",
		"Monit\\/",
		"monitis",
		"Monitority\\/[0-9]",
		"montastic",
		"MonTools",
		"Moreover",
		"Morfeus\\ Fucking\\ Scanner",
		"Morning Paper",
		"MovableType",
		"mowser",
		"Mrcgiguy",
		"MSFrontPage",
		"mShots",
		"msnbot",
		"MSNBot|msrbot|bingbot|BingPreview|msnbot-(UDiscovery|NewsBlogs)|adidxbot",
		"msrbot",
		"MS\\ Web\\ Services\\ Client\\ Protocol",
		"MuckRack",
		"MuckRack\\/",
		"muhstik-scan",
		"Multiviewbot",
		"munin",
		"MVAClient",
		"MxToolbox\\/",
		"nagios",
		"Najdi\\.si\\/",
		"NalezenCzBot",
		"Name\\ Intelligence",
		"Nameprotect",
		"Navroad",
		"NearSite",
		"Needle",
		"NerdByNature.Bot",
		"nerdybot",
		"Nessus",
		"NetAnts",
		"NETCRAFT",
		"NetcraftSurveyAgent",
		"Netcraft( Web Server Survey| SSL Server Survey|SurveyAgent)",
		"netEstate NE Crawler",
		"NetLyzer",
		"NetLyzer FastProbe",
		"NetMechanic",
		"Netpursual",
		"netresearch",
		"netresearchserver",
		"NetShelter ContentScan",
		"Netsparker",
		"NetTrack",
		"Net\\ Vampire",
		"Netvibes",
		"NetZIP",
		"Neustar WPM",
		"NeutrinoAPI",
		"NewRelicPinger",
		"NewsBlur .*(Fetcher|Finder)",
		"NewsBlur .*Finder",
		"NewsGator",
		"NewsGatorOnline",
		"newsharecounts",
		"newsme",
		"newspaper\\/",
		"Nexgate Ruby Client",
		"^NG\\/[0-9\\.]",
		"NG-Search",
		"Nibbler",
		"NICErsPRO",
		"niki-bot",
		"Nikto",
		"Nimbostratus-Bot",
		"nineconnections\\.com",
		"^NING\\/",
		"NING\\/",
		"nlcrawler",
		"NLNZ_IAHarvester",
		"Nmap Scripting Engine",
		"node\\.io",
		"node-superagent",
		"node-urllib\\/",
		"nominet\\.org\\.uk",
		"Norton-Safeweb",
		"Notifixious",
		"notifyninja",
		"nrsbot|netresearch",
		"nuhk",
		"nutch",
		"Nutch",
		"Nuzzel",
		"nWormFeedFinder",
		"Nymesis",
		"NYU",
		"Ocarinabot",
		"Ocelli\\/[0-9]",
		"Octopus",
		"Octopus [0-9]",
		"oegp",
		"Offline Explorer",
		"Offline\\ Navigator",
		"okhttp",
		"Omea Reader",
		"omgili",
		"omgili(?:bot)?",
		"OMSC",
		"Online Domain Tools",
		"OpenCalaisSemanticProxy",
		"Openfind",
		"OpenHoseBot",
		"openindexspider",
		"OpenindexSpider",
		"OpenLinkProfiler",
		"Openstat\\/",
		"OpenVAS",
		"OpenWebSpider",
		"Optimizer",
		"OrangeBot\\/",
		"OrangeBot|VoilaBot",
		"Orbiter",
		"OrgProbe\\/[0-9]",
		"orion-semantics",
		"OutclicksBot",
		"Outlook-Express",
		"Owler",
		"ow\\.ly",
		"ownCloud News",
		"OxfordCloudService\\/[0-9]",
		"page2rss",
		"Page Analyzer",
		"PageAnalyzer",
		"PageGrabber",
		"PagePeeker",
		"page\\ scorer",
		"PageScorer",
		"Pagespeed\\/[0-9]",
		"Page Valet",
		"page_verifier",
		"Panopta",
		"panscient",
		"Papa\\ Foto",
		"PaperLiBot",
		"parsijoo",
		"Pavuk",
		"PayPal IPN",
		"pcBrowser",
		"Pcore-HTTP",
		"PEAR HTTPRequest",
		"Pearltrees",
		"PECL::HTTP",
		"peerindex",
		"Peew",
		"PeoplePal",
		"Perlu -",
		"phantomas/",
		"PhantomJS",
		"PhantomJS\\/",
		"PhantomJS Screenshoter",
		"Photon\\/",
		"^PHP\\/[0-9]",
		"phpcrawl",
		"phpservermon",
		"Picscout",
		"Picsearch",
		"PictureFinder",
		"Pimonster",
		"Pi-Monster",
		"Pingability",
		"ping\\.blo\\.gs\\/",
		"pingdom",
		"Pingdom",
		"Pingdom\\.com",
		"Pingoscope",
		"PingSpot",
		"pinterest",
		"pinterest\\.com",
		"Pinterest/\\d\\.\\d.*www\\.pinterest\\.com.*",
		"PiplBot",
		"Pixray",
		"Pizilla",
		"PleaseCrawl",
		"Ploetz \\+ Zeller",
		"Plukkie",
		"plumanalytics",
		"PocketParser",
		"Pockey",
		"POE-Component-Client-HTTP",
		"Pompos",
		"Porkbun",
		"Port Monitor",
		"postano",
		"PostmanRuntime\\/",
		"PostPost",
		"postrank",
		"PowerPoint\\/",
		"PR-CY.RU",
		"Priceonomics Analysis Engine",
		"Primalbot",
		"PrintFriendly\\.com",
		"PritTorrent",
		"PritTorrent\\/[0-9]",
		"PrivacyAwareBot",
		"Prlog",
		"probethenet",
		"Project 25499",
		"Promotion_Tools_www.searchenginepromotionhelp.com",
		"prospectb2b",
		"Protopage",
		"ProWebWalker",
		"proximic",
		"PRTG Network Monitor",
		"psbot",
		"psbot(-page)?",
		"pshtt, https scanning",
		"PTST ",
		"PTST/",
		"PTST\\/[0-9]+",
		"Pulsepoint XT3 web scraper",
		"Pump",
		"purebot",
		"Python-httplib2",
		"python-requests",
		"Python-urllib",
		"Qirina Hurdler",
		"QQDownload",
		"QrafterPro",
		"Qseero",
		"Qualidator.com SiteAnalyzer",
		"QueryN\\ Metasearch",
		"QuerySeekerSpider",
		"Quora Link Preview",
		"Qwantify",
		"Radian6",
		"Rainmeter",
		"RamblerMail",
		"RankActive",
		"RankActiveLinkBot",
		"RankFlex",
		"RankSonicSiteAuditor",
		"Readability",
		"RealDownload",
		"RealPlayer%20Downloader",
		"RebelMouse",
		"Recorder",
		"RecurPost\\/",
		"redback\\/",
		"redditbot",
		"Redirect Checker Tool",
		"ReederForMac",
		"ReGet",
		"RepoMonkey",
		"request\\.js",
		"ResponseCodeTest\\/[0-9]",
		"RestSharp",
		"RetrevoPageAnalyzer",
		"Riddler",
		"Rival IQ",
		"^RMA\\/",
		"Robosourcer",
		"Robozilla\\/[0-9]",
		"rogerbot",
		"ROI Hunter",
		"RPT-HTTPClient",
		"RSSOwl",
		"^Ruby|Ruby\\/[0-9]",
		"safe-agent-scanner",
		"SafeDNSBot",
		"SafeSearch microdata crawler",
		"SalesIntelligent",
		"Saleslift",
		"SauceNAO",
		"SBIder",
		"SBL-BOT",
		"scalaj-http",
		"ScanAlert",
		"scan\\.lol",
		"Scoop",
		"scooter",
		"ScoutJet",
		"ScoutURLMonitor",
		"Scrapy",
		"Screaming",
		"Screaming Frog SEO Spider",
		"ScreenerBot",
		"ScreenShotService\\/[0-9]",
		"scribdbot",
		"Scrubby",
		"^scrutiny\\/",
		"Search37\\/",
		"Searchestate",
		"SearchSight",
		"search\\.thunderstone",
		"seekbot",
		"Seeker",
		"S[eE][mM]rushBot",
		"semanticbot",
		"semanticdiscovery",
		"semanticjuice",
		"SemanticScholarBot",
		"Semiocast HTTP client",
		"Semrush",
		"SemrushBot",
		"s Encrypt validation server",
		"SensikaBot",
		"^sentry",
		"(^| )sentry\\/",
		"sentry\\/",
		"Seobility",
		"SEO Browser",
		"SEOCentro",
		"SeoCheck",
		"SEOENG(World)?Bot",
		"SEOkicks",
		"SEOkicks-Robot",
		"Seomoz",
		"seo-nastroj.cz",
		"SEOprofiler",
		"SeopultContentAnalyzer",
		"seoscanners",
		"seoscanners\\.net",
		"Seo Servis",
		"SEOstats",
		"Server Density Service Monitoring.*",
		"Server Density Service Monitoring",
		"servernfo\\.com",
		"SetCronJob\\/",
		"sexsearcher",
		"Seznam",
		"seznambot",
		"SeznamBot|SklikBot|Seznam screenshot-generator",
		"SeznamEmailProxy",
		"Seznam-Zbozi-robot",
		"Shelob",
		"Shodan",
		"ShopAlike",
		"ShoppimonAgent\\/[0-9]",
		"Shoppimon Analyzer",
		"ShopWiki",
		"ShortLinkTranslate",
		"shrinktheweb",
		"Sideqik",
		"SilverReader",
		"SimpleCrawler",
		"SimplePie",
		"SimplyFast",
		"Siphon",
		"SISTRIX",
		"sistrix crawler",
		"SISTRIX Crawler",
		"Site24x7",
		"SiteBar",
		"Sitebeam",
		"sitebot",
		"Sitebulb\\/",
		"SiteCondor",
		"SiteExplorer",
		"siteexplorer.info",
		"SiteGuardian",
		"Siteimprove",
		"Siteimprove.com",
		"SiteIndexed",
		"Sitemap(s)? Generator",
		"SiteMonitor",
		"Siteshooter B0t",
		"Site-Shot\\/",
		"SiteSnagger",
		"Site\\ Sucker",
		"SiteSucker",
		"SiteTruth",
		"Sitevigil",
		"sitexy\\.com",
		"sixy.ch",
		"SkypeUriPreview",
		"Slack\\/",
		"Slackbot",
		"Slackbot|Slack-ImgProxy",
		"Slack-ImgProxy",
		"slider\\.com",
		"slurp",
		"Slurp",
		"SlySearch",
		"SmartDownload",
		"SMRF URL Expander",
		"smtbot",
		"SMUrlExpander",
		"Snacktory",
		"Snake",
		"Snappy",
		"SniffRSS",
		"sniptracker",
		"Snoopy",
		"SnowHaze Search",
		"SocialRankIOBot",
		"Sogou",
		"sogou web",
		"(Sogou (web|inst|Pic) spider)|New-Sogou-Spider",
		"Sonic",
		"SortSite",
		"Sosospider|Sosoimagespider",
		"Sottopop",
		"sovereign\\.ai",
		"SpaceBison",
		"SpamExperts",
		"Spammen",
		"Spanner",
		"Sparkler/[0-9]",
		"spaziodati",
		"spbot",
		"SPDYCheck",
		"Specificfeeds",
		"speedy",
		"Speedy Spider",
		"SPEng",
		"Spinn3r",
		"spray-can",
		"Sprinklr ",
		"Sputnik(Image)?Bot",
		"spyonweb",
		"sqlmap/",
		"sqlmap",
		"Sqlworm",
		"Sqworm",
		"SSL Labs",
		"ssl-tools",
		"s~snapchat-proxy",
		"StackRambler",
		"Statastico\\/",
		"StatusCake",
		"Steeler",
		"StorygizeBot",
		"Stratagems Kumo",
		"Stroke.cz",
		"StudioFACA",
		"suchen",
		"Sucuri",
		"summify",
		"Superfeedr bot",
		"SuperHTTP",
		"Super Monitoring",
		"Surphace Scout",
		"SurveyBot",
		"Suzuran",
		"SWIMGBot",
		"SwiteScraper",
		"Symfony2 BrowserKit",
		"Symfony BrowserKit",
		"SynHttpClient-Built",
		"Sysomos",
		"sysscan",
		"Szukacz",
		"T0PHackTeam",
		"tagoobot",
		"tAkeOut",
		"TangibleeBot",
		"Tarantula\\/",
		"Taringa UGC",
		"TarmotGezgin",
		"TelegramBot",
		"Teleport",
		"Telesoft",
		"Telesphoreo",
		"Telesphorep",
		"Tenon\\.io",
		"teoma",
		"Teoma",
		"terrainformatica\\.com",
		"Test Certificate Info",
		"Tetrahedron\\/[0-9]",
		"The Drop Reaper",
		"The Expert HTML Source Viewer",
		"theinternetrules",
		"The\\ Intraformant",
		"The Knowledge AI",
		"TheNomad",
		"theoldreader\\.com",
		"Thinklab",
		"Thumbshots",
		"ThumbSniper",
		"TinEye",
		"TinEye-bot",
		"Tiny Tiny RSS",
		"TLSProbe",
		"TLSProbe\\/",
		"Toata",
		"toplistbot",
		"topster",
		"touche.com",
		"ToutiaoSpider",
		"Traackr.com",
		"tracemyfile",
		"TrapitAgent",
		"Trendiction",
		"trendictionbot",
		"Trendsmap Resolver",
		"trendspottr\\.com",
		"Trove",
		"truwoGPS",
		"TulipChain",
		"Turingos",
		"Turnitin",
		"turnitinbot",
		"TurnitinBot",
		"TweetedTimes Bot",
		"tweetedtimes\\.com",
		"TweetmemeBot",
		"Tweetminster",
		"Tweezler\\/",
		"twengabot",
		"twibble",
		"Twice",
		"Twikle",
		"Twingly",
		"Twisted PageGetter",
		"Twitterbot",
		"Twurly",
		"Typhoeus",
		"ubermetrics-technologies",
		"uclassify",
		"uCrawlr\\/",
		"UdmSearch",
		"um-LN",
		"UniversalFeedParser",
		"Unshorten\\.It",
		"Untiny",
		"UnwindFetchor",
		"updated",
		"updown\\.io daemon",
		"Upflow",
		"Uptimebot",
		"UptimeRobot",
		"Uptimia",
		"urlappendbot",
		"URLAppendBot",
		"URLChecker",
		"URLitor.com",
		"urlresolver",
		"Urlstat",
		"UrlTrends Ranking Updater",
		"URL Verifier",
		"URLy\\ Warning",
		"URLy\\.Warning",
		"UsineNouvelleCrawler",
		"Vacuum",
		"Vagabondo",
		"VB\\ Project",
		"vBSEO",
		"VCI",
		"vebidoobot",
		"Veoozbot",
		"via ggpht\\.com GoogleImageProxy",
		"via secureurl\\.fwdcdn\\.com",
		"VidibleScraper",
		"Virusdie",
		"visionutils",
		"vkShare; ",
		"vkShare",
		"VoidEYE",
		"Voil",
		"voilabot",
		"voltron",
		"voyager\\/",
		"VSAgent\\/[0-9]",
		"VSB-TUO\\/[0-9]",
		"^VSE\\/[0-9]",
		"VSMCrawler",
		"Vulnbusters Meter",
		"VYU2",
		"w3af\\.org",
		"W3C-checklink",
		"W3C_I18n-Checker",
		"W3C-mobileOK",
		"W3C_Unicorn",
		"W3C_Validator|Validator.nu",
		"Wallpapers\\/[0-9]+",
		"WallpapersHD",
		"wangling",
		"Wappalyzer",
		"WatchMouse",
		"wbsearchbot",
		"WbSrch\\/",
		"Webalta",
		"web-archive-net.com.bot",
		"Webauskunft",
		"Web\\ Auto",
		"WebAuto",
		"WebbCrawler",
		"WebCapture",
		"web-capture\\.net",
		"WebClient\\/",
		"webcollage",
		"Web\\ Collage",
		"webcompanycrawler",
		"WebCookies",
		"WebCopier",
		"WebCorp",
		"WebDoc",
		"Web\\ Enhancer",
		"WebEnhancer",
		"Web\\ Fetch",
		"WebFetch",
		"Web\\ Fuck",
		"WebFuck",
		"WebGo\\ IS",
		"WebImageCollector",
		"WebImages",
		"WebIndex",
		"webkit2png",
		"WebLeacher",
		"webmastercoffee",
		"webmon ",
		"Web-Monitoring",
		"Web\\ Pix",
		"WebPix",
		"WebReaper",
		"Web\\ Sauger",
		"WebSauger",
		"webscreenie",
		"Webshag",
		"Webshot",
		"Website Analyzer\\/",
		"WebsiteExtractor",
		"websitepulse agent",
		"websitepulse[+ ]checker",
		"Website\\ Quester",
		"WebsiteQuester",
		"Websnapr\\/",
		"Web-sniffer",
		"Webster",
		"WebStripper",
		"Web\\ Sucker",
		"WebSucker",
		"Webthumb\\/[0-9]",
		"WebThumbnail",
		"WebWhacker",
		"WebZIP",
		"WeCrawlForThePeace",
		"WeLikeLinks",
		"WEPA",
		"WeSEE",
		"WeSEE:Search",
		"WeSEE(:Search)?",
		"wf84",
		"Wfuzz\\/",
		"wget",
		"WhatsApp",
		"WhatsMyIP",
		"WhatWeb",
		"WhereGoes\\?",
		"Whibse",
		"WhoRunsCoinHive",
		"Whynder Magnet",
		"Willow Internet Crawler",
		"Windows-RSS-Platform",
		"WinHttpRequest",
		"wkhtmlto",
		"wmtips",
		"wocbot",
		"Woko",
		"woobot",
		"Word\\/",
		"WordPress",
		"WordPress\\/",
		"^WordPress\\.com",
		"woriobot",
		"wotbox",
		"Wotbox",
		"WP Engine Install Performance API",
		"wpif",
		"wprecon\\.com survey",
		"WPScan",
		"wscheck",
		"Wtrace",
		"[wW]get",
		"WWW-Collector-E",
		"WWW-Mechanize",
		"WWW::Mechanize",
		"www\\.monitor\\.us",
		"WWWOFFLE",
		"x09Mozilla",
		"x22Mozilla",
		"XaxisSemanticsClassifier",
		"Xenu Link Sleuth",
		"XING-contenttabreceiver\\/[0-9]",
		"XmlSitemapGenerator",
		"xovibot",
		"xpymep([0-9]?)\\.exe",
		"^XRL\\/[0-9]",
		"Yaanb",
		"yacy",
		"yacybot",
		"Yahoo Ad monitoring",
		"Yahoo Ad monitoring.*yahoo-ad-monitoring-SLN24857.*",
		"YahooCacheSystem",
		"Yahoo Link Preview",
		"Yahoo Link Preview|Yahoo:LinkExpander:Slingstone",
		"Yahoo! Slurp|Yahoo!-AdCrawler",
		"YahooYSMcm",
		"YaK\\/",
		"YandeG",
		"YandexBot",
		"yandex.com\\/bots",
		"Yandex([^Search]*)",
		"Yandex(SpravBot|ScreenshotBot|MobileBot|AccessibilityBot|ForDomain|Vertis|Market|Catalog|Calendar|Sitelinks|AdNet|Pagechecker|Webmaster|Media|Video|Bot|Images|Antivirus|Direct|Blogs|Favicons|ImageResizer|Verticals|News(links)?|Metrika|\\.Gazeta Bot)|YaDirectFetcher",
		"yanga",
		"yeti",
		"Yeti",
		"YisouSpider",
		"Y!J",
		"Y!J-(ASR|BSC)",
		" YLT",
		"Yoleo Consumer",
		"yoogliFetchAgent",
		"yoozBot",
		"YottaaMonitor",
		"YoudaoBot",
		"yourls\\.org",
		"YOURLS v[0-9]",
		"Your-Website-Sucks\\/[0-9]",
		"Yo-yo",
		"YRSpider|YYSpider",
		"Zabbix",
		"Zade",
		"Zao/",
		"Zao",
		"Zauba",
		"Zemanta Aggregator",
		"Zend_Http_Client",
		"Zend\\\\\\\\Http\\\\\\\\Client",
		"Zermelo",
		"Zeus",
		"zgrab",
		"^ZmEu",
		"ZnajdzFoto",
		"Zombie\\.js",
		"Zookabot",
		"ZoominfoBot",
		"ZumBot",
		"ZuperlistBot\\/",
		"ZyBorg",
	}
	tmpfile, err := os.OpenFile("/tmp/aki.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	btmpfile, err := os.OpenFile("/tmp/bki.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	for _, test := range testCrawlers {
		t.Run(test, func(tt *testing.T) {
			response := tester.FindAllString(test, -1)
			if len(response) == 0 {
				if _, err := tmpfile.Write([]byte(test + "\n")); err != nil {
					log.Fatal(err)
				}
			} else {
				if _, err := btmpfile.Write([]byte(test + "\n")); err != nil {
					log.Fatal(err)
				}
			}
			Convey("Subject: Test ShortCrawlersList response\n", t, func() {
				Convey("Should Be a bot", func() {
					So(len(response), ShouldNotEqual, 0)
				})
			})
		},
		)
	}
}
