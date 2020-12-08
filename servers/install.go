package servers

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Eldius/minecraft-manager-go/config"
	"github.com/Eldius/minecraft-manager-go/model"
	ansibler "github.com/apenella/go-ansible"
)

/*
Install installs MC on a new MC server
*/
func Install(server *model.MCServer) error {
	fmt.Println(server)

	roleFolder := filepath.Join(config.GetWorkspaceFolder(), "role", "minecraft-java-edition-ansible-role")
	if _, err := os.Stat(roleFolder); err != nil {
		fmt.Println("Creating roles folder and cloning repo...")
		_ = clone(roleFolder, config.GetAnsibleRoleGitURL())
	} else {
		fmt.Println("Role folder already exists... Starting to update repo...")
		_ = pull(roleFolder)
	}
	playbookFile, err := GeneratePlaybookFile(roleFolder)
	if err != nil {
		fmt.Println("Failed to generate the Ansible playbook\n---\n", err.Error())
		return err
	}
	return Execute(*server, playbookFile)
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
Execute executes the playbook
*/
func Execute(server model.MCServer, playbookFile string) error {
	/*
		var connType string
		switch server.ConnectionType {
		case model.LOCALConType:
			connType = "local"
		case model.SSHConnType:
			connType = "ssh"
		}
	*/
	ansiblePlaybookConnectionOptions := &ansibler.AnsiblePlaybookConnectionOptions{
		Connection: "ssh",
		PrivateKey: server.PrivateKeyPath,
		User:       server.User,
	}
	ansiblePlaybookOptions := &ansibler.AnsiblePlaybookOptions{
		Inventory: fmt.Sprintf("%s,", server.Host),
	}
	//privilegeEscalationOptions := &ansibler.AnsiblePlaybookPrivilegeEscalationOptions{}

	fmt.Println("running playbook", playbookFile)
	playbook := &ansibler.AnsiblePlaybookCmd{
		Playbook:          playbookFile,
		ConnectionOptions: ansiblePlaybookConnectionOptions,
		Options:           ansiblePlaybookOptions,
		//PrivilegeEscalationOptions: privilegeEscalationOptions,
		ExecPrefix: "Go-ansible example",
	}

	err := playbook.Run()
	if err != nil {
		fmt.Println("Failed to execute Ansible playbook\n---\n", err.Error())
		return err
	}

	return nil
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

/*
GeneratePlaybookFile generates an Ansible playbook
*/
func GeneratePlaybookFile(roleFolder string) (path string, err error) {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		return
	}
	playbookContent := fmt.Sprintf(`---
- name: This is a hello-world example
  hosts: all
  tasks:
    - name: debug host
      debug:
        msg: "{{ ansible_facts['nodename'] }}"
    - name: Install MC
      include_role:
        name: %s
      vars:
        minecraft_server_download_url: https://launcher.mojang.com/v1/objects/35139deedbd5182953cf1caa23835da59ca3d7cd/server.jar
`,
		roleFolder,
	)
	err = ioutil.WriteFile(f.Name(), []byte(playbookContent), os.ModePerm)
	if err != nil {
		return
	}
	path = f.Name()
	return
}
