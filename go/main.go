package main

func main() {
	scans := scanInput("./2day.txt")
	scan := scans.scan
	defer scans.file.Close()

	d2 := Day2{}
	d2.part2(scan)
}
