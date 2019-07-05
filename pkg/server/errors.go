package server

import (
	"net/http"
)

/*
Standard Errors in Matrix errcode format. */
const CodeBadRequest = "M_BAD_REQUEST"                                       // 400
const CodeUnauthorized = "M_UNAUTHORIZED"                                    // 401
const CodePaymentRequired = "M_PAYMENT_REQUIRED"                             // 402
const CodeForbidden = "M_FORBIDDEN"                                          // 403
const CodeNotFound = "M_NOT_FOUND"                                           // 404
const CodeMethodNotAllowed = "M_METHOD_NOT_ALLOWED"                          // 405
const CodeNotAcceptable = "M_NOT_ACCEPTABLE"                                 // 406
const CodeProxyAuthRequired = "M_PROXY_AUTH_REQUIRED"                        // 407
const CodeRequestTimeout = "M_REQUEST_TIMEOUT"                               // 408
const CodeConflict = "M_CONFLICT"                                            // 409
const CodeGone = "M_GONE"                                                    // 410
const CodeLengthRequired = "M_LENGTH_REQUIRED"                               // 411
const CodePreconditionFailed = "M_PRECONDITION_FAILED"                       // 412
const CodeRequestEntityTooLarge = "M_REQUEST_ENTITY_TOO_LARGE"               // 413
const CodeRequestURITooLong = "M_REQUEST_URI_TOO_LONG"                       // 414
const CodeUnsupportedMediaType = "M_UNSUPPORTED_MEDIA_TYPE"                  // 415
const CodeRequestedRangeNotSatisfiable = "M_REQUESTED_RANGE_NOT_SATISFIABLE" // 416
const CodeExpectationFailed = "M_EXPECTATION_FAILED"                         // 417
const CodeTeapot = "M_TEAPOT"                                                // 418
const CodeMisdirectedRequest = "M_MISDIRECTED_REQUEST"                       // 421
const CodeUnprocessableEntity = "M_UNPROCESSABLE_ENTITY"                     // 422
const CodeLocked = "M_LOCKED"                                                // 423
const CodeFailedDependency = "M_FAILED_DEPENDENCY"                           // 424
const CodeTooEarly = "M_TOO_EARLY"                                           // 425
const CodeUpgradeRequired = "M_UPGRADE_REQUIRED"                             // 426
const CodePreconditionRequired = "M_PRECONDITION_REQUIRED"                   // 428
const CodeTooManyRequests = "M_TOO_MANY_REQUESTS"                            // 429
const CodeRequestHeaderFieldsTooLarge = "M_REQUEST_HEADER_FIELDS_TOO_LARGE"  // 431
const CodeUnavailableForLegalReasons = "M_UNAVAILABLE_FOR_LEGAL_REASONS"     // 451

const CodeInternalServerError = "M_INTERNAL_SERVER_ERROR"                     // 500
const CodeNotImplemented = "M_NOT_IMPLEMENTED"                                // 501
const CodeBadGateway = "M_BAD_GATEWAY"                                        // 502
const CodeServiceUnavailable = "M_SERVICE_UNAVAILABLE"                        // 503
const CodeGatewayTimeout = "M_GATEWAY_TIMEOUT"                                // 504
const CodeHTTPVersionNotSupported = "M_HTTP_VERSION_NOT_SUPPORTED"            // 505
const CodeVariantAlsoNegotiates = "M_VARIANT_ALSO_NEGOTIATES"                 // 506
const CodeInsufficientStorage = "M_INSUFFICIENT_STORAGE"                      // 507
const CodeLoopDetected = "M_LOOP_DETECTED"                                    // 508
const CodeNotExtended = "M_NOT_EXTENDED"                                      // 510
const CodeNetworkAuthenticationRequired = "M_NETWORK_AUTHENTICATION_REQUIRED" // 511

/*
Error provides the basic JSON fields that are expected by the Matrix
specification, and Extra for additional details.  */
type Error struct {
	Extra   interface{} `json:"extra,omitempty"`
	Code    string      `json:"errcode"`
	Message string      `json:"error"`
}

/*
Errors is a mapping of net/http status codes to Error*'s. */
var Errors = map[int]Error{
	http.StatusBadRequest:                   ErrorBadRequest,
	http.StatusUnauthorized:                 ErrorUnauthorized,
	http.StatusPaymentRequired:              ErrorPaymentRequired,
	http.StatusForbidden:                    ErrorForbidden,
	http.StatusNotFound:                     ErrorNotFound,
	http.StatusMethodNotAllowed:             ErrorMethodNotAllowed,
	http.StatusNotAcceptable:                ErrorNotAcceptable,
	http.StatusProxyAuthRequired:            ErrorProxyAuthRequired,
	http.StatusRequestTimeout:               ErrorRequestTimeout,
	http.StatusConflict:                     ErrorConflict,
	http.StatusGone:                         ErrorGone,
	http.StatusLengthRequired:               ErrorLengthRequired,
	http.StatusPreconditionFailed:           ErrorPreconditionFailed,
	http.StatusRequestEntityTooLarge:        ErrorRequestEntityTooLarge,
	http.StatusRequestURITooLong:            ErrorRequestURITooLong,
	http.StatusUnsupportedMediaType:         ErrorUnsupportedMediaType,
	http.StatusRequestedRangeNotSatisfiable: ErrorRequestedRangeNotSatisfiable,
	http.StatusExpectationFailed:            ErrorExpectationFailed,
	http.StatusTeapot:                       ErrorTeapot,
	http.StatusMisdirectedRequest:           ErrorMisdirectedRequest,
	http.StatusUnprocessableEntity:          ErrorUnprocessableEntity,
	http.StatusLocked:                       ErrorLocked,
	http.StatusFailedDependency:             ErrorFailedDependency,
	http.StatusTooEarly:                     ErrorTooEarly,
	http.StatusRequestHeaderFieldsTooLarge:  ErrorRequestHeaderFieldsTooLarge,
	http.StatusUnavailableForLegalReasons:   ErrorUnavailableForLegalReasons,

	// 500s
	http.StatusInternalServerError:           ErrorInternalServerError,
	http.StatusNotImplemented:                ErrorNotImplemented,
	http.StatusBadGateway:                    ErrorBadGateway,
	http.StatusServiceUnavailable:            ErrorServiceUnavailable,
	http.StatusGatewayTimeout:                ErrorGatewayTimeout,
	http.StatusHTTPVersionNotSupported:       ErrorHTTPVersionNotSupported,
	http.StatusVariantAlsoNegotiates:         ErrorVariantAlsoNegotiates,
	http.StatusInsufficientStorage:           ErrorInsufficientStorage,
	http.StatusLoopDetected:                  ErrorLoopDetected,
	http.StatusNotExtended:                   ErrorNotExtended,
	http.StatusNetworkAuthenticationRequired: ErrorNetworkAuthenticationRequired,
}

/*
Error* are net/http status codes to Error objects. */
var ErrorBadRequest = Error{Message: "Bad Request", Code: CodeBadRequest}
var ErrorUnauthorized = Error{Message: "Unauthorized", Code: CodeUnauthorized}
var ErrorPaymentRequired = Error{Message: "Payment Required", Code: CodePaymentRequired}
var ErrorForbidden = Error{Message: "Forbidden", Code: CodeForbidden}
var ErrorNotFound = Error{Message: "Not Found", Code: CodeNotFound}
var ErrorMethodNotAllowed = Error{Message: "Method Not Allowed", Code: CodeMethodNotAllowed}
var ErrorNotAcceptable = Error{Message: "Not Acceptable", Code: CodeNotAcceptable}
var ErrorProxyAuthRequired = Error{Message: "Proxy Auth Required", Code: CodeProxyAuthRequired}
var ErrorRequestTimeout = Error{Message: "Request Timeout", Code: CodeRequestTimeout}
var ErrorConflict = Error{Message: "Conflict", Code: CodeConflict}
var ErrorGone = Error{Message: "Gone", Code: CodeGone}
var ErrorLengthRequired = Error{Message: "Length Required", Code: CodeLengthRequired}
var ErrorPreconditionFailed = Error{Message: "Precondition Failed", Code: CodePreconditionFailed}
var ErrorRequestEntityTooLarge = Error{Message: "Request Entity Too Large", Code: CodeRequestEntityTooLarge}
var ErrorRequestURITooLong = Error{Message: "Request URI Too Long", Code: CodeRequestURITooLong}
var ErrorUnsupportedMediaType = Error{Message: "Unsupported Media Type", Code: CodeUnsupportedMediaType}
var ErrorRequestedRangeNotSatisfiable = Error{Message: "Requested Range Not Satisfiable", Code: CodeRequestedRangeNotSatisfiable}
var ErrorExpectationFailed = Error{Message: "Expectation Failed", Code: CodeExpectationFailed}
var ErrorTeapot = Error{Message: "Teapot", Code: CodeTeapot}
var ErrorMisdirectedRequest = Error{Message: "Misdirected Request", Code: CodeMisdirectedRequest}
var ErrorUnprocessableEntity = Error{Message: "Unprocessable Entity", Code: CodeUnprocessableEntity}
var ErrorLocked = Error{Message: "Locked", Code: CodeLocked}
var ErrorFailedDependency = Error{Message: "Failed Dependency", Code: CodeFailedDependency}
var ErrorTooEarly = Error{Message: "Too Early", Code: CodeTooEarly}
var ErrorUpgradeRequired = Error{Message: "Upgrade Required", Code: CodeUpgradeRequired}
var ErrorPreconditionRequired = Error{Message: "Precondition Required", Code: CodePreconditionRequired}
var ErrorTooManyRequests = Error{Message: "Too Many Requests", Code: CodeTooManyRequests}
var ErrorRequestHeaderFieldsTooLarge = Error{Message: "Request Header Fields Too Large", Code: CodeRequestHeaderFieldsTooLarge}
var ErrorUnavailableForLegalReasons = Error{Message: "Unavailable For Legal Reasons", Code: CodeUnavailableForLegalReasons}

var ErrorInternalServerError = Error{Message: "Internal Server Error", Code: CodeInternalServerError}
var ErrorNotImplemented = Error{Message: "Not Implemented", Code: CodeNotImplemented}
var ErrorBadGateway = Error{Message: "Bad Gateway", Code: CodeBadGateway}
var ErrorServiceUnavailable = Error{Message: "Service Unavailable", Code: CodeServiceUnavailable}
var ErrorGatewayTimeout = Error{Message: "Gateway Timeout", Code: CodeGatewayTimeout}
var ErrorHTTPVersionNotSupported = Error{Message: "HTTP Version Not Supported", Code: CodeHTTPVersionNotSupported}
var ErrorVariantAlsoNegotiates = Error{Message: "Variant Also Negotiates", Code: CodeVariantAlsoNegotiates}
var ErrorInsufficientStorage = Error{Message: "Insufficient Storage", Code: CodeInsufficientStorage}
var ErrorLoopDetected = Error{Message: "Loop Detected", Code: CodeLoopDetected}
var ErrorNotExtended = Error{Message: "Not Extended", Code: CodeNotExtended}
var ErrorNetworkAuthenticationRequired = Error{Message: "Network Authentication Required", Code: CodeNetworkAuthenticationRequired}
