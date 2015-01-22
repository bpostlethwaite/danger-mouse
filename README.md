# danger-mouse
state mutating server for stress testing your monitoring setup
![image](http://www.whatsupwhatson.com/wp-content/uploads/2014/12/danger-mouse.jpg)

## Install
```shell
go get github.com/bpostlethwaite/danger-mouse
go install
```
**danger-mouse** requires a configuration file `danger.json` in one of these directories:

1.  `./`
2.  `/etc`
3.  `/var/lib/danger/`

for the following example we will copy the default config into `/etc`
```shell
cp $GOPATH/src/github.com/bpostlethwaite/danger-mouse/danger.json /etc/danger.json
```
finally start the daemon with (provided you have included your `GOPATH/bin` in your path) with
```
danger-mouse
```

### permissions
If the configuration file or the working directory of **danger-mouse** is set in a file with higher than user permissions you will have to elevate permissions with
```shell
sudo danger-mouse
```

For a more concise start script symlink **danger-mouse** to **danger** with
```shell
ln -s $GOPATH/bin/danger-mouse /usr/local/bin/danger
```

## Usage
the following assumes you have symlinked to **danger**

### Daemon control
#### start the daemon process
```shell
danger
```

#### stop the daemon process
```shell
danger -s stop
```

### Danger Commands
#### memup
Increase the memory of the server process by n *mb*
```shell
danger memup 60
```

#### memdown
Decrease the memory of the server back to running requirements
```shell
danger memdown
```

#### ping
Change the returned status code of the url `ping/`
```shell
danger ping 500
```

#### cpu
Burn 100% CPU for n *seconds*
```shell
danger cpu 30
```

#### dbup
Set the file size of the db located at (the configurable) `work-dir/db-file` by n `mb`
```shell
danger dbup 100
```

#### dbdown
Truncate the file size of the db file to 0 bytes
```shell
danger dbdown
```

## License MIT
