package main

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"time"
)

func gitInit() {
	fmt.Println("initialized")
	os.Mkdir(".minigit", 0755)
	os.Mkdir(".minigit/objects", 0755)
	os.Mkdir(".minigit/refs", 0755)
	os.Mkdir(".minigit/refs/heads", 0755)
	os.WriteFile(".minigit/HEAD", []byte("ref: refs/heads/main"), 0644)
}

func gitCommit(message string) {
	author := "author"
	date := strconv.Itoa(int(time.Now().Unix()))

	//TODO: go to the same directory as .minigit
	files, _ := os.ReadDir(".")
	entries := make([][]byte, 0)
	for _, file := range files {
		//TODO: search recursively
		if !file.IsDir() {
			filename := file.Name()
			content, _ := os.ReadFile(filename)
			// calculate the filename of blob object
			checksum := sha1.Sum(append([]byte("blob"+" "+strconv.Itoa(len(content))+"\x00"), content...))
			blobHash := hex.EncodeToString(checksum[:])

			// calculate the content of blob object
			var b bytes.Buffer
			w := zlib.NewWriter(&b)
			w.Write(append([]byte("blob"+" "+strconv.Itoa(len(content))+"\x00"), content...))
			w.Close()
			os.Mkdir(".minigit/objects/"+blobHash[:2], 0755)
			os.WriteFile(".minigit/objects/"+blobHash[:2]+"/"+blobHash[2:], b.Bytes(), 0644)

			mode := "100644"
			entries = append(entries, append([]byte(mode+" "+filename+"\x00"), checksum[:]...))
		}
	}

	// calculate the filename of tree object
	combined := make([]byte, 0)
	for _, entry := range entries {
		combined = append(combined, entry...)
	}
	checksum := sha1.Sum(append([]byte("tree"+" "+strconv.Itoa(len(combined))+"\x00"), combined...))
	treeHash := hex.EncodeToString(checksum[:])

	// calculate the content of tree object
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write(append([]byte("tree"+" "+strconv.Itoa(len(combined))+"\x00"), combined...))
	w.Close()

	os.Mkdir(".minigit/objects/"+treeHash[:2], 0755)
	os.WriteFile(".minigit/objects/"+treeHash[:2]+"/"+treeHash[2:], b.Bytes(), 0644)

	// calculate the filename of commit object
	// simplifies format
	commitData := []byte("tree" + " " + treeHash + "\n" + "author" + " " + author + " " + date + "\n" + message)
	checksum = sha1.Sum(append([]byte("commit"+" "+strconv.Itoa(len(commitData))+"\x00"), commitData...))
	commitHash := hex.EncodeToString(checksum[:])

	// calculate the content of commit object
	var cb bytes.Buffer
	cw := zlib.NewWriter(&cb)
	cw.Write(append([]byte("commit"+" "+strconv.Itoa(len(commitData))+"\x00"), []byte(commitData)...))
	cw.Close()

	os.Mkdir(".minigit/objects/"+commitHash[:2], 0755)
	os.WriteFile(".minigit/objects/"+commitHash[:2]+"/"+commitHash[2:], cb.Bytes(), 0644)

	os.WriteFile(".minigit/refs/heads/main", []byte(commitHash+"\n"), 0644)
}

func main() {
	command := os.Args[1]
	switch command {
	case "init":
		gitInit()
	case "commit":
		if os.Args[2] == "-m" {
			message := os.Args[3]
			// git add -A && git commit -m <message>
			gitCommit(message)
		}
	default:
		panic("no command")
	}
}
