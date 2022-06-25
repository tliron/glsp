package protocol

import "github.com/tliron/glsp"

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#window_showMessage

const ServerWindowShowMessage = Method("window/showMessage")

type ShowMessageParams struct {
	/**
	 * The message type. See {@link MessageType}.
	 */
	Type MessageType `json:"type"`

	/**
	 * The actual message.
	 */
	Message string `json:"message"`
}

type MessageType Integer

const (
	/**
	 * An error message.
	 */
	MessageTypeError = MessageType(1)

	/**
	 * A warning message.
	 */
	MessageTypeWarning = MessageType(2)

	/**
	 * An information message.
	 */
	MessageTypeInfo = MessageType(3)

	/**
	 * A log message.
	 */
	MessageTypeLog = MessageType(4)
)

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#window_showMessageRequest

type ShowMessageRequestClientCapabilities struct {
	/**
	 * Capabilities specific to the `MessageActionItem` type.
	 */
	MessageActionItem *struct {
		/**
		 * Whether the client supports additional attributes which
		 * are preserved and sent back to the server in the
		 * request's response.
		 */
		AdditionalPropertiesSupport *bool `json:"additionalPropertiesSupport,omitempty"`
	} `json:"messageActionItem,omitempty"`
}

const ServerWindowShowMessageRequest = Method("window/showMessageRequest")

type ShowMessageRequestParams struct {
	/**
	 * The message type. See {@link MessageType}
	 */
	Type MessageType `json:"type"`

	/**
	 * The actual message
	 */
	Message string `json:"message"`

	/**
	 * The message action items to present.
	 */
	Actions []MessageActionItem `json:"actions,omitempty"`
}

type MessageActionItem struct {
	/**
	 * A short title like 'Retry', 'Open Log' etc.
	 */
	Title string `json:"title"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#window_showDocument

/**
 * Client capabilities for the show document request.
 *
 * @since 3.16.0
 */
type ShowDocumentClientCapabilities struct {
	/**
	 * The client has support for the show document
	 * request.
	 */
	Support bool `json:"support"`
}

const ServerWindowShowDocument = Method("window/showDocument")

/**
 * Params to show a document.
 *
 * @since 3.16.0
 */
type ShowDocumentParams struct {
	/**
	 * The document uri to show.
	 */
	URI URI `json:"uri"`

	/**
	 * Indicates to show the resource in an external program.
	 * To show for example `https://code.visualstudio.com/`
	 * in the default WEB browser set `external` to `true`.
	 */
	External *bool `json:"external,omitempty"`

	/**
	 * An optional property to indicate whether the editor
	 * showing the document should take focus or not.
	 * Clients might ignore this property if an external
	 * program is started.
	 */
	TakeFocus *bool `json:"takeFocus,omitempty"`

	/**
	 * An optional selection range if the document is a text
	 * document. Clients might ignore the property if an
	 * external program is started or the file is not a text
	 * file.
	 */
	Selection *Range `json:"selection,omitempty"`
}

/**
 * The result of an show document request.
 *
 * @since 3.16.0
 */
type ShowDocumentResult struct {
	/**
	 * A boolean indicating if the show was successful.
	 */
	Success bool `json:"success"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#window_logMessage

const ServerWindowLogMessage = Method("window/logMessage")

type LogMessageParams struct {
	/**
	 * The message type. See {@link MessageType}
	 */
	Type MessageType `json:"type"`

	/**
	 * The actual message
	 */
	Message string `json:"message"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#window_workDoneProgress_create

const ServerWindowWorkDoneProgressCreate = Method("window/workDoneProgress/create")

type WorkDoneProgressCreateParams struct {
	/**
	 * The token to be used to report progress.
	 */
	Token ProgressToken `json:"token"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#window_workDoneProgress_cancel

const MethodWindowWorkDoneProgressCancel = Method("window/workDoneProgress/cancel")

type WindowWorkDoneProgressCancelFunc func(context *glsp.Context, params *WorkDoneProgressCancelParams) error

type WorkDoneProgressCancelParams struct {
	/**
	 * The token to be used to report progress.
	 */
	Token ProgressToken `json:"token"`
}
