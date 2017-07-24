package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"

	"github.com/AlekSi/gocovermerge/internal/lib"
)

var (
	coverprofileF = flag.String("coverprofile", "cover.out", "Output file.")

	testFlagSet = flag.NewFlagSet("test", flag.ExitOnError)

	// The build flags are shared by the build, clean, get, install, list, run, and test commands:
	aF    = testFlagSet.Bool("a", false, "force rebuilding of packages that are already up-to-date.")
	nF    = testFlagSet.Bool("n", false, "print the commands but do not run them.")
	raceF = testFlagSet.Bool("race", false, "enable data race detection.")
	msanF = testFlagSet.Bool("msan", false, "enable interoperation with memory sanitizer.")
	workF = testFlagSet.Bool("work", false, "print the name of the temporary work directory and do not delete it when exiting.")
	xF    = testFlagSet.Bool("x", false, "print the commands.")
	tagsF = testFlagSet.String("tags", "", "a list of build tags to consider satisfied during the build.")
	// -p is not supported
	// -v is redefined below

	// The test binary flags:
	covermodeF = testFlagSet.String("covermode", "", "set the mode for coverage analysis for the package[s] being tested.")
	shortF     = testFlagSet.Bool("short", false, "tell long-running tests to shorten their run time.")
	timeoutF   = testFlagSet.Duration("timeout", 0, "if a test runs longer than t, panic.")
	vF         = testFlagSet.Bool("v", false, "verbose output: log all tests as they are run.")

	// TODO add more flags
)

func main() {
	log.SetFlags(0)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	var err error
	switch flag.Arg(0) {
	case "merge":
		err = lib.Merge(flag.Args()[1:], *coverprofileF)

	case "test":
		testFlagSet.Parse(flag.Args()[1:])
		err = lib.Test(testFlagSet, *coverprofileF)

	default:
		flag.Usage()
		log.Fatalf("Unexpected command %q.", flag.Arg(0))
	}

	if err != nil {
		log.Print(err)
		if eErr, ok := err.(*exec.ExitError); ok {
			if ws, ok := eErr.Sys().(*syscall.WaitStatus); ok {
				os.Exit(ws.ExitStatus())
			}
		}
		os.Exit(1)
	}
}
