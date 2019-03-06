package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/dweomer/semver-cli/version"
)

var (
	cli  = flag.NewFlagSet("", flag.ExitOnError)
	from = cli.String("from", "", "<semver-source>")
	bump = cli.String("bump", "", "format: major|minor|patch|final|pre[=<prefix>]")
	meta = cli.String("meta", "", "[build-metadata]")
)

func init() {
	// parse the command line
	if err := cli.Parse(os.Args[1:]); err != nil {
		fail(err)
	}

	// validate required flags
	cli.VisitAll(func(f *flag.Flag) {
		switch f.Name {
		case "from":
			if f.Value.String() == "" {
				fail(fmt.Errorf("flag is required: -%s", f.Name))
			}
		}
	})
}

func main() {
	v, err := version.Make(*from, options(cli)...)
	if err != nil {
		fail(err)
	}
	fmt.Println(v.String())
}

func options(cli *flag.FlagSet) []version.MakeOpt {
	opts := []version.MakeOpt{}

	// validate flag values and translate to options
	cli.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "meta":
			opts = append(opts, version.WithBuild(*meta))
		case "bump":
			switch p := strings.SplitN(*bump, "=", 2); p[0] {
			case "final":
				opts = append(opts, version.BumpFinal())
			case "major":
				opts = append(opts, version.BumpMajor())
			case "minor":
				opts = append(opts, version.BumpMinor())
			case "patch":
				opts = append(opts, version.BumpPatch())
			case "pre":
				pre := "pre"
				if len(p) == 2 {
					pre = p[1]
				}
				opts = append(opts, version.BumpPre(pre))
			default:
				fail(fmt.Errorf("flag with bad format: -bump=%s", *bump))
			}
		}
	})

	return opts
}

func fail(err error) {
	fmt.Fprintf(os.Stderr, "%s", err)
	fmt.Fprintln(os.Stderr)
	cli.Usage()
	os.Exit(1)
}
