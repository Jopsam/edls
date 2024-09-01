package main

import "time"

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

// Define file struct.
type file struct {
	name             string
	fileType         int
	isDirectory      bool
	isHidden         bool
	userName         string
	groupName        string
	size             int64
	modificationTime time.Time
	mode             string
}

// Define file type style struct.
type fileTypeStyle struct {
	icon   string
	color  string
	symbol string
}

var fileTypeStyleMap = map[int]fileTypeStyle{
	regularFile: {
		icon: "📄",
	},
	directoryFile: {
		icon:   "📁",
		color:  "blue",
		symbol: "/",
	},
	executableFile: {
		icon:   "🚀",
		color:  "green",
		symbol: "*",
	},
	compressedFile: {
		icon:  "🗜️",
		color: "red",
	},
	imageFile: {
		icon:  "🖼️",
		color: "magenta",
	},
	linkFile: {
		icon:  "🔗",
		color: "cyan",
	},
}
