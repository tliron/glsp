package protocol

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#client_registerCapability

/**
 * General parameters to register for a capability.
 */
type Registration struct {
	/**
	 * The id used to register the request. The id can be used to deregister
	 * the request again.
	 */
	ID string `json:"id"`

	/**
	 * The method / capability to register for.
	 */
	Method string `json:"method"`

	/**
	 * Options necessary for the registration.
	 */
	RegisterOptions any `json:"registerOptions,omitempty"`
}

const ServerClientRegisterCapability = Method("client/registerCapability")

type RegistrationParams struct {
	Registrations []Registration `json:"registrations"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#client_unregisterCapability

/**
 * General parameters to unregister a capability.
 */
type Unregistration struct {
	/**
	 * The id used to unregister the request or notification. Usually an id
	 * provided during the register request.
	 */
	ID string `json:"id"`

	/**
	 * The method / capability to unregister for.
	 */
	Method string `json:"method"`
}

const ServerClientUnregisterCapability = Method("client/unregisterCapability")

type UnregistrationParams struct {
	// This should correctly be named `unregistrations`. However changing this
	// is a breaking change and needs to wait until we deliver a 4.x version
	// of the specification.
	Unregisterations []Unregistration `json:"unregisterations"`
}
