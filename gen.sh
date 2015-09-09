#!/usr/bin/env bash
mkdir -p langs

langGo="langs.go"

echo "package gitignores

// Gitignores is the generated map of language -> gitignore
var Gitignores = map[string]string{" > $langGo

for gitignore in gitignore-cloned/*.gitignore; do
  lang=`basename $gitignore .gitignore`

  printf "  \"%s\": \`" $lang >> $langGo
  cat $gitignore | sed 's/`/` + "`" + `/g'>> $langGo
  printf "\`,\n" >> $langGo
done

echo "}" >> $langGo
