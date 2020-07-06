# The GO Programming Language

## TODO:
- [ ] Move code into packages, index, scrape etc
- [ ] Write tests


## Getting Started

### Build 

```
go build -o bin/xkcd
```

### Scrape XKCD

```
bin/xkcd --index
```

This will create a file named `db.jsonl` and write responses from xkcd to this file.
The program calculate the url for each comic and then fetches it fom xkcd.
By default it will fetch all available comics as of July 6th.


### Search for terms

```
bin/xkcd space
```

```
bin/xkcd "space elevators"
```
