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

This interpreter was built by following the book "Writing an Interpreter in Go" by Thorsten Ball but has been modified slightly for my own purposes.
