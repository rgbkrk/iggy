//go:generate ./gen.sh
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/rgbkrk/iggy/langs"
)

var languageAliases = map[string]string{
	"js":         "node",
	"javascript": "node",
	"nodejs":     "node",
}

func normalizeLanguage(lang string) string {
	l := strings.ToLower(lang)
	alias, ok := languageAliases[l]
	if !ok {
		return l
	}
	return alias
}

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		fmt.Println("Usage: iggy <language>")
	}
	lang := normalizeLanguage(flag.Arg(0))
	gitignore, ok := langs.Gitignores[lang]

	if !ok {
		fmt.Printf("Language %s not available\n", lang)
		os.Exit(1)
	}

	f, err := os.Create(".gitignore")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString(gitignore)
	f.Sync()

}
