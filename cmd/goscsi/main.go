/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"log"
	"os"

	"github.com/talos-systems/go-scsi/scsi"
)

func main() {
	device, err := scsi.NewDevice(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	err = device.Identify()
	if err != nil {
		log.Fatal(err)
	}

	err = device.StandardInquiry()
	if err != nil {
		log.Fatal(err)
	}

	_, err = device.Page0Inquiry()
	if err != nil {
		log.Fatal(err)
	}

	err = device.Page80Inquiry()
	if err != nil {
		log.Fatal(err)
	}

	err = device.Page83Inquiry()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", device.ID)
}
