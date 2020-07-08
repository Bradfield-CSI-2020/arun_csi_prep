#!/usr/bin/env bash
title=$1
section_title=$2

if [ $# -eq 0 ]
  then
    echo "No arguments supplied"
    exit 1
fi

# Get search hits
url="https://en.wikipedia.org/w/api.php?action=opensearch&search=${title}&limit=3&namespace=0&format=json"
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
echo "${summary}"


if [ $# -eq 2 ]
  then
    echo ""
    echo "sections:"
    echo ""
    # get sections titles
    url="https://en.wikipedia.org/w/api.php?action=parse&format=json&page=${title}&prop=sections&disabletoc=1"
    response=$(curl -s -X GET ${url})
    sections=$(echo "${response}" | jq '.parse.sections' | jq '.[] | select(.toclevel == 1) | .line')
    sections_index=$(echo "${response}" | jq '.parse.sections' | jq '.[] | select(.toclevel == 1) | .number')

    # echo "${sections_index}"
    # echo "${sections}"

    # hash[${sections_index[i]}]=${sections[i]}


    # for i in "${sections_index[@]}"; do 
    #     echo "${hash[$i]}":; 
    # done

    section_detail_url="https://en.wikipedia.org/w/api.php?action=parse&format=json&page=${title}&prop=text&section=1&disabletoc=1"
    response=$(curl -s -X GET "${section_detail_url}")
    detail=$(echo "${response}" | jq '.parse.text."*"')

    # todo:
    # the html returned has all the details we need
    # the section title and the subsection have this before them `class="mw-headline"`
    # the section details can be found by looking for the first <p>
    # we can then look for the first ". " ? maybe...
    # but we need to strip all html in there

    echo "${detail}"
    
    # | map(select(.toclevel == 1))
    # | map(select(.toclevel = 1)) |.line' 
    exit 0
fi


# get sections titles
# url="https://en.wikipedia.org/w/api.php?action=parse&format=json&page=${search_term}&prop=sections&disabletoc=1"
# curl -X GET ${url} | jq '.parse.sections' | jq '.[] | .line' 

# section_title=$2
# section_detail="https://en.wikipedia.org/w/api.php?action=parse&format=json&page=${search_term}&prop=text&sectiontitle=${section_title}&disabletoc=1"
# echo $section_detail
# # curl -X GET ${section_detail} | jq '.parse.text."*"' | grep -o '<p>.*</p>' | sed 's/\(<p>\|<\/p>\)//g'
# curl -X GET ${section_detail} | jq '.parse.text."*"' > example.html


# better section detail 
# https://en.wikipedia.org/w/api.php?action=parse&format=json&page=Walrus&prop=text&section=3&disabletoc=1



# curl -sL https://www.iana.org/ | xargs | egrep -o "<tr>.*?</tr>"
