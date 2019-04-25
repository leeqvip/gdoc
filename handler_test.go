package gindoc

import "testing"

func TestGetSidebar(t *testing.T)  {
	getSidebar("default")
}

func TestGetContent(t *testing.T)  {
	c := getContent("default", "index.html")
	t.Log(c)
}