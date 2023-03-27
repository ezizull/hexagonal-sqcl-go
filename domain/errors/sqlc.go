// Package errors defines the domain errors used in the application.
package errors

// SqlcErr is a struct that contains the error number and message for Sqlc errors
type SqlcErr struct {
	Number  int    `json:"Number"`
	Message string `json:"Message"`
}
