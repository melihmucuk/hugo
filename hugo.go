package hugo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
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

func (h HurriyetAPI) ListArticles(query *Query) ([]Article, error) {
	data, err := h.request("articles", query)
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

func (h HurriyetAPI) SingleArticle(Id string, query *Query) (Article, error) {
	data, err := h.request("articles/"+Id, query)
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

func (h HurriyetAPI) Search(keyword string, query *Query) (Search, error) {
	data, err := h.request("search/"+keyword, query)
	if err != nil {
		return Search{}, err
	}

	var search Search
	if jsonErr := json.Unmarshal(data, &search); jsonErr != nil {
		// check response for error
		return Search{}, jsonErr
	}

	return search, nil
}

func (h HurriyetAPI) ListPaths(query *Query) ([]Path, error) {
	data, err := h.request("paths", query)
	if err != nil {
		return nil, err
	}

	var paths []Path
	if jsonErr := json.Unmarshal(data, &paths); jsonErr != nil {
		// check response for error
		return nil, jsonErr
	}

	return paths, nil
}

func (h HurriyetAPI) SinglePath(Id string, query *Query) (Path, error) {
	data, err := h.request("paths/"+Id, query)
	if err != nil {
		return Path{}, err
	}

	var path Path
	if jsonErr := json.Unmarshal(data, &path); jsonErr != nil {
		// check response for error
		return Path{}, jsonErr
	}

	return path, nil
}

func (h HurriyetAPI) ListWriters(query *Query) ([]Writer, error) {
	data, err := h.request("writers", query)
	if err != nil {
		return nil, err
	}

	var writers []Writer
	if jsonErr := json.Unmarshal(data, &writers); jsonErr != nil {
		// check response for error
		return nil, jsonErr
	}

	return writers, nil
}

func (h HurriyetAPI) SingleWriter(Id string, query *Query) (Writer, error) {
	data, err := h.request("writers/"+Id, query)
	if err != nil {
		return Writer{}, err
	}

	var writer Writer
	if jsonErr := json.Unmarshal(data, &writer); jsonErr != nil {
		// check response for error
		return Writer{}, jsonErr
	}

	return writer, nil
}

func (h HurriyetAPI) request(url string, query *Query) ([]byte, error) {
	req, err := http.NewRequest("GET", h.baseUrl+url, nil)
	req.Header.Add("apikey", h.apiKey)
	req.Header.Add("accept", "application/json")

	if query != nil {
		q := req.URL.Query()
		if len(query.Filter) > 0 {
			q.Add("$filter", strings.Join(query.Filter, ","))
		}

		if len(query.Select) > 0 {
			q.Add("$select", strings.Join(query.Select, ","))
		}

		if query.Top != "" {
			q.Add("$top", query.Top)
		}

		if query.Skip != "" {
			q.Add("$skip", query.Skip)
		}

		if query.S != "" {
			q.Add("s", query.S)
		}

		req.URL.RawQuery = q.Encode()
	}

	resp, err := h.client.Do(req)

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
