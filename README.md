<h1 align="center">✅ TestWE VM enforcer</h1>
<p align="center">Because running Windows natively is unquestionable</p>

## Install
You'll need first an installation of the go compiler (you can find it [here](https://golang.org/)), you'll also need a git client (that you can find [here](https://git-scm.com/downloads))

Then run a command prompt and type
```powershell
▶ go get -u github.com/keyboard-slayer/testwe-vm-enforcer
```

## How to use it ?
First verify that you have installed TestWe on your machine than simply run the command `testwe-vm-enforcer`

## Why did I do that ?
In 2020 because of the Covid pendamic my college used a remote examination software called Testwe that only accepte Windows and MacOS. Me as a GNU/Linux user don't really like Microsoft Windows. And this wonderful piece of software doesn't allows us to run it inside of a virtual machine so I've decided to see how can I bypass that without breaking the EULA (for instance I couldn't reverse engineer the software)

## How did I do that ?
As I've said previously I can't reverse engineer the software, it was a lot of watching the raw binary with my bare eyes and pray for the best but it didn't take too long. While looking at the raw binary I saw something that looked like a regex and there was the name of a lot of hypervisors like Qemu and HyperV. So I've tried to change one letter in each one of theme and BINGO ! It worked ! This software just do that process automaticly (Yup super easy and didn't required a decompiler or a disassembler :wink: )
