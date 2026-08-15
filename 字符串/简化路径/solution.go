package main

import "strings"

func simplifyPath(path string) string {
	paths := strings.Split(path, "/")
	var ans string
	stack := []string{}
	for _, p := range paths {
		if p != ".." && p != "." {
			if p != "" {
				stack = append(stack, p)
			}
			continue
		} else if p == ".." {
			if len(stack) != 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			continue
		}
	}
	for len(stack) != 0 {
		if len(stack) != 0 && len(stack[0]) != 0 {
			ans += "/"
			ans += stack[0]
			stack = stack[1:]
		} else {
			stack = stack[1:]
		}
	}
	if len(ans) == 0 {
		return "/"
	}
	return ans
}

func simplifyPath2(path string) string {
	parts := strings.Split(path, "/")
	stack := make([]string, 0, len(parts))

	for _, part := range parts {
		switch part {
		case "", ".":
			// 忽略空路径和当前目录
			continue

		case "..":
			// 返回上一级目录
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}

		default:
			// 普通目录名
			stack = append(stack, part)
		}
	}

	return "/" + strings.Join(stack, "/")
} //优化版本，使用 strings.Join 实现
/*
给你一个字符串 path ，表示指向某一文件或目录的 Unix 风格 绝对路径 （以 '/' 开头），请你将其转化为 更加简洁的规范路径。

在 Unix 风格的文件系统中规则如下：

一个点 '.' 表示当前目录本身。
此外，两个点 '..' 表示将目录切换到上一级（指向父目录）。
任意多个连续的斜杠（即，'//' 或 '///'）都被视为单个斜杠 '/'。
任何其他格式的点（例如，'...' 或 '....'）均被视为有效的文件/目录名称。
返回的 简化路径 必须遵循下述格式：

始终以斜杠 '/' 开头。
两个目录名之间必须只有一个斜杠 '/' 。
最后一个目录名（如果存在）不能 以 '/' 结尾。
此外，路径仅包含从根目录到目标文件或目录的路径上的目录（即，不含 '.' 或 '..'）。
返回简化后得到的 规范路径 。

思路：先用 strings.Split 将路径分割成多个部分，然后遍历每个部分，如果是 ".." 则弹出栈顶元素，如果是 "." 则忽略，否则将元素压入栈中，最后将栈中元素用 "/" 连接起来。
连接使用 strings.Join 实现，性能更高
*/
