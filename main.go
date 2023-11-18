package main

import (
	"rc_iot_agent/storages"
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