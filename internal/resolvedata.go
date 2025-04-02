package internal

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

// TableInfo 表示表的基本信息
type TableInfo struct {
	TableName    string      `json:"tableName"`
	TableComment string      `json:"tableComment"`
	Fields       []FieldInfo `json:"fields"`
}

// FieldInfo 表示字段的基本信息
type FieldInfo struct {
	FieldName    string `json:"fieldName"`
	FieldComment string `json:"fieldComment"`
	FieldType    string `json:"fieldType"`
	FieldLength  string `json:"fieldLength"`
}

func resolve(data string) TableInfo {
	// 解析 SQL 语句
	tableInfo := parseSQLCreateTable(data)

	// 转换为 JSON
	jsonData, err := json.MarshalIndent(tableInfo, "", "    ")
	if err != nil {
		fmt.Println("JSON 转换错误:", err)
		return TableInfo{}
	}

	// 输出 JSON
	fmt.Println(string(jsonData))
	return tableInfo
}

func parseSQLCreateTable(sql string) TableInfo {
	var tableInfo TableInfo

	// 提取表名
	tableNameRegex := regexp.MustCompile(`create\s+table\s+(\w+)`)
	tableNameMatches := tableNameRegex.FindStringSubmatch(sql)
	if len(tableNameMatches) > 1 {
		tableInfo.TableName = tableNameMatches[1]
	}

	// 提取表注释
	tableCommentRegex := regexp.MustCompile(`comment\s+'([^']*)'`)
	tableCommentMatches := tableCommentRegex.FindStringSubmatch(sql)
	if len(tableCommentMatches) > 1 {
		tableInfo.TableComment = tableCommentMatches[1]
	}

	// 提取字段信息
	// 首先分割出字段定义部分
	fieldsPart := sql[strings.Index(sql, "(")+1 : strings.LastIndex(sql, ")")]

	// 按逗号分割每个字段定义，但需要处理括号内的逗号
	fieldLines := splitFieldLines(fieldsPart)

	for _, line := range fieldLines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "primary key") || strings.HasPrefix(line, "key") {
			continue // 跳过空行和主键/索引定义
		}

		field := parseFieldLine(line)
		if field.FieldName != "" {
			tableInfo.Fields = append(tableInfo.Fields, field)
		}
	}

	return tableInfo
}

// 分割字段行，处理括号内的逗号
func splitFieldLines(fieldsPart string) []string {
	var result []string
	var currentLine strings.Builder
	bracketCount := 0

	for _, char := range fieldsPart {
		if char == '(' {
			bracketCount++
		} else if char == ')' {
			bracketCount--
		}

		if char == ',' && bracketCount == 0 {
			result = append(result, currentLine.String())
			currentLine.Reset()
		} else {
			currentLine.WriteRune(char)
		}
	}

	if currentLine.Len() > 0 {
		result = append(result, currentLine.String())
	}

	return result
}

func parseFieldLine(line string) FieldInfo {
	var field FieldInfo

	// 提取字段名
	parts := strings.Fields(line)
	if len(parts) < 2 {
		return field
	}
	field.FieldName = parts[0]

	// 提取字段类型和长度
	typeRegex := regexp.MustCompile(`(\w+)(?:\(([^)]*)\))?`)
	typeMatches := typeRegex.FindStringSubmatch(parts[1])
	if len(typeMatches) > 1 {
		field.FieldType = typeMatches[1]
		if len(typeMatches) > 2 {
			field.FieldLength = typeMatches[2]
		}
	}

	// 提取字段注释
	commentRegex := regexp.MustCompile(`comment\s+'([^']*)'`)
	commentMatches := commentRegex.FindStringSubmatch(line)
	if len(commentMatches) > 1 {
		field.FieldComment = commentMatches[1]
	}

	return field
}
