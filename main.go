package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/aybabtme/rgbterm"
	"github.com/hink/go-blink1"
)

type remoteTemp struct {
	Temperature float64 `json:"temperature"`
}

type colorRange struct {
	MaxTemp float64
	RGB     color.RGBA
}

var colorRanges = []colorRange{
	colorRange{
		MaxTemp: 90,
		RGB: color.RGBA{
			R: 207,
			G: 57,
			B: 39,
		},
	},
	colorRange{
		MaxTemp: 80,
		RGB: color.RGBA{
			R: 238,
			G: 126,
			B: 38,
		},
	},
	colorRange{
		MaxTemp: 70,
		RGB: color.RGBA{
			R: 253,
			G: 249,
			B: 41,
		},
	},
	colorRange{
		MaxTemp: 60,
		RGB: color.RGBA{
			R: 110,
			G: 210,
			B: 40,
		},
	},
	colorRange{
		MaxTemp: 50,
		RGB: color.RGBA{
			R: 90,
			G: 219,
			B: 140,
		},
	},
	colorRange{
		MaxTemp: 40,
		RGB: color.RGBA{
			R: 68,
			G: 184,
			B: 219,
		},
	},
	colorRange{
		MaxTemp: 30,
		RGB: color.RGBA{
			R: 75,
			G: 56,
			B: 156,
		},
	},
	colorRange{
		MaxTemp: 20,
		RGB: color.RGBA{
			R: 148,
			G: 69,
			B: 188,
		},
	},
	colorRange{
		MaxTemp: 10,
		RGB: color.RGBA{
			R: 213,
			G: 117,
			B: 219,
		},
	},
}

func main() {
	url := "http://weather.bigboy.us/weatherstation/current.json"
	if envURL := os.Getenv("TEMP_URL"); envURL != "" {
		url = envURL
	}
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("fetch error :(", err)
	}
	defer resp.Body.Close()
	asBytes, _ := ioutil.ReadAll(resp.Body)

	var obs remoteTemp

	if err := json.Unmarshal(asBytes, &obs); err != nil {
		log.Fatal("bad JSON :(", err)
	}

	c := tempToRGBA(obs.Temperature)

	device, err := blink1.OpenNextDevice()
	if err != nil {
		log.Fatal("bad device :(", err)
	}
	if err := device.SetState(blink1.State{
		Red:   c.R,
		Green: c.G,
		Blue:  c.B,
	}); err != nil {
		log.Fatal("bad light setting :( ", err)
	}

	debugColor(c)
}

func tempToRGBA(temp float64) color.RGBA {
	lower := colorRanges[len(colorRanges)-1]
	upper := colorRanges[0]
	for _, cr := range colorRanges {
		if cr.MaxTemp > temp && cr.MaxTemp < upper.MaxTemp {
			upper = cr
		}
		if cr.MaxTemp < temp && cr.MaxTemp > lower.MaxTemp {
			lower = cr
		}
	}
	pct := (temp - lower.MaxTemp) / (upper.MaxTemp - lower.MaxTemp)
	return midpointRGBA(upper.RGB, lower.RGB, pct)
}

func midpointRGBA(upper, lower color.RGBA, pct float64) color.RGBA {
	return color.RGBA{
		R: uint8((float64(upper.R)-float64(lower.R))*pct) + lower.R,
		G: uint8((float64(upper.G)-float64(lower.G))*pct) + lower.G,
		B: uint8((float64(upper.B)-float64(lower.B))*pct) + lower.B,
	}
}

func debugColor(c color.RGBA) {
	fmt.Println("Color: ", rgbterm.BgString(strconv.Itoa(int(c.R))+","+strconv.Itoa(int(c.G))+","+strconv.Itoa(int(c.B)), c.R, c.G, c.B))
}
