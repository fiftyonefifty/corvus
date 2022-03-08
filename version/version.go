// The version package provides a location to set the release versions for all
// packages to consume. This is a clone of HashiCorp's implementation.
// Do Not import any other Reaper packages in this module.

package version

import (
	"fmt"
	"github.com/hashicorp/go-version"
)

// Version - the application version number
const Version = "0.1.0"

// Prerelease marker for the version (Thanks HashiCorp). If this is "" (empty string)
// then this is a final release version. Otherwise, this is a pre-release such as
// "dev" (in development), "beta", "rc1", etc.
var Prerelease = "alpha"

// SemVer is an instance of version.Version. This has the secondary
// benefit of verifying during tests and init time that our version is a
// proper semantic version, which should always be the case.
var SemVer *version.Version

func init() {
	SemVer = version.Must(version.NewVersion(Version))
}

// Header is the header name used to send the current terraform version
// in http requests.
const Header = "Corvus-Version"

// String returns the complete version string, including prerelease
func String() string {
	if Prerelease != "" {
		return fmt.Sprintf("%s-%s", Version, Prerelease)
	}
	return Version
}
