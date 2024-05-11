package protocol

import (
	"encoding/json"

	"github.com/tliron/glsp"
	protocol316 "github.com/tliron/glsp/protocol_3_16"
)

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#initialize

const MethodInitialize = protocol316.Method("initialize")

// Returns: InitializeResult | InitializeError
type InitializeFunc func(context *glsp.Context, params *InitializeParams) (any, error)

type InitializeParams struct {
	protocol316.InitializeParams

	/**
	 * The capabilities provided by the client (editor or tool)
	 */
	Capabilities ClientCapabilities `json:"capabilities"`
}

type ClientCapabilities struct {
	protocol316.ClientCapabilities

	TextDocument *TextDocumentClientCapabilities `json:"textDocument,omitempty"`
}

/**
 * Text document specific client capabilities.
 */
type TextDocumentClientCapabilities struct {
	protocol316.TextDocumentClientCapabilities

	/**
	 * Capabilities specific to the diagnostic pull model.
	 *
	 * @since 3.17.0
	 */
	Diagnostic *DiagnosticClientCapabilities `json:"diagnostic,omitempty"`
}

type ServerCapabilities struct {
	protocol316.ServerCapabilities

	/**
	 * The server has support for pull model diagnostics.
	 *
	 * @since 3.17.0
	 */
	DiagnosticProvider any `json:"diagnosticProvider,omitempty"` // nil | DiagnosticOptions | DiagnosticRegistrationOptions
}

func (self *ServerCapabilities) UnmarshalJSON(data []byte) error {
	var value struct {
		TextDocumentSync                 json.RawMessage                              `json:"textDocumentSync,omitempty"` // nil | TextDocumentSyncOptions | TextDocumentSyncKind
		CompletionProvider               *protocol316.CompletionOptions               `json:"completionProvider,omitempty"`
		HoverProvider                    json.RawMessage                              `json:"hoverProvider,omitempty"` // nil | bool | HoverOptions
		SignatureHelpProvider            *protocol316.SignatureHelpOptions            `json:"signatureHelpProvider,omitempty"`
		DeclarationProvider              json.RawMessage                              `json:"declarationProvider,omitempty"`       // nil | bool | DeclarationOptions | DeclarationRegistrationOptions
		DefinitionProvider               json.RawMessage                              `json:"definitionProvider,omitempty"`        // nil | bool | DefinitionOptions
		TypeDefinitionProvider           json.RawMessage                              `json:"typeDefinitionProvider,omitempty"`    // nil | bool | TypeDefinitionOption | TypeDefinitionRegistrationOptions
		ImplementationProvider           json.RawMessage                              `json:"implementationProvider,omitempty"`    // nil | bool | ImplementationOptions | ImplementationRegistrationOptions
		ReferencesProvider               json.RawMessage                              `json:"referencesProvider,omitempty"`        // nil | bool | ReferenceOptions
		DocumentHighlightProvider        json.RawMessage                              `json:"documentHighlightProvider,omitempty"` // nil | bool | DocumentHighlightOptions
		DocumentSymbolProvider           json.RawMessage                              `json:"documentSymbolProvider,omitempty"`    // nil | bool | DocumentSymbolOptions
		CodeActionProvider               json.RawMessage                              `json:"codeActionProvider,omitempty"`        // nil | bool | CodeActionOptions
		CodeLensProvider                 *protocol316.CodeLensOptions                 `json:"codeLensProvider,omitempty"`
		DocumentLinkProvider             *protocol316.DocumentLinkOptions             `json:"documentLinkProvider,omitempty"`
		ColorProvider                    json.RawMessage                              `json:"colorProvider,omitempty"`                   // nil | bool | DocumentColorOptions | DocumentColorRegistrationOptions
		DocumentFormattingProvider       json.RawMessage                              `json:"documentFormattingProvider,omitempty"`      // nil | bool | DocumentFormattingOptions
		DocumentRangeFormattingProvider  json.RawMessage                              `json:"documentRangeFormattingProvider,omitempty"` // nil | bool | DocumentRangeFormattingOptions
		DocumentOnTypeFormattingProvider *protocol316.DocumentOnTypeFormattingOptions `json:"documentOnTypeFormattingProvider,omitempty"`
		RenameProvider                   json.RawMessage                              `json:"renameProvider,omitempty"`       // nil | bool | RenameOptions
		FoldingRangeProvider             json.RawMessage                              `json:"foldingRangeProvider,omitempty"` // nil | bool | FoldingRangeOptions | FoldingRangeRegistrationOptions
		ExecuteCommandProvider           *protocol316.ExecuteCommandOptions           `json:"executeCommandProvider,omitempty"`
		SelectionRangeProvider           json.RawMessage                              `json:"selectionRangeProvider,omitempty"`     // nil | bool | SelectionRangeOptions | SelectionRangeRegistrationOptions
		LinkedEditingRangeProvider       json.RawMessage                              `json:"linkedEditingRangeProvider,omitempty"` // nil | bool | LinkedEditingRangeOptions | LinkedEditingRangeRegistrationOptions
		CallHierarchyProvider            json.RawMessage                              `json:"callHierarchyProvider,omitempty"`      // nil | bool | CallHierarchyOptions | CallHierarchyRegistrationOptions
		SemanticTokensProvider           json.RawMessage                              `json:"semanticTokensProvider,omitempty"`     // nil | SemanticTokensOptions | SemanticTokensRegistrationOptions
		MonikerProvider                  json.RawMessage                              `json:"monikerProvider,omitempty"`            // nil | bool | MonikerOptions | MonikerRegistrationOptions
		WorkspaceSymbolProvider          json.RawMessage                              `json:"workspaceSymbolProvider,omitempty"`    // nil | bool | WorkspaceSymbolOptions
		Workspace                        *protocol316.ServerCapabilitiesWorkspace     `json:"workspace,omitempty"`
		Experimental                     *any                                         `json:"experimental,omitempty"`
		DiagnosticProvider               json.RawMessage                              `json:"diagnosticProvider,omitempty"` // nil | DiagnosticOptions | DiagnosticRegistrationOptions
	}

	if err := json.Unmarshal(data, &value); err == nil {
		self.CompletionProvider = value.CompletionProvider
		self.SignatureHelpProvider = value.SignatureHelpProvider
		self.CodeLensProvider = value.CodeLensProvider
		self.DocumentLinkProvider = value.DocumentLinkProvider
		self.DocumentOnTypeFormattingProvider = value.DocumentOnTypeFormattingProvider
		self.ExecuteCommandProvider = value.ExecuteCommandProvider
		self.Workspace = value.Workspace

		if value.TextDocumentSync != nil {
			var value_ protocol316.TextDocumentSyncOptions
			if err = json.Unmarshal(value.TextDocumentSync, &value_); err == nil {
				self.TextDocumentSync = value_
			} else {
				var value_ protocol316.TextDocumentSyncKind
				if err = json.Unmarshal(value.TextDocumentSync, &value_); err == nil {
					self.TextDocumentSync = value_
				} else {
					return err
				}
			}
		}

		if value.HoverProvider != nil {
			var value_ bool
			if err = json.Unmarshal(value.HoverProvider, &value_); err == nil {
				self.HoverProvider = value_
			} else {
				var value_ protocol316.HoverOptions
				if err = json.Unmarshal(value.HoverProvider, &value_); err == nil {
					self.HoverProvider = value_
				} else {
					return err
				}
			}
		}

		if value.DeclarationProvider != nil {
			var value_ bool
			if err = json.Unmarshal(value.DeclarationProvider, &value_); err == nil {
				self.DeclarationProvider = value_
			} else {
				var value_ protocol316.DeclarationOptions
				if err = json.Unmarshal(value.DeclarationProvider, &value_); err == nil {
					self.DeclarationProvider = value_
				} else {
					var value_ protocol316.DeclarationRegistrationOptions
					if err = json.Unmarshal(value.DeclarationProvider, &value_); err == nil {
						self.DeclarationProvider = value_
					} else {
						return err
					}
				}
			}
		}

		if value.DefinitionProvider != nil {
			var value_ bool
			if err = json.Unmarshal(value.DefinitionProvider, &value_); err == nil {
				self.DefinitionProvider = value_
			} else {
				var value_ protocol316.DefinitionOptions
				if err = json.Unmarshal(value.DefinitionProvider, &value_); err == nil {
					self.DefinitionProvider = value_
				} else {
					return err
				}
			}
		}

		if value.TypeDefinitionProvider != nil {
			var value_ bool
			if err = json.Unmarshal(value.TypeDefinitionProvider, &value_); err == nil {
				self.TypeDefinitionProvider = value_
			} else {
				var value_ protocol316.TypeDefinitionOptions
				if err = json.Unmarshal(value.TypeDefinitionProvider, &value_); err == nil {
					self.TypeDefinitionProvider = value_
				} else {
					var value_ protocol316.TypeDefinitionRegistrationOptions
					if err = json.Unmarshal(value.TypeDefinitionProvider, &value_); err == nil {
						self.TypeDefinitionProvider = value_
					} else {
						return err
					}
				}
			}
		}

		if value.ImplementationProvider != nil {
			var value_ bool
			if err = json.Unmarshal(value.ImplementationProvider, &value_); err == nil {
				self.ImplementationProvider = value_
			} else {
				var value_ protocol316.ImplementationOptions
				if err = json.Unmarshal(value.ImplementationProvider, &value_); err == nil {
					self.ImplementationProvider = value_
				} else {
					var value_ protocol316.ImplementationRegistrationOptions
					if err = json.Unmarshal(value.ImplementationProvider, &value_); err == nil {
						self.ImplementationProvider = value_
					} else {
						return err
					}
				}
			}
		}

		if value.ReferencesProvider != nil {
			var value_ bool
			if err = json.Unmarshal(value.ReferencesProvider, &value_); err == nil {
				self.ReferencesProvider = value_
			} else {
				var value_ protocol316.ReferenceOptions
				if err = json.Unmarshal(value.ReferencesProvider, &value_); err == nil {
					self.ReferencesProvider = value_
				} else {
					return err
				}
			}
		}

		if value.DocumentHighlightProvider != nil {
			var value_ bool
			if err = json.Unmarshal(value.DocumentHighlightProvider, &value_); err == nil {
				self.DocumentHighlightProvider = value_
			} else {
				var value_ protocol316.DocumentHighlightOptions
				if err = json.Unmarshal(value.DocumentHighlightProvider, &value_); err == nil {
					self.DocumentHighlightProvider = value_
				} else {
					return err
				}
			}
		}

		if value.DocumentSymbolProvider != nil {
			var value_ bool
			if err = json.Unmarshal(value.DocumentSymbolProvider, &value_); err == nil {
				self.DocumentSymbolProvider = value_
			} else {
				var value_ protocol316.DocumentSymbolOptions
				if err = json.Unmarshal(value.DocumentSymbolProvider, &value_); err == nil {
					self.DocumentSymbolProvider = value_
				} else {
					return err
				}
			}
		}

		if value.CodeActionProvider != nil {
			var value_ bool
			if err = json.Unmarshal(value.CodeActionProvider, &value_); err == nil {
				self.CodeActionProvider = value_
			} else {
				var value_ protocol316.CodeActionOptions
				if err = json.Unmarshal(value.CodeActionProvider, &value_); err == nil {
					self.CodeActionProvider = value_
				} else {
					return err
				}
			}
		}

		if value.ColorProvider != nil {
			var value_ bool
			if err = json.Unmarshal(value.ColorProvider, &value_); err == nil {
				self.ColorProvider = value_
			} else {
				var value_ protocol316.DocumentColorOptions
				if err = json.Unmarshal(value.ColorProvider, &value_); err == nil {
					self.ColorProvider = value_
				} else {
					var value_ protocol316.DocumentColorRegistrationOptions
					if err = json.Unmarshal(value.ColorProvider, &value_); err == nil {
						self.ColorProvider = value_
					} else {
						return err
					}
				}
			}
		}

		if value.DocumentFormattingProvider != nil {
			var value_ bool
			if err = json.Unmarshal(value.DocumentFormattingProvider, &value_); err == nil {
				self.DocumentFormattingProvider = value_
			} else {
				var value_ protocol316.DocumentFormattingOptions
				if err = json.Unmarshal(value.DocumentFormattingProvider, &value_); err == nil {
					self.DocumentFormattingProvider = value_
				} else {
					return err
				}
			}
		}

		if value.DocumentRangeFormattingProvider != nil {
			var value_ bool
			if err = json.Unmarshal(value.DocumentRangeFormattingProvider, &value_); err == nil {
				self.DocumentRangeFormattingProvider = value_
			} else {
				var value_ protocol316.DocumentRangeFormattingOptions
				if err = json.Unmarshal(value.DocumentRangeFormattingProvider, &value_); err == nil {
					self.DocumentRangeFormattingProvider = value_
				} else {
					return err
				}
			}
		}

		if value.RenameProvider != nil {
			var value_ bool
			if err = json.Unmarshal(value.RenameProvider, &value_); err == nil {
				self.RenameProvider = value_
			} else {
				var value_ protocol316.RenameOptions
				if err = json.Unmarshal(value.RenameProvider, &value_); err == nil {
					self.RenameProvider = value_
				} else {
					return err
				}
			}
		}

		if value.FoldingRangeProvider != nil {
			var value_ bool
			if err = json.Unmarshal(value.FoldingRangeProvider, &value_); err == nil {
				self.FoldingRangeProvider = value_
			} else {
				var value_ protocol316.FoldingRangeOptions
				if err = json.Unmarshal(value.FoldingRangeProvider, &value_); err == nil {
					self.FoldingRangeProvider = value_
				} else {
					var value_ protocol316.FoldingRangeRegistrationOptions
					if err = json.Unmarshal(value.FoldingRangeProvider, &value_); err == nil {
						self.FoldingRangeProvider = value_
					} else {
						return err
					}
				}
			}
		}

		if value.SelectionRangeProvider != nil {
			var value_ bool
			if err = json.Unmarshal(value.SelectionRangeProvider, &value_); err == nil {
				self.SelectionRangeProvider = value_
			} else {
				var value_ protocol316.SelectionRangeOptions
				if err = json.Unmarshal(value.SelectionRangeProvider, &value_); err == nil {
					self.SelectionRangeProvider = value_
				} else {
					var value_ protocol316.SelectionRangeRegistrationOptions
					if err = json.Unmarshal(value.SelectionRangeProvider, &value_); err == nil {
						self.SelectionRangeProvider = value_
					} else {
						return err
					}
				}
			}
		}

		if value.LinkedEditingRangeProvider != nil {
			var value_ bool
			if err = json.Unmarshal(value.LinkedEditingRangeProvider, &value_); err == nil {
				self.LinkedEditingRangeProvider = value_
			} else {
				var value_ protocol316.LinkedEditingRangeOptions
				if err = json.Unmarshal(value.LinkedEditingRangeProvider, &value_); err == nil {
					self.LinkedEditingRangeProvider = value_
				} else {
					var value_ protocol316.LinkedEditingRangeRegistrationOptions
					if err = json.Unmarshal(value.LinkedEditingRangeProvider, &value_); err == nil {
						self.LinkedEditingRangeProvider = value_
					} else {
						return err
					}
				}
			}
		}

		if value.CallHierarchyProvider != nil {
			var value_ bool
			if err = json.Unmarshal(value.CallHierarchyProvider, &value_); err == nil {
				self.CallHierarchyProvider = value_
			} else {
				var value_ protocol316.CallHierarchyOptions
				if err = json.Unmarshal(value.CallHierarchyProvider, &value_); err == nil {
					self.CallHierarchyProvider = value_
				} else {
					var value_ protocol316.CallHierarchyRegistrationOptions
					if err = json.Unmarshal(value.CallHierarchyProvider, &value_); err == nil {
						self.CallHierarchyProvider = value_
					} else {
						return err
					}
				}
			}
		}

		if value.SemanticTokensProvider != nil {
			var value_ protocol316.SemanticTokensOptions
			if err = json.Unmarshal(value.SemanticTokensProvider, &value_); err == nil {
				self.SemanticTokensProvider = value_
			} else {
				var value_ protocol316.SemanticTokensRegistrationOptions
				if err = json.Unmarshal(value.SemanticTokensProvider, &value_); err == nil {
					self.SemanticTokensProvider = value_
				} else {
					return err
				}
			}
		}

		if value.MonikerProvider != nil {
			var value_ bool
			if err = json.Unmarshal(value.MonikerProvider, &value_); err == nil {
				self.MonikerProvider = value_
			} else {
				var value_ protocol316.MonikerOptions
				if err = json.Unmarshal(value.MonikerProvider, &value_); err == nil {
					self.MonikerProvider = value_
				} else {
					var value_ protocol316.MonikerRegistrationOptions
					if err = json.Unmarshal(value.MonikerProvider, &value_); err == nil {
						self.MonikerProvider = value_
					} else {
						return err
					}
				}
			}
		}

		if value.WorkspaceSymbolProvider != nil {
			var value_ bool
			if err = json.Unmarshal(value.WorkspaceSymbolProvider, &value_); err == nil {
				self.WorkspaceSymbolProvider = value_
			} else {
				var value_ protocol316.WorkspaceSymbolOptions
				if err = json.Unmarshal(value.WorkspaceSymbolProvider, &value_); err == nil {
					self.WorkspaceSymbolProvider = value_
				} else {
					return err
				}
			}
		}

		if value.DiagnosticProvider != nil {
			var value_ DiagnosticOptions
			if err = json.Unmarshal(value.DiagnosticProvider, &value_); err == nil {
				self.DiagnosticProvider = value_
			} else {
				var value_ DiagnosticRegistrationOptions
				if err = json.Unmarshal(value.DiagnosticProvider, &value_); err == nil {
					self.DiagnosticProvider = value_
				} else {
					return err
				}
			}
		}

		return nil
	} else {
		return err
	}
}

type InitializeResult struct {
	/**
	 * The capabilities the language server provides.
	 */
	Capabilities ServerCapabilities `json:"capabilities"`

	/**
	 * Information about the server.
	 *
	 * @since 3.15.0
	 */
	ServerInfo *protocol316.InitializeResultServerInfo `json:"serverInfo,omitempty"`
}
