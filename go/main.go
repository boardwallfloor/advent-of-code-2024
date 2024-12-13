package main

func main() {
	scans := scanInput("./4day_sample.txt")
	scan := scans.scan
	defer scans.file.Close()

	d4 := Day4{}
	d4.part1(scan)
}
