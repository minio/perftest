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
	"os"
	"os/exec"
	"strconv"
	"sync"
	"syscall"
)

// For all unixes we need to bump allowed number of open files to a
// higher value than its usual default of '1024'. The reasoning is
// that this value is too small.
func setMaxOpenFiles() error {
	var rLimit syscall.Rlimit
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		return err
	}
	// Set the current limit to Max, it is usually around 4096.
	// TO increase this limit further user has to manually edit
	// `/etc/security/limits.conf`
	rLimit.Cur = rLimit.Max
	return syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
}

func init() {
	os.Setenv("ACCESS_KEY", os.Getenv("ACCESS_KEY"))
	os.Setenv("SECRET_KEY", os.Getenv("SECRET_KEY"))
	os.Setenv("ENDPOINT", os.Getenv("ENDPOINT"))
	os.Setenv("CONCURRENCY", os.Getenv("CONCURRENCY"))
}

func main() {

	setErr := setMaxOpenFiles()
	if setErr != nil {
		fmt.Println("Error bumping up open file limits: <ERROR> ", setErr)
		return
	}

	concurrency, err := strconv.Atoi(os.Getenv("CONCURRENCY"))
	if err != nil {
		fmt.Println("Please set a valid integer for concurrency level. ex: `export CONCURRENCY=100`: ", err)
		return
	}
	var wg sync.WaitGroup
	f, _ := os.Create("output.log")
	defer f.Close()
	testCmd := "./minio.test -test.timeout 3600s"
	if len(os.Args[1]) != 0 {
		testCmd = fmt.Sprintf("./minio.test -test.timeout 3600s -test.run %s", os.Args[1])
	}
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(routineId int) {
			defer wg.Done()

			out, err := exec.Command("sh", "-c", testCmd).Output()
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Fprintf(f, "\nGoroutine: %d\n", routineId+1)
			fmt.Fprintf(f, string(out))

		}(i)
	}
	wg.Wait()

}
