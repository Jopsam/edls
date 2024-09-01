package main

// Define file types constants with iota.
const (
	regularFile int = iota /* The regular file type is for files that are not defined by the others file type constants. */
	directoryFile
	executableFile
	compressedFile
	imageFile
	linkFile
)

// Define file extensions constants.
const (
	exe    = ".exe"
	debian = ".deb"
	zip    = ".zip"
	gz     = ".gz"
	tar    = ".tar"
	rar    = ".rar"
	png    = ".png"
	jpg    = ".jpg"
	gif    = ".gif"
)
