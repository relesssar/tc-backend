# REST API на GoLang, для Total Control

## Документация запросов swagger
- http://localhost:8000/swagger/index.html
- Access Group: список интерфейсов не имеющих access-group
  
- Сверка скоростей на оборудовании: на роутерах и Ониме

- Active Directory: лог файл 

### Для запуска приложения:

в терминале выполнить:
```
go env -w GO111MODULE=off
go get -u github.com/gin-gonic/gin
go get -u github.com/joho/godotenv
go get -u github.com/sirupsen/logrus github.com/dgrijalva/jwt-go github.com/gin-contrib/cors
go get -u github.com/swaggo/swag/cmd/swag
go get github.com/go-sql-driver/mysql
go get -u github.com/google/uuid github.com/jmoiron/sqlx github.com/swaggo/gin-swagger github.com/swaggo/swag
```
After running go get -v -u github.com/swaggo/swag/cmd/swag
Please check /usr/local/go/bin or $HOME/go/bin where those swag executable is present.
Make sure you have that PATH in /etc/profile or $HOME/.profile
e.g. export PATH=$PATH:$HOME/go/bin

Далее 
```
cp .env-example .env
```
Отредактируйте файл настроек .env
``` 
#Port для запуска API 
port: "9000"
```
Настроить доступ к базе данных в файле .env
```
#Настройки базы PostgreSql
DB_HOST="localhost"
DB_PORT="5433"
DB_USER="user"
DB_NAME="db"
DB_PASSWORD="pass"
DB_SSLMODE="disable"

#"debug" mode. Switch to "release" mode in production.
GIN_MODE=debug
```

Если приложение запускается впервые, необходимо применить миграции к базе данных:

```
-- Установка пакета migrate
curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | sudo bash
sudo apt install migrate
-- Запуск применения миграций
 migrate -path ./schema -database 'postgres://tl_user_db_2021:MY4y5p6h7MebBZb@localhost:5433/tl_ktc_db?sslmode=disable' up

```

Для свагера запустить в корне 
```
swag init -g cmd/main.go
```
Свагер будет доступен на http://localhost:9000/swagger/index.html

### В Cron добавить 
```
# Добьавляет в базу tc.ktc данные с роутеров и онимы
0 6 * * * python /srv/ctc_onyma/check_router_speed.py

# Серышев tc.ktc.kz, запуск проверки строк роутера с данными онимы.

# запускать после скрипта добавления строк check_router_speed.py
30 7 * * * curl -X GET "https://tc.kaztranscom.kz:9000/api/router_onyma/problem/check_router_onyma_speed" -H "accept: application/json" > /var/log/tc.kaztranscom.kz/check_router_onyma_speed_$(date +\%Y\%m\%d\%H\%M\%S).log

# удаление ексель файлов старше 1 дня
0 5 * * * find /home/ubuntu/go/src/tc_kaztranscom_backend_go/download_files/ -type f -mtime +1 -delete

#поиск исправленных строк router_onyma_speed
0 8 * * * go run /home/ubuntu/go/src/tc_kaztranscom_backend_go/cron/check_close_pros.go > /var/log/tc.kaztranscom.kz/check_close_pros_$(date +\%Y\%m\%d\%H\%M\%S).log

#############  END tc.ktc.kz ######################


```