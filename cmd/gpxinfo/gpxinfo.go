package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sio4/geo/gpx"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "usage: gpxinfo filename")
		os.Exit(1)
	}

	filename := args[0]

	gx, err := gpx.ParseFile(filename)
	if err != nil {
		fmt.Printf("parse error: %v\n", err)
	}

	fmt.Println("version:", gx.Version)
	fmt.Println("creator:", gx.Creator)
	fmt.Println("metadata:")
	fmt.Println("  name:", gx.Metadata.Name)
	fmt.Println("  description:", gx.Metadata.Desc)
	fmt.Println("  author:", gx.Metadata.Author)
	fmt.Println("  copyright:", gx.Metadata.Copyright)
	fmt.Println("  time:", gx.Metadata.Time.Local())
	fmt.Println("  keywords:", gx.Metadata.Keywords)
	fmt.Println("  bounds:", gx.Metadata.Bounds)

	for _, link := range gx.Metadata.Link {
		fmt.Println("  link:", link.HREF)
	}

	for _, wpt := range gx.Wpt {
		fmt.Println("waypoint:", wpt.Name, wpt)
	}

	for i, rte := range gx.Rte {
		fmt.Printf("route #%v name: %v\n", i, rte.Name)
		fmt.Printf("route #%v comment: %v\n", i, rte.Cmt)
		fmt.Printf("route #%v description: %v\n", i, rte.Desc)
		fmt.Printf("route #%v source: %v\n", i, rte.Src)
		fmt.Printf("route #%v number: %v\n", i, rte.Number)
		fmt.Printf("route #%v type: %v\n", i, rte.Type)
		fmt.Printf("route #%v number of route points: %v\n", i, len(rte.RtePt))
	}

	for i, trk := range gx.Trk {
		fmt.Printf("track #%v name: %v\n", i, trk.Name)
		fmt.Printf("track #%v comment: %v\n", i, trk.Cmt)
		fmt.Printf("track #%v description: %v\n", i, trk.Desc)
		fmt.Printf("track #%v source: %v\n", i, trk.Src)
		fmt.Printf("track #%v number: %v\n", i, trk.Number)
		fmt.Printf("track #%v type: %v\n", i, trk.Type)
		fmt.Printf("track #%v number of track segments: %v\n", i, len(trk.TrkSeg))
	}
}
