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

The db file is already in the repo, so NO need to do this.
This takes a while to run.

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


### Open Repl

```
bin/xkcd
```

Calling the executable without an arg will open a repl that can use used for many look ups