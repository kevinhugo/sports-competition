package helpers

import (
	"math/rand"
	"sports-competition/app/consts"
	"strings"

	"github.com/gofrs/uuid"
)

/*
uuid.NewV4() wrapper with additional paramtere to remove strip on first parameter
*/
func GetUUIDString(additionalParam ...interface{}) string {
	uuid, _ := uuid.NewV4()
	var removeStrip bool = false
	if len(additionalParam) > 0 {
		removeStrip = additionalParam[0].(bool)
	}
	if removeStrip {
		return strings.ReplaceAll(uuid.String(), "-", "")
	} else {
		return uuid.String()
	}
}

/*
Returns UUID, but with fully uppercased UUID, shorthand of GetUUIDString(true)
*/
func GetUppercasedUUIDString() string {
	return strings.ToUpper(GetUUIDString(true))
}

/*
Returns unique random string

First param assumed length of random string (int)

Second param assumed if use character as charset (bool)

Third param assumed if use digit as charset (bool)

Fourth param assumed if use lowercased character, ignored if second param is false (bool)

Fifth param assumed if use uppercased character, ignored if second param is false (bool)
*/
func GetRandomString(additionalConfig ...interface{}) string {

	var length int = 10
	var useChars bool = true
	var useDigits bool = true
	var useLowercaseChars bool = true
	var useUppercaseChars bool = true

	var additionalConfigLength int = len(additionalConfig)
	if additionalConfigLength > 0 {
		length = additionalConfig[0].(int)
	}
	if additionalConfigLength > 1 {
		useChars = additionalConfig[1].(bool)
	}
	if additionalConfigLength > 2 {
		useDigits = additionalConfig[2].(bool)
	}
	if additionalConfigLength > 3 {
		useLowercaseChars = additionalConfig[3].(bool)
	}
	if additionalConfigLength > 4 {
		useUppercaseChars = additionalConfig[4].(bool)
	}

	var randomCharset string = ""
	if useChars {
		if useLowercaseChars {
			randomCharset += consts.ASCII_LOWER_CHARS
		}
		if useUppercaseChars {
			randomCharset += consts.ASCII_UPPER_CHARS
		}
	}
	if useDigits {
		randomCharset += consts.ASCII_DIGITS
	}
	var charsetLength int = len(randomCharset)
	var randomStringResult string = ""
	for i := 0; i < length; i++ {
		randomStringResult += string(randomCharset[rand.Intn(charsetLength)])
	}

	return randomStringResult
}

func GetUniqueString(additionalData ...interface{}) string {
	var shortMode bool = true

	var additionalDataLength int = len(additionalData)
	if additionalDataLength > 0 {
		shortMode = additionalData[0].(bool)
	}

	var uuid string = GetUppercasedUUIDString()
	var timestamp string = GetTimestampUnixNanoNow()

	var uuidLength int = len(uuid)
	var timestampLength = len(timestamp)

	var maxCharLength int
	if shortMode {
		if uuidLength < timestampLength {
			maxCharLength = uuidLength
		} else {
			maxCharLength = timestampLength
		}
	} else {
		if uuidLength > timestampLength {
			maxCharLength = uuidLength
		} else {
			maxCharLength = timestampLength
		}
	}

	var uniqueString string
	for i := 0; i < maxCharLength; i++ {
		uniqueString += string(uuid[i]) + string(timestamp[i])
	}

	return uniqueString
}

/*
Returns unique random string, but only numbers
First param is assumed length of random string (int)
*/

func GetRandomNumberString(additionalConfig ...interface{}) string {
	// Random generator config
	var length int = 6
	// \ Random generator config

	// Override settings based on additionalConfig param
	var additionalConfigLength int = len(additionalConfig)
	if additionalConfigLength > 0 {
		length = additionalConfig[0].(int)
	}

	return GetRandomString(length, false, true)
	// \ Override settings based ona dditionalConfig param

}
