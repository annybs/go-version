package version

import (
	"strconv"
	"strings"
)

const (
	sectionMajor     = 0
	sectionMinor     = 1
	sectionPatch     = 2
	sectionExtension = 3
)

func MustParse(str string) *Version {
	v, err := Parse(str)
	if err != nil {
		panic(err)
	}
	return v
}

func Parse(str string) (*Version, error) {
	v := &Version{Text: string(str)}

	if len(str) == 0 {
		return nil, invalid(str)
	}

	section := sectionMajor
	chars := []byte{}

	commit := func() error {
		if len(chars) == 0 {
			return invalid(str)
		}

		if section < sectionExtension {
			n, err := strconv.Atoi(string(chars))
			if err != nil {
				return invalid(str)
			}

			switch section {
			case sectionMajor:
				v.Major = n
			case sectionMinor:
				v.Minor = n
			case sectionPatch:
				v.Patch = n
			}
		} else {
			v.Extension = string(chars)
		}

		chars = []byte{}
		return nil
	}

	for i := 0; i < len(str); i++ {
		c := str[i]

		if i == 0 && strings.IndexByte("vV", c) > -1 {
			continue
		}

		if section < sectionExtension {
			if strings.IndexByte("0123456789", c) > -1 {
				chars = append(chars, c)
			} else {
				if err := commit(); err != nil {
					return nil, err
				}
				if c == '.' {
					section++
				} else {
					section = sectionExtension
					chars = append(chars, c)
				}
			}
		} else {
			chars = append(chars, c)
		}
	}
	if len(chars) > 0 {
		if err := commit(); err != nil {
			return nil, err
		}
	}

	return v, nil
}
