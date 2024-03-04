package core

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/cespare/xxhash"
)

func GetHash(path string) uint64 {
	f, err := os.Open(path)

	if err != nil {
		log.Printf(err.Error())
		return 0
	}
	defer f.Close()

	h := xxhash.New()

	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	return h.Sum64()
}

func GetMapOfDuplicates(path string) map[uint64][]string {
	hashes := make(map[uint64][]string) // hash and list of filepaths where it occurs
	duplicates := make(map[uint64][]string)

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf(err.Error())
			return nil
		}

		if info.IsDir() {
			return nil
		}

		hash := GetHash(path)
		if hash == 0 {
			return nil
		}

		_, exists := hashes[hash]

		hashes[hash] = append(hashes[hash], path)

		if exists {
			duplicates[hash] = hashes[hash]
		}

		return nil
	})
	return duplicates
}
