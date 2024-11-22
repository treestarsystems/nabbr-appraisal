# Nabbr-Appraisal

Nabbr Appraisal Web Application

## Requirements

- [GoLang 1.23.1 or higher](https://go.dev/doc/install)
- [MongoDB 7.0.14 or higher (optional)](https://www.mongodb.com/docs/manual/administration/install-community/)
- Properly formatted .env file (see below)

## Installation

This will build the binary in the destination directory. I suggest you run this app in it's own directory as it will create a SQLite database file in the same directory.

##### Note: This may change in the future to allow for a specified path for the SQLite database file.

```bash
git clone https://github.com/treestarsystems/nabbr-appraisal.git
cd ./nabbr-appraisal/
go build -o <destination dir>nabbr-appraisal
cd <destination dir>
```

To run this app it requires an .env file in the same directory or you must specify a path (Ex: -e=/path/to/.env) to a .env file.

The .env file should be formatted as follows (template included repository):

```bash
# .env
PORT="8081"
GIN_MODE="release"
DB_NAME="nabbrAppraisal"
DB_TABLE_NAME="appraisals"
DB_MONGODB_URI="mongodb://localhost:27017"
```

## Usage

Without -e flag, the program will use the .env file in the same directory

```bash
./nabbr-appraisal
```

With -e flag, the program will use the .env file at the specified path

```bash
./nabbr-appraisal -e=/path/to/.env
```
