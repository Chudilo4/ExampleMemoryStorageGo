package main

import (
	"storage_test/storages"
)

func main() {
	StorageRam := storages.CreateStoregeRam()
	go storages.TestSet(StorageRam)
	go storages.TestGet(StorageRam)
	go storages.TestUnset(StorageRam)
	for {
		switch {
			default:
				continue
		}
	}
}