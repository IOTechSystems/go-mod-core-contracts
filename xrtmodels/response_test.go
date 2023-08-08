package xrtmodels

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestXrtErrorCode(t *testing.T) {
	assert.Equal(t, XrtErrorCode(errors.NewCommonEdgeX(errors.KindEntityDoesNotExist, "", nil)), XrtSdkStatusNotFound)
	assert.Equal(t, XrtErrorCode(errors.NewCommonEdgeX(errors.KindNotImplemented, "", nil)), XrtSdkStatusNotSupported)
	assert.Equal(t, XrtErrorCode(errors.NewCommonEdgeX(errors.KindInvalidId, "", nil)), XrtSdkStatusInvalidOperation)
	assert.Equal(t, XrtErrorCode(errors.NewCommonEdgeX(errors.KindDuplicateName, "", nil)), XrtSdkStatusAlreadyExists)
	assert.Equal(t, XrtErrorCode(errors.NewCommonEdgeX(errors.KindServerError, "", nil)), XrtSdkStatusServerError)
}
