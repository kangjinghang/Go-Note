package homework

import "fmt"

func Register() {
	fmt.Println("欢迎注册!")

	for {
		var username, pwd, cpwd string
		fmt.Print("请输入用户名:")
		fmt.Scanf("%s\n", &username)
		fmt.Print("请输入密码:")
		fmt.Scanf("%s\n", &pwd)
		fmt.Print("请再次输入密码:")
		fmt.Scanf("%s\n", &cpwd)
		if username == "" || pwd == "" || cpwd == "" {
			fmt.Println("输入内容不得为空")
			continue
		}

		if pwd != cpwd {
			fmt.Println("两次输入密码不一致")
			continue
		}
		fmt.Println("注册成功!")
		break
	}
}
