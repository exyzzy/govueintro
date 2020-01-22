# GoVueIntro

GoVueIntro demo code

To create the local database:

```
cd govueintro/data
createuser -P -d govueintro  *pass: govueintro*
createdb govueintro
psql -U govueintro -f setup.sql -d govueintro
cd ..
```

Config in in data/configlocaldb.json and config.json

PORT and DATABASE_URL will be set by heroku

before push to heroku:

```
go mod init
#if needed
go mod tidy
go mod vendor
go install
```