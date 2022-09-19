package main

import (
	"encoding/csv"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
)

var DataCount = 1000000

func generateCSVData() (csvData [][]string) {
	coinList := []string{"BTC", "INR", "DENT", "CTXC", "AST"}
	typeList := []string{"Buy", "Sell"}
	headerList := []string{"UserId", "Coin", "Type", "Date", "Time", "Volume", "Rate", "Amount"}
	csvData = append(csvData, headerList)
	for i := 0; i < DataCount; i++ {
		coinData := coinList[rand.Intn(len(coinList))]
		typeData := typeList[rand.Intn(len(typeList))]
		volume := rand.Float64() * 5
		rate := (rand.Intn(20) + 1) * 100000
		amount := int64((math.Round(volume * 100)) / 100 * float64(rate))
		// volumeString := strconv.FormatFloat(volume, 'f', 2, 32)
		// rateString := strconv.Itoa(rate)
		// amountString := strconv.Itoa(int(amount))

		csvEntry := []string{
			uuid.NewString(),
			coinData,
			typeData,
			time.Now().Format("02-01-2006"),
			time.Now().Format("15:04:05"),
			strconv.FormatFloat(volume, 'f', 2, 64),
			strconv.Itoa(rate),
			strconv.Itoa(int(amount)),
		}

		csvData = append(csvData, csvEntry)
	}

	return csvData
}

func main() {
	csvData := generateCSVData()

	csvFile, err := os.Create("transaction.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvFile)

	for _, empRow := range csvData {
		err = csvwriter.Write(empRow)
		if err != nil {
			log.Fatalf("failed writing data: %s", err)
		}
	}

	csvwriter.Flush()
	csvFile.Close()
}
