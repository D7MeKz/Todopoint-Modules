package logging_test

import (
	"github.com/stretchr/testify/assert"
	"modules/v2/common/logging"
	"modules/v2/common/testutils"
	"testing"
)

func TestSetup(t *testing.T) {
	logger := logging.NewAppLogger()

	assert.NotNil(t, logger)
}

func TestLogging(t *testing.T) {
	logger := logging.NewAppLogger()
	ctx := testutils.GetTestGinContext()
	logger.Debug(ctx, 200, "Test Info")
}
