package main

import (
	"flag"
	"fmt"
	"os"
	"time"

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

	fmt.Println("name:", gx.XMLName)
	fmt.Println("version:", gx.Version)
	fmt.Println("creator:", gx.Creator)
	printMeta(gx)
	printWayPoints(gx)
	printRoutes(gx)
	printTracks(gx)
	printExtensions(gx.Extensions)
}

func printMeta(gx *gpx.GPX) {
	if gx.Metadata != nil {
		fmt.Println("metadata:")
		if gx.Metadata.Name != "" {
			fmt.Println("  name:", gx.Metadata.Name)
		}
		if gx.Metadata.Desc != "" {
			fmt.Println("  description:", gx.Metadata.Desc)
		}
		if gx.Metadata.Author != nil {
			fmt.Println("  author:", gx.Metadata.Author)
		}
		if gx.Metadata.Copyright != nil {
			fmt.Println("  copyright:", gx.Metadata.Copyright)
		}
		if !gx.Metadata.Time.Equal(time.Time{}) {
			fmt.Println("  time:", gx.Metadata.Time.Local())
		}
		if gx.Metadata.Keywords != "" {
			fmt.Println("  keywords:", gx.Metadata.Keywords)
		}
		if gx.Metadata.Bounds != nil {
			fmt.Println("  bounds:", gx.Metadata.Bounds)
		}
		for _, link := range gx.Metadata.Link {
			fmt.Printf("  link: %v: %v\n", link.Text, link.HRef)
		}
	} else {
		fmt.Println("no metadata")
	}
}

func printWayPoints(gx *gpx.GPX) {
	for _, wpt := range gx.Wpt {
		fmt.Println("waypoint:", wpt.Name, wpt)
	}
}

func printRoutes(gx *gpx.GPX) {
	for i, rte := range gx.Rte {
		fmt.Printf("route #%v name: %v\n", i, rte.Name)
		fmt.Printf("route #%v comment: %v\n", i, rte.Cmt)
		fmt.Printf("route #%v description: %v\n", i, rte.Desc)
		fmt.Printf("route #%v source: %v\n", i, rte.Src)
		fmt.Printf("route #%v number: %v\n", i, rte.Number)
		fmt.Printf("route #%v type: %v\n", i, rte.Type)
		fmt.Printf("route #%v number of route points: %v\n", i, len(rte.RtePt))
	}
}

func printTracks(gx *gpx.GPX) {
	for i, trk := range gx.Trk {
		fmt.Printf("track #%v:", i)
		if trk.Name != "" {
			fmt.Printf(" name: '%v'", trk.Name)
		}
		if trk.Type != "" {
			fmt.Printf(" type: '%v'", trk.Type)
		}
		if trk.Number != 0 {
			fmt.Printf(" number: '%v'", trk.Number)
		}
		fmt.Printf(" number of track segments: %v\n", len(trk.TrkSeg))

		if trk.Cmt != "" {
			fmt.Printf("  comment: %v\n", trk.Cmt)
		}
		if trk.Desc != "" {
			fmt.Printf("  description: %v\n", trk.Desc)
		}
		if trk.Src != "" {
			fmt.Printf("  source: %v\n", trk.Src)
		}
		for j, link := range trk.Link {
			fmt.Printf("  link #%d: %v: %v\n", j, link.Text, link.HRef)
		}
		printExtensions(trk.Extensions)

		for j, seg := range trk.TrkSeg {
			sPt := seg.TrkPt[0]
			ePt := seg.TrkPt[len(seg.TrkPt)-1]
			fmt.Printf("  track segment #%v: number of points: %v, time length: %v\n", j, len(seg.TrkPt), ePt.Time.Sub(sPt.Time))
		}

		printTrackPoints(&trk)
	}
}

func printExtensions(ext *gpx.Extensions) {
	fmt.Println("extensions:", ext)
}

func printTrackPoints(trk *gpx.Trk) {
	for j, seg := range trk.TrkSeg {
		sPt := seg.TrkPt[0]
		ePt := seg.TrkPt[len(seg.TrkPt)-1]
		fmt.Printf("track segment #%v: number of points: %v, time length: %v\n", j, len(seg.TrkPt), ePt.Time.Sub(sPt.Time))

		for k, curr := range seg.TrkPt {
			fmt.Printf("-- %#v, %#v\n", curr, curr.Extensions)
			fmt.Printf("%04d: %s %6.3fm\n",
				k+1, curr.Time.Local().Format("15:04:05"), curr.Ele)
		}
	}

}
