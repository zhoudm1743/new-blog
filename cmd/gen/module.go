package gen

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func GenerateModule(args []string) error {
	dir, _ := os.Getwd()
	if len(args) == 0 {
		fmt.Println("Please specify the module name")
		return fmt.Errorf("please specify the module name")
	}
	moduleName := args[0]
	tempDir := path.Join(dir, "cmd", "gen", "template", "module")
	distDir := path.Join(dir, "app", moduleName)

	// 创建目标目录
	if err := os.MkdirAll(distDir, 0755); err != nil {
		return fmt.Errorf("创建目录失败: %v", err)
	}

	// 处理模块名称
	originModuleName := moduleName
	distModuleName := handleName(moduleName)

	// 遍历模板目录
	err := filepath.Walk(tempDir, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 计算相对路径
		relPath, _ := filepath.Rel(tempDir, filePath)
		targetPath := path.Join(distDir, strings.ReplaceAll(relPath, "{module_name}", moduleName))

		// 转换 .temp 扩展名为 .go
		if filepath.Ext(targetPath) == ".temp" {
			targetPath = strings.TrimSuffix(targetPath, ".temp") + ".go"
		}

		if info.IsDir() {
			return os.MkdirAll(targetPath, 0755)
		}

		// 读取模板文件
		content, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}

		// 替换模板变量
		newContent := string(content)
		newContent = strings.ReplaceAll(newContent, "{module_name}", distModuleName)
		newContent = strings.ReplaceAll(newContent, "{origin_module_name}", originModuleName)

		// 写入目标文件
		return os.WriteFile(targetPath, []byte(newContent), 0755)
	})

	if err != nil {
		return fmt.Errorf("模块生成失败: %v", err)
	}
	fmt.Printf("模块 %s 生成成功，路径: %s\n", distModuleName, distDir)
	return nil
}
