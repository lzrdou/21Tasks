### Hello, everyone!

It's time to apply the skills we learned to implement a simple iOS app.
We will use our recipe reader app as the foundation of the project.

## Topics:
- network layer development
- database layer development
- implementation of UI and data reactive connection 

### Project: Yummy notes for iOS. Advanced level

On the first team day, we developed a basic interface for displaying the list of recipes and their descriptions.
Now let's bind the logic to our interface - let the user manage and work with recipes
from the network. Development will be on the client-server interaction:
- a client that requests recipes and allows you to edit them
- a server that stores recipes and handles client requests

We will use recipe API as an external data source. (URL)

**API**
- URL for a list of recipes

**Used technologies**
- CocoaPods
- SwiftUI
- SwiftRX
- Realm
- Alamofire
- Swinject
- XcTest

### Task 1. Client application development

**Project**
- Create a new project or use a project from a previous team day

**Storage**
- We will use Realm as the data storage
- To retrieve data from external sources, use Alamofire

**Logic**
- Connect the network layer with the local storage via the domain layer
- Add logic for displaying data according to the offline first principle

As a result of this task, you will have an application for displaying recipe lists with their details
based on network data.

### Task 2. Recipe editing

- Let's add a recipe editing screen

- It can be accessed from the detailed description screen

- The recipe must be saved as a local copy

### Bonus task: Editing recipe steps

- Add a recipe step editing screen

- It can be accessed from the detailed description screen

- The step must be saved as a local copy
