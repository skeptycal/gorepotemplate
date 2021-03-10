package gorepo

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/skeptycal/zsh"
)

const (
	defaultGitignoreItems = "macos linux windows ssh vscode go zsh node vue nuxt python django"
)

var (
	repoName string
	repoPath string
)

func init() {
	pwd, _ := os.Getwd()
	repoPath, repoName = path.Split(path.Clean(pwd))
}

// Setup initializes the repo, creates files, prompts as needed, creates the
// github.com repository, and pushes the initial commit.
func Setup() error {
	err := gitRepoSetup()
	if err != nil {
		log.Printf("gitRepoSetup failed with %v", err)
		return err
	}
	err = createAutomatedFiles()
	if err != nil {
		log.Printf("createAutomatedFiles failed with %v", err)
		return err
	}
}

// gitRepoSetup initializes the repo, prompts as needed, creates the
// github.com repository, and pushes the initial commit.
func gitRepoSetup() error {
	err := gitInit()
	if err != nil {
		return err
	}
	// todo - stuff
	return err
}

// CreateAutomatedFiles creates the automated files.
func createAutomatedFiles() error {
	zsh.Sh("touch main.go")
	return nil
}

// gitIgnore writes a .gitignore file, including default items followed by the response from
// the www.gitignore.io API containing standard .gitignore items for the args given.
//
//      default: "macos linux windows ssh vscode go zsh node vue nuxt python django"
//
// using: https://www.toptal.com/developers/gitignore/api/macos,linux,windows,ssh,vscode,go,zsh,node,vue,nuxt,python,django
func gitIgnore(args string) error {
	// notes - .gitignore header
	/*
	   # gorepo - .gitignore file

	   # generic secure items:
	   *private*
	   *secret*
	   *bak

	   # repo specific items
	   coverage.txt
	   profile.out
	*/

	if args == "" {
		args = defaultGitignoreItems
	}

	var sb strings.Builder
	defer sb.Reset()

	sb.WriteString(fmt.Sprintf("# %s - .gitignore file\n\n", repoName))

	sb.WriteString("# generic secure items:\n")
	sb.WriteString("*private*\n*secret*\n*bak\n\n")

	sb.WriteString("# repo specific items:\n")
	sb.WriteString("coverage.txt\nprofile.out\n\n")

	// add .gitignore contents from gitignore.io API
	sb.WriteString(gi(args))

	return WriteFile(".gitignore", sb.String())
}

// GitInit initializes the Git environment
func gitInit() error {
	if !fileExists(".gitignore") {
		gitIgnore("")
	}

	Shell("git init")
	GitCommitAll("initial commit")
	return nil
}

// GoMod creates and initializes the repo go.mod file.
func GoMod() error {
	Shell("go mod init")
	Shell("go mod tidy")
	Shell("go mod download")
	GitCommit("go mod setup")
	return nil
}

// GitCommit creates a commit with message
func GitCommitAll(message string) error {
	Shell("git add --all")
	Shell("git commit -m '" + message + "'")
}

// GoSum creates the go.sum file.
func GoSum() error {
	return nil
}

// GoTestSh creates the go.test.sh script.
func GoTestSh() error {
	return nil
}

// GoDoc creates the go.doc file.
func GoDoc() error {
	return nil
}

// BugReportMd creates the .github/ISSUE_TEMPLATE/bug_report.md file.
func BugReportMd() error {
	return nil
}

// FeatureRequestMd creates the .github/ISSUE_TEMPLATE/feature_request.md file.
func FeatureRequestMd() error {
	return nil
}

// GitWorkflows creates initial .github/workflows/... files:
// codeql-analysis.yml go.yml greetings.yml label.yml stale.yml
func GitWorkflows() error {
	return nil
}

// CodeCovYml creates the initial .codecov.yml file.
func CodeCovYml() error {
	return nil
}

// FundingYml creates the initial FUNDING.yml file.
func FundingYml() error {
	return nil
}

// PreCommitYaml creates the initial .pre-commit-config.yaml file.
func PreCommitYaml() error {
	return nil
}

// TravisYml creates the initial .travis.yml file.
func TravisYml() error {
	return nil
}

// DocGo creates the initial doc.go file.
func DocGo() error {
	return nil
}

// ReadMeMd creates the initial README.md file.
func ReadMeMd() error {
	return nil
}

// SecurityMd creates the initial SECURITY.md file.
func SecurityMd() error {
	return nil
}

// CodeOfConduct creates the initial CODE_OF_CONDUCT.md file.
func CodeOfConduct() error {
	return nil
}

// License creates the initial LICENSE file.
func License(license string) error {
	return nil
}
