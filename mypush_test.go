package pushtest

import (
	"fmt"
	"os/exec"
	"testing"
	"time"
)

func TestPushInit(t *testing.T) {
	t1 := time.Now()
	fmt.Printf("##### t1: %v #####\r\n", t1)
	time.Sleep(time.Second)

	t2 := time.Now()
	fmt.Printf("##### t2: %v #####\r\n", t2)

	cmd := exec.Command("git", "push", "origin", "master")
	if out, err := ParseCmdCtx(cmd); err != nil {
		t.Fatalf("Push testing preparation is failed: [Info]%v, [Error]%v", out, err)
	}

	t3 := time.Now()
	fmt.Printf("##### t3: %v #####\r\n", t3)
	d := t3.Sub(t2)
	fmt.Printf("##### time d: %v #####\r\n", d)
}

func ParseCmdCtx(cmd *exec.Cmd) (output string, err error) {
	out, err := cmd.CombinedOutput()
	output = string(out)
	return output, err
}
