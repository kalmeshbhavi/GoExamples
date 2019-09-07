package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// fmt.Println(noOf(`1234567890 123456789 THESE TERMS AND CONDITIONS OF SERVICE (the Terms) ARE A LEGAL AND BINDING AGREEMENT BETWEEN YOU AND NATIONAL GEOGRAPHIC governing your use of this site, which includes but is not limited to products, software and services offered by way of the website such as the Video Player, Uploader, and other applications that link to these Terms (the Site). Please review the Terms fully before you continue to use the Site. By using the Site, you agree to be bound by the Terms. You shall also be subject to any additional terms posted with respect to individual sections of the Site. Please review our Privacy Policy, which also governs your use of the Site, to understand our practices. If you do not agree, please discontinue using the Site. National Geographic reserves the right to change the Terms at any time without prior notice. Your continued access or use of the Site after such changes indicates your acceptance of the Terms as modified. It is your responsibility to review the Terms regularly. The Terms were last updated on 18 July 2011.`, 20))
	fmt.Println(noOf("hello i am going to", 5))
	// fmt.Println(test(2, "1A 2F 1C"))
}

func test(N int, S string) int {

	seats := strings.Fields(S)

	seatMap := make(map[string]string)

	for _, seat := range seats {
		seatMap[seat] = seat
	}

	count := 0
	c1 := "A"
	c2 := "B"
	c3 := "C"
	c4 := "D"
	c5 := "E"
	c6 := "F"
	c7 := "E"
	c8 := "F"
	c9 := "J"
	c10 := "K"
	for i := 1; i <= N; i++ {
		row := make([]int, 10)

		curSeat := strconv.Itoa(i) + c1
		if _, ok := seatMap[curSeat]; ok {
			row[0] = 1
		} else {
			row[0] = 0
		}

		curSeat = strconv.Itoa(i) + c2
		if _, ok := seatMap[curSeat]; ok {
			row[1] = 1
		} else {
			row[1] = 0
		}

		curSeat = strconv.Itoa(i) + c3
		if _, ok := seatMap[curSeat]; ok {
			row[2] = 1
		} else {
			row[2] = 0
		}

		curSeat = strconv.Itoa(i) + c4
		if _, ok := seatMap[curSeat]; ok {
			row[3] = 1
		} else {
			row[3] = 0
		}

		curSeat = strconv.Itoa(i) + c5
		if _, ok := seatMap[curSeat]; ok {
			row[4] = 1
		} else {
			row[4] = 0
		}

		curSeat = strconv.Itoa(i) + c6
		if _, ok := seatMap[curSeat]; ok {
			row[5] = 1
		} else {
			row[5] = 0
		}

		curSeat = strconv.Itoa(i) + c7
		if _, ok := seatMap[curSeat]; ok {
			row[7] = 1
		} else {
			row[7] = 0
		}

		curSeat = strconv.Itoa(i) + c8
		if _, ok := seatMap[curSeat]; ok {
			row[7] = 1
		} else {
			row[7] = 0
		}

		curSeat = strconv.Itoa(i) + c9
		if _, ok := seatMap[curSeat]; ok {
			row[8] = 1
		} else {
			row[8] = 0
		}

		curSeat = strconv.Itoa(i) + c10
		if _, ok := seatMap[curSeat]; ok {
			row[9] = 1
		} else {
			row[9] = 0
		}

		if row[3]+row[4]+row[5]+row[6] == 0 && (row[2]+row[1]+row[7]+row[8] > 0) {
			count++
		} else {

			if row[1]+row[2]+row[3]+row[4] == 0 {
				count++
			}

			if row[5]+row[6]+row[7]+row[8] == 0 {
				count++
			}
		}
	}

	return count
}

func noOf(S string, K int) int {
	fields := strings.Fields(S)
	lineLength := 0
	count := 1
	for _, value := range fields {
		if len(value) > K {
			return -1
		}

		if lineLength+len(value) > K {
			count++
			lineLength = 0
		}
		lineLength = lineLength + len(value) + 1
	}
	return count
}

func noOfMsgs(str string, n int) int {
	fmt.Println(str)
	start := 0
	count := 0
	if n >= len(str) {
		return count
	}
	for start < len(str) {
		prevIndex := start
		start = start + n
		fmt.Println(str[0:start], " ind ", start)

		if start < len(str) {
			if str[start] == ' ' {
				start = start + 1
			} else if str[start+1] == ' ' {
				start = start + 2
			} else {
				for prevIndex < start && str[start] != ' ' {
					start--
				}
				if prevIndex <= start {
					count = 0
					break
				}

			}
		}

		/*if start < len(str)-1 {
			// if str[start+1] != ' ' {
			for prevIndex < start && str[start] != ' ' && str[start+1] == ' ' {
				start--
			}
			if prevIndex <= start {
				count = 0
				break
			}
			// }
		}*/
		count++
	}

	return count
}
