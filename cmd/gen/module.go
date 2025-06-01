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
	// 在boot/bootstarp.go中注册模块, 在fx.Invoke(bootstrap),前面插入一行:
	// {moduleName}.Module,
	bootFile := path.Join(dir, "boot", "bootstrap.go")
	content, err := os.ReadFile(bootFile)
	if err != nil {
		return fmt.Errorf("读取文件失败: %v", err)
	}
	newContent := string(content)

	// 寻找 fx.Invoke(bootstrap), 的位置
	target := "fx.Invoke(bootstrap),"
	start := strings.Index(newContent, target)
	if start == -1 {
		return fmt.Errorf("找不到 %s 位置", target)
	}

	// 计算插入位置（直接在目标语句前插入）
	insertPos := start

	// 构造带缩进的新内容（假设原内容使用4空格缩进）
	insertText := fmt.Sprintf("\t%s.Module,\n\t", originModuleName)

	// 在目标位置前插入新行
	newContent = newContent[:insertPos] + insertText + newContent[insertPos:]

	// 写入文件
	if err := os.WriteFile(bootFile, []byte(newContent), 0755); err != nil {
		return fmt.Errorf("写入文件失败: %v", err)
	}
	fmt.Printf("模块 %s 生成成功，路径: %s\n", distModuleName, distDir)
	return nil
}
