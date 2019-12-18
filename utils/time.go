// MIT License

// Copyright (c) 2019 gonethopper

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// * @Author: ankye
// * @Date: 2019-12-18 10:46:52
// * @Last Modified by:   ankye
// * @Last Modified time: 2019-12-18 10:46:52

package utils

import (
	"time"
)

// TimeYMDHIS get current time
// return format yearmouthday hour:minute:second
func TimeYMDHIS() string {
	return time.Now().Format("20060102 15:04:05")
}

// TimeYMDH get current time
// return format yearmouthday-hour
func TimeYMDH() string {
	return time.Now().Format("20060102-15")
}

// TimeYMD get current time
// return format yearmouthday
func TimeYMD() string {
	return time.Now().Format("20060102")
}
