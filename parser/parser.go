package parser

import (
	"fmt"
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/gocolly/colly"
)

type TestCase struct {
	Inputs  string
	Outputs string
}

func GetProblemLinks(competitionUrl string) ([]string, []error) {
	problemLinks := make(map[string]struct{})
	var links []string
	var errors []error

	c := colly.NewCollector(
		colly.AllowedDomains("codeforces.com", "www.codeforces.com"),
	)

	c.OnHTML("table.problems a", func(e *colly.HTMLElement) {
		link := "https://codeforces.com/" + e.Attr("href")
		if strings.Contains(link, "/problem/") {
			problemLinks[link] = struct{}{}
		}
	})

	c.OnError(func(_ *colly.Response, err error) {
		errors = append(errors, err)
	})

	c.OnScraped(func(_ *colly.Response) {
		for link, _ := range problemLinks {
			links = append(links, link)
		}
	})

	c.Visit(competitionUrl)

	return links, errors
}

func GetTestCases(problemUrl string) ([]TestCase, []error) {
	var testInputs []string
	var testOutputs []string
	var testCases []TestCase
	var errors []error

	c := colly.NewCollector(
		colly.AllowedDomains("codeforces.com", "www.codeforces.com"),
	)

	c.OnHTML("div.input pre", func(e *colly.HTMLElement) {
		testInputs = append(testInputs, e.Text)
	})

	c.OnHTML("div.output pre", func(e *colly.HTMLElement) {
		testOutputs = append(testOutputs, e.Text)
	})

	c.OnScraped(func(_ *colly.Response) {
		if len(testInputs) != len(testOutputs) {
			err := fmt.Errorf("Inputs and outputs don't match. Found %d inputs and %d outputs",
				len(testInputs), len(testOutputs))
			errors = append(errors, err)
		} else {
			for i := 0; i < len(testInputs); i++ {
				testCases = append(testCases, TestCase{
					Inputs:  testInputs[i],
					Outputs: testOutputs[i],
				})
			}
		}
	})

	c.Visit("https://codeforces.com/contest/1228/problem/B")

	return testCases, errors
}

func GetProblemPdf(problemUrl string, saveLocation string) error {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return err
	}

	linkParts := strings.Split(problemUrl, "/")
	filename := linkParts[len(linkParts)-1]

	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.Grayscale.Set(false)

	page := wkhtmltopdf.NewPage(problemUrl)

	pdfg.AddPage(page)

	err = pdfg.Create()
	if err != nil {
		return err
	}

	err = pdfg.WriteFile("./" + saveLocation + filename + ".pdf")
	if err != nil {
		return err
	}

	return nil
}
