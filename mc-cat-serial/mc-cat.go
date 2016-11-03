/*
 * Minio Cloud Storage, (C) 2015, 2016 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func execCommand(command string) ([]byte, error) {
	return exec.Command("sh", "-c", command).Output()
}

func main() {
	f, err := os.Create("output.log")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	obj := os.Getenv("OBJECT")
	if obj == "" {
		log.Fatal("Set Object to be downloaded, `export OBJECT=<mc-alias>/<bucket>/<object>`.")
	}

	expectedMD5 := os.Getenv("MD5")
	if expectedMD5 == "" {
		log.Fatal("Set Expected MD5 `export MD5=xxxxx`.")
	}

	countStr := os.Getenv("COUNT")
	if countStr == "" {
		log.Fatal("COUNT not set, `export COUNT=100")
	}

	count, err := strconv.Atoi(countStr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(f, "Expected Md5Sum: "+expectedMD5)

	for i := 0; i < count; i++ {
		log.Println("\nLoop ", i+1, ": Copy operation and verifying Md5sum\n")
		cpBackCmd := "mc cat " + obj + " | md5sum"
		out, err := execCommand(cpBackCmd)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if strings.Contains(string(out), expectedMD5) {
			fmt.Fprintf(f, "\nSuccess Match: Loop: %d\n", i+1)
			fmt.Fprintf(f, string(out))
		} else {
			fmt.Fprintf(f, "\nFailed Match: loop: %d\n", i+1)
			fmt.Fprintf(f, string(out))
		}
	}
}
