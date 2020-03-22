package href

import (
	"golang.org/x/net/html"
	"io"
	"log"
	"strings"
)

type Link struct {
	Href string
	Text string
}

func findLinkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var r []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		r = append(r, findLinkNodes(c)...)
	}
	return r
}

func getText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var r string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		r += getText(c) + " "
	}
	return strings.Join(strings.Fields(r), " ")
}

func buildLink(n *html.Node) Link {
	var r Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			r.Href = attr.Val
			break
		}
	}
	r.Text = getText(n)
	return r
}

func Parse(r io.Reader) ([]Link, error) {
	links := []Link{}
	doc, err := html.Parse(r)
	if err != nil {
		log.Fatal(err)
	}
	linkNodes := findLinkNodes(doc)
	for _, el := range linkNodes {
		links = append(links, buildLink(el))
	}
	return links, nil
}
