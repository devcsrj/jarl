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

import "fmt"

// Interface for creating import style for artifacts
type ImportStyle interface {
	// Constructs the import string for the provided artifact and version
	Apply(artifact Artifact, version Version) string
}

// https://maven.apache.org
type MavenImportStyle struct{}

func (m MavenImportStyle) Apply(artifact Artifact, version Version) string {
	return fmt.Sprintf(`<dependency>
    <groupId>%s</groupId>
    <artifactId>%s</artifactId>
    <version>%s</version>
</dependency>`, artifact.Group, artifact.Id, version.Value)
}

// https://gradle.org
type GradleImportStyle struct{}

func (g GradleImportStyle) Apply(artifact Artifact, version Version) string {
	return fmt.Sprintf(`implementation("%s:%s:%s")`, artifact.Group, artifact.Id, version.Value)
}

// https://www.scala-sbt.org
type SbtImportStyle struct{}

func (s SbtImportStyle) Apply(artifact Artifact, version Version) string {
	return fmt.Sprintf(`libraryDependencies += "%s" %% "%s" %% "%s"`, artifact.Group, artifact.Id, version.Value)
}

// https://ant.apache.org/ivy/
type IvyImportStyle struct{}

func (i IvyImportStyle) Apply(artifact Artifact, version Version) string {
	return fmt.Sprintf(`<dependency org="%s" name="%s" rev="%s"/>`, artifact.Group, artifact.Id, version.Value)
}

// http://docs.groovy-lang.org/latest/html/documentation/grape.html
type GrapeImportStyle struct{}

func (g GrapeImportStyle) Apply(artifact Artifact, version Version) string {
	return fmt.Sprintf(`@Grapes(
    @Grab(group='%s', module='%s', version='%s')
)`, artifact.Group, artifact.Id, version.Value)
}

// https://leiningen.org
type LeiningenImportStyle struct{}

func (l LeiningenImportStyle) Apply(artifact Artifact, version Version) string {
	return fmt.Sprintf(`[%s/%s "%s"]`, artifact.Group, artifact.Id, version.Value)
}

// https://buildr.apache.org
type BuildrImportStyle struct{}

func (b BuildrImportStyle) Apply(artifact Artifact, version Version) string {
	return fmt.Sprintf(`'%s:%s:jar:%s'`, artifact.Group, artifact.Id, version.Value)
}
