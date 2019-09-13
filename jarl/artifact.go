/*
Copyright © 2019 Reijhanniel Jearl Campos <devcsrj@apache.org>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package jarl

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"strconv"
)

type SearchResults struct {
	Total     int
	Page      int
	Artifacts []Artifact
}

type Artifact struct {
	Group       string
	Id          string
	Description string
}

type Repository struct {
	Name string
	Url  string
}

type Details struct {
	License      string
	Repositories []Repository
	Versions     []Version
}

type Version struct {
	Value      string
	Repository Repository
	Date       string
}

// Interface for querying maven artifacts
type ArtifactApi interface {
	// Queries for artifacts using the provided query 'q'
	SearchArtifacts(q string, page int) SearchResults

	// Queries the details of the provided group and id
	GetArtifactDetails(group string, id string) Details
}

// Implementation for fetching artifacts from https://mvnrepository.com
type Mvnrepository struct {
	url string
}

func (e *Mvnrepository) Init(url string) {
	e.url = url
}

func (e *Mvnrepository) SearchArtifacts(q string, page int) SearchResults {
	path := fmt.Sprintf("/search?q=%s&p=%d&sort=relevance", q, page)
	resp, err := http.Get(e.url + path)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("Expecting 200 OK from %s, but got %d", e.url, resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("Could not parse response from server", err)
	}

	var results []Artifact
	total, err := strconv.Atoi(doc.Find("#maincontent > h2 > b").Text())
	if err != nil {
		log.Fatalf("Could not parse the total of number of results")
	}

	doc.Find("#maincontent > div.im").Each(func(i int, elem *goquery.Selection) {
		groupId := elem.Find("div.im-header > p > a:nth-child(1)").Text()
		artifactId := elem.Find("div.im-header > p > a:nth-child(2)").Text()
		description := elem.Find("div.im-description").Text()
		if len(groupId) != 0 { // skip the ad row
			results = append(results, Artifact{Group: groupId, Id: artifactId, Description: description})
		}
	})
	return SearchResults{
		Total:     total,
		Page:      page,
		Artifacts: results,
	}
}

func (e *Mvnrepository) GetArtifactDetails(group string, id string) Details {
	path := fmt.Sprintf("/artifact/%s/%s", group, id)
	resp, err := http.Get(e.url + path)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("The artifact '%s:%s' was not found", group, id)
	}

	return e.parseArtifactDetails(resp.Body)
}

func (e *Mvnrepository) parseArtifactDetails(r io.Reader) Details {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Fatal("Could not parse response from server", err)
	}

	var versions []Version
	doc.Find("#snippets > div > div > div > table > tbody").Each(func(i int, tbody *goquery.Selection) {
		// 3.2.x, 3.1.x, etc
		tbody.Children().Each(func(i int, tr *goquery.Selection) {
			version := tr.Find("a.vbtn").Text()
			a := tr.Find("a.b.lic")
			date := tr.Find("td:last-child").Text()
			versions = append(versions, Version{
				Value: version,
				Repository: Repository{
					Name: a.Text(),
					Url:  a.AttrOr("href", "#unknown"),
				},
				Date: date,
			})
		})
	})

	license := doc.Find("#maincontent > table > tbody > tr:nth-child(1) > td > span").Text()
	var repos []Repository
	doc.Find("#snippets > ul.tabs > li > a").Each(func(i int, a *goquery.Selection) {
		repos = append(repos, Repository{
			Name: a.Text(),
			Url:  a.AttrOr("href", "#unknown"),
		})
	})

	return Details{
		License:      license,
		Repositories: repos,
		Versions:     versions,
	}
}
