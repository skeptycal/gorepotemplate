package gorepo

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var (
	redb, _        = hex.DecodeString("1b5b33316d0a") // byte code for ANSI red
	red     string = string(redb)                     // ANSI red
)

// WriteFile creates the file 'fileName' and writes 'data' to it.
// It returns any error encountered. If the file already exists, it
// will be TRUNCATED and OVERWRITTEN.
func WriteFile(fileName string, data string) error {
	dataFile, err := OpenTrunc(fileName)
	if err != nil {
		log.Println(err)
		return err
	}
	defer dataFile.Close()

	n, err := dataFile.WriteString(data)
	if err != nil {
		log.Println(err)
		return err
	}
	if n != len(data) {
		log.Printf("incorrect string length written (wanted %d): %d\n", len(data), n)
		return fmt.Errorf("incorrect string length written (wanted %d): %d", len(data), n)
	}
	return nil
}

// type result struct {
// 	stdout string
// 	stderr string
// 	retval int
// 	err    error
// }

// gi returns a string response from the www.gitignore.io API containing
// standard .gitignore items for the args given.
//
//      default: "macos linux windows ssh vscode go zsh node vue nuxt python django"
//
// using: https://www.toptal.com/developers/gitignore/api/macos,linux,windows,ssh,vscode,go,zsh,node,vue,nuxt,python,django
func gi(args string) string {

	if len(args) == 0 {
		args = []string{"macos linux windows ssh vscode go zsh node vue nuxt python django"}
	}

	command := "curl -fLw '\n' https://www.gitignore.io/api/\"${(j:,:)@}\" "
	command += strings.Join(args, " ")

	return Shell(command)
}

// Shell executes a command line string and returns the result.
func Shell(command string) string {

	cmd := cmdPrep(command)
	stdout, err := cmd.Output()

	if err != nil {
		return fmt.Errorf("%Terror: %v", red, err).Error()
	}

	return string(stdout)
}

// OpenTrunc creates and opens the named file for writing. If successful, methods on
// the returned file can be used for writing; the associated file descriptor has mode
//      O_WRONLY|O_CREATE|O_TRUNC
// If the file does not exist, it is created with mode o644;
//
// If the file already exists, it is TRUNCATED and overwritten
//
// If there is an error, it will be of type *PathError.
func OpenTrunc(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 644)
}

// cmdPrep prepares a Cmd struct from a command line string.
func cmdPrep(command string) *exec.Cmd {
	commandSlice := strings.Split(command, " ")
	app := commandSlice[0]
	args := strings.Join(commandSlice[1:], " ")
	return exec.Command(app, args)
}

// fileExists checks if a file exists and is not a directory
func fileExists(fileName string) bool {
	info, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Notes: Cmd struct summary:
/*
type Cmd struct {
	Path            string
	Args            []string
	Env             []string
	Dir             string
	Stdin           io.Reader
	Stdout          io.Writer
	Stderr          io.Writer
	ExtraFiles      []*os.File
	SysProcAttr     *syscall.SysProcAttr
	Process         *os.Process
	ProcessState    *os.ProcessState
	ctx             context.Context // nil means none
	lookPathErr     error           // LookPath error, if any.
	finished        bool            // when Wait was called
	childFiles      []*os.File
	closeAfterStart []io.Closer
	closeAfterWait  []io.Closer
	goroutine       []func() error
	errch           chan error // one send per goroutine
	waitDone        chan struct{}
}
*/
