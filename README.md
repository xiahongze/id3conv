# id3conv: a commandline tool to fix encoding error in music files

Currently, **id3conv** supports only conversion from `GBK` to `UTF-8` assuming
the `ID3` tags were incorrectly encoded in `latin-1` or `ISO-8859-1` and decoded
as `UTF-8` on Unix-like systems.

## requirements

- go >= 1.8
- dep >= 0.5
  
## installation

1. clone this repo
2. `dep ensure`
3. `cd main && go build`
4. `cp main where_you_wanna_be/id3conv`
5. run it like `id3conv ${musicFile1} ${musicFile2}`
6. enjoy correct `UTF-8` meta info in your music

## integrate with **macOS** context menu with **automator**

1. open **automator**
2. create an empty service
3. set `workflow receives current` <- `audio files` in `Finder`
4. look for `run shell script` on the left menu and drag it in
5. input a script like this:

```bash
for f in "$@"
do
	/usr/local/bin/id3conv "$( echo "$f" | sed 's/ /\ /g' )"
done
```

Suppose you save your `main` binary as `/usr/local/bin/id3conv`.

6. save this service and you can use it in `Finder`