package blogposts

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

type StubFailingsFS struct {
}

func (s StubFailingsFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, I am always failed")
}

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd,go
---
Hello
World`
		secondBody = `Title: Post 2
Description: Description 2
Tags: rust,borrow-checker
---
B
L
M`
	)

	fs := fstest.MapFS{
		"hello-world.md":   {Data: []byte(firstBody)},
		"hello-world-2.md": {Data: []byte(secondBody)},
	}
	posts, err := NewPostsFromFs(fs)

	if err != nil {
		t.Fatal(err)
	}

	got := posts[0]
	want := Post{
		Title:       "Post 2",
		Description: "Description 2",
		Tags:        []string{"rust", "borrow-checker"},
		Body: `B
L
M`,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
