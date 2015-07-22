// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package webp implements a decoder and encoder for WEBP images.

WEBP is defined at:
https://developers.google.com/speed/webp/docs/riff_container

Install CGO Version

Install `GCC` or `MinGW` (http://tdm-gcc.tdragon.net/download) at first,
and then run these commands:

	1. Assure set the `CGO_ENABLED` environment variable to `1` to enable `CGO` (Default is enabled).
	2. `go get github.com/chai2010/webp`
	3. `go run hello.go`


Install Pure Go Version

Pure Go versin only support `Decode` and `DecodeConfig`, `go test` will failed on some other api test.

	1. Assure set the `CGO_ENABLED` environment variable to `0` to disable `CGO` (Default is enabled).
	2. `go get github.com/chai2010/webp`
	3. `go run hello.go`


Examples

This is a simple example:

	package main

	import (
		"bytes"
		"fmt"
		"io/ioutil"
		"log"

		"github.com/chai2010/webp"
	)

	func main() {
		var buf bytes.Buffer
		var width, height int
		var data []byte
		var err error

		// Load file data
		if data, err = ioutil.ReadFile("./testdata/1_webp_ll.webp"); err != nil {
			log.Fatal(err)
		}

		// GetInfo
		if width, height, _, err = webp.GetInfo(data); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("width = %d, height = %d\n", width, height)

		// GetMetadata
		if metadata, err := webp.GetMetadata(data, "ICCP"); err != nil {
			fmt.Printf("Metadata: err = %v\n", err)
		} else {
			fmt.Printf("Metadata: %s\n", string(metadata))
		}

		// Decode webp
		m, err := webp.Decode(bytes.NewReader(data))
		if err != nil {
			log.Fatal(err)
		}

		// Encode lossless webp
		if err = webp.Encode(&buf, m, &webp.Options{Lossless: true}); err != nil {
			log.Fatal(err)
		}
		if err = ioutil.WriteFile("output.webp", buf.Bytes(), 0666); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Save output.webp ok\n")
	}

Decode and Encode as RGB format:

	m, err := webp.DecodeRGB(data)
	if err != nil {
		log.Fatal(err)
	}

	data, err := webp.EncodeRGB(m)
	if err != nil {
		log.Fatal(err)
	}

BUGS

Report bugs to <chaishushan@gmail.com>.

Thanks!
*/
package webp
