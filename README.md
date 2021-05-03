# GoSDL
This project is basically a starting point for me to do SDL2-based projects in Go.

The plan is to follow a basic framework:
* The application is broken up into Pages, with one page being displayed at a time.
* Each page is made up of Components (Page itself being a special type of component)
* A component will have a specific section of the screen that it is responsible for
* A component (including Page) can contain child components, and it will be responsible
  for doing the following:
    * Calling each child component on each frame to draw it's portion of the screen
    * Calling each child component with any messages that it may have to handle
