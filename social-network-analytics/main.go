package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Post struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

type User struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Posts []*Post `json:"posts"`
}

type Users struct {
	Users []*User `json:"users"`
}

func main() {
	f, err := os.Open("users_posts.json")
	if err != nil {
		log.Fatalf("Unable to open file: %d", err)
	}
	var u Users
	json.NewDecoder(f).Decode(&u)

	p := []Post{}
	for _, v := range u.Users {
		for _, po := range v.Posts {
			if strings.Contains(po.Content, "morning") {
				p = append(p, *po)
			}
		}
	}
	fmt.Println("Found posts")
	for _, v := range p {
		fmt.Println(v)
	}
}
