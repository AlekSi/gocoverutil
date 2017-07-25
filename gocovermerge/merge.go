package gocovermerge

import (
	"fmt"
	"os"
	"sort"

	"golang.org/x/tools/cover"
)

// Merge combines several coverage files into single file.
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

	// sort files
	files = make([]string, 0, len(blocks))
	for file := range blocks {
		files = append(files, file)
	}
	sort.Strings(files)

	for _, file := range files {
		// sort blocks
		sort.Slice(blocks[file], func(i, j int) bool {
			bi, bj := blocks[file][i], blocks[file][j]
			return bi.StartLine < bj.StartLine || bi.StartLine == bj.StartLine && bi.StartCol < bj.StartCol
		})

		// merge blocks
		var newBlocks []cover.ProfileBlock
		var prev cover.ProfileBlock
		for _, b := range blocks[file] {
			// skip full duplicate
			if prev == b {
				continue
			}

			// change count inside previous block if only count changed
			prev.Count = b.Count
			if prev == b {
				if mode == "set" {
					newBlocks[len(newBlocks)-1].Count = 1
				} else {
					newBlocks[len(newBlocks)-1].Count += b.Count
				}
				prev = b
				continue
			}

			newBlocks = append(newBlocks, b)
			prev = b
		}
		blocks[file] = newBlocks
	}

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
			l := fmt.Sprintf("%s:%d.%d,%d.%d %d %d\n", file, b.StartLine, b.StartCol, b.EndLine, b.EndCol, b.NumStmt, b.Count)
			if _, err = f.WriteString(l); err != nil {
				return err
			}
		}
	}
	return nil
}
