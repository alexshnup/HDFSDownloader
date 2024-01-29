package main

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/colinmarc/hdfs"
)

func main() {
	// get the HDFS namenode address from an environment variable
	namenode := os.Getenv("HDFS_NAMENODE")
	if namenode == "" {
		log.Fatal("HDFS_NAMENODE environment variable is required")
	}

	// get the HDFS path from an environment variable
	hdfsPath := os.Getenv("HDFS_PATH")
	if hdfsPath == "" {
		log.Fatal("HDFS_PATH environment variable is required")
	}

	// get the local path from an environment variable
	localPath := os.Getenv("LOCAL_PATH")
	if localPath == "" {
		log.Fatal("LOCAL_PATH environment variable is required")
	}

	// create the HDFS client
	client, err := hdfs.New(namenode)
	if err != nil {
		log.Fatal(err)
	}

	// start the recursive download
	err = downloadDir(client, hdfsPath, localPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Directory downloaded successfully")
}

func downloadDir(client *hdfs.Client, hdfsPath string, localPath string) error {
	// read the HDFS directory
	infos, err := client.ReadDir(hdfsPath)
	if err != nil {
		return err
	}

	for _, info := range infos {
		hdfsFilePath := filepath.Join(hdfsPath, info.Name())
		localFilePath := filepath.Join(localPath, info.Name())

		if info.IsDir() {
			// create the local directory
			err = os.MkdirAll(localFilePath, 0755)
			if err != nil {
				return err
			}

			// recursively download the directory
			err = downloadDir(client, hdfsFilePath, localFilePath)
			if err != nil {
				return err
			}
		} else {
			// download the file
			err = downloadFile(client, hdfsFilePath, localFilePath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func downloadFile(client *hdfs.Client, hdfsPath string, localPath string) error {
	// open the HDFS file
	r, err := client.Open(hdfsPath)
	if err != nil {
		return err
	}
	defer r.Close()

	// create the local file
	localFile, err := os.Create(localPath)
	if err != nil {
		return err
	}
	defer localFile.Close()

	// copy the HDFS file to the local file
	if _, err := io.Copy(localFile, r); err != nil {
		return err
	}

	return nil
}
