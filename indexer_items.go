package main

type recordItem struct {
	id int
	content string
}

type indexItem struct {
	id int
	term string
	records []int
}
