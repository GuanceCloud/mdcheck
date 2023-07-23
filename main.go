package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/GuanceCloud/mdcheck/check"
)

var (
	autofix,
	jsonOutput bool
	markdownDir,
	metaDir string
)

//nolint:gochecknoinits
func init() {
	flag.StringVar(&markdownDir, "md-dir", "", "markdown dirs")
	flag.StringVar(&metaDir, "meta-dir", "", "markdown meta dir, with meta dir set, only checking meta info of markdown")
	flag.BoolVar(&autofix, "autofix", false, "auto fix error")
	flag.BoolVar(&jsonOutput, "json", false, "set output in json format(2 space indent)")
}

func main() {
	flag.Parse()

	log.SetFlags(log.Lshortfile | log.LstdFlags)

	if markdownDir != "" {
		res, err := check.Check(
			check.WithMarkdownDir(markdownDir),
			check.WithMetaDir(metaDir),
			check.WithAutofix(autofix),
		)
		if err != nil {
			log.Printf("[E] %s", err.Error())
			return
		}

		if jsonOutput {
			if j, err := json.MarshalIndent(res, "", "  "); err != nil {
				log.Printf("[E] json.MarshalIndent: %s", err.Error())
			} else {
				fmt.Printf("%s", string(j))
			}
		} else {
			for _, r := range res {
				switch {
				case r.Err != "":
					fmt.Printf("%s: %q | Err: %s\n", r.Path, r.Text, r.Err)
				case r.Warn != "":
					fmt.Printf("%s: Warn: %s\n", r.Path, r.Warn)
				}
			}
		}
	}
}
