package version

import "github.com/coreos/go-semver/semver"

var (
	// VersionMajor is for an API incompatible changes.
	VersionMajor int64 = 0
	// VersionMinor is for functionality in a backwards-compatible manner.
	VersionMinor int64 = 1
	// VersionPatch is for backwards-compatible bug fixes.
	VersionPatch int64 = 0
	// VersionPre indicates prerelease.
	VersionPre = "rc.1"
	// VersionDev indicates development branch. Releases will be empty string.
	VersionDev string

	// ApiVersionMajor is for an API incompatible changes.
	ApiVersionMajor int64 = 0
	// ApiVersionMinor is for functionality in a backwards-compatible manner.
	ApiVersionMinor int64 = 1
	// ApiVersionPatch is for backwards-compatible bug fixes.
	ApiVersionPatch int64 = 0
	// ApiVersionPre indicates prerelease.
	ApiVersionPre = "rc.1"
	// ApiVersionDev indicates development branch. Releases will be empty string.
	ApiVersionDev string

	// GitHash Value will be set during build
	GitHash = "Not provided"

	// BuildTime Value will be set during build
	BuildTime = "Not provided"
)

// Version is the specification version that the package types support.
var Version = semver.Version{
	Major:      VersionMajor,
	Minor:      VersionMinor,
	Patch:      VersionPatch,
	PreRelease: semver.PreRelease(VersionPre),
	Metadata:   VersionDev,
}

// ApiVersion is the specification version that the package types support.
var ApiVersion = semver.Version{
	Major:      VersionMajor,
	Minor:      VersionMinor,
	Patch:      VersionPatch,
	PreRelease: semver.PreRelease(ApiVersionPre),
	Metadata:   VersionDev,
}
