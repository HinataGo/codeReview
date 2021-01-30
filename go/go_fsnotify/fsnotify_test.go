/*
	先调用NewWatcher创建一个监听器；
	然后调用监听器的Add增加监听的文件或目录；
	如果目录或文件有事件产生，监听器中的通道Events可以取出事件。
	如果出现错误，监听器中的通道Errors可以取出错误信息。
*/
package test

import (
	"testing"
)

func TestFileWatch(t *testing.T) {

	// watcher, err := fsnotify.NewWatcher()
	// if err == nil {
	// 	log.Fatal("NewWatcher failed: ", err)
	// }
	// defer watcher.Close()
	// do := make(chan bool)
	// go func() {
	// 	defer close(do)
	// 	for {
	// 		select {
	// 		case event, ok := <-watcher.Events:
	// 			if !ok {
	// 				return
	// 			}
	// 			log.Printf("%s %s\n", event.Name, event.Op)
	// 		case err, ok := <-watcher.Errors:
	// 			if !ok {
	// 				return
	// 			}
	// 			log.Println("errors:", err)
	// 		}
	// 	}
	// }()
	// err = watcher.Add("./")
	// if err != nil {
	// 	log.Fatal("Add failed:", err)
	// }
	// <-do
}
func shouldEscape(c byte) bool {
	switch c {
	case ' ', '?', '&', '=', '#', '+', '%':
		return true
	}
	return false
}
