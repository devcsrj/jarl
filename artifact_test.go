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
	"testing"
)

func TestSearch(t *testing.T) {
	repo := new(Mvnrepository)
	repo.Init("https://mvnrepository.com")

	result := repo.SearchArtifacts("reactor-core", 1)
	if len(result.Artifacts) != 10 {
		t.Errorf("Expecting 10 result, but got %d", len(result.Artifacts))
	}
	if result.Total < 10 {
		t.Errorf("Expecting more than 10 total, but got %d", result.Total)
	}
	if result.Page != 1 {
		t.Errorf("Expecting page 1, but got %d", result.Page)
	}

	actual := result.Artifacts[0]
	if actual.Group != "io.projectreactor" {
		t.Errorf("Expecting 'io.projectreactor', but got %s", actual.Group)
	}
	if actual.Id != "reactor-core" {
		t.Errorf("Expecting 'reactor-core', but got %s", actual.Id)
	}
	if len(actual.Description) == 0 {
		t.Errorf("Expecting a description, but got none")
	}
}

func TestGet(t *testing.T) {
	repo := new(Mvnrepository)
	repo.Init("https://mvnrepository.com")

	result := repo.GetArtifactDetails("io.projectreactor", "reactor-core")
	if result.License != "Apache 2.0" {
		t.Errorf("Expecting a 'Apache 2.0', but got %s", result.License)
	}
	if len(result.Repositories) < 5 {
		t.Errorf("Expecting at least 5 repositories hosting the artifact, "+
			"but got %d", len(result.Repositories))
	}

	versions := result.Versions
	actual := versions[len(versions)-1]
	expected := Version{
		Value: "2.0.0.RC1",
		Repository: Repository{
			Name: "Central",
			Url:  "/repos/central",
		},
		Date: "Feb, 2015",
	}
	if actual != expected {
		t.Errorf("Expecting %+v, but got %+v", expected, actual)
	}

}
