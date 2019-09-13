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
package cmd

import (
	"errors"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/devcsrj/jarl"
	"github.com/gookit/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "jarl",
	Short: "Your trusty 'Jar l'ocator",
	Long: `Locate jar coordinates right from your terminal.

Example:
$ jarl reactor-core

The ending coordinates is automatically copied to the clipboard.
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a search term")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		repo := new(jarl.Mvnrepository)
		repo.Init("https://mvnrepository.com")

		q := args[0]
		fmt.Print("\U0001F50E  Searching for ")
		color.Cyan.Printf("'%s'\n", q)
		results := repo.SearchArtifacts(q, 1)
		artifact := selectArtifact(results.Artifacts)

		details := repo.GetArtifactDetails(artifact.Group, artifact.Id)
		version := selectVersion(details)

		style := selectImportStyle(artifact, version)
		fmt.Println()
		color.Magenta.Println(style)
		fmt.Println()
		clipboard.WriteAll(style)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func selectArtifact(artifacts []jarl.Artifact) jarl.Artifact {
	if len(artifacts) == 0 {
		color.Red.Println("No results found")
		os.Exit(0)
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}:",
		Selected: "\U0001F4CD  {{ .Value | red | cyan }}",
		Details: "\U0001F4CD  " + `{{ .Group | cyan }}:{{ .Id | cyan }}
-------------------------------------
{{ .Description }}`,
	}
	artifactPrompt := promptui.Select{
		Label:     "Artifacts",
		Items:     artifacts,
		Size:      5,
		Templates: templates,
	}

	i, _, err := artifactPrompt.Run()
	if err != nil {
		color.Red.Println("Cancelled")
		os.Exit(0)
	}
	return artifacts[i]
}

func selectVersion(details jarl.Details) jarl.Version {
	header := `{{ "License:" | faint }}	` + details.License

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}:",
		Active:   "\U0001F4CD  {{ .Value | cyan }}",
		Inactive: "   {{ .Value | cyan }}",
		Selected: "\U0001F4CD  {{ .Value | cyan }}",
		Details: header + `
{{ "Repository:" | faint }}	{{ .Repository.Name }}
{{ "Date:" | faint }}	{{ .Date }}
`,
	}

	prompt := promptui.Select{
		Label:     "Versions",
		Items:     details.Versions,
		Templates: templates,
		Size:      4,
		Searcher: func(input string, index int) bool {
			return strings.HasPrefix(details.Versions[index].Value, input)
		},
		StartInSearchMode: true,
	}

	i, _, err := prompt.Run()
	if err != nil {
		color.Red.Println("Cancelled")
		os.Exit(0)
	}
	return details.Versions[i]
}

func selectImportStyle(artifact jarl.Artifact, version jarl.Version) string {
	type Import struct {
		Style string
		Value string
	}
	imports := []Import{
		{Style: "maven", Value: new(jarl.MavenImportStyle).Apply(artifact, version)},
		{Style: "gradle", Value: new(jarl.GradleImportStyle).Apply(artifact, version)},
		{Style: "sbt", Value: new(jarl.SbtImportStyle).Apply(artifact, version)},
		{Style: "ivy", Value: new(jarl.IvyImportStyle).Apply(artifact, version)},
		{Style: "grape", Value: new(jarl.GrapeImportStyle).Apply(artifact, version)},
		{Style: "leiningen", Value: new(jarl.LeiningenImportStyle).Apply(artifact, version)},
		{Style: "buildr", Value: new(jarl.BuildrImportStyle).Apply(artifact, version)},
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}:",
		Active:   "\U0001F4CD  {{ .Style | cyan }}",
		Inactive: "   {{ .Style | cyan }}",
		Selected: "\U0001F4CD  {{ .Style | cyan }}",
		Details: `
{{ .Value | magenta }}`,
	}

	prompt := promptui.Select{
		Label:     "Style",
		Items:     imports,
		Templates: templates,
		Size:      4,
	}

	i, _, err := prompt.Run()
	if err != nil {
		color.Red.Println("Cancelled")
		os.Exit(0)
	}

	return imports[i].Value
}
