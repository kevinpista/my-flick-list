package error_constants

// Error constants to call upon to keep consistent with frontend error constants
// which will be used to validate the backend thrown error which will then trigger a specific frontend action
const (
	EmailExists        = "email_already_exists"
	InvalidEmail       = "invalid_email"
	PasswordWhitespace = "password_whitespace"
	PasswordEmpty      = "password_empty"
	InvalidName        = "invalid_name" // empty or whitespace-only name
	BadRequest         = "bad_request"  // general
	Server             = "server_error"
)
