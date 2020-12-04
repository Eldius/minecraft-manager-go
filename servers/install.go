package servers

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Eldius/minecraft-manager-go/model"
	ansibler "github.com/apenella/go-ansible"
)

func Setup(server model.MCServer)  {
	ansiblePlaybookConnectionOptions := &ansibler.AnsiblePlaybookConnectionOptions{
		Connection: "local",
	}
	ansiblePlaybookOptions := &ansibler.AnsiblePlaybookOptions{
		Inventory: "127.0.0.1,",
	}
	privilegeEscalationOptions := &ansibler.AnsiblePlaybookPrivilegeEscalationOptions{
		Become:        true,
		BecomeMethod:  "sudo",
	}

	playbookFile, err := GeneratePlaybookFile()
	if err != nil {
		fmt.Println("Failed to generate playbook")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	playbook := &ansibler.AnsiblePlaybookCmd{
		Playbook:          playbookFile,
		ConnectionOptions: ansiblePlaybookConnectionOptions,
		Options:           ansiblePlaybookOptions,
		PrivilegeEscalationOptions: privilegeEscalationOptions,
		ExecPrefix:        "Go-ansible example",
	}

	err = playbook.Run()
	if err != nil {
		panic(err)
	}
}

func GeneratePlaybookFile() (path string, err error) {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		return
	}
	ioutil.WriteFile(f.Name(), []byte(`---
- name: This is a hello-world example
  hosts: all
  tasks:
    - name: Create a file called '/tmp/testfile.txt' with the content 'hello world'.
      copy:
        content: hello worldn
        dest: /tmp/testfile.txt
`), os.ModePerm)
	path = f.Name()
	return
}
