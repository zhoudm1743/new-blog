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
	// 生成req和resp文件
	genReqandResp(module, groupName, serviceName)
	// 生成路由文件
	genRoutes(module, groupName, serviceName, originServiceName)
	return nil
}

func genReqandResp(module, groupName, serviceName string) {
	dir, _ := os.Getwd()
	reqFile := path.Join(dir, "app", module, "schemas", "req", fmt.Sprintf("%s.go", groupName))
	// {service_name}QueryReq {service_name}AddReq {service_name}EditReq
	// 判断文件是否存在，不存在则创建
	if _, err := os.Stat(reqFile); os.IsNotExist(err) {
		os.MkdirAll(path.Dir(reqFile), 0755)
		f, err := os.Create(reqFile)
		if err != nil {
			fmt.Println("Failed to create file", err.Error())
			return
		}
		defer f.Close()
		// reqTemplate 为读取的模板内容，
		temp := "package req\n\ntype {service_name}QueryReq struct {}\n\ntype {service_name}AddReq struct {}\n\ntype {service_name}EditReq struct {\n\tID uint `json:\"id\" form:\"id\" validate:\"required\"`\n}"
		temp = strings.ReplaceAll(temp, "{service_name}", serviceName)
		_, err = f.WriteString(temp)
		if err != nil {
			fmt.Println("Failed to write file", err.Error())
			return
		}
		fmt.Printf("Req file %s generated successfully\n", reqFile)
	} else {
		// 文件存在，则在末尾追加
		f, err := os.OpenFile(reqFile, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			fmt.Println("Failed to open file", err.Error())
			return
		}
		defer f.Close()
		temp := "\n\ntype {service_name}QueryReq struct {}\n\ntype {service_name}AddReq struct {}\n\ntype {service_name}EditReq struct {\n\tID uint `json:\"id\" form:\"id\" validate:\"required\"`\n}"
		temp = strings.ReplaceAll(temp, "{service_name}", serviceName)
		_, err = f.WriteString(temp)
		if err != nil {
			fmt.Println("Failed to write file", err.Error())
			return
		}
		fmt.Printf("Req file %s updated successfully\n", reqFile)
	}

	respFile := path.Join(dir, "app", module, "schemas", "resp", fmt.Sprintf("%s.go", groupName))
	// 判断文件是否存在，不存在则创建
	if _, err := os.Stat(respFile); os.IsNotExist(err) {
		os.MkdirAll(path.Dir(respFile), 0755)
		f, err := os.Create(respFile)
		if err != nil {
			fmt.Println("Failed to create file", err.Error())
			return
		}
		defer f.Close()
		// respTemplate 为读取的模板内容，
		temp := `package resp

type {service_name}Resp struct {}
		`
		temp = strings.ReplaceAll(temp, "{service_name}", serviceName)
		_, err = f.WriteString(temp)
		if err != nil {
			fmt.Println("Failed to write file", err.Error())
			return
		}
		fmt.Printf("Resp file %s generated successfully\n", respFile)
	} else {
		// 文件存在，则在末尾追加
		f, err := os.OpenFile(respFile, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			fmt.Println("Failed to open file", err.Error())
			return
		}
		defer f.Close()
		temp := `
type {service_name}Resp struct {}
				`
		temp = strings.ReplaceAll(temp, "{service_name}", serviceName)
		_, err = f.WriteString(temp)
		if err != nil {
			fmt.Println("Failed to write file", err.Error())
			return
		}
		fmt.Printf("Resp file %s updated successfully\n", respFile)
	}
	return
}

// genRoutes 用于生成路由文件
func genRoutes(module, groupName, serviceName string, originServiceName string) {
	dir, _ := os.Getwd()
	tempFile := path.Join(dir, "cmd", "gen", "template", "routes.temp")
	targetFile := path.Join(dir, "app", module, "routes", groupName, fmt.Sprintf("%s.go", originServiceName))
	// 创建目标文件
	os.MkdirAll(path.Dir(targetFile), 0755)
	f, err := os.Create(targetFile)
	if err != nil {
		fmt.Println("Failed to create file", err.Error())
		return
	}
	defer f.Close()
	// 读取模板内容
	temp, err := os.ReadFile(tempFile)
	if err != nil {
		fmt.Println("Failed to read template", err.Error())
		return
	}
	// 替换模板内容
	content := string(temp)
	content = strings.ReplaceAll(content, "{module_name}", handleName(module))
	content = strings.ReplaceAll(content, "{group_name}", groupName)
	content = strings.ReplaceAll(content, "{service_name}", serviceName)
	content = strings.ReplaceAll(content, "{origin_service_name}", originServiceName)
	// 写入文件
	_, err = f.WriteString(content)
	if err != nil {
		fmt.Println("Failed to write file", err.Error())
		return
	}
	fmt.Printf("Routes file %s generated successfully\n", targetFile)
	// 在路由 app/{module}/routes/routes.go 中注册路由
	registerFile := path.Join(dir, "app", module, "routes", "routes.go")
	registerContent, err := os.ReadFile(registerFile)
	if err != nil {
		fmt.Println("Failed to read register file", err.Error())
		return
	}
	// 找到 fx.Provide(NewRoutes), 并在其后添加一行
	fxIndex := strings.Index(string(registerContent), "fx.Provide(NewRoutes),")
	if fxIndex == -1 {
		fmt.Println("Failed to find fx.Provide(NewRoutes)")
		return
	}
	// 构建新的注册行（带正确缩进）格式：
	// fx.Invoke({group_name}.{origin_service_name}Routes)
	newRegistration := fmt.Sprintf("\n\tfx.Invoke(%s.%sRoutes),", groupName, serviceName)
	// 在最后一个换行+闭合括号前插入
	updatedContent := string(registerContent[:fxIndex+23])
	updatedContent += newRegistration
	updatedContent += string(registerContent[fxIndex+23:])
	// 将更新后的内容写回文件
	err = os.WriteFile(registerFile, []byte(updatedContent), 0755)
	if err != nil {
		fmt.Println("Failed to write register file", err.Error())
		return
	}
	fmt.Printf("Routes registered in %s\n", registerFile)
	return
}
