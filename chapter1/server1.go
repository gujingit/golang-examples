package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex

var count int

func v1() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func v2() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", countHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func v3() {
	http.HandleFunc("/", handlerV3)
	http.HandleFunc("/count", countHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func v4() {
	http.HandleFunc("/", handlerV4)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}

func v5(){
	http.HandleFunc("/",handlerV5)
	log.Fatal(http.ListenAndServe("localhost:8080",nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	fmt.Printf("count: %d\n", count)
	mu.Unlock()
	_, _ = fmt.Fprintf(w, "URL.PATH = %q\n", r.URL.Path)
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	_, _ = fmt.Fprintf(w, "count: %d\n", count)
	mu.Unlock()
}

func handlerV3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q \n", k, v)
	}

}

func handlerV4(w http.ResponseWriter, r *http.Request) {
	lissajousV2(w)
}

func handlerV5(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm();
	if err != nil {
		fmt.Fprintf(w, "parse request %s error: %s", r.URL, err.Error())
	}
	var cycles int
	for k, v := range r.Form {
		if k == "cycles" {
			cycles, err = strconv.Atoi(v[0])
			if err != nil {
				fmt.Fprintf(w, err.Error())
			}
		}
	}
	lissajousV3(w, float64(cycles))
}

func lissajousV2(w io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	palette := []color.Color{color.White, color.Black}
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), 1)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	_ = gif.EncodeAll(w, &anim) // NOTE: ignoring encoding errors
}

func lissajousV3(w io.Writer, cycles float64) {
	if cycles == 0 {
		cycles = 5
	}
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	palette := []color.Color{color.White, color.Black}
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), 1)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	_ = gif.EncodeAll(w, &anim) // NOTE: ignoring encoding errors
}

func main() {
	//v2() //chrome访问时 count会加两次
	v5()
}
