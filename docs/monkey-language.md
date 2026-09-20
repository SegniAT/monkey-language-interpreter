# Monkey Programming Language Reference

## Overview

Monkey is a dynamically-typed programming language with a C-like syntax. It features first-class functions, closures, arrays, hashes, and built-in utility functions.

The language and its interpreter are based on [Writing An Interpreter In Go](https://interpreterbook.com/) by Thorsten Ball.

---

## Token Types

### Keywords

| Token | Description |
|-------|-------------|
| `fn` | Function declaration |
| `let` | Variable binding |
| `if` | Conditional expression |
| `else` | Alternative branch in conditional |
| `return` | Return value from function |
| `true` | Boolean true value |
| `false` | Boolean false value |

### Operators

| Operator | Description |
|----------|-------------|
| `+` | Addition |
| `-` | Subtraction |
| `*` | Multiplication |
| `/` | Division |
| `==` | Equality comparison |
| `!=` | Inequality comparison |
| `<` | Less than |
| `>` | Greater than |
| `!` | Logical negation (prefix) |

### Delimiters

| Token | Description |
|-------|-------------|
| `(` `)` | Function calls, grouped expressions |
| `{` `}` | Function body, block expressions |
| `[` `]` | Array literals and indexing |
| `,` | List separator |
| `;` | Statement terminator |
| `:` | Hash key-value separator |

### Data Types

| Type | Description |
|------|-------------|
| `ident` | Identifier (variable/function names) |
| `int` | Integer literal (e.g., `42`, `100`) |
| `string` | String literal (e.g., `"hello"`) |
| `boolean` | Boolean literal (`true` / `false`) |
| `array` | Array literal (e.g., `[1, 2, 3]`) |
| `hash` | Hash literal (e.g., `{"key": "value"}`) |
| `function` | Function literal (`fn(x) { ... }`) |

### Escape Sequences

String literals are wrapped in double quotes. A backslash inside a string
starts an escape sequence:

| Escape | Meaning |
|--------|---------|
| `\n` | Newline (LF) |
| `\t` | Tab |
| `\r` | Carriage return |
| `\"` | Double quote |
| `\\` | Backslash |

Any other `\x` sequence keeps the backslash as a literal character.

```monkey
let multiLine = "line one\nline two";
let path = "C:\\monkey";
```

---

## Syntax Examples

### Variables

```monkey
let x = 5;
let name = "Monkey";
let isActive = true;
```

### Functions

```monkey
let greet = fn() { puts("Hello!"); };

let add = fn(a, b) { return a + b; };

let fib = fn(n) {
    if (n < 2) {
        return n;
    }
    return fib(n - 1) + fib(n - 2);
};
```

Note: Monkey has no `<=` operator - only `<`, `>`, `==`, and `!=`.

### Conditionals

```monkey
let result = if (x > 10) { "big" } else { "small" };
```

### Arrays

```monkey
let numbers = [1, 2, 3, 4, 5];
let firstEl = numbers[0];
let lastEl = numbers[4];
```

### Hashes (Dictionaries)

```monkey
let person = {
    "name": "Alice",
    "age": 30,
    "city": "Boston"
};

let name = person["name"];
```

### Closures

```monkey
let newAdder = fn(x) {
    return fn(y) { return x + y; };
};

let addFive = newAdder(5);
addFive(10);
```

(`addFive(10)` returns `15`.)

### Higher-Order Functions

```monkey
let map = fn(arr, f) {
    let iter = fn(arr, acc) {
        if (len(arr) == 0) {
            return acc;
        }
        return iter(rest(arr), push(acc, f(first(arr))));
    };
    return iter(arr, []);
};

map([1, 2, 3], fn(x) { return x * 2; });
```

(`map([1, 2, 3], ...)` returns `[2, 4, 6]`.)

---

## Built-in Functions

| Function | Description | Example |
|----------|-------------|---------|
| `len(array)` | Returns length of array or string | `len([1,2,3])` → `3` |
| `first(array)` | Returns first element | `first([1,2,3])` → `1` |
| `last(array)` | Returns last element | `last([1,2,3])` → `3` |
| `rest(array)` | Returns array without first element | `rest([1,2,3])` → `[2,3]` |
| `push(array, value)` | Appends value to array (returns new array) | `push([1,2], 3)` → `[1,2,3]` |
| `puts(...)` | Prints each argument, one per line | `puts("hello")` |
| `readFile(path)` | Reads the file at `path` and returns its contents as a string | `readFile("input.txt")` |
| `splitString(str, sep)` | Splits `str` by `sep` and returns an array of strings | `splitString("a,b,c", ",")` → `["a","b","c"]` |
| `atoi(str)` | Parses a string into an integer | `atoi("42")` → `42` |
| `sortInts(array)` | Sorts an array of integers in ascending order (returns a new array) | `sortInts([3, 1, 2])` → `[1, 2, 3]` |
| `trimSpace(str)` | Removes leading and trailing whitespace from a string | `trimSpace("  hi  ")` → `"hi"` |
| `trimSuffix(str, suffix)` | Removes `suffix` from the end of `str` | `trimSuffix("hello.", ".")` → `"hello"` |

The file-I/O and string built-ins (`readFile`, `splitString`, `atoi`,
`sortInts`, `trimSpace`, `trimSuffix`) are handy for data-processing
scripts - see [Monkey BNF Grammar](./monkey-bnf.md) for the strict lexical
rules, and `examples/AoC_2024/` for a data-crunching example.

---

## Parsing Pipeline

1. **Lexer** (`lexer.New(input)`) → tokenizes input string into stream of `Token{Type, Literal}`
2. **Parser** (`parser.New(lexer)`) → parses tokens into AST (`Program` → `Statement` → `Expression`)
3. **Evaluator** (`evaluator.Eval(program, env)`) → evaluates AST and returns `Object`

Common usage:

```go
l := lexer.New(code)
p := parser.New(l)
program := p.ParseProgram()

if len(p.Diagnostics()) > 0 {
    // Handle syntax errors
}

result := evaluator.Eval(program, env)
```

---

## Error Messages

Common parser errors indicate:
- Unexpected tokens
- Incorrect syntax (unmatched brackets, missing `=` in `let`)
- Invalid prefix/infix expressions

The evaluator can produce runtime errors:
- Type errors (e.g., calling non-function)
- Unknown built-in function
- Index out of bounds

---

## Summary

Monkey is a minimal, educational language demonstrating:
- Lexical analysis (tokenization)
- Recursive descent parsing (AST generation)
- AST evaluation (interpretation)
- First-class functions and closures
- Built-in higher-order functions
