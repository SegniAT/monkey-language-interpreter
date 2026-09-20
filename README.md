<p align="center">
  <img src="assets/mascot.png" alt="Monkey language mascot, Punch, the sensetive young monkey." width="240">
</p>

<h1 align="center">Monkey Programming Language</h1>

<p align="center">
  A tree-walking interpreter written in Go, complete with a language server (LSP)
  and a <code>tree-sitter</code> grammar for first-class editor support.
</p>

**Monkey** is a minimal, dynamically-typed language with a C-like syntax. It is
the result of building a real interpreter from scratch - a lexer, a Pratt
(recursive-descent) parser, an AST, and a tree-walking evaluator - plus the
developer-experience layer around it: diagnostics, hover, completions, and
go-to-definition, and a full `tree-sitter` grammar with syntax highlighting,
folding, and indentation queries.

The interpreter core follows [Writing An Interpreter In Go](https://interpreterbook.com/)
by Thorsten Ball.

## Features

- First-class functions and closures
- Arrays and hashes (dictionaries) as native data types
- Everything is an expression: `if`, function literals, and calls all produce values
- **No loops and no assignment** - recursion is the only iteration primitive
- **12 built-in functions**, including string/file helpers (`readFile`,
  `splitString`, `atoi`, `sortInts`, `trimSpace`, `trimSuffix`)
- A Language Server Protocol (`LSP`) implementation with semantic diagnostics
- A `tree-sitter` grammar with captures for highlighting, folds, and indents

## Requirements

- Go 1.24 or newer (`go.mod` declares `go 1.24`)

## Getting Started

Build the `monkey` executable:

```bash
go build -o bin/monkey ./cmd/monkey
```

### REPL

Run with no arguments to enter the interactive REPL:

```bash
./bin/monkey
```

```
>> let add = fn(a, b) { a + b; };
>> add(2, 3)
5
```

### Running a script

```bash
./bin/monkey run examples/AoC_2024/day1.monkey
# `run` is optional - passing a bare filename works too:
./bin/monkey examples/AoC_2024/day1.monkey
```

> Note: `readFile` resolves paths relative to the **current working directory**.
> The AoC examples read `input.txt` from the same folder, so run them from
> inside it:
>
> ```bash
> cd examples/AoC_2024 && ../../bin/monkey day1.monkey
> ```

### Quick tour

```monkey
let fib = fn(n) {
  if (n < 2) { return n; }
  return fib(n - 1) + fib(n - 2);
};

let map = fn(arr, f) {
  let iter = fn(arr, acc) {
    if (len(arr) == 0) { return acc; }
    return iter(rest(arr), push(acc, f(first(arr))));
  };
  return iter(arr, []);
};

map([1, 2, 3], fn(x) { return x * 2; });
```

(`map(...)` above yields `[2, 4, 6]`.)

Full language reference and exact grammar:

- [`docs/monkey-language.md`](docs/monkey-language.md) - the user-facing reference
- [`docs/monkey-bnf.md`](docs/monkey-bnf.md) - the formal BNF grammar
  (mechanically verified by the `tree-sitter-monkey` corpus tests)

## Repository Layout

| Directory | Purpose |
|-----------|---------|
| `lexer/` | Turns source text into tokens (with string escape sequences) |
| `parser/` | Pratt parser producing the AST |
| `ast/` | Abstract syntax tree node definitions |
| `evaluator/` | Tree-walking evaluator and built-in functions |
| `object/` | Runtime object types |
| `token/` | Token type definitions |
| `repl/` | Interactive read-eval-print loop |
| `cmd/monkey/` | CLI entry point (`monkey`, `monkey run <file>`, `monkey lsp`) |
| `lsp/` | Language Server Protocol implementation |
| `tree-sitter-monkey/` | The `tree-sitter` grammar, queries, and corpus tests |
| `docs/` | Language reference and BNF grammar |
| `examples/` | Runnable example scripts (`AoC_2024/`) |

## Language Server (LSP)

The `monkey` binary speaks LSP over stdio. It currently provides:

- **Diagnostics** - parse errors plus semantic analysis: redeclared variables,
  undefined variables, and unused-variable warnings
- **Hover** - details for variables (kind + definition site), literals
  (integer hex/octal/binary, string length/size, booleans), and full docs for
  built-in functions
- **Go to definition** - jump from any identifier use to its `let` binding or parameter definition
- **Completions** - identifiers in scope (variables, parameters, built-ins)
  and keywords, prefix-matched against the symbol under the cursor

### Starting the server

```bash
./bin/monkey lsp
```

The server speaks JSON-RPC 2.0 over `stdin`/`stdout`, so any LSP client works.
Logs are appended to `/tmp/monkey-lsp.log` (set with `slog` to debug level).

### Neovim (`vim.lsp.config`)

Register the hand-built server with Neovim's native LSP framework (there is no
bundled `monkey` server or `nvim-lspconfig` config for it), map the filetype,
then enable it:

```lua
vim.filetype.add({ extension = { monkey = "monkey" } })

local servers = {
  monkey = {
    name = "monkey-lsp",
    cmd = { "/absolute/path/to/bin/monkey", "lsp" },
    filetypes = { "monkey" },
  },
  -- your other servers...
}

for name, server in pairs(servers) do
  vim.lsp.config(name, server)
  vim.lsp.enable(name)
end
```

If you use Mason, exclude the custom server from auto-installation (in both the
installer list and the `default_setup` handlers):

```lua
local ensure_installed = vim.tbl_filter(function(name)
  return name ~= "monkey"
end, vim.tbl_keys(servers))

mason_lspconfig.setup {
  ensure_installed = ensure_installed,
  handlers = {
    function(name)
      if name ~= "monkey" then require('mason-lspconfig').default_setup(name) end
    end,
  },
}
```

## Tree-sitter & Syntax Highlighting

The grammar lives in [`tree-sitter-monkey/`](tree-sitter-monkey/) with a
`tree-sitter.json` manifest, node/c bindings, and three queries:

| Query | Purpose |
|-------|---------|
| `queries/highlights.scm` | Syntax highlighting captures |
| `queries/folds.scm` | Fold ranges for `block`, `function_literal`, `if_expression`, hashes, arrays, and call argument lists |
| `queries/indents.scm` | Indentation after multiline opening delimiters |

### Installing with nvim-treesitter

Register the grammar as a custom parser, then install it:

```lua
-- nvim-treesitter wipes the `parsers` table on every install/update and
-- re-fires `User TSUpdate`, so re-apply the registration in an autocmd:
vim.api.nvim_create_autocmd('User', {
  pattern = 'TSUpdate',
  callback = function()
    require('nvim-treesitter.parsers').monkey = {
      install_info = {
        url = 'https://github.com/SegniAT/monkey-language-interpreter',
        location = 'tree-sitter-monkey',
        queries = 'tree-sitter-monkey/queries',
      },
    }
  end,
})

require('nvim-treesitter').install({ 'monkey', ... })
```

Then add `monkey` to the highlighted languages (or just edit a `.monkey` file —
`vim.filetype.add` above will attach the parser automatically).

### Highlights

Built-in function calls render as `@function.builtin`, user-defined calls as
`@function.call`, all-caps identifiers as `@constant`, hash keys as
`@property`, parameter lists as `@variable.parameter`, and everything else as
`@variable`. Keywords, booleans, numbers, strings (including escapes), and
operators get the standard captures (`@keyword`, `@boolean`, `@number`,
`@string`, `@string.escape`, `@operator`, ...).

**Capture resolution rule** (matters if you tweak the queries): when several
patterns capture the same node, the one with the highest `#priority` wins
(default is 100); at equal priority the capture defined **later** in the file
wins. The builtin rule therefore carries `priority 130` and is placed *after*
the generic `@function.call` rule, so built-ins always beat both it and the
catch-all `@variable`.

Reloading: queries are re-read when the buffer reopens, so after a queries
change simply run `:edit`. After a **grammar** change, reinstall the parser.

## Examples

- [`examples/AoC_2024/day1.monkey`](examples/AoC_2024/day1.monkey) — Advent of
  Code 2024 Day 1: total distance between two sorted lists.
- [`examples/AoC_2024/day2.monkey`](examples/AoC_2024/day2.monkey) — AoC 2024
  Day 1 part 2: a similarity score, counting occurrences.

Both solve the problem recursively (Monkey has no loops) using the string and
array built-ins: `readFile`, `trimSuffix`, `splitString`, `atoi`, `sortInts`,
`len`, `push`, and `first`/`rest`, plus integer indexing like `list[i]`.

## Testing

```bash
go test ./...
go vet ./...
```

The parser grammar is verified end-to-end too: `tree-sitter-monkey/test/corpus/`
holds 10 corpus files of valid and invalid snippets that double as an executable
specification of `docs/monkey-bnf.md`. Run them from `tree-sitter-monkey/` with
the tree-sitter CLI:

```bash
npm install
tree-sitter test
```

## Acknowledgements

Monkey is based on [Writing An Interpreter In Go](https://interpreterbook.com/)
by Thorsten Ball. The pipeline - lexer, Pratt parser, AST, and tree-walking
evaluator - and the language's functional, expression-oriented style follow
the book, extended here with 12 built-ins, a language server (LSP), and a
`tree-sitter` grammar.

## License

MIT — see [LICENSE](LICENSE). The interpreter, the `monkey-lsp` server, and
the `tree-sitter-monkey` grammar are all released under the MIT license.
