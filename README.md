# Nabbr-Appraisal

Nabbr Appraisal Web Application

## Requirements

- [GoLang 1.23.1 or higher](https://go.dev/doc/install)
- [MongoDB 7.0.14 or higher (optional)](https://www.mongodb.com/docs/manual/administration/install-community/)
- Properly formatted .env file (see below)

## Installation/Build

To run the frontend and backend app it requires an .env file with your own configuration info. The backend requires the .env file be in the same directory or you must specify a path (Ex: -e=/path/to/.env) to a .env file.

```diff
- Secure the backend .env file by assigning permissions using chown/chmod so only authorized users can view the file contents.
```

The "./backend/.env" file should be formatted as follows (template included repository):

```bash
# .env
PORT="8081"
GIN_MODE="release"
DB_NAME="nabbrAppraisal"
DB_TABLE_NAME_APPRAISALS="appraisals"
DB_MONGODB_URI="mongodb://localhost:27017"
AUTH_SECRET_KEY = ""
AUTH_REGISTRATION_KEY_ADMIN = "<something here>"
AUTH_REGISTRATION_KEY_APPRAISER = "<something here>"
AUTH_REGISTRATION_KEY_PETOWNER = "<something here>"
```

The "./frontend/.env" file should be formatted as follows (template included repository):

```bash
VITE_BACKEND_API_BASE_URL="http://localhost:8081"
```

### Build - Frontend/Backend

```bash
git clone https://github.com/treestarsystems/nabbr-appraisal.git
cd ./nabbr-appraisal/
cp ./backend/.env-rename ./backend/.env
cp ./frontend/.env-rename ./frontend/.env
./scripts/build-all.sh
```

- The build script will create a dist folder in the project root directory. The dist folder contains the frontend and backend build files for you to place in your desired directory.

## Deployment

- From the project's root directory.

### Frontend

- Clear destination directory before copying new files.

```bash
rm -rf <destination dir served by web server (Ex: nginx root dir)>
cp -R ./dist/frontend/* <destination dir served by web server (Ex: nginx root dir)>
```

### Backend

#### Deployment

```bash
cp -R ./dist/backend/* <destination dir (Ex: /opt/nabbr-appraisal)>
```

#### Usage

Without -e flag, the program will expect the .env file to be in the same directory.

```bash
cd <destination dir>
nohup ./nabbr-appraisal-tool-backend-api &
```

With -e flag, the program will use the .env file at the specified path

```bash
cd <destination dir>
nohup ./nabbr-appraisal-tool-backend-api -e=/path/to/.env &
```

### Let's Encrypt SSL Certificates

```bash
# Help: https://www.digitalocean.com/community/tutorials/how-to-secure-nginx-with-let-s-encrypt-on-ubuntu-20-04
service nginx restart
apt install certbot python3-certbot-nginx
certbot --nginx -d <domain fqdn> -d <sub domain fqdn>
service nginx restart
```

### Nginx Configuration

```nginx
server {
  listen [::]:443 ssl; # managed by Certbot
  listen 443 ssl; # managed by Certbot
  ssl_certificate /etc/letsencrypt/live/<domain fqdn>/fullchain.pem; # managed by Certbot
  ssl_certificate_key /etc/letsencrypt/live/<domain fqdn>/privkey.pem; # managed by Certbot
  include /etc/letsencrypt/options-ssl-nginx.conf; # managed by Certbot
  ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem; # managed by Certbot
  server_name <sub domain fqdn>;
  location / {
    root /var/www/<sub domain folder>;
    index index.html index.htm index.nginx-debian.html;
    try_files $uri $uri/ /index.html;
  }
  location ^~ /api/ {
    proxy_set_header Host $http_host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header X-Forwarded-Proto $scheme;
    proxy_pass http://localhost:8081;
  }
}
```
