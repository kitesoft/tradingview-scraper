package tradingview

// InitErrorContext ...
const InitErrorContext = "Initializing the connection"

// ReadFirstMessageErrorContext
const ReadFirstMessageErrorContext = "Reading the first message after stablishing the connection"

// DecodeFirstMessageErrorContext ...
const DecodeFirstMessageErrorContext = "Decoding the first message after stablishing the connection"

// FirstMessageWithoutSessionIdErrorContext ...
const FirstMessageWithoutSessionIdErrorContext = "Does not have 'session_id' property"

// ConnectionSetupMessagesErrorContext ...
const ConnectionSetupMessagesErrorContext = "Sending connection setup messages"

// SendMessageErrorContext ...
const SendMessageErrorContext = "Sending a message"

// SendKeepAliveMessageErrorContext ...
const SendKeepAliveMessageErrorContext = "Sending the keep alive message"

// GetPayloadLengthErrorContext ...
const GetPayloadLengthErrorContext = "Getting the payload length"

// DecodeMessageErrorContext ...
const DecodeMessageErrorContext = "Decoding the JSON message"

// DecodedMessageHasErrorPropertyErrorContext ...
const DecodedMessageHasErrorPropertyErrorContext = "JSON message has an error message"

// DecodedMessageDoesNotIncludePayloadErrorContext ...
const DecodedMessageDoesNotIncludePayloadErrorContext = "JSON message does not include the payload"

// PayloadCantBeParsedErrorContext ...
const PayloadCantBeParsedErrorContext = "JSON payload couldn't be parsed"

// FinalPayloadCantBeParsedErrorContext ...
const FinalPayloadCantBeParsedErrorContext = "The final JSON payload of the socket message couldn't be parsed"

// FinalPayloadHasMissingPropertiesErrorContext ...
const FinalPayloadHasMissingPropertiesErrorContext = "The final JSON payload doesn't have the expected data"

// ReadMessageErrorContext ...
const ReadMessageErrorContext = "Error while reading new messages through the socket connection"

