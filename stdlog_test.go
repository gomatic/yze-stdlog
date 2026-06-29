package stdlog_test

import (
	"testing"

	stdlog "github.com/gomatic/yze-go-stdlog"
	"github.com/stretchr/testify/assert"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestStandardLogImportIsReported(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), stdlog.Analyzer, "a")
}

func TestRegistrationIsWellFormed(t *testing.T) {
	assert.NoError(t, stdlog.Registration.Validate())
	assert.Equal(t, "yze/go/stdlog", stdlog.Registration.RuleID())
	assert.Same(t, stdlog.Analyzer, stdlog.Registration.Analyzer)
}
