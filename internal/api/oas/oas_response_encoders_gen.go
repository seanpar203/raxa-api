// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	ht "github.com/ogen-go/ogen/http"
)

func encodeV1AuthLoginResponse(response V1AuthLoginRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *V1AuthLoginResponse:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *V1ErrorResponseStatusCode:
		w.Header().Set("Content-Type", "application/json")
		code := response.StatusCode
		if code == 0 {
			// Set default status code.
			code = http.StatusOK
		}
		w.WriteHeader(code)
		st := http.StatusText(code)
		if code >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}

		e := jx.GetEncoder()
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		if code >= http.StatusInternalServerError {
			return errors.Wrapf(ht.ErrInternalServerErrorResponse, "code: %d, message: %s", code, http.StatusText(code))
		}
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeV1AuthRefreshResponse(response V1AuthRefreshRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *V1AuthRefreshResponse:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *V1ErrorResponseStatusCode:
		w.Header().Set("Content-Type", "application/json")
		code := response.StatusCode
		if code == 0 {
			// Set default status code.
			code = http.StatusOK
		}
		w.WriteHeader(code)
		st := http.StatusText(code)
		if code >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}

		e := jx.GetEncoder()
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		if code >= http.StatusInternalServerError {
			return errors.Wrapf(ht.ErrInternalServerErrorResponse, "code: %d, message: %s", code, http.StatusText(code))
		}
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeV1OTPCodeEnterResponse(response V1OTPCodeEnterRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *V1OTPCodeEnterNoContent:
		w.WriteHeader(204)
		span.SetStatus(codes.Ok, http.StatusText(204))

		return nil

	case *V1ErrorResponseStatusCode:
		w.Header().Set("Content-Type", "application/json")
		code := response.StatusCode
		if code == 0 {
			// Set default status code.
			code = http.StatusOK
		}
		w.WriteHeader(code)
		st := http.StatusText(code)
		if code >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}

		e := jx.GetEncoder()
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		if code >= http.StatusInternalServerError {
			return errors.Wrapf(ht.ErrInternalServerErrorResponse, "code: %d, message: %s", code, http.StatusText(code))
		}
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeV1OTPCodeSendResponse(response V1OTPCodeSendRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *V1OTPCodeSendNoContent:
		w.WriteHeader(204)
		span.SetStatus(codes.Ok, http.StatusText(204))

		return nil

	case *V1ErrorResponseStatusCode:
		w.Header().Set("Content-Type", "application/json")
		code := response.StatusCode
		if code == 0 {
			// Set default status code.
			code = http.StatusOK
		}
		w.WriteHeader(code)
		st := http.StatusText(code)
		if code >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}

		e := jx.GetEncoder()
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		if code >= http.StatusInternalServerError {
			return errors.Wrapf(ht.ErrInternalServerErrorResponse, "code: %d, message: %s", code, http.StatusText(code))
		}
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeV1UsersCreateResponse(response V1UsersCreateRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *V1AuthLoginResponse:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *V1ErrorResponseStatusCode:
		w.Header().Set("Content-Type", "application/json")
		code := response.StatusCode
		if code == 0 {
			// Set default status code.
			code = http.StatusOK
		}
		w.WriteHeader(code)
		st := http.StatusText(code)
		if code >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}

		e := jx.GetEncoder()
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		if code >= http.StatusInternalServerError {
			return errors.Wrapf(ht.ErrInternalServerErrorResponse, "code: %d, message: %s", code, http.StatusText(code))
		}
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeV1UsersMeResponse(response V1UsersMeRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *V1User:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *V1401UnauthorizedResponse:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(401)
		span.SetStatus(codes.Error, http.StatusText(401))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *V1ErrorResponseStatusCode:
		w.Header().Set("Content-Type", "application/json")
		code := response.StatusCode
		if code == 0 {
			// Set default status code.
			code = http.StatusOK
		}
		w.WriteHeader(code)
		st := http.StatusText(code)
		if code >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}

		e := jx.GetEncoder()
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		if code >= http.StatusInternalServerError {
			return errors.Wrapf(ht.ErrInternalServerErrorResponse, "code: %d, message: %s", code, http.StatusText(code))
		}
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeV1UsersMeContactsCreateResponse(response V1UsersMeContactsCreateRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *V1UsersMeContactsCreateNoContent:
		w.WriteHeader(204)
		span.SetStatus(codes.Ok, http.StatusText(204))

		return nil

	case *V1ErrorResponseStatusCode:
		w.Header().Set("Content-Type", "application/json")
		code := response.StatusCode
		if code == 0 {
			// Set default status code.
			code = http.StatusOK
		}
		w.WriteHeader(code)
		st := http.StatusText(code)
		if code >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}

		e := jx.GetEncoder()
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		if code >= http.StatusInternalServerError {
			return errors.Wrapf(ht.ErrInternalServerErrorResponse, "code: %d, message: %s", code, http.StatusText(code))
		}
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeV1UsersMeUpdateResponse(response V1UsersMeUpdateRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *V1User:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *V1ErrorResponseStatusCode:
		w.Header().Set("Content-Type", "application/json")
		code := response.StatusCode
		if code == 0 {
			// Set default status code.
			code = http.StatusOK
		}
		w.WriteHeader(code)
		st := http.StatusText(code)
		if code >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}

		e := jx.GetEncoder()
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		if code >= http.StatusInternalServerError {
			return errors.Wrapf(ht.ErrInternalServerErrorResponse, "code: %d, message: %s", code, http.StatusText(code))
		}
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeErrorResponse(response *V1ErrorResponseStatusCode, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	code := response.StatusCode
	if code == 0 {
		// Set default status code.
		code = http.StatusOK
	}
	w.WriteHeader(code)
	st := http.StatusText(code)
	if code >= http.StatusBadRequest {
		span.SetStatus(codes.Error, st)
	} else {
		span.SetStatus(codes.Ok, st)
	}

	e := jx.GetEncoder()
	response.Response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	if code >= http.StatusInternalServerError {
		return errors.Wrapf(ht.ErrInternalServerErrorResponse, "code: %d, message: %s", code, http.StatusText(code))
	}
	return nil

}
