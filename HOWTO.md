# How to Use This Framework

The purpose of this library/framework is to form the foundation for developing applications with Go and SDL2.  Look at BUILD.md to figure out how to setup the environment.  As of this writing, this has only been tested on Windows, though I've done nothing that is Windows-specific so there no specific reason it shouldn't work on Mac or Linux.

# main.go

You can do whatever is needed here, but to work with the library you should do the following at a minimum:
1. Call `component.SetupSDL()` to initialize the SDL library
2. Create a `component.Window` (this could be done in the game controller)
3. Create a game instance (this could also be done in the game controller)
4. Create a game controller instance
5. call `Run()` on the game controller 

# App folder

The intention is that all application-specific source files will go in this directory in the following subdirectories:
* components - for custom components
* pages - for page objects
The application's game object can go in the root `app` folder.

# GameController

The game controller is the object that manages the game model and the game interface.  Developers should create their own game controller, but it should extend the GameController.  The only thing that the developer needs to do is provide a factory method to instantiate the game controller, giving it the game and window object.  In addition, the game controller should register all the pages that will be needed in the game.  The rest of the controller logic (specifically the basic game loop) will be provided by the library's GameController.

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

The basic method for drawing a component is the `Draw()`.  However, in general, classes that derive from `BaseComponent` should provide a boilerplate implementationof this method (see example in `RectangleComponent.go`).  Instead, the developer should put his actual drawing logic in `DrawComponent()`.  The boilerplate `Draw()` method will call the `DrawComponent()`, and then call the `Draw()` on any children.  In this way, the parent will draw first, and the children will draw after so that they are on top of the parent component.  The order that the will draw is the order in which they are added to the parent during initialization, so if they overlap, the later children will overlay on top of the earlier child components.

## BaseComponent

This is intended to be the base of components, and handles the common functionality of all components:
* It contains members for location and size (X, Y, Width, Height)
* It conatins a list of child components

# Fonts

The library provides a class to manage fonts of type `FontsManager`.  It can be accessed globally using the global variable `resources.Fonts`.  

To use the font manager, the developer should first initialize the font manager by calling `resources.FontsInits`, passing in a reference to the `embed.FS` instance of the embedded fonts (currently `FontsManager` only supports embedded fonts using Go's 1.16+ `embed` package).  This can be done anytime after `component.SetupSDL()` is called.

The developer should then indicate which fonts they will use by calling `resources.Fonts.RegisterFont()`, passing in a key name, the full filename of the font (including relative path from root of the project), and the font size.  The key can be any arbitrary string the developer chooses, but something like `"Arial-18"` that indicates both the font and size is recommended.  Note that the same font used at different sizes will have to be registered for each size.  The font does not need to be registered at the launch of the application, though it may be convenient to do so; it only needs to be registered before first use.  Once registered, the same font/size combination will remain in memory for the life of the application.

To use a font, you can simply call `resources.Fonts.GetFont()`, passing in the key name.  Or you can render a font to a texture by calling `resources.Fonts.CreateTexture()`, passing in the message to render and the color.  Use the `Query()` method on the resultant texture to determine the size of the rendered text.

# Event handling

The component library is capable of handling keyboard and mouse events and routing them to the responsible component.

For keyboard events, the framework will pass the event to the page, which will then pass to all contained components, which should then pass to their child component, on down the tree.  This will stop as soon as one component returns `true` from the `KeyEvent()` method.

For mouse up/down events, this will similarly be passed from parent to child events, except that it will first check if the event occurred within the bounds of the components.  If it did not, then the event will not be sent to the child.  Note that, as of this writing, both up and down events are passed, so a component will receive two messages for each button click.  It is up to the component to pick which event to actually trigger the functionality, though in general it is recommended on the up event.