// Package plugin defines the protocol layer for communication between notation
// and notation external plugin.
package plugin

// Capability is a feature available in the plugin contract.
type Capability string

const (
	// CapabilitySignatureGenerator is the name of the capability
	// for a plugin to support generating raw signatures.
	CapabilitySignatureGenerator Capability = "SIGNATURE_GENERATOR.RAW"

	// CapabilityEnvelopeGenerator is the name of the capability
	// for a plugin to support generating envelope signatures.
	CapabilityEnvelopeGenerator Capability = "SIGNATURE_GENERATOR.ENVELOPE"

	// CapabilityTrustedIdentityVerifier is the name of the
	// capability for a plugin to support verifying trusted identities.
	CapabilityTrustedIdentityVerifier Capability = "SIGNATURE_VERIFIER.TRUSTED_IDENTITY"

	// CapabilityRevocationCheckVerifier is the name of the
	// capability for a plugin to support verifying revocation checks.
	CapabilityRevocationCheckVerifier Capability = "SIGNATURE_VERIFIER.REVOCATION_CHECK"
)

// Command is a CLI command available in the plugin contract.
type Command string

const (
	// CommandGetMetadata is the name of the plugin command
	// which must be supported by every plugin and returns the
	// plugin metadata.
	CommandGetMetadata Command = "get-plugin-metadata"

	// CommandDescribeKey is the name of the plugin command
	// which must be supported by every plugin that has the
	// SIGNATURE_GENERATOR.RAW capability.
	CommandDescribeKey Command = "describe-key"

	// CommandGenerateSignature is the name of the plugin command
	// which must be supported by every plugin that has the
	// SIGNATURE_GENERATOR.RAW capability.
	CommandGenerateSignature Command = "generate-signature"

	// CommandGenerateEnvelope is the name of the plugin command
	// which must be supported by every plugin that has the
	// SIGNATURE_GENERATOR.ENVELOPE capability.
	CommandGenerateEnvelope Command = "generate-envelope"

	// CommandVerifySignature is the name of the plugin command
	// which must be supported by every plugin that has
	// any SIGNATURE_VERIFIER.* capability
	CommandVerifySignature Command = "verify-signature"

	Version Command = "version"
)
