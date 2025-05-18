# My Todo CLI App
This is a simple command line tool built with Go that enables you to create a todo-list directly on your computer

### About
0. version 1.0.0
1. Built with GO, (must install Go on your machine)
2. run todo.exe depending on your operatig system 
4. ***-t*** command sets ***type*** of action (add, complete, delete or list)
5. ***-v*** sets the ***value***, must be a string or text only

### Example
1. Create a new TODO
```bash
.\todo.exe -t add -v "Go to the movies"
```

2. Complete a TODO
```bash
.\todo.exe -t complete -v "Go to the movies"
```

3. Delete a TODO
```bash
.\todo.exe -t delete -v "Go to the movies"
```

4. List all TODOS
```bash
.\todo.exe -t list
```

### Check the json file to see your data