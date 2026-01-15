# Changelog

All notable changes to this project will be documented in this file.

## [2.0.0] - 2026-01-15

### Breaking Changes

- Updated minimum Go version to 1.21
- Removed global `sync.WaitGroup` variable (fixes race conditions)

### Added

- Lazy regex compilation with `sync.Once` for significant performance improvements
- Proper error handling in `PiwikCrawlersList()` and `PiwikMobilesList()`
- Better logging with contextual prefixes
- Comprehensive documentation for all methods
- Thread-safety guarantees documented

### Changed

- **CRITICAL FIX**: Replaced global `sync.WaitGroup` with local instances in `Parse()` and `ParseUnsafe()` methods
  - Previous implementation caused race conditions when called concurrently
  - Now fully thread-safe for concurrent usage
- Optimized regex compilation - patterns now compiled once and cached
  - Separate caching for Full, Piwik, and Short detection modes
  - Dramatically improved initialization performance
- Improved memory allocation patterns
  - Pre-allocated slices with proper capacity
  - Filtered empty regex patterns
- Enhanced error handling
  - Graceful degradation on JSON unmarshal failures
  - Returns empty pattern "()" instead of potentially invalid patterns
- Updated `go.mod` with proper dependency management

### Performance Improvements

- Regex compilation now happens once per pattern set (not per instance)
- Reduced memory allocations in pattern building
- Faster instance creation (New(), NewPiwik(), NewShort())
- Better slice pre-allocation reduces GC pressure

### Fixed

- Race condition in concurrent Parse() calls
- Memory leaks from repeated regex compilation
- Inconsistent error handling
- Missing nil checks for empty patterns

### Documentation

- Added ADR-0007: Use Regex-Based User Agent Detection
- Added ADR-0008: Refactor Global WaitGroup to Instance-Level
- Added ADR-0009: Implement Lazy Regex Compilation with sync.Once
- Improved inline documentation for all public methods
- Added thread-safety guarantees to method documentation

## [1.0.0] - 2018

### Initial Release

- Basic crawler detection functionality
- Support for Piwik patterns
- Mobile device detection
- Browser exclusion patterns
