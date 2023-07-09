# Code Organization

- Some codebases are easy to navigate
- Others are hard to follow
- Code organization is often the difference

A good code structure will:

- Make it easier to guess where bugs are
- Make it easier to add new features
- And more

## I am going to use MVC

It isn't perfect, but it:

- Is predictable
- Will provide with some clear rules to follow when organizing code
- It will prevent from getting stuck trying to decide how to organize code
- Is easy to pickup and utilize when learning. (Many others are not and are better suited to experienced developers and teams)

Later I can explore how I might refactor to something different from MVC _when I understand the needs better_

## Model-View-Controller (MVC)

`Model-View-Controller` aka MVC is an architectural pattern for code.

Organizes code based on responsibilities. Often called "concerns" (separation of concerns).

Developers are still ultimately responsible for putting code into the right location, so differing opinions and mistakes can occur.

Three distinct roles of MVC:

**Models** are about data, logic, and rules.

This typically means interacting with your database, but it could also mean interacting with data that comes from other services or APIs.

Often includes validating data, normalizing it, etc.

For example, this web application is going to have user accounts, and logic for validating passwords and authenticating users will all be in the models package.

**Views** render data.

In this case, we are rendering HTML.

An API could use MVC and the views could be responsible for generating JSON.

As little logic as possible. Only logic required to render data.

eg:

- "if next page exists, show next page link" is okay
- logic to calculate a bunch of graphs should probably be handled elsewhere, and then passed into a view as raw data to render.

Too much logic in views makes code very hard to maintain.

In my apps I also like to have common layouts in my views package. Eg a "theme" with some shared elements, like a navbar.
Not a requirement of MVC, but not uncommon since it relates to rendering.

**Controllers** connects it all.

It won't directly render HTML, it won't directly touch the DB, but it will call code from the models and views packages to do these things.

You can think of controllers as your [air traffic controllers](https://en.wikipedia.org/wiki/Air_traffic_control). Air traffic controllers are the people that inform each plane at an airport where to fly, when to land, and on which runway to land. They don't actually do any piloting, but instead are in charge of telling everyone what to do so that an airport can operate smoothly.

Similarly, your controller shouldn't have too much logic in it, but will instead pass data around to different pieces of your application that actually handle performing whatever work needs done.

Will start placing pretty much all of the code in these packages, but long term it might make sense to use MVC with other packages as well.
