#!/usr/bin/env bash
title=$1
section_title=$2

if [ $# -eq 0 ]
  then
    echo "No arguments supplied"
    exit 1
fi

# Get search hits
# url="https://en.wikipedia.org/w/api.php?action=opensearch&search=${title}&limit=3&namespace=0&format=json"
# echo "$url"
# result=$(curl -s -X GET ${url} | jq -r '.[1] | map(sub(" "; "_")) | join(" ")')

# arr=($result)

# for i in ${arr[@]}; do echo $i; done
# todo: use jq map to replace space with string

# get summary
summary_url="https://en.wikipedia.org/api/rest_v1/page/summary/${title}"
summary=$(curl -s -X GET "${summary_url}" | jq '.extract')

echo ""
echo "Summary:"
echo ""
echo "${summary}" | lynx -stdin -dump

if [ $# -eq 1 ]
  then
    # get sections titles
    url="https://en.wikipedia.org/w/api.php?action=parse&format=json&page=${title}&prop=sections&disabletoc=1"
    response=$(curl -s -X GET ${url})
    sections=$(echo "${response}" | jq '.parse.sections' | jq '.[] | select(.toclevel == 1) | .line')
    # sections_index=$(echo "${response}" | jq '.parse.sections' | jq '.[] | select(.toclevel == 1) | .number')
    echo ""
    echo "Sections:"
    echo ""
    echo "${sections}"
fi

if [ $# -eq 2 ]
  then

    # todo:
    # the html returned has all the details we need
    # the section title and the subsection have this before them `class="mw-headline"`
    # the section details can be found by looking for the first <p>
    # we can then look for the first ". " ? maybe...
    # but we need to strip all html in there
    url="https://en.wikipedia.org/w/api.php?action=parse&format=json&page=${title}&prop=sections&disabletoc=1"
    response=$(curl -s -X GET ${url})
    # sections=$(echo "${response}" | jq -r '.parse.sections | .[] | select(.toclevel == 1) | .line')
    readarray -t sections < <(echo "${response}" | jq -r '.parse.sections | .[] | select(.toclevel == 1) | .line')

    # todo: use global substitution to replace spaces
    # | gsub(" "; "_")
    
    for i in "${!sections[@]}"; do
      if [[ "${sections[$i]}" = "${section_title}" ]]; then
        index=$((i+1))
        echo ""
        echo "${sections[$i]}:"
        echo ""
        section_detail_url="https://en.wikipedia.org/w/api.php?action=parse&format=json&page=${title}&prop=text&section=${index}&disabletoc=1"
        response=$(curl -s -X GET "${section_detail_url}")
        detail=$(echo "${response}" | jq '.parse.text."*"')
        
        # NICE: this works
        echo "${detail//\\}" | \
        gsed -n 's:.*<p>\(.*\)</p>.*:\1:p' | \
        gsed -e 's/<[^>]*>//g' | \
        lynx -stdin -dump
        # or use the following
        # sed 's/&#160;//g' | \
        # sed 's/&#91;//g' | \
        # sed 's/2&#93;//g'
      fi
    done

    exit 0
fi

