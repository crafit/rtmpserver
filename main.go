package main

import (

)

func main() {
	s := NewServer()

	s.ListenAndServe(":1935")
}
