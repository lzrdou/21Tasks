# Swift. Mobile application

Development and implementation of a mobile application for ios operating system in the Swift language.

## Contents

1. [Chapter 1](#chapter-1)
2. [Chapter 2](#chapter-2) \
    2.1. [Rules](#rules) \
    2.2. [Information](#information)
3. [Chapter 3](Chapter-3) \
    3.1. [Part 1. Description and requirements](#part-1-description-and-requirements) \
    3.2. [Part 2. Data and business logic layers implementation](#part-2-data-and-business-logic-layers-implementation) \
    3.3. [Part 3. User interface layer implementation](#part-3-user-interface-layer-implementation)

## Chapter 1

## Introduction

You will be developing and implementing a client mobile application to display information about objects using the MVVM pattern. The application will support:
  - retrieving a collection of data from the server with a list of objects
  - saving the received information into the database
  - displaying the list of objects 
  - displaying detailed information about an object

## Chapter 2

## Rules
  - Source code must be uploaded to a git repository
  - Do you have a question? Ask your neighbor on the right. Otherwise, try with your neighbor on the left.
  - Your reference guide: peers / internet / documentation

## Information


### MVVM pattern

The MVVM (Model-View-ViewModel) is an architectural pattern for designing an application. It consists of three components: business logic model (Model), view model (ViewModel) and view (View).
The bindings mechanism provides a link between the View and ViewModel components. 


#### Business logic model

Describes the data used in the application. Models can contain logic directly linked by this data, such as validation logic for model properties. At the same time, the model must not contain any logic related to data mapping and interaction with visual controls.

#### View

Defines the visual interface through which the user interacts with the application. 
The view passes the user's actions to the view model via commands.

#### View model

Links the view and the business logic model via the data binding mechanism. If property values change in the business logic model, the displayed data in the view automatically changes, although the business logic model and the view are not directly linked.

It contains the logic for retrieving data from the business logic model, which is then transferred to the view. The view model defines the logic for updating the data in the business logic model.

Visual components, such as the entry field, do not use events, and the view interacts with the view model via commands. For example, a user wants to save data entered in a text field. He clicks a button and thereby sends a command to the view model. And the view model receives the data and updates the model accordingly.

<div align="center">
  <img src="misc/images/MVVM.png"/>
</div>

## Chapter 3

### Part 1. Description and requirements

**General requirements:**
  - The program code must be located in the src folder and not contradict the requirements below
  - As a design reference, you can use [SwiftUI tutorial](https://developer.apple.com/tutorials/swiftui)
  - The application design must be intuitive

**To implement a mobile application, you need to use the following standard tools:**

  - Xcode 14.2
  - Swift 5.7
  - iOS SDK V6
  - Cocaopods 1.11.3 (dependency management)
  - SwiftUI 4 (creating a user interface)
  - Codable (JSON serialization, Swift standard library)

**The following dependencies should be added to the project:**

  - Realm 10.35.0 (working with the database)
  - Alamofire 5.6.4 (networking)
  - Swinject 2.8.3 (dependency injection)
  - RxSwift 6.5.0 (asynchronous operation)

**The application must meet the following requirements in terms of architecture:**
  - Use of the MVVM architecture pattern
  - Presence of three main application layers: data, business logic and user interface

### Part 2. Data and business logic layers implementation

**Implementation requirements:**
1. Data layer (datasource)
    - Database
    - Separate models (Entity) for the database tables must be created
    - The model name must be formed as a NameEntity (e.g. UserEntity)
    - The CRUD operations must be implemented
    - Network service
    - Separate models (Dto) with a serializer must be created
    - The model name must be formed as NameDto (e.g. UserDto)

2. Business logic layer (domain)
    - Separate models must be created (Domain)
    - The name of the model must be formed as a Name (e.g. User)
    - Mappers must be implemented for Dto<-> Domain, Entity<->Domain
    - Mappers must be formed as [extension methods](https://docs.swift.org/swift-book/LanguageGuide/Extensions.html)
    - A repository for working with the database must be implemented
    - A repository for network service must be implemented
3. Dependencies Implementation (DI)

You can check whether the data is correctly retrieved and saved by logging

### Part 3. User interface layer implementation
In this part you need to:

1. User interface layer (presentation)
    - Main screen
        - Button to update information about objects:
            - When you click the button, data about the objects should be retrieved from the server
            - Retrieving data from the server must be performed asynchronously and not block the main thread
            - If there is an error when retrieving data from the server, a pop-up notification must appear on the screen
            - Data retrieved from the server must be saved to the database
            - Operation of saving data from the server must be performed asynchronously and not block the main thread
        - A list with brief information about the objects:
            - The data to be displayed must be taken from the database
            - Operation of retrieving information from the database must be performed asynchronously and not block the main thread
            - The list element must contain several text fields and a picture
            - Clicking on the list item must take you to a screen with details of the object
        - Horizontal and vertical orientation of the device must be supported
    - Screen with object details
        - Detailed information about the object:
            - Full information about the object must be displayed
            - Data must be downloaded from the network to the database
            - The operation of downloading data from the network to the database must be asynchronous and not block the main thread
            - The data to be displayed must be retrieved from the database
            - Operation of retrieving data from the database must be performed asynchronously and not block main thread
            - It should be possible to scroll through the displayed data if not all of the content fits on the screen
            - Horizontal and vertical orientation of the device must be supported
    - Separate models (ViewData) must be created
    - The model name must be formed as NameViewData (e.g., UserViewData)
    - ViewData<->Domain mappers must be implemented
    - Mappers must be formed as [extension methods](https://docs.swift.org/swift-book/LanguageGuide/Extensions.html)
2. Provide communication between the user interface layer and the business logic layer
