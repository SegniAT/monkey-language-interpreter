package benchmark

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/SegniAT/monkey-language-interpreter/evaluator"
	"github.com/SegniAT/monkey-language-interpreter/examples/AoC_2024/go_soln"
	"github.com/SegniAT/monkey-language-interpreter/lexer"
	"github.com/SegniAT/monkey-language-interpreter/object"
	"github.com/SegniAT/monkey-language-interpreter/parser"
)

func aocDir() string {
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		panic("benchmark: runtime.Caller(0) failed")
	}

	return filepath.Join(filepath.Dir(thisFile), "..")
}

func monkeyBench(b *testing.B, script string) {
	dir := aocDir()

	source, err := os.ReadFile(filepath.Join(dir, script))
	if err != nil {
		b.Fatalf("os.ReadFile(): %s", err.Error())
	}

	wd, err := os.Getwd()
	if err != nil {
		b.Fatalf("os.Getwd(): %s", err.Error())
	}

	if err := os.Chdir(dir); err != nil {
		b.Fatalf("os.Chdir(): %s", err.Error())
	}
	defer os.Chdir(wd)

	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		env := object.NewEnvironment()
		l := lexer.New(string(source))
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Diagnostics()) != 0 {
			b.Fatalf("parse errors in %s", script)
		}

		if evaluated := evaluator.Eval(program, env); evaluated != nil && evaluated.Type() == object.ERROR_OBJ {
			b.Fatalf("Runtime Error: %s", evaluated.Inspect())
		}
	}
}

func goBench(b *testing.B, solve func(col1, col2 []int) int) {
	inputPath := filepath.Join(aocDir(), "input.txt")

	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		col1, col2 := soln.Parse(inputPath)
		solve(col1, col2)
	}
}

func BenchmarkGoDay1Part1(b *testing.B) {
	goBench(b, soln.Part1)
}

func BenchmarkGoDay1Part2(b *testing.B) {
	goBench(b, soln.Part2)
}

func BenchmarkMonkeyDay1Part1(b *testing.B) {
	monkeyBench(b, "day1_part1.monkey")
}

func BenchmarkMonkeyDay1Part2(b *testing.B) {
	monkeyBench(b, "day1_part2.monkey")
}
