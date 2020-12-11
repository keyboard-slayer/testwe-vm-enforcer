/**
 * Copyright (C) 2020 Jordan DALCQ & Contributors
 *
 * This file is part of TestWe VM enforcer.
 *
 * TestWe VM enforcer is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * TestWe VM enforcer is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with TestWe VM enforcer.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	directory := fmt.Sprintf("%s\\TestWe\\TestWe.exe", os.Getenv("APPDATA"))

	_, err := os.Stat(directory)
	if os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "You're braindead, you forgot to install TestWe !\n")
		os.Exit(1)
	}

	exe, err := ioutil.ReadFile(directory);
	if err != nil {
		panic(err)
	}

	data := string(exe)
	regex, _ := regexp.Compile("\\(.*QEMU.*\\)")
	matches := regex.FindAllString(data, -1)[0]
	matches = matches[1:len(matches)-1]
	patch := "("

	for _, vm := range strings.Split(matches, "|") {
		vm := strings.Replace(vm, string(vm[0]), "$", -1)
		patch += vm + "|"
	}

	patch = patch[:len(patch)-1] + ")"
	newdata := strings.Replace(data, "("+matches+")", patch, -1)

	_ = os.Rename(directory, directory+".bak")
	file, err := os.Create(directory)
	if err != nil {
		panic(err)
	}

	_, err = file.Write([]byte(newdata))
	if err != nil {
		panic(err)
	}

	_ = file.Close()
}
