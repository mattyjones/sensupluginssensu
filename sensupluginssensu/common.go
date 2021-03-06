// Copyright © 2016 Yieldbot <devops@yieldbot.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package sensupluginssensu

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/Sirupsen/logrus"
	//"github.com/yieldbot/sensupluginssensu/version"
)

// AcquireLocalChecks will retrieve the currently running configuration and
// return a list of all checks it knows about
func AcquireLocalChecks() {
	var jsonOut map[string][]string
	localChecks := exec.Command("/opt/sensu/embedded/bin/sensu-client", "-L", "error", "-d", "/etc/sensu/conf.d", "-P")

	out, err := localChecks.Output()
	if err != nil {
		syslogLog.WithFields(logrus.Fields{
			"check":   "sensupluginssensu",
			"client":  host,
			//"version": version.AppVersion(),
			"error":   err,
		}).Error(`Local Checks returned invalid output`)
	}

	err = json.Unmarshal(out, &jsonOut)
	if err != nil {
		syslogLog.WithFields(logrus.Fields{
			"check":   "sensupluginssensu",
			"client":  host,
			//"version": version.AppVersion(),
			"error":   err,
			"output":  out,
		}).Error(`Could not unmarshall the json.`)
	}

	fmt.Println(jsonOut)
	fmt.Println(jsonOut["transport"])
}
