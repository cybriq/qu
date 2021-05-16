package version

//go:generate go run ./update/.

import (
	"fmt"
)

var (

	// URL is the git URL for the repository
	URL = "github.com/cybriq/qu"
	// GitRef is the gitref, as in refs/heads/branchname
	GitRef = "refs/heads/main"
	// GitCommit is the commit hash of the current HEAD
	GitCommit = "9b4158bd7b653d7c4e4c4711d6978f353c4bae65"
	// BuildTime stores the time when the current binary was built
	BuildTime = "2021-05-16T11:21:48+03:00"
	// Tag lists the Tag on the build, adding a + to the newest Tag if the commit is
	// not that commit
	Tag = "v0.0.12+"
	// PathBase is the path base returned from runtime caller
	PathBase = "/home/loki/src/github.com/cybriq/qu/"
	// Major is the major number from the tag
	Major = 0
	// Minor is the minor number from the tag
	Minor = 0
	// Patch is the patch version number from the tag
	Patch = 12
	// Meta is the extra arbitrary string field from Semver spec
	Meta = ""
)

// Get returns a pretty printed version information string
func Get() string {
	return fmt.Sprint(
		"\nRepository Information\n"+
			"\tGit repository: "+URL+"\n",
		"\tBranch: "+GitRef+"\n"+
			"\tCommit: "+GitCommit+"\n"+
			"\tBuilt: "+BuildTime+"\n"+
			"\tTag: "+Tag+"\n",
		"\tMajor:", Major, "\n",
		"\tMinor:", Minor, "\n",
		"\tPatch:", Patch, "\n",
		"\tMeta: ", Meta, "\n",
	)
}
