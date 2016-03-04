# A word search example written in golang

This is just a quick example implementation to a word search assignment.

Some things to note:
  - There are quite a few nested loops, so efficiency could be improved.
  - The sanitation of the input could be improved

### Version
0.1

### Usage
You need the go runtime installed, and your GOPATH configured. Afterwards you can run:
```sh
$ go run *.go puzzle.csv words.csv out.csv
```
>Arguments:
* _*puzzle.csv*_ the puzzle in csv format.
* _*words.csv*_ all words to search for [including the not available words.
* _*out.csv*_ the output containing just the words that were actually found.
