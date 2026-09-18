; Highlights for Monkey
;
; Neovim capture resolution: when multiple captures match the same node,
; the one with the highest #priority wins (default 100); at equal priority
; the capture defined LATER in the file wins. Specific rules therefore
; carry an explicit higher priority, and the builtin rule comes AFTER the
; generic call rule so it always wins over both @function.call and the
; catch-all (identifier) @variable at the bottom.
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

(escape_sequence) @string.escape

(string) @string

; User-defined function calls
; Monkey has no named function definitions, so @function is unused;
; callees are `let`-bound variables, hence @function.call.
((call_expression
  function: (identifier) @function.call)
  (#set! priority 120))

; Built-in function calls
; A builtin is also a call, so this rule comes AFTER @function.call and
; gets the highest priority: it must win over @function.call (120) and the
; catch-all @variable (100) below.
((call_expression
  function: (identifier) @function.builtin
  (#any-of? @function.builtin "len" "first" "last" "rest" "push" "puts"
    "readFile" "splitString" "atoi" "sortInts" "trimSpace" "trimSuffix"))
  (#set! priority 130))

; Function parameters
((function_literal
  parameters: (parameter_list
    (identifier) @variable.parameter))
  (#set! priority 120))

; Hash keys are properties (strings and bare identifiers)
((hash_pair
  key: (string) @property)
  (#set! priority 120))

((hash_pair
  key: (identifier) @property)
  (#set! priority 120))

; All-caps identifiers are constants by convention
((identifier) @constant
  (#match? @constant "^[A-Z][A-Z_]*$")
  (#set! priority 120))

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
