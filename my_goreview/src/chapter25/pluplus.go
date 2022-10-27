//just nil Select,
// my poor English! the orignal text call it Empty select
//select 语句没有任何情况，因此它将永远阻塞导致死锁。这个程序运行会得到报错，
//所以这个代码就不运行了吧!
package main

func main() {
	select {}
}
