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

import "testing"

func TestMavenImportStyle_Apply(t *testing.T) {
	artifact := Artifact{
		Group:       "io.projectreactor",
		Id:          "reactor-core",
		Description: "Non-Blocking Reactive Foundation for the JVM",
	}
	version := Version{
		Value: "3.2.12.RELEASE",
		Repository: Repository{
			Name: "Central",
			Url:  "https://repo1.maven.org/maven2/",
		},
		Date: "Sep, 2019",
	}

	style := new(MavenImportStyle)
	actual := style.Apply(artifact, version)
	expected := `<dependency>
    <groupId>io.projectreactor</groupId>
    <artifactId>reactor-core</artifactId>
    <version>3.2.12.RELEASE</version>
</dependency>`
	if actual != expected {
		t.Errorf("Expecting '%s', but got '%s'", expected, actual)
	}
}

func TestGradleImportStyle_Apply(t *testing.T) {
	artifact := Artifact{
		Group:       "io.projectreactor",
		Id:          "reactor-core",
		Description: "Non-Blocking Reactive Foundation for the JVM",
	}
	version := Version{
		Value: "3.2.12.RELEASE",
		Repository: Repository{
			Name: "Central",
			Url:  "https://repo1.maven.org/maven2/",
		},
		Date: "Sep, 2019",
	}

	style := new(GradleImportStyle)
	actual := style.Apply(artifact, version)
	expected := `implementation("io.projectreactor:reactor-core:3.2.12.RELEASE")`
	if actual != expected {
		t.Errorf("Expecting '%s', but got '%s'", expected, actual)
	}
}

func TestSbtImportStyle_Apply(t *testing.T) {
	artifact := Artifact{
		Group:       "io.projectreactor",
		Id:          "reactor-core",
		Description: "Non-Blocking Reactive Foundation for the JVM",
	}
	version := Version{
		Value: "3.2.12.RELEASE",
		Repository: Repository{
			Name: "Central",
			Url:  "https://repo1.maven.org/maven2/",
		},
		Date: "Sep, 2019",
	}

	style := new(SbtImportStyle)
	actual := style.Apply(artifact, version)
	expected := `libraryDependencies += "io.projectreactor" % "reactor-core" % "3.2.12.RELEASE"`
	if actual != expected {
		t.Errorf("Expecting '%s', but got '%s'", expected, actual)
	}
}

func TestIvyImportStyle_Apply(t *testing.T) {
	artifact := Artifact{
		Group:       "io.projectreactor",
		Id:          "reactor-core",
		Description: "Non-Blocking Reactive Foundation for the JVM",
	}
	version := Version{
		Value: "3.2.12.RELEASE",
		Repository: Repository{
			Name: "Central",
			Url:  "https://repo1.maven.org/maven2/",
		},
		Date: "Sep, 2019",
	}

	style := new(IvyImportStyle)
	actual := style.Apply(artifact, version)
	expected := `<dependency org="io.projectreactor" name="reactor-core" rev="3.2.12.RELEASE"/>`
	if actual != expected {
		t.Errorf("Expecting '%s', but got '%s'", expected, actual)
	}
}

func TestGrapeImportStyle_Apply(t *testing.T) {
	artifact := Artifact{
		Group:       "io.projectreactor",
		Id:          "reactor-core",
		Description: "Non-Blocking Reactive Foundation for the JVM",
	}
	version := Version{
		Value: "3.2.12.RELEASE",
		Repository: Repository{
			Name: "Central",
			Url:  "https://repo1.maven.org/maven2/",
		},
		Date: "Sep, 2019",
	}

	style := new(GrapeImportStyle)
	actual := style.Apply(artifact, version)
	expected := `@Grapes(
    @Grab(group='io.projectreactor', module='reactor-core', version='3.2.12.RELEASE')
)`
	if actual != expected {
		t.Errorf("Expecting '%s', but got '%s'", expected, actual)
	}
}

func TestLeiningenImportStyle_Apply(t *testing.T) {
	artifact := Artifact{
		Group:       "io.projectreactor",
		Id:          "reactor-core",
		Description: "Non-Blocking Reactive Foundation for the JVM",
	}
	version := Version{
		Value: "3.2.12.RELEASE",
		Repository: Repository{
			Name: "Central",
			Url:  "https://repo1.maven.org/maven2/",
		},
		Date: "Sep, 2019",
	}

	style := new(LeiningenImportStyle)
	actual := style.Apply(artifact, version)
	expected := `[io.projectreactor/reactor-core "3.2.12.RELEASE"]`
	if actual != expected {
		t.Errorf("Expecting '%s', but got '%s'", expected, actual)
	}
}

func TestBuildrImportStyle_Apply(t *testing.T) {
	artifact := Artifact{
		Group:       "io.projectreactor",
		Id:          "reactor-core",
		Description: "Non-Blocking Reactive Foundation for the JVM",
	}
	version := Version{
		Value: "3.2.12.RELEASE",
		Repository: Repository{
			Name: "Central",
			Url:  "https://repo1.maven.org/maven2/",
		},
		Date: "Sep, 2019",
	}

	style := new(BuildrImportStyle)
	actual := style.Apply(artifact, version)
	expected := `'io.projectreactor:reactor-core:jar:3.2.12.RELEASE'`
	if actual != expected {
		t.Errorf("Expecting '%s', but got '%s'", expected, actual)
	}
}
