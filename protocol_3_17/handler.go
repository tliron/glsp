package protocol

import (
	"encoding/json"
	"errors"
	"sync"

	"github.com/tliron/glsp"
	protocol316 "github.com/tliron/glsp/protocol_3_16"
)

type Handler struct {
	protocol316.Handler

	Initialize             InitializeFunc
	TextDocumentDiagnostic TextDocumentDiagnosticFunc

	initialized bool
	lock        sync.Mutex
}

func (self *Handler) Handle(context *glsp.Context) (r any, validMethod bool, validParams bool, err error) {
	if !self.IsInitialized() && (context.Method != protocol316.MethodInitialize) {
		return nil, true, true, errors.New("server not initialized")
	}

	switch context.Method {
	case protocol316.MethodCancelRequest:
		if self.CancelRequest != nil {
			validMethod = true
			var params protocol316.CancelParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.CancelRequest(context, &params)
			}
		}

	case protocol316.MethodProgress:
		if self.Progress != nil {
			validMethod = true
			var params protocol316.ProgressParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.Progress(context, &params)
			}
		}

	// General Messages

	case MethodInitialize:
		if self.Initialize != nil {
			validMethod = true
			var params InitializeParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				if r, err = self.Initialize(context, &params); err == nil {
					self.SetInitialized(true)
				}
			}
		}

	case protocol316.MethodInitialized:
		if self.Initialized != nil {
			validMethod = true
			var params protocol316.InitializedParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.Initialized(context, &params)
			}
		}

	case protocol316.MethodShutdown:
		self.SetInitialized(false)
		if self.Shutdown != nil {
			validMethod = true
			validParams = true
			err = self.Shutdown(context)
		}

	case protocol316.MethodExit:
		// Note that the server will close the connection after we handle it here
		if self.Exit != nil {
			validMethod = true
			validParams = true
			err = self.Exit(context)
		}

	case protocol316.MethodLogTrace:
		if self.LogTrace != nil {
			validMethod = true
			var params protocol316.LogTraceParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.LogTrace(context, &params)
			}
		}

	case protocol316.MethodSetTrace:
		if self.SetTrace != nil {
			validMethod = true
			var params protocol316.SetTraceParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.SetTrace(context, &params)
			}
		}

	// Window

	case protocol316.MethodWindowWorkDoneProgressCancel:
		if self.WindowWorkDoneProgressCancel != nil {
			validMethod = true
			var params protocol316.WorkDoneProgressCancelParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.WindowWorkDoneProgressCancel(context, &params)
			}
		}

	// Workspace

	case protocol316.MethodWorkspaceDidChangeWorkspaceFolders:
		if self.WorkspaceDidChangeWorkspaceFolders != nil {
			validMethod = true
			var params protocol316.DidChangeWorkspaceFoldersParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.WorkspaceDidChangeWorkspaceFolders(context, &params)
			}
		}

	case protocol316.MethodWorkspaceDidChangeConfiguration:
		if self.WorkspaceDidChangeConfiguration != nil {
			validMethod = true
			var params protocol316.DidChangeConfigurationParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.WorkspaceDidChangeConfiguration(context, &params)
			}
		}

	case protocol316.MethodWorkspaceDidChangeWatchedFiles:
		if self.WorkspaceDidChangeWatchedFiles != nil {
			validMethod = true
			var params protocol316.DidChangeWatchedFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.WorkspaceDidChangeWatchedFiles(context, &params)
			}
		}

	case protocol316.MethodWorkspaceSymbol:
		if self.WorkspaceSymbol != nil {
			validMethod = true
			var params protocol316.WorkspaceSymbolParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.WorkspaceSymbol(context, &params)
			}
		}

	case protocol316.MethodWorkspaceExecuteCommand:
		if self.WorkspaceExecuteCommand != nil {
			validMethod = true
			var params protocol316.ExecuteCommandParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.WorkspaceExecuteCommand(context, &params)
			}
		}

	case protocol316.MethodWorkspaceWillCreateFiles:
		if self.WorkspaceWillCreateFiles != nil {
			validMethod = true
			var params protocol316.CreateFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.WorkspaceWillCreateFiles(context, &params)
			}
		}

	case protocol316.MethodWorkspaceDidCreateFiles:
		if self.WorkspaceDidCreateFiles != nil {
			validMethod = true
			var params protocol316.CreateFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.WorkspaceDidCreateFiles(context, &params)
			}
		}

	case protocol316.MethodWorkspaceWillRenameFiles:
		if self.WorkspaceWillRenameFiles != nil {
			validMethod = true
			var params protocol316.RenameFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.WorkspaceWillRenameFiles(context, &params)
			}
		}

	case protocol316.MethodWorkspaceDidRenameFiles:
		if self.WorkspaceDidRenameFiles != nil {
			validMethod = true
			var params protocol316.RenameFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.WorkspaceDidRenameFiles(context, &params)
			}
		}

	case protocol316.MethodWorkspaceWillDeleteFiles:
		if self.WorkspaceWillDeleteFiles != nil {
			validMethod = true
			var params protocol316.DeleteFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.WorkspaceWillDeleteFiles(context, &params)
			}
		}

	case protocol316.MethodWorkspaceDidDeleteFiles:
		if self.WorkspaceDidDeleteFiles != nil {
			validMethod = true
			var params protocol316.DeleteFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.WorkspaceDidDeleteFiles(context, &params)
			}
		}

	// Text Document Synchronization

	case protocol316.MethodTextDocumentDidOpen:
		if self.TextDocumentDidOpen != nil {
			validMethod = true
			var params protocol316.DidOpenTextDocumentParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.TextDocumentDidOpen(context, &params)
			}
		}

	case protocol316.MethodTextDocumentDidChange:
		if self.TextDocumentDidChange != nil {
			validMethod = true
			var params protocol316.DidChangeTextDocumentParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.TextDocumentDidChange(context, &params)
			}
		}

	case protocol316.MethodTextDocumentWillSave:
		if self.TextDocumentWillSave != nil {
			validMethod = true
			var params protocol316.WillSaveTextDocumentParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.TextDocumentWillSave(context, &params)
			}
		}

	case protocol316.MethodTextDocumentWillSaveWaitUntil:
		if self.TextDocumentWillSaveWaitUntil != nil {
			validMethod = true
			var params protocol316.WillSaveTextDocumentParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentWillSaveWaitUntil(context, &params)
			}
		}

	case protocol316.MethodTextDocumentDidSave:
		if self.TextDocumentDidSave != nil {
			validMethod = true
			var params protocol316.DidSaveTextDocumentParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.TextDocumentDidSave(context, &params)
			}
		}

	case protocol316.MethodTextDocumentDidClose:
		if self.TextDocumentDidClose != nil {
			validMethod = true
			var params protocol316.DidCloseTextDocumentParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.TextDocumentDidClose(context, &params)
			}
		}

	// Language Features

	case protocol316.MethodTextDocumentCompletion:
		if self.TextDocumentCompletion != nil {
			validMethod = true
			var params protocol316.CompletionParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentCompletion(context, &params)
			}
		}

	case protocol316.MethodCompletionItemResolve:
		if self.CompletionItemResolve != nil {
			validMethod = true
			var params protocol316.CompletionItem
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.CompletionItemResolve(context, &params)
			}
		}

	case protocol316.MethodTextDocumentHover:
		if self.TextDocumentHover != nil {
			validMethod = true
			var params protocol316.HoverParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentHover(context, &params)
			}
		}

	case protocol316.MethodTextDocumentSignatureHelp:
		if self.TextDocumentSignatureHelp != nil {
			validMethod = true
			var params protocol316.SignatureHelpParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentSignatureHelp(context, &params)
			}
		}

	case protocol316.MethodTextDocumentDeclaration:
		if self.TextDocumentDeclaration != nil {
			validMethod = true
			var params protocol316.DeclarationParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentDeclaration(context, &params)
			}
		}

	case protocol316.MethodTextDocumentDefinition:
		if self.TextDocumentDefinition != nil {
			validMethod = true
			var params protocol316.DefinitionParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentDefinition(context, &params)
			}
		}

	case protocol316.MethodTextDocumentTypeDefinition:
		if self.TextDocumentTypeDefinition != nil {
			validMethod = true
			var params protocol316.TypeDefinitionParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentTypeDefinition(context, &params)
			}
		}

	case protocol316.MethodTextDocumentImplementation:
		if self.TextDocumentImplementation != nil {
			validMethod = true
			var params protocol316.ImplementationParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentImplementation(context, &params)
			}
		}

	case protocol316.MethodTextDocumentReferences:
		if self.TextDocumentReferences != nil {
			validMethod = true
			var params protocol316.ReferenceParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentReferences(context, &params)
			}
		}

	case protocol316.MethodTextDocumentDocumentHighlight:
		if self.TextDocumentDocumentHighlight != nil {
			validMethod = true
			var params protocol316.DocumentHighlightParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentDocumentHighlight(context, &params)
			}
		}

	case protocol316.MethodTextDocumentDocumentSymbol:
		if self.TextDocumentDocumentSymbol != nil {
			validMethod = true
			var params protocol316.DocumentSymbolParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentDocumentSymbol(context, &params)
			}
		}

	case protocol316.MethodTextDocumentCodeAction:
		if self.TextDocumentCodeAction != nil {
			validMethod = true
			var params protocol316.CodeActionParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentCodeAction(context, &params)
			}
		}

	case protocol316.MethodCodeActionResolve:
		if self.CodeActionResolve != nil {
			validMethod = true
			var params protocol316.CodeAction
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.CodeActionResolve(context, &params)
			}
		}

	case protocol316.MethodTextDocumentCodeLens:
		if self.TextDocumentCodeLens != nil {
			validMethod = true
			var params protocol316.CodeLensParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentCodeLens(context, &params)
			}
		}

	case protocol316.MethodCodeLensResolve:
		if self.TextDocumentDidClose != nil {
			validMethod = true
			var params protocol316.CodeLens
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.CodeLensResolve(context, &params)
			}
		}

	case protocol316.MethodTextDocumentDocumentLink:
		if self.TextDocumentDocumentLink != nil {
			validMethod = true
			var params protocol316.DocumentLinkParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentDocumentLink(context, &params)
			}
		}

	case protocol316.MethodDocumentLinkResolve:
		if self.DocumentLinkResolve != nil {
			validMethod = true
			var params protocol316.DocumentLink
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.DocumentLinkResolve(context, &params)
			}
		}

	case protocol316.MethodTextDocumentColor:
		if self.TextDocumentColor != nil {
			validMethod = true
			var params protocol316.DocumentColorParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentColor(context, &params)
			}
		}

	case protocol316.MethodTextDocumentColorPresentation:
		if self.TextDocumentColorPresentation != nil {
			validMethod = true
			var params protocol316.ColorPresentationParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentColorPresentation(context, &params)
			}
		}

	case protocol316.MethodTextDocumentFormatting:
		if self.TextDocumentFormatting != nil {
			validMethod = true
			var params protocol316.DocumentFormattingParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentFormatting(context, &params)
			}
		}

	case protocol316.MethodTextDocumentRangeFormatting:
		if self.TextDocumentRangeFormatting != nil {
			validMethod = true
			var params protocol316.DocumentRangeFormattingParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentRangeFormatting(context, &params)
			}
		}

	case protocol316.MethodTextDocumentOnTypeFormatting:
		if self.TextDocumentOnTypeFormatting != nil {
			validMethod = true
			var params protocol316.DocumentOnTypeFormattingParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentOnTypeFormatting(context, &params)
			}
		}

	case protocol316.MethodTextDocumentRename:
		if self.TextDocumentRename != nil {
			validMethod = true
			var params protocol316.RenameParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentRename(context, &params)
			}
		}

	case protocol316.MethodTextDocumentPrepareRename:
		if self.TextDocumentPrepareRename != nil {
			validMethod = true
			var params protocol316.PrepareRenameParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentPrepareRename(context, &params)
			}
		}

	case protocol316.MethodTextDocumentFoldingRange:
		if self.TextDocumentFoldingRange != nil {
			validMethod = true
			var params protocol316.FoldingRangeParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentFoldingRange(context, &params)
			}
		}

	case protocol316.MethodTextDocumentSelectionRange:
		if self.TextDocumentSelectionRange != nil {
			validMethod = true
			var params protocol316.SelectionRangeParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentSelectionRange(context, &params)
			}
		}

	case protocol316.MethodTextDocumentPrepareCallHierarchy:
		if self.TextDocumentPrepareCallHierarchy != nil {
			validMethod = true
			var params protocol316.CallHierarchyPrepareParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentPrepareCallHierarchy(context, &params)
			}
		}

	case protocol316.MethodCallHierarchyIncomingCalls:
		if self.CallHierarchyIncomingCalls != nil {
			validMethod = true
			var params protocol316.CallHierarchyIncomingCallsParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.CallHierarchyIncomingCalls(context, &params)
			}
		}

	case protocol316.MethodCallHierarchyOutgoingCalls:
		if self.CallHierarchyOutgoingCalls != nil {
			validMethod = true
			var params protocol316.CallHierarchyOutgoingCallsParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.CallHierarchyOutgoingCalls(context, &params)
			}
		}

	case protocol316.MethodTextDocumentSemanticTokensFull:
		if self.TextDocumentSemanticTokensFull != nil {
			validMethod = true
			var params protocol316.SemanticTokensParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentSemanticTokensFull(context, &params)
			}
		}

	case protocol316.MethodTextDocumentSemanticTokensFullDelta:
		if self.TextDocumentSemanticTokensFullDelta != nil {
			validMethod = true
			var params protocol316.SemanticTokensDeltaParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentSemanticTokensFullDelta(context, &params)
			}
		}

	case protocol316.MethodTextDocumentSemanticTokensRange:
		if self.TextDocumentSemanticTokensRange != nil {
			validMethod = true
			var params protocol316.SemanticTokensRangeParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentSemanticTokensRange(context, &params)
			}
		}

	case protocol316.MethodWorkspaceSemanticTokensRefresh:
		if self.WorkspaceSemanticTokensRefresh != nil {
			validMethod = true
			validParams = true
			err = self.WorkspaceSemanticTokensRefresh(context)
		}

	case protocol316.MethodTextDocumentLinkedEditingRange:
		if self.TextDocumentLinkedEditingRange != nil {
			validMethod = true
			var params protocol316.LinkedEditingRangeParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentLinkedEditingRange(context, &params)
			}
		}

	case protocol316.MethodTextDocumentMoniker:
		if self.TextDocumentMoniker != nil {
			validMethod = true
			var params protocol316.MonikerParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentMoniker(context, &params)
			}
		}
	case MethodTextDocumentDiagnostic:
		if self.TextDocumentDiagnostic != nil {
			validMethod = true
			var params DocumentDiagnosticParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentDiagnostic(context, &params)
			}
		}
	}

	return

}

func (self *Handler) IsInitialized() bool {
	self.lock.Lock()
	defer self.lock.Unlock()
	return self.initialized
}

func (self *Handler) SetInitialized(initialized bool) {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.initialized = initialized
}

func (self *Handler) CreateServerCapabilities() ServerCapabilities {
	var capabilities ServerCapabilities

	if (self.TextDocumentDidOpen != nil) || (self.TextDocumentDidClose != nil) {
		if _, ok := capabilities.TextDocumentSync.(*protocol316.TextDocumentSyncOptions); !ok {
			capabilities.TextDocumentSync = &protocol316.TextDocumentSyncOptions{}
		}
		capabilities.TextDocumentSync.(*protocol316.TextDocumentSyncOptions).OpenClose = &protocol316.True
	}

	if self.TextDocumentDidChange != nil {
		if _, ok := capabilities.TextDocumentSync.(*protocol316.TextDocumentSyncOptions); !ok {
			capabilities.TextDocumentSync = &protocol316.TextDocumentSyncOptions{}
		}
		// This can be overriden to TextDocumentSyncKindFull
		value := protocol316.TextDocumentSyncKindIncremental
		capabilities.TextDocumentSync.(*protocol316.TextDocumentSyncOptions).Change = &value
	}

	if self.TextDocumentWillSave != nil {
		if _, ok := capabilities.TextDocumentSync.(*protocol316.TextDocumentSyncOptions); !ok {
			capabilities.TextDocumentSync = &protocol316.TextDocumentSyncOptions{}
		}
		capabilities.TextDocumentSync.(*protocol316.TextDocumentSyncOptions).WillSave = &protocol316.True
	}

	if self.TextDocumentWillSaveWaitUntil != nil {
		if _, ok := capabilities.TextDocumentSync.(*protocol316.TextDocumentSyncOptions); !ok {
			capabilities.TextDocumentSync = &protocol316.TextDocumentSyncOptions{}
		}
		capabilities.TextDocumentSync.(*protocol316.TextDocumentSyncOptions).WillSaveWaitUntil = &protocol316.True
	}

	if self.TextDocumentDidSave != nil {
		if _, ok := capabilities.TextDocumentSync.(*protocol316.TextDocumentSyncOptions); !ok {
			capabilities.TextDocumentSync = &protocol316.TextDocumentSyncOptions{}
		}
		capabilities.TextDocumentSync.(*protocol316.TextDocumentSyncOptions).Save = &protocol316.True
	}

	if self.TextDocumentCompletion != nil {
		capabilities.CompletionProvider = &protocol316.CompletionOptions{}
	}

	if self.TextDocumentHover != nil {
		capabilities.HoverProvider = true
	}

	if self.TextDocumentSignatureHelp != nil {
		capabilities.SignatureHelpProvider = &protocol316.SignatureHelpOptions{}
	}

	if self.TextDocumentDeclaration != nil {
		capabilities.DeclarationProvider = true
	}

	if self.TextDocumentDefinition != nil {
		capabilities.DefinitionProvider = true
	}

	if self.TextDocumentTypeDefinition != nil {
		capabilities.TypeDefinitionProvider = true
	}

	if self.TextDocumentImplementation != nil {
		capabilities.ImplementationProvider = true
	}

	if self.TextDocumentReferences != nil {
		capabilities.ReferencesProvider = true
	}

	if self.TextDocumentDocumentHighlight != nil {
		capabilities.DocumentHighlightProvider = true
	}

	if self.TextDocumentDocumentSymbol != nil {
		capabilities.DocumentSymbolProvider = true
	}

	if self.TextDocumentCodeAction != nil {
		capabilities.CodeActionProvider = true
	}

	if self.TextDocumentCodeLens != nil {
		capabilities.CodeLensProvider = &protocol316.CodeLensOptions{}
	}

	if self.TextDocumentDocumentLink != nil {
		capabilities.DocumentLinkProvider = &protocol316.DocumentLinkOptions{}
	}

	if self.TextDocumentColor != nil {
		capabilities.ColorProvider = true
	}

	if self.TextDocumentFormatting != nil {
		capabilities.DocumentFormattingProvider = true
	}

	if self.TextDocumentRangeFormatting != nil {
		capabilities.DocumentRangeFormattingProvider = true
	}

	if self.TextDocumentOnTypeFormatting != nil {
		capabilities.DocumentOnTypeFormattingProvider = &protocol316.DocumentOnTypeFormattingOptions{}
	}

	if self.TextDocumentRename != nil {
		capabilities.RenameProvider = true
	}

	if self.TextDocumentFoldingRange != nil {
		capabilities.FoldingRangeProvider = true
	}

	if self.WorkspaceExecuteCommand != nil {
		capabilities.ExecuteCommandProvider = &protocol316.ExecuteCommandOptions{}
	}

	if self.TextDocumentSelectionRange != nil {
		capabilities.SelectionRangeProvider = true
	}

	if self.TextDocumentLinkedEditingRange != nil {
		capabilities.LinkedEditingRangeProvider = true
	}

	if self.TextDocumentPrepareCallHierarchy != nil {
		capabilities.CallHierarchyProvider = true
	}

	if self.TextDocumentSemanticTokensFull != nil {
		if _, ok := capabilities.SemanticTokensProvider.(*protocol316.SemanticTokensOptions); !ok {
			capabilities.SemanticTokensProvider = &protocol316.SemanticTokensOptions{}
		}
		if self.TextDocumentSemanticTokensFullDelta != nil {
			capabilities.SemanticTokensProvider.(*protocol316.SemanticTokensOptions).Full = &protocol316.SemanticDelta{}
			capabilities.SemanticTokensProvider.(*protocol316.SemanticTokensOptions).Full.(*protocol316.SemanticDelta).Delta = &protocol316.True
		} else {
			capabilities.SemanticTokensProvider.(*protocol316.SemanticTokensOptions).Full = true
		}
	}

	if self.TextDocumentSemanticTokensRange != nil {
		if _, ok := capabilities.SemanticTokensProvider.(*protocol316.SemanticTokensOptions); !ok {
			capabilities.SemanticTokensProvider = &protocol316.SemanticTokensOptions{}
		}
		capabilities.SemanticTokensProvider.(*protocol316.SemanticTokensOptions).Range = true
	}

	// TODO: self.TextDocumentSemanticTokensRefresh?

	if self.TextDocumentMoniker != nil {
		capabilities.MonikerProvider = true
	}

	if self.WorkspaceSymbol != nil {
		capabilities.WorkspaceSymbolProvider = true
	}

	if self.WorkspaceDidCreateFiles != nil {
		if capabilities.Workspace == nil {
			capabilities.Workspace = &protocol316.ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &protocol316.ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.DidCreate = &protocol316.FileOperationRegistrationOptions{
			Filters: []protocol316.FileOperationFilter{},
		}
	}

	if self.WorkspaceWillCreateFiles != nil {
		if capabilities.Workspace == nil {
			capabilities.Workspace = &protocol316.ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &protocol316.ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.WillCreate = &protocol316.FileOperationRegistrationOptions{
			Filters: []protocol316.FileOperationFilter{},
		}
	}

	if self.WorkspaceDidRenameFiles != nil {
		capabilities.RenameProvider = true
		if capabilities.Workspace == nil {
			capabilities.Workspace = &protocol316.ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &protocol316.ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.DidRename = &protocol316.FileOperationRegistrationOptions{
			Filters: []protocol316.FileOperationFilter{},
		}
	}

	if self.WorkspaceWillRenameFiles != nil {
		capabilities.RenameProvider = true
		if capabilities.Workspace == nil {
			capabilities.Workspace = &protocol316.ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &protocol316.ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.WillRename = &protocol316.FileOperationRegistrationOptions{
			Filters: []protocol316.FileOperationFilter{},
		}
	}

	if self.WorkspaceDidDeleteFiles != nil {
		if capabilities.Workspace == nil {
			capabilities.Workspace = &protocol316.ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &protocol316.ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.DidDelete = &protocol316.FileOperationRegistrationOptions{
			Filters: []protocol316.FileOperationFilter{},
		}
	}

	if self.WorkspaceWillDeleteFiles != nil {
		if capabilities.Workspace == nil {
			capabilities.Workspace = &protocol316.ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &protocol316.ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.WillDelete = &protocol316.FileOperationRegistrationOptions{
			Filters: []protocol316.FileOperationFilter{},
		}
	}

	if self.TextDocumentDiagnostic != nil {
		capabilities.DiagnosticProvider = DiagnosticOptions{
			InterFileDependencies: true,
			WorkspaceDiagnostics:  false,
		}
	}

	return capabilities
}
