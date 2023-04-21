package cmd

import (
	"flag"
	"fmt"
	"os"
)

type CMDArgs struct {
	appID          string
	authToken      string
	collectionID   string
	metadataViewID string
	iconikUrl      string
}

var args CMDArgs

func Execute() error {
	argParse()
	return nil
}

func argParse() error {
	flag.StringVar(&args.appID, "app-id", "", "The app ID provided by Iconik")
	flag.StringVar(&args.authToken, "auth-token", "", "The auth token provided by Iconik")
	flag.StringVar(&args.collectionID, "collection-id", "", "The collection ID provided by Iconik")
	flag.StringVar(&args.metadataViewID, "metadata-view-id", "", "The metadata view ID provided by Iconik")
	flag.StringVar(&args.iconikUrl, "iconik-url", "https://preview.iconik.cloud", "Iconik URL (default https://preview.iconik.cloud)")
	flag.Parse()
	if args.appID == "" {
		fmt.Println("No app ID provided.")
		os.Exit(1)
	}
	if args.authToken == "" {
		fmt.Println("No auth token provided.")
		os.Exit(1)
	}
	if args.collectionID == "" {
		fmt.Println("No collection ID provided.")
		os.Exit(1)
	}
	if args.metadataViewID == "" {
		fmt.Println("No metadata view ID provided.")
		os.Exit(1)
	}
	if args.iconikUrl == "" || args.iconikUrl == "preview.iconik.cloud" || args.iconikUrl == "https://preview.iconik.cloud" {
		args.iconikUrl = "https://preview.iconik.cloud"
	} else {
		fmt.Println("not a valid Iconik URL.")
		os.Exit(1)
	}
	fmt.Println(args)
	return nil
}
