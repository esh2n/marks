/*
Copyright © 2020 Shunya Endoh

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// for docker
var lang string

var dirName string
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "starting project by using marks.",
	Run: func(cmd *cobra.Command, args []string) {
		if dirName != "" {
			fmt.Println("marks: " + dirName + " creating...")
			if lang != "" {
				cloneCmd := exec.Command("git", "clone", "-b", lang, "https://github.com/shunyaendoh1215/docker-file.git", dirName)
				if err := cloneCmd.Start(); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				cloneCmd.Wait()
			} else {
				cloneCmd := exec.Command("git", "clone", "https://github.com/shunyaendoh1215/markup-template.git", dirName)
				if err := cloneCmd.Start(); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				cloneCmd.Wait()
			}
			os.Chdir("./" + dirName)
			remoteCmd := exec.Command("git", "remote", "remove", "origin")
			if err := remoteCmd.Start(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			remoteCmd.Wait()
			removeCmd := exec.Command("rm", "-rf", ".git")
			if err := removeCmd.Start(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			removeCmd.Wait()
			fmt.Println("marks: finish!")
		} else {
			fmt.Println("Error: expected 1 argument but not existed.")
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.PersistentFlags().StringVarP(&dirName, "dirName", "d", "", "directory name")
	initCmd.PersistentFlags().StringVarP(&lang, "lang", "l", "", "docker template flag")
}
