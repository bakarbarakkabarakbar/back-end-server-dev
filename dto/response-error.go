package dto

type ErrorResponse struct {
	ResponseMeta
	Data   any `json:"data"`
	Errors any `json:"errors,omitempty"`
}

func DefaultErrorResponse() ErrorResponse {
	return DefaultErrorResponseWithMessage("")
}

func DefaultErrorResponseWithMessage(msg string) ErrorResponse {
	return ErrorResponse{
		ResponseMeta: ResponseMeta{
			Success:      false,
			MessageTitle: "Oops, something went wrong.",
			Message:      msg,
			ResponseTime: "",
		},
		Data: nil,
	}
}

func DefaultErrorWithResponse(response ResponseMeta) ErrorResponse {
	return ErrorResponse{
		ResponseMeta: response,
		Data:         nil,
	}
}

func DefaultErrorInvalidDataWithMessage(title string, msg string) ErrorResponse {
	return ErrorResponse{
		ResponseMeta: ResponseMeta{
			Success:      false,
			MessageTitle: title,
			Message:      msg,
			ResponseTime: "",
		},
		Data: nil,
	}
}

func DefaultDataInvalidResponse(validationErrors any) ErrorResponse {
	return ErrorResponse{
		ResponseMeta: ResponseMeta{
			MessageTitle: "Oops, something went wrong.",
			Message:      "Data invalid.",
		},
		Errors: validationErrors,
	}
}

func DefaultBadRequestResponse() ErrorResponse {
	return DefaultErrorResponseWithMessage("Bad request")
}
