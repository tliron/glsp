package protocol

import (
	"encoding/json"

	"github.com/tliron/glsp"
)

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#initialize

const MethodInitialize = Method("initialize")

// Returns: InitializeResult | InitializeError
type InitializeFunc func(context *glsp.Context, params *InitializeParams) (any, error)

type InitializeParams struct {
	WorkDoneProgressParams

	/**
	 * The process Id of the parent process that started the server. Is null if
	 * the process has not been started by another process. If the parent
	 * process is not alive then the server should exit (see exit notification)
	 * its process.
	 */
	ProcessID *Integer `json:"processId"`

	/**
	 * Information about the client
	 *
	 * @since 3.15.0
	 */
	ClientInfo *struct {
		/**
		 * The name of the client as defined by the client.
		 */
		Name string `json:"name"`

		/**
		 * The client's version as defined by the client.
		 */
		Version *string `json:"version,omitempty"`
	} `json:"clientInfo,omitempty"`

	/**
	 * The locale the client is currently showing the user interface
	 * in. This must not necessarily be the locale of the operating
	 * system.
	 *
	 * Uses IETF language tags as the value's syntax
	 * (See https://en.wikipedia.org/wiki/IETF_language_tag)
	 *
	 * @since 3.16.0
	 */
	Locale *string `json:"locale,omitempty"`

	/**
	 * The rootPath of the workspace. Is null
	 * if no folder is open.
	 *
	 * @deprecated in favour of `rootUri`.
	 */
	RootPath *string `json:"rootPath,omitempty"`

	/**
	 * The rootUri of the workspace. Is null if no
	 * folder is open. If both `rootPath` and `rootUri` are set
	 * `rootUri` wins.
	 *
	 * @deprecated in favour of `workspaceFolders`
	 */
	RootURI *DocumentUri `json:"rootUri"`

	/**
	 * User provided initialization options.
	 */
	InitializationOptions any `json:"initializationOptions,omitempty"`

	/**
	 * The capabilities provided by the client (editor or tool)
	 */
	Capabilities ClientCapabilities `json:"capabilities"`

	/**
	 * The initial trace setting. If omitted trace is disabled ('off').
	 */
	Trace *TraceValue `json:"trace,omitempty"`

	/**
	 * The workspace folders configured in the client when the server starts.
	 * This property is only available if the client supports workspace folders.
	 * It can be `null` if the client supports workspace folders but none are
	 * configured.
	 *
	 * @since 3.6.0
	 */
	WorkspaceFolders []WorkspaceFolder `json:"workspaceFolders,omitempty"`
}

/**
 * Text document specific client capabilities.
 */
type TextDocumentClientCapabilities struct {
	Synchronization *TextDocumentSyncClientCapabilities `json:"synchronization,omitempty"`

	/**
	 * Capabilities specific to the `textDocument/completion` request.
	 */
	Completion *CompletionClientCapabilities `json:"completion,omitempty"`

	/**
	 * Capabilities specific to the `textDocument/hover` request.
	 */
	Hover *HoverClientCapabilities `json:"hover,omitempty"`

	/**
	 * Capabilities specific to the `textDocument/signatureHelp` request.
	 */
	SignatureHelp *SignatureHelpClientCapabilities `json:"signatureHelp,omitempty"`

	/**
	 * Capabilities specific to the `textDocument/declaration` request.
	 *
	 * @since 3.14.0
	 */
	Declaration *DeclarationClientCapabilities `json:"declaration,omitempty"`

	/**
	 * Capabilities specific to the `textDocument/definition` request.
	 */
	Definition *DefinitionClientCapabilities `json:"definition,omitempty"`

	/**
	 * Capabilities specific to the `textDocument/typeDefinition` request.
	 *
	 * @since 3.6.0
	 */
	TypeDefinition *TypeDefinitionClientCapabilities `json:"typeDefinition,omitempty"`

	/**
	 * Capabilities specific to the `textDocument/implementation` request.
	 *
	 * @since 3.6.0
	 */
	Implementation *ImplementationClientCapabilities `json:"implementation,omitempty"`

	/**
	 * Capabilities specific to the `textDocument/references` request.
	 */
	References *ReferenceClientCapabilities `json:"references,omitempty"`

	/**
	 * Capabilities specific to the `textDocument/documentHighlight` request.
	 */
	DocumentHighlight *DocumentHighlightClientCapabilities `json:"documentHighlight,omitempty"`

	/**
	 * Capabilities specific to the `textDocument/documentSymbol` request.
	 */
	DocumentSymbol *DocumentSymbolClientCapabilities `json:"documentSymbol,omitempty"`

	/**
	 * Capabilities specific to the `textDocument/codeAction` request.
	 */
	CodeAction *CodeActionClientCapabilities `json:"codeAction,omitempty"`

	/**
	 * Capabilities specific to the `textDocument/codeLens` request.
	 */
	CodeLens *CodeLensClientCapabilities `json:"codeLens,omitempty"`

	/**
	 * Capabilities specific to the `textDocument/documentLink` request.
	 */
	DocumentLink *DocumentLinkClientCapabilities `json:"documentLink,omitempty"`

	/**
	 * Capabilities specific to the `textDocument/documentColor` and the
	 * `textDocument/colorPresentation` request.
	 *
	 * @since 3.6.0
	 */
	ColorProvider *DocumentColorClientCapabilities `json:"colorProvider,omitempty"`

	/**
	 * Capabilities specific to the `textDocument/formatting` request.
	 */
	Formatting *DocumentFormattingClientCapabilities `json:"formatting,omitempty"`

	/**
	 * Capabilities specific to the `textDocument/rangeFormatting` request.
	 */
	RangeFormatting *DocumentRangeFormattingClientCapabilities `json:"rangeFormatting,omitempty"`

	/** request.
	 * Capabilities specific to the `textDocument/onTypeFormatting` request.
	 */
	OnTypeFormatting *DocumentOnTypeFormattingClientCapabilities `json:"onTypeFormatting,omitempty"`

	/**
	 * Capabilities specific to the `textDocument/rename` request.
	 */
	Rename *RenameClientCapabilities `json:"rename,omitempty"`

	/**
	 * Capabilities specific to the `textDocument/publishDiagnostics`
	 * notification.
	 */
	PublishDiagnostics *PublishDiagnosticsClientCapabilities `json:"publishDiagnostics,omitempty"`

	/**
	 * Capabilities specific to the `textDocument/foldingRange` request.
	 *
	 * @since 3.10.0
	 */
	FoldingRange *FoldingRangeClientCapabilities `json:"foldingRange,omitempty"`

	/**
	 * Capabilities specific to the `textDocument/selectionRange` request.
	 *
	 * @since 3.15.0
	 */
	SelectionRange *SelectionRangeClientCapabilities `json:"selectionRange,omitempty"`

	/**
	 * Capabilities specific to the `textDocument/linkedEditingRange` request.
	 *
	 * @since 3.16.0
	 */
	LinkedEditingRange *LinkedEditingRangeClientCapabilities `json:"linkedEditingRange,omitempty"`

	/**
	 * Capabilities specific to the various call hierarchy requests.
	 *
	 * @since 3.16.0
	 */
	CallHierarchy *CallHierarchyClientCapabilities `json:"callHierarchy,omitempty"`

	/**
	 * Capabilities specific to the various semantic token requests.
	 *
	 * @since 3.16.0
	 */
	SemanticTokens *SemanticTokensClientCapabilities `json:"semanticTokens,omitempty"`

	/**
	 * Capabilities specific to the `textDocument/moniker` request.
	 *
	 * @since 3.16.0
	 */
	Moniker *MonikerClientCapabilities `json:"moniker,omitempty"`
}

type ClientCapabilities struct {
	/**
	 * Workspace specific client capabilities.
	 */
	Workspace *struct {
		/**
		 * The client supports applying batch edits
		 * to the workspace by supporting the request
		 * 'workspace/applyEdit'
		 */
		ApplyEdit *bool `json:"applyEdit,omitempty"`

		/**
		 * Capabilities specific to `WorkspaceEdit`s
		 */
		WorkspaceEdit *WorkspaceEditClientCapabilities `json:"workspaceEdit,omitempty"`

		/**
		 * Capabilities specific to the `workspace/didChangeConfiguration`
		 * notification.
		 */
		DidChangeConfiguration *DidChangeConfigurationClientCapabilities `json:"didChangeConfiguration,omitempty"`

		/**
		 * Capabilities specific to the `workspace/didChangeWatchedFiles`
		 * notification.
		 */
		DidChangeWatchedFiles *DidChangeWatchedFilesClientCapabilities `json:"didChangeWatchedFiles,omitempty"`

		/**
		 * Capabilities specific to the `workspace/symbol` request.
		 */
		Symbol *WorkspaceSymbolClientCapabilities `json:"symbol,omitempty"`

		/**
		 * Capabilities specific to the `workspace/executeCommand` request.
		 */
		ExecuteCommand *ExecuteCommandClientCapabilities `json:"executeCommand,omitempty"`

		/**
		 * The client has support for workspace folders.
		 *
		 * @since 3.6.0
		 */
		WorkspaceFolders *bool `json:"workspaceFolders,omitempty"`

		/**
		 * The client supports `workspace/configuration` requests.
		 *
		 * @since 3.6.0
		 */
		Configuration *bool `json:"configuration,omitempty"`

		/**
		 * Capabilities specific to the semantic token requests scoped to the
		 * workspace.
		 *
		 * @since 3.16.0
		 */
		SemanticTokens *SemanticTokensWorkspaceClientCapabilities `json:"semanticTokens,omitempty"`

		/**
		 * Capabilities specific to the code lens requests scoped to the
		 * workspace.
		 *
		 * @since 3.16.0
		 */
		CodeLens *CodeLensWorkspaceClientCapabilities `json:"codeLens,omitempty"`

		/**
		 * The client has support for file requests/notifications.
		 *
		 * @since 3.16.0
		 */
		FileOperations *struct {
			/**
			 * Whether the client supports dynamic registration for file
			 * requests/notifications.
			 */
			DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`

			/**
			 * The client has support for sending didCreateFiles notifications.
			 */
			DidCreate *bool `json:"didCreate,omitempty"`

			/**
			 * The client has support for sending willCreateFiles requests.
			 */
			WillCreate *bool `json:"willCreate,omitempty"`

			/**
			 * The client has support for sending didRenameFiles notifications.
			 */
			DidRename *bool `json:"didRename,omitempty"`

			/**
			 * The client has support for sending willRenameFiles requests.
			 */
			WillRename *bool `json:"willRename,omitempty"`

			/**
			 * The client has support for sending didDeleteFiles notifications.
			 */
			DidDelete *bool `json:"didDelete,omitempty"`

			/**
			 * The client has support for sending willDeleteFiles requests.
			 */
			WillDelete *bool `json:"willDelete,omitempty"`
		} `json:"fileOperations,omitempty"`
	} `json:"workspace,omitempty"`

	/**
	 * Text document specific client capabilities.
	 */
	TextDocument *TextDocumentClientCapabilities `json:"textDocument,omitempty"`

	/**
	 * Window specific client capabilities.
	 */
	Window *struct {
		/**
		 * Whether client supports handling progress notifications. If set
		 * servers are allowed to report in `workDoneProgress` property in the
		 * request specific server capabilities.
		 *
		 * @since 3.15.0
		 */
		WorkDoneProgress *bool `json:"workDoneProgress,omitempty"`

		/**
		 * Capabilities specific to the showMessage request
		 *
		 * @since 3.16.0
		 */
		ShowMessage *ShowMessageRequestClientCapabilities `json:"showMessage,omitempty"`

		/**
		 * Client capabilities for the show document request.
		 *
		 * @since 3.16.0
		 */
		ShowDocument *ShowDocumentClientCapabilities `json:"showDocument,omitempty"`
	} `json:"window,omitempty"`

	/**
	 * General client capabilities.
	 *
	 * @since 3.16.0
	 */
	General *struct {
		/**
		 * Client capabilities specific to regular expressions.
		 *
		 * @since 3.16.0
		 */
		RegularExpressions *RegularExpressionsClientCapabilities `json:"regularExpressions,omitempty"`

		/**
		 * Client capabilities specific to the client's markdown parser.
		 *
		 * @since 3.16.0
		 */
		Markdown *MarkdownClientCapabilities `json:"markdown,omitempty"`
	} `json:"general,omitempty"`

	/**
	 * Experimental client capabilities.
	 */
	Experimental any `json:"experimental,omitempty"`
}

func (self *ClientCapabilities) SupportsSymbolKind(kind SymbolKind) bool {
	var kinds []SymbolKind
	if (self.TextDocument != nil) && (self.TextDocument.DocumentSymbol != nil) && (self.TextDocument.DocumentSymbol.SymbolKind != nil) {
		kinds = self.TextDocument.DocumentSymbol.SymbolKind.ValueSet
	}
	if kinds == nil {
		return kind <= 19
	} else {
		for _, kind_ := range kinds {
			if kind == kind_ {
				return true
			}
		}
		return false
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
	ServerInfo *InitializeResultServerInfo `json:"serverInfo,omitempty"`
}

type InitializeResultServerInfo struct {
	/**
	 * The name of the server as defined by the server.
	 */
	Name string `json:"name"`

	/**
	 * The server's version as defined by the server.
	 */
	Version *string `json:"version,omitempty"`
}

/**
 * Known error codes for an `InitializeError`;
 */
type InitializeErrorCode Integer

const (
	/**
	 * If the protocol version provided by the client can't be handled by the
	 * server.
	 *
	 * @deprecated This initialize error got replaced by client capabilities.
	 * There is no version handshake in version 3.0x
	 */
	InitializeErrorCodeUnknownProtocolVersion = InitializeErrorCode(1)
)

type InitializeError struct {
	/**
	 * Indicates whether the client execute the following retry logic:
	 * (1) show the message provided by the ResponseError to the user
	 * (2) user selects retry or cancel
	 * (3) if user selected retry the initialize method is sent again.
	 */
	Retry bool `json:"retry"`
}

type ServerCapabilities struct {
	/**
	 * Defines how text documents are synced. Is either a detailed structure
	 * defining each notification or for backwards compatibility the
	 * TextDocumentSyncKind number. If omitted it defaults to
	 * `TextDocumentSyncKind.None`.
	 */
	TextDocumentSync any `json:"textDocumentSync,omitempty"` // nil | TextDocumentSyncOptions | TextDocumentSyncKind

	/**
	 * The server provides completion support.
	 */
	CompletionProvider *CompletionOptions `json:"completionProvider,omitempty"`

	/**
	 * The server provides hover support.
	 */
	HoverProvider any `json:"hoverProvider,omitempty"` // nil | bool | HoverOptions

	/**
	 * The server provides signature help support.
	 */
	SignatureHelpProvider *SignatureHelpOptions `json:"signatureHelpProvider,omitempty"`

	/**
	 * The server provides go to declaration support.
	 *
	 * @since 3.14.0
	 */
	DeclarationProvider any `json:"declarationProvider,omitempty"` // nil | bool | DeclarationOptions | DeclarationRegistrationOptions

	/**
	 * The server provides goto definition support.
	 */
	DefinitionProvider any `json:"definitionProvider,omitempty"` // nil | bool | DefinitionOptions

	/**
	 * The server provides goto type definition support.
	 *
	 * @since 3.6.0
	 */
	TypeDefinitionProvider any `json:"typeDefinitionProvider,omitempty"` // nil | bool | TypeDefinitionOption | TypeDefinitionRegistrationOptions

	/**
	 * The server provides goto implementation support.
	 *
	 * @since 3.6.0
	 */
	ImplementationProvider any `json:"implementationProvider,omitempty"` // nil | bool | ImplementationOptions | ImplementationRegistrationOptions

	/**
	 * The server provides find references support.
	 */
	ReferencesProvider any `json:"referencesProvider,omitempty"` // nil | bool | ReferenceOptions

	/**
	 * The server provides document highlight support.
	 */
	DocumentHighlightProvider any `json:"documentHighlightProvider,omitempty"` // nil | bool | DocumentHighlightOptions

	/**
	 * The server provides document symbol support.
	 */
	DocumentSymbolProvider any `json:"documentSymbolProvider,omitempty"` // nil | bool | DocumentSymbolOptions

	/**
	 * The server provides code actions. The `CodeActionOptions` return type is
	 * only valid if the client signals code action literal support via the
	 * property `textDocument.codeAction.codeActionLiteralSupport`.
	 */
	CodeActionProvider any `json:"codeActionProvider,omitempty"` // nil | bool | CodeActionOptions

	/**
	 * The server provides code lens.
	 */
	CodeLensProvider *CodeLensOptions `json:"codeLensProvider,omitempty"`

	/**
	 * The server provides document link support.
	 */
	DocumentLinkProvider *DocumentLinkOptions `json:"documentLinkProvider,omitempty"`

	/**
	 * The server provides color provider support.
	 *
	 * @since 3.6.0
	 */
	ColorProvider any `json:"colorProvider,omitempty"` // nil | bool | DocumentColorOptions | DocumentColorRegistrationOptions

	/**
	 * The server provides document formatting.
	 */
	DocumentFormattingProvider any `json:"documentFormattingProvider,omitempty"` // nil | bool | DocumentFormattingOptions

	/**
	 * The server provides document range formatting.
	 */
	DocumentRangeFormattingProvider any `json:"documentRangeFormattingProvider,omitempty"` // nil | bool | DocumentRangeFormattingOptions

	/**
	 * The server provides document formatting on typing.
	 */
	DocumentOnTypeFormattingProvider *DocumentOnTypeFormattingOptions `json:"documentOnTypeFormattingProvider,omitempty"`

	/**
	 * The server provides rename support. RenameOptions may only be
	 * specified if the client states that it supports
	 * `prepareSupport` in its initial `initialize` request.
	 */
	RenameProvider any `json:"renameProvider,omitempty"` // nil | bool | RenameOptions

	/**
	 * The server provides folding provider support.
	 *
	 * @since 3.10.0
	 */
	FoldingRangeProvider any `json:"foldingRangeProvider,omitempty"` // nil | bool | FoldingRangeOptions | FoldingRangeRegistrationOptions

	/**
	 * The server provides execute command support.
	 */
	ExecuteCommandProvider *ExecuteCommandOptions `json:"executeCommandProvider,omitempty"`

	/**
	 * The server provides selection range support.
	 *
	 * @since 3.15.0
	 */
	SelectionRangeProvider any `json:"selectionRangeProvider,omitempty"` // nil | bool | SelectionRangeOptions | SelectionRangeRegistrationOptions

	/**
	 * The server provides linked editing range support.
	 *
	 * @since 3.16.0
	 */
	LinkedEditingRangeProvider any `json:"linkedEditingRangeProvider,omitempty"` // nil | bool | LinkedEditingRangeOptions | LinkedEditingRangeRegistrationOptions

	/**
	 * The server provides call hierarchy support.
	 *
	 * @since 3.16.0
	 */
	CallHierarchyProvider any `json:"callHierarchyProvider,omitempty"` // nil | bool | CallHierarchyOptions | CallHierarchyRegistrationOptions

	/**
	 * The server provides semantic tokens support.
	 *
	 * @since 3.16.0
	 */
	SemanticTokensProvider any `json:"semanticTokensProvider,omitempty"` // nil | SemanticTokensOptions | SemanticTokensRegistrationOptions

	/**
	 * Whether server provides moniker support.
	 *
	 * @since 3.16.0
	 */
	MonikerProvider any `json:"monikerProvider,omitempty"` // nil | bool | MonikerOptions | MonikerRegistrationOptions

	/**
	 * The server provides workspace symbol support.
	 */
	WorkspaceSymbolProvider any `json:"workspaceSymbolProvider,omitempty"` // nil | bool | WorkspaceSymbolOptions

	/**
	 * Workspace specific server capabilities
	 */
	Workspace *ServerCapabilitiesWorkspace `json:"workspace,omitempty"`

	/**
	 * Experimental server capabilities.
	 */
	Experimental any `json:"experimental,omitempty"`
}

type ServerCapabilitiesWorkspace struct {
	/**
	 * The server supports workspace folder.
	 *
	 * @since 3.6.0
	 */
	WorkspaceFolders *WorkspaceFoldersServerCapabilities `json:"workspaceFolders,omitempty"`

	/**
	 * The server is interested in file notifications/requests.
	 *
	 * @since 3.16.0
	 */
	FileOperations *ServerCapabilitiesWorkspaceFileOperations `json:"fileOperations,omitempty"`
}

type ServerCapabilitiesWorkspaceFileOperations struct {
	/**
	 * The server is interested in receiving didCreateFiles
	 * notifications.
	 */
	DidCreate *FileOperationRegistrationOptions `json:"didCreate,omitempty"`

	/**
	 * The server is interested in receiving willCreateFiles requests.
	 */
	WillCreate *FileOperationRegistrationOptions `json:"willCreate,omitempty"`

	/**
	 * The server is interested in receiving didRenameFiles
	 * notifications.
	 */
	DidRename *FileOperationRegistrationOptions `json:"didRename,omitempty"`

	/**
	 * The server is interested in receiving willRenameFiles requests.
	 */
	WillRename *FileOperationRegistrationOptions `json:"willRename,omitempty"`

	/**
	 * The server is interested in receiving didDeleteFiles file
	 * notifications.
	 */
	DidDelete *FileOperationRegistrationOptions `json:"didDelete,omitempty"`

	/**
	 * The server is interested in receiving willDeleteFiles file
	 * requests.
	 */
	WillDelete *FileOperationRegistrationOptions `json:"willDelete,omitempty"`
}

// json.Unmarshaler interface
func (self *ServerCapabilities) UnmarshalJSON(data []byte) error {
	var value struct {
		TextDocumentSync                 json.RawMessage                  `json:"textDocumentSync,omitempty"` // nil | TextDocumentSyncOptions | TextDocumentSyncKind
		CompletionProvider               *CompletionOptions               `json:"completionProvider,omitempty"`
		HoverProvider                    json.RawMessage                  `json:"hoverProvider,omitempty"` // nil | bool | HoverOptions
		SignatureHelpProvider            *SignatureHelpOptions            `json:"signatureHelpProvider,omitempty"`
		DeclarationProvider              json.RawMessage                  `json:"declarationProvider,omitempty"`       // nil | bool | DeclarationOptions | DeclarationRegistrationOptions
		DefinitionProvider               json.RawMessage                  `json:"definitionProvider,omitempty"`        // nil | bool | DefinitionOptions
		TypeDefinitionProvider           json.RawMessage                  `json:"typeDefinitionProvider,omitempty"`    // nil | bool | TypeDefinitionOption | TypeDefinitionRegistrationOptions
		ImplementationProvider           json.RawMessage                  `json:"implementationProvider,omitempty"`    // nil | bool | ImplementationOptions | ImplementationRegistrationOptions
		ReferencesProvider               json.RawMessage                  `json:"referencesProvider,omitempty"`        // nil | bool | ReferenceOptions
		DocumentHighlightProvider        json.RawMessage                  `json:"documentHighlightProvider,omitempty"` // nil | bool | DocumentHighlightOptions
		DocumentSymbolProvider           json.RawMessage                  `json:"documentSymbolProvider,omitempty"`    // nil | bool | DocumentSymbolOptions
		CodeActionProvider               json.RawMessage                  `json:"codeActionProvider,omitempty"`        // nil | bool | CodeActionOptions
		CodeLensProvider                 *CodeLensOptions                 `json:"codeLensProvider,omitempty"`
		DocumentLinkProvider             *DocumentLinkOptions             `json:"documentLinkProvider,omitempty"`
		ColorProvider                    json.RawMessage                  `json:"colorProvider,omitempty"`                   // nil | bool | DocumentColorOptions | DocumentColorRegistrationOptions
		DocumentFormattingProvider       json.RawMessage                  `json:"documentFormattingProvider,omitempty"`      // nil | bool | DocumentFormattingOptions
		DocumentRangeFormattingProvider  json.RawMessage                  `json:"documentRangeFormattingProvider,omitempty"` // nil | bool | DocumentRangeFormattingOptions
		DocumentOnTypeFormattingProvider *DocumentOnTypeFormattingOptions `json:"documentOnTypeFormattingProvider,omitempty"`
		RenameProvider                   json.RawMessage                  `json:"renameProvider,omitempty"`       // nil | bool | RenameOptions
		FoldingRangeProvider             json.RawMessage                  `json:"foldingRangeProvider,omitempty"` // nil | bool | FoldingRangeOptions | FoldingRangeRegistrationOptions
		ExecuteCommandProvider           *ExecuteCommandOptions           `json:"executeCommandProvider,omitempty"`
		SelectionRangeProvider           json.RawMessage                  `json:"selectionRangeProvider,omitempty"`     // nil | bool | SelectionRangeOptions | SelectionRangeRegistrationOptions
		LinkedEditingRangeProvider       json.RawMessage                  `json:"linkedEditingRangeProvider,omitempty"` // nil | bool | LinkedEditingRangeOptions | LinkedEditingRangeRegistrationOptions
		CallHierarchyProvider            json.RawMessage                  `json:"callHierarchyProvider,omitempty"`      // nil | bool | CallHierarchyOptions | CallHierarchyRegistrationOptions
		SemanticTokensProvider           json.RawMessage                  `json:"semanticTokensProvider,omitempty"`     // nil | SemanticTokensOptions | SemanticTokensRegistrationOptions
		MonikerProvider                  json.RawMessage                  `json:"monikerProvider,omitempty"`            // nil | bool | MonikerOptions | MonikerRegistrationOptions
		WorkspaceSymbolProvider          json.RawMessage                  `json:"workspaceSymbolProvider,omitempty"`    // nil | bool | WorkspaceSymbolOptions
		Workspace                        *ServerCapabilitiesWorkspace     `json:"workspace,omitempty"`
		Experimental                     *any                             `json:"experimental,omitempty"`
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
			var value_ TextDocumentSyncOptions
			if err = json.Unmarshal(value.TextDocumentSync, &value_); err == nil {
				self.TextDocumentSync = value_
			} else {
				var value_ TextDocumentSyncKind
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
				var value_ HoverOptions
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
				var value_ DeclarationOptions
				if err = json.Unmarshal(value.DeclarationProvider, &value_); err == nil {
					self.DeclarationProvider = value_
				} else {
					var value_ DeclarationRegistrationOptions
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
				var value_ DefinitionOptions
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
				var value_ TypeDefinitionOptions
				if err = json.Unmarshal(value.TypeDefinitionProvider, &value_); err == nil {
					self.TypeDefinitionProvider = value_
				} else {
					var value_ TypeDefinitionRegistrationOptions
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
				var value_ ImplementationOptions
				if err = json.Unmarshal(value.ImplementationProvider, &value_); err == nil {
					self.ImplementationProvider = value_
				} else {
					var value_ ImplementationRegistrationOptions
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
				var value_ ReferenceOptions
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
				var value_ DocumentHighlightOptions
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
				var value_ DocumentSymbolOptions
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
				var value_ CodeActionOptions
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
				var value_ DocumentColorOptions
				if err = json.Unmarshal(value.ColorProvider, &value_); err == nil {
					self.ColorProvider = value_
				} else {
					var value_ DocumentColorRegistrationOptions
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
				var value_ DocumentFormattingOptions
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
				var value_ DocumentRangeFormattingOptions
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
				var value_ RenameOptions
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
				var value_ FoldingRangeOptions
				if err = json.Unmarshal(value.FoldingRangeProvider, &value_); err == nil {
					self.FoldingRangeProvider = value_
				} else {
					var value_ FoldingRangeRegistrationOptions
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
				var value_ SelectionRangeOptions
				if err = json.Unmarshal(value.SelectionRangeProvider, &value_); err == nil {
					self.SelectionRangeProvider = value_
				} else {
					var value_ SelectionRangeRegistrationOptions
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
				var value_ LinkedEditingRangeOptions
				if err = json.Unmarshal(value.LinkedEditingRangeProvider, &value_); err == nil {
					self.LinkedEditingRangeProvider = value_
				} else {
					var value_ LinkedEditingRangeRegistrationOptions
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
				var value_ CallHierarchyOptions
				if err = json.Unmarshal(value.CallHierarchyProvider, &value_); err == nil {
					self.CallHierarchyProvider = value_
				} else {
					var value_ CallHierarchyRegistrationOptions
					if err = json.Unmarshal(value.CallHierarchyProvider, &value_); err == nil {
						self.CallHierarchyProvider = value_
					} else {
						return err
					}
				}
			}
		}

		if value.SemanticTokensProvider != nil {
			var value_ SemanticTokensOptions
			if err = json.Unmarshal(value.SemanticTokensProvider, &value_); err == nil {
				self.SemanticTokensProvider = value_
			} else {
				var value_ SemanticTokensRegistrationOptions
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
				var value_ MonikerOptions
				if err = json.Unmarshal(value.MonikerProvider, &value_); err == nil {
					self.MonikerProvider = value_
				} else {
					var value_ MonikerRegistrationOptions
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
				var value_ WorkspaceSymbolOptions
				if err = json.Unmarshal(value.WorkspaceSymbolProvider, &value_); err == nil {
					self.WorkspaceSymbolProvider = value_
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

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#initialized

const MethodInitialized = Method("initialized")

type InitializedFunc func(context *glsp.Context, params *InitializedParams) error

type InitializedParams struct{}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#shutdown

const MethodShutdown = Method("shutdown")

type ShutdownFunc func(context *glsp.Context) error

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#exit

const MethodExit = Method("exit")

type ExitFunc func(context *glsp.Context) error

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#logTrace

const MethodLogTrace = Method("$/logTrace")

type LogTraceFunc func(context *glsp.Context, params *LogTraceParams) error

type LogTraceParams struct {
	/**
	 * The message to be logged.
	 */
	Message string `json:"message"`

	/**
	 * Additional information that can be computed if the `trace` configuration
	 * is set to `'verbose'`
	 */
	Verbose *string `json:"verbose,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#setTrace

const MethodSetTrace = Method("$/setTrace")

type SetTraceFunc func(context *glsp.Context, params *SetTraceParams) error

type SetTraceParams struct {
	/**
	 * The new value that should be assigned to the trace setting.
	 */
	Value TraceValue `json:"value"`
}
