package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/gocolly/colly/v2"
)

// ImdbMovie :
type ImdbMovie struct {
	Title    string `json:"title"`
	Year     string `json:"movie_release_year"`
	Rating   string `json:"imdb_rating"`
	Summary  string `json:"summary"`
	Duration string `json:"duration"`
	Genre    string `json:"genre"`
}

// Result :
type Result struct {
	Movies []ImdbMovie
	// Since parallelism used so we must protect our result with a mutex.
	*sync.Mutex
}

// InputValidation : Validates command line argument
func InputValidation() (string, int, error) {

	if len(os.Args) != 3 {
		return "", 0, errors.New("incorrect parameters please enter <imdb_url> <item_count>")
	}

	imdbURL := os.Args[1]
	_, err := url.Parse(imdbURL)
	if err != nil {
		return "", 0, errors.New("invalid url please enter valid imdb_url")
	}

	limit, err := strconv.Atoi(os.Args[2])
	if err != nil || limit <= 0 {
		return "", 0, errors.New("invalid count please enter positive integer count > 0")
	}

	return imdbURL, limit, nil

}

func main() {
	// Taking valid input from os args
	imdbURL, count, err := InputValidation()
	if err != nil {
		fmt.Print(err)
		return
	}

	// Collector Component :
	c := colly.NewCollector(
		// Allow only IMDB links to be crawled
		colly.AllowedDomains("imdb.com", "www.imdb.com"),
		// Sets the recursion depth for links to visit
		colly.MaxDepth(2),
		// Enables asynchronous network requests
		colly.Async(),
	)

	// Setting limit :
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 5,
	})

	var result = Result{Mutex: &sync.Mutex{}}

	c.OnHTML("td.posterColumn > a", func(e *colly.HTMLElement) {
		if e.Index+1 > count {
			return
		}

		err := e.Request.Visit(e.Attr("href"))
		if err != nil {
			fmt.Println(err)
		}
	})

	// For removing date from title
	var yearRegex = regexp.MustCompile(`\(\d{4}\)`)

	c.OnHTML("#title-overview-widget", func(element *colly.HTMLElement) {

		year := element.ChildText("#titleYear")
		title := element.ChildText(".titleBar h1")
		rating := element.ChildText("div.ratingValue > strong > span")
		summary := element.ChildText(".summary_text")
		duration := element.ChildText("time")
		genre := element.ChildText("div.subtext > a:nth-child(4)")

		// Remove () from year
		year = strings.ReplaceAll(year, "(", "")
		year = strings.ReplaceAll(year, ")", "")

		// Remove year from title
		title = yearRegex.ReplaceAllString(title, "")
		title = strings.TrimLeft(title, " ")

		var movie = ImdbMovie{
			Title:    title,
			Year:     year,
			Rating:   rating,
			Summary:  summary,
			Duration: duration,
			Genre:    genre,
		}

		result.Lock()
		defer result.Unlock()

		result.Movies = append(result.Movies, movie)

	})

	c.Visit(imdbURL)
	c.Wait()

	jsonResult, err := json.Marshal(&result.Movies)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Print(string(jsonResult))
}
