package miniblog

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func Test_run(t *testing.T) {
	// 保存真实的标准输出
	old := os.Stdout // keep backup of the real stdout.

	// 创建一个管道
	r, w, _ := os.Pipe()
	// 将标准输出重定向到写入端
	os.Stdout = w

	// 调用待测试的函数
	err := run()

	// 关闭写入端
	w.Close()

	// 创建一个缓冲区
	var buf bytes.Buffer
	// 将管道中的数据复制到缓冲区
	io.Copy(&buf, r)

	// 恢复真实的标准输出
	os.Stdout = old // restoring the real stdout.

	if err != nil {
		// 如果待测试函数出错，则打印错误信息并返回
		t.Errorf("run() error = %v, wantErr %v", err, nil)
		return
	}

	// 比较缓冲区中的数据与期望的输出
	if got := buf.String(); got != "Hello MiniBlog!\n" {
		t.Errorf("run() = %v, want %v", got, "Hello MiniBlog!\n")
	}
}
