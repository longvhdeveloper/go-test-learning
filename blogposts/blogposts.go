package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
	"testing/fstest"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func NewPostsFromFs(fileSystem fstest.MapFS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, f := range dir {
		post, err := getPost(fileSystem, f)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
func getPost(fileSystem fs.FS, f fs.DirEntry) (Post, error) {
	postFile, err := fileSystem.Open(f.Name())
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()

	return newPost(postFile)
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagSeparator         = "Tags: "
)

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	title := readMetaLine(titleSeparator)
	description := readMetaLine(descriptionSeparator)

	tags := strings.Split(readMetaLine(tagSeparator), ",")

	body := readBody(scanner)

	post := Post{Title: title, Description: description, Tags: tags, Body: body}
	return post, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan()

	buffer := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buffer, scanner.Text())
	}
	body := strings.TrimSuffix(buffer.String(), "\n")
	return body
}
