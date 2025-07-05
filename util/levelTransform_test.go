package util_test

import (
	"testing"

	"github.com/Alonza0314/logger-go/util"
	"github.com/stretchr/testify/assert"
)

var testLevelTransformCases = []struct {
	name        string
	levelString util.LogLevelString
	level       util.LogLevel
	wantPanic   bool
}{
	{
		name:        "error",
		levelString: util.LEVEL_STRING_ERROR,
		level:       util.LEVEL_ERROR,
		wantPanic:   false,
	},
	{
		name:        "warn",
		levelString: util.LEVEL_STRING_WARN,
		level:       util.LEVEL_WARN,
		wantPanic:   false,
	},
	{
		name:        "info",
		levelString: util.LEVEL_STRING_INFO,
		level:       util.LEVEL_INFO,
		wantPanic:   false,
	},
	{
		name:        "debug",
		levelString: util.LEVEL_STRING_DEBUG,
		level:       util.LEVEL_DEBUG,
		wantPanic:   false,
	},
	{
		name:        "trace",
		levelString: util.LEVEL_STRING_TRACE,
		level:       util.LEVEL_TRACE,
		wantPanic:   false,
	},
	{
		name:        "test",
		levelString: util.LEVEL_STRING_TEST,
		level:       util.LEVEL_TEST,
		wantPanic:   false,
	},
	{
		name:        "invalid",
		levelString: "invalid",
		level:       -1,
		wantPanic:   true,
	},
}

func TestLevelStringToLevel(t *testing.T) {
	for _, testCase := range testLevelTransformCases {
		t.Run(testCase.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != testCase.wantPanic {
					t.Errorf("LevelStringToLevel() panic = %v, wantPanic = %v", r != nil, testCase.wantPanic)
				}
			}()
			level := util.LevelStringToLevel(testCase.levelString)
			if !testCase.wantPanic {
				assert.Equal(t, testCase.level, level)
			}
		})
	}
}
