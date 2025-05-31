package gen

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func GenerateModel(args []string) interface{} {
	dir, _ := os.Getwd()
	templateDir := path.Join(dir, "cmd", "gen", "template")
	distDir := path.Join(dir, "app", "models")
	if len(args) == 0 {
		fmt.Println("Please specify the model name")
		return fmt.Errorf("Please specify the model name")
	}
	modelTemp := path.Join(templateDir, "model.temp")
	modelContent, err := os.ReadFile(modelTemp)
	if err != nil {
		fmt.Println("Failed to read template file:", err)
		return fmt.Errorf("Failed to read template file: %v", err)
	}
	fileName := args[0] + ".go"
	modelName := args[0]
	if len(args) > 1 && args[1] != "" {
		modelName = args[1]
	}
	modelPath := path.Join(distDir, fileName)
	modelName = handleName(modelName)
	// 判断文件是否存在,不存在则创建
	if _, err := os.Stat(modelPath); os.IsNotExist(err) {
		f, err := os.Create(modelPath)
		if err != nil {
			fmt.Println("Failed to create file:", err)
			return fmt.Errorf("Failed to create file: %v", err)
		}
		defer f.Close()
		// 替换模板内容
		content := string(modelContent)
		content = strings.ReplaceAll(content, "{model_name}", modelName)
		_, err = f.WriteString(content)
		if err != nil {
			fmt.Println("Failed to write file:", err)
			return fmt.Errorf("Failed to write file: %v", err)
		}
		fmt.Println("Model file generated:", modelPath)
	} else {
		// 文件已存在,则在末尾追加
		fmt.Println("Model file already exists:", modelPath)
		f, err := os.OpenFile(modelPath, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			fmt.Println("Failed to open file:", err)
			return fmt.Errorf("Failed to open file: %v", err)
		}
		defer f.Close()
		// 如果文件中已经有了同名的模型,则返回
		if strings.Contains(f.Name(), modelName) {
			fmt.Println("Model already exists in file:", modelPath)
			return fmt.Errorf("Model already exists in file: %v", modelPath)
		}
		content := string(modelContent)
		// 追加模型内容，去除第一行的package 声明
		content = content[strings.Index(content, "\n")+1:]
		content = strings.ReplaceAll(content, "{model_name}", modelName)
		_, err = f.WriteString("\n" + content)
		if err != nil {
			fmt.Println("Failed to write file:", err)
			return fmt.Errorf("Failed to write file: %v", err)
		}
		fmt.Println("Model file appended:", modelPath)
	}
	return nil
}
