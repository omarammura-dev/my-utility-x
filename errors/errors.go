// errors/errors.go
package errors

import "net/http"

type APIError struct {
    Code    string `json:"code"`
    Status  int    `json:"status"`
    Message string `json:"message"`
}

var (
    ErrBindingUserData = APIError{
        Code:    "BINDING_USER_DATA_ERROR",
        Status:  http.StatusInternalServerError,
        Message: "Error while binding the user data",
    }
    ErrSavingUser = APIError{
        Code:    "SAVING_USER_ERROR",
        Status:  http.StatusInternalServerError,
        Message: "Could not register the user",
    }
    ErrGeneratingToken = APIError{
        Code:    "TOKEN_GENERATION_ERROR",
        Status:  http.StatusInternalServerError,
        Message: "Could not generate the token",
    }
    ErrSendingVerificationEmail = APIError{
        Code:    "EMAIL_VERIFICATION_ERROR",
        Status:  http.StatusInternalServerError,
        Message: "Could not send verification email",
    }
    ErrIncompleteRequest = APIError{
        Code:    "INCOMPLETE_REQUEST_ERROR",
        Status:  http.StatusInternalServerError,
        Message: "Incomplete request!",
    }
    ErrIncorrectOrExpiredToken = APIError{
        Code:    "TOKEN_VERIFICATION_ERROR",
        Status:  http.StatusInternalServerError,
        Message: "Incorrect or expired token!",
    }
    ErrUpdatingUser = APIError{
        Code:    "UPDATING_USER_ERROR",
        Status:  http.StatusInternalServerError,
        Message: "Oops!",
    }
    ErrEmptyEmail = APIError{
        Code:    "EMPTY_EMAIL_ERROR",
        Status:  http.StatusInternalServerError,
        Message: "Email is empty!",
    }
    ErrFindingByEmail = APIError{
        Code:    "FINDING_BY_EMAIL_ERROR",
        Status:  http.StatusInternalServerError,
        Message: "There seems to be a discrepancy with the information provided. To ensure account security, we can't assist with password resets for unrecognized information.",
    }
    ErrSendingResetPasswordEmail = APIError{
        Code:    "RESET_PASSWORD_EMAIL_ERROR",
        Status:  http.StatusInternalServerError,
        Message: "Email could not send!",
    }
    ErrParsing = APIError{
        Code:    "PARSING_ERROR",
        Status:  http.StatusInternalServerError,
        Message: "Could not parse!",
    }
    ErrVerifyingAndUpdatePassword = APIError{
        Code:    "VERIFYING_AND_UPDATE_PASSWORD_ERROR",
        Status:  http.StatusUnauthorized,
        Message: "Please double check your password!",
    }
    ErrHashingPassword = APIError{
        Code:    "HASHING_PASSWORD_ERROR",
        Status:  http.StatusOK,
        Message: "Error!",
    }
	ErrSomethingWentWrong = APIError{
        Code:    "INTERNAL_SERVER_ERROR",
        Status:  http.StatusInternalServerError,
        Message: "Error!",
    }
	ErrUnAuthorized = APIError{
        Code:    "UNAUTHORIZED",
        Status:  http.StatusUnauthorized,
        Message: "Error!",
    }
    ErrInvalidExpenseType = APIError{
        Code:    "INVALID_EXPENSE_TYPE",
        Status:  http.StatusInternalServerError,
        Message: "Error!",
    }
)