package gapi

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/url"
)

type SearchResult struct {
	Id          int64    `json:"id"`
	Uid         string   `json:"uid"`
	Title       string   `json:"title"`
	Uri         string   `json:"uri`
	Url         string   `json:"url`
	Slug        string   `json:"slug`
	Type        string   `json:"type`
	Tags        []string `json:"tags`
	IsStarred   bool     `json:"isStarred`
	FolderId    int64    `json:"folderId,omitempty`
	FolderUid   string   `json:"folderUid,omitempty`
	FolderTitle string   `json:"folderTitle,omitempty`
	FolderUrl   string   `json:"folderUrl,omitempty`
}

func (c *Client) Search(params map[string]string) ([]SearchResult, error) {
	searchResults := make([]SearchResult, 0)

	vals := url.Values{}
	for k, v := range params {
		vals.Add(k, v)
	}

	req, err := c.newRequest("GET", "/api/search", vals, nil)

	if err != nil {
		return searchResults, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return searchResults, err
	}
	if resp.StatusCode != 200 {
		return searchResults, errors.New(resp.Status)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return searchResults, err
	}
	err = json.Unmarshal(data, &searchResults)
	return searchResults, err
}
