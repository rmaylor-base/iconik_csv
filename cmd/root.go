package cmd

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strconv"
)

type CMDArgs struct {
	AppID          string
	AuthToken      string
	CollectionID   string
	MetadataViewID string
	IconikUrl      string
}

var args CMDArgs

func Execute() error {
	err := argParse()
	if err != nil {
		panic(err)
	}

	var a *Assets
	a, err = getAssets(&args)
	if err != nil {
		panic(err)
	}

	err = createCSV(a)
	if err != nil {
		panic(err)
	}
	
	return nil
}

func createCSV(a *Assets) error {
	var keySlice []string
	var csvFile [][]string
	
	aDeref := reflect.ValueOf(*a.Objects[0])
	for i := 0; i < aDeref.NumField(); i++ {
		key := aDeref.Type().Field(i).Name
		keySlice = append(keySlice, key)
	}
	csvFile = append(csvFile, keySlice)
	
	for _, o := range a.Objects {
		var valuesLine []string
		valuesLine = append(valuesLine, o.DateCreated, o.DateModified, o.FileNames[0], o.Files[0].OriginalName, o.Formats[0].Status, strconv.FormatFloat(float64(o.FrameRate), 'f', 2, 32), o.ID, o.InCollections[0], o.Keyframes[0].URL, (strconv.Itoa(o.OriginalResolution["width"]) + "x" + strconv.Itoa(o.OriginalResolution["height"])))
		csvFile = append(csvFile, valuesLine)
	}

	// Create a new CSV writer
	file, err := os.Create("output.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	// Write the data to the CSV file
	for _, row := range csvFile {
		if err := writer.Write(row); err != nil {
			panic(err)
		}
	}

	// Flush the CSV writer to ensure any buffered data is written to the file
	writer.Flush()

	fmt.Println("CSV file created successfully")
	return nil
}

func getAssets(args *CMDArgs) (*Assets, error) {
	var a *Assets
	url := args.IconikUrl + "/API/assets/v1/collections/" + args.CollectionID + "/contents/?object_types=assets"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("App-ID", args.AppID)
	req.Header.Set("Auth-Token", args.AuthToken)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bytestream, _ := io.ReadAll(res.Body)

	err = json.Unmarshal(bytestream, &a)
	if err != nil {
		panic(err)
	}

	return a, nil
}

func argParse() error {
	flag.StringVar(&args.AppID, "app-id", "", "The app ID provided by Iconik")
	flag.StringVar(&args.AuthToken, "auth-token", "", "The auth token provided by Iconik")
	flag.StringVar(&args.CollectionID, "collection-id", "", "The collection ID provided by Iconik")
	flag.StringVar(&args.MetadataViewID, "metadata-view-id", "", "The metadata view ID provided by Iconik")
	flag.StringVar(&args.IconikUrl, "iconik-url", "https://preview.iconik.cloud", "Iconik URL (default https://preview.iconik.cloud)")
	flag.Parse()
	if args.AppID == "" {
		fmt.Println("No app ID provided.")
		os.Exit(1)
	}
	if args.AuthToken == "" {
		fmt.Println("No auth token provided.")
		os.Exit(1)
	}
	if args.CollectionID == "" {
		fmt.Println("No collection ID provided.")
		os.Exit(1)
	}
	if args.MetadataViewID == "" {
		fmt.Println("No metadata view ID provided.")
		os.Exit(1)
	}
	if args.IconikUrl == "" || args.IconikUrl == "preview.iconik.cloud" || args.IconikUrl == "https://preview.iconik.cloud" {
		args.IconikUrl = "https://preview.iconik.cloud"
	} else {
		fmt.Println("not a valid Iconik URL.")
		os.Exit(1)
	}
	return nil
}

type Assets struct {
	Objects []*Object `json:"objects"`
}

type Object struct {
	DateCreated        string         `json:"date_created"`
	DateModified       string         `json:"date_modified"`
	FileNames          []string       `json:"file_names"`
	Files              []*File        `json:"files"`
	Formats            []*Format      `json:"formats"`
	FrameRate          float32        `json:"frame_rate"`
	ID                 string         `json:"id"`
	InCollections      []string       `json:"in_collections"`
	Keyframes          []*Keyframe    `json:"keyframes"`
	OriginalResolution map[string]int `json:"original_resolution"`
}

type File struct {
	OriginalName string `json:"original_name"`
}

type Keyframe struct {
	URL string `json:"url"`
}

type Format struct {
	Status string `json:"status"`
}
