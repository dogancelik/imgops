package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func isUrl(targetPath string) bool {
	return strings.Contains(targetPath, "http:") || strings.Contains(targetPath, "https:")
}

func getIdMap() map[string]string {
	return map[string]string{
		"google": "#t85",
		"bing":   "#t101",
		"tineye": "#t11",
		"reddit": "#t97",
		"yandex": "#t72",
		"baidu":  "#t74",
		"so":     "#t109",
		"sogou":  "#t110",
	}
}

func findHref(document, targetStr, finalUrl string) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(document))
	if err != nil {
		return strings.Fields(""), err
	} else {
		queryList := getQueryList(targetStr)
		debug("Queries: %v", queryList)
		foundUrls := make([]string, 0, len(queryList))

		for _, query := range queryList {
			href, attrOk := doc.Find(query).Attr("href")
			debug("Get href from query '%s': %v", query, attrOk)

			if attrOk {
				foundUrls = append(foundUrls, href)
			}
		}

		if len(foundUrls) > 0 {
			return foundUrls, nil
		} else {
			return strings.Fields(finalUrl), errors.New(fmt.Sprintf("No link is found in targets: %s", queryList))
		}
	}
}

// Returns queries of found targets
func getQueryList(s string) []string {
	split := strings.Split(s, ",")

	found := make([]string, 0, len(split))
	for _, val := range split {
		key := strings.TrimSpace(val)
		query, mapOk := QueryMap[key]
		if mapOk {
			debug("Query for '%s' is found: %s", key, query)
			found = append(found, query)
		}
	}

	return found
}
