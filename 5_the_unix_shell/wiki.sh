#!/usr/bin/env bash
title=$1
section_title=$2

if [ $# -eq 0 ]
  then
    echo "No arguments supplied"
    exit 1
fi

if [ $# -gt 2 ]
  then
    echo "Too many arguments"
    exit 1
fi

summary_url="https://en.wikipedia.org/api/rest_v1/page/summary/${title}"

# get summary
# -s silent, -L follow redirects
# get property 'extract' from json response
summary=$(curl -s -L GET "${summary_url}" | jq '.extract')

echo ""
echo "Summary:"
echo ""

# print formatted summary
# todo: quit early if summary is empty
echo "${summary}" | lynx -stdin -dump


# get top level sections titles
url="https://en.wikipedia.org/w/api.php?action=parse&format=json&page=${title}&prop=sections&disabletoc=1"
response=$(curl -s -X GET ${url})
readarray -t sections < <(echo "${response}" | jq -r '.parse.sections | .[] | select(.toclevel == 1) | .line')

# we could use the following jq filter to get the string value snake cased to use as url params
# | gsub(" "; "_")

# check if arguments size is 1
# and print sections in line
if [ $# -eq 1 ]
  then
    echo ""
    echo "Sections:"
    echo ""
    for i in "${!sections[@]}"; do
        number=$((i+1))
        echo "${number}. ${sections[$i]}" | lynx -stdin -dump
    done
fi

# check if arguments size is 2
# Life history must be passed as "Life history" to be counted a single arg
if [ $# -eq 2 ]
  then

    # todo:
    # the html returned has all the details we need
    # the section title and the subsection have this before them `class="mw-headline"`
    # the section details can be found by looking for the first <p>
    # we can then look for the first ". " ? maybe...
    # but we need to strip all html in there

    # todo: the history subsection for walmart takes a long time...why?
    for i in "${!sections[@]}"; do
      if [[ "${sections[$i]}" = "${section_title}" ]]; then

        index=$((i+1))
        section_detail_url="https://en.wikipedia.org/w/api.php?action=parse&format=json&page=${title}&prop=text&section=${index}&disabletoc=1"
        response=$(curl -s -X GET "${section_detail_url}")
        detail=$(echo "${response}" | jq '.parse.text."*"')

        echo ""
        echo "${sections[$i]}:"
        echo ""
        
        # NICE: this works
        echo "${detail//\\}" | \
        gsed -n 's:.*<p>\(.*\)</p>.*:\1:p' | \
        gsed -e 's/<[^>]*>//g' | \
        lynx -stdin -dump
        # or use the following
        # gsed 's/&#160;//g' | \
        # gsed 's/&#91;//g' | \
        # gsed 's/2&#93;//g'
        break
      fi
    done

    exit 0
fi

