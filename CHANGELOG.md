# Changelog

All notable changes to MrRSS will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.1.3] - 2025-11-22

### Added
- **Automatic Update System**: One-click download and installation of updates
  - Automatically detects user's operating system and CPU architecture
  - Downloads appropriate installer from GitHub releases
  - Launches installer and prepares for update
  - Fallback to manual download for unsupported platforms
- **Multi-Platform Support**: Expanded from 3 to 8 build combinations
  - Windows: x64 (amd64), ARM64, x86 (386)
  - Linux: x64 (amd64), ARM64 (aarch64), x86 (386)
  - macOS: Apple Silicon (arm64), Intel (amd64)
- **Download Progress Indicators**: Visual feedback during update download and installation
- **Internationalization**: Added i18n strings for update process (English and Chinese)
  - downloading, downloadComplete, installingUpdate
  - updateWillRestart, downloadFailed, installFailed

### Changed
- **Refactored SettingsModal**: Split large component (653 lines) into focused tab components
  - SettingsModal.vue (380 lines) - Main orchestrator
  - GeneralTab.vue (140 lines) - Appearance, updates, database, translation settings
  - FeedsTab.vue (138 lines) - Data management and feed management
  - AboutTab.vue (120 lines) - Version info, update checking, auto-update UI
- **Removed macOS Universal Build**: Now only builds architecture-specific versions
  - Improved build efficiency
  - Users get optimized binaries for their specific architecture

### Security
- **Enhanced Update Security**:
  - URL validation restricted to official repository (`https://github.com/WCY-dt/MrRSS/releases/download/`)
  - Path traversal protection on asset names and file paths
  - File type validation per platform (.exe/.appimage/.dmg)
  - Regular file verification to prevent symlink exploitation
  - CodeQL security scan passed with 0 vulnerabilities

### Documentation
- Updated README.md and README_zh.md with platform-specific download instructions
- Added architecture guidance (x64, ARM64, x86, Apple Silicon, Intel)
- Documented all 8 supported platform/architecture combinations

## [1.1.2] - 2025-11-22

### Added
- **Theme Auto-Detection**: Automatically detect and apply system theme preference
- **Translation Feature Improvements**: Better defaults for translation settings
- Version check functionality in Settings → About tab

### Changed
- Simplified update check UI
- Improved theme switching mechanism
- Better handling of translation provider selection

### Fixed
- Theme switching issues between light and dark modes
- Translation default language selection
- Update notification display

## [1.1.1] - 2025-11-21

### Added
- Initial release preparation
- OPML import/export functionality
- Feed category organization

### Fixed
- Various bug fixes and stability improvements
- UI refinements for better user experience

## [1.1.0] - 2025-11-20

### Added
- **Initial Public Release** of MrRSS
- **Cross-Platform Support**: Native desktop app for Windows, macOS, and Linux
- **RSS Feed Management**: Add, edit, and delete RSS feeds
- **Article Reading**: Clean, distraction-free reading interface
- **Smart Organization**: Organize feeds into categories
- **Favorites & Reading Tracking**: Save articles and track read/unread status
- **Modern UI**: Clean, responsive interface with dark mode support
- **Auto-Translation**: Translate article titles using Google Translate or DeepL API
- **OPML Support**: Import and export feed subscriptions
- **Auto-Update**: Configurable interval for fetching new articles
- **Database Cleanup**: Automatic removal of old articles
- **Multi-Language Support**: English and Chinese interface
- **Theme Support**: Light, dark, and auto (system) themes

### Technical Stack
- Backend: Go 1.21+ with Wails v2
- Frontend: Vue.js 3 with Composition API
- Styling: Tailwind CSS
- Database: SQLite
- Build: Cross-platform builds for Windows, macOS, and Linux

---

## Release Notes

### Version Numbering
MrRSS follows [Semantic Versioning](https://semver.org/):
- **MAJOR** version for incompatible API changes
- **MINOR** version for backwards-compatible functionality additions
- **PATCH** version for backwards-compatible bug fixes

### Download
Downloads for all platforms are available on the [GitHub Releases](https://github.com/WCY-dt/MrRSS/releases) page.

### Upgrade Notes
When upgrading from a previous version:
1. Your data (feeds, articles, settings) is preserved in platform-specific directories
2. Database migrations are applied automatically on first launch
3. For major version upgrades, please review the changelog for breaking changes

### Support
- Report bugs: [GitHub Issues](https://github.com/WCY-dt/MrRSS/issues)
- Feature requests: [GitHub Issues](https://github.com/WCY-dt/MrRSS/issues)
- Documentation: [README](README.md)
