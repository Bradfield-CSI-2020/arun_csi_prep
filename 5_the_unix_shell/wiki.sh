#!/usr/bin/env bash
search_term=$1

# Get search hits
url="https://en.wikipedia.org/w/api.php?action=opensearch&search=${search_term}&limit=3&namespace=0&format=json"
# echo "$url"
result=$(curl -s -X GET ${url} | jq -r '.[1] | map(sub(" "; "_")) | join(" ")')

# arr=($result)

compgen -W "$result"
# for i in ${arr[@]}; do echo $i; done

# todo: use jq map to replace space with string


# get summary
# url="https://en.wikipedia.org/api/rest_v1/page/summary/${search_term}"
# echo "$url"
# curl -X GET ${url} | jq '.extract'


# get sections titles
# url="https://en.wikipedia.org/w/api.php?action=parse&format=json&page=${search_term}&prop=sections&disabletoc=1"
# curl -X GET ${url} | jq '.parse.sections' | jq '.[] | .line' 

# section_title=$2
# section_detail="https://en.wikipedia.org/w/api.php?action=parse&format=json&page=${search_term}&prop=text&sectiontitle=${section_title}&disabletoc=1"
# echo $section_detail
# # curl -X GET ${section_detail} | jq '.parse.text."*"' | grep -o '<p>.*</p>' | sed 's/\(<p>\|<\/p>\)//g'
# curl -X GET ${section_detail} | jq '.parse.text."*"' > example.html





# curl -sL https://www.iana.org/ | xargs | egrep -o "<tr>.*?</tr>"
