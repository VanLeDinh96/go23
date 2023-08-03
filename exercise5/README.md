# Web Scraping with Golang

![Golang Logo](https://golang.org/doc/gopher/frontpage.png)

This repository contains a Golang web scraping application to extract top news articles from the BBC News website. It demonstrates how to use the "colly" library to scrape data from a web page in a structured and automated manner.

## Prerequisites

Before running the scraper, make sure you have the following installed on your system:

- Golang (https://golang.org/dl/)

## Installation

1. Install the "colly" library using `go get`:

```bash
go get github.com/gocolly/colly/v2
```

## How to Use

1. Navigate to the project's root directory.

2. Run the Golang scraper:

```bash
go run main.go
```

The scraper will start visiting the BBC News website (https://www.bbc.com/news) and extract the titles of the top news articles.

3. View the output:

The scraped titles will be printed to the console.

## How It Works

The Golang web scraper uses the "colly" library to perform the scraping. Here's a step-by-step explanation of how it works:

1. Import the required packages:

```go
package main

import (
    "fmt"
    "log"
    "github.com/gocolly/colly/v2"
)
```

2. Create a new collector:

```go
c := colly.NewCollector()
```

3. Define a callback function for when the scraper encounters a specific HTML element containing the article titles:

```go
c.OnHTML(".gs-c-promo-heading__title", func(e *colly.HTMLElement) {
    title := e.Text
    fmt.Println(title)
})
```

4. Set up a callback for handling errors, in case any occur during the scraping process:

```go
c.OnError(func(r *colly.Response, err error) {
    log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
})
```

5. Start the scraping process by visiting the BBC News website:

```go
err := c.Visit("https://www.bbc.com/news")
if err != nil {
    log.Fatal(err)
}
```

6. Run the scraper:

```go
go run main.go
```

The application will send HTTP requests to the BBC News website, download the HTML content, and parse it to extract the top news article titles. The titles will then be printed to the console.

## Legal Disclaimer

Please note that web scraping may be subject to legal and ethical restrictions. Always review the website's terms of service and robots.txt file before scraping. Respect rate limits and avoid overloading the website's servers with too many requests.

## Contributing

Contributions are welcome! If you have any improvements or feature suggestions, feel free to open a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

[This](README.md) provides an overview of the Golang web scraper and includes instructions on how to install and use it. It also outlines the steps involved in scraping the BBC News website and provides a legal disclaimer to ensure responsible usage.

Feel free to customize [this](README.md) to suit your specific project or repository needs. Happy scraping!

Or navigate [back](../README.md) to see the overview of the Golang Learning Repository!
