package lib

import (
	"fmt"
	"os"
	"sort"

	"golang.org/x/tools/cover"
)

func Merge(files []string, output string) error {
	blocks := make(map[string][]cover.ProfileBlock)
	var mode string
	for _, f := range files {
		profiles, err := cover.ParseProfiles(f)
		if err != nil {
			return err
		}
		for _, p := range profiles {
			if mode == "" {
				mode = p.Mode
			}
			if mode != p.Mode {
				return fmt.Errorf("different modes: %s and %s", mode, p.Mode)
			}

			blocks[p.FileName] = append(blocks[p.FileName], p.Blocks...)
		}
	}

	files = make([]string, 0, len(blocks))
	for file := range blocks {
		files = append(files, file)
	}
	sort.Strings(files)

	f, err := os.Create(output)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.WriteString(fmt.Sprintf("mode: %s\n", mode)); err != nil {
		return err
	}
	for _, file := range files {
		for _, b := range blocks[file] {
			// encoding/base64/base64.go:34.44,37.40 3 1
			// where the fields are: name.go:line.column,line.column numberOfStatements count
			l := fmt.Sprintf("%s:%d.%d,%d.%d %d %d\n", file, b.StartLine, b.StartCol, b.EndLine, b.EndLine, b.NumStmt, b.Count)
			if _, err = f.WriteString(l); err != nil {
				return err
			}
		}
	}
	return nil
}
