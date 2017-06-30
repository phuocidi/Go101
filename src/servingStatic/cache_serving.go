package main 

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
	)

// Data structure to store a file in Memory
type cacheFile struc {
	content io.ReadSeeker
	modTime time.Time
}

var cache map[string]*cacheFile
var mutex = new(sync.RWMutex) // mutex to handle parallel cache changes

func main() {
	cache = make(map[string]*cacheFile)

	http.HandleFunc("/", serveFiles)
	http.ListenAndServe(":8080", nil)
}

func serveFiles(res http.ResponseWriter, req *http.Request) {
	mutex.RLock()
	v, found := cache[req.URL.Path] // Loads from the cache if it's already poulated
	mutex.RUnlock()

	// When the file isn't in the cache, starts loading process
	if !found {
		mutex.Lock() // map can't be written to concurrently or be read while being written to. 
		defer mutex.Unlock()
		fileName := "./files" + req.URL.Path
		f, err := os.Open(file)
		defer f.Close()

		if err != nil {
			http.NotFound(res, req)
			return
		}

		var b bytes.Buffer
		_, err = io.Copy(&b, f) // Copy the file to an in-memory  buffer
		if err != nil {
			http.NotFound(res, req)
			return
		}

		// puts the bytes into a Reader for later use
		r := bytes.NewReader(b.Bytes())

		info, _ := f.Stat()
		v := &cacheFile {
			content: r,
			modTime: info.ModTime(),
		}

		// store cache
		cache[req.URL.Path] = v
	}

	http.ServeContent(res, req, req.URL.Path, v.modTime, v.content)
}



