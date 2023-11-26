windows_amd64_garble_nopack:
	go clean -cache
	GOOS=windows GOARCH=amd64 GARBLE_EXPERIMENTAL_CONTROLFLOW=1 garble -debug -tiny -literals -seed=random build -o badger.exe badger.go
clean:
	rm badger.exe