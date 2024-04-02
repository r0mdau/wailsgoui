# wailsgoui

Experimenting [Wails.io](https://wails.io) and building a todolist.

## Features to develop

* Update item
* Show error log to the user
* Mark item done, allow to remove all items done
* Sort item alphabetically
* Add a deadline date to items
* Sort items by deadline date
* Add tags to item e.g.: #work, #market
* search tasks by name or by tag
* Save user preferences
* Ssh synchronization of datastore
* Switch between light and dark mode
* Internationalization of the app through os language preferences
* Analytics and insights into item completion rates...
* ML to predict task completion date or future tasks ?

## About

This is the official Wails Vue template.

You can configure the project by editing `wails.json`. More information about the project settings can be found
here: https://wails.io/docs/reference/project-config

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.
