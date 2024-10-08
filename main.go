package main

import (
	// import Colly

	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"

	// "os"

	"github.com/davecgh/go-spew/spew"
	"github.com/gocolly/colly"
)

// initialize a data structure to keep the scraped data
type DataTask struct {
	Ticket, Subject, Description, AssigneeName, Status, Hour, PlanStart, PlanEnd, SlaDesc, ReportedDate, EncDesc, ReleaseDate, ProjectId, ProjectName, ReportedBy string
}

func main() {
	var dataTasks []DataTask

	var dateSelected = "2024-09-17"
	var emp_id = "41265"

	var linkUrl = "https://support.dataon.com/dashboard/devtimelinebydeveloper.cfm?dept=HR&txtStartDate=" + dateSelected + "&txtEndDate=" + dateSelected + "&selEmp=" + emp_id + "&btnSubmit=View&chktasktype=E&chktasktype=BE&chktasktype=BI&chktasktype=I&chktasktype=S&chktasktype=CRQ&chkonlycurrent=1"

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
	c.OnHTML("body>script:last-child", func(e *colly.HTMLElement) {
		// printing all URLs associated with the <a> tag on the page
		// initialize a new Product instance
		s := e.DOM.Text()
		re := regexp.MustCompile(`[\n\t]+`)
		formattedString := re.ReplaceAllString(s, "")

		// with a word with regex
		reTCK := regexp.MustCompile(`tiket\["([^"]+)"\]\s*=\s*{([^}]*)}`)
		matchesTCK := reTCK.FindAllString(formattedString, -1)

		// spew.Dump(matchesTCK)

		if matchesTCK == nil {
			fmt.Println("No matches found.")
			return
		}
		spew.Dump(matchesTCK)

		for _, dataTCK := range matchesTCK {
			dataTask := DataTask{}

			// find the regex specific data for ticket
			reTicket := regexp.MustCompile(`tiket\["(TCK[0-9]{4}-[0-9]{7})"\]`)
			matchTicket := reTicket.FindStringSubmatch(dataTCK)
			if matchTicket == nil {
				dataTask.Ticket = ""
			} else {
				dataTask.Ticket = matchTicket[1]
			}

			// find the regex specific data for subject
			reSubject := regexp.MustCompile(`subject:\s*"([^"]+)"`)
			matchSubject := reSubject.FindStringSubmatch(dataTCK)
			if matchSubject == nil {
				dataTask.Subject = ""
			} else {
				dataTask.Subject = matchSubject[1]
			}

			// find the regex specific data for description
			reDescription := regexp.MustCompile(`description:\s*"([^"]+)"`)
			matchDescription := reDescription.FindStringSubmatch(dataTCK)
			if matchSubject == nil {
				dataTask.Description = ""
			} else {
				dataTask.Description = matchDescription[1]
			}

			// find the regex specific data for assignee name
			reAssigneeName := regexp.MustCompile(`assignee_name:\s*"([^"]+)"`)
			matchAssigneeName := reAssigneeName.FindStringSubmatch(dataTCK)
			if matchAssigneeName == nil {
				dataTask.AssigneeName = ""
			} else {
				dataTask.AssigneeName = matchAssigneeName[1]
			}

			// find the regex specific data for status
			reStatus := regexp.MustCompile(`status:\s*"([^"]+)"`)
			matchStatus := reStatus.FindStringSubmatch(dataTCK)
			if matchStatus == nil {
				dataTask.Status = ""
			} else {
				dataTask.Status = matchStatus[1]
			}

			// find the regex specific data for hour
			reHour := regexp.MustCompile(`hour:\s*"([^"]+)"`)
			matchHour := reHour.FindStringSubmatch(dataTCK)
			if matchHour == nil {
				dataTask.Hour = ""
			} else {
				dataTask.Hour = matchHour[1]
			}

			// find the regex specific data for plan start
			rePlanStart := regexp.MustCompile(`planStart:\s*"([^"]+)"`)
			matchPlanStart := rePlanStart.FindStringSubmatch(dataTCK)
			if matchPlanStart == nil {
				dataTask.PlanStart = ""
			} else {
				dataTask.PlanStart = matchPlanStart[1]
			}

			// find the regex specific data for status
			rePlanEnd := regexp.MustCompile(`planEnd:\s*"([^"]+)"`)
			matchPlanEnd := rePlanEnd.FindStringSubmatch(dataTCK)
			if matchPlanEnd == nil {
				dataTask.PlanEnd = ""
			} else {
				dataTask.PlanEnd = matchPlanEnd[1]
			}

			// find the regex specific data for status
			reSlaDesc := regexp.MustCompile(`slaDesc:\s*"([^"]+)"`)
			matchSlaDesc := reSlaDesc.FindStringSubmatch(dataTCK)
			if matchSlaDesc == nil {
				dataTask.SlaDesc = ""
			} else {
				dataTask.SlaDesc = matchSlaDesc[1]
			}

			// find the regex specific data for status
			reReportedDate := regexp.MustCompile(`reportedDate:\s*"([^"]+)"`)
			matchReportedDate := reReportedDate.FindStringSubmatch(dataTCK)
			if matchReportedDate == nil {
				dataTask.ReportedDate = ""
			} else {
				dataTask.ReportedDate = matchReportedDate[1]
			}
			// find the regex specific data for status
			reEncDesc := regexp.MustCompile(`encDesc:\s*"([^"]+)"`)
			matchEncDesc := reEncDesc.FindStringSubmatch(dataTCK)
			if matchEncDesc == nil {
				dataTask.EncDesc = ""
			} else {
				dataTask.EncDesc = matchEncDesc[1]
			}
			// find the regex specific data for status
			reReleaseDate := regexp.MustCompile(`releaseDate:\s*"([^"]+)"`)
			matchReleaseDate := reReleaseDate.FindStringSubmatch(dataTCK)
			if matchReleaseDate == nil {
				dataTask.ReleaseDate = ""
			} else {
				dataTask.ReleaseDate = matchReleaseDate[1]
			}
			// find the regex specific data for status
			reProjectId := regexp.MustCompile(`projectid:\s*"([^"]+)"`)
			matchProjectId := reProjectId.FindStringSubmatch(dataTCK)
			if matchProjectId == nil {
				dataTask.ProjectId = ""
			} else {
				dataTask.ProjectId = matchProjectId[1]
			}
			// find the regex specific data for status
			reProjectName := regexp.MustCompile(`projectName:\s*"([^"]+)"`)
			matchProjectName := reProjectName.FindStringSubmatch(dataTCK)
			if matchProjectName == nil {
				dataTask.ProjectName = ""
			} else {
				dataTask.ProjectName = matchProjectName[1]
			}
			// find the regex specific data for status
			reReportedBy := regexp.MustCompile(`reportedBy:\s*"([^"]+)"`)
			matchReportedBy := reReportedBy.FindStringSubmatch(dataTCK)
			if matchReportedBy == nil {
				dataTask.ReportedBy = ""
			} else {
				dataTask.ReportedBy = matchReportedBy[1]
			}

			// add the product instance with scraped data to the list of products
			dataTasks = append(dataTasks, dataTask)
		}

	})

	// triggered once scraping is done (e.g., write the data to a CSV file)
	c.OnScraped(func(r *colly.Response) {
		// Convert the slice to JSON and write to a file
		// spew.Dump(dataTasks)
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
	c.Visit(linkUrl)

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
		"Subject",
		"Description",
		"AssigneeName",
		"Status",
	}
	writer.Write(headers)

	// write each product as a CSV row
	for _, dataTask := range data.([]DataTask) {
		// convert a Product to an array of strings
		record := []string{
			dataTask.Subject,
			dataTask.Description,
			dataTask.AssigneeName,
			dataTask.Status,
		}

		// add a CSV record to the output file
		writer.Write(record)
	}
	defer writer.Flush()

	return nil
}
