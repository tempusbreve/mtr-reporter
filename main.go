package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"time"

	"github.com/pkg/errors"
)

var (
	destFlag    string
	urlFlag     string
	dbFlag      string
	verboseFlag bool
)

func init() {
	flag.StringVar(&destFlag, "destination", "1.1.1.1", "Destination to report on")
	flag.StringVar(&urlFlag, "url", "http://localhost:8086", "InfluxDB reporting URL")
	flag.StringVar(&dbFlag, "db", "db0", "InfluxDB DB")
	flag.BoolVar(&verboseFlag, "verbose", false, "Verbose logging")
}

func main() {
	flag.Parse()

	res, err := collectMetrics(destFlag)
	if err != nil {
		log.Fatal(err)
	}
	if verboseFlag {
		log.Printf("Collect %q:\n%s\n", destFlag, res.Report.String())
	}

	if err = publishMetrics(urlFlag, dbFlag, res); err != nil {
		log.Fatal(err)
	}
}

func publishMetrics(influxURL string, influxDB string, result *mtrResult) error {
	u := fmt.Sprintf("%s/write?db=%s&precision=s", influxURL, influxDB)
	contentType := "text/plain"
	body := bytes.NewReader([]byte(result.Report.String()))
	_, err := http.Post(u, contentType, body)
	return err
}

func collectMetrics(host string) (*mtrResult, error) {
	var (
		err  error
		sout bytes.Buffer
		serr bytes.Buffer
		res  = &mtrResult{}
	)
	cmd := exec.Command("mtr", "-r", "-c", "10", "--json", host)
	cmd.Stdout = &sout
	cmd.Stderr = &serr
	if err = cmd.Run(); err != nil {
		return nil, errors.Wrap(err, serr.String())
	}
	if err = json.Unmarshal(sout.Bytes(), &res); err != nil {
		return nil, errors.Wrap(err, serr.String())
	}
	res.Report.Timestamp = time.Now().Unix()
	return res, nil
}
