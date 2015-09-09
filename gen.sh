#!/usr/bin/env bash
# Generates the go code for ALL the gitignores,
# via github.com/github/gitignore

preparebar() {
# $1 - bar length
# $2 - bar char
    barlen=$1
    barspaces=$(printf "%*s" "$1")
    barchars=$(printf "%*s" "$1" | tr ' ' "$2")
}

progressbar() {
# $1 - number (-1 for clearing the bar)
# $2 - max number
    if [ $1 -eq -1 ]; then
        printf "\r  $barspaces\r"
    else
        barch=$(($1*barlen/$2))
        barsp=$((barlen-barch))
        printf "\r[%.${barch}s%.${barsp}s]\r" "$barchars" "$barspaces"
    fi
}

clonedLocation="ghignore"

if [ ! -d $clonedLocation ]
then
  git clone https://github.com/github/gitignore $clonedLocation
else
  pushd . >> /dev/null
  cd $clonedLocation
  git pull >> /dev/null
  popd >> /dev/null
fi

mkdir -p langs
langGo="langs/langs.go"

echo "package langs

// Gitignores is the generated map of language -> gitignore
var Gitignores = map[string][]byte{" > $langGo

gitignores=($clonedLocation/*.gitignore)
count=${#gitignores[@]}

preparebar $count "#"
let "start=0"

for gitignore in $clonedLocation/*.gitignore; do
  lang=`basename $gitignore .gitignore`

  printf "  \"%s\": []byte(\`" $lang >> $langGo
  cat $gitignore | sed 's/`/` + "`" + `/g'>> $langGo
  printf "\`),\n" >> $langGo
  let "start++"
  progressbar $start $count
done

echo "}" >> $langGo

echo ""
echo "Generated gitignore for go!"
