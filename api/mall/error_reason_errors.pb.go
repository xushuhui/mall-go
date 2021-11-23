// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package mall

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsNotFound(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_Not_Found.String() && e.Code == 404
}

func ErrorNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, ErrorReason_Not_Found.String(), fmt.Sprintf(format, args...))
}

func IsUnauthorized(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_Unauthorized.String() && e.Code == 401
}

func ErrorUnauthorized(format string, args ...interface{}) *errors.Error {
	return errors.New(401, ErrorReason_Unauthorized.String(), fmt.Sprintf(format, args...))
}

func IsForbidden(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_Forbidden.String() && e.Code == 403
}

func ErrorForbidden(format string, args ...interface{}) *errors.Error {
	return errors.New(403, ErrorReason_Forbidden.String(), fmt.Sprintf(format, args...))
}

func IsInternalServerError(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_Internal_Server_Error.String() && e.Code == 500
}

func ErrorInternalServerError(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_Internal_Server_Error.String(), fmt.Sprintf(format, args...))
}

func IsInvalidParams(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_Invalid_Params.String() && e.Code == 400
}

func ErrorInvalidParams(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_Invalid_Params.String(), fmt.Sprintf(format, args...))
}

func IsInvalidToken(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_Invalid_Token.String() && e.Code == 400
}

func ErrorInvalidToken(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_Invalid_Token.String(), fmt.Sprintf(format, args...))
}

func IsLoginFail(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_Login_Fail.String() && e.Code == 400
}

func ErrorLoginFail(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_Login_Fail.String(), fmt.Sprintf(format, args...))
}