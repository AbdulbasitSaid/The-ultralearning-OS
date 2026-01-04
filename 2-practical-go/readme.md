# Practical go foundation

This repo is a documentation of lessons learned from the Arden labs **Practical go foundation** course.

## Module 1 lessons

- any folder can become a go project initialized by running the `` go mod init <project name > `` command
- you can import package using the `` import "<package name >" `` statement at the top file after the package name statement.
- to run the go file, one can simply run the ``go run <filename.go>.go`` or first build the module using ``go build <file name>`` after that execute the compiled output by simply calling the file ``./<file name>``
- for the best developer experience learn go using an IDE or a rich text editor like VScode(recommended).
- ``ls -lh`` can be used to know file and folder size and more.
- ``file <file.name>`` can be used to get file type information about a file. for example running `` file .vscode/launch.json `` will output **``vscode/launch.json: ASCII text``**

> note: Some commands only works on mac or linux

## module 2 lessons

1. String: This lesson introduced us to to how strings work in go.
    - go strings are do not have a ``.lower()`` functions or any functions like on      language (Dart, JavaScript...) instead we have to use the **strings** import for string manipulations. In the example we used the ``string.Repeat("text", <number>)`` to repeatedly print a given character the number of time we want.
    - In go, we have a built in function call ``len()`` which can be used to get the length of string. However the length given is in bytes. For example ```go
       fmt.Println(len("world🌎"))``` **would print out =>** 9 instead of perhaps 6.
    This is because the char; 🌍 is a unicode character with byte a size of 4 while each other character is one byte sized. Thus why we get 9. Tips google more about **ASCII** and **UNICODE**

    - Now in other to get the actual length of visual characters in **go**, you need to use the ``utf8`` imported module and call the ``utf8.RuneCountInString("world🌎")`` which will give you 6 *runs*. Notice how in go, you get **byte** using ``len()`` function and **Rune** using the  ``utf8.RuneCountInString``.
