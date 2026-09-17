package server

import (
	"testing"

	"github.com/SegniAT/monkey-language-interpreter/lsp/analysis"
	"github.com/SegniAT/monkey-language-interpreter/lsp/lsp/protocol"
)

func TestToProtocolDiagnosticsExclusiveEnd(t *testing.T) {
	tests := []struct {
		name        string
		content     string
		wantMessage string
		wantStart   protocol.Position
		wantEnd     protocol.Position
	}{
		{
			name:        "undefined multi-char identifier",
			content:     "let y = foo;",
			wantMessage: "undefined variable: foo",
			wantStart:   protocol.Position{Line: 0, Character: 8},
			wantEnd:     protocol.Position{Line: 0, Character: 11},
		},
		{
			name:        "undefined single-char identifier",
			content:     "let y = x;",
			wantMessage: "undefined variable: x",
			wantStart:   protocol.Position{Line: 0, Character: 8},
			wantEnd:     protocol.Position{Line: 0, Character: 9},
		},
		{
			name:        "unused variable",
			content:     "let longName = 5;",
			wantMessage: "unused variable: longName",
			wantStart:   protocol.Position{Line: 0, Character: 4},
			wantEnd:     protocol.Position{Line: 0, Character: 12},
		},
		{
			name:        "redeclaration",
			content:     "let a = 1;\nlet again = 2;\nlet again = 3;",
			wantMessage: "redeclaration of again",
			wantStart:   protocol.Position{Line: 2, Character: 4},
			wantEnd:     protocol.Position{Line: 2, Character: 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			state := analysis.NewState()
			diags := state.DidOpen(1, "file:///tmp/test.monkey", tt.content)

			var found *protocol.Diagnostic
			for _, d := range toProtocolDiagnostics(diags) {
				if d.Message == tt.wantMessage {
					found = &d
					break
				}
			}
			if found == nil {
				t.Fatalf("diagnostic %q not found in %v", tt.wantMessage, diags)
			}

			if found.Range.Start != tt.wantStart {
				t.Errorf("start = %+v, want %+v", found.Range.Start, tt.wantStart)
			}
			if found.Range.End != tt.wantEnd {
				t.Errorf("end = %+v, want %+v", found.Range.End, tt.wantEnd)
			}
		})
	}
}
