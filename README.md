# HDFS Downloader

This application loads a directory from Hadoop (HDFS) to a local directory using the Hadoop REST API, also known as WebHDFS.

## Requirements

- Go 1.17 (minimum)
- Access to namenode HDFS
- Docker (optional)

## Usage

Set the following environment variables:

- `HDFS_NAMENODE`: the address of your HDFS namenode (for example, `namenode:8020`)
- `HDFS_PATH`: path to the directory on HDFS you want to load (e.g. `/path/to/hdfs/directory`)
- `LOCAL_PATH`: path to the local directory where you want to load the HDFS directory (e.g. `/path/to/local/directory`)


## Usage (in Docker)

Build the Docker image using the following command:

```shell
docker build -t hdfs-downloader .
```

Then start the Docker container with the following command, replacing the environment variables with your actual values:

```shell
docker run --rm -e HDFS_NAMENODE=namenode:8020 -e HDFS_PATH=/path/to/hdfs/directory -e LOCAL_PATH=/path/to/local/directory hdfs-downloader
```

## Under the hood the following

The downloadDir function reads a directory in HDFS, creates a corresponding local directory (using os.MkdirAll to create any necessary parent directories), and then downloads each file and recursively downloads each subdirectory.

As in the previous script, replace namenode:8020 with the actual namenode address and port, and replace /path/to/hdfs/directory and /path/to/local/directory with the actual HDFS and local directory paths.

Additionally, this script does not handle the Hadoop security model and does not download files in parallel.# HDFSDownloader
