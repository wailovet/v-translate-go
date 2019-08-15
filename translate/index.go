package translate

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wailovet/osmanthuswine/src/core"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Index struct {
}

func (n *Index) Index(req core.Request, res core.Response) {
	text := req.GET["query"]
	isEn := req.GET["isEn"]
	result := ""
	if isEn != "" {
		result, _ = Translate(text, "en", "zh")
	} else {
		result, _ = Translate(text, "zh", "en")
	}
	res.DisplayByData(result)
}

func Translate(source, sourceLang, targetLang string) (string, error) {
	var text []string
	var result []interface{}

	encodedSource := url.QueryEscape(source)
	client := &http.Client{}
	url := "https://translate.googleapis.com/translate_a/single?client=gtx&sl=" +
		sourceLang + "&tl=" + targetLang + "&dt=t&q=" + encodedSource

	reqest, err := http.NewRequest("GET", url, nil)

	reqest.Header.Add("accept-charset", "utf-8")
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0")
	r, err := client.Do(reqest) //提交
	if err != nil {
		return "err", errors.New("Error getting translate.googleapis.com")
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	print(string(url))
	if err != nil {
		return "err", errors.New("Error reading response body")
	}

	bReq := strings.Contains(string(body), `<title>Error 400 (Bad Request)`)
	if bReq {
		return "err", errors.New("Error 400 (Bad Request)")
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return "err", errors.New("Error unmarshaling data")
	}

	if len(result) > 0 {
		inner := result[0]
		for _, slice := range inner.([]interface{}) {
			for _, translatedText := range slice.([]interface{}) {
				text = append(text, fmt.Sprintf("%v", translatedText))
				break
			}
		}
		cText := strings.Join(text, "")

		return cText, nil
	} else {
		return "err", errors.New("No translated data in responce")
	}
}
