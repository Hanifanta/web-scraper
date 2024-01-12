package handlers

import "log"

func ErrorVisit(err error) {
	if err != nil {
		log.Fatalln("Failed to make GET request", err)
	}
}

func ErrorCreateFile(err error) {
	if err != nil {
		log.Fatalln("Failed to create output CSV file", err)
	}
}

func ErrorWriteFile(err error) {
	if err != nil {
		log.Fatalln("Failed to write record to CSV", err)
	}
}
