package main
//必填字段
 if len(r.Form("username")[0]) == 0{
 	//为空的处理
}

//数字
getnum,err := strconv.Atoi(r.Form.Get("age"))
 if err != nil {

}
 if getnum > 100 {
 	//不太现实
}
//对数字的正则表达：
if m ,_ := regexp.MatchString("^[0-9]+$",r.Form.Get("age")); !m{
	return false
}

//中文
if m,_ := regexp.MatchString("^[\\x{4e00}-\\x{9fa5}]+$",r.Form.Det("realname")); !m{
	return false
}
//英文
if m,_ := regexp.MatchString("^[a-zA-Z],r.Form.Det("engname")); !m{
	return false
}
//

