package utils

import (
    "net/http"
    "your-jwt-library"
)

// AuthenticateMiddleware is a middleware function to validate JWT token
func AuthenticateMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        tokenString := // extract token from request headers or cookies

        // Verify the token
        _, err := yourJWTLibrary.VerifyToken(tokenString)
        if err != nil {
            RespondWithError(w, http.StatusUnauthorized, "Invalid or expired token")
            return
        }

        // Proceed to the next handler if the token is valid
        next(w, r)
    }
}

// GetURLParam extracts a parameter from the URL path
func GetURLParam(r *http.Request, param string) string {
    // Extract and return the parameter value from the request URL path
}
