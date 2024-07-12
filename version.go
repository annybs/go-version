package version

import (
	"fmt"
)

// Version is a structured representation of a version number.
type Version struct {
	Major     int    // Major version number.
	Minor     int    // Minor version number.
	Patch     int    // Patch version number.
	Extension string // Version extension, such as pre-release number or build metdata.

	Text string // Original version string, if this version was created via the Parse function.
}

// Compare this version (a) with another version (b).
// This function returns -1 if a is less than b, 1 if a is greater than b, or 0 if a is equal to b.
//
// Extensions such as pre-release version or build metadata are ignored when comparing versions.
func (a *Version) Compare(b *Version) int {
	if a == nil || b == nil {
		return 0
	}

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

// Match tests the version against a constraint.
// Gt and Lt take precedence over Gte and Lte.
func (v *Version) Match(c *Constraint) bool {
	if v == nil {
		return false
	}

	if c == nil {
		return true
	}

	if c.Gt != nil {
		if v.Compare(c.Gt) <= 0 {
			return false
		}
	} else if c.Gte != nil {
		if v.Compare(c.Gte) < 0 {
			return false
		}
	}

	if c.Lt != nil {
		if v.Compare(c.Lt) >= 0 {
			return false
		}
	} else if c.Lte != nil {
		if v.Compare(c.Lte) > 0 {
			return false
		}
	}

	return true
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
	if v == nil {
		return ""
	}

	return fmt.Sprintf("%d.%d.%d%s", v.Major, v.Minor, v.Patch, v.Extension)
}

func (v *Version) String() string {
	if v == nil {
		return ""
	}

	if v.Text != "" {
		return v.Text
	}

	return v.SemanticString()
}
