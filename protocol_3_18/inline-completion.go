package protocol

import (
	"github.com/tliron/glsp"
	protocol316 "github.com/tliron/glsp/protocol_3_16"
)

/**
 * Client capabilities specific to inline completion requests.
 *
 * @since 3.18.0
 */
type InlineCompletionClientCapabilities struct {

	/**
	 * Whether implementation supports dynamic registration. If this is set to
	 * `true` the client supports the new
	 * `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
	 * return value for the corresponding server capability as well.
	 */
	DynamicRegistration bool `json:"dynamicRegistration"`
}

/**
 * Inline completion options.
 *
 * @since 3.18.0
 */
type InlineCompletionOptions struct {
	protocol316.WorkDoneProgressOptions
}

/**
 * Inline completion registration options.
 *
 * @since 3.18.0
 */
type InlineCompletionRegistrationOptions struct {
	protocol316.TextDocumentRegistrationOptions
	protocol316.StaticRegistrationOptions
	InlineCompletionOptions
}

const MethodTextDocumentInlineCompletion = protocol316.Method("textDocument/inlineCompletion")

// Returns: []InlineCompletionItem | InlineCompletionList | nil
type TextDocumentInlineCompletionFunc func(context *glsp.Context, params *InlineCompletionParams) (any, error)

/**
 * Parameters of the inline completion request.
 *
 * @since 3.18.0
 */
type InlineCompletionParams struct {
	protocol316.TextDocumentPositionParams
	protocol316.WorkDoneProgressParams

	Context InlineCompletionContext `json:"context"`
}

/**
 * Provides information about the context in which an inline completion was
 * requested.
 *
 * @since 3.18.0
 */
type InlineCompletionContext struct {
	/**
	 * Describes how the inline completion was triggered.
	 */
	InlineCompletionTriggerKind InlineCompletionTriggerKind `json:"inlineCompletionTriggerKind"`

	/**
	 * Provides information about the currently selected item in the
	 * autocomplete widget if it is visible.
	 *
	 * If set, provided inline completions must extend the text of the
	 * selected item and use the same range, otherwise they are not shown as
	 * preview.
	 * As an example, if the document text is `console.` and the selected item
	 * is `.log` replacing the `.` in the document, the inline completion must
	 * also replace `.` and start with `.log`, for example `.log()`.
	 *
	 * Inline completion providers are requested again whenever the selected
	 * item changes.
	 */
	SelectedCompletionInfo *SelectedCompletionInfo `json:"selectedCompletionInfo:omitempty"`
}

/**
 * The inline completion trigger kinds.
 *
 * @since 3.18.0
 */
type InlineCompletionTriggerKind protocol316.Integer

const (
	/**
	 * Completion was triggered explicitly by a user gesture.
	 * Return multiple completion items to enable cycling through them.
	 */
	InlineCompletionTriggerKindInvoked = InlineCompletionTriggerKind(1)

	/**
	 * Completion was triggered automatically while editing.
	 * It is sufficient to return a single completion item in this case.
	 */
	InlineCompletionTriggerKindAutomatic = InlineCompletionTriggerKind(2)
)

/**
 * @since 3.18.0
 */
type SelectedCompletionInfo struct {
	Range protocol316.Range `json:"range"`

	Text string `json:"text"`
}

/**
 * Represents a collection of inline completion items to be
 * presented in the editor.
 *
 * @since 3.18.0
 */
type InlineCompletionList struct {
	/**
	 * The inline completion items.
	 */
	InlineCompletionItem []InlineCompletionItem `json:"items"`
}

/**
 * An inline completion item represents a text snippet that is proposed inline
 * to complete text that is being typed.
 *
 * @since 3.18.0
 */
type InlineCompletionItem struct {
	/**
	 * The text to replace the range with. Must be set.
	 * Is used both for the preview and the accept operation.
	 */
	InsertText string `json:"insertText"`

	/**
	 * A text that is used to decide if this inline completion should be
	 * shown. When `falsy`, the InlineCompletionItem.insertText is used.
	 *
	 * An inline completion is shown if the text to replace is a prefix of the
	 * filter text.
	 */
	FilterText *string `json:"filterText,omitempty"`

	/**
	 * The range to replace.
	 * Must begin and end on the same line.
	 *
	 * Prefer replacements over insertions to provide a better experience when
	 * the user deletes typed text.
	 */
	Range *protocol316.Range `json:"range,omitempty"`

	/**
	 * An optional Command that is executed *after* inserting this
	 * completion.
	 */
	Command *protocol316.Command `json:"command,omitempty"`
}
