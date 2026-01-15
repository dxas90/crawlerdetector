# About Crawler Detector

Crawler Detector is a high-performance, thread-safe Golang package for detecting bots/crawlers/spiders and mobile devices via the user agent string.

## Features

- 🚀 **High Performance**: Lazy regex compilation with caching for optimal speed
- 🔒 **Thread-Safe**: Concurrent-safe operations, perfect for high-traffic servers
- 📱 **Mobile Detection**: Comprehensive mobile device pattern recognition
- 🤖 **Bot Detection**: Extensive crawler/spider/bot patterns including Piwik dataset
- ⚙️ **Flexible Modes**: Choose between Full, Piwik, or Short detection for your needs

### Installation

```bash
go get github.com/dxas90/crawlerdetector
```

### Quick Start

```go
package main

import (
    "fmt"
    "github.com/dxas90/crawlerdetector"
)

func main() {
    // Create a detector instance
    detector := crawlerdetector.New()

    // Analyze a user agent
    userAgent := "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"
    detector.Parse(userAgent)

    if detector.Crawler {
        fmt.Println("Bot detected!")
        fmt.Printf("Matched patterns: %v\n", detector.GetMatched())
    }
}
```

### Detection Modes

Choose the detection mode that fits your needs:

#### Full Mode (Most Comprehensive)

```go
detector := crawlerdetector.New()
// Uses complete pattern lists for maximum coverage
// Best for: Production servers requiring thorough bot detection
```

#### Piwik Mode (Standard)

```go
detector := crawlerdetector.NewPiwik()
// Uses Piwik bot and mobile databases
// Best for: Standard bot detection with Piwik compatibility
```

#### Short Mode (Fastest)

```go
detector := crawlerdetector.NewShort()
// Uses minimal pattern sets for quick detection
// Best for: High-performance APIs, rate limiting, basic filtering
```

### Usage Examples

#### Basic Bot Detection

```go
detector := crawlerdetector.New()
isCrawler := detector.IsCrawler(userAgent)

if isCrawler {
    // Handle bot traffic
    log.Printf("Bot detected: %v", detector.GetMatched())
}
```

#### Mobile Device Detection

```go
detector := crawlerdetector.New()
isMobile := detector.IsMobile(userAgent)

if isMobile {
    // Serve mobile-optimized content
    fmt.Println("Mobile device detected")
}
```

#### Full Analysis

```go
detector := crawlerdetector.New()
detector.Parse(userAgent)

fmt.Printf("Browser: %v\n", detector.Browser)
fmt.Printf("Crawler: %v\n", detector.Crawler)
fmt.Printf("Mobile: %v\n", detector.Mobile)
fmt.Printf("Matches: %v\n", detector.GetMatched())
```

#### HTTP Middleware Example

```go
func BotDetectionMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        detector := crawlerdetector.NewShort()
        userAgent := r.Header.Get("User-Agent")

        if detector.IsCrawler(userAgent) {
            // Apply rate limiting for bots
            w.Header().Set("X-Bot-Detected", "true")
            log.Printf("Bot request: %v", detector.GetMatched())
        }

        next.ServeHTTP(w, r)
    })
}
```

### Performance

Version 2.0.0 introduces significant performance improvements:

- ✅ Regex patterns compiled once and cached (not per instance)
- ✅ Thread-safe concurrent operations with no race conditions
- ✅ Optimized memory allocation with pre-sized slices
- ✅ Lazy initialization reduces startup time

### Thread Safety

All detector methods are thread-safe and can be called concurrently:

```go
detector := crawlerdetector.New()

// Safe to call from multiple goroutines
go detector.Parse(userAgent1)
go detector.Parse(userAgent2)
```

### Contributing

If you find a bot/spider/crawler user agent that Crawler Detector doesn't detect, please submit a pull request with the regex pattern.

### License

MIT License - see [LICENSE](LICENSE) file for details.

### Changelog

See [CHANGELOG.md](CHANGELOG.md) for version history and updates.
