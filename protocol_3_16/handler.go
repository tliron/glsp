package protocol

import (
	"encoding/json"
	"errors"
	"sync"

	"github.com/tliron/glsp"
)

type Handler struct {
	// Base Protocol
	CancelRequest CancelRequestFunc
	Progress      ProgressFunc

	// General Messages
	Initialize  InitializeFunc
	Initialized InitializedFunc
	Shutdown    ShutdownFunc
	Exit        ExitFunc
	LogTrace    LogTraceFunc
	SetTrace    SetTraceFunc

	// Window
	WindowWorkDoneProgressCancel WindowWorkDoneProgressCancelFunc

	// Workspace
	WorkspaceDidChangeWorkspaceFolders WorkspaceDidChangeWorkspaceFoldersFunc
	WorkspaceDidChangeConfiguration    WorkspaceDidChangeConfigurationFunc
	WorkspaceDidChangeWatchedFiles     WorkspaceDidChangeWatchedFilesFunc
	WorkspaceSymbol                    WorkspaceSymbolFunc
	WorkspaceExecuteCommand            WorkspaceExecuteCommandFunc
	WorkspaceWillCreateFiles           WorkspaceWillCreateFilesFunc
	WorkspaceDidCreateFiles            WorkspaceDidCreateFilesFunc
	WorkspaceWillRenameFiles           WorkspaceWillRenameFilesFunc
	WorkspaceDidRenameFiles            WorkspaceDidRenameFilesFunc
	WorkspaceWillDeleteFiles           WorkspaceWillDeleteFilesFunc
	WorkspaceDidDeleteFiles            WorkspaceDidDeleteFilesFunc
	WorkspaceSemanticTokensRefresh     WorkspaceSemanticTokensRefreshFunc

	// Text Document Synchronization
	TextDocumentDidOpen           TextDocumentDidOpenFunc
	TextDocumentDidChange         TextDocumentDidChangeFunc
	TextDocumentWillSave          TextDocumentWillSaveFunc
	TextDocumentWillSaveWaitUntil TextDocumentWillSaveWaitUntilFunc
	TextDocumentDidSave           TextDocumentDidSaveFunc
	TextDocumentDidClose          TextDocumentDidCloseFunc

	// Language Features
	TextDocumentCompletion              TextDocumentCompletionFunc
	CompletionItemResolve               CompletionItemResolveFunc
	TextDocumentHover                   TextDocumentHoverFunc
	TextDocumentSignatureHelp           TextDocumentSignatureHelpFunc
	TextDocumentDeclaration             TextDocumentDeclarationFunc
	TextDocumentDefinition              TextDocumentDefinitionFunc
	TextDocumentTypeDefinition          TextDocumentTypeDefinitionFunc
	TextDocumentImplementation          TextDocumentImplementationFunc
	TextDocumentReferences              TextDocumentReferencesFunc
	TextDocumentDocumentHighlight       TextDocumentDocumentHighlightFunc
	TextDocumentDocumentSymbol          TextDocumentDocumentSymbolFunc
	TextDocumentCodeAction              TextDocumentCodeActionFunc
	CodeActionResolve                   CodeActionResolveFunc
	TextDocumentCodeLens                TextDocumentCodeLensFunc
	CodeLensResolve                     CodeLensResolveFunc
	TextDocumentDocumentLink            TextDocumentDocumentLinkFunc
	DocumentLinkResolve                 DocumentLinkResolveFunc
	TextDocumentColor                   TextDocumentColorFunc
	TextDocumentColorPresentation       TextDocumentColorPresentationFunc
	TextDocumentFormatting              TextDocumentFormattingFunc
	TextDocumentRangeFormatting         TextDocumentRangeFormattingFunc
	TextDocumentOnTypeFormatting        TextDocumentOnTypeFormattingFunc
	TextDocumentRename                  TextDocumentRenameFunc
	TextDocumentPrepareRename           TextDocumentPrepareRenameFunc
	TextDocumentFoldingRange            TextDocumentFoldingRangeFunc
	TextDocumentSelectionRange          TextDocumentSelectionRangeFunc
	TextDocumentPrepareCallHierarchy    TextDocumentPrepareCallHierarchyFunc
	CallHierarchyIncomingCalls          CallHierarchyIncomingCallsFunc
	CallHierarchyOutgoingCalls          CallHierarchyOutgoingCallsFunc
	TextDocumentSemanticTokensFull      TextDocumentSemanticTokensFullFunc
	TextDocumentSemanticTokensFullDelta TextDocumentSemanticTokensFullDeltaFunc
	TextDocumentSemanticTokensRange     TextDocumentSemanticTokensRangeFunc
	TextDocumentLinkedEditingRange      TextDocumentLinkedEditingRangeFunc
	TextDocumentMoniker                 TextDocumentMonikerFunc

	initialized bool
	lock        sync.Mutex
}

// glsp.Handler interface
func (self *Handler) Handle(context *glsp.Context) (r any, validMethod bool, validParams bool, err error) {
	if !self.IsInitialized() && (context.Method != MethodInitialize) {
		return nil, true, true, errors.New("server not initialized")
	}

	switch context.Method {
	// Base Protocol

	case MethodCancelRequest:
		if self.CancelRequest != nil {
			validMethod = true
			var params CancelParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.CancelRequest(context, &params)
			}
		}

	case MethodProgress:
		if self.Progress != nil {
			validMethod = true
			var params ProgressParams
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

	case MethodInitialized:
		if self.Initialized != nil {
			validMethod = true
			var params InitializedParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.Initialized(context, &params)
			}
		}

	case MethodShutdown:
		self.SetInitialized(false)
		if self.Shutdown != nil {
			validMethod = true
			validParams = true
			err = self.Shutdown(context)
		}

	case MethodExit:
		// Note that the server will close the connection after we handle it here
		if self.Exit != nil {
			validMethod = true
			validParams = true
			err = self.Exit(context)
		}

	case MethodLogTrace:
		if self.LogTrace != nil {
			validMethod = true
			var params LogTraceParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.LogTrace(context, &params)
			}
		}

	case MethodSetTrace:
		if self.SetTrace != nil {
			validMethod = true
			var params SetTraceParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.SetTrace(context, &params)
			}
		}

	// Window

	case MethodWindowWorkDoneProgressCancel:
		if self.WindowWorkDoneProgressCancel != nil {
			validMethod = true
			var params WorkDoneProgressCancelParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.WindowWorkDoneProgressCancel(context, &params)
			}
		}

	// Workspace

	case MethodWorkspaceDidChangeWorkspaceFolders:
		if self.WorkspaceDidChangeWorkspaceFolders != nil {
			validMethod = true
			var params DidChangeWorkspaceFoldersParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.WorkspaceDidChangeWorkspaceFolders(context, &params)
			}
		}

	case MethodWorkspaceDidChangeConfiguration:
		if self.WorkspaceDidChangeConfiguration != nil {
			validMethod = true
			var params DidChangeConfigurationParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.WorkspaceDidChangeConfiguration(context, &params)
			}
		}

	case MethodWorkspaceDidChangeWatchedFiles:
		if self.WorkspaceDidChangeWatchedFiles != nil {
			validMethod = true
			var params DidChangeWatchedFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.WorkspaceDidChangeWatchedFiles(context, &params)
			}
		}

	case MethodWorkspaceSymbol:
		if self.WorkspaceSymbol != nil {
			validMethod = true
			var params WorkspaceSymbolParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.WorkspaceSymbol(context, &params)
			}
		}

	case MethodWorkspaceExecuteCommand:
		if self.WorkspaceExecuteCommand != nil {
			validMethod = true
			var params ExecuteCommandParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.WorkspaceExecuteCommand(context, &params)
			}
		}

	case MethodWorkspaceWillCreateFiles:
		if self.WorkspaceWillCreateFiles != nil {
			validMethod = true
			var params CreateFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.WorkspaceWillCreateFiles(context, &params)
			}
		}

	case MethodWorkspaceDidCreateFiles:
		if self.WorkspaceDidCreateFiles != nil {
			validMethod = true
			var params CreateFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.WorkspaceDidCreateFiles(context, &params)
			}
		}

	case MethodWorkspaceWillRenameFiles:
		if self.WorkspaceWillRenameFiles != nil {
			validMethod = true
			var params RenameFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.WorkspaceWillRenameFiles(context, &params)
			}
		}

	case MethodWorkspaceDidRenameFiles:
		if self.WorkspaceDidRenameFiles != nil {
			validMethod = true
			var params RenameFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.WorkspaceDidRenameFiles(context, &params)
			}
		}

	case MethodWorkspaceWillDeleteFiles:
		if self.WorkspaceWillDeleteFiles != nil {
			validMethod = true
			var params DeleteFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.WorkspaceWillDeleteFiles(context, &params)
			}
		}

	case MethodWorkspaceDidDeleteFiles:
		if self.WorkspaceDidDeleteFiles != nil {
			validMethod = true
			var params DeleteFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.WorkspaceDidDeleteFiles(context, &params)
			}
		}

	// Text Document Synchronization

	case MethodTextDocumentDidOpen:
		if self.TextDocumentDidOpen != nil {
			validMethod = true
			var params DidOpenTextDocumentParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.TextDocumentDidOpen(context, &params)
			}
		}

	case MethodTextDocumentDidChange:
		if self.TextDocumentDidChange != nil {
			validMethod = true
			var params DidChangeTextDocumentParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.TextDocumentDidChange(context, &params)
			}
		}

	case MethodTextDocumentWillSave:
		if self.TextDocumentWillSave != nil {
			validMethod = true
			var params WillSaveTextDocumentParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.TextDocumentWillSave(context, &params)
			}
		}

	case MethodTextDocumentWillSaveWaitUntil:
		if self.TextDocumentWillSaveWaitUntil != nil {
			validMethod = true
			var params WillSaveTextDocumentParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentWillSaveWaitUntil(context, &params)
			}
		}

	case MethodTextDocumentDidSave:
		if self.TextDocumentDidSave != nil {
			validMethod = true
			var params DidSaveTextDocumentParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.TextDocumentDidSave(context, &params)
			}
		}

	case MethodTextDocumentDidClose:
		if self.TextDocumentDidClose != nil {
			validMethod = true
			var params DidCloseTextDocumentParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = self.TextDocumentDidClose(context, &params)
			}
		}

	// Language Features

	case MethodTextDocumentCompletion:
		if self.TextDocumentCompletion != nil {
			validMethod = true
			var params CompletionParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentCompletion(context, &params)
			}
		}

	case MethodCompletionItemResolve:
		if self.CompletionItemResolve != nil {
			validMethod = true
			var params CompletionItem
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.CompletionItemResolve(context, &params)
			}
		}

	case MethodTextDocumentHover:
		if self.TextDocumentHover != nil {
			validMethod = true
			var params HoverParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentHover(context, &params)
			}
		}

	case MethodTextDocumentSignatureHelp:
		if self.TextDocumentSignatureHelp != nil {
			validMethod = true
			var params SignatureHelpParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentSignatureHelp(context, &params)
			}
		}

	case MethodTextDocumentDeclaration:
		if self.TextDocumentDeclaration != nil {
			validMethod = true
			var params DeclarationParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentDeclaration(context, &params)
			}
		}

	case MethodTextDocumentDefinition:
		if self.TextDocumentDefinition != nil {
			validMethod = true
			var params DefinitionParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentDefinition(context, &params)
			}
		}

	case MethodTextDocumentTypeDefinition:
		if self.TextDocumentTypeDefinition != nil {
			validMethod = true
			var params TypeDefinitionParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentTypeDefinition(context, &params)
			}
		}

	case MethodTextDocumentImplementation:
		if self.TextDocumentImplementation != nil {
			validMethod = true
			var params ImplementationParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentImplementation(context, &params)
			}
		}

	case MethodTextDocumentReferences:
		if self.TextDocumentReferences != nil {
			validMethod = true
			var params ReferenceParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentReferences(context, &params)
			}
		}

	case MethodTextDocumentDocumentHighlight:
		if self.TextDocumentDocumentHighlight != nil {
			validMethod = true
			var params DocumentHighlightParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentDocumentHighlight(context, &params)
			}
		}

	case MethodTextDocumentDocumentSymbol:
		if self.TextDocumentDocumentSymbol != nil {
			validMethod = true
			var params DocumentSymbolParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentDocumentSymbol(context, &params)
			}
		}

	case MethodTextDocumentCodeAction:
		if self.TextDocumentCodeAction != nil {
			validMethod = true
			var params CodeActionParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentCodeAction(context, &params)
			}
		}

	case MethodCodeActionResolve:
		if self.CodeActionResolve != nil {
			validMethod = true
			var params CodeAction
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.CodeActionResolve(context, &params)
			}
		}

	case MethodTextDocumentCodeLens:
		if self.TextDocumentCodeLens != nil {
			validMethod = true
			var params CodeLensParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentCodeLens(context, &params)
			}
		}

	case MethodCodeLensResolve:
		if self.TextDocumentDidClose != nil {
			validMethod = true
			var params CodeLens
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.CodeLensResolve(context, &params)
			}
		}

	case MethodTextDocumentDocumentLink:
		if self.TextDocumentDocumentLink != nil {
			validMethod = true
			var params DocumentLinkParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentDocumentLink(context, &params)
			}
		}

	case MethodDocumentLinkResolve:
		if self.DocumentLinkResolve != nil {
			validMethod = true
			var params DocumentLink
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.DocumentLinkResolve(context, &params)
			}
		}

	case MethodTextDocumentColor:
		if self.TextDocumentColor != nil {
			validMethod = true
			var params DocumentColorParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentColor(context, &params)
			}
		}

	case MethodTextDocumentColorPresentation:
		if self.TextDocumentColorPresentation != nil {
			validMethod = true
			var params ColorPresentationParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentColorPresentation(context, &params)
			}
		}

	case MethodTextDocumentFormatting:
		if self.TextDocumentFormatting != nil {
			validMethod = true
			var params DocumentFormattingParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentFormatting(context, &params)
			}
		}

	case MethodTextDocumentRangeFormatting:
		if self.TextDocumentRangeFormatting != nil {
			validMethod = true
			var params DocumentRangeFormattingParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentRangeFormatting(context, &params)
			}
		}

	case MethodTextDocumentOnTypeFormatting:
		if self.TextDocumentOnTypeFormatting != nil {
			validMethod = true
			var params DocumentOnTypeFormattingParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentOnTypeFormatting(context, &params)
			}
		}

	case MethodTextDocumentRename:
		if self.TextDocumentRename != nil {
			validMethod = true
			var params RenameParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentRename(context, &params)
			}
		}

	case MethodTextDocumentPrepareRename:
		if self.TextDocumentPrepareRename != nil {
			validMethod = true
			var params PrepareRenameParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentPrepareRename(context, &params)
			}
		}

	case MethodTextDocumentFoldingRange:
		if self.TextDocumentFoldingRange != nil {
			validMethod = true
			var params FoldingRangeParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentFoldingRange(context, &params)
			}
		}

	case MethodTextDocumentSelectionRange:
		if self.TextDocumentSelectionRange != nil {
			validMethod = true
			var params SelectionRangeParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentSelectionRange(context, &params)
			}
		}

	case MethodTextDocumentPrepareCallHierarchy:
		if self.TextDocumentPrepareCallHierarchy != nil {
			validMethod = true
			var params CallHierarchyPrepareParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentPrepareCallHierarchy(context, &params)
			}
		}

	case MethodCallHierarchyIncomingCalls:
		if self.CallHierarchyIncomingCalls != nil {
			validMethod = true
			var params CallHierarchyIncomingCallsParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.CallHierarchyIncomingCalls(context, &params)
			}
		}

	case MethodCallHierarchyOutgoingCalls:
		if self.CallHierarchyOutgoingCalls != nil {
			validMethod = true
			var params CallHierarchyOutgoingCallsParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.CallHierarchyOutgoingCalls(context, &params)
			}
		}

	case MethodTextDocumentSemanticTokensFull:
		if self.TextDocumentSemanticTokensFull != nil {
			validMethod = true
			var params SemanticTokensParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentSemanticTokensFull(context, &params)
			}
		}

	case MethodTextDocumentSemanticTokensFullDelta:
		if self.TextDocumentSemanticTokensFullDelta != nil {
			validMethod = true
			var params SemanticTokensDeltaParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentSemanticTokensFullDelta(context, &params)
			}
		}

	case MethodTextDocumentSemanticTokensRange:
		if self.TextDocumentSemanticTokensRange != nil {
			validMethod = true
			var params SemanticTokensRangeParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentSemanticTokensRange(context, &params)
			}
		}

	case MethodWorkspaceSemanticTokensRefresh:
		if self.WorkspaceSemanticTokensRefresh != nil {
			validMethod = true
			validParams = true
			err = self.WorkspaceSemanticTokensRefresh(context)
		}

	case MethodTextDocumentLinkedEditingRange:
		if self.TextDocumentLinkedEditingRange != nil {
			validMethod = true
			var params LinkedEditingRangeParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentLinkedEditingRange(context, &params)
			}
		}

	case MethodTextDocumentMoniker:
		if self.TextDocumentMoniker != nil {
			validMethod = true
			var params MonikerParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = self.TextDocumentMoniker(context, &params)
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
		if _, ok := capabilities.TextDocumentSync.(*TextDocumentSyncOptions); !ok {
			capabilities.TextDocumentSync = &TextDocumentSyncOptions{}
		}
		capabilities.TextDocumentSync.(*TextDocumentSyncOptions).OpenClose = &True
	}

	if self.TextDocumentDidChange != nil {
		if _, ok := capabilities.TextDocumentSync.(*TextDocumentSyncOptions); !ok {
			capabilities.TextDocumentSync = &TextDocumentSyncOptions{}
		}
		// This can be overriden to TextDocumentSyncKindFull
		value := TextDocumentSyncKindIncremental
		capabilities.TextDocumentSync.(*TextDocumentSyncOptions).Change = &value
	}

	if self.TextDocumentWillSave != nil {
		if _, ok := capabilities.TextDocumentSync.(*TextDocumentSyncOptions); !ok {
			capabilities.TextDocumentSync = &TextDocumentSyncOptions{}
		}
		capabilities.TextDocumentSync.(*TextDocumentSyncOptions).WillSave = &True
	}

	if self.TextDocumentWillSaveWaitUntil != nil {
		if _, ok := capabilities.TextDocumentSync.(*TextDocumentSyncOptions); !ok {
			capabilities.TextDocumentSync = &TextDocumentSyncOptions{}
		}
		capabilities.TextDocumentSync.(*TextDocumentSyncOptions).WillSaveWaitUntil = &True
	}

	if self.TextDocumentDidSave != nil {
		if _, ok := capabilities.TextDocumentSync.(*TextDocumentSyncOptions); !ok {
			capabilities.TextDocumentSync = &TextDocumentSyncOptions{}
		}
		capabilities.TextDocumentSync.(*TextDocumentSyncOptions).Save = &True
	}

	if self.TextDocumentCompletion != nil {
		capabilities.CompletionProvider = &CompletionOptions{}
	}

	if self.TextDocumentHover != nil {
		capabilities.HoverProvider = true
	}

	if self.TextDocumentSignatureHelp != nil {
		capabilities.SignatureHelpProvider = &SignatureHelpOptions{}
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
		capabilities.CodeLensProvider = &CodeLensOptions{}
	}

	if self.TextDocumentDocumentLink != nil {
		capabilities.DocumentLinkProvider = &DocumentLinkOptions{}
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
		capabilities.DocumentOnTypeFormattingProvider = &DocumentOnTypeFormattingOptions{}
	}

	if self.TextDocumentRename != nil {
		capabilities.RenameProvider = true
	}

	if self.TextDocumentFoldingRange != nil {
		capabilities.FoldingRangeProvider = true
	}

	if self.WorkspaceExecuteCommand != nil {
		capabilities.ExecuteCommandProvider = &ExecuteCommandOptions{}
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
		if _, ok := capabilities.SemanticTokensProvider.(*SemanticTokensOptions); !ok {
			capabilities.SemanticTokensProvider = &SemanticTokensOptions{}
		}
		if self.TextDocumentSemanticTokensFullDelta != nil {
			capabilities.SemanticTokensProvider.(*SemanticTokensOptions).Full = &SemanticDelta{}
			capabilities.SemanticTokensProvider.(*SemanticTokensOptions).Full.(*SemanticDelta).Delta = &True
		} else {
			capabilities.SemanticTokensProvider.(*SemanticTokensOptions).Full = true
		}
	}

	if self.TextDocumentSemanticTokensRange != nil {
		if _, ok := capabilities.SemanticTokensProvider.(*SemanticTokensOptions); !ok {
			capabilities.SemanticTokensProvider = &SemanticTokensOptions{}
		}
		capabilities.SemanticTokensProvider.(*SemanticTokensOptions).Range = true
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
			capabilities.Workspace = &ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.DidCreate = &FileOperationRegistrationOptions{
			Filters: []FileOperationFilter{},
		}
	}

	if self.WorkspaceWillCreateFiles != nil {
		if capabilities.Workspace == nil {
			capabilities.Workspace = &ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.WillCreate = &FileOperationRegistrationOptions{
			Filters: []FileOperationFilter{},
		}
	}

	if self.WorkspaceDidRenameFiles != nil {
		capabilities.RenameProvider = true
		if capabilities.Workspace == nil {
			capabilities.Workspace = &ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.DidRename = &FileOperationRegistrationOptions{
			Filters: []FileOperationFilter{},
		}
	}

	if self.WorkspaceWillRenameFiles != nil {
		capabilities.RenameProvider = true
		if capabilities.Workspace == nil {
			capabilities.Workspace = &ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.WillRename = &FileOperationRegistrationOptions{
			Filters: []FileOperationFilter{},
		}
	}

	if self.WorkspaceDidDeleteFiles != nil {
		if capabilities.Workspace == nil {
			capabilities.Workspace = &ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.DidDelete = &FileOperationRegistrationOptions{
			Filters: []FileOperationFilter{},
		}
	}

	if self.WorkspaceWillDeleteFiles != nil {
		if capabilities.Workspace == nil {
			capabilities.Workspace = &ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.WillDelete = &FileOperationRegistrationOptions{
			Filters: []FileOperationFilter{},
		}
	}

	return capabilities
}
