# Croncache

Croncache is an automation tool which helps developers to cache data actively
with JSON macros in any programming language.


## Installation

You can use Croncache both as a `Go` library (If you want to change a specific module
inside it, thanks to `IoC`) or standalone application (with a `supervisor` app.).

```
go get -u -v github.com/ertgl/croncache
```

```
cd $GOPATH/src/github.com/ertgl/croncache/cmd
go build croncache.go
```



Example
-----------


Before starting to use Croncache, you need to have installed a backend that Croncache's
 `CacheEngine` can use. For now, Croncache only can run with `Redis`. But, you are free to
 initialize a custom `CacheEngine`.


#### 1. Creating config and task files


- Create a config file with the content below.

`node.json`
```
{
  "LogFilePath": "node.log",
  "TaskManager": {
    "Tasks": {
      "test": "task_test.json"
    }
  }
}
```


- Create a task file with the content below.

`task_test.json`
```
{
  "Name": "test",
  "Command": "./test.py",
  "Args": [],
  "Interval": "4s",
  "Timeout": "2s",
  "LogFilePath": "test.log",
  "IterationOnFail": 3,
  "CacheEngineModuleName": "cache_engine/radix/v1",
  "CacheEngineCredentials": {
    "Host": "localhost",
    "Port": 6379,
    "Database": "0",
    "Password": "",
    "ConnectionPoolSize": 5
  }
}
```

> A duration string is a possibly signed sequence of decimal numbers,
> each with optional fraction and a unit suffix, such as "300ms",
> "-1.5h" or "2h45m". Valid time units are "ns", "us" (or "µs"), "ms",
> "s", "m", "h".


- Create the script file to make some datas cached.

`test.py`
```
#!/usr/bin/env python
# -*- coding: UTF-8 -*-

import sys
import json
import random
import time


def main(args):
    val = random.randint(0, 50)
    sys.stdout.write(json.dumps({
        'Macros': [
            {
                'Command': 'SET',
                'Args': [
                    'rand',
                    val,
                ],
            },
            {
                'Command': 'INCR',
                'Args': [
                    'rand',
                ],
            },
            {
                'Command': 'HMSET',
                'Args': [
                    'vals',
                    [
                        'rand',
                        val,
                    ],
                ],
            },
        ],
    }))


if __name__ == '__main__':
    main(sys.argv)
```


- Make the script file executable by typing the following code.

```
chmod a+x test.py
```

#### 2. Running Croncache


- Run an instance of Croncache specifying the config file by typing the following code.

```
./croncache --config="node.json"
```


#### 3. Correct everything works


- Type the following codes.

    -
      ```
      redis-cli
      ```
    -
      ```
      HMGET vals rand
      ```
    -
      ```
      GET rand
      ```


## License

```
The MIT License (MIT)

Copyright (c) 2017 Ertuğrul Keremoğlu <ertugkeremoglu@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
