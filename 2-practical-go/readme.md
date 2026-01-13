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

## module 2 Data Structures & Rest APIs

1. String: This lesson introduced us to to how strings work in go.
    - go strings are do not have a ``.lower()`` functions or any functions like on      language (Dart, JavaScript...) instead we have to use the **strings** import for string manipulations. In the example we used the ``string.Repeat("text", <number>)`` to repeatedly print a given character the number of time we want.
    - In go, we have a built in function call ``len()`` which can be used to get the length of string. However the length given is in bytes. For example ```go
       fmt.Println(len("world🌎"))``` **would print out =>** 9 instead of perhaps 6.
    This is because the char; 🌍 is a unicode character with byte a size of 4 while each other character is one byte sized. Thus why we get 9. Tips google more about **ASCII** and **UNICODE**

    - Now in other to get the actual length of visual characters in **go**, you need to use the ``utf8`` imported module and call the ``utf8.RuneCountInString("world🌎")`` which will give you 6 *runs*. Notice how in go, you get **byte** using ``len()`` function and **Rune** using the  ``utf8.RuneCountInString``.
1. Calling REST APIs:
   - the http package what you need to make various HTTP requests. example ``http.Get("url")`` or ``http.NewRequest(<http method>, url, ...)`` haven't used the **http.NewRequest** yet but will take a look at it later.
   - learnt a bit about json decoding and encoding in go. you decode using ``json.Decode()`` and most time use as struct type to return the decoded data ( much magic handle by go) example:
      > json

      ``` json
      {
         "login": "octocat",
         "id": 1,
         "avatar_url": "<https://github.com/images/error/octocat_happy.gif>",
         "html_url": "<https://github.com/octocat>",
         "name": "monalisa octocat",
         "company": "GitHub",
         "blog": "<https://github.com/blog>",
         "location": "San Francisco",
         "bio": "There once was...",
         "twitter_username": "monatheoctocat",
         "public_repos": 2,
         "followers": 20,
         "following": 0,
         "created_at": "2008-01-14T04:33:35Z",
         "updated_at": "2008-01-14T04:33:35Z"
      }
      
      ```

      > go

      ```go
      func parseResponse(readCloser io.ReadCloser)(result any, err error) {
         responseData struct {
            LoginName string `json:"login"` // some magical to let the decoder know the json key is login. So that you can use what makes sense to you in your go struct.
            Id int 
            AvatarUrl string
            HtmlUrl string
            Name string
            CompanyName string `json:"company"` // same magic 🪄 here
            Blog string
            Location int
            Bio string
            TwitterUsername string
            ... // And so on: you get the gist 😉
         }

      // next we decode
         dec := json.NewDecoder(readCloser)
         if err := dec.Decode(&responseData); err != nil{
            // login to cry, panic 😱 and return the err
            ...
         }
         return responseData, err
      }
      ```

1. Working with files (basics)
   In this example we created pretend server killer that when given a file path it kills the process by getting process id and deleting the file after killing pretend program.
   - the **os** has many features and of them is Playing ▶️ with files 🗂️
   - use the ``file, err := os.Open(pipFile)`` to open a file. Use the ``file.Fscanf(<path>, <format string>, a ...any)`` to read the file.
   - Good lesson learned is that almost all **os** ops return error and the error needs to be handled correctly.
   - use the ``defer`` keyword to close open file connections example: ``defer file.Close()``
   - defer happens when function exits, no matter what ( even when a panic 😱 happens)
   - defer works at the function  level ( don't us in a for loop ➰)
   - defer works in reverse other (kind of like a Stack 📚 or think LIFO 🥞)
   - Idiomatic use: try to to first acquire a resource, check for error, defer and releases
   - The **slog** import can be used to print out Info(information) ℹ️, Warn(warding) ⚠️ and so on...
   - even if a function has a return type of *error* you can return *nil*
   > note: nil is like null in other programming languages (dart, rust)
   > note: functions returning **error** are meant to force to thing about error handling.
