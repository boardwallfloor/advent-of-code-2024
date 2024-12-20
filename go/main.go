package main

func main() {
	scans := scanInput("./5day_sample.txt")
	scan := scans.scan
	defer scans.file.Close()

	d4 := Day5{}
	d4.part1(scan)
}
