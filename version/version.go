package version

import "fmt"

var (
	gitTag    string
	gitCommit string
	gitBranch string
	buildTime string
	goVersion string
)

// FullVersion show the version info
func FullVersion() string {
	version := fmt.Sprintf("Version   : %s\nBuild Time: %s\nGit Branch: %s\nGit Commit: %s\nGo Version: %s\n", gitTag, buildTime, gitBranch, gitCommit, goVersion)
	return version
}

// Short 版本缩写
func Short() string {
	commit := ""
	if len(gitCommit) > 8 {
		commit = gitCommit[:8]
	}
	return fmt.Sprintf("%s[%s %s]", gitTag, buildTime, commit)
}
