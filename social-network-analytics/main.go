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

type PostInteraction struct {
	PostID int    `json:"post_id"`
	Action string `json:"action"`
	UserID int    `json:"user_id"`
}

type PostPopularity [][2]int

func main() {
	ps := []*PostInteraction{
		{
			PostID: 101,
			Action: "like",
			UserID: 2,
		},
		{
			PostID: 101,
			Action: "share",
			UserID: 3,
		},
		{
			PostID: 102,
			Action: "like",
			UserID: 1,
		},
	}

	mr := MapReducer[PostInteraction, int, string, PostPopularity]{
		PoolNumber: 5,
		MapperIn:   make(chan *PostInteraction),
		MapperOut:  make(chan *Pair[int, *string]),
		Mapper: func(pi *PostInteraction) (int, *string) {
			i := pi.PostID
			a := pi.Action
			return i, &a
		},
		Reducer: func(pp *PostPopularity, i int, s []*string) *PostPopularity {
			if pp == nil {
				sl := PostPopularity([][2]int{})
				pp = &sl
			}
			*pp = append(*pp, [2]int{i, len(s)})
			return pp
		},
	}

	mr.Accept(ps)
	mapped := mr.Get()
	fmt.Println(mapped)
}

func imperativeQuery() {
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
