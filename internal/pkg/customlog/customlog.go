package customlog

import (
	"runtime"

	"github.com/rs/zerolog/log"
)

func HandleError(activity string, err error) {
	if err != nil {
		// notice that we're using 1, so it will actually log where
		// the error happened, 0 = this function, we don't want that.
		_, fn, line, _ := runtime.Caller(1)
		log.Error().Msgf("%s [error] %s:%d %v", activity, fn, line, err)
	}
}

// this logs the function name as well.
func FancyHandleError(err error) (b bool) {
	if err != nil {
		// notice that we're using 1, so it will actually log the where
		// the error happened, 0 = this function, we don't want that.
		pc, fn, line, _ := runtime.Caller(1)

		log.Error().Msgf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, err)
		b = true
	}
	return
}
