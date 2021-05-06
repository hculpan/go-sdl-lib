# How to Use This Framework

The purpose of this library/framework is to form the foundation for developing applications with Go and SDL2.  Look at BUILD.md to figure out how to setup the environment.  As of this writing, this has only been tested on Windows, though I've done nothing that is Windows-specific so there no specific reason it shouldn't work on Mac or Linux.

# main.go

The changes here are pretty minimal.  If you desire a default window background color other than white, you'll need to change the assignment to `windowBackground`.  Otherwise, change the `GetGame()` function to instantiate your custom game object (see below).

# App folder

The intention is that all application-specific source files will go in this directory in the following subdirectories:
* components - for custom components
* pages - for page objects
The application's game object can go in the root `app` folder.

# Game

The game object contains the basic setup and state information for the game.  The object will be instantiated at launch in `main.go` and the same instance will remain in memory until the application terminates, and there will only ever be one instance of the game.

## BaseGame

This implementation of the Game interface is provided as a convenience for the developer.  The developer can use this as an example, though the intention is to use encapsulate it into your own Game instance.  Assuming you use the BaseGame for your own Game object, in the custom game's factory function, you should make sure to call BaseGame.Initialize() and BaseGame.RegisterPages().  BaseGame.LoadPages() returns an empty list of pages; you should override this method and include your list of pages there.

## MyGame

This is a basic example implementation of the Game interface using BaseGame.  You can rename it and change the name of the object in `main.go`, extending it with your own functionality.  Or you can write your own using this as an example.

# Page

A page is screen of your application/game.  The main game screen, a settings dialog, and a quit/cancel dialog are all example of pages.  Only one page can be displayed at a time, and the active page has complete control of the drawing area of the window.  A page instance may draw to the window renderer directly, but it is intended that a page will contain a set of components that handle all the drawing.

Note that a page is created at the launch of the application and remains in memory until the application is closed, and there will only ever be one instance of each page.  

## BasePage

This contains basic functionality for a page and is intended to be the base for any pages the developer creates.  The `LoadPage()` and `UnloadPage()` are called when a page is first created and finally destroyed; they are not called whenever a page becomes active.

### MainPage and RedPage

These are basic examples of pages and can be used as the foundation for any pages the developer may create.  Look at `MyGame.Update()` for an example of how to switch pages.

# Component

A component is the basic building block of the application.  A Page (which is a component) can contain any number of components, and a component can contain any number of other components.  When instantiating a component and assigning it to a Page, you specify the location and size of the component.  Nothing in the framework checks this, and there is nothing stopping the developer from hiding a component by putting another component on top of it, or by putting part or all of a component outside the bounds of the window.  A Page or component can contain any number of instance of a component type, as the developer desires.  A component instance is created as the page is loaded, and will only be unloaded when the application terminates.  Components that are part of an inactive page will not have their `Draw()` method called and will not receive events.

It is up to the developer to decide how big or little a component is.  A Page could have only a single component that draws everything, or every UI element could be its own component.  The framework does not require either approach to be used.

The basic method for drawing a component is the `Draw()`.  However, in general, classes implementing `Component` should provide a boilerplate `Draw()` implementation that calls `BaseComponent.DrawWithChildren()`.  This will then call all child `Draw()` methods and the component's `DrawComponent()`, which should actually handle the component's drawing logic.  For components with no children, following this same pattern is recommended, but the developer may choose to provide their own implementation of `Draw()`; this is the only method called directly by the framework.

## BaseComponent

This is intended to be the base of components, and handles the common functionality of all components:
* It contains members for location and size (X, Y, Width, Height)
* It conatins a list of child components
* Using the `DrawWithChildren()` method, it calls the `Draw()` method on all child components
    * Each components `Draw()` method should call `BaseComponent.DrawWithChildren()`, and this will call each child's `Draw()` method and then the component's `DrawComponent()`