package version

import (
	"strings"

	"github.com/blang/semver"
)

type MakeOpt func(*semver.Version) error

func Make(s string, m ...MakeOpt) (semver.Version, error) {
	if len(s) > 0 && s[0] == 'v' {
		s = s[1:]
	}
	v, e := semver.Parse(s)
	if e != nil {
		return v, e
	}
	for _, o := range m {
		if e := o(&v); e != nil {
			return v, e
		}
	}
	return v, nil
}

func BumpFinal() MakeOpt {
	return func(v *semver.Version) error {
		v.Pre = nil
		return nil
	}
}

func BumpMajor() MakeOpt {
	return func(v *semver.Version) error {
		v.Major++
		v.Minor = 0
		v.Patch = 0
		v.Pre = nil
		return nil
	}
}

func BumpMinor() MakeOpt {
	return func(v *semver.Version) error {
		v.Minor++
		v.Patch = 0
		v.Pre = nil
		return nil
	}
}

func BumpPatch() MakeOpt {
	return func(v *semver.Version) error {
		v.Patch++
		v.Pre = nil
		return nil
	}
}

func BumpPre(pre string) MakeOpt {
	return func(v *semver.Version) error {
		if v.Pre == nil || v.Pre[0].VersionStr != pre || len(v.Pre) <= 1 {
			v.Pre = []semver.PRVersion{
				{VersionStr: pre},
				{VersionNum: 1, IsNum: true},
			}
		} else {
			v.Pre = []semver.PRVersion{
				{VersionStr: pre},
				{VersionNum: v.Pre[1].VersionNum + 1, IsNum: true},
			}
		}
		return nil
	}
}

func WithBuild(s string) MakeOpt {
	return func(v *semver.Version) error {
		v.Build = strings.Split(s, ".")
		return nil
	}
}
