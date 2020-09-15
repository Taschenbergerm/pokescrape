package pkg

import "github.com/taschenbergerm/pokescraper/log"

// handleError will log away any error that happened and if it should be strict it will panic
func handleError(err error, strict bool) {
	if err != nil {
		log.Errorln(err)
		if strict {
			panic(err)
		}
	}
}

// HandleErrorSoftly will log away any error that happened
func HandleErrorSoftly(err error) {
	handleError(err, false)
}

// HandleErrorStrictly will log away any error that happened and panic in that case
func HandleErrorStrictly(err error) {
	handleError(err, true)
}
