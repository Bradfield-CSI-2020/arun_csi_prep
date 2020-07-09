### Research

Search route
```
https://en.wikipedia.org/w/api.php?action=opensearch&search=${title}&limit=${num}&namespace=0&format=json
```


Summary route
```
https://en.wikipedia.org/api/rest_v1/page/summary/${title}
```

Section Titles from TOC
```
https://en.wikipedia.org/w/api.php?action=parse&format=json&page=${title}&prop=sections&disabletoc=1
```


Section Detail
```
https://en.wikipedia.org/w/api.php?action=parse&format=json&page=${title}&prop=text&section=1&disabletoc=1
```