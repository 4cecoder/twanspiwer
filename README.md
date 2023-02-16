# 👉uwu👈 Twanspiwer

A package that uses uwu keywords as an alternative syntax for React.

Creating a transpiler to convert regular code to uwu language would require a lot of complex programming and natural language processing.

## Usage:

### Pwewequisites

Golang needs to be installed.

### Installation

Cwone the project:

`git clone https://github.com/4cecoder/twanspiwer` or `gh repo clone 4cecoder/twanspiwer` 

## get depenwencies

## `cd twanspiwer && go mod tidy`

## Wun the twanspiwer
```shell
 go run main.go -input=react-code-with-uwu.js -output=react-code-without-uwu.js
```

## How it works

To transpile a React code with uwu keywords using the key-value pairs in a `uwu.json` file:

1. Read the `uwu.json` file and store its contents in a Go map.
2. Read the React code from a `react-code.js` file.
3. Replace all the uwu keywords with the React keywords using `strings.ReplaceAll`.
4. Write the transpiled code to a new file named `transpiled-code.js`.

## Running the transpiler

```shell
go run main.go -input=react-code-with-uwu.js -output=react-code-without-uwu.js
```
