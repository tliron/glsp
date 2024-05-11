package protocol317

import (
	"encoding/json"
	"errors"
	"sync"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type Handler struct {
	protocol.Handler

	Initialize             InitializeFunc
	TextDocumentDiagnostic TextDocumentDiagnosticFunc

	initialized bool
	lock        sync.Mutex
}

func (self *Handler) Handle(context *glsp.Context) (r any, validMethod bool, validParams bool, err error) {
	if !self.IsInitialized() && (context.Method != protocol.MethodInitialize) {
		return nil, true, true, errors.New("server not initialized")
	}

	switch context.Method {
	case protocol.MethodCancelRequest:
		if self.CancelRequest != nil {
			validMethod = true
			var params protocol.CancelParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.CancelRequest(context, &params)
			}
		}

	case protocol.MethodProgress:
		if self.Progress != nil {
			validMethod = true
			var params protocol.ProgressParams
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

	case protocol.MethodInitialized:
		if self.Initialized != nil {
			validMethod = true
			var params protocol.InitializedParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.Initialized(context, &params)
			}
		}

	case protocol.MethodShutdown:
		self.SetInitialized(false)
		if self.Shutdown != nil {
			validMethod = true
			validParams = true
			err = self.Shutdown(context)
		}

	case protocol.MethodExit:
		// Note that the server will close the connection after we handle it here
		if self.Exit != nil {
			validMethod = true
			validParams = true
			err = self.Exit(context)
		}

	case protocol.MethodLogTrace:
		if self.LogTrace != nil {
			validMethod = true
			var params protocol.LogTraceParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.LogTrace(context, &params)
			}
		}

	case protocol.MethodSetTrace:
		if self.SetTrace != nil {
			validMethod = true
			var params protocol.SetTraceParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.SetTrace(context, &params)
			}
		}

	// Window

	case protocol.MethodWindowWorkDoneProgressCancel:
		if self.WindowWorkDoneProgressCancel != nil {
			validMethod = true
			var params protocol.WorkDoneProgressCancelParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.WindowWorkDoneProgressCancel(context, &params)
			}
		}

	// Workspace

	case protocol.MethodWorkspaceDidChangeWorkspaceFolders:
		if self.WorkspaceDidChangeWorkspaceFolders != nil {
			validMethod = true
			var params protocol.DidChangeWorkspaceFoldersParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.WorkspaceDidChangeWorkspaceFolders(context, &params)
			}
		}

	case protocol.MethodWorkspaceDidChangeConfiguration:
		if self.WorkspaceDidChangeConfiguration != nil {
			validMethod = true
			var params protocol.DidChangeConfigurationParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.WorkspaceDidChangeConfiguration(context, &params)
			}
		}

	case protocol.MethodWorkspaceDidChangeWatchedFiles:
		if self.WorkspaceDidChangeWatchedFiles != nil {
			validMethod = true
			var params protocol.DidChangeWatchedFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.WorkspaceDidChangeWatchedFiles(context, &params)
			}
		}

	case protocol.MethodWorkspaceSymbol:
		if self.WorkspaceSymbol != nil {
			validMethod = true
			var params protocol.WorkspaceSymbolParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.WorkspaceSymbol(context, &params)
			}
		}

	case protocol.MethodWorkspaceExecuteCommand:
		if self.WorkspaceExecuteCommand != nil {
			validMethod = true
			var params protocol.ExecuteCommandParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.WorkspaceExecuteCommand(context, &params)
			}
		}

	case protocol.MethodWorkspaceWillCreateFiles:
		if self.WorkspaceWillCreateFiles != nil {
			validMethod = true
			var params protocol.CreateFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.WorkspaceWillCreateFiles(context, &params)
			}
		}

	case protocol.MethodWorkspaceDidCreateFiles:
		if self.WorkspaceDidCreateFiles != nil {
			validMethod = true
			var params protocol.CreateFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.WorkspaceDidCreateFiles(context, &params)
			}
		}

	case protocol.MethodWorkspaceWillRenameFiles:
		if self.WorkspaceWillRenameFiles != nil {
			validMethod = true
			var params protocol.RenameFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.WorkspaceWillRenameFiles(context, &params)
			}
		}

	case protocol.MethodWorkspaceDidRenameFiles:
		if self.WorkspaceDidRenameFiles != nil {
			validMethod = true
			var params protocol.RenameFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.WorkspaceDidRenameFiles(context, &params)
			}
		}

	case protocol.MethodWorkspaceWillDeleteFiles:
		if self.WorkspaceWillDeleteFiles != nil {
			validMethod = true
			var params protocol.DeleteFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.WorkspaceWillDeleteFiles(context, &params)
			}
		}

	case protocol.MethodWorkspaceDidDeleteFiles:
		if self.WorkspaceDidDeleteFiles != nil {
			validMethod = true
			var params protocol.DeleteFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.WorkspaceDidDeleteFiles(context, &params)
			}
		}

	// Text Document Synchronization

	case protocol.MethodTextDocumentDidOpen:
		if self.TextDocumentDidOpen != nil {
			validMethod = true
			var params protocol.DidOpenTextDocumentParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.TextDocumentDidOpen(context, &params)
			}
		}

	case protocol.MethodTextDocumentDidChange:
		if self.TextDocumentDidChange != nil {
			validMethod = true
			var params protocol.DidChangeTextDocumentParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.TextDocumentDidChange(context, &params)
			}
		}

	case protocol.MethodTextDocumentWillSave:
		if self.TextDocumentWillSave != nil {
			validMethod = true
			var params protocol.WillSaveTextDocumentParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.TextDocumentWillSave(context, &params)
			}
		}

	case protocol.MethodTextDocumentWillSaveWaitUntil:
		if self.TextDocumentWillSaveWaitUntil != nil {
			validMethod = true
			var params protocol.WillSaveTextDocumentParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentWillSaveWaitUntil(context, &params)
			}
		}

	case protocol.MethodTextDocumentDidSave:
		if self.TextDocumentDidSave != nil {
			validMethod = true
			var params protocol.DidSaveTextDocumentParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.TextDocumentDidSave(context, &params)
			}
		}

	case protocol.MethodTextDocumentDidClose:
		if self.TextDocumentDidClose != nil {
			validMethod = true
			var params protocol.DidCloseTextDocumentParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.TextDocumentDidClose(context, &params)
			}
		}

	// Language Features

	case protocol.MethodTextDocumentCompletion:
		if self.TextDocumentCompletion != nil {
			validMethod = true
			var params protocol.CompletionParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentCompletion(context, &params)
			}
		}

	case protocol.MethodCompletionItemResolve:
		if self.CompletionItemResolve != nil {
			validMethod = true
			var params protocol.CompletionItem
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.CompletionItemResolve(context, &params)
			}
		}

	case protocol.MethodTextDocumentHover:
		if self.TextDocumentHover != nil {
			validMethod = true
			var params protocol.HoverParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentHover(context, &params)
			}
		}

	case protocol.MethodTextDocumentSignatureHelp:
		if self.TextDocumentSignatureHelp != nil {
			validMethod = true
			var params protocol.SignatureHelpParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentSignatureHelp(context, &params)
			}
		}

	case protocol.MethodTextDocumentDeclaration:
		if self.TextDocumentDeclaration != nil {
			validMethod = true
			var params protocol.DeclarationParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentDeclaration(context, &params)
			}
		}

	case protocol.MethodTextDocumentDefinition:
		if self.TextDocumentDefinition != nil {
			validMethod = true
			var params protocol.DefinitionParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentDefinition(context, &params)
			}
		}

	case protocol.MethodTextDocumentTypeDefinition:
		if self.TextDocumentTypeDefinition != nil {
			validMethod = true
			var params protocol.TypeDefinitionParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentTypeDefinition(context, &params)
			}
		}

	case protocol.MethodTextDocumentImplementation:
		if self.TextDocumentImplementation != nil {
			validMethod = true
			var params protocol.ImplementationParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentImplementation(context, &params)
			}
		}

	case protocol.MethodTextDocumentReferences:
		if self.TextDocumentReferences != nil {
			validMethod = true
			var params protocol.ReferenceParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentReferences(context, &params)
			}
		}

	case protocol.MethodTextDocumentDocumentHighlight:
		if self.TextDocumentDocumentHighlight != nil {
			validMethod = true
			var params protocol.DocumentHighlightParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentDocumentHighlight(context, &params)
			}
		}

	case protocol.MethodTextDocumentDocumentSymbol:
		if self.TextDocumentDocumentSymbol != nil {
			validMethod = true
			var params protocol.DocumentSymbolParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentDocumentSymbol(context, &params)
			}
		}

	case protocol.MethodTextDocumentCodeAction:
		if self.TextDocumentCodeAction != nil {
			validMethod = true
			var params protocol.CodeActionParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentCodeAction(context, &params)
			}
		}

	case protocol.MethodCodeActionResolve:
		if self.CodeActionResolve != nil {
			validMethod = true
			var params protocol.CodeAction
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.CodeActionResolve(context, &params)
			}
		}

	case protocol.MethodTextDocumentCodeLens:
		if self.TextDocumentCodeLens != nil {
			validMethod = true
			var params protocol.CodeLensParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentCodeLens(context, &params)
			}
		}

	case protocol.MethodCodeLensResolve:
		if self.TextDocumentDidClose != nil {
			validMethod = true
			var params protocol.CodeLens
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.CodeLensResolve(context, &params)
			}
		}

	case protocol.MethodTextDocumentDocumentLink:
		if self.TextDocumentDocumentLink != nil {
			validMethod = true
			var params protocol.DocumentLinkParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentDocumentLink(context, &params)
			}
		}

	case protocol.MethodDocumentLinkResolve:
		if self.DocumentLinkResolve != nil {
			validMethod = true
			var params protocol.DocumentLink
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.DocumentLinkResolve(context, &params)
			}
		}

	case protocol.MethodTextDocumentColor:
		if self.TextDocumentColor != nil {
			validMethod = true
			var params protocol.DocumentColorParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentColor(context, &params)
			}
		}

	case protocol.MethodTextDocumentColorPresentation:
		if self.TextDocumentColorPresentation != nil {
			validMethod = true
			var params protocol.ColorPresentationParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentColorPresentation(context, &params)
			}
		}

	case protocol.MethodTextDocumentFormatting:
		if self.TextDocumentFormatting != nil {
			validMethod = true
			var params protocol.DocumentFormattingParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentFormatting(context, &params)
			}
		}

	case protocol.MethodTextDocumentRangeFormatting:
		if self.TextDocumentRangeFormatting != nil {
			validMethod = true
			var params protocol.DocumentRangeFormattingParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentRangeFormatting(context, &params)
			}
		}

	case protocol.MethodTextDocumentOnTypeFormatting:
		if self.TextDocumentOnTypeFormatting != nil {
			validMethod = true
			var params protocol.DocumentOnTypeFormattingParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentOnTypeFormatting(context, &params)
			}
		}

	case protocol.MethodTextDocumentRename:
		if self.TextDocumentRename != nil {
			validMethod = true
			var params protocol.RenameParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentRename(context, &params)
			}
		}

	case protocol.MethodTextDocumentPrepareRename:
		if self.TextDocumentPrepareRename != nil {
			validMethod = true
			var params protocol.PrepareRenameParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentPrepareRename(context, &params)
			}
		}

	case protocol.MethodTextDocumentFoldingRange:
		if self.TextDocumentFoldingRange != nil {
			validMethod = true
			var params protocol.FoldingRangeParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentFoldingRange(context, &params)
			}
		}

	case protocol.MethodTextDocumentSelectionRange:
		if self.TextDocumentSelectionRange != nil {
			validMethod = true
			var params protocol.SelectionRangeParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentSelectionRange(context, &params)
			}
		}

	case protocol.MethodTextDocumentPrepareCallHierarchy:
		if self.TextDocumentPrepareCallHierarchy != nil {
			validMethod = true
			var params protocol.CallHierarchyPrepareParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentPrepareCallHierarchy(context, &params)
			}
		}

	case protocol.MethodCallHierarchyIncomingCalls:
		if self.CallHierarchyIncomingCalls != nil {
			validMethod = true
			var params protocol.CallHierarchyIncomingCallsParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.CallHierarchyIncomingCalls(context, &params)
			}
		}

	case protocol.MethodCallHierarchyOutgoingCalls:
		if self.CallHierarchyOutgoingCalls != nil {
			validMethod = true
			var params protocol.CallHierarchyOutgoingCallsParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.CallHierarchyOutgoingCalls(context, &params)
			}
		}

	case protocol.MethodTextDocumentSemanticTokensFull:
		if self.TextDocumentSemanticTokensFull != nil {
			validMethod = true
			var params protocol.SemanticTokensParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentSemanticTokensFull(context, &params)
			}
		}

	case protocol.MethodTextDocumentSemanticTokensFullDelta:
		if self.TextDocumentSemanticTokensFullDelta != nil {
			validMethod = true
			var params protocol.SemanticTokensDeltaParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentSemanticTokensFullDelta(context, &params)
			}
		}

	case protocol.MethodTextDocumentSemanticTokensRange:
		if self.TextDocumentSemanticTokensRange != nil {
			validMethod = true
			var params protocol.SemanticTokensRangeParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentSemanticTokensRange(context, &params)
			}
		}

	case protocol.MethodWorkspaceSemanticTokensRefresh:
		if self.WorkspaceSemanticTokensRefresh != nil {
			validMethod = true
			validParams = true
			err = self.WorkspaceSemanticTokensRefresh(context)
		}

	case protocol.MethodTextDocumentLinkedEditingRange:
		if self.TextDocumentLinkedEditingRange != nil {
			validMethod = true
			var params protocol.LinkedEditingRangeParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentLinkedEditingRange(context, &params)
			}
		}

	case protocol.MethodTextDocumentMoniker:
		if self.TextDocumentMoniker != nil {
			validMethod = true
			var params protocol.MonikerParams
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
		if _, ok := capabilities.TextDocumentSync.(*protocol.TextDocumentSyncOptions); !ok {
			capabilities.TextDocumentSync = &protocol.TextDocumentSyncOptions{}
		}
		capabilities.TextDocumentSync.(*protocol.TextDocumentSyncOptions).OpenClose = &protocol.True
	}

	if self.TextDocumentDidChange != nil {
		if _, ok := capabilities.TextDocumentSync.(*protocol.TextDocumentSyncOptions); !ok {
			capabilities.TextDocumentSync = &protocol.TextDocumentSyncOptions{}
		}
		// This can be overriden to TextDocumentSyncKindFull
		value := protocol.TextDocumentSyncKindIncremental
		capabilities.TextDocumentSync.(*protocol.TextDocumentSyncOptions).Change = &value
	}

	if self.TextDocumentWillSave != nil {
		if _, ok := capabilities.TextDocumentSync.(*protocol.TextDocumentSyncOptions); !ok {
			capabilities.TextDocumentSync = &protocol.TextDocumentSyncOptions{}
		}
		capabilities.TextDocumentSync.(*protocol.TextDocumentSyncOptions).WillSave = &protocol.True
	}

	if self.TextDocumentWillSaveWaitUntil != nil {
		if _, ok := capabilities.TextDocumentSync.(*protocol.TextDocumentSyncOptions); !ok {
			capabilities.TextDocumentSync = &protocol.TextDocumentSyncOptions{}
		}
		capabilities.TextDocumentSync.(*protocol.TextDocumentSyncOptions).WillSaveWaitUntil = &protocol.True
	}

	if self.TextDocumentDidSave != nil {
		if _, ok := capabilities.TextDocumentSync.(*protocol.TextDocumentSyncOptions); !ok {
			capabilities.TextDocumentSync = &protocol.TextDocumentSyncOptions{}
		}
		capabilities.TextDocumentSync.(*protocol.TextDocumentSyncOptions).Save = &protocol.True
	}

	if self.TextDocumentCompletion != nil {
		capabilities.CompletionProvider = &protocol.CompletionOptions{}
	}

	if self.TextDocumentHover != nil {
		capabilities.HoverProvider = true
	}

	if self.TextDocumentSignatureHelp != nil {
		capabilities.SignatureHelpProvider = &protocol.SignatureHelpOptions{}
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
		capabilities.CodeLensProvider = &protocol.CodeLensOptions{}
	}

	if self.TextDocumentDocumentLink != nil {
		capabilities.DocumentLinkProvider = &protocol.DocumentLinkOptions{}
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
		capabilities.DocumentOnTypeFormattingProvider = &protocol.DocumentOnTypeFormattingOptions{}
	}

	if self.TextDocumentRename != nil {
		capabilities.RenameProvider = true
	}

	if self.TextDocumentFoldingRange != nil {
		capabilities.FoldingRangeProvider = true
	}

	if self.WorkspaceExecuteCommand != nil {
		capabilities.ExecuteCommandProvider = &protocol.ExecuteCommandOptions{}
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
		if _, ok := capabilities.SemanticTokensProvider.(*protocol.SemanticTokensOptions); !ok {
			capabilities.SemanticTokensProvider = &protocol.SemanticTokensOptions{}
		}
		if self.TextDocumentSemanticTokensFullDelta != nil {
			capabilities.SemanticTokensProvider.(*protocol.SemanticTokensOptions).Full = &protocol.SemanticDelta{}
			capabilities.SemanticTokensProvider.(*protocol.SemanticTokensOptions).Full.(*protocol.SemanticDelta).Delta = &protocol.True
		} else {
			capabilities.SemanticTokensProvider.(*protocol.SemanticTokensOptions).Full = true
		}
	}

	if self.TextDocumentSemanticTokensRange != nil {
		if _, ok := capabilities.SemanticTokensProvider.(*protocol.SemanticTokensOptions); !ok {
			capabilities.SemanticTokensProvider = &protocol.SemanticTokensOptions{}
		}
		capabilities.SemanticTokensProvider.(*protocol.SemanticTokensOptions).Range = true
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
			capabilities.Workspace = &protocol.ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &protocol.ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.DidCreate = &protocol.FileOperationRegistrationOptions{
			Filters: []protocol.FileOperationFilter{},
		}
	}

	if self.WorkspaceWillCreateFiles != nil {
		if capabilities.Workspace == nil {
			capabilities.Workspace = &protocol.ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &protocol.ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.WillCreate = &protocol.FileOperationRegistrationOptions{
			Filters: []protocol.FileOperationFilter{},
		}
	}

	if self.WorkspaceDidRenameFiles != nil {
		capabilities.RenameProvider = true
		if capabilities.Workspace == nil {
			capabilities.Workspace = &protocol.ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &protocol.ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.DidRename = &protocol.FileOperationRegistrationOptions{
			Filters: []protocol.FileOperationFilter{},
		}
	}

	if self.WorkspaceWillRenameFiles != nil {
		capabilities.RenameProvider = true
		if capabilities.Workspace == nil {
			capabilities.Workspace = &protocol.ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &protocol.ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.WillRename = &protocol.FileOperationRegistrationOptions{
			Filters: []protocol.FileOperationFilter{},
		}
	}

	if self.WorkspaceDidDeleteFiles != nil {
		if capabilities.Workspace == nil {
			capabilities.Workspace = &protocol.ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &protocol.ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.DidDelete = &protocol.FileOperationRegistrationOptions{
			Filters: []protocol.FileOperationFilter{},
		}
	}

	if self.WorkspaceWillDeleteFiles != nil {
		if capabilities.Workspace == nil {
			capabilities.Workspace = &protocol.ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &protocol.ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.WillDelete = &protocol.FileOperationRegistrationOptions{
			Filters: []protocol.FileOperationFilter{},
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
