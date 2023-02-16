# 👉uwu👈 twanspiwer

# weact to react
Package that uses uwu keywords as an alternative syntax for React.

Creating a transpiler to convert regular code to uwu language would require a lot of complex programming and natural language processing. 

# Golang needs to be installed

# Usage:

# clone project 

`git clone https://github.com/4cecoder/twanspiwer` or `gh repo clone 4cecoder/twanspiwer` 

## `cd twanspiwer && go mod tidy`

## Wun the twanspiwer
```shell
 go run main.go -input=react-code-with-uwu.js -output=react-code-without-uwu.js
```


To transpile a React code with uwu keywords using the key-value pairs in a `uwu.json` file. 

This code reads the `uwu.json` file and stores its contents in a Go map. 

It then reads the React code from a `react-code.js` file, 

replaces all the uwu keywords with the React keywords using `strings.ReplaceAll`, 

and then writes the transpiled code to a new file named `transpiled-code.js`.
