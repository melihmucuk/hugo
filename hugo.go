package hugo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	baseUrl = "https://api.hurriyet.com.tr/v1/"
)

type HurriyetAPI struct {
	apiKey  string
	client  *http.Client
	baseUrl string
}

func NewHurriyetAPI(apiKey string) *HurriyetAPI {
	return &HurriyetAPI{
		apiKey:  apiKey,
		baseUrl: baseUrl,
		client:  http.DefaultClient,
	}
}

func (h HurriyetAPI) GetArticles() ([]Article, error) {
	data, err := h.request("articles")
	if err != nil {
		return nil, err
	}

	var articles []Article
	if jsonErr := json.Unmarshal(data, &articles); jsonErr != nil {
		// check response for error
		return nil, jsonErr
	}

	return articles, nil
}

func (h HurriyetAPI) GetArticleById(Id string) (Article, error) {
	data, err := h.request("articles/" + Id)
	if err != nil {
		return Article{}, err
	}

	var article Article
	if jsonErr := json.Unmarshal(data, &article); jsonErr != nil {
		// check response for error
		return Article{}, jsonErr
	}

	return article, nil
}

func (h HurriyetAPI) request(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", h.baseUrl+url, nil)
	req.Header.Add("apikey", h.apiKey)
	req.Header.Add("accept", "application/json")
	resp, err := h.client.Do(req)

	resp.StatusCode

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
