# Go Interpreter

An tree walking interpreter written in Go. The interpreter is used for a made up language just used for the purposes of learning. 
The language is made up of these features:
- C-like syntax
- integers and booleans
- arithmetic expressions
- built-in functions
- first-class and higher-order functions
- closures
- a string data structure
- an array data structure
- a hash data structure

## Testing out the language in a REPL
To use the language and test out the interpreter you need to start the repl built for it.
As of 01 / 03 / 2018, the repl only reads, prints out the tokens read from the input, and loops.
Evaluating the statements will be completed soon.
Run the repl by:
1. Cloning the repository
2. Navigating to the repository
2. Running `go build`
3. Running `./go-interpreter` or run `go-interpreter.exe`, depending on which OS you are on

## Future Plans
The overall goal is to use what was learned in this exercise to create a language that is approachable
to kids. It will be a language like scratch but will allow you to actually type, instead of dragging and
dropping. It is meant to be a slight step up from scratch so students can get the feel of typing simple code statements. The plan is to couple this language in a web interface similar to scratch, game on one side, editor on the other. The game will have challenges that involve a character needing to move, jump, and break through obstacles. Moving, jumping, and breaking will be global functions part of the language so that students can use them alongside normal code statements to control a character to complete challenges and progress through the game.

The interpreter was built by following the book "Writing an Interpreter in Go" by Thorsten Ball but has been modified for my own purposes.