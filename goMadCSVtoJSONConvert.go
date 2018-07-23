package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type PodStats struct {
	ServerName string
	DateTime   string
	Project    string
	PodName    string
	CpuReq     string
	CpuPercent string
	CpuLim     string
	MemReq     string
	MemPercent string
	MemLimit   string
	CpuUsage   string
}

var Matched int
var root string
var recorddate string

func init() {
	flag.StringVar(&root, "root", "", "Enable recursive/directory mode")
	flag.StringVar(&recorddate, "recorddate", "", "Set the search date")
	flag.Parse()
}

func main() {

	log.SetFlags(log.LstdFlags | log.Ldate | log.Lmicroseconds | log.Lshortfile)

	if strings.Contains(root, "/") || strings.Contains(root, "\\") {
		root = filepath.FromSlash(root)
	} else {
		root = root + "/"
		root = filepath.FromSlash(root)
	}

	d, err := time.Parse("20060102", recorddate)
	if err != nil {
		log.Printf("ERROR: invalid date specified! %s\n %s\n", d, err)
		return
	}

	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}

	for i, file := range files {
		if strings.Contains(file.Name(), recorddate) {
			Matched++
			reportcontents(root + file.Name())
		} else if len(files) == i+1 && Matched < 1 {
			log.Println("ERROR: Unable to match date - no files found")
			return
		}
	}

}

func Split(r rune) bool {
	return r == '.' || r == '\\' || r == '/'
}

func reportcontents(source string) {
	sourcefilename := source
	csvFile, err := os.Open(sourcefilename)
	PodServerName := strings.FieldsFunc(sourcefilename, Split)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.Comma = ' '

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var pod PodStats
	var pods []PodStats

	for _, each := range csvData {
		pod.ServerName = PodServerName[1]
		pod.DateTime = each[0] + " " + each[1]
		pod.Project = each[2]
		pod.PodName = each[3]
		pod.CpuReq = each[4]
		pod.CpuPercent = each[5]
		pod.CpuLim = each[6]
		pod.MemReq = each[7]
		pod.MemPercent = each[8]
		pod.MemLimit = each[9]
		pod.CpuUsage = each[10]
		pods = append(pods, pod)
	}

	jsonData, err := json.MarshalIndent(pods, "", "\t")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(jsonData))
}
