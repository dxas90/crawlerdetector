package crawlerdetector

import (
	"regexp"
	"sync"
)

// Compiled regex cache with lazy initialization
var (
	crawlersRegex     *regexp.Regexp
	mobilesRegex      *regexp.Regexp
	exclusionsRegex   *regexp.Regexp
	piwikCrawlersRx   *regexp.Regexp
	piwikMobilesRx    *regexp.Regexp
	shortCrawlersRx   *regexp.Regexp
	shortMobilesRx    *regexp.Regexp
	shortExclusionsRx *regexp.Regexp

	compileOnce      sync.Once
	piwikCompileOnce sync.Once
	shortCompileOnce sync.Once
)

// Piwik struct to parse the yml
type Piwik struct {
	Name  string `json:"name"`
	Regex string `json:"regex"`
}

// String function to dump the Regex
func (m *Piwik) String() string {
	return m.Regex
}

// CrawlerDetector is crawler detector structure
type CrawlerDetector struct {
	Crawlers   *regexp.Regexp
	Mobiles    *regexp.Regexp
	Exclusions *regexp.Regexp
	Matched    []string
	Browser    bool
	Crawler    bool
	Mobile     bool
}

// New returns a new initialized CrawlerDetector with full pattern lists.
// Uses lazy compilation for optimal performance.
func New() *CrawlerDetector {
	compileOnce.Do(func() {
		crawlersRegex = regexp.MustCompile(CrawlersList())
		mobilesRegex = regexp.MustCompile(MobilesList())
		exclusionsRegex = regexp.MustCompile(ExclusionsList())
	})

	return &CrawlerDetector{
		Crawlers:   crawlersRegex,
		Mobiles:    mobilesRegex,
		Exclusions: exclusionsRegex,
		Matched:    make([]string, 0),
	}
}

// NewPiwik returns a new initialized CrawlerDetector from Piwik patterns.
// Uses Piwik bot/mobile lists with short exclusions.
func NewPiwik() *CrawlerDetector {
	piwikCompileOnce.Do(func() {
		piwikCrawlersRx = regexp.MustCompile(PiwikCrawlersList())
		piwikMobilesRx = regexp.MustCompile(PiwikMobilesList())
		if shortExclusionsRx == nil {
			shortExclusionsRx = regexp.MustCompile(ShortExclusionsList())
		}
	})

	return &CrawlerDetector{
		Crawlers:   piwikCrawlersRx,
		Mobiles:    piwikMobilesRx,
		Exclusions: shortExclusionsRx,
		Matched:    make([]string, 0),
	}
}

// NewShort returns a new basic initialized CrawlerDetector with minimal patterns.
// Fastest initialization and detection for performance-critical applications.
func NewShort() *CrawlerDetector {
	shortCompileOnce.Do(func() {
		shortCrawlersRx = regexp.MustCompile(ShortCrawlersList())
		shortMobilesRx = regexp.MustCompile(ShortMobilesList())
		shortExclusionsRx = regexp.MustCompile(ShortExclusionsList())
	})

	return &CrawlerDetector{
		Crawlers:   shortCrawlersRx,
		Mobiles:    shortMobilesRx,
		Exclusions: shortExclusionsRx,
		Matched:    make([]string, 0),
	}
}

// Parse performs all detection operations concurrently on the user agent.
// This method is thread-safe and can be called concurrently.
func (cd *CrawlerDetector) Parse(userAgent string) *CrawlerDetector {
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		cd.Browser = (len(cd.Exclusions.ReplaceAllString(userAgent, "")) == 0)
	}()

	go func() {
		defer wg.Done()
		matches := cd.Mobiles.FindAllString(userAgent, -1)
		cd.Mobile = len(matches) != 0
	}()

	go func() {
		defer wg.Done()
		matches := cd.Crawlers.FindAllString(userAgent, -1)
		cd.Crawler = len(matches) != 0
		if cd.Crawler {
			cd.Matched = matches
		}
	}()

	wg.Wait()
	return cd
}

// ParseUnsafe performs browser and mobile detection concurrently.
// Unlike Parse(), this assumes any non-browser is a crawler.
// This method is thread-safe and can be called concurrently.
func (cd *CrawlerDetector) ParseUnsafe(userAgent string) *CrawlerDetector {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		cd.IsCrawler(userAgent)
	}()

	go func() {
		defer wg.Done()
		cd.IsMobile(userAgent)
	}()

	wg.Wait()
	return cd
}

// IsCrawler detects crawlers/spiders/bots by user agent.
// Returns false if the user agent matches browser exclusion patterns.
// Populates Matched field with detected crawler patterns.
func (cd *CrawlerDetector) IsCrawler(userAgent string) bool {
	// Browser exclusion check
	if cd.IsExclusion(userAgent) {
		cd.Crawler = false
		cd.Matched = nil
		return false
	}

	// Find crawler patterns
	cd.Matched = cd.Crawlers.FindAllString(userAgent, -1)
	cd.Crawler = len(cd.Matched) > 0

	return cd.Crawler
}

// IsMobile detects mobile devices by user agent.
// Populates Matched field with detected mobile patterns.
func (cd *CrawlerDetector) IsMobile(userAgent string) bool {
	cd.Matched = cd.Mobiles.FindAllString(userAgent, -1)
	cd.Mobile = len(cd.Matched) > 0
	return cd.Mobile
}

// IsExclusion checks if the user agent matches browser exclusion patterns.
// Returns true if the entire user agent consists of exclusion patterns (is a browser).
func (cd *CrawlerDetector) IsExclusion(userAgent string) bool {
	// If removing all exclusion patterns leaves nothing, it's a browser
	remaining := cd.Exclusions.ReplaceAllString(userAgent, "")
	cd.Browser = len(remaining) == 0
	return cd.Browser
}

// GetMatched is getter of matched result
func (cd *CrawlerDetector) GetMatched() []string {
	return cd.Matched
}
