package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"task2/internal"
)

func main() {
	fileName := flag.String("file", "file.json", "DataJSON")
	flag.Parse()

	if *fileName == "" {
		log.Fatal("Необходимо указать имя файла для конвертации")
	}

	jsonFile, err := os.ReadFile(*fileName)
	if err != nil {
		log.Fatalf("Ошибка чтения файла: %v", err)
	}

	yamlData, err := internal.ConvertJSONtoYAML(jsonFile)
	if err != nil {
		log.Fatalf("Ошибка конвертации JSON в YAML: %v", err)
	}

	err = os.WriteFile("output.yaml", yamlData, 0644)
	if err != nil {
		log.Fatalf("Ошибка записи в файл: %v", err)
	}

	fmt.Println("Конвертация завершена. Результат сохранен в output.yaml")
}
