# TERTI-OPTIMIZER

## INTRO
Terti-optimizer is a program that receives a single input from the command line arguments—a text file containing tetrominos represented by #. The program arranges the tetrominos to form the smallest square possible.

## How to run 
1. clone this repo with the following command

```
git clone https://github.com/Siaka385/Terti-optimizier.git
```
2. Move into the repo
```
cd Terti-optimizier
```
3. Run the program with any of the provided files or any other file that contains valid tetrominos:
  - goodexample00.txt
  - goodexample01.txt
  - goodexample02.txt
  - goodexample03.txt
  - hardexample.txt
  - Badexample02.txt

Example one;
```
go run . hardexample.txt
```
it's output;
```
FFFBLLL
AAFBBL.
AAEBKGG
CCEEKKG
CCEJJKG
DDHHJII
DDHHJII
Time elapsed: 30.638791594s
```
Example two;invalid tertiminos
```
go run . Badexample02.txt
```
it's output;
```
ERROR
exit status 1
```

## RUles for valid Tetrominos

- Each tetromino must be represented by exactly 4 # characters.
- The tetrominos must form valid shapes such as L, T, Z, O, or I.
- Each tetromino should occupy a continuous block in the grid, with no disjointed parts.
- Invalid configurations or shapes will result in an error.

## why program will return Error
In case of bad format on the tetrominoes or bad file format  or invalid file name it should print ERROR

## Contribution
I welcome all contributors! Feel free to submit issues, pull requests, or suggestions to improve Terti-optimizer. Your contributions help make this project better.