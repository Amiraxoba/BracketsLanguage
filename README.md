# BracketsLanguage

**BracketsLanguage** is an esoteric programming language similar to Brainfuck and can compile to it natively.

## Numbers

- `()` = 1
- `((` = 10
- `)(` = 50

## Commands

- `> NUMBER` Increases the NUMBER of the value of the current cell
- `< NUMBER` Decreases the NUMBER of the value of the current cell
- `[ NUMBER` Increases the cell pointer
- `] NUMBER` Increases the Cell pointer
- `(((` Prints the Value of the current cell
- `())` Gets input

## Run your CustomDictionaries

`%BINARY` is the path to the binary file.

- #### Run the code using the cli
  
  > `%BINARY run ENTIRE_PATH_TO_YOUR_SCRIPT`

- #### Compile your code to Brainfuck
  
  > `%BINARY build ENTIRE_PATH_TO_YOUR_SCRIPT`

## Support

- [x] Linux
- [x] Windows
- [ ] MacOs

Request support for your operating system through GitHub issues.
