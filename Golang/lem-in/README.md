# lem-in
Lem-in is a program lem-in that will read from a file (describing the ants and the colony) given in the arguments. The goal of the program is to find the quickest way to get `n` ants across a colony (composed of rooms and tunnels).
## Usage
```sh
$ go run . <filename.txt>
```
If valid path(s) is found, program will print out content of the text file given in arguments and the moves ants need to take.
## Testing
There is a test file included with the program which you can run with
```sh
$ ./test.sh
```
You can use questions from [lem-in/audit](https://github.com/01-edu/public/tree/master/subjects/lem-in/audit) to test if the program is working correctly.
## Authors
[Margus Toompere](https://01.kood.tech/git/MargusT)

[Ingvar Leerimaa](https://01.kood.tech/git/IngvarLeerimaa)