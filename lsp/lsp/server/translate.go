package server

import (
	"github.com/SegniAT/monkey-language-interpreter/lsp/analysis"
	"github.com/SegniAT/monkey-language-interpreter/lsp/lsp/protocol"
)

func toProtocolHover(h *analysis.Hover) *protocol.Hover {
	if h == nil {
		return nil
	}

	var r *protocol.Range
	if h.Range != nil {
		converted := toProtocolRange(*h.Range)
		r = &converted
	}

	return &protocol.Hover{
		Contents: toProtocolMarkupContent(h.Contents),
		Range:    r,
	}
}

func toProtocolLocation(l *analysis.Location) *protocol.Location {
	if l == nil {
		return nil
	}

	return &protocol.Location{
		URI:   l.URI,
		Range: toProtocolRange(l.Range),
	}
}

func toProtocolCompletionItems(items []analysis.CompletionItem) []protocol.CompletionItem {
	result := make([]protocol.CompletionItem, 0, len(items))
	for _, item := range items {
		result = append(result, protocol.CompletionItem{
			Label:         item.Label,
			Detail:        item.Detail,
			Kind:          (*protocol.CompletionItemKind)(&item.Kind),
			Documentation: toProtocolMarkupContent(item.Documentation),
		})
	}
	return result
}

// toProtocolRange converts an internal inclusive (1-indexed) analysis range into LSP's exclusive (0-indexed) range
func toProtocolRange(r analysis.Range) protocol.Range {
	return protocol.Range{
		Start: protocol.Position{
			Line:      r.Start.Line - 1,
			Character: r.Start.Character,
		},
		End: protocol.Position{
			Line:      r.End.Line - 1,
			Character: r.End.Character, // LSP end character is 0-indexed exclusive
		},
	}
}

func toProtocolMarkupContent(c analysis.MarkupContent) protocol.MarkupContent {
	return protocol.MarkupContent{
		Kind:  protocol.MarkupKind(c.Kind),
		Value: c.Value,
	}
}
