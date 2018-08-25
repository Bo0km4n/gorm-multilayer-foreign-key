package main

import (
	"log"

	"github.com/k0kubun/pp"
)

func init() {
	InitMysql()
}

func main() {
	r := Region{
		Shops: []Shop{
			Shop{
				Name: "shop1",
				Books: []Book{
					Book{
						Name:  "book1",
						Price: 100,
					},
					Book{
						Name:  "book2",
						Price: 200,
					},
				},
			},
		},
	}

	if err := mysql.Create(&r).Error; err != nil {
		log.Fatal(err)
	}

	pp.Println(r)
}
