// Command slider builds docs/img/slider.svg: a CSS-only slideshow of the PNG
// screenshots in docs/img, for the README.
//
// GitHub strips <style>, style attributes and JavaScript from Markdown, but it
// renders an SVG as an image and runs the CSS animations inside it. Images an
// SVG links to are blocked in that context, so every screenshot is embedded as
// a data URI.
//
// Regenerate after adding, removing or renaming a screenshot, from the repo root:
//
//	go run docs/slider/main.go
//
// Slides are shown in file-name order; prefix names (01-, 02-…) to reorder.
package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const (
	background = "#272822" // the TUI background, so letterboxing blends in
	dotIdle    = "#555555"
	dotActive  = "#00a8d6" // the TUI accent
	dotsBand   = 28        // height of the progress-dot strip under the slides
	dotGap     = 18
	dotRadius  = 4
)

type slide struct {
	name          string
	data          string
	width, height int
}

func main() {
	dir := flag.String("dir", "docs/img", "directory with the PNG screenshots")
	out := flag.String("out", "docs/img/slider.svg", "SVG file to write")
	width := flag.Int("width", 973, "slide width in pixels")
	seconds := flag.Float64("seconds", 4, "seconds each slide stays on screen")
	fade := flag.Float64("fade", 0.6, "crossfade duration in seconds")
	flag.Parse()

	slides, err := load(*dir)
	if err != nil {
		log.Fatal(err)
	}
	if len(slides) == 0 {
		log.Fatalf("no PNG files in %s", *dir)
	}

	svg := render(slides, *width, *seconds, *fade)
	if err := os.WriteFile(*out, []byte(svg), 0o644); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s: %d slides, %d KB\n", *out, len(slides), len(svg)/1024)
}

func load(dir string) ([]slide, error) {
	paths, err := filepath.Glob(filepath.Join(dir, "*.png"))
	if err != nil {
		return nil, err
	}
	sort.Strings(paths)

	var slides []slide
	for _, p := range paths {
		raw, err := os.ReadFile(p)
		if err != nil {
			return nil, err
		}
		cfg, _, err := image.DecodeConfig(strings.NewReader(string(raw)))
		if err != nil {
			return nil, fmt.Errorf("%s: %w", p, err)
		}
		slides = append(slides, slide{
			name:   filepath.Base(p),
			data:   base64.StdEncoding.EncodeToString(raw),
			width:  cfg.Width,
			height: cfg.Height,
		})
	}
	return slides, nil
}

func render(slides []slide, width int, seconds, fade float64) string {
	// Every slide is scaled to the same width; the canvas is as tall as the
	// tallest one and shorter slides are centered vertically.
	height := 0
	for _, s := range slides {
		height = max(height, s.height*width/s.width)
	}

	n := len(slides)
	cycle := seconds * float64(n)
	slot := 100 / float64(n)      // % of the cycle each slide owns
	fadePct := fade / cycle * 100 // % of the cycle a crossfade takes

	var b strings.Builder
	total := height + dotsBand
	fmt.Fprintf(&b, `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 %d %d" width="%d" height="%d">`+"\n", width, total, width, total)
	fmt.Fprintf(&b, `<style>
.s{opacity:0;animation:fade %.2fs linear infinite both}
.d{fill:%s;animation:dot %.2fs step-end infinite both}
@keyframes fade{0%%{opacity:0}%.3f%%{opacity:1}%.3f%%{opacity:1}%.3f%%{opacity:0}100%%{opacity:0}}
@keyframes dot{0%%{fill:%s}%.3f%%{fill:%s}100%%{fill:%s}}
@media (prefers-reduced-motion:reduce){.s,.d{animation:none}.s0{opacity:1}}
</style>
`, cycle, dotIdle, cycle,
		fadePct, slot, slot+fadePct,
		dotActive, slot, dotIdle, dotIdle)

	fmt.Fprintf(&b, `<rect width="%d" height="%d" rx="10" fill="%s"/>`+"\n", width, total, background)
	fmt.Fprintf(&b, `<clipPath id="c"><rect width="%d" height="%d" rx="10"/></clipPath>`+"\n", width, height)
	b.WriteString(`<g clip-path="url(#c)">` + "\n")
	for i, s := range slides {
		h := s.height * width / s.width
		delay := phase(i, seconds, fade, cycle)
		fmt.Fprintf(&b, `<image class="s s%d" style="animation-delay:%.2fs" x="0" y="%d" width="%d" height="%d" href="data:image/png;base64,%s"><title>%s</title></image>`+"\n",
			i, delay, (height-h)/2, width, h, s.data, s.name)
	}
	b.WriteString("</g>\n")

	x0 := width/2 - (n-1)*dotGap/2
	for i := range slides {
		delay := phase(i, seconds, fade, cycle)
		fmt.Fprintf(&b, `<circle class="d" style="animation-delay:%.2fs" cx="%d" cy="%d" r="%d"/>`+"\n",
			delay, x0+i*dotGap, height+dotsBand/2, dotRadius)
	}
	b.WriteString("</svg>\n")
	return b.String()
}

// phase is slide i's animation-delay. It is always negative, so every
// animation is already running at t=0: a positive delay would show the first
// keyframe until it started (every dot lit, for instance). Shifting by the fade
// makes slide 0 fully visible at t=0 instead of fading in from an empty frame.
func phase(i int, seconds, fade, cycle float64) float64 {
	return float64(i)*seconds - fade - cycle
}
