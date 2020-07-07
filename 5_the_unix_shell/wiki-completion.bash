#!/usr/bin/env bash

_wiki_completions()
{

  url="https://en.wikipedia.org/w/api.php?action=opensearch&search=${COMP_WORDS[1]}&limit=3&namespace=0&format=json"
  result=$(curl -s -X GET "${url}" | jq -r '.[1] | map(sub(" "; "_")) | join(" ")')
  COMPREPLY=("$(compgen -W "${result}" -- "${COMP_WORDS[1]}")")
}

complete -F _wiki_completions wiki


# _wiki_completions()
# {
#     # Get search hits
#     url="https://en.wikipedia.org/w/api.php?action=opensearch&search=${COMP_WORDS[0]}&limit=3&namespace=0&format=json"
#     result=$(curl -s -X GET ${url} | jq '.[1] | join(" ")')
#     COMPREPLY=($(compgen -W ${result} "${COMP_WORDS[1]}"))
# }

# complete -F _wiki_completions wiki