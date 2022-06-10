package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var daysMap = map[string]uint8{
	"MONDAY":    0,
	"TUESDAY":   1,
	"WEDNESDAY": 2,
	"THURSDAY":  3,
	"FRIDAY":    4,
	"SATURDAY":  5,
	"SUNDAY":    6,
}

var daysList = []string{
	"MONDAY",
	"TUESDAY",
	"WEDNESDAY",
	"THURSDAY",
	"FRIDAY",
	"SATURDAY",
	"SUNDAY",
}

type engineer struct {
	busy     uint8
	sessions map[uint8][]session // sessions on each day
}

type session struct {
	start int
	end   int
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	line, _ := reader.ReadString('\n')
	line = strings.TrimSuffix(line, "\n")
	engineersCount, _ := strconv.Atoi(line)

	engineers := make([]engineer, engineersCount)

	for engineerIndex := 0; engineerIndex < engineersCount; engineerIndex++ {
		// user_id
		_, _ = reader.ReadString('\n')

		// busy_day
		line, _ = reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		busy := daysMap[line]

		// number_of_meetings
		line, _ := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		sessionsCount, _ := strconv.Atoi(line)

		sessions := make(map[uint8][]session, 7)

		// iterate over meetings
		for index := 0; index < sessionsCount; index++ {
			line, _ = reader.ReadString('\n')
			line = strings.TrimSuffix(line, "\n")
			splits := strings.Split(line, " ")
			day := daysMap[splits[1]]
			start, _ := strconv.Atoi(splits[2])
			end, _ := strconv.Atoi(splits[3])
			sessions[day] = append(sessions[day], session{start: start, end: end})
		}

		engineers[engineerIndex] = engineer{busy: busy, sessions: sessions}
	}

	line, _ = reader.ReadString('\n')
	line = strings.TrimSuffix(line, "\n")
	splits := strings.Split(line, " ")
	duration, _ := strconv.Atoi(splits[1])

	for day := 0; day < 7; day++ {
		isOk := true
		for index := 0; index < engineersCount; index++ {
			if engineers[index].busy == uint8(day) {
				isOk = false
				break
			}
		}
		if !isOk {
			continue
		}

		daySessions := make([]session, 0)

		for index := 0; index < engineersCount; index++ {
			daySessions = append(daySessions, engineers[index].sessions[uint8(day)]...)
		}

		daySessions = MergeSort(daySessions) // sort daySessions

		if duration < daySessions[0].start {
			fmt.Println(daysList[day], 0, duration)
			return
		}

		for index := 0; index < len(daySessions); index++ {
			if diff := daySessions[index+1].start - daySessions[index].end; diff > duration {
				fmt.Println(daysList[day], daySessions[index].end+1, daySessions[index].end+1+duration)
				return
			}
		}

		if daySessions[len(daySessions)-1].end+1+duration > 32400000 {
			continue
		}

		fmt.Println(daysList[day], daySessions[len(daySessions)-1].end+1, daySessions[len(daySessions)-1].end+1+duration)
	}
}

// =================================================================================> SORT

func MergeSort(sessions []session) []session {
	if len(sessions) < 2 {
		return sessions
	}

	middle := int(len(sessions) / 2)
	return merge(MergeSort(sessions[middle:]), MergeSort(sessions[:middle]))
}

func merge(a, b []session) []session {
	size, i, j := len(a)+len(b), 0, 0
	result := make([]session, size)
	for k := 0; k < size; k++ {
		switch true {
		case i == len(a):
			//all the elements of a already been judged
			//assuming a and b both are sorted, this happens because
			//some cases will have not equal a and b, so it might
			// be a possibility that one array got finished earlier.
			result[k].start = b[j].start
			j += 1
		case j == len(b):
			//alll the elements of a already been judged
			//assuming a nd b both are sorted
			result[k].start = a[i].start
			i += 1
		case a[i].start > b[j].start:
			result[k].start = b[j].start
			//increment the index of b array because its present index
			//is already appended to the result array
			j += 1
		case a[i].start < b[j].start:
			//increment the index of a array because its present index
			//element is already appended to the result array
			result[k].start = a[i].start
			i += 1
		case a[i].start == b[j].start:
			//in case of equality
			result[k].start = b[j].start
			j += 1
		}
	}
	return result
}
