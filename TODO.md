### DP1 — Project Foundation & Configuration

Set up the Go project, CLI structure, YAML configuration, configuration validation, and logging. Implement init, validate, and version commands.

### DP2 — GitHub Releases Integration

Implement GitHub API integration to discover releases, support public and private repositories, select release assets, and download artifacts securely. Support authentication, checksums, and error handling.

###  DP3 — Deployment Engine

Implement versioned release directories, safe artifact extraction, atomic symlink switching, systemd service restarts, health checks, rollback, and release retention.

### DP4 — State & Reconciliation

Track deployed versions and compare them with configured GitHub Releases. Implement idempotent deployments, locking, failure recovery, and retry handling.

### DP5 — Linux Service Integration

Implement installation and uninstallation commands that configure a dedicated service account and systemd service/timer. Support configurable polling intervals and automatic startup after reboot.

### DP6 — Distribution & Installation

Configure GitHub Actions to build and publish precompiled Linux binaries for AMD64 and ARM64. Provide an installation script, checksums, documentation, and diagnostic commands.

### DP7 — Testing & Initial Release

Add unit, integration, and end-to-end tests covering successful deployments, failures, and rollback. Validate the installation process on supported Linux distributions and publish the initial release.
