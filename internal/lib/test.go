package lib

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func list(packages []string) ([]string, error) {
	args := append([]string{"list"}, packages...)
	cmd := exec.Command("go", args...)
	cmd.Stderr = os.Stderr
	b, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var res []string
	lines := strings.Split(string(b), "\n")
	for _, l := range lines {
		l = strings.TrimSpace(l)
		if l != "" {
			res = append(res, l)
		}
	}
	return res, nil
}

func Test(flagSet *flag.FlagSet, output string) error {
	packages, err := list(flagSet.Args())
	if err != nil {
		return err
	}

	var flags []string
	flagSet.Visit(func(f *flag.Flag) {
		flags = append(flags, fmt.Sprintf("-%s=%s", f.Name, f.Value.String()))
	})
	flags = append(flags, fmt.Sprintf("-coverpkg=%s", strings.Join(packages, ",")))

	f, err := ioutil.TempFile("", "gocovermerge-")
	if err != nil {
		return err
	}
	if err = f.Close(); err != nil {
		return err
	}
	dir := f.Name()
	if err = os.Remove(dir); err != nil {
		return err
	}
	if err = os.Mkdir(dir, 0777); err != nil {
		return err
	}

	files := make([]string, 0, len(packages))
	for _, p := range packages {
		if f, err = ioutil.TempFile(dir, filepath.Base(p)+"-"); err != nil {
			return err
		}
		files = append(files, f.Name())
		if err = f.Close(); err != nil {
			return err
		}

		args := append([]string{"test"}, flags...)
		args = append(args, fmt.Sprintf("-coverprofile=%s", f.Name()))
		args = append(args, p)
		cmd := exec.Command("go", args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		log.Printf(strings.Join(cmd.Args, " "))
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	if err = Merge(files, output); err != nil {
		return err
	}
	return os.RemoveAll(dir)
}
