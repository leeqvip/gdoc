package gdoc

import "testing"

func TestGetSidebar(t *testing.T)  {
	h := DefaultHandler
	h.GetSidebar("default")
}

func TestGetContent(t *testing.T)  {
	h := DefaultHandler
	c := h.GetContent("default", "index.html")
	t.Log(c)
}