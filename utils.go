package main

import (
	"github.com/valyala/fastrand"
	"encoding/base64"
	"math/rand"
	"strings"
	"strconv"
	"unsafe"
	"bufio"
	"math"
	"time"
	"fmt"
	"os"
)

func openFile(filename string) (slice []string) {
	file, _ := os.Open(filename)
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		slice = append(slice, scan.Text())
	}
	file.Close()
	return slice
}

func createFile(filename string, slice []string) {
	file, _ := os.Create(filename)
	for i := range slice {
		fmt.Fprintln(file, slice[i])
	}
	file.Close()
}

func writeFile(filename, content string) {
	file, _ := os.OpenFile(filename, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	fmt.Fprintln(file, content)
	file.Close()
}

func chunkSlice[T any](slice []T, size int) (chunks [][]T) {
	for size < len(slice) {
		slice, chunks = slice[size:], append(chunks, slice[0:size:size])
	}
	return append(chunks, slice)
}

func createIndexSlice(slice []string) (indexes []int) {
	for i := range slice {
		indexes = append(indexes, i)
	}
	return indexes
}

func getUsernameId(username string) string {
	for i := range fullUsernames {
		parts := strings.Split(fullUsernames[i], ":")
		if parts[0] == username {
			return parts[1]
		}
	}
	return ""
}

func findClaimedSession(usernameID string) []string {
	for i := range activeSessionPool {
		bearer := strings.Split(activeSessionPool[i][4], "Bearer IGT:2:")[1]
		if strings.Contains(decodeBase64(bearer), usernameID) {
			return activeSessionPool[i]
		}
	}
	return nil
}

func findStringSplit(slice []string, str string, delimiter string, index int) []string {
	for i := range slice {
		parts := strings.Split(slice[i], delimiter)
		if strings.Contains(parts[index], str) {
			return parts
		}
	}
	return nil
}

func containsStringSplit(slice []string, str string, delimiter string, index int) bool {
	for i := range slice {
		part := strings.Split(slice[i], delimiter)[index]
		if strings.Contains(part, str) {
			return true
		}
	}
	return false
}

func shuffleSlice(slice []string) []string {
	rand.New(rand.NewSource(time.Now().Unix())).Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
	return slice
}

func splitSliceAtDelimiter(slice []string, delimiter string, index int) (splitSlice []string) {
	for i := range slice {
		splitSlice = append(splitSlice, strings.Split(slice[i], delimiter)[index])
	}
	return splitSlice
}

func removeIndex[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index + 1:]...)
}

func removeString(slice []string, str string) []string {
	for i := range slice {
		if slice[i] == str {
			return append(slice[:i], slice[i + 1:]...)
		}
	}
	return slice
}

func decodeBase64(str string) string {
	bytes, _ := base64.StdEncoding.DecodeString(str)
	return string(bytes)
}

func randomIntString(length int) string {
    bytes := make([]byte, length)    
    for i := range bytes {
        bytes[i] = byte(49 + fastrand.Uint32n(9))
    }
    return unsafe.String(&bytes[0], length)
}

func formatNumber(number int64) string {
    var in string = strconv.FormatInt(number, 10)
	var out []byte = make([]byte, len(in) + (len(in) - 1) / 3)
	for i, j, k := len(in) - 1, len(out) - 1, 0; ; i, j = i - 1, j - 1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
        if k++; (k == 3) {
			j, k = j - 1, 0
			out[j] = ','
        }
    }
}

func randomRGBColor() int {
	h := rand.Float64() * 360
	s := 0.85 + rand.Float64() * 0.15
	v := 0.85 + rand.Float64() * 0.15 
	r, g, b := hsvToRGB(h, s, v)
	return (int(r * 255) << 16) | (int(g * 255) << 8) | int(b * 255)
}

func hsvToRGB(h, s, v float64) (float64, float64, float64) {
	var r, g, b float64
	c := v * s
	x := c * (1 - math.Abs(math.Mod( h / 60.0, 2) - 1))
	m := v - c
	switch {
		case h < 60:
			r, g, b = c, x, 0
		case h < 120:
			r, g, b = x, c, 0
		case h < 180:
			r, g, b = 0, c, x
		case h < 240:
			r, g, b = 0, x, c
		case h < 300:
			r, g, b = x, 0, c
		default:
			r, g, b = c, 0, x
	}
	return r + m, g + m, b + m
}