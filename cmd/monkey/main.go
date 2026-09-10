package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/SegniAT/monkey-language-interpreter/evaluator"
	"github.com/SegniAT/monkey-language-interpreter/lexer"
	"github.com/SegniAT/monkey-language-interpreter/lsp/lsp/server"
	"github.com/SegniAT/monkey-language-interpreter/object"
	"github.com/SegniAT/monkey-language-interpreter/parser"
	"github.com/SegniAT/monkey-language-interpreter/repl"
)

func main() {
	if len(os.Args) == 1 {
		repl.Start(os.Stdin, os.Stdout)
		return
	}

	switch os.Args[1] {
	case "lsp":
		runLSP()
	case "run":
		if len(os.Args) < 3 {
			fmt.Println("Usage: monkey run <file.monkey>")
			os.Exit(1)
		}
		runFile(os.Args[2])
	default:
		// Fallback: assume the user just passed a filename without "run"
		runFile(os.Args[1])
	}
}

func runFile(filename string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err.Error())
		os.Exit(1)
	}

	env := object.NewEnvironment()
	l := lexer.New(string(content))
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Diagnostics()) != 0 {
		repl.PrintParserDiagnostics(os.Stdout, filename, p.Diagnostics())
		os.Exit(1)
	}

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil && evaluated.Type() == object.ERROR_OBJ {
		fmt.Printf("Runtime Error: %s\n", evaluated.Inspect())
		os.Exit(1)
	}
}

func runLSP() {
	file, err := os.OpenFile(filepath.Join(os.TempDir(), "monkey-lsp.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		panic(err)
	}

	slog.SetDefault(slog.New(slog.NewTextHandler(file, &slog.HandlerOptions{
		AddSource: false,
		Level:     slog.LevelDebug,
	})))

	srv := server.NewServer(os.Stdin, os.Stdout)
	slog.Info("Started LSP server")

	if err := srv.Run(); err != nil {
		log.Fatalf("Server stopped: %v", err)
	}
}
