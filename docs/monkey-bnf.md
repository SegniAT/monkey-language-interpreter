# Monkey Language BNF Grammar Specification

This document describes the Monkey language as implemented in this
repository. Where this document and the code disagree, the code wins.

> **Executable spec**: the grammar is also mechanically verified by
> `tree-sitter-monkey/` (grammar.js) and its 70 corpus tests in
> `tree-sitter-monkey/test/corpus/`.

## Lexical Structure

```bnf
letter      ::= "a" | "b" | ... | "z" | "A" | ... | "Z" | "_"
digit       ::= "0" | "1" | ... | "9"
identifier  ::= letter+                        (* no digits allowed! *)
integer     ::= digit+
string      ::= '"' (any-character-except-quote-or-backslash | escape)* '"'
escape      ::= '\' ('n' | 't' | 'r' | '"' | '\')

whitespace  ::= " " | "\t" | "\n" | "\r"
comment     ::= (none - Monkey has no comments)
```

Notes:

- Identifiers consist solely of letters and `_` - `add5`, `foo2` etc. are
  **not** valid identifiers. `add5` lexes as the identifier `add` followed
  by the integer `5`.
- String literals are opened by `"` and closed by the next `"`. A backslash
  starts an escape sequence: `\n` (newline), `\t` (tab), `\r` (carriage
  return), `\"` (double quote) and `\\` (backslash). Any other `\x` sequence
  keeps the backslash as an ordinary character. An unterminated string runs
  to the end of input.
- Whitespace separates tokens; no comments exist.

## Program Structure

```bnf
program     ::= statement*

statement   ::= let_statement
             | return_statement
             | expression_statement
```

## Statements

```bnf
let_statement        ::= "let" identifier "=" expression ";?"
return_statement     ::= "return" expression ";?"
expression_statement ::= expression ";?"
```

Semicolons are **optional** at the end of any statement: the parser simply
skips a `;` when one is present.

## Expressions (by precedence, lowest to highest)

```bnf
expression          ::= equality_expression

equality_expression ::= comparison_expression (("==" | "!=") comparison_expression)*

comparison_expression ::= sum_expression (("<" | ">") sum_expression)*

sum_expression      ::= product_expression (("+" | "-") product_expression)*

product_expression  ::= prefix_expression (("*" | "/") prefix_expression)*

prefix_expression   ::= ("!" | "-") prefix_expression
                     | postfix_expression

postfix_expression  ::= primary_expression (
                         "(" expression_list? ")"    (* call: precedence 7 *)
                       )*
                     | primary_expression (
                         "[" expression "]"          (* index: precedence 8 *)
                       )*

primary_expression  ::= identifier
                      | integer
                      | string
                      | boolean
                      | array_literal
                      | hash_literal
                      | function_literal
                      | if_expression
                      | "(" expression ")"
```

Notes:

- `(` call and `[` index are **separate precedence levels** (7 and 8).
  Index binds tighter than call, so `add(1, 2)[0]` parses as
  `(add(1, 2))[0]` while `a[0](1)` parses as `(a[0])(1)`.
- There is **no assignment operator** and no compound operators. The `=`
  token appears only in `let` statements.
- Prefix operators apply to a single operand: `-5 * 3` is `(-5) * 3`,
  `!(x == y)` is `!(x == y)`.
- The unary chain and the call/index chain are left-associative.

## Literals

```bnf
boolean       ::= "true" | "false"

array_literal ::= "[" expression_list? "]"
expression_list ::= expression ("," expression)*

hash_literal  ::= "{" hash_pair_list? "}"
hash_pair_list ::= hash_pair ("," hash_pair)* ","?
hash_pair     ::= expression ":" expression

function_literal ::= "fn" "(" parameter_list? ")" "{" statement* "}"
parameter_list   ::= identifier ("," identifier)*
```

Notes:

- Hash keys are **any expression**. At run time only strings, integers,
  and booleans can be used as keys; using any other value raises
  `unusable as hash key: <type>`.
- A **trailing comma** after the last hash pair is allowed:
  `{"a": 1, "b": 2,}`.
- Function parameters are plain identifiers; `fn()` with no parameters is
  valid.

## Control Flow

```bnf
if_expression ::= "if" "(" expression ")" "{" statement* "}"
                  ("else" "{" statement* "}")?
```

**Monkey does NOT support**:

- `else if` — nest instead: `else { if (...) { ... } }`
- `for` loops
- `while` loops — **no loops exist at all**; use recursion
- Assignment (`i = i + 1`)
- The ternary operator `? :`
- `switch`/`case` statements
- Comments

## Built-in Functions (runtime, not syntax)

Built-ins are ordinary call expressions; the evaluator dispatches them at
run time. Unknown identifiers in call position are an error.

| Function | Arguments | Behavior |
|----------|-----------|----------|
| `len`      | 1: array or string | Length of array or string |
| `first`    | 1: array | First element (or `null` if empty) |
| `last`     | 1: array | Last element (or `null` if empty) |
| `rest`     | 1: array | New array without the first element |
| `push`     | 2: array, value | New array with value appended |
| `puts`     | any number | Prints each argument, one per line |
| `readFile` | 1: string path | Contents of the file as a string |
| `splitString` | 2: string, separator | Array of strings split by separator |
| `atoi`     | 1: string | Integer parsed from the string |
| `sortInts` | 1: array of integers | New array sorted in ascending order |
| `trimSpace` | 1: string | String without leading/trailing whitespace |
| `trimSuffix` | 2: string, suffix | String with `suffix` removed from the end |

## Valid Monkey Code Examples

```monkey
let x = 5;
let name = "Monkey";
let isReady = true;

let add = fn(a, b) { return a + b; };
let multiply = fn(x, y) { x * y; };

let result = add(3, 4);

let arr = [1, 2, 3, 4];
let firstElement = first(arr);
let restElements = rest(arr);
let newArr = push(arr, 5);

let person = {"name": "Alice", "age": 30};
let nameVal = person["name"];

let outcome = if (x > 10) { "big" } else { "small" };

let makeAdder = fn(x) {
  return fn(y) { return x + y; };
};
let addFive = makeAdder(5);
addFive(10);

let fib = fn(n) {
  if (n < 2) {
    return n;
  }
  return fib(n - 1) + fib(n - 2);
};
```

## Common Syntax Rules

1. **Semicolons are OPTIONAL** at the end of statements
2. **No comments** - `//`, `/* */`, `#` are all unsupported
3. **No loops** - no `for`, no `while`; recurse instead
4. **No `else if`** - use `else { if (...) { } }`
5. **No assignment expressions** - `=` appears only in `let`
6. **Hash keys are expressions** - any expression is syntactically valid;
   strings, integers, and booleans are usable at run time
7. **Functions are values** - assignable to variables, passable as arguments
8. **Everything is an expression** - `if` and function literals produce values

## Error Patterns to Avoid

```monkey
let x = 5;                    ok
let x = 5                     ok (semicolon optional)

for (i = 0; i < 10; i++) {}   error - no for loops
while (i < 10) {}             error - no while loops
i = i + 1                     error - no assignment expressions

let add5 = fn(x) { x + 5; };  error - add5 contains a digit
let add5 = ...                use: let addFive = fn(x) { x + 5; };

let h = {name: "Alice"};      ok - identifier key, evaluates "name"
let h = {"name": "Alice"};    ok - string key

arr.length                    error - no dot notation
len(arr)                      ok

if (n <= 1) { ... }           error - no <= operator
if (n < 2) { ... }            ok
```

## Token Reference

```
Keywords:   let, fn, if, else, return, true, false
Operators:  +, -, *, /, ==, !=, <, >, !, =
Delimiters: (, ), {, }, [, ], ,, ;, :
Identifiers: [a-zA-Z_]+            (no digits)
Integers:   [0-9]+
Strings:    "..." (double quotes only, escapes: \n \t \r \" \\)
```
