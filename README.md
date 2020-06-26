
## How to build

Prerequisites:
- java
- nodejs
- golang

Install depedencies:
```
gradlew npmInstall
gradlew goGetAll    # skip errors on missing dependencies
```

Build
```
gradlew distZip
```

Find a distributive in build/distributions/transtur_card_atm.zip

Or start directly bin/transtur_card_atm.exe

# How to develop UI

UI is located under the web folder. web/htmlpage/index.html.template is an entry page that execute a js bundle.

UI is implemented in React/TypeScript, e.g. web/src/components/AppForm.tsx

1. Start a standalone dev server to compile 
```
npm run start
```
2. Open http://localhost:8080/


## How to use

Create a file transtur_card_atm.json in the application folder.

For example:

```
{
    "debug": true,
    "card_reader_name": "Smart",
    "local_save_path": "TCReaderUpload",
    "ftp_server": "<server>:<port>",
    "ftp_upload_path": "card_data_atm",
    "ftp_user": "<user>",
    "ftp_pass": "<pass>"
}
```

*debug* - show a menu with developer tools

*card_reader_name_pattern* - Full or partial name of the card reader to use.

*ftp_server* - if not empty the program will upload on FTP, otherwise will store into *local_save_path*




