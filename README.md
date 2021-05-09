# Go SDL Lib
This project is basically a starting point for me to do SDL2-based projects in Go.

The plan is to follow a basic framework:
* The application is broken up into Pages, with one page being displayed at a time.
* Each page is made up of Components (Page itself being a special type of component)
* A component will have a specific section of the screen that it is responsible for
* A component (including Page) can contain child components, and it will be responsible
  for doing the following:
    * Calling each child component on each frame to draw it's portion of the screen
    * Calling each child component with any messages/events that it may have to handle

This framework is written specifically for me and the way I approach projects, so I've made many choices and compromises that work for me.  For example, because of the amount of work that would be required to do otherwise, components all use absolute screen position, not relative to their parent.  So a component's drawing function must take into account the component's position or it will draw outside it's bounds.  Another example: Events that are handled are limited, and mostly driven by what I tend to use.
