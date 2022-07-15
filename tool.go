package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func ReplaceVar(template string, vars map[string]string, newFile string) {
	contentByte, _ := os.ReadFile(template)
	content := string(contentByte)
	for key, value := range vars {
		key = fmt.Sprintf("{%s}", key)
		content = strings.Replace(content, key, value, -1)
	}
	os.Remove(newFile)
	file, err := os.OpenFile(newFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString(content)
	write.Flush()
}

func Html2Image(input string, output string) {
	cmd := exec.Command("wkhtmltoimage", "--quality", "100", "--disable-smart-width", "--width", "1600", "--zoom", "2", "--enable-local-file-access", input, output)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func Insert(filePath string, value string) {
	file, err := os.OpenFile(fmt.Sprintf("./db/%s", filePath), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString(value + "\n")
	write.Flush()
}

func CalcRate(filePath string, currentValue string) string {
	value := FindLatest(filePath)
	if value == "" {
		return ""
	}
	if cur, err := strconv.ParseFloat(currentValue, 64); err == nil {
		if s, err := strconv.ParseFloat(value, 64); err == nil {
			diff := cur - s
			return fmt.Sprintf("%.2f%%", diff/s*100)
		}
	}
	return ""
}

func CalcDiff(filePath string, currentValue string) string {
	value := FindLatest(filePath)
	if value == "" {
		return ""
	}
	if cur, err := strconv.ParseFloat(currentValue, 64); err == nil {
		if s, err := strconv.ParseFloat(value, 64); err == nil {
			diff := cur - s
			return fmt.Sprintf("%.2f", diff)
		}
	}
	return ""
}

func InsertAndCalcRate(slug string, key string, curValue float64) string {
	fileName := fmt.Sprintf("%s_%s", slug, key)
	curValueStr := fmt.Sprintf("%.2f", curValue)
	rate := CalcRate(fileName, curValueStr)
	Insert(fileName, curValueStr)
	return rate
}

func InsertAndCalcDiff(slug string, key string, curValue float64) string {
	fileName := fmt.Sprintf("%s_%s", slug, key)
	curValueStr := fmt.Sprintf("%.2f", curValue)
	diff := CalcDiff(fileName, curValueStr)
	Insert(fileName, curValueStr)
	return diff
}

func FindLatest(filePath string) string {
	file, err := os.OpenFile(fmt.Sprintf("./db/%s", filePath), os.O_RDONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var lineText string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText = scanner.Text()
	}
	return lineText
}

func SetFloatColor(rate string, prefix string, suffix string) string {
	var output = ""
	if strings.Contains(rate, "-") {
		output = `<label class="red">` + fmt.Sprintf(`%s%s%s`, prefix, rate, suffix) + `</label>`
	} else {
		output = `<label class="green">` + fmt.Sprintf(`%s+%s%s`, prefix, rate, suffix) + `</label>`
	}
	return output
}
