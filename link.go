package main

import (
    "fmt"
    "net/http"
    "golang.org/x/net/html"
)

func main() {
    fmt.Println("tool")

    var target string
    var depth int

    fmt.Print("Digite o alvo: ")
    fmt.Scanln(&target)

    fmt.Print("Digite a profundidade: ")
    fmt.Scanln(&depth)

    visited := make(map[string]bool)
    links := []string{target}
    traverseLinks(links, visited, depth)
}

func traverseLinks(links []string, visited map[string]bool, depth int) {
    if depth == 0 {
        return
    }

    var nextLinks []string

    for _, link := range links {
        if visited[link] {
            continue
        }

        visited[link] = true

        response, err := http.Get(link)
        if err != nil {
            continue
        }

        defer response.Body.Close()

        z := html.NewTokenizer(response.Body)

        for {
            tt := z.Next()

            switch {
            case tt == html.ErrorToken:
                break

            case tt == html.StartTagToken:
                t := z.Token()

                if t.Data == "a" {
                    for _, a := range t.Attr {
                        if a.Key == "href" {
                            fmt.Println(a.Val)
                            nextLinks = append(nextLinks, a.Val)
                            break
                        }
                    }
                }
            }
        }
    }

    traverseLinks(nextLinks, visited, depth-1)
}
