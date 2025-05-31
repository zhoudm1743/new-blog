package gen

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func GenerateService(args []string) interface{} {
	if len(args) != 3 {
		fmt.Println("Usage: go run gen.go <module> <name> <model>")
		return nil
	}
	dir, _ := os.Getwd()
	module := args[0]
	name := args[1]
	modelName := handleName(args[2])
	moduleDir := path.Join(dir, "app", module)
	if _, err := os.Stat(moduleDir); os.IsNotExist(err) {
		fmt.Printf("Module %s does not exist\n", module)
		return fmt.Errorf("Module %s does not exist", module)
	}
	d := strings.Split(name, "/")

	groupName := d[0]
	serviceName := d[0]
	if len(d) > 1 {
		serviceName = d[1]
	}
	originServiceName := serviceName
	serviceName = handleName(serviceName)
	distDir := path.Join(moduleDir, "service", groupName)
	distFile := path.Join(distDir, fmt.Sprintf("%s.go", originServiceName))
	if _, err := os.Stat(distDir); os.IsNotExist(err) {
		os.MkdirAll(distDir, 0755)
	}
	f, err := os.Create(distFile)
	if err != nil {
		fmt.Println("Failed to create file", err.Error())
		return err
	}
	defer f.Close()
	// serviceTemplate 为读取的模板内容，
	tempDir := path.Join(dir, "cmd", "gen", "template", "service.temp")
	serviceTemplate, err := os.ReadFile(tempDir)
	if err != nil {
		fmt.Println("Failed to read template", err.Error())
		return err
	}
	// 替换模板内容
	content := string(serviceTemplate)
	content = strings.ReplaceAll(content, "{group_name}", groupName)
	content = strings.ReplaceAll(content, "{service_name}", serviceName)
	content = strings.ReplaceAll(content, "{model_name}", modelName)
	content = strings.ReplaceAll(content, "{origin_service_name}", originServiceName)
	_, err = f.WriteString(content)
	if err != nil {
		fmt.Println("Failed to write file", err.Error())
		return err
	}
	// 读入注册文件
	registerFile := path.Join(moduleDir, "service", "service.go")
	registerContent, err := os.ReadFile(registerFile)
	if err != nil {
		fmt.Println("Failed to read register file", err.Error())
		return err
	}
	/**
	// 获取注册文件内容
	var Module = fx.Options(
	    // Provide your dependencies here
	)
	// 注册服务
	需要最后一行添加：
	fx.Provide({group_name}.New{service_name}Service)
	*/
	// 解析注册文件
	closingBraceIndex := strings.LastIndex(string(registerContent), "\n)")
	if closingBraceIndex == -1 {
		return fmt.Errorf("invalid service.go format: missing closing brace")
	}

	// 构建新的注册行（带正确缩进）
	newRegistration := fmt.Sprintf("\n\tfx.Provide(%s.New%sService),", groupName, serviceName)

	// 在最后一个换行+闭合括号前插入
	updatedContent := string(registerContent[:closingBraceIndex])
	updatedContent += newRegistration
	updatedContent += string(registerContent[closingBraceIndex:])
	// 将更新后的内容写回文件
	err = os.WriteFile(registerFile, []byte(updatedContent), 0755)
	if err != nil {
		fmt.Println("Failed to write register file", err.Error())
		return err
	}
	fmt.Printf("Service %s generated successfully\n", serviceName)
	return nil
}
