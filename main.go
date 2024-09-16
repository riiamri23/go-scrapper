package main

import (
	// import Colly

	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"

	// "os"

	"github.com/davecgh/go-spew/spew"
	"github.com/gocolly/colly"
)

// initialize a data structure to keep the scraped data
type Product struct {
	Url, Image, Name, Price string
}

type DataTask struct {
	subject, description, assignee_name, status string
}

func main() {
	var dataTasks []DataTask

	// instantiate a new collector object
	c := colly.NewCollector(
		colly.AllowedDomains("support.dataon.com"),
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
	c.OnHTML("body script:last-child", func(e *colly.HTMLElement) {
		// printing all URLs associated with the <a> tag on the page
		// initialize a new Product instance
		s := e.DOM.Find("script").Text()

		// spew.Dump(s)

		dataTask := DataTask{}
		dataTask.description = s
		// add the product instance with scraped data to the list of products
		dataTasks = append(dataTasks, dataTask)
		spew.Dump(dataTasks)
	})

	// triggered once scraping is done (e.g., write the data to a CSV file)
	c.OnScraped(func(r *colly.Response) {
		// Convert the slice to JSON and write to a file
		if err := writeToJSONFile(dataTasks, "dataTasks.json"); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Data has been written to dataTasks.json")
		}

		// if err := writeToCSVFile(dataTasks, "dataTasks.csv"); err != nil {
		// 	fmt.Println("Error:", err)
		// } else {
		// 	fmt.Println("Data Has been written to product.json")
		// }

	})

	// open the target URL
	c.Visit("https://support.dataon.com/dashboard/devtimelinebydeveloper.cfm?dept=HR&txtStartDate=2024-09-09&txtEndDate=2024-09-09&selEmp=41265&btnSubmit=View&chktasktype=E&chktasktype=BE&chktasktype=BI&chktasktype=I&chktasktype=S&chktasktype=CRQ&chkonlycurrent=1")

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
