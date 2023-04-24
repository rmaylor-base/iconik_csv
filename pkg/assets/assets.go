package assets

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strconv"

	"github.com/rmaylor-base/iconik_csv/cmd"
)

func CreateCSV(a *Assets) error {
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

func GetAssets(args *cmd.CMDArgs) (*Assets, error) {
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

	// // Unmarshal the JSON into a map
	// var data map[string]interface{}
	// if err := json.Unmarshal(bytestream, &data); err != nil {
	// 		panic(err)
	// }

	// // Create struct fields dynamically
	// var fields []reflect.StructField
	// for k, v := range data {
	// 		fieldName := strings.Title(k)
	// 		fields = append(fields, reflect.StructField{
	// 				Name: fieldName,
	// 				Type: reflect.TypeOf(v),
	// 		})
	// }

	// fmt.Println(fields)

	// // Create struct type dynamically
	// personType := reflect.StructOf(fields)

	// // Create a new instance of the struct
	// personValue := reflect.New(personType).Elem()

	// // Set values for the struct fields
	// for k, v := range data {
	// 		fieldValue := personValue.FieldByName(strings.Title(k))
	// 		if fieldValue.IsValid() {
	// 				fieldKind := fieldValue.Type().Kind()
	// 				switch fieldKind {
	// 				case reflect.String:
	// 						strValue, ok := v.(string)
	// 						if ok {
	// 								fieldValue.SetString(strValue)
	// 						}
	// 				case reflect.Int:
	// 						intValue, ok := v.(int)
	// 						if ok {
	// 								fieldValue.SetInt(int64(intValue))
	// 						}
	// 				// Handle other types as needed
	// 				}
	// 		}
	// }

	// // Print the struct
	// fmt.Println(personValue.Interface())

	err = json.Unmarshal(bytestream, &a)
	if err != nil {
		panic(err)
	}

	return a, nil
}