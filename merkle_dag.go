package main

import (
	"crypto"
)

type hash []byte

type hashable interface {
	Hash() hash
}

type directory struct {
	Entries map[string]hash
}

func (d *directory) Hash() hash {
	hash := crypto.SHA256()
	for (k, v := range d.Entries) {
		hash.Write(k)
		hash.Write(v)
	}
	return hash.Digest()
}

func (d *directory) Add(h hashable) {
	d.Entries[h.Hash()] = h
}

type file struct {
	Data []byte
}

func (f *file) Hash() hash {
	hash := crypto.SHA256()
	hash.Write(f.Data)
	return hash.Digest()
}

func main() {
	
}