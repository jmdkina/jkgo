package main

import (
	"fmt"
	"jk/jklog"

	"github.com/anaskhan96/soup"
	"github.com/tebeka/selenium"
	"os"
	"strings"
	"time"
)

type IstStruct struct {
	starturl string
}

func newIST(url string) *IstStruct {
	return &IstStruct{
		starturl: url,
	}
}

func (ist *IstStruct) queryGlobal() string {
	resp, err := soup.Get(ist.starturl)
	if err != nil {
		return ""
	}
	jklog.L().Infoln(resp)
	doc := soup.HTMLParse(resp)
	links := doc.FindAll("div", "class", "thumb-container-big")
	jklog.L().Infoln("len links ", len(links))

	for _, link := range links {
		imgu := link.Find("img")
		fmt.Println("imgu: ", imgu.Text())
		url := link.Find("img").Attrs()["href"]
		fmt.Println(link.Text(), "| Link :", url)
	}
	return ""
}

func (ist *IstStruct) queryWithSele() string {
	// Start a Selenium WebDriver server instance (if one is not already
	// running).
	const (
		// These paths will be different on your system.
		seleniumPath    = "/opt/data/proj/apps/selenium-core-1.0-20080914.225453.jar"
		geckoDriverPath = "/opt/data/proj/apps/geckodriver"
		port            = 8011
	)
	opts := []selenium.ServiceOption{
		selenium.StartFrameBuffer(),           // Start an X frame buffer for the browser to run in.
		selenium.GeckoDriver(geckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
		selenium.Output(os.Stderr),            // Output debug information to STDERR.
	}
	selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}
	defer service.Stop()

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	// Navigate to the simple playground interface.
	if err := wd.Get("http://www.baidu.com"); err != nil {
		panic(err)
	}

	// Get a reference to the text box containing code.
	elem, err := wd.FindElement(selenium.ByCSSSelector, "#code")
	if err != nil {
		panic(err)
	}
	// Remove the boilerplate code already in the text box.
	if err := elem.Clear(); err != nil {
		panic(err)
	}

	// Enter some new code in text box.
	err = elem.SendKeys(`
		package main
		import "fmt"
		func main() {
			fmt.Println("Hello WebDriver!\n")
		}
	`)
	if err != nil {
		panic(err)
	}

	// Click the run button.
	btn, err := wd.FindElement(selenium.ByCSSSelector, "#run")
	if err != nil {
		panic(err)
	}
	if err := btn.Click(); err != nil {
		panic(err)
	}

	// Wait for the program to finish running and get the output.
	outputDiv, err := wd.FindElement(selenium.ByCSSSelector, "#output")
	if err != nil {
		panic(err)
	}

	var output string
	for {
		output, err = outputDiv.Text()
		if err != nil {
			panic(err)
		}
		if output != "Waiting for remote server..." {
			break
		}
		time.Sleep(time.Millisecond * 100)
	}

	fmt.Printf("%s", strings.Replace(output, "\n\n", "\n", -1))

	// Example Output:
	// Hello WebDriver!
	//
	// Program exited.
	return ""
}
