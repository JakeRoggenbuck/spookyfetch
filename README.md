# spookyfetch
Fetch tool but spooky for [hacktoberfest](https://hacktoberfest.com/) and [hacktoberfest github topic](https://github.com/topics/hacktoberfest)

![image](https://user-images.githubusercontent.com/35516367/193947726-3b234a67-4109-4c08-91b8-a5c27b799930.png)

## Install
```
go install github.com/jakeroggenbuck/spookyfetch@latest
```

## Display automatically in October

This code will run on shell startup and check if it's October and display the spookyfetch if it is. I put this in my `.bashrc` file.

#### Bash Version

```sh
# Get the date
MONTH=$(date "+%m")
MONTH=10#${MONTH}

# Check if it's October
if [[ $MONTH -eq 10 ]]; then
	SPOOKY=1
fi

# Check if the October flag was set
if [[ $SPOOKY -eq 1 ]]; then
	# Check if the spookyfetch program is installed
	if [[ $(which spookyfetch & >/dev/null 2>&1) ]]; then
		spookyfetch
	else
		echo "Install spookyfetch and add it to your PATH"
	fi
fi
```

#### Zsh Version

```zsh
#!/bin/zsh
# Get the date
MONTH=$(date "+%m")
MONTH=10#${MONTH}

# Check if it's October
if [[ $MONTH -eq 10 ]]; then
	SPOOKY=1
fi

# Check if the October flag was set
if [[ $SPOOKY -eq 1 ]]; then
	# Check if the spookyfetch program is installed
	if command -v spookyfetch >/dev/null 2>&1; then
		spookyfetch
	else
		echo "Install spookyfetch and add it to your PATH"
	fi
fi
```
