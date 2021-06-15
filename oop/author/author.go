package author

import "fmt"

type Author struct {
	FirstName string
	LastName  string
	Bio       string
}

func (a Author) FullName() string {
	return fmt.Sprintf("%s %s", a.FirstName, a.LastName)
}

type BlogPost struct {
	Title   string
	Content string
	Author
}

func (b BlogPost) Details() {
	fmt.Println("Title: ", b.Title)
	fmt.Println("Content: ", b.Content)
	fmt.Println("Author: ", b.FullName())
	fmt.Println("Bio: ", b.Bio)
}
