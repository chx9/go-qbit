package qbit

import (
	wrapper "github.com/pkg/errors"
)

var NotLogin = wrapper.New("Please Login First")

func FailedToBuildRequest(err error) error {
	return wrapper.Wrap(err, "request Build Failed")
}

func FailedToDecodeResponse(err error) error {
	return wrapper.Wrap(err, "failed to decode response")
}

func RequestFailed(err error) error {
	return wrapper.Wrap(err, "request Failed")
}
