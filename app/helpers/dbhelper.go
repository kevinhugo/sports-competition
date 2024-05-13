package helpers

import (
	"fmt"
	"strings"
)

/*
Search query builder, simplified mutliple search query builder

First param of additionalData is presumed bool type and decide if search is by case insensitive or not ( default is false, set to true for case insensitive )
*/
func BuildSearchString(searchByList []string, search string, additionalData ...interface{}) (string, []interface{}) {
	var additionalDataLength = len(additionalData)
	var toUpper bool = false
	if additionalDataLength > 0 {
		toUpper = additionalData[0].(bool)
	}
	var formattedSearchList []string
	for _, each := range searchByList {
		if toUpper {
			formattedSearchList = append(formattedSearchList, fmt.Sprintf("UPPER(%s) LIKE ?", each))
		} else {
			formattedSearchList = append(formattedSearchList, fmt.Sprintf("%s LIKE ?", each))
		}
	}

	var searchArray []interface{}
	for i := 0; i < len(searchByList); i++ {
		if toUpper {
			searchArray = append(searchArray, strings.ToUpper(search))
		} else {
			searchArray = append(searchArray, search)
		}
	}

	return fmt.Sprintf("( %s )", strings.Join(formattedSearchList, " OR ")), searchArray
}
