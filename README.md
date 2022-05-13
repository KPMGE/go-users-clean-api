# go users clean-api

## Goal
This is a really simple project, the aim here is applying the concepts of clean architecture, design patterns,
tdd and solid principles in order to create an api which is not coupled with external agents. 


## Api structure
Basically, this api has got 3 entities: accounts, users and books. Every user is gonna get an account 
when using the application. Furthermore, each user has its properties, one of them is the books that user 
has uploaded to the api.


## How is this project structured?
Basically in this project, we have got 7 principal folders. The __tests__ folder is where i have been doing the 
tests. Naturally, as i am using TDD, i have been testing first!

The __src__ folder is where all the production code is. There we will find all the clean architecture layers plus
a main layer. The main layer is just the place where all the code is coupled. There is the place you are going 
to find a connection to a database or an external framework for example. All the remaining layers are decoupled. 

Now i am going to explain briefly what each of the layers do. If you already know about the clean architecture, this is exactly what Robert C. Marting proposes. 

#### domain
The domain layer is where the core of our application is. You shall find the tiniest pieces of the application 
there, specifically in our case, the entities of the application: Account, User and Book

#### application
the application layer is where the implementation of the useCases are gonna be found. More than that, in this
layer, we do not care about how the system is dealing persistence. So we abstract away any type of database connection for example. In order to do that, we just depend on interfaces instead of concrete implementations. So, later we can replace the implementations with the less of effort.

#### infrastructure
the infrastructure layer is where we do implement the persistence system and connection with external providers. We do that following the 'rules' defined in the domain layer. Please note that by 'rules' i mean that we implement an interface defined in the domain model, so that later we can pass in that concrete implementation to the useCase and it shall know how to deal with the data the way we want.

#### presentation
the presentation layer is where we deal with how our api is going to serve its data. In our case, we are using http request/response. But it is important to notice that we are not depending on any framework in this layer. Instead, we create a representation of what is a http request and http response. So that, our application does not care how other frameworks deal with http, we just care about how OUR api is gonna represent and deal with it. 

#### main
As i said before, the main layer is where we couple all the components together. The nicest thing in this case is that, as we have decoupled everything else, it's kind of easy to assemble all the components together. More than that, we can use some design patterns to do that, just like the factory  design patter for example. 
Finally, to serve the that over http, i am using a library of go called [Fiber](https://gofiber.io/)


## How do i use the api routes?
If you wanna know the api routes and stuff, there is a folder called __documentation__ follow the steps there and you are going to
get your docs!
