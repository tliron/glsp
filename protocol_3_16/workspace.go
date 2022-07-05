package protocol

import "github.com/tliron/glsp"

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#workspace_workspaceFolders

const ServerWorkspaceWorkspaceFolders = Method("workspace/workspaceFolders")

type WorkspaceFoldersServerCapabilities struct {
	/**
	 * The server has support for workspace folders
	 */
	Supported *bool `json:"supported"`

	/**
	 * Whether the server wants to receive workspace folder
	 * change notifications.
	 *
	 * If a string is provided, the string is treated as an ID
	 * under which the notification is registered on the client
	 * side. The ID can be used to unregister for these events
	 * using the `client/unregisterCapability` request.
	 */
	ChangeNotifications *BoolOrString `json:"changeNotifications,omitempty"`
}

type WorkspaceFolder struct {
	/**
	 * The associated URI for this workspace folder.
	 */
	URI DocumentUri `json:"uri"`

	/**
	 * The name of the workspace folder. Used to refer to this
	 * workspace folder in the user interface.
	 */
	Name string `json:"name"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#workspace_didChangeWorkspaceFolders

const MethodWorkspaceDidChangeWorkspaceFolders = Method("workspace/didChangeWorkspaceFolders")

type WorkspaceDidChangeWorkspaceFoldersFunc func(context *glsp.Context, params *DidChangeWorkspaceFoldersParams) error

type DidChangeWorkspaceFoldersParams struct {
	/**
	 * The actual workspace folder change event.
	 */
	Event WorkspaceFoldersChangeEvent `json:"event"`
}

/**
 * The workspace folder change event.
 */
type WorkspaceFoldersChangeEvent struct {
	/**
	 * The array of added workspace folders
	 */
	Added []WorkspaceFolder `json:"added"`

	/**
	 * The array of the removed workspace folders
	 */
	Removed []WorkspaceFolder `json:"removed"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#workspace_didChangeConfiguration

type DidChangeConfigurationClientCapabilities struct {
	/**
	 * Did change configuration notification supports dynamic registration.
	 */
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

const MethodWorkspaceDidChangeConfiguration = Method("workspace/didChangeConfiguration")

type WorkspaceDidChangeConfigurationFunc func(context *glsp.Context, params *DidChangeConfigurationParams) error

type DidChangeConfigurationParams struct {
	/**
	 * The actual changed settings
	 */
	Settings any `json:"settings"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#workspace_configuration

const ServerWorkspaceConfiguration = Method("workspace/configuration")

type ConfigurationParams struct {
	Items []ConfigurationItem `json:"items"`
}

type ConfigurationItem struct {
	/**
	 * The scope to get the configuration section for.
	 */
	ScopeURI *DocumentUri `json:"scopeUri,omitempty"`

	/**
	 * The configuration section asked for.
	 */
	Section *string `json:"section,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#workspace_didChangeWatchedFiles

type DidChangeWatchedFilesClientCapabilities struct {
	/**
	 * Did change watched files notification supports dynamic registration.
	 * Please note that the current protocol doesn't support static
	 * configuration for file changes from the server side.
	 */
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

/**
 * Describe options to be used when registering for file system change events.
 */
type DidChangeWatchedFilesRegistrationOptions struct {
	/**
	 * The watchers to register.
	 */
	Watchers []FileSystemWatcher `json:"watchers"`
}

type FileSystemWatcher struct {
	/**
	 * The  glob pattern to watch.
	 *
	 * Glob patterns can have the following syntax:
	 * - `*` to match one or more characters in a path segment
	 * - `?` to match on one character in a path segment
	 * - `**` to match any number of path segments, including none
	 * - `{}` to group conditions (e.g. `**​/*.{ts,js}` matches all TypeScript
	 *   and JavaScript files)
	 * - `[]` to declare a range of characters to match in a path segment
	 *   (e.g., `example.[0-9]` to match on `example.0`, `example.1`, …)
	 * - `[!...]` to negate a range of characters to match in a path segment
	 *   (e.g., `example.[!0-9]` to match on `example.a`, `example.b`, but not
	 *   `example.0`)
	 */
	GlobPattern string `json:"globPattern"`

	/**
	 * The kind of events of interest. If omitted it defaults
	 * to WatchKind.Create | WatchKind.Change | WatchKind.Delete
	 * which is 7.
	 */
	Kind *UInteger `json:"kind,omitempty"`
}

const (
	/**
	 * Interested in create events.
	 */
	WatchKindCreate = UInteger(1)

	/**
	 * Interested in change events
	 */
	WatchKindChange = UInteger(2)

	/**
	 * Interested in delete events
	 */
	WatchKindDelete = UInteger(4)
)

const MethodWorkspaceDidChangeWatchedFiles = Method("workspace/didChangeWatchedFiles")

type WorkspaceDidChangeWatchedFilesFunc func(context *glsp.Context, params *DidChangeWatchedFilesParams) error

type DidChangeWatchedFilesParams struct {
	/**
	 * The actual file events.
	 */
	Changes []FileEvent `json:"changes"`
}

/**
 * An event describing a file change.
 */
type FileEvent struct {
	/**
	 * The file's URI.
	 */
	URI DocumentUri `json:"uri"`
	/**
	 * The change type.
	 */
	Type UInteger `json:"type"`
}

/**
 * The file event type.
 */
const (
	/**
	 * The file got created.
	 */
	FileChangeTypeCreated = UInteger(1)

	/**
	 * The file got changed.
	 */
	FileChangeTypeChanged = UInteger(2)

	/**
	 * The file got deleted.
	 */
	FileChangeTypeDeleted = UInteger(3)
)

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#workspace_symbol

type WorkspaceSymbolClientCapabilities struct {
	/**
	 * Symbol request supports dynamic registration.
	 */
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`

	/**
	 * Specific capabilities for the `SymbolKind` in the `workspace/symbol`
	 * request.
	 */
	SymbolKind *struct {
		/**
		 * The symbol kind values the client supports. When this
		 * property exists the client also guarantees that it will
		 * handle values outside its set gracefully and falls back
		 * to a default value when unknown.
		 *
		 * If this property is not present the client only supports
		 * the symbol kinds from `File` to `Array` as defined in
		 * the initial version of the protocol.
		 */
		ValueSet []SymbolKind `json:"valueSet,omitempty"`
	} `json:"symbolKind,omitempty"`

	/**
	 * The client supports tags on `SymbolInformation`.
	 * Clients supporting tags have to handle unknown tags gracefully.
	 *
	 * @since 3.16.0
	 */
	TagSupport *struct {
		/**
		 * The tags supported by the client.
		 */
		ValueSet []SymbolTag `json:"valueSet"`
	} `json:"tagSupport,omitempty"`
}

type WorkspaceSymbolOptions struct {
	WorkDoneProgressOptions
}

type WorkspaceSymbolRegistrationOptions struct {
	WorkspaceSymbolOptions
}

const MethodWorkspaceSymbol = Method("workspace/symbol")

type WorkspaceSymbolFunc func(context *glsp.Context, params *WorkspaceSymbolParams) ([]SymbolInformation, error)

type WorkspaceSymbolParams struct {
	WorkDoneProgressParams
	PartialResultParams

	/**
	 * A query string to filter symbols by. Clients may send an empty
	 * string here to request all symbols.
	 */
	Query string `json:"query"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#workspace_executeCommand

type ExecuteCommandClientCapabilities struct {
	/**
	 * Execute command supports dynamic registration.
	 */
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

type ExecuteCommandOptions struct {
	WorkDoneProgressOptions

	/**
	 * The commands to be executed on the server
	 */
	Commands []string `json:"commands"`
}

/**
 * Execute command registration options.
 */
type ExecuteCommandRegistrationOptions struct {
	ExecuteCommandOptions
}

const MethodWorkspaceExecuteCommand = Method("workspace/executeCommand")

type WorkspaceExecuteCommandFunc func(context *glsp.Context, params *ExecuteCommandParams) (any, error)

type ExecuteCommandParams struct {
	WorkDoneProgressParams

	/**
	 * The identifier of the actual command handler.
	 */
	Command string `json:"command"`

	/**
	 * Arguments that the command should be invoked with.
	 */
	Arguments []any `json:"arguments,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#workspace_applyEdit

const ServerWorkspaceApplyEdit = Method("workspace/applyEdit")

type ApplyWorkspaceEditParams struct {
	/**
	 * An optional label of the workspace edit. This label is
	 * presented in the user interface for example on an undo
	 * stack to undo the workspace edit.
	 */
	Label *string `json:"label,omitempty"`

	/**
	 * The edits to apply.
	 */
	Edit WorkspaceEdit `json:"edit"`
}

type ApplyWorkspaceEditResponse struct {
	/**
	 * Indicates whether the edit was applied or not.
	 */
	Applied bool `json:"applied"`

	/**
	 * An optional textual description for why the edit was not applied.
	 * This may be used by the server for diagnostic logging or to provide
	 * a suitable error for a request that triggered the edit.
	 */
	FailureReason *string `json:"failureReason,omitempty"`

	/**
	 * Depending on the client's failure handling strategy `failedChange`
	 * might contain the index of the change that failed. This property is
	 * only available if the client signals a `failureHandlingStrategy`
	 * in its client capabilities.
	 */
	FailedChange *UInteger `json:"failedChange,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#workspace_willCreateFiles

/**
 * The options to register for file operations.
 *
 * @since 3.16.0
 */
type FileOperationRegistrationOptions struct {
	/**
	 * The actual filters.
	 */
	Filters []FileOperationFilter `json:"filters"`
}

type FileOperationPatternKind string

/**
 * A pattern kind describing if a glob pattern matches a file a folder or
 * both.
 *
 * @since 3.16.0
 */
const (
	/**
	 * The pattern matches a file only.
	 */
	FileOperationPatternKindFile = FileOperationPatternKind("file")

	/**
	 * The pattern matches a folder only.
	 */
	FileOperationPatternKindFolder = FileOperationPatternKind("folder")
)

/**
 * Matching options for the file operation pattern.
 *
 * @since 3.16.0
 */
type FileOperationPatternOptions struct {

	/**
	 * The pattern should be matched ignoring casing.
	 */
	IgnoreCase *bool `json:"ignoreCase,omitempty"`
}

/**
 * A pattern to describe in which file operation requests or notifications
 * the server is interested in.
 *
 * @since 3.16.0
 */
type FileOperationPattern struct {
	/**
	 * The glob pattern to match. Glob patterns can have the following syntax:
	 * - `*` to match one or more characters in a path segment
	 * - `?` to match on one character in a path segment
	 * - `**` to match any number of path segments, including none
	 * - `{}` to group conditions (e.g. `**​/*.{ts,js}` matches all TypeScript
	 *   and JavaScript files)
	 * - `[]` to declare a range of characters to match in a path segment
	 *   (e.g., `example.[0-9]` to match on `example.0`, `example.1`, …)
	 * - `[!...]` to negate a range of characters to match in a path segment
	 *   (e.g., `example.[!0-9]` to match on `example.a`, `example.b`, but
	 *   not `example.0`)
	 */
	Glob string `json:"glob"`

	/**
	 * Whether to match files or folders with this pattern.
	 *
	 * Matches both if undefined.
	 */
	Matches *FileOperationPatternKind `json:"matches,omitempty"`

	/**
	 * Additional options used during matching.
	 */
	Options *FileOperationPatternOptions `json:"options,omitempty"`
}

/**
 * A filter to describe in which file operation requests or notifications
 * the server is interested in.
 *
 * @since 3.16.0
 */
type FileOperationFilter struct {
	/**
	 * A Uri like `file` or `untitled`.
	 */
	Scheme *string `json:"scheme,omitempty"`

	/**
	 * The actual file operation pattern.
	 */
	Pattern FileOperationPattern `json:"pattern"`
}

const MethodWorkspaceWillCreateFiles = Method("workspace/willCreateFiles")

type WorkspaceWillCreateFilesFunc func(context *glsp.Context, params *CreateFilesParams) (*WorkspaceEdit, error)

/**
 * The parameters sent in notifications/requests for user-initiated creation
 * of files.
 *
 * @since 3.16.0
 */
type CreateFilesParams struct {
	/**
	 * An array of all files/folders created in this operation.
	 */
	Files []FileCreate `json:"files"`
}

/**
 * Represents information on a file/folder create.
 *
 * @since 3.16.0
 */
type FileCreate struct {
	/**
	 * A file:// URI for the location of the file/folder being created.
	 */
	URI string `json:"uri"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#workspace_didCreateFiles

const MethodWorkspaceDidCreateFiles = Method("workspace/didCreateFiles")

type WorkspaceDidCreateFilesFunc func(context *glsp.Context, params *CreateFilesParams) error

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#workspace_willRenameFiles

const MethodWorkspaceWillRenameFiles = Method("workspace/willRenameFiles")

type WorkspaceWillRenameFilesFunc func(context *glsp.Context, params *RenameFilesParams) (*WorkspaceEdit, error)

/**
 * The parameters sent in notifications/requests for user-initiated renames
 * of files.
 *
 * @since 3.16.0
 */
type RenameFilesParams struct {
	/**
	 * An array of all files/folders renamed in this operation. When a folder
	 * is renamed, only the folder will be included, and not its children.
	 */
	Files []FileRename `json:"files"`
}

/**
 * Represents information on a file/folder rename.
 *
 * @since 3.16.0
 */
type FileRename struct {
	/**
	 * A file:// URI for the original location of the file/folder being renamed.
	 */
	OldURI string `json:"oldUri"`

	/**
	 * A file:// URI for the new location of the file/folder being renamed.
	 */
	NewURI string `json:"newUri"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#workspace_didRenameFiles

const MethodWorkspaceDidRenameFiles = Method("workspace/didRenameFiles")

type WorkspaceDidRenameFilesFunc func(context *glsp.Context, params *RenameFilesParams) error

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#workspace_willDeleteFiles

const MethodWorkspaceWillDeleteFiles = Method("workspace/willDeleteFiles")

type WorkspaceWillDeleteFilesFunc func(context *glsp.Context, params *DeleteFilesParams) (*WorkspaceEdit, error)

/**
 * The parameters sent in notifications/requests for user-initiated deletes
 * of files.
 *
 * @since 3.16.0
 */
type DeleteFilesParams struct {
	/**
	 * An array of all files/folders deleted in this operation.
	 */
	Files []FileDelete `json:"files"`
}

/**
 * Represents information on a file/folder delete.
 *
 * @since 3.16.0
 */
type FileDelete struct {
	/**
	 * A file:// URI for the location of the file/folder being deleted.
	 */
	URI string `json:"uri"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#workspace_didDeleteFiles

const MethodWorkspaceDidDeleteFiles = Method("workspace/didDeleteFiles")

type WorkspaceDidDeleteFilesFunc func(context *glsp.Context, params *DeleteFilesParams) error

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16/#textDocument_semanticTokens
const MethodWorkspaceSemanticTokensRefresh = Method("workspace/semanticTokens/refresh")

type WorkspaceSemanticTokensRefreshFunc func(context *glsp.Context) error

type SemanticTokensWorkspaceClientCapabilities struct {
	/**
	 * Whether the client implementation supports a refresh request sent from
	 * the server to the client.
	 *
	 * Note that this event is global and will force the client to refresh all
	 * semantic tokens currently shown. It should be used with absolute care
	 * and is useful for situation where a server for example detect a project
	 * wide change that requires such a calculation.
	 */
	RefreshSupport *bool `json:"refreshSupport,omitempty"`
}
