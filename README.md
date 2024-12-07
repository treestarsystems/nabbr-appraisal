# Nabbr-Appraisal

Nabbr Appraisal Web Application

## Requirements

- [GoLang 1.23.1 or higher](https://go.dev/doc/install)
- [MongoDB 7.0.14 or higher (optional)](https://www.mongodb.com/docs/manual/administration/install-community/)
- Properly formatted .env file (see below)

## Installation

### Backend

```bash
git clone https://github.com/treestarsystems/nabbr-appraisal.git
cd ./nabbr-appraisal/
go build -o <destination dir>/nabbr-appraisal
cd <destination dir>
```

### Frontend

```bash
cd ./frontend
npm install
npm run build
```

To run the backend app it requires an .env file in the same directory or you must specify a path (Ex: -e=/path/to/.env) to a .env file.

The ./backend/.env file should be formatted as follows (template included repository):

```bash
# .env
PORT="8081"
GIN_MODE="release"
DB_NAME="nabbrAppraisal"
DB_TABLE_NAME_APPRAISALS="appraisals"
DB_MONGODB_URI="mongodb://localhost:27017"
AUTH_SECRET_KEY = ""
AUTH_REGISTRATION_KEY_ADMIN = ""
AUTH_REGISTRATION_KEY_APPRAISER = ""
AUTH_REGISTRATION_KEY_PETOWNER = ""
```

The ./frontend/.env file should be formatted as follows (template included repository):

```bash
VITE_BACKEND_API_BASE_URL="http://localhost:8081"
```

## Usage

### Backend

Without -e flag, the program will expect the .env file to be in the same directory.

```bash
<destination dir>/nabbr-appraisal
```

With -e flag, the program will use the .env file at the specified path

```bash
<destination dir>/nabbr-appraisal -e=/path/to/.env
```

### Frontend

```bash
cd ./frontend
mv dist <destination dir>
```
