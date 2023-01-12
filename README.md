# BU-app
This project was created to read and store data from Ballet Reports(Boletins de Urna) reading QR Codes.

## Dependencies
- [html5-qrcode](https://github.com/mebjas/html5-qrcode)
- [PocketBase](https://pocketbase.io/)

## Usage
Simple run:
```
go build -o buapp main.go // Use buapp.exe in Windows instead
./buapp serve
```
PocketBase will take care of the authentication and storage, read the [docs](https://pocketbase.io/docs/) for more information

## Disclaimer
This code was never used in a real election, just on examples. For more information on the subject you can read the [official guides](https://www.tse.jus.br/eleicoes/eleicoes-2018/qr-code-no-boletim-de-urna-manual-para-criacao-de-aplicativos-de-leitura)