# FormulaOne-CLI

Simple CLI written in GoLang to keep up to date with events in Formula 1

## Install

- Download f1.exe from this repository
- Make this executables(f1.exe) globally accessible

One way to do that for Windows users is:

1. Create folder on disk C where all your simple CLI tools will be kept. For example C:\utils. In cmd it is command `md ulils`

Add that folder to PATH

1. Open the Run window (`Win + R`)
2. Type `sysdm.cpl` in Run window
3. Go to the advanced tab in the pop-up window
4. Chose Environmental Variables
5. Select “Path” variable → click Edit button 
6. Then click “New” and type path to the folder we created earlier - C:\utils or any other name you`ve chosen
7. Save all changes by clicking OK
8. Try to run in cmd `f1` command. If not working → reboot computer

## Usage

This tool supports such commands as:

- `f1` - shows quick up-to-date info about current season

![Untitled](FormulaOne-CLI%201693b79f4c35403bb5545d7a84be3b3f/Untitled.png)

- `f1 ds` - shows current drivers standings

![Untitled](FormulaOne-CLI%201693b79f4c35403bb5545d7a84be3b3f/Untitled%201.png)

- `f1 cs` - shows current constructors standings

![Untitled](FormulaOne-CLI%201693b79f4c35403bb5545d7a84be3b3f/Untitled%202.png)

- `f1 schedule` - shows schedule for current season

![Untitled](FormulaOne-CLI%201693b79f4c35403bb5545d7a84be3b3f/Untitled%203.png)

- `f1 history` - shows all time history of Formula 1

![Untitled](FormulaOne-CLI%201693b79f4c35403bb5545d7a84be3b3f/Untitled%204.png)

.

.

.

### **Credits**

This project was created by **[Ustym Hentosh](https://github.com/ustymhentosh)**.