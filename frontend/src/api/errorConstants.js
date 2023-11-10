// Error constants to better map with backend returned errors 
// along with translating to allow components to display a message or perform an action based on the error
// This is more so for our component, so that it is doing checks for the constant name versus strings

// USER REGISTRATION ERRORS
export const ERROR_EMAIL_EXISTS = "email_already_exists";
export const ERROR_INVALID_EMAIL = "invalid_email";
export const ERROR_PASSWORD_WHITESPACE = "password_whitespace";
export const ERROR_PASSWORD_EMPTY = "password_empty";
export const ERROR_INVALID_NAME = "invalid_name";
export const ERROR_BAD_REQUEST = "bad_request"; // general
export const ERROR_SERVER = "server_error";
export const ERROR_INVALID_LOGIN = "invalid_login";

