
go get -u github.com/gin-gonic/gin@latest
go get -u github.com/google/wire@latest
go get -u gorm.io/gorm@latest
go get -u github.com/go-ozzo/ozzo-validation/v4@latest
go get -u github.com/BurntSushi/toml@latest
go get -u gorm.io/driver/mysql@latest
go get -u github.com/go-gormigrate/gormigrate/v2@latest
go get -u github.com/lib/pq@latest  
go get -u github.com/mattn/go-sqlite3@latest
go get -u gorm.io/driver/postgres@latest
go get -u gorm.io/driver/sqlite@latest
go get -u github.com/sirupsen/logrus@latest
go get -u gopkg.in/natefinch/lumberjack.v2@latest
go get -u github.com/gin-contrib/cors@latest
go get -u github.com/golang-jwt/jwt/v4@latest

❯ wire ./pkg/wire/wire.go
❯ air -c ./pkg/air/.air.toml

>  go build -o rss  ./cmd/app/


