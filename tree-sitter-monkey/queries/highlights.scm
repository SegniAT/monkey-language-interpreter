; Highlights for Monkey
; Order matters: the first matching pattern wins for a given node, so
; specific rules (builtins, call position, hash keys, constants) come
; before the catch-all (identifier) @variable at the bottom.
;
; Standard Nvim captures that are intentionally unused: Monkey has no
; comments, types, modules, labels, string escapes, or markup, so
; @comment, @type, @module, @label, @string.escape, @markup, and friends
; never apply.

; Keywords
"let" @keyword

[
  "if"
  "else"
] @keyword.conditional

"return" @keyword.return

"fn" @keyword.function

; Literals
[
  "true"
  "false"
] @boolean

(integer) @number

(string) @string

; Built-in function calls
(call_expression
  function: (identifier) @function.builtin
  (#any-of? @function.builtin "len" "first" "last" "rest" "push" "puts"))

; User-defined function calls
; Monkey has no named function definitions, so @function is unused;
; callees are `let`-bound variables, hence @function.call.
(call_expression
  function: (identifier) @function.call)

; Function parameters
(function_literal
  parameters: (parameter_list
    (identifier) @variable.parameter))

; Hash keys are properties (strings and bare identifiers)
(hash_pair
  key: (string) @property)

(hash_pair
  key: (identifier) @property)

; All-caps identifiers are constants by convention
((identifier) @constant
  (#match? @constant "^[A-Z][A-Z_]*$"))

; Operators
[
  "+"
  "-"
  "*"
  "/"
  "=="
  "!="
  "<"
  ">"
  "="
  "!"
] @operator

; Brackets
[
  "("
  ")"
  "{"
  "}"
  "["
  "]"
] @punctuation.bracket

; Delimiters
[
  ","
  ";"
  ":"
] @punctuation.delimiter

; Catch-all: every other identifier is a variable
(identifier) @variable
