package protocol

import (
	"encoding/json"

	"github.com/tliron/glsp"
)

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#textDocument_synchronization

type TextDocumentSyncKind Integer

/**
 * Defines how the host (editor) should sync document changes to the language
 * server.
 */
const (
	/**
	 * Documents should not be synced at all.
	 */
	TextDocumentSyncKindNone = TextDocumentSyncKind(0)

	/**
	 * Documents are synced by always sending the full content
	 * of the document.
	 */
	TextDocumentSyncKindFull = TextDocumentSyncKind(1)

	/**
	 * Documents are synced by sending the full content on open.
	 * After that only incremental updates to the document are
	 * send.
	 */
	TextDocumentSyncKindIncremental = TextDocumentSyncKind(2)
)

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#textDocument_didOpen

const MethodTextDocumentDidOpen = Method("textDocument/didOpen")

type TextDocumentDidOpenFunc func(context *glsp.Context, params *DidOpenTextDocumentParams) error

type DidOpenTextDocumentParams struct {
	/**
	 * The document that was opened.
	 */
	TextDocument TextDocumentItem `json:"textDocument"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#textDocument_didChange

/**
 * Describe options to be used when registering for text document change events.
 */
type TextDocumentChangeRegistrationOptions struct {
	TextDocumentRegistrationOptions

	/**
	 * How documents are synced to the server. See TextDocumentSyncKind.Full
	 * and TextDocumentSyncKind.Incremental.
	 */
	SyncKind TextDocumentSyncKind `json:"syncKind"`
}

const MethodTextDocumentDidChange = Method("textDocument/didChange")

type TextDocumentDidChangeFunc func(context *glsp.Context, params *DidChangeTextDocumentParams) error

type DidChangeTextDocumentParams struct {
	/**
	 * The document that did change. The version number points
	 * to the version after all provided content changes have
	 * been applied.
	 */
	TextDocument VersionedTextDocumentIdentifier `json:"textDocument"`

	/**
	 * The actual content changes. The content changes describe single state
	 * changes to the document. So if there are two content changes c1 (at
	 * array index 0) and c2 (at array index 1) for a document in state S then
	 * c1 moves the document from S to S' and c2 from S' to S''. So c1 is
	 * computed on the state S and c2 is computed on the state S'.
	 *
	 * To mirror the content of a document using change events use the following
	 * approach:
	 * - start with the same initial content
	 * - apply the 'textDocument/didChange' notifications in the order you
	 *   receive them.
	 * - apply the `TextDocumentContentChangeEvent`s in a single notification
	 *   in the order you receive them.
	 */
	ContentChanges []any `json:"contentChanges"` // TextDocumentContentChangeEvent or TextDocumentContentChangeEventWhole
}

// json.Unmarshaler interface
func (self *DidChangeTextDocumentParams) UnmarshalJSON(data []byte) error {
	var value struct {
		TextDocument   VersionedTextDocumentIdentifier `json:"textDocument"`
		ContentChanges []json.RawMessage               `json:"contentChanges"` // TextDocumentContentChangeEvent or TextDocumentContentChangeEventWhole
	}

	if err := json.Unmarshal(data, &value); err == nil {
		self.TextDocument = value.TextDocument

		for _, contentChange := range value.ContentChanges {
			var changeEvent TextDocumentContentChangeEvent
			if err = json.Unmarshal(contentChange, &changeEvent); err == nil {
				if changeEvent.Range != nil {
					self.ContentChanges = append(self.ContentChanges, changeEvent)
				} else {
					changeEventWhole := TextDocumentContentChangeEventWhole{
						Text: changeEvent.Text,
					}
					self.ContentChanges = append(self.ContentChanges, changeEventWhole)
				}
			} else {
				return err
			}
		}

		return nil
	} else {
		return err
	}
}

/**
 * An event describing a change to a text document. If range and rangeLength are
 * omitted the new text is considered to be the full content of the document.
 */
type TextDocumentContentChangeEvent struct {
	/**
	 * The range of the document that changed.
	 */
	Range *Range `json:"range"`

	/**
	 * The optional length of the range that got replaced.
	 *
	 * @deprecated use range instead.
	 */
	RangeLength *UInteger `json:"rangeLength,omitempty"`

	/**
	 * The new text for the provided range.
	 */
	Text string `json:"text"`
}

type TextDocumentContentChangeEventWhole struct {
	/**
	 * The new text of the whole document.
	 */
	Text string `json:"text"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#textDocument_willSave

const MethodTextDocumentWillSave = Method("textDocument/willSave")

type TextDocumentWillSaveFunc func(context *glsp.Context, params *WillSaveTextDocumentParams) error

/**
 * The parameters send in a will save text document notification.
 */
type WillSaveTextDocumentParams struct {
	/**
	 * The document that will be saved.
	 */
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	/**
	 * The 'TextDocumentSaveReason'.
	 */
	Reason TextDocumentSaveReason `json:"reason"`
}

type TextDocumentSaveReason Integer

/**
 * Represents reasons why a text document is saved.
 */
const (
	/**
	 * Manually triggered, e.g. by the user pressing save, by starting
	 * debugging, or by an API call.
	 */
	TextDocumentSaveReasonManual = TextDocumentSaveReason(1)

	/**
	 * Automatic after a delay.
	 */
	TextDocumentSaveReasonAfterDelay = TextDocumentSaveReason(2)

	/**
	 * When the editor lost focus.
	 */
	TextDocumentSaveReasonFocusOut = TextDocumentSaveReason(3)
)

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#textDocument_willSaveWaitUntil

const MethodTextDocumentWillSaveWaitUntil = Method("textDocument/willSaveWaitUntil")

type TextDocumentWillSaveWaitUntilFunc func(context *glsp.Context, params *WillSaveTextDocumentParams) ([]TextEdit, error)

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#textDocument_didSave

type SaveOptions struct {
	/**
	 * The client is supposed to include the content on save.
	 */
	IncludeText *bool `json:"includeText,omitempty"`
}

type TextDocumentSaveRegistrationOptions struct {
	TextDocumentRegistrationOptions

	/**
	 * The client is supposed to include the content on save.
	 */
	IncludeText *bool `json:"includeText"`
}

const MethodTextDocumentDidSave = Method("textDocument/didSave")

type TextDocumentDidSaveFunc func(context *glsp.Context, params *DidSaveTextDocumentParams) error

type DidSaveTextDocumentParams struct {
	/**
	 * The document that was saved.
	 */
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	/**
	 * Optional the content when saved. Depends on the includeText value
	 * when the save notification was requested.
	 */
	Text *string `json:"text,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#textDocument_didClose

type TextDocumentSyncClientCapabilities struct {
	/**
	 * Whether text document synchronization supports dynamic registration.
	 */
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`

	/**
	 * The client supports sending will save notifications.
	 */
	WillSave *bool `json:"willSave,omitempty"`

	/**
	 * The client supports sending a will save request and
	 * waits for a response providing text edits which will
	 * be applied to the document before it is saved.
	 */
	WillSaveWaitUntil *bool `json:"willSaveWaitUntil,omitempty"`

	/**
	 * The client supports did save notifications.
	 */
	DidSave *bool `json:"didSave,omitempty"`
}

type TextDocumentSyncOptions struct {
	/**
	 * Open and close notifications are sent to the server. If omitted open
	 * close notification should not be sent.
	 */
	OpenClose *bool `json:"openClose,omitempty"`

	/**
	 * Change notifications are sent to the server. See
	 * TextDocumentSyncKind.None, TextDocumentSyncKind.Full and
	 * TextDocumentSyncKind.Incremental. If omitted it defaults to
	 * TextDocumentSyncKind.None.
	 */
	Change *TextDocumentSyncKind `json:"change,omitempty"`

	/**
	 * If present will save notifications are sent to the server. If omitted
	 * the notification should not be sent.
	 */
	WillSave *bool `json:"willSave,omitempty"`

	/**
	 * If present will save wait until requests are sent to the server. If
	 * omitted the request should not be sent.
	 */
	WillSaveWaitUntil *bool `json:"willSaveWaitUntil,omitempty"`

	/**
	 * If present save notifications are sent to the server. If omitted the
	 * notification should not be sent.
	 */
	Save any `json:"save,omitempty"` // nil | bool | SaveOptions
}

// json.Unmarshaler interface
func (self *TextDocumentSyncOptions) UnmarshalJSON(data []byte) error {
	var value struct {
		OpenClose         *bool                 `json:"openClose"`
		Change            *TextDocumentSyncKind `json:"change"`
		WillSave          *bool                 `json:"willSave"`
		WillSaveWaitUntil *bool                 `json:"willSaveWaitUntil"`
		Save              json.RawMessage       `json:"save"` // nil | bool | SaveOptions
	}

	if err := json.Unmarshal(data, &value); err == nil {
		self.OpenClose = value.OpenClose
		self.Change = value.Change
		self.WillSave = value.WillSave
		self.WillSaveWaitUntil = value.WillSaveWaitUntil

		if value.Save != nil {
			var value_ bool
			if err = json.Unmarshal(value.Save, &value_); err == nil {
				self.Save = value_
			} else {
				var value_ SaveOptions
				if err = json.Unmarshal(value.Save, &value_); err == nil {
					self.Save = value_
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

const MethodTextDocumentDidClose = Method("textDocument/didClose")

type TextDocumentDidCloseFunc func(context *glsp.Context, params *DidCloseTextDocumentParams) error

type DidCloseTextDocumentParams struct {
	/**
	 * The document that was closed.
	 */
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}
