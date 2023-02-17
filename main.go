package main

type Rsvp struct {
	Name, Email, Phone string
	WillAttend         bool
}

// Slice of pointers to instances of the Rsvp struct
var responses = make([]*Rsvp, 0, 10)

func main() {

}
