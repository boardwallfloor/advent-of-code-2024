package main

func main() {
	scans := scanInput("./4day.txt")
	scan := scans.scan
	defer scans.file.Close()

	d4 := Day4{}
	d4.part2(scan)
}
