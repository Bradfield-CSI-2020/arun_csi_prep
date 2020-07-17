_wiki_completions()
{
  if [ ${#COMP_WORDS[@]} -lt 2 ]; then
    return
  fi
  search_url="https://en.wikipedia.org/w/api.php?action=opensearch&search=${COMP_WORDS[1]}&limit=1&namespace=0&format=json"
  response=$(curl -s -L -X GET "${search_url}")
  results=$(echo "${response}" | jq -r '.[1] | join(" ")')

  COMPREPLY=("$(compgen -W "${results}" -- "${COMP_WORDS[1]}")")
}
complete -F _wiki_completions wiki

