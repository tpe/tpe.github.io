package main

import (
	"bytes"
	"github.com/russross/blackfriday"
	"path"
	"strings"
)

type Post struct {
	file string
	html []byte
	name string
	path string
	slug string
}

type Posts []Post

// Extract the post name from the pathname, which holds the meta data. This
// avoids including foreign data in the raw post markdown file.
// Filename format is YYYYMMDD-name-of-post.md
func (p *Post) SetNameFromFilename(pathname string) {
	info := strings.Split(pathname, ".md")[0]
	name := strings.Replace(info, "-", " ", -1)

	words := strings.Fields(name)
	smallwords := " a an if in of on the to "

	for index, word := range words {
		if strings.Contains(smallwords, " "+word+" ") {
			words[index] = word
		} else {
			words[index] = strings.Title(word)
		}
	}

	p.name = strings.Join(words, " ")
}

// Extract the post slug from the pathname, which holds the meta data. This
// avoids including foreign data in the raw post markdown file.
// Filename format is YYYYMMDD-name-of-post.md
func (p *Post) SetSlugFromFilename(pathname string) {
	info := strings.Split(pathname, ".md")[0]
	slug := strings.ToLower(info)

	p.slug = slug
}

// Set the post html by first converting the markdown to html and then
// decorating the result with the base post template.
func (p *Post) SetHtmlFromMarkdown(html []byte, md []byte) {
	text := blackfriday.MarkdownBasic(md)

	html = bytes.Replace(html, []byte("{{ name }}"), []byte(p.name), -1)
	html = bytes.Replace(html, []byte("{{ text }}"), []byte(text), -1)

	p.html = html
}

// Loop through source directories recursively to create category directories,
// post directories, and index html files. Also works for uncategorized posts.
func ProcessPosts(base []byte, in string, out string) {
	for _, info := range GetDirectoryListing(in) {
		pathname := info.Name()

		if info.IsDir() {
			// Create category directory
			CreateDirectory(path.Join(out, pathname))

			// Process category posts recursively
			ProcessPosts(base, path.Join(in, pathname), path.Join(out, pathname))
		} else {
			// Get the markdown
			md := GetFile(path.Join(in, pathname))

			// Build the post
			post := new(Post)
			post.SetNameFromFilename(pathname)
			post.SetSlugFromFilename(pathname)
			post.SetHtmlFromMarkdown(base, md)
			post.path = "/" + path.Join(out, post.slug)

			// Create the output directory
			CreateDirectory(path.Join(out, post.slug))

			// Create the output file
			CreateFile(path.Join(out, post.slug, "index.html"), post.html)
		}
	}
}
