# imdb-chart-fetcher [![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org) [![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gauravICT/imdb-chart-fetcher)](https://github.com/gauravICT/imdb-chart-fetcher) [![GoReportCard example](https://goreportcard.com/badge/github.com/gauravICT/imdb-chart-fetcher)](https://goreportcard.com/report/github.com/gauravICT/imdb-chart-fetcher) [![HitCount](http://hits.dwyl.com/gauravICT/imdb-chart-fetcher.svg)](http://hits.dwyl.com/gauravICT/imdb-chart-fetcher)
A  command-line script in golang that takes input the chart_url and items_count , where chart_url is one of IMDB Top Indian Charts. Output produce by the script in a JSON format with top indian movies.

## How to use :
```
Installation: `git clone github.com/gauravICT/imdb-chart-fetcher`
```     
### 1. Running Directly using go run after going to the directry
```
`cd imdb-chart-fetcher`
Syntax: `go run main.go <chart_url> <items_count>`


go run main.go 'https://www.imdb.com/india/top-rated-indian-movies/' 2
```

### 2. Build : 
Syntax: `go build -o executable_name path/to/main/directory`

```
go build -o application ./imdb-chart-fetcher

or you can simply go to the directory and then type :-

go build -o application

then run using this syntax

./application 'https://www.imdb.com/india/top-rated-indian-movies/' 2

```

### 3. Output :runner:
```
$ go run main.go 'https://www.imdb.com/india/top-rated-indian-movies/' 2
[{"title":"Pather Panchali ","movie_release_year":"1955","imdb_rating":"8.6","summary":"Impoverished priest Harihar Ray, dreaming of a better life for himself and his family, leaves his rural Bengal village in search of work.","duration":"2h 5min","genre":"Drama"},{"title":"Gol Maal ","movie_release_year":"1979","imdb_rating":"8.6","summary":"A man's simple lie to secure his job escalates into more complex lies when his orthodox boss gets suspicious.","duration":"2h 24min","genre":"Comedy"}]

```
