package main

import (
	// import Colly

	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"

	// "os"

	"github.com/gocolly/colly"
)

// initialize a data structure to keep the scraped data
type Product struct {
	Url, Image, Name, Price string
}

func main() {
	var products []Product

	// instantiate a new collector object
	c := colly.NewCollector(
		colly.AllowedDomains("www.scrapingcourse.com"),
	)

	// called before an HTTP request is triggered
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	// triggered when the scraper encounters an error
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong: ", err)
	})

	// fired when the server responds
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
	})

	// triggered when a CSS selector matches an element
	c.OnHTML("a", func(e *colly.HTMLElement) {
		// printing all URLs associated with the <a> tag on the page
		// initialize a new Product instance
		product := Product{}

		// scrape the target data
		product.Url = e.ChildAttr("a", "href")
		product.Image = e.ChildAttr("img", "src")
		product.Name = e.ChildText(".product-name")
		product.Price = e.ChildText(".price")

		// add the product instance with scraped data to the list of products
		products = append(products, product)

	})

	// triggered once scraping is done (e.g., write the data to a CSV file)
	c.OnScraped(func(r *colly.Response) {
		// Convert the slice to JSON and write to a file
		if err := writeToJSONFile(products, "product.json"); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Data has been written to people.json")
		}

		if err := writeToCSVFile(products, "product.csv"); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Data Has been written to product.json")
		}

	})

	// open the target URL
	c.Visit("https://www.scrapingcourse.com/ecommerce")

}

func writeToJSONFile(data interface{}, filename string) error {
	// Create the JSON file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode data to JSON and write to the file
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Optional: pretty-print JSON with indentation
	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}

func writeToCSVFile(data interface{}, filename string) error {
	// open the CSV file
	file, err := os.Create("products.csv")
	if err != nil {
		log.Fatalln("Failed to create output CSV file", err)
	}
	defer file.Close()

	// initialize a file writer
	writer := csv.NewWriter(file)

	// write the CSV headers
	headers := []string{
		"Url",
		"Image",
		"Name",
		"Price",
	}
	writer.Write(headers)

	// write each product as a CSV row
	for _, product := range data.([]Product) {
		// convert a Product to an array of strings
		record := []string{
			product.Url,
			product.Image,
			product.Name,
			product.Price,
		}

		// add a CSV record to the output file
		writer.Write(record)
	}
	defer writer.Flush()

	return nil
}
