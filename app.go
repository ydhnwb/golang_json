package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/ydhnwb/converting-json-go/entity"
)

func main() {

	//You have ':' in key json, 'articles:' shouldbe 'articles', but i leave as it is
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Fatal(err)
	}

	var users []entity.User
	err = json.Unmarshal([]byte(data), &users)
	if err != nil {
		log.Fatal(err)
	}

	var usersWithoutPhoneNumber []entity.User
	var usersThatHaveArticles []entity.User
	var usersThatHaveAnnisWord []entity.User
	var usersThatHaveArticlesPostedIn2020 []entity.User
	var usersThatBornIn1986 []entity.User
	var articlesThatContainsTipsWord []entity.Article
	var articlesBeforeAugust2019 []entity.Article

	for _, value := range users {

		if isBornIn1986(value.Profile.Birthday) {
			usersThatBornIn1986 = append(usersThatBornIn1986, value)
		}

		if len(value.Profile.Phones) == 0 {
			usersWithoutPhoneNumber = append(usersWithoutPhoneNumber, value)
		}

		if len(value.Articles) > 0 {
			usersThatHaveArticles = append(usersThatHaveArticles, value)

			if isHaveArticlesPostedIn2020(value.Articles) {
				usersThatHaveArticlesPostedIn2020 = append(usersThatHaveArticlesPostedIn2020, value)
			}

			articlesBeforeAugust2019 = append(articlesBeforeAugust2019, articlesThatPostedBeforeAugust2019(value.Articles)...)
			articlesThatContainsTipsWord = append(articlesThatContainsTipsWord, isArticleTitleContainsTipsWord(value.Articles)...)
		}

		//assumption : annis and Annis are the same
		if strings.Contains(strings.ToLower(value.Profile.FullName), "annis") {
			usersThatHaveAnnisWord = append(usersThatHaveAnnisWord, value)
		}
	}

	println("===============USER WITHOUT PHONE NUM=====================")
	data1, err := json.MarshalIndent(usersWithoutPhoneNumber, "", " ")
	if err != nil {
		log.Fatal("Cannot encode to JSON ", err)
	}
	fmt.Fprintf(os.Stdout, "%s", data1)

	println("===============USER THAT HAVE ARTICLES=====================")
	data2, err := json.MarshalIndent(usersThatHaveArticles, "", " ")
	if err != nil {
		log.Fatal("Cannot encode to JSON ", err)
	}
	fmt.Fprintf(os.Stdout, "%s", data2)

	println("===============USER THAT HAVE ANNIS on HIS/HER NAME=====================")
	data3, err := json.MarshalIndent(usersThatHaveAnnisWord, "", " ")
	if err != nil {
		log.Fatal("Cannot encode to JSON ", err)
	}
	fmt.Fprintf(os.Stdout, "%s", data3)

	println("===============USERS THAT HAVE ARTICLE POSTED ON 2020=====================")
	if len(usersThatHaveArticlesPostedIn2020) == 0 {
		println("No users that have artiicles posted in 2020")
	} else {
		data4, err := json.MarshalIndent(usersThatHaveArticlesPostedIn2020, "", " ")
		if err != nil {
			log.Fatal("Cannot encode to JSON ", err)
		}
		fmt.Fprintf(os.Stdout, "%s", data4)
	}

	println("===============ARTCLES THAT POSTED ON BEFORE AUGUST 2019=====================")
	if len(articlesBeforeAugust2019) == 0 {
		println("No articles that posted before 2019")
	} else {
		data5, err := json.MarshalIndent(articlesBeforeAugust2019, "", " ")
		if err != nil {
			log.Fatal("Cannot encode to JSON ", err)
		}
		fmt.Fprintf(os.Stdout, "%s", data5)
	}

	println("===============ARTCLES THAT HAVE 'tips' ON IT TITLES=====================")
	if len(articlesThatContainsTipsWord) == 0 {
		println("GHERE IS NO ARTICLES THAT HAVE 'tips' ON IT TITLES")
	} else {
		data6, err := json.MarshalIndent(articlesThatContainsTipsWord, "", " ")
		if err != nil {
			log.Fatal("Cannot encode to JSON ", err)
		}
		fmt.Fprintf(os.Stdout, "%s", data6)
	}

}

//assumption that "tips" and "Tips" are same
func isArticleTitleContainsTipsWord(articles []entity.Article) []entity.Article {
	var results []entity.Article
	for _, v := range articles {
		if strings.Contains(strings.ToLower(v.Title), "tips") {
			results = append(results, v)
		}
	}
	return results
}

func isBornIn1986(birthdate string) bool {
	layout := "2006-01-02"
	d, err := time.Parse(layout, birthdate)
	if err != nil {
		log.Fatalf("Cannot parse birthdate %v", err.Error())
	}
	return d.Year() == 1986
}

func isHaveArticlesPostedIn2020(articles []entity.Article) bool {
	pivot := false
	layout := "2006-01-02T15:04:05"
	for _, v := range articles {
		currentTime, err := time.Parse(layout, v.PublishedAt)
		if err != nil {
			log.Println("Error parsing date")
		}
		if currentTime.Year() == 2020 {
			pivot = true
			break
		}
	}
	return pivot
}

func articlesThatPostedBeforeAugust2019(articles []entity.Article) []entity.Article {
	//Article with posted in AUGUST 2019 will be not included
	layout := "2006-01-02T15:04:05"
	dateToCompare := time.Date(2019, 7, 0, 0, 0, 0, 0, time.UTC)
	var results []entity.Article
	for _, v := range articles {
		PublishedAt, err := time.Parse(layout, v.PublishedAt)
		if err != nil {
			log.Fatalf("sss: %s", err)
		}
		if PublishedAt.Before(dateToCompare) {
			results = append(results, v)
		}
	}
	return results
}
