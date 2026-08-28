package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/SegniAT/monkey-language-interpreter/evaluator"
	"github.com/SegniAT/monkey-language-interpreter/lexer"
	"github.com/SegniAT/monkey-language-interpreter/object"
	"github.com/SegniAT/monkey-language-interpreter/parser"
	"github.com/SegniAT/monkey-language-interpreter/token"
)

const PROMT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Printf(PROMT)
		scanned := scanner.Scan()
		if !scanned {
			// Check if the scanner stopped due to an error rather than just EOF
			if err := scanner.Err(); err != nil {
				fmt.Fprintf(out, "Error reading input: %s\n", err)
			}

			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Diagnostics()) != 0 {
			printParserDiagnostics(out, p.Diagnostics())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

const MONKEY_FACE = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

func printParserDiagnostics(out io.Writer, diagnostics []token.Diagnostic) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser diagnostics:\n")
	for _, diag := range diagnostics {
		io.WriteString(out, "\t"+diag.Severity.String()+": "+diag.Message+"\n")
	}
}
