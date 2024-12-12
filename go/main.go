package main

func main() {
	scans := scanInput("./3day.txt")
	scan := scans.scan
	defer scans.file.Close()

	d3 := Day3{}
	d3.part2(scan)
}
