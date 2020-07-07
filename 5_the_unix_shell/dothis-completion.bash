#!/usr/bin/env bash

# _dothis_completions()
# {
#   if [ "${#COMP_WORDS[@]}" != "2" ]; then
#     return
#   fi

#   COMPREPLY=($(compgen -W "$(fc -l -50 | sed 's/\t//')" -- "${COMP_WORDS[1]}"))
# }

# complete -F _dothis_completions dothis



#/usr/bin/env bash
_dothis_completions()
{
  COMPREPLY+=("now")
  COMPREPLY+=("tomorrow")
  COMPREPLY+=("never")
}

complete -F _dothis_completions wiki