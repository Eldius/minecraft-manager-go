package servers

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Eldius/minecraft-manager-go/model"
	ansibler "github.com/apenella/go-ansible"
)

/*
Install installs MC on a new MC server
*/
func Install(server *model.MCServer) error {
	return nil
}

/*
Test just a test method...
*/
func Test(server model.MCServer) {
	ansiblePlaybookConnectionOptions := &ansibler.AnsiblePlaybookConnectionOptions{
		Connection: "local",
	}
	ansiblePlaybookOptions := &ansibler.AnsiblePlaybookOptions{
		Inventory: "127.0.0.1,",
	}
	privilegeEscalationOptions := &ansibler.AnsiblePlaybookPrivilegeEscalationOptions{}

	playbookFile, err := GenerateTestPlaybookFile()
	if err != nil {
		fmt.Println("Failed to generate playbook")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("running playbook", playbookFile)
	playbook := &ansibler.AnsiblePlaybookCmd{
		Playbook:                   playbookFile,
		ConnectionOptions:          ansiblePlaybookConnectionOptions,
		Options:                    ansiblePlaybookOptions,
		PrivilegeEscalationOptions: privilegeEscalationOptions,
		ExecPrefix:                 "Go-ansible example",
	}

	err = playbook.Run()
	if err != nil {
		panic(err)
	}
}

/*
GenerateTestPlaybookFile generates a test Ansible playbook
*/
func GenerateTestPlaybookFile() (path string, err error) {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		return
	}
	err = ioutil.WriteFile(f.Name(), []byte(`---
- name: This is a hello-world example
  hosts: all
  tasks:
    - name: Create a file called '/tmp/testfile.txt' with the content 'hello world'.
      copy:
        content: hello worldn
        dest: /tmp/testfile.txt
`), os.ModePerm)
	if err != nil {
		return
	}
	path = f.Name()
	return
}
