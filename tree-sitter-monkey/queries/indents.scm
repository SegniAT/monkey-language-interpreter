; Indents for Monkey
; @indent only fires when the node spans multiple lines, so single-line
; literals are not affected. Deliberately excludes if_expression and
; function_literal: their inner (block) would double-indent the body.

; Indent after multiline opening delimiters
[
  (block)
  (array_literal)
  (hash_literal)
  (paren_expression)
  (argument_list)
  (parameter_list)
] @indent

; Dedent on closing delimiters
[
  "}"
  ")"
  "]"
] @outdent
