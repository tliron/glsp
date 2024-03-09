package protocol

import (
	"encoding/json"

	"github.com/tliron/glsp"
)

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#textDocument_completion

type CompletionClientCapabilities struct {
	/**
	 * Whether completion supports dynamic registration.
	 */
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`

	/**
	 * The client supports the following `CompletionItem` specific
	 * capabilities.
	 */
	CompletionItem *struct {
		/**
		 * Client supports snippets as insert text.
		 *
		 * A snippet can define tab stops and placeholders with `$1`, `$2`
		 * and `${3:foo}`. `$0` defines the final tab stop, it defaults to
		 * the end of the snippet. Placeholders with equal identifiers are
		 * linked, that is typing in one will update others too.
		 */
		SnippetSupport *bool `json:"snippetSupport,omitempty"`

		/**
		 * Client supports commit characters on a completion item.
		 */
		CommitCharactersSupport *bool `json:"commitCharactersSupport,omitempty"`

		/**
		 * Client supports the following content formats for the documentation
		 * property. The order describes the preferred format of the client.
		 */
		DocumentationFormat []MarkupKind `json:"documentationFormat,omitempty"`

		/**
		 * Client supports the deprecated property on a completion item.
		 */
		DeprecatedSupport *bool `json:"deprecatedSupport,omitempty"`

		/**
		 * Client supports the preselect property on a completion item.
		 */
		PreselectSupport *bool `json:"preselectSupport,omitempty"`

		/**
		 * Client supports the tag property on a completion item. Clients
		 * supporting tags have to handle unknown tags gracefully. Clients
		 * especially need to preserve unknown tags when sending a completion
		 * item back to the server in a resolve call.
		 *
		 * @since 3.15.0
		 */
		TagSupport *struct {
			/**
			 * The tags supported by the client.
			 */
			ValueSet []CompletionItemTag `json:"valueSet"`
		} `json:"tagSupport,omitempty"`

		/**
		 * Client supports insert replace edit to control different behavior if
		 * a completion item is inserted in the text or should replace text.
		 *
		 * @since 3.16.0
		 */
		InsertReplaceSupport *bool `json:"insertReplaceSupport,omitempty"`

		/**
		 * Indicates which properties a client can resolve lazily on a
		 * completion item. Before version 3.16.0 only the predefined properties
		 * `documentation` and `details` could be resolved lazily.
		 *
		 * @since 3.16.0
		 */
		ResolveSupport *struct {
			/**
			 * The properties that a client can resolve lazily.
			 */
			Properties []string `json:"properties"`
		} `json:"resolveSupport,omitempty"`

		/**
		 * The client supports the `insertTextMode` property on
		 * a completion item to override the whitespace handling mode
		 * as defined by the client (see `insertTextMode`).
		 *
		 * @since 3.16.0
		 */
		InsertTextModeSupport *struct {
			ValueSet []InsertTextMode `json:"valueSet"`
		} `json:"insertTextModeSupport,omitempty"`
	} `json:"completionItem,omitempty"`

	CompletionItemKind *struct {
		/**
		 * The completion item kind values the client supports. When this
		 * property exists the client also guarantees that it will
		 * handle values outside its set gracefully and falls back
		 * to a default value when unknown.
		 *
		 * If this property is not present the client only supports
		 * the completion items kinds from `Text` to `Reference` as defined in
		 * the initial version of the protocol.
		 */
		ValueSet []CompletionItemKind `json:"valueSet,omitempty"`
	} `json:"completionItemKind,omitempty"`

	/**
	 * The client supports to send additional context information for a
	 * `textDocument/completion` request.
	 */
	ContextSupport *bool `json:"contextSupport,omitempty"`
}

/**
 * Completion options.
 */
type CompletionOptions struct {
	WorkDoneProgressOptions

	/**
	 * Most tools trigger completion request automatically without explicitly
	 * requesting it using a keyboard shortcut (e.g. Ctrl+Space). Typically they
	 * do so when the user starts to type an identifier. For example if the user
	 * types `c` in a JavaScript file code complete will automatically pop up
	 * present `console` besides others as a completion item. Characters that
	 * make up identifiers don't need to be listed here.
	 *
	 * If code complete should automatically be trigger on characters not being
	 * valid inside an identifier (for example `.` in JavaScript) list them in
	 * `triggerCharacters`.
	 */
	TriggerCharacters []string `json:"triggerCharacters,omitempty"`

	/**
	 * The list of all possible characters that commit a completion. This field
	 * can be used if clients don't support individual commit characters per
	 * completion item. See client capability
	 * `completion.completionItem.commitCharactersSupport`.
	 *
	 * If a server provides both `allCommitCharacters` and commit characters on
	 * an individual completion item the ones on the completion item win.
	 *
	 * @since 3.2.0
	 */
	AllCommitCharacters []string `json:"allCommitCharacters,omitempty"`

	/**
	 * The server provides support to resolve additional
	 * information for a completion item.
	 */
	ResolveProvider *bool `json:"resolveProvider,omitempty"`
}

type CompletionRegistrationOptions struct {
	TextDocumentRegistrationOptions
	CompletionOptions
}

const MethodTextDocumentCompletion = Method("textDocument/completion")

// Returns: []CompletionItem | CompletionList | nil
type TextDocumentCompletionFunc func(context *glsp.Context, params *CompletionParams) (any, error)

type CompletionParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams

	/**
	 * The completion context. This is only available if the client specifies
	 * to send this using the client capability
	 * `completion.contextSupport === true`
	 */
	Context *CompletionContext `json:"context,omitempty"`
}

/**
 * How a completion was triggered
 */
type CompletionTriggerKind Integer

const (
	/**
	 * Completion was triggered by typing an identifier (24x7 code
	 * complete), manual invocation (e.g Ctrl+Space) or via API.
	 */
	CompletionTriggerKindInvoked = CompletionTriggerKind(1)

	/**
	 * Completion was triggered by a trigger character specified by
	 * the `triggerCharacters` properties of the
	 * `CompletionRegistrationOptions`.
	 */
	CompletionTriggerKindTriggerCharacter = CompletionTriggerKind(2)

	/**
	 * Completion was re-triggered as the current completion list is incomplete.
	 */
	CompletionTriggerKindTriggerForIncompleteCompletions = CompletionTriggerKind(3)
)

/**
 * Contains additional information about the context in which a completion
 * request is triggered.
 */
type CompletionContext struct {
	/**
	 * How the completion was triggered.
	 */
	TriggerKind CompletionTriggerKind `json:"triggerKind"`

	/**
	 * The trigger character (a single character) that has trigger code
	 * complete. Is undefined if
	 * `triggerKind !== CompletionTriggerKind.TriggerCharacter`
	 */
	TriggerCharacter *string `json:"triggerCharacter,omitempty"`
}

/**
 * Represents a collection of [completion items](#CompletionItem) to be
 * presented in the editor.
 */
type CompletionList struct {
	/**
	 * This list it not complete. Further typing should result in recomputing
	 * this list.
	 */
	IsIncomplete bool `json:"isIncomplete"`

	/**
	 * The completion items.
	 */
	Items []CompletionItem `json:"items"`
}

/**
 * Defines whether the insert text in a completion item should be interpreted as
 * plain text or a snippet.
 */
type InsertTextFormat Integer

const (
	/**
	 * The primary text to be inserted is treated as a plain string.
	 */
	InsertTextFormatPlainText = InsertTextFormat(1)

	/**
	 * The primary text to be inserted is treated as a snippet.
	 *
	 * A snippet can define tab stops and placeholders with `$1`, `$2`
	 * and `${3:foo}`. `$0` defines the final tab stop, it defaults to
	 * the end of the snippet. Placeholders with equal identifiers are linked,
	 * that is typing in one will update others too.
	 */
	InsertTextFormatSnippet = InsertTextFormat(2)
)

/**
 * Completion item tags are extra annotations that tweak the rendering of a
 * completion item.
 *
 * @since 3.15.0
 */
type CompletionItemTag Integer

const (
	/**
	 * Render a completion as obsolete, usually using a strike-out.
	 */
	CompletionItemTagDeprecated = CompletionItemTag(1)
)

/**
 * A special text edit to provide an insert and a replace operation.
 *
 * @since 3.16.0
 */
type InsertReplaceEdit struct {
	/**
	 * The string to be inserted.
	 */
	NewText string `json:"newText"`

	/**
	 * The range if the insert is requested
	 */
	Insert Range `json:"insert"`

	/**
	 * The range if the replace is requested.
	 */
	Replace Range `json:"replace"`
}

/**
 * How whitespace and indentation is handled during completion
 * item insertion.
 *
 * @since 3.16.0
 */
type InsertTextMode Integer

const (
	/**
	 * The insertion or replace strings is taken as it is. If the
	 * value is multi line the lines below the cursor will be
	 * inserted using the indentation defined in the string value.
	 * The client will not apply any kind of adjustments to the
	 * string.
	 */
	InsertTextModeAsIs = InsertTextMode(1)

	/**
	 * The editor adjusts leading whitespace of new lines so that
	 * they match the indentation up to the cursor of the line for
	 * which the item is accepted.
	 *
	 * Consider a line like this: <2tabs><cursor><3tabs>foo. Accepting a
	 * multi line completion item is indented using 2 tabs and all
	 * following lines inserted will be indented using 2 tabs as well.
	 */
	InsertTextModeAdjustIndentation = InsertTextMode(2)
)

type CompletionItem struct {
	/**
	 * The label of this completion item. By default
	 * also the text that is inserted when selecting
	 * this completion.
	 */
	Label string `json:"label"`

	/**
	 * The kind of this completion item. Based of the kind
	 * an icon is chosen by the editor. The standardized set
	 * of available values is defined in `CompletionItemKind`.
	 */
	Kind *CompletionItemKind `json:"kind,omitempty"`

	/**
	 * Tags for this completion item.
	 *
	 * @since 3.15.0
	 */
	Tags []CompletionItemTag `json:"tags,omitempty"`

	/**
	 * A human-readable string with additional information
	 * about this item, like type or symbol information.
	 */
	Detail *string `json:"detail,omitempty"`

	/**
	 * A human-readable string that represents a doc-comment.
	 */
	Documentation any `json:"documentation,omitempty"` // nil | string | MarkupContent

	/**
	 * Indicates if this item is deprecated.
	 *
	 * @deprecated Use `tags` instead if supported.
	 */
	Deprecated *bool `json:"deprecated,omitempty"`

	/**
	 * Select this item when showing.
	 *
	 * *Note* that only one completion item can be selected and that the
	 * tool / client decides which item that is. The rule is that the *first*
	 * item of those that match best is selected.
	 */
	Preselect *bool `json:"preselect,omitempty"`

	/**
	 * A string that should be used when comparing this item
	 * with other items. When `falsy` the label is used.
	 */
	SortText *string `json:"sortText,omitempty"`

	/**
	 * A string that should be used when filtering a set of
	 * completion items. When `falsy` the label is used.
	 */
	FilterText *string `json:"filterText,omitempty"`

	/**
	 * A string that should be inserted into a document when selecting
	 * this completion. When `falsy` the label is used.
	 *
	 * The `insertText` is subject to interpretation by the client side.
	 * Some tools might not take the string literally. For example
	 * VS Code when code complete is requested in this example
	 * `con<cursor position>` and a completion item with an `insertText` of
	 * `console` is provided it will only insert `sole`. Therefore it is
	 * recommended to use `textEdit` instead since it avoids additional client
	 * side interpretation.
	 */
	InsertText *string `json:"insertText,omitempty"`

	/**
	 * The format of the insert text. The format applies to both the
	 * `insertText` property and the `newText` property of a provided
	 * `textEdit`. If omitted defaults to `InsertTextFormat.PlainText`.
	 */
	InsertTextFormat *InsertTextFormat `json:"insertTextFormat,omitempty"`

	/**
	 * How whitespace and indentation is handled during completion
	 * item insertion. If not provided the client's default value depends on
	 * the `textDocument.completion.insertTextMode` client capability.
	 *
	 * @since 3.16.0
	 */
	InsertTextMode *InsertTextMode `json:"insertTextMode,omitempty"`

	/**
	 * An edit which is applied to a document when selecting this completion.
	 * When an edit is provided the value of `insertText` is ignored.
	 *
	 * *Note:* The range of the edit must be a single line range and it must
	 * contain the position at which completion has been requested.
	 *
	 * Most editors support two different operations when accepting a completion
	 * item. One is to insert a completion text and the other is to replace an
	 * existing text with a completion text. Since this can usually not be
	 * predetermined by a server it can report both ranges. Clients need to
	 * signal support for `InsertReplaceEdits` via the
	 * `textDocument.completion.insertReplaceSupport` client capability
	 * property.
	 *
	 * *Note 1:* The text edit's range as well as both ranges from an insert
	 * replace edit must be a [single line] and they must contain the position
	 * at which completion has been requested.
	 * *Note 2:* If an `InsertReplaceEdit` is returned the edit's insert range
	 * must be a prefix of the edit's replace range, that means it must be
	 * contained and starting at the same position.
	 *
	 * @since 3.16.0 additional type `InsertReplaceEdit`
	 */
	TextEdit any `json:"textEdit,omitempty"` // nil | TextEdit | InsertReplaceEdit

	/**
	 * An optional array of additional text edits that are applied when
	 * selecting this completion. Edits must not overlap (including the same
	 * insert position) with the main edit nor with themselves.
	 *
	 * Additional text edits should be used to change text unrelated to the
	 * current cursor position (for example adding an import statement at the
	 * top of the file if the completion item will insert an unqualified type).
	 */
	AdditionalTextEdits []TextEdit `json:"additionalTextEdits,omitempty"`

	/**
	 * An optional set of characters that when pressed while this completion is
	 * active will accept it first and then type that character. *Note* that all
	 * commit characters should have `length=1` and that superfluous characters
	 * will be ignored.
	 */
	CommitCharacters []string `json:"commitCharacters,omitempty"`

	/**
	 * An optional command that is executed *after* inserting this completion.
	 * *Note* that additional modifications to the current document should be
	 * described with the additionalTextEdits-property.
	 */
	Command *Command `json:"command,omitempty"`

	/**
	 * A data entry field that is preserved on a completion item between
	 * a completion and a completion resolve request.
	 */
	Data any `json:"data,omitempty"`
}

// json.Unmarshaler interface
func (self *CompletionItem) UnmarshalJSON(data []byte) error {
	var value struct {
		Label               string              `json:"label"`
		Kind                *CompletionItemKind `json:"kind,omitempty"`
		Tags                []CompletionItemTag `json:"tags,omitempty"`
		Detail              *string             `json:"detail,omitempty"`
		Documentation       json.RawMessage     `json:"documentation,omitempty"` // nil | string | MarkupContent
		Deprecated          *bool               `json:"deprecated,omitempty"`
		Preselect           *bool               `json:"preselect,omitempty"`
		SortText            *string             `json:"sortText,omitempty"`
		FilterText          *string             `json:"filterText,omitempty"`
		InsertText          *string             `json:"insertText,omitempty"`
		InsertTextFormat    *InsertTextFormat   `json:"insertTextFormat,omitempty"`
		InsertTextMode      *InsertTextMode     `json:"insertTextMode,omitempty"`
		TextEdit            json.RawMessage     `json:"textEdit,omitempty"` // nil | TextEdit | InsertReplaceEdit
		AdditionalTextEdits []TextEdit          `json:"additionalTextEdits,omitempty"`
		CommitCharacters    []string            `json:"commitCharacters,omitempty"`
		Command             *Command            `json:"command,omitempty"`
		Data                any                 `json:"data,omitempty"`
	}

	if err := json.Unmarshal(data, &value); err == nil {
		self.Label = value.Label
		self.Kind = value.Kind
		self.Tags = value.Tags
		self.Detail = value.Detail
		self.Deprecated = value.Deprecated
		self.Preselect = value.Preselect
		self.SortText = value.SortText
		self.FilterText = value.FilterText
		self.InsertText = value.InsertText
		self.InsertTextFormat = value.InsertTextFormat
		self.InsertTextMode = value.InsertTextMode
		self.AdditionalTextEdits = value.AdditionalTextEdits
		self.CommitCharacters = value.CommitCharacters
		self.Command = value.Command
		self.Data = value.Data

		if value.Documentation != nil {
			var value_ string
			if err = json.Unmarshal(value.Documentation, &value_); err == nil {
				self.Documentation = value_
			} else {
				var value_ MarkupContent
				if err = json.Unmarshal(value.Documentation, &value_); err == nil {
					self.Documentation = value_
				} else {
					return err
				}
			}
		}

		if value.TextEdit != nil {
			var value_ TextEdit
			if err = json.Unmarshal(value.TextEdit, &value_); err == nil {
				self.TextEdit = value_
			} else {
				var value_ InsertReplaceEdit
				if err = json.Unmarshal(value.TextEdit, &value_); err == nil {
					self.TextEdit = value_
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

/**
 * The kind of a completion entry.
 */
type CompletionItemKind Integer

const (
	CompletionItemKindText          = CompletionItemKind(1)
	CompletionItemKindMethod        = CompletionItemKind(2)
	CompletionItemKindFunction      = CompletionItemKind(3)
	CompletionItemKindConstructor   = CompletionItemKind(4)
	CompletionItemKindField         = CompletionItemKind(5)
	CompletionItemKindVariable      = CompletionItemKind(6)
	CompletionItemKindClass         = CompletionItemKind(7)
	CompletionItemKindInterface     = CompletionItemKind(8)
	CompletionItemKindModule        = CompletionItemKind(9)
	CompletionItemKindProperty      = CompletionItemKind(10)
	CompletionItemKindUnit          = CompletionItemKind(11)
	CompletionItemKindValue         = CompletionItemKind(12)
	CompletionItemKindEnum          = CompletionItemKind(13)
	CompletionItemKindKeyword       = CompletionItemKind(14)
	CompletionItemKindSnippet       = CompletionItemKind(15)
	CompletionItemKindColor         = CompletionItemKind(16)
	CompletionItemKindFile          = CompletionItemKind(17)
	CompletionItemKindReference     = CompletionItemKind(18)
	CompletionItemKindFolder        = CompletionItemKind(19)
	CompletionItemKindEnumMember    = CompletionItemKind(20)
	CompletionItemKindConstant      = CompletionItemKind(21)
	CompletionItemKindStruct        = CompletionItemKind(22)
	CompletionItemKindEvent         = CompletionItemKind(23)
	CompletionItemKindOperator      = CompletionItemKind(24)
	CompletionItemKindTypeParameter = CompletionItemKind(25)
)

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#completionItem_resolve

const MethodCompletionItemResolve = Method("completionItem/resolve")

type CompletionItemResolveFunc func(context *glsp.Context, params *CompletionItem) (*CompletionItem, error)

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#textDocument_hover

type HoverClientCapabilities struct {
	/**
	 * Whether hover supports dynamic registration.
	 */
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`

	/**
	 * Client supports the following content formats if the content
	 * property refers to a `literal of type MarkupContent`.
	 * The order describes the preferred format of the client.
	 */
	ContentFormat []MarkupKind `json:"contentFormat,omitempty"`
}

type HoverOptions struct {
	WorkDoneProgressOptions
}

type HoverRegistrationOptions struct {
	TextDocumentRegistrationOptions
	HoverOptions
}

const MethodTextDocumentHover = Method("textDocument/hover")

type TextDocumentHoverFunc func(context *glsp.Context, params *HoverParams) (*Hover, error)

type HoverParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
}

/**
 * The result of a hover request.
 */
type Hover struct {
	/**
	 * The hover's content
	 */
	Contents any `json:"contents"` // MarkupContent | MarkedString | []MarkedString

	/**
	 * An optional range is a range inside a text document
	 * that is used to visualize a hover, e.g. by changing the background color.
	 */
	Range *Range `json:"range,omitempty"`
}

// json.Unmarshaler interface
func (self *Hover) UnmarshalJSON(data []byte) error {
	var value struct {
		Contents json.RawMessage `json:"contents"` // MarkupContent | MarkedString | []MarkedString
		Range    *Range          `json:"range,omitempty"`
	}

	if err := json.Unmarshal(data, &value); err == nil {
		self.Range = value.Range

		var value_ MarkupContent
		if err = json.Unmarshal(value.Contents, &value_); err == nil {
			self.Contents = value_
		} else {
			var value_ MarkedString
			if err = json.Unmarshal(value.Contents, &value_); err == nil {
				self.Contents = value_
			} else {
				var value_ []MarkedString
				if err = json.Unmarshal(value.Contents, &value_); err == nil {
					self.Contents = value_
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

/**
 * MarkedString can be used to render human readable text. It is either a
 * markdown string or a code-block that provides a language and a code snippet.
 * The language identifier is semantically equal to the optional language
 * identifier in fenced code blocks in GitHub issues.
 *
 * The pair of a language and a value is an equivalent to markdown:
 * ```${language}
 * ${value}
 * ```
 *
 * Note that markdown strings will be sanitized - that means html will be
 * escaped.
 *
 * @deprecated use MarkupContent instead.
 */
type MarkedString struct {
	value any // string | MarkedStringStruct
}

type MarkedStringStruct struct {
	Language string `json:"language"`
	Value    string `json:"value"`
}

// ([json.Marshaler] interface)
func (self MarkedString) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.value)
}

// json.Unmarshaler interface
func (self MarkedString) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err == nil {
		self.value = value
		return nil
	} else {
		var value MarkedStringStruct
		if err := json.Unmarshal(data, &value); err == nil {
			self.value = value
			return nil
		} else {
			return err
		}
	}
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#textDocument_signatureHelp

type SignatureHelpClientCapabilities struct {
	/**
	 * Whether signature help supports dynamic registration.
	 */
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`

	/**
	 * The client supports the following `SignatureInformation`
	 * specific properties.
	 */
	SignatureInformation *struct {
		/**
		 * Client supports the following content formats for the documentation
		 * property. The order describes the preferred format of the client.
		 */
		DocumentationFormat []MarkupKind `json:"documentationFormat,omitempty"`

		/**
		 * Client capabilities specific to parameter information.
		 */
		ParameterInformation *struct {
			/**
			 * The client supports processing label offsets instead of a
			 * simple label string.
			 *
			 * @since 3.14.0
			 */
			LabelOffsetSupport *bool `json:"labelOffsetSupport,omitempty"`
		} `json:"parameterInformation,omitempty"`

		/**
		 * The client supports the `activeParameter` property on
		 * `SignatureInformation` literal.
		 *
		 * @since 3.16.0
		 */
		ActiveParameterSupport *bool `json:"activeParameterSupport,omitempty"`
	} `json:"signatureInformation,omitempty"`

	/**
	 * The client supports to send additional context information for a
	 * `textDocument/signatureHelp` request. A client that opts into
	 * contextSupport will also support the `retriggerCharacters` on
	 * `SignatureHelpOptions`.
	 *
	 * @since 3.15.0
	 */
	ContextSupport *bool `json:"contextSupport,omitempty"`
}

type SignatureHelpOptions struct {
	WorkDoneProgressOptions

	/**
	 * The characters that trigger signature help
	 * automatically.
	 */
	TriggerCharacters []string `json:"triggerCharacters,omitempty"`

	/**
	 * List of characters that re-trigger signature help.
	 *
	 * These trigger characters are only active when signature help is already
	 * showing. All trigger characters are also counted as re-trigger
	 * characters.
	 *
	 * @since 3.15.0
	 */
	RetriggerCharacters []string `json:"retriggerCharacters,omitempty"`
}

type SignatureHelpRegistrationOptions struct {
	TextDocumentRegistrationOptions
	SignatureHelpOptions
}

const MethodTextDocumentSignatureHelp = Method("textDocument/signatureHelp")

type TextDocumentSignatureHelpFunc func(context *glsp.Context, params *SignatureHelpParams) (*SignatureHelp, error)

type SignatureHelpParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams

	/**
	 * The signature help context. This is only available if the client
	 * specifies to send this using the client capability
	 * `textDocument.signatureHelp.contextSupport === true`
	 *
	 * @since 3.15.0
	 */
	Context *SignatureHelpContext `json:"context,omitempty"`
}

/**
 * How a signature help was triggered.
 *
 * @since 3.15.0
 */
type SignatureHelpTriggerKind Integer

const (
	/**
	 * Signature help was invoked manually by the user or by a command.
	 */
	SignatureHelpTriggerKindInvoked = SignatureHelpTriggerKind(1)

	/**
	 * Signature help was triggered by a trigger character.
	 */
	SignatureHelpTriggerKindTriggerCharacter = SignatureHelpTriggerKind(2)

	/**
	 * Signature help was triggered by the cursor moving or by the document
	 * content changing.
	 */
	SignatureHelpTriggerKindContentChange = SignatureHelpTriggerKind(3)
)

/**
 * Additional information about the context in which a signature help request
 * was triggered.
 *
 * @since 3.15.0
 */
type SignatureHelpContext struct {
	/**
	 * Action that caused signature help to be triggered.
	 */
	TriggerKind SignatureHelpTriggerKind `json:"triggerKind"`

	/**
	 * Character that caused signature help to be triggered.
	 *
	 * This is undefined when triggerKind !==
	 * SignatureHelpTriggerKind.TriggerCharacter
	 */
	TriggerCharacter *string `json:"triggerCharacter,omitempty"`

	/**
	 * `true` if signature help was already showing when it was triggered.
	 *
	 * Retriggers occur when the signature help is already active and can be
	 * caused by actions such as typing a trigger character, a cursor move, or
	 * document content changes.
	 */
	IsRetrigger bool `json:"isRetrigger"`

	/**
	 * The currently active `SignatureHelp`.
	 *
	 * The `activeSignatureHelp` has its `SignatureHelp.activeSignature` field
	 * updated based on the user navigating through available signatures.
	 */
	ActiveSignatureHelp *SignatureHelp `json:"activeSignatureHelp,omitempty"`
}

/**
 * Signature help represents the signature of something
 * callable. There can be multiple signature but only one
 * active and only one active parameter.
 */
type SignatureHelp struct {
	/**
	 * One or more signatures. If no signatures are available the signature help
	 * request should return `null`.
	 */
	Signatures []SignatureInformation `json:"signatures"`

	/**
	 * The active signature. If omitted or the value lies outside the
	 * range of `signatures` the value defaults to zero or is ignored if
	 * the `SignatureHelp` has no signatures.
	 *
	 * Whenever possible implementors should make an active decision about
	 * the active signature and shouldn't rely on a default value.
	 *
	 * In future version of the protocol this property might become
	 * mandatory to better express this.
	 */
	ActiveSignature *UInteger `json:"activeSignature,omitempty"`

	/**
	 * The active parameter of the active signature. If omitted or the value
	 * lies outside the range of `signatures[activeSignature].parameters`
	 * defaults to 0 if the active signature has parameters. If
	 * the active signature has no parameters it is ignored.
	 * In future version of the protocol this property might become
	 * mandatory to better express the active parameter if the
	 * active signature does have any.
	 */
	ActiveParameter *UInteger `json:"activeParameter,omitempty"`
}

/**
 * Represents the signature of something callable. A signature
 * can have a label, like a function-name, a doc-comment, and
 * a set of parameters.
 */
type SignatureInformation struct {
	/**
	 * The label of this signature. Will be shown in
	 * the UI.
	 */
	Label string `json:"label"`

	/**
	 * The human-readable doc-comment of this signature. Will be shown
	 * in the UI but can be omitted.
	 */
	Documentation any `json:"documentation,omitempty"` // nil | string | MarkupContent

	/**
	 * The parameters of this signature.
	 */
	Parameters []ParameterInformation `json:"parameters,omitempty"`

	/**
	 * The index of the active parameter.
	 *
	 * If provided, this is used in place of `SignatureHelp.activeParameter`.
	 *
	 * @since 3.16.0
	 */
	ActiveParameter *UInteger `json:"activeParameter,omitempty"`
}

// json.Unmarshaler interface
func (self *SignatureInformation) UnmarshalJSON(data []byte) error {
	var value struct {
		Label           string                 `json:"label"`
		Documentation   json.RawMessage        `json:"documentation"` // nil | string | MarkupContent
		Parameters      []ParameterInformation `json:"parameters"`
		ActiveParameter *UInteger              `json:"activeParameter"`
	}

	if err := json.Unmarshal(data, &value); err == nil {
		self.Label = value.Label
		self.Parameters = value.Parameters
		self.ActiveParameter = value.ActiveParameter

		if value.Documentation != nil {
			var value_ string
			if err = json.Unmarshal(value.Documentation, &value_); err == nil {
				self.Documentation = value_
			} else {
				var value_ MarkupContent
				if err = json.Unmarshal(value.Documentation, &value_); err == nil {
					self.Documentation = value_
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

/**
 * Represents a parameter of a callable-signature. A parameter can
 * have a label and a doc-comment.
 */
type ParameterInformation struct {
	/**
	 * The label of this parameter information.
	 *
	 * Either a string or an inclusive start and exclusive end offsets within
	 * its containing signature label. (see SignatureInformation.label). The
	 * offsets are based on a UTF-16 string representation as `Position` and
	 * `Range` does.
	 *
	 * *Note*: a label of type string should be a substring of its containing
	 * signature label. Its intended use case is to highlight the parameter
	 * label part in the `SignatureInformation.label`.
	 */
	Label any `json:"label"` // string | [2]UInteger

	/**
	 * The human-readable doc-comment of this parameter. Will be shown
	 * in the UI but can be omitted.
	 */
	Documentation any `json:"documentation,omitempty"` // nil | string | MarkupContent
}

// json.Unmarshaler interface
func (self *ParameterInformation) UnmarshalJSON(data []byte) error {
	var value struct {
		Label         json.RawMessage `json:"label"`         // string | [2]UInteger
		Documentation json.RawMessage `json:"documentation"` // nil | string | MarkupContent
	}

	if err := json.Unmarshal(data, &value); err == nil {
		var value_ string
		if err = json.Unmarshal(value.Label, &value_); err == nil {
			self.Label = value_
		} else {
			var value_ []UInteger
			if err = json.Unmarshal(value.Label, &value_); err == nil {
				self.Label = value_
			} else {
				return err
			}
		}

		if value.Documentation != nil {
			var value_ string
			if err = json.Unmarshal(value.Documentation, &value_); err == nil {
				self.Documentation = value_
			} else {
				var value_ MarkupContent
				if err = json.Unmarshal(value.Documentation, &value_); err == nil {
					self.Documentation = value_
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

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#textDocument_declaration

type DeclarationClientCapabilities struct {
	/**
	 * Whether declaration supports dynamic registration. If this is set to
	 * `true` the client supports the new `DeclarationRegistrationOptions`
	 * return value for the corresponding server capability as well.
	 */
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`

	/**
	 * The client supports additional metadata in the form of declaration links.
	 */
	LinkSupport *bool `json:"linkSupport,omitempty"`
}

type DeclarationOptions struct {
	WorkDoneProgressOptions
}

type DeclarationRegistrationOptions struct {
	DeclarationOptions
	TextDocumentRegistrationOptions
	StaticRegistrationOptions
}

const MethodTextDocumentDeclaration = Method("textDocument/declaration")

// Returns: Location | []Location | []LocationLink | nil
type TextDocumentDeclarationFunc func(context *glsp.Context, params *DeclarationParams) (any, error)

type DeclarationParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#textDocument_definition

type DefinitionClientCapabilities struct {
	/**
	 * Whether definition supports dynamic registration.
	 */
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`

	/**
	 * The client supports additional metadata in the form of definition links.
	 *
	 * @since 3.14.0
	 */
	LinkSupport *bool `json:"linkSupport,omitempty"`
}

type DefinitionOptions struct {
	WorkDoneProgressOptions
}

type DefinitionRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DefinitionOptions
}

const MethodTextDocumentDefinition = Method("textDocument/definition")

// Returns: Location | []Location | []LocationLink | nil
type TextDocumentDefinitionFunc func(context *glsp.Context, params *DefinitionParams) (any, error)

type DefinitionParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#textDocument_typeDefinition

type TypeDefinitionClientCapabilities struct {
	/**
	 * Whether implementation supports dynamic registration. If this is set to
	 * `true` the client supports the new `TypeDefinitionRegistrationOptions`
	 * return value for the corresponding server capability as well.
	 */
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`

	/**
	 * The client supports additional metadata in the form of definition links.
	 *
	 * @since 3.14.0
	 */
	LinkSupport *bool `json:"linkSupport,omitempty"`
}

type TypeDefinitionOptions struct {
	WorkDoneProgressOptions
}

type TypeDefinitionRegistrationOptions struct {
	TextDocumentRegistrationOptions
	TypeDefinitionOptions
	StaticRegistrationOptions
}

const MethodTextDocumentTypeDefinition = Method("textDocument/typeDefinition")

// Returns: Location | []Location | []LocationLink | nil
type TextDocumentTypeDefinitionFunc func(context *glsp.Context, params *TypeDefinitionParams) (any, error)

type TypeDefinitionParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#textDocument_implementation

type ImplementationClientCapabilities struct {
	/**
	 * Whether implementation supports dynamic registration. If this is set to
	 * `true` the client supports the new `ImplementationRegistrationOptions`
	 * return value for the corresponding server capability as well.
	 */
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`

	/**
	 * The client supports additional metadata in the form of definition links.
	 *
	 * @since 3.14.0
	 */
	LinkSupport *bool `json:"linkSupport,omitempty"`
}

type ImplementationOptions struct {
	WorkDoneProgressOptions
}

type ImplementationRegistrationOptions struct {
	TextDocumentRegistrationOptions
	TypeDefinitionOptions
	StaticRegistrationOptions
}

const MethodTextDocumentImplementation = Method("textDocument/implementation")

// Returns: Location | []Location | []LocationLink | nil
type TextDocumentImplementationFunc func(context *glsp.Context, params *ImplementationParams) (any, error)

type ImplementationParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#textDocument_references

type ReferenceClientCapabilities struct {
	/**
	 * Whether references supports dynamic registration.
	 */
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

type ReferenceOptions struct {
	WorkDoneProgressOptions
}

type ReferenceRegistrationOptions struct {
	TextDocumentRegistrationOptions
	ReferenceOptions
}

const MethodTextDocumentReferences = Method("textDocument/references")

type TextDocumentReferencesFunc func(context *glsp.Context, params *ReferenceParams) ([]Location, error)

type ReferenceParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams

	Context ReferenceContext `json:"context"`
}

type ReferenceContext struct {
	/**
	 * Include the declaration of the current symbol.
	 */
	IncludeDeclaration bool `json:"includeDeclaration"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#textDocument_documentHighlight

type DocumentHighlightClientCapabilities struct {
	/**
	 * Whether document highlight supports dynamic registration.
	 */
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

type DocumentHighlightOptions struct {
	WorkDoneProgressOptions
}

type DocumentHighlightRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DocumentHighlightOptions
}

const MethodTextDocumentDocumentHighlight = Method("textDocument/documentHighlight")

type TextDocumentDocumentHighlightFunc func(context *glsp.Context, params *DocumentHighlightParams) ([]DocumentHighlight, error)

type DocumentHighlightParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
}

/**
 * A document highlight is a range inside a text document which deserves
 * special attention. Usually a document highlight is visualized by changing
 * the background color of its range.
 *
 */
type DocumentHighlight struct {
	/**
	 * The range this highlight applies to.
	 */
	Range Range `json:"range"`

	/**
	 * The highlight kind, default is DocumentHighlightKind.Text.
	 */
	Kind *DocumentHighlightKind `json:"kind,omitempty"`
}

/**
 * A document highlight kind.
 */
type DocumentHighlightKind Integer

const (
	/**
	 * A textual occurrence.
	 */
	DocumentHighlightKindText = DocumentHighlightKind(1)

	/**
	 * Read-access of a symbol, like reading a variable.
	 */
	DocumentHighlightKindRead = DocumentHighlightKind(2)

	/**
	 * Write-access of a symbol, like writing to a variable.
	 */
	DocumentHighlightKindWrite = DocumentHighlightKind(3)
)

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#textDocument_documentSymbol

type DocumentSymbolClientCapabilities struct {
	/**
	 * Whether document symbol supports dynamic registration.
	 */
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`

	/**
	 * Specific capabilities for the `SymbolKind` in the
	 * `textDocument/documentSymbol` request.
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
	 * The client supports hierarchical document symbols.
	 */
	HierarchicalDocumentSymbolSupport *bool `json:"hierarchicalDocumentSymbolSupport,omitempty"`

	/**
	 * The client supports tags on `SymbolInformation`. Tags are supported on
	 * `DocumentSymbol` if `hierarchicalDocumentSymbolSupport` is set to true.
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

	/**
	 * The client supports an additional label presented in the UI when
	 * registering a document symbol provider.
	 *
	 * @since 3.16.0
	 */
	LabelSupport *bool `json:"labelSupport,omitempty"`
}

type DocumentSymbolOptions struct {
	WorkDoneProgressOptions

	/**
	 * A human-readable string that is shown when multiple outlines trees
	 * are shown for the same document.
	 *
	 * @since 3.16.0
	 */
	Label *string `json:"label,omitempty"`
}

type DocumentSymbolRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DocumentSymbolOptions
}

const MethodTextDocumentDocumentSymbol = Method("textDocument/documentSymbol")

// Returns: []DocumentSymbol | []SymbolInformation | nil
type TextDocumentDocumentSymbolFunc func(context *glsp.Context, params *DocumentSymbolParams) (any, error)

type DocumentSymbolParams struct {
	WorkDoneProgressParams
	PartialResultParams

	/**
	 * The text document.
	 */
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

/**
 * A symbol kind.
 */
type SymbolKind Integer

const (
	SymbolKindFile          = SymbolKind(1)
	SymbolKindModule        = SymbolKind(2)
	SymbolKindNamespace     = SymbolKind(3)
	SymbolKindPackage       = SymbolKind(4)
	SymbolKindClass         = SymbolKind(5)
	SymbolKindMethod        = SymbolKind(6)
	SymbolKindProperty      = SymbolKind(7)
	SymbolKindField         = SymbolKind(8)
	SymbolKindConstructor   = SymbolKind(9)
	SymbolKindEnum          = SymbolKind(10)
	SymbolKindInterface     = SymbolKind(11)
	SymbolKindFunction      = SymbolKind(12)
	SymbolKindVariable      = SymbolKind(13)
	SymbolKindConstant      = SymbolKind(14)
	SymbolKindString        = SymbolKind(15)
	SymbolKindNumber        = SymbolKind(16)
	SymbolKindBoolean       = SymbolKind(17)
	SymbolKindArray         = SymbolKind(18)
	SymbolKindObject        = SymbolKind(19)
	SymbolKindKey           = SymbolKind(20)
	SymbolKindNull          = SymbolKind(21)
	SymbolKindEnumMember    = SymbolKind(22)
	SymbolKindStruct        = SymbolKind(23)
	SymbolKindEvent         = SymbolKind(24)
	SymbolKindOperator      = SymbolKind(25)
	SymbolKindTypeParameter = SymbolKind(26)
)

/**
 * Symbol tags are extra annotations that tweak the rendering of a symbol.
 *
 * @since 3.16.0
 */
type SymbolTag Integer

const (
	/**
	 * Render a symbol as obsolete, usually using a strike-out.
	 */
	SymbolTagDeprecated = SymbolTag(1)
)

/**
 * Represents programming constructs like variables, classes, interfaces etc.
 * that appear in a document. Document symbols can be hierarchical and they
 * have two ranges: one that encloses its definition and one that points to its
 * most interesting range, e.g. the range of an identifier.
 */
type DocumentSymbol struct {
	/**
	 * The name of this symbol. Will be displayed in the user interface and
	 * therefore must not be an empty string or a string only consisting of
	 * white spaces.
	 */
	Name string `json:"name"`

	/**
	 * More detail for this symbol, e.g the signature of a function.
	 */
	Detail *string `json:"detail,omitempty"`

	/**
	 * The kind of this symbol.
	 */
	Kind SymbolKind `json:"kind"`

	/**
	 * Tags for this document symbol.
	 *
	 * @since 3.16.0
	 */
	Tags []SymbolTag `json:"tags,omitempty"`

	/**
	 * Indicates if this symbol is deprecated.
	 *
	 * @deprecated Use tags instead
	 */
	Deprecated *bool `json:"deprecated,omitempty"`

	/**
	 * The range enclosing this symbol not including leading/trailing whitespace
	 * but everything else like comments. This information is typically used to
	 * determine if the clients cursor is inside the symbol to reveal in the
	 * symbol in the UI.
	 */
	Range Range `json:"range"`

	/**
	 * The range that should be selected and revealed when this symbol is being
	 * picked, e.g. the name of a function. Must be contained by the `range`.
	 */
	SelectionRange Range `json:"selectionRange"`

	/**
	 * Children of this symbol, e.g. properties of a class.
	 */
	Children []DocumentSymbol `json:"children,omitempty"`
}

/**
 * Represents information about programming constructs like variables, classes,
 * interfaces etc.
 */
type SymbolInformation struct {
	/**
	 * The name of this symbol.
	 */
	Name string `json:"name"`

	/**
	 * The kind of this symbol.
	 */
	Kind SymbolKind `json:"kind"`

	/**
	 * Tags for this completion item.
	 *
	 * @since 3.16.0
	 */
	Tags []SymbolTag `json:"tags,omitempty"`

	/**
	 * Indicates if this symbol is deprecated.
	 *
	 * @deprecated Use tags instead
	 */
	Deprecated *bool `json:"deprecated,omitempty"`

	/**
	 * The location of this symbol. The location's range is used by a tool
	 * to reveal the location in the editor. If the symbol is selected in the
	 * tool the range's start information is used to position the cursor. So
	 * the range usually spans more then the actual symbol's name and does
	 * normally include things like visibility modifiers.
	 *
	 * The range doesn't have to denote a node range in the sense of a abstract
	 * syntax tree. It can therefore not be used to re-construct a hierarchy of
	 * the symbols.
	 */
	Location Location `json:"location"`

	/**
	 * The name of the symbol containing this symbol. This information is for
	 * user interface purposes (e.g. to render a qualifier in the user interface
	 * if necessary). It can't be used to re-infer a hierarchy for the document
	 * symbols.
	 */
	ContainerName *string `json:"containerName,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#textDocument_codeAction

type CodeActionClientCapabilities struct {
	/**
	 * Whether code action supports dynamic registration.
	 */
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`

	/**
	 * The client supports code action literals as a valid
	 * response of the `textDocument/codeAction` request.
	 *
	 * @since 3.8.0
	 */
	CodeActionLiteralSupport *struct {
		/**
		 * The code action kind is supported with the following value
		 * set.
		 */
		CodeActionKind struct {
			/**
			 * The code action kind values the client supports. When this
			 * property exists the client also guarantees that it will
			 * handle values outside its set gracefully and falls back
			 * to a default value when unknown.
			 */
			ValueSet []CodeActionKind `json:"valueSet"`
		} `json:"codeActionKind"`
	} `json:"codeActionLiteralSupport,omitempty"`

	/**
	 * Whether code action supports the `isPreferred` property.
	 *
	 * @since 3.15.0
	 */
	IsPreferredSupport *bool `json:"isPreferredSupport,omitempty"`

	/**
	 * Whether code action supports the `disabled` property.
	 *
	 * @since 3.16.0
	 */
	DisabledSupport *bool `json:"disabledSupport,omitempty"`

	/**
	 * Whether code action supports the `data` property which is
	 * preserved between a `textDocument/codeAction` and a
	 * `codeAction/resolve` request.
	 *
	 * @since 3.16.0
	 */
	DataSupport *bool `json:"dataSupport,omitempty"`

	/**
	 * Whether the client supports resolving additional code action
	 * properties via a separate `codeAction/resolve` request.
	 *
	 * @since 3.16.0
	 */
	ResolveSupport *struct {
		/**
		 * The properties that a client can resolve lazily.
		 */
		Properties []string `json:"properties"`
	} `json:"resolveSupport,omitempty"`

	/**
	 * Whether the client honors the change annotations in
	 * text edits and resource operations returned via the
	 * `CodeAction#edit` property by for example presenting
	 * the workspace edit in the user interface and asking
	 * for confirmation.
	 *
	 * @since 3.16.0
	 */
	HonorsChangeAnnotations *bool `json:"honorsChangeAnnotations,omitempty"`
}

type CodeActionOptions struct {
	WorkDoneProgressOptions

	/**
	 * CodeActionKinds that this server may return.
	 *
	 * The list of kinds may be generic, such as `CodeActionKind.Refactor`,
	 * or the server may list out every specific kind they provide.
	 */
	CodeActionKinds []CodeActionKind `json:"codeActionKinds,omitempty"`

	/**
	 * The server provides support to resolve additional
	 * information for a code action.
	 *
	 * @since 3.16.0
	 */
	ResolveProvider *bool `json:"resolveProvider,omitempty"`
}

type CodeActionRegistrationOptions struct {
	TextDocumentRegistrationOptions
	CodeActionOptions
}

const MethodTextDocumentCodeAction = Method("textDocument/codeAction")

// Returns: Command | []CodeAction | nil
type TextDocumentCodeActionFunc func(context *glsp.Context, params *CodeActionParams) (any, error)

/**
 * Params for the CodeActionRequest
 */
type CodeActionParams struct {
	WorkDoneProgressParams
	PartialResultParams

	/**
	 * The document in which the command was invoked.
	 */
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	/**
	 * The range for which the command was invoked.
	 */
	Range Range `json:"range"`

	/**
	 * Context carrying additional information.
	 */
	Context CodeActionContext `json:"context"`
}

/**
 * The kind of a code action.
 *
 * Kinds are a hierarchical list of identifiers separated by `.`,
 * e.g. `"refactor.extract.function"`.
 *
 * The set of kinds is open and client needs to announce the kinds it supports
 * to the server during initialization.
 */
type CodeActionKind = string

/**
 * A set of predefined code action kinds.
 */
const (
	/**
	 * Empty kind.
	 */
	CodeActionKindEmpty = CodeActionKind("")

	/**
	 * Base kind for quickfix actions: 'quickfix'.
	 */
	CodeActionKindQuickFix = CodeActionKind("quickfix")

	/**
	 * Base kind for refactoring actions: 'refactor'.
	 */
	CodeActionKindRefactor = CodeActionKind("refactor")

	/**
	 * Base kind for refactoring extraction actions: 'refactor.extract'.
	 *
	 * Example extract actions:
	 *
	 * - Extract method
	 * - Extract function
	 * - Extract variable
	 * - Extract interface from class
	 * - ...
	 */
	CodeActionKindRefactorExtract = CodeActionKind("refactor.extract")

	/**
	 * Base kind for refactoring inline actions: 'refactor.inline'.
	 *
	 * Example inline actions:
	 *
	 * - Inline function
	 * - Inline variable
	 * - Inline constant
	 * - ...
	 */
	CodeActionKindRefactorInline = CodeActionKind("refactor.inline")

	/**
	 * Base kind for refactoring rewrite actions: 'refactor.rewrite'.
	 *
	 * Example rewrite actions:
	 *
	 * - Convert JavaScript function to class
	 * - Add or remove parameter
	 * - Encapsulate field
	 * - Make method static
	 * - Move method to base class
	 * - ...
	 */
	CodeActionKindRefactorRewrite = CodeActionKind("refactor.rewrite")

	/**
	 * Base kind for source actions: `source`.
	 *
	 * Source code actions apply to the entire file.
	 */
	CodeActionKindSource = CodeActionKind("source")

	/**
	 * Base kind for an organize imports source action:
	 * `source.organizeImports`.
	 */
	CodeActionKindSourceOrganizeImports = CodeActionKind("source.organizeImports")
)

/**
* Contains additional diagnostic information about the context in which
* a code action is run.
 */
type CodeActionContext struct {
	/**
	 * An array of diagnostics known on the client side overlapping the range
	 * provided to the `textDocument/codeAction` request. They are provided so
	 * that the server knows which errors are currently presented to the user
	 * for the given range. There is no guarantee that these accurately reflect
	 * the error state of the resource. The primary parameter
	 * to compute code actions is the provided range.
	 */
	Diagnostics []Diagnostic `json:"diagnostics"`

	/**
	 * Requested kind of actions to return.
	 *
	 * Actions not of this kind are filtered out by the client before being
	 * shown. So servers can omit computing them.
	 */
	Only []CodeActionKind `json:"only,omitempty"`
}

/**
 * A code action represents a change that can be performed in code, e.g. to fix
 * a problem or to refactor code.
 *
 * A CodeAction must set either `edit` and/or a `command`. If both are supplied,
 * the `edit` is applied first, then the `command` is executed.
 */
type CodeAction struct {
	/**
	 * A short, human-readable, title for this code action.
	 */
	Title string `json:"title"`

	/**
	 * The kind of the code action.
	 *
	 * Used to filter code actions.
	 */
	Kind *CodeActionKind `json:"kind,omitempty"`

	/**
	 * The diagnostics that this code action resolves.
	 */
	Diagnostics []Diagnostic `json:"diagnostics,omitempty"`

	/**
	 * Marks this as a preferred action. Preferred actions are used by the
	 * `auto fix` command and can be targeted by keybindings.
	 *
	 * A quick fix should be marked preferred if it properly addresses the
	 * underlying error. A refactoring should be marked preferred if it is the
	 * most reasonable choice of actions to take.
	 *
	 * @since 3.15.0
	 */
	IsPreferred *bool `json:"isPreferred,omitempty"`

	/**
	 * Marks that the code action cannot currently be applied.
	 *
	 * Clients should follow the following guidelines regarding disabled code
	 * actions:
	 *
	 * - Disabled code actions are not shown in automatic lightbulbs code
	 *   action menus.
	 *
	 * - Disabled actions are shown as faded out in the code action menu when
	 *   the user request a more specific type of code action, such as
	 *   refactorings.
	 *
	 * - If the user has a keybinding that auto applies a code action and only
	 *   a disabled code actions are returned, the client should show the user
	 *   an error message with `reason` in the editor.
	 *
	 * @since 3.16.0
	 */
	Disabled *struct {
		/**
		 * Human readable description of why the code action is currently
		 * disabled.
		 *
		 * This is displayed in the code actions UI.
		 */
		Reason string `json:"reason"`
	} `json:"disabled,omitempty"`

	/**
	 * The workspace edit this code action performs.
	 */
	Edit *WorkspaceEdit `json:"edit,omitempty"`

	/**
	 * A command this code action executes. If a code action
	 * provides an edit and a command, first the edit is
	 * executed and then the command.
	 */
	Command *Command `json:"command,omitempty"`

	/**
	 * A data entry field that is preserved on a code action between
	 * a `textDocument/codeAction` and a `codeAction/resolve` request.
	 *
	 * @since 3.16.0
	 */
	Data any `json:"data,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#codeAction_resolve

const MethodCodeActionResolve = Method("codeAction/resolve")

type CodeActionResolveFunc func(context *glsp.Context, params *CodeAction) (*CodeAction, error)

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#textDocument_codeLens

type CodeLensClientCapabilities struct {
	/**
	 * Whether code lens supports dynamic registration.
	 */
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

type CodeLensOptions struct {
	WorkDoneProgressOptions

	/**
	 * Code lens has a resolve provider as well.
	 */
	ResolveProvider *bool `json:"resolveProvider,omitempty"`
}

type CodeLensRegistrationOptions struct {
	TextDocumentRegistrationOptions
	CodeLensOptions
}

const MethodTextDocumentCodeLens = Method("textDocument/codeLens")

type TextDocumentCodeLensFunc func(context *glsp.Context, params *CodeLensParams) ([]CodeLens, error)

type CodeLensParams struct {
	WorkDoneProgressParams
	PartialResultParams

	/**
	 * The document to request code lens for.
	 */
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

/**
 * A code lens represents a command that should be shown along with
 * source text, like the number of references, a way to run tests, etc.
 *
 * A code lens is _unresolved_ when no command is associated to it. For
 * performance reasons the creation of a code lens and resolving should be done
 * in two stages.
 */
type CodeLens struct {
	/**
	 * The range in which this code lens is valid. Should only span a single
	 * line.
	 */
	Range Range `json:"range"`

	/**
	 * The command this code lens represents.
	 */
	Command *Command `json:"command,omitempty"`

	/**
	 * A data entry field that is preserved on a code lens item between
	 * a code lens and a code lens resolve request.
	 */
	Data any `json:"data,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#codeLens_resolve

const MethodCodeLensResolve = Method("codeLens/resolve")

type CodeLensResolveFunc func(context *glsp.Context, params *CodeLens) (*CodeLens, error)

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#codeLens_refresh

type CodeLensWorkspaceClientCapabilities struct {
	/**
	 * Whether the client implementation supports a refresh request sent from the
	 * server to the client.
	 *
	 * Note that this event is global and will force the client to refresh all
	 * code lenses currently shown. It should be used with absolute care and is
	 * useful for situation where a server for example detect a project wide
	 * change that requires such a calculation.
	 */
	RefreshSupport *bool `json:"refreshSupport,omitempty"`
}

const ServerWorkspaceCodeLensRefresh = Method("workspace/codeLens/refresh")

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#textDocument_documentLink

type DocumentLinkClientCapabilities struct {
	/**
	 * Whether document link supports dynamic registration.
	 */
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`

	/**
	 * Whether the client supports the `tooltip` property on `DocumentLink`.
	 *
	 * @since 3.15.0
	 */
	TooltipSupport *bool `json:"tooltipSupport,omitempty"`
}

type DocumentLinkOptions struct {
	WorkDoneProgressOptions

	/**
	 * Document links have a resolve provider as well.
	 */
	ResolveProvider *bool `json:"resolveProvider,omitempty"`
}

type DocumentLinkRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DocumentLinkOptions
}

const MethodTextDocumentDocumentLink = Method("textDocument/documentLink")

type TextDocumentDocumentLinkFunc func(context *glsp.Context, params *DocumentLinkParams) ([]DocumentLink, error)

type DocumentLinkParams struct {
	WorkDoneProgressParams
	PartialResultParams

	/**
	 * The document to provide document links for.
	 */
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

/**
 * A document link is a range in a text document that links to an internal or
 * external resource, like another text document or a web site.
 */
type DocumentLink struct {
	/**
	 * The range this link applies to.
	 */
	Range Range `json:"range"`

	/**
	 * The uri this link points to. If missing a resolve request is sent later.
	 */
	Target *DocumentUri `json:"target,omitempty"`

	/**
	 * The tooltip text when you hover over this link.
	 *
	 * If a tooltip is provided, is will be displayed in a string that includes
	 * instructions on how to trigger the link, such as `{0} (ctrl + click)`.
	 * The specific instructions vary depending on OS, user settings, and
	 * localization.
	 *
	 * @since 3.15.0
	 */
	Tooltip *string `json:"tooltip,omitempty"`

	/**
	 * A data entry field that is preserved on a document link between a
	 * DocumentLinkRequest and a DocumentLinkResolveRequest.
	 */
	Data any `json:"data,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#documentLink_resolve

const MethodDocumentLinkResolve = Method("documentLink/resolve")

type DocumentLinkResolveFunc func(context *glsp.Context, params *DocumentLink) (*DocumentLink, error)

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#textDocument_documentColor

type DocumentColorClientCapabilities struct {
	/**
	 * Whether document color supports dynamic registration.
	 */
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

type DocumentColorOptions struct {
	WorkDoneProgressOptions
}

type DocumentColorRegistrationOptions struct {
	TextDocumentRegistrationOptions
	StaticRegistrationOptions
	DocumentColorOptions
}

const MethodTextDocumentColor = Method("textDocument/documentColor")

type TextDocumentColorFunc func(context *glsp.Context, params *DocumentColorParams) ([]ColorInformation, error)

type DocumentColorParams struct {
	WorkDoneProgressParams
	PartialResultParams

	/**
	 * The text document.
	 */
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

type ColorInformation struct {
	/**
	 * The range in the document where this color appears.
	 */
	Range Range `json:"range"`

	/**
	 * The actual color value for this color range.
	 */
	Color Color `json:"color"`
}

/**
 * Represents a color in RGBA space.
 */
type Color struct {
	/**
	 * The red component of this color in the range [0-1].
	 */
	Red Decimal `json:"red"`

	/**
	 * The green component of this color in the range [0-1].
	 */
	Green Decimal `json:"green"`

	/**
	 * The blue component of this color in the range [0-1].
	 */
	Blue Decimal `json:"blue"`

	/**
	 * The alpha component of this color in the range [0-1].
	 */
	Alpha Decimal `json:"alpha"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#textDocument_colorPresentation

const MethodTextDocumentColorPresentation = Method("textDocument/colorPresentation")

type TextDocumentColorPresentationFunc func(context *glsp.Context, params *ColorPresentationParams) ([]ColorPresentation, error)

type ColorPresentationParams struct {
	WorkDoneProgressParams
	PartialResultParams

	/**
	 * The text document.
	 */
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	/**
	 * The color information to request presentations for.
	 */
	Color Color `json:"color"`

	/**
	 * The range where the color would be inserted. Serves as a context.
	 */
	Range Range `json:"range"`
}

type ColorPresentation struct {
	/**
	 * The label of this color presentation. It will be shown on the color
	 * picker header. By default this is also the text that is inserted when
	 * selecting this color presentation.
	 */
	Label string `json:"label"`

	/**
	 * An [edit](#TextEdit) which is applied to a document when selecting
	 * this presentation for the color.  When `falsy` the
	 * [label](#ColorPresentation.label) is used.
	 */
	TextEdit *TextEdit `json:"textEdit,omitempty"`

	/**
	 * An optional array of additional [text edits](#TextEdit) that are applied
	 * when selecting this color presentation. Edits must not overlap with the
	 * main [edit](#ColorPresentation.textEdit) nor with themselves.
	 */
	AdditionalTextEdits []TextEdit `json:"additionalTextEdits,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#textDocument_formatting

type DocumentFormattingClientCapabilities struct {
	/**
	 * Whether formatting supports dynamic registration.
	 */
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

type DocumentFormattingOptions struct {
	WorkDoneProgressOptions
}

type DocumentFormattingRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DocumentFormattingOptions
}

const MethodTextDocumentFormatting = Method("textDocument/formatting")

type TextDocumentFormattingFunc func(context *glsp.Context, params *DocumentFormattingParams) ([]TextEdit, error)

type DocumentFormattingParams struct {
	WorkDoneProgressParams

	/**
	 * The document to format.
	 */
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	/**
	 * The format options.
	 */
	Options FormattingOptions `json:"options"`
}

type FormattingOptions map[string]any // bool | Integer | string

/**
 * Value-object describing what options formatting should use.
 */
const (
	/**
	 * Size of a tab in spaces.
	 */
	FormattingOptionTabSize = "tabSize"

	/**
	 * Prefer spaces over tabs.
	 */
	FormattingOptionInsertSpaces = "insertSpaces"

	/**
	 * Trim trailing whitespace on a line.
	 *
	 * @since 3.15.0
	 */
	FormattingOptionTrimTrailingWhitespace = "trimTrailingWhitespace"

	/**
	 * Insert a newline character at the end of the file if one does not exist.
	 *
	 * @since 3.15.0
	 */
	FormattingOptionInsertFinalNewline = "insertFinalNewline"

	/**
	 * Trim all newlines after the final newline at the end of the file.
	 *
	 * @since 3.15.0
	 */
	FormattingOptionTrimFinalNewlines = "trimFinalNewlines"
)

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#textDocument_rangeFormatting

type DocumentRangeFormattingClientCapabilities struct {
	/**
	 * Whether formatting supports dynamic registration.
	 */
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

type DocumentRangeFormattingOptions struct {
	WorkDoneProgressOptions
}

type DocumentRangeFormattingRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DocumentRangeFormattingOptions
}

const MethodTextDocumentRangeFormatting = Method("textDocument/rangeFormatting")

type TextDocumentRangeFormattingFunc func(context *glsp.Context, params *DocumentRangeFormattingParams) ([]TextEdit, error)

type DocumentRangeFormattingParams struct {
	WorkDoneProgressParams

	/**
	 * The document to format.
	 */
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	/**
	 * The range to format
	 */
	Range Range `json:"range"`

	/**
	 * The format options
	 */
	Options FormattingOptions `json:"options"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#textDocument_onTypeFormatting

type DocumentOnTypeFormattingClientCapabilities struct {
	/**
	 * Whether on type formatting supports dynamic registration.
	 */
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

type DocumentOnTypeFormattingOptions struct {
	/**
	 * A character on which formatting should be triggered, like `}`.
	 */
	FirstTriggerCharacter string `json:"firstTriggerCharacter"`

	/**
	 * More trigger characters.
	 */
	MoreTriggerCharacter []string `json:"moreTriggerCharacter,omitempty"`
}

type DocumentOnTypeFormattingRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DocumentOnTypeFormattingOptions
}

const MethodTextDocumentOnTypeFormatting = Method("textDocument/onTypeFormatting")

type TextDocumentOnTypeFormattingFunc func(context *glsp.Context, params *DocumentOnTypeFormattingParams) ([]TextEdit, error)

type DocumentOnTypeFormattingParams struct {
	TextDocumentPositionParams
	/**
	 * The character that has been typed.
	 */
	Ch string `json:"ch"`

	/**
	 * The format options.
	 */
	Options FormattingOptions `json:"options"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#textDocument_rename

type PrepareSupportDefaultBehavior Integer

const (
	/**
	 * The client's default behavior is to select the identifier
	 * according the to language's syntax rule.
	 */
	PrepareSupportDefaultBehaviorIdentifier = PrepareSupportDefaultBehavior(1)
)

type RenameClientCapabilities struct {
	/**
	 * Whether rename supports dynamic registration.
	 */
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`

	/**
	 * Client supports testing for validity of rename operations
	 * before execution.
	 *
	 * @since 3.12.0
	 */
	PrepareSupport *bool `json:"prepareSupport,omitempty"`

	/**
	 * Client supports the default behavior result
	 * (`{ defaultBehavior: boolean }`).
	 *
	 * The value indicates the default behavior used by the
	 * client.
	 *
	 * @since 3.16.0
	 */
	PrepareSupportDefaultBehavior *PrepareSupportDefaultBehavior `json:"prepareSupportDefaultBehavior,omitempty"`

	/**
	 * Whether th client honors the change annotations in
	 * text edits and resource operations returned via the
	 * rename request's workspace edit by for example presenting
	 * the workspace edit in the user interface and asking
	 * for confirmation.
	 *
	 * @since 3.16.0
	 */
	HonorsChangeAnnotations *bool `json:"honorsChangeAnnotations,omitempty"`
}

type RenameOptions struct {
	WorkDoneProgressOptions

	/**
	 * Renames should be checked and tested before being executed.
	 */
	PrepareProvider *bool `json:"prepareProvider,omitempty"`
}

type RenameRegistrationOptions struct {
	TextDocumentRegistrationOptions
	RenameOptions
}

const MethodTextDocumentRename = Method("textDocument/rename")

type TextDocumentRenameFunc func(context *glsp.Context, params *RenameParams) (*WorkspaceEdit, error)

type RenameParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams

	/**
	 * The new name of the symbol. If the given name is not valid the
	 * request must return a [ResponseError](#ResponseError) with an
	 * appropriate message set.
	 */
	NewName string `json:"newName"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#textDocument_prepareRename

const MethodTextDocumentPrepareRename = Method("textDocument/prepareRename")

// Returns: Range | RangeWithPlaceholder | DefaultBehavior | nil
type TextDocumentPrepareRenameFunc func(context *glsp.Context, params *PrepareRenameParams) (any, error)

type PrepareRenameParams struct {
	TextDocumentPositionParams
}

type RangeWithPlaceholder struct {
	Range       Range  `json:"range"`
	Placeholder string `json:"placeholder"`
}

type DefaultBehavior struct {
	DefaultBehavior bool `json:"defaultBehavior"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#textDocument_foldingRange

type FoldingRangeClientCapabilities struct {
	/**
	 * Whether implementation supports dynamic registration for folding range
	 * providers. If this is set to `true` the client supports the new
	 * `FoldingRangeRegistrationOptions` return value for the corresponding
	 * server capability as well.
	 */
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`

	/**
	 * The maximum number of folding ranges that the client prefers to receive
	 * per document. The value serves as a hint, servers are free to follow the
	 * limit.
	 */
	RangeLimit *UInteger `json:"rangeLimit,omitempty"`

	/**
	 * If set, the client signals that it only supports folding complete lines.
	 * If set, client will ignore specified `startCharacter` and `endCharacter`
	 * properties in a FoldingRange.
	 */
	LineFoldingOnly *bool `json:"lineFoldingOnly,omitempty"`
}

type FoldingRangeOptions struct {
	WorkDoneProgressOptions
}

type FoldingRangeRegistrationOptions struct {
	TextDocumentRegistrationOptions
	FoldingRangeOptions
	StaticRegistrationOptions
}

const MethodTextDocumentFoldingRange = Method("textDocument/foldingRange")

type TextDocumentFoldingRangeFunc func(context *glsp.Context, params *FoldingRangeParams) ([]FoldingRange, error)

type FoldingRangeParams struct {
	WorkDoneProgressParams
	PartialResultParams

	/**
	 * The text document.
	 */
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

/**
 * Enum of known range kinds
 */
type FoldingRangeKind string

const (
	/**
	 * Folding range for a comment
	 */
	FoldingRangeKindComment = FoldingRangeKind("comment")

	/**
	 * Folding range for a imports or includes
	 */
	FoldingRangeKindImports = FoldingRangeKind("imports")

	/**
	 * Folding range for a region (e.g. `#region`)
	 */
	FoldingRangeKindRegion = FoldingRangeKind("region")
)

/**
 * Represents a folding range. To be valid, start and end line must be bigger
 * than zero and smaller than the number of lines in the document. Clients
 * are free to ignore invalid ranges.
 */
type FoldingRange struct {
	/**
	 * The zero-based start line of the range to fold. The folded area starts
	 * after the line's last character. To be valid, the end must be zero or
	 * larger and smaller than the number of lines in the document.
	 */
	StartLine UInteger `json:"startLine"`

	/**
	 * The zero-based character offset from where the folded range starts. If
	 * not defined, defaults to the length of the start line.
	 */
	StartCharacter *UInteger `json:"startCharacter,omitempty"`

	/**
	 * The zero-based end line of the range to fold. The folded area ends with
	 * the line's last character. To be valid, the end must be zero or larger
	 * and smaller than the number of lines in the document.
	 */
	EndLine UInteger `json:"endLine"`

	/**
	 * The zero-based character offset before the folded range ends. If not
	 * defined, defaults to the length of the end line.
	 */
	EndCharacter *UInteger `json:"endCharacter,omitempty"`

	/**
	 * Describes the kind of the folding range such as `comment` or `region`.
	 * The kind is used to categorize folding ranges and used by commands like
	 * 'Fold all comments'. See [FoldingRangeKind](#FoldingRangeKind) for an
	 * enumeration of standardized kinds.
	 */
	Kind *string `json:"kind,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#textDocument_selectionRange

type SelectionRangeClientCapabilities struct {
	/**
	 * Whether implementation supports dynamic registration for selection range
	 * providers. If this is set to `true` the client supports the new
	 * `SelectionRangeRegistrationOptions` return value for the corresponding
	 * server capability as well.
	 */
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

type SelectionRangeOptions struct {
	WorkDoneProgressOptions
}

type SelectionRangeRegistrationOptions struct {
	SelectionRangeOptions
	TextDocumentRegistrationOptions
	StaticRegistrationOptions
}

const MethodTextDocumentSelectionRange = Method("textDocument/selectionRange")

type TextDocumentSelectionRangeFunc func(context *glsp.Context, params *SelectionRangeParams) ([]SelectionRange, error)

type SelectionRangeParams struct {
	WorkDoneProgressParams
	PartialResultParams

	/**
	 * The text document.
	 */
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	/**
	 * The positions inside the text document.
	 */
	Positions []Position `json:"positions"`
}

type SelectionRange struct {
	/**
	 * The [range](#Range) of this selection range.
	 */
	Range Range `json:"range"`

	/**
	 * The parent selection range containing this range. Therefore
	 * `parent.range` must contain `this.range`.
	 */
	Parent *SelectionRange `json:"parent,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#textDocument_prepareCallHierarchy

type CallHierarchyClientCapabilities struct {
	/**
	 * Whether implementation supports dynamic registration. If this is set to
	 * `true` the client supports the new `(TextDocumentRegistrationOptions &
	 * StaticRegistrationOptions)` return value for the corresponding server
	 * capability as well.
	 */
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

type CallHierarchyOptions struct {
	WorkDoneProgressOptions
}

type CallHierarchyRegistrationOptions struct {
	TextDocumentRegistrationOptions
	CallHierarchyOptions
	StaticRegistrationOptions
}

const MethodTextDocumentPrepareCallHierarchy = Method("textDocument/prepareCallHierarchy")

type TextDocumentPrepareCallHierarchyFunc func(context *glsp.Context, params *CallHierarchyPrepareParams) ([]CallHierarchyItem, error)

type CallHierarchyPrepareParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
}

type CallHierarchyItem struct {
	/**
	 * The name of this item.
	 */
	Name string `json:"name"`

	/**
	 * The kind of this item.
	 */
	Kind SymbolKind `json:"kind"`

	/**
	 * Tags for this item.
	 */
	Tags []SymbolTag `json:"tags,omitempty"`

	/**
	 * More detail for this item, e.g. the signature of a function.
	 */
	Detail *string `json:"detail,omitempty"`

	/**
	 * The resource identifier of this item.
	 */
	URI DocumentUri `json:"uri"`

	/**
	 * The range enclosing this symbol not including leading/trailing whitespace
	 * but everything else, e.g. comments and code.
	 */
	Range Range `json:"range"`

	/**
	 * The range that should be selected and revealed when this symbol is being
	 * picked, e.g. the name of a function. Must be contained by the
	 * [`range`](#CallHierarchyItem.range).
	 */
	SelectionRange Range `json:"selectionRange"`

	/**
	 * A data entry field that is preserved between a call hierarchy prepare and
	 * incoming calls or outgoing calls requests.
	 */
	Data any `json:"data,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#callHierarchy_incomingCalls

const MethodCallHierarchyIncomingCalls = Method("callHierarchy/incomingCalls")

type CallHierarchyIncomingCallsFunc func(context *glsp.Context, params *CallHierarchyIncomingCallsParams) ([]CallHierarchyIncomingCall, error)

type CallHierarchyIncomingCallsParams struct {
	WorkDoneProgressParams
	PartialResultParams

	Item CallHierarchyItem `json:"item"`
}

type CallHierarchyIncomingCall struct {
	/**
	 * The item that makes the call.
	 */
	From CallHierarchyItem `json:"from"`

	/**
	 * The ranges at which the calls appear. This is relative to the caller
	 * denoted by [`this.from`](#CallHierarchyIncomingCall.from).
	 */
	FromRanges []Range `json:"fromRanges"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#callHierarchy_outgoingCalls

const MethodCallHierarchyOutgoingCalls = Method("callHierarchy/outgoingCalls")

type CallHierarchyOutgoingCallsFunc func(context *glsp.Context, params *CallHierarchyOutgoingCallsParams) ([]CallHierarchyOutgoingCall, error)

type CallHierarchyOutgoingCallsParams struct {
	WorkDoneProgressParams
	PartialResultParams

	Item CallHierarchyItem `json:"item"`
}

type CallHierarchyOutgoingCall struct {
	/**
	 * The item that is called.
	 */
	To CallHierarchyItem `json:"to"`

	/**
	 * The range at which this item is called. This is the range relative to
	 * the caller, e.g the item passed to `callHierarchy/outgoingCalls` request.
	 */
	FromRanges []Range `json:"fromRanges"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#textDocument_semanticTokens

type SemanticTokenType string

const (
	SemanticTokenTypeNamespace = SemanticTokenType("namespace")
	/**
	 * Represents a generic type. Acts as a fallback for types which
	 * can't be mapped to a specific type like class or enum.
	 */
	SemanticTokenTypeType          = SemanticTokenType("type")
	SemanticTokenTypeClass         = SemanticTokenType("class")
	SemanticTokenTypeEnum          = SemanticTokenType("enum")
	SemanticTokenTypeInterface     = SemanticTokenType("interface")
	SemanticTokenTypeStruct        = SemanticTokenType("struct")
	SemanticTokenTypeTypeParameter = SemanticTokenType("typeParameter")
	SemanticTokenTypeParameter     = SemanticTokenType("parameter")
	SemanticTokenTypeVariable      = SemanticTokenType("variable")
	SemanticTokenTypeProperty      = SemanticTokenType("property")
	SemanticTokenTypeEnumMember    = SemanticTokenType("enumMember")
	SemanticTokenTypeEvent         = SemanticTokenType("event")
	SemanticTokenTypeFunction      = SemanticTokenType("function")
	SemanticTokenTypeMethod        = SemanticTokenType("method")
	SemanticTokenTypeMacro         = SemanticTokenType("macro")
	SemanticTokenTypeKeyword       = SemanticTokenType("keyword")
	SemanticTokenTypeModifier      = SemanticTokenType("modifier")
	SemanticTokenTypeComment       = SemanticTokenType("comment")
	SemanticTokenTypeString        = SemanticTokenType("string")
	SemanticTokenTypeNumber        = SemanticTokenType("number")
	SemanticTokenTypeRegexp        = SemanticTokenType("regexp")
	SemanticTokenTypeOperator      = SemanticTokenType("operator")
)

type SemanticTokenModifier string

const (
	SemanticTokenModifierDeclaration    = SemanticTokenModifier("declaration")
	SemanticTokenModifierDefinition     = SemanticTokenModifier("definition")
	SemanticTokenModifierReadonly       = SemanticTokenModifier("readonly")
	SemanticTokenModifierStatic         = SemanticTokenModifier("static")
	SemanticTokenModifierDeprecated     = SemanticTokenModifier("deprecated")
	SemanticTokenModifierAbstract       = SemanticTokenModifier("abstract")
	SemanticTokenModifierAsync          = SemanticTokenModifier("async")
	SemanticTokenModifierModification   = SemanticTokenModifier("modification")
	SemanticTokenModifierDocumentation  = SemanticTokenModifier("documentation")
	SemanticTokenModifierDefaultLibrary = SemanticTokenModifier("defaultLibrary")
)

type TokenFormat string

const (
	TokenFormatRelative = TokenFormat("relative")
)

type SemanticTokensLegend struct {
	/**
	 * The token types a server uses.
	 */
	TokenTypes []string `json:"tokenTypes"`

	/**
	 * The token modifiers a server uses.
	 */
	TokenModifiers []string `json:"tokenModifiers"`
}

type SemanticDelta struct {
	Delta *bool `json:"delta,omitempty"`
}

type SemanticTokensClientCapabilities struct {
	/**
	 * Whether implementation supports dynamic registration. If this is set to
	 * `true` the client supports the new `(TextDocumentRegistrationOptions &
	 * StaticRegistrationOptions)` return value for the corresponding server
	 * capability as well.
	 */
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`

	/**
	 * Which requests the client supports and might send to the server
	 * depending on the server's capability. Please note that clients might not
	 * show semantic tokens or degrade some of the user experience if a range
	 * or full request is advertised by the client but not provided by the
	 * server. If for example the client capability `requests.full` and
	 * `request.range` are both set to true but the server only provides a
	 * range provider the client might not render a minimap correctly or might
	 * even decide to not show any semantic tokens at all.
	 */
	Requests struct {
		/**
		 * The client will send the `textDocument/semanticTokens/range` request
		 * if the server provides a corresponding handler.
		 */
		Range any `json:"Range,omitempty"` // nil | bool | struct{}

		/**
		 * The client will send the `textDocument/semanticTokens/full` request
		 * if the server provides a corresponding handler.
		 */
		Full any `json:"full,omitempty"` // nil | bool | SemanticDelta
	} `json:"requests"`

	/**
	 * The token types that the client supports.
	 */
	TokenTypes []string `json:"tokenTypes"`

	/**
	 * The token modifiers that the client supports.
	 */
	TokenModifiers []string `json:"tokenModifiers"`

	/**
	 * The formats the clients supports.
	 */
	Formats []TokenFormat `json:"formats"`

	/**
	 * Whether the client supports tokens that can overlap each other.
	 */
	OverlappingTokenSupport *bool `json:"overlappingTokenSupport,omitempty"`

	/**
	 * Whether the client supports tokens that can span multiple lines.
	 */
	MultilineTokenSupport *bool `json:"multilineTokenSupport,omitempty"`
}

// json.Unmarshaler interface
func (self *SemanticTokensClientCapabilities) UnmarshalJSON(data []byte) error {
	var value struct {
		DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
		Requests            struct {
			Range json.RawMessage `json:"Range,omitempty"` // nil | bool | struct{}
			Full  json.RawMessage `json:"full,omitempty"`  // nil | bool | SemanticDelta
		} `json:"requests"`
		TokenTypes              []string      `json:"tokenTypes"`
		TokenModifiers          []string      `json:"tokenModifiers"`
		Formats                 []TokenFormat `json:"formats"`
		OverlappingTokenSupport *bool         `json:"overlappingTokenSupport,omitempty"`
		MultilineTokenSupport   *bool         `json:"multilineTokenSupport,omitempty"`
	}

	if err := json.Unmarshal(data, &value); err == nil {
		self.DynamicRegistration = value.DynamicRegistration
		self.TokenTypes = value.TokenTypes
		self.TokenModifiers = value.TokenModifiers
		self.Formats = value.Formats
		self.OverlappingTokenSupport = value.OverlappingTokenSupport
		self.MultilineTokenSupport = value.MultilineTokenSupport

		if value.Requests.Range != nil {
			var value_ bool
			if err = json.Unmarshal(value.Requests.Range, &value_); err == nil {
				self.Requests.Range = value_
			} else {
				var value_ struct{}
				if err = json.Unmarshal(value.Requests.Range, &value_); err == nil {
					self.Requests.Range = value_
				} else {
					return err
				}
			}
		}

		if value.Requests.Full != nil {
			var value_ bool
			if err = json.Unmarshal(value.Requests.Full, &value_); err == nil {
				self.Requests.Full = value_
			} else {
				var value_ SemanticDelta
				if err = json.Unmarshal(value.Requests.Full, &value_); err == nil {
					self.Requests.Full = value_
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

type SemanticTokensOptions struct {
	WorkDoneProgressOptions

	/**
	 * The legend used by the server
	 */
	Legend SemanticTokensLegend `json:"legend"`

	/**
	 * Server supports providing semantic tokens for a specific range
	 * of a document.
	 */
	Range any `json:"range,omitempty"` // nil | bool | struct{}

	/**
	 * Server supports providing semantic tokens for a full document.
	 */
	Full any `json:"full,omitempty"` // nil | bool | SemanticDelta
}

// json.Unmarshaler interface
func (self *SemanticTokensOptions) UnmarshalJSON(data []byte) error {
	var value struct {
		Legend SemanticTokensLegend `json:"legend"`
		Range  json.RawMessage      `json:"range,omitempty"` // nil | bool | struct{}
		Full   json.RawMessage      `json:"full,omitempty"`  // nil | bool | SemanticDelta
	}

	if err := json.Unmarshal(data, &value); err == nil {
		self.Legend = value.Legend

		if value.Range != nil {
			var value_ bool
			if err = json.Unmarshal(value.Range, &value_); err == nil {
				self.Range = value_
			} else {
				var value_ struct{}
				if err = json.Unmarshal(value.Range, &value_); err == nil {
					self.Range = value_
				} else {
					return err
				}
			}
		}

		if value.Full != nil {
			var value_ bool
			if err = json.Unmarshal(value.Full, &value_); err == nil {
				self.Full = value_
			} else {
				var value_ SemanticDelta
				if err = json.Unmarshal(value.Full, &value_); err == nil {
					self.Full = value_
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

type SemanticTokensRegistrationOptions struct {
	TextDocumentRegistrationOptions
	SemanticTokensOptions
	StaticRegistrationOptions
}

const MethodTextDocumentSemanticTokensFull = Method("textDocument/semanticTokens/full")

type TextDocumentSemanticTokensFullFunc func(context *glsp.Context, params *SemanticTokensParams) (*SemanticTokens, error)

type SemanticTokensParams struct {
	WorkDoneProgressParams
	PartialResultParams

	/**
	 * The text document.
	 */
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

type SemanticTokens struct {
	/**
	 * An optional result id. If provided and clients support delta updating
	 * the client will include the result id in the next semantic token request.
	 * A server can then instead of computing all semantic tokens again simply
	 * send a delta.
	 */
	ResultID *string `json:"resultId,omitempty"`

	/**
	 * The actual tokens.
	 */
	Data []UInteger `json:"data"`
}

type SemanticTokensPartialResult struct {
	Data []UInteger `json:"data"`
}

const MethodTextDocumentSemanticTokensFullDelta = Method("textDocument/semanticTokens/full/delta")

// Returns: SemanticTokens | SemanticTokensDelta | SemanticTokensDeltaPartialResult | nil
type TextDocumentSemanticTokensFullDeltaFunc func(context *glsp.Context, params *SemanticTokensDeltaParams) (any, error)

type SemanticTokensDeltaParams struct {
	WorkDoneProgressParams
	PartialResultParams

	/**
	 * The text document.
	 */
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	/**
	 * The result id of a previous response. The result Id can either point to
	 * a full response or a delta response depending on what was received last.
	 */
	PreviousResultID string `json:"previousResultId"`
}

type SemanticTokensDelta struct {
	ResultId *string `json:"resultId,omitempty"`

	/**
	 * The semantic token edits to transform a previous result into a new
	 * result.
	 */
	Edits []SemanticTokensEdit `json:"edits"`
}

type SemanticTokensEdit struct {
	/**
	 * The start offset of the edit.
	 */
	Start UInteger `json:"start"`

	/**
	 * The count of elements to remove.
	 */
	DeleteCount UInteger `json:"deleteCount"`

	/**
	 * The elements to insert.
	 */
	Data []UInteger `json:"data,omitempty"`
}

type SemanticTokensDeltaPartialResult struct {
	Edits []SemanticTokensEdit `json:"edits"`
}

const MethodTextDocumentSemanticTokensRange = Method("textDocument/semanticTokens/range")

// Returns: SemanticTokens | SemanticTokensPartialResult | nil
type TextDocumentSemanticTokensRangeFunc func(context *glsp.Context, params *SemanticTokensRangeParams) (any, error)

type SemanticTokensRangeParams struct {
	WorkDoneProgressParams
	PartialResultParams

	/**
	 * The text document.
	 */
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	/**
	 * The range the semantic tokens are requested for.
	 */
	Range Range `json:"range"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#textDocument_linkedEditingRange

type LinkedEditingRangeClientCapabilities struct {
	/**
	 * Whether implementation supports dynamic registration.
	 * If this is set to `true` the client supports the new
	 * `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
	 * return value for the corresponding server capability as well.
	 */
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

type LinkedEditingRangeOptions struct {
	WorkDoneProgressOptions
}

type LinkedEditingRangeRegistrationOptions struct {
	TextDocumentRegistrationOptions
	LinkedEditingRangeOptions
	StaticRegistrationOptions
}

const MethodTextDocumentLinkedEditingRange = Method("textDocument/linkedEditingRange")

type TextDocumentLinkedEditingRangeFunc func(context *glsp.Context, params *LinkedEditingRangeParams) (*LinkedEditingRanges, error)

type LinkedEditingRangeParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
}

type LinkedEditingRanges struct {
	/**
	 * A list of ranges that can be renamed together. The ranges must have
	 * identical length and contain identical text content. The ranges cannot overlap.
	 */
	Ranges []Range `json:"ranges"`

	/**
	 * An optional word pattern (regular expression) that describes valid contents for
	 * the given ranges. If no pattern is provided, the client configuration's word
	 * pattern will be used.
	 */
	WordPattern *string `json:"wordPattern,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#textDocument_moniker

type MonikerClientCapabilities struct {
	/**
	 * Whether implementation supports dynamic registration. If this is set to
	 * `true` the client supports the new `(TextDocumentRegistrationOptions &
	 * StaticRegistrationOptions)` return value for the corresponding server
	 * capability as well.
	 */
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

type MonikerOptions struct {
	WorkDoneProgressOptions
}

type MonikerRegistrationOptions struct {
	TextDocumentRegistrationOptions
	MonikerOptions
}

const MethodTextDocumentMoniker = Method("textDocument/moniker")

type TextDocumentMonikerFunc func(context *glsp.Context, params *MonikerParams) ([]Moniker, error)

type MonikerParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
}

/**
 * Moniker uniqueness level to define scope of the moniker.
 */
type UniquenessLevel string

const (
	/**
	 * The moniker is only unique inside a document
	 */
	UniquenessLevelDocument = UniquenessLevel("document")

	/**
	 * The moniker is unique inside a project for which a dump got created
	 */
	UniquenessLevelProject = UniquenessLevel("project")

	/**
	 * The moniker is unique inside the group to which a project belongs
	 */
	UniquenessLevelGroup = UniquenessLevel("group")

	/**
	 * The moniker is unique inside the moniker scheme.
	 */
	UniquenessLevelScheme = UniquenessLevel("scheme")

	/**
	 * The moniker is globally unique
	 */
	UniquenessLevelGlobal = UniquenessLevel("global")
)

/**
 * The moniker kind.
 */
type MonikerKind string

const (
	/**
	 * The moniker represent a symbol that is imported into a project
	 */
	MonikerKindImport = MonikerKind("import")

	/**
	 * The moniker represents a symbol that is exported from a project
	 */
	MonikerKindExport = MonikerKind("export")

	/**
	 * The moniker represents a symbol that is local to a project (e.g. a local
	 * variable of a function, a class not visible outside the project, ...)
	 */
	MonikerKindLocal = MonikerKind("local")
)

/**
 * Moniker definition to match LSIF 0.5 moniker definition.
 */
type Moniker struct {
	/**
	 * The scheme of the moniker. For example tsc or .Net
	 */
	Scheme string `json:"scheme"`

	/**
	 * The identifier of the moniker. The value is opaque in LSIF however
	 * schema owners are allowed to define the structure if they want.
	 */
	Identifier string `json:"identifier"`

	/**
	 * The scope in which the moniker is unique
	 */
	Unique UniquenessLevel `json:"unique"`

	/**
	 * The moniker kind if known.
	 */
	Kind *MonikerKind `json:"kind,omitempty"`
}
