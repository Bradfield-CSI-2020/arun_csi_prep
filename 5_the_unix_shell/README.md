### Research

Search route
```
https://en.wikipedia.org/w/api.php?action=opensearch&search=Walrus&limit=3&namespace=0&format=json
```


Summary route
```
https://en.wikipedia.org/api/rest_v1/page/summary/Walrus
```

Section Titles from TOC
```
https://en.wikipedia.org/w/api.php?action=parse&format=json&page=Walrus&prop=sections&disabletoc=1
```


Section Detail
```
https://en.wikipedia.org/w/api.php?action=parse&format=json&page=Walrus&prop=text&section=1&disabletoc=1
```


section details but while using section tile as id
but does not work well
```
https://en.wikipedia.org/w/api.php?action=parse&format=json&page=house&prop=text&sectiontitle=Etymology&disabletoc=1&sectionpreview=1
```

mobile sections
```
https://en.wikipedia.org/api/rest_v1/page/mobile-sections/Walrus
```


### Reading List

1. https://www.shellscript.sh/tips/spinner/
2. https://gist.github.com/mohanpedala/1e2ff5661761d3abd0385e8223e16425
3. https://iridakos.com/programming/2018/03/01/bash-programmable-completion-tutorial

### TODO:

- [ ] spinner while waiting for requests?
- [ ] use gnu sed, awk etc
- [ ] use IFS=\n for array set


### Details

1. had to upgrade to bash 5
2. had to install jq
3. had to install lynx
4. had enable bash completion for zsh
```
autoload -U +X compinit && compinit
autoload -U +X bashcompinit && bashcompinit
```