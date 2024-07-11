package version

import (
	"fmt"
)

// Version is a structured representation of a version number
type Version struct {
	Major     int
	Minor     int
	Patch     int
	Extension string

	Text string
}

// Compare this version (a) with another version (b).
// This function returns -1 if a is less than b, 1 if a is greater than b, or 0 if a is equal to b.
//
// Extensions such as pre-release version or build metadata are ignored when comparing versions.
func (a *Version) Compare(b *Version) int {
	if a.Major == b.Major {
		if a.Minor == b.Minor {
			if a.Patch == b.Patch {
				return 0
			} else if a.Patch > b.Patch {
				return 1
			}
		} else if a.Minor > b.Minor {
			return 1
		}
	} else if a.Major > b.Major {
		return 1
	}

	return -1
}

// Less performs a simple comparison of this version (a) with another version (b).
// This function returns true if a is less than b, or false otherwise.
//
// Extensions such as pre-release version or build metadata are ignored when comparing versions.
func (a *Version) Less(b *Version) bool {
	return a.Compare(b) < 0
}

// SemanticString returns a version string conforming to the standard described in Semantic Versioning 2.0.0.
//
// See https://semver.org/#is-v123-a-semantic-version
func (v *Version) SemanticString() string {
	return fmt.Sprintf("%d.%d.%d%s", v.Major, v.Minor, v.Patch, v.Extension)
}

func (v *Version) String() string {
	if v.Text != "" {
		return v.Text
	}
	return v.SemanticString()
}
