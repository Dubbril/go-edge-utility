package services

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"os"
	"sort"
	"strings"
)

func main() {
	// Replace the following variables with your SFTP connection details
	host := "192.168.1.17"
	port := 22
	username := "dubbril"
	password := "bit@1234"
	remoteDir := "/home/dubbril/Desktop/data/"
	localDir := "C:\\Users\\dubbril\\Desktop\\remote\\"

	// Establish an SSH connection
	configx := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, port), configx)
	if err != nil {
		fmt.Println("Error connecting to SSH:", err)
		return
	}
	defer func(conn *ssh.Client) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	// Create an SFTP client
	sftpClient, err := sftp.NewClient(conn)
	if err != nil {
		fmt.Println("Error creating SFTP client:", err)
		return
	}
	defer func(sftpClient *sftp.Client) {
		err := sftpClient.Close()
		if err != nil {
			return
		}
	}(sftpClient)

	// List files in the remote directory
	files, err := sftpClient.ReadDir(remoteDir)
	if err != nil {
		fmt.Println("Error listing remote directory:", err)
		return
	}

	// Filter files starting with "EIM_EDGE_BLACKLIST"
	filteredFiles := make([]os.FileInfo, 0)
	for _, file := range files {
		if strings.HasPrefix(file.Name(), "EIM_EDGE_BLACKLIST") {
			filteredFiles = append(filteredFiles, file)
		}
	}

	// Sort filtered files by filename
	sort.Slice(filteredFiles, func(i, j int) bool {
		return strings.Compare(filteredFiles[i].Name(), filteredFiles[j].Name()) > 0
	})

	// Download the latest file
	if len(filteredFiles) > 0 {
		latestFile := filteredFiles[0]
		remoteFilePath := remoteDir + latestFile.Name()
		localFilePath := localDir + latestFile.Name()

		// Open the remote file for reading
		remoteFile, err := sftpClient.Open(remoteFilePath)
		if err != nil {
			fmt.Println("Error opening remote file:", err)
			return
		}
		defer func(remoteFile *sftp.File) {
			err := remoteFile.Close()
			if err != nil {
				return
			}
		}(remoteFile)

		// Create the local file for writing
		localFile, err := os.Create(localFilePath)
		if err != nil {
			fmt.Println("Error creating local file:", err)
			return
		}
		defer func(localFile *os.File) {
			err := localFile.Close()
			if err != nil {
				return
			}
		}(localFile)

		// Copy the contents from the remote file to the local file
		_, err = io.Copy(localFile, remoteFile)
		if err != nil {
			fmt.Println("Error copying file contents:", err)
			return
		}

		fmt.Printf("Latest file downloaded successfully from %s to %s\n", remoteFilePath, localFilePath)
	} else {
		fmt.Println("No files found in the remote directory starting with 'EIM_EDGE_BLACKLIST'")
	}
}