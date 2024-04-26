# WAILS-CMDREF

## About

A desktop gui app to store and manage a reference of CLI commands.
Build a centralized reference for CLI commands of any type with name, program, and description.
Use the built in fuzzy finder to filter commands easily.

It uses SMUI, light/dark themes and transparent mica option on Windows to create a visually impressive app.

## TODO:
- Add Fuzzy finder logic
- Add Tooltips to buttons
- Add New Command Modal Dialog Form and connect to Go function to create and save new command.
- Add Edit Command Modal Dialog Form and connect to Go function to save the command.
- Add Delete Command Modal Confirmation Dialog and connect to Go function to remove command from file.
- Add Import button at the top to import from .json file.

## Screenshots

![Example1](/screenshots/Example1.png)
![Example2](/screenshots/Example2.png)

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.
