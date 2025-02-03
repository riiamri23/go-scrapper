package main

import (
	// import Colly

	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"

	// "time"

	"webscraper/data/timesheet"
	// "os"

	"github.com/davecgh/go-spew/spew"
	"github.com/gocolly/colly"
)

// initialize a data structure to keep the scraped data
type DataTask struct {
	Ticket       string `json:"Ticket"`
	Subject      string `json:"Subject"`
	Description  string `json:"Description"`
	AssigneeName string `json:"AssigneeName"`
	Status       string `json:"Status"`
	Hour         string `json:"Hour"`
	PlanStart    string `json:"PlanStart"`
	PlanEnd      string `json:"PlanEnd"`
	SlaDesc      string `json:"SlaDesc"`
	ReportedDate string `json:"ReportedDate"`
	EncDesc      string `json:"EncDesc"`
	ReleaseDate  string `json:"ReleaseDate"`
	ProjectId    string `json:"ProjectId"`
	ProjectName  string `json:"ProjectName"`
	ReportedBy   string `json:"ReportedBy"`
}

func main() {

	isTimesheet := flag.String("timesheet", "", "")
	// Parse command-line flags
	flag.Parse()
	if *isTimesheet != "" { // for run create api to timesheet json go run .\main.go -timesheet yes
		// fmt.Println(*isTimesheet)
		genPayloadTS()
		return
	}

	var dataTasks []DataTask

	var dateSelected = "2025-02-03"
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

func genPayloadTS() {
	// fmt.Println("testing")

	// currentDate := time.Now().Format("20060102") // e.g., "02032025"
	// startTime := "2025-02-03T01:00:00.000Z"
	// endTime := "2025-02-03T10:00:00.000Z"

	data, err := os.ReadFile("dataTasks.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// spew.Dump(data)

	var dataTasks []DataTask
	err = json.Unmarshal(data, &dataTasks)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Print the loaded data
	for _, task := range dataTasks {
		fmt.Printf("Ticket: %s\nSubject: %s\nAssignee: %s\nStatus: %s\n\n",
			task.Ticket, task.Subject, task.AssigneeName, task.Status)
	}

	var entries timesheet.Entries
	entries = timesheet.Entries{
		RequestNo:  "",
		ActRequest: "draft",
		StartDate:  "2025-02-03",
		EndDate:    "2025-02-03",
		IsMax:      false,
		RequestBy:  "DO215572",
		RequestFor: "DO215572",
		RefDoc:     "",
		Remark:     "",
		Template: timesheet.Template{
			HdnRowCount: 0,
			Detail:      []timesheet.EntryDetail{},
		},
		HdnField: []string{},
		Dates:    map[string]timesheet.DateEntry{},
	}

	spew.Dump(entries)

	return
}
