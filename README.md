<h1 align="center">✅ TestWE VM enforcer</h1>
<p align="center">Because running Windows natively is unquestionable</p>

# Installation and usage
1. Download the latest release from the [releases page](https://github.com/keyboard-slayer/testwe-vm-enforcer/releases) and select `Source code (zip)`
2. Extract the archive to your desired location (for example, your desktop)
3. Install node.js from the [official website](https://nodejs.org/)
4. Double click on `run.bat`
5. Enjoy!

# Q&A
## Why did you replace the go version with node.js?
It was brought to my attention that the go version was not working as intended,
after some investigation I discovered that TestWE was now written in javascript with
the electron framework. This mean that I have to use [Electron's ASAR](https://github.com/electron/asar) to extract the files.

## Why did you write this patcher in the first place?
Because of the infamous COVID-19 pandemic, we had to take exams from home. The software used for
my school's exams was TestWE, which is a Windows and macOS only software. As a firm GNU/Linux user,
I had to use a VM to run the software. However, the software was detecting that I was running it
in a VM and refused to start. The first patcher written in go was made for this purpose, by modifying
some regex present in the raw binary.

## What's the difference between the old patcher and this one?
Here is a list of differences between the old patcher and this one:
- The old patcher was checking for a specific installation path, this one can automatically find the installation path by using the shortcut on the desktop
- The old patcher was checking for a specific regex in the binary, this one extract the `app.asar` file and modify the `main.js` file and removes the VM detection code
