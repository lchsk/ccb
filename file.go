package main

import (
    "io/ioutil"
    "os"
)

func ReadDir(path string) []os.FileInfo {
    files, _ := ioutil.ReadDir(path)

    return files
}

func AddPathsToBuffer(b *Buffer) {
    b.data = ReadDir(b.path)
}
