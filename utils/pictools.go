package utils

import (
	"bytes"
	"net/http"
	"strings"
)

func ConverArceeURLToWeedUrl(url string) (resultUrl string) {

	split := strings.Split(url, "v2")
	if len(split) == 1 {
		resultUrl = strings.Replace(url, "8501/api/file", "9333", -1)
	} else {
		temp := strings.Replace(url, "8501/api/v2/file", "9333", -1)
		s1 := temp[:strings.LastIndex(temp, "/")]
		s2 := temp[strings.LastIndex(temp, "/")+1:]
		var buffer bytes.Buffer
		buffer.WriteString(s1)
		buffer.WriteString(",")
		buffer.WriteString(s2)
		resultUrl = buffer.String()
	}
	resp, err := http.Head(resultUrl)
	if err != nil {

	}
	resultUrl = resp.Request.URL.String()
	return resultUrl

}
