![](./flaskr/static/images/home.png)

![](./flaskr/static/images/films.png)

![](./flaskr/static/images/client.png)

# CRUD application simulating the operation of a videotape rental shop.

Project objective:

* Interaction frontend - backend - database
* Work with the database (creating, reading, updating, deleting records) and specifically one-to-many relationship
* Populating the database with data from an external site. This is an unexpected goal, I didn't plan it originally, just after creating all the functionality of the database it turned out to be necessary to fill it, and I don't want to fill it manually (I'm lazy). Implemented the ability to download the top 250 movies from [https://www.imdb.com/](https://www.imdb.com).
* Learn Flask and basically how frameworks work 
* Work with SQLAlchemy
* Write a frontend that will approximate a typical site, try to make the site look nice
* Creating your own package, working with \_\_init__.py
* Working with project architecture

# A story from a customer I came up with

We need to create a website where a movie rental store can keep track of customers and tapes. We need the ability to add a movie (also add some number of tapes of this movie to the warehouse), register a client, give the movie to the client, return the movie to the client. And output statistics, both overall and by client separately. One requirement of the customer - "I want to create a client or find by name on the site - fall into it - and there enter the id of the movie to give him a cassette"

# Development progress and related challenges

* **Frontend** - it was the longest in terms of time, I wanted to make it beautiful, but it turned out.... I am generally satisfied, but this is due to the fact that I improved my knowledge of HTML and CSS. I used new tags, made a navbar, worked with copying the basic template in other pages. About CSS, the main is FlexBox, which here used for the first time (while studying it, learned about tabular and block models of sites, old-school). On the site is realized almost everything in FlexBox or in FlexBox, within other FlexBox. Otherwise - work with positioning (position: relative / absolute), content alignment, element size by content. Very little JavaScript, the simplest functions. On the downside: the site is not adaptive for devices, I opened it on my phone, it's terrible. I didn't want to deal with it. Frontend, it's not my thing, I learned the base and it's fine. 

* **Databases and SQLAlchemy** - this was the main goal, before that I did not work with a database, within a web application, only pure SQL (learned, for understanding), but it is not very good to write queries on it, good thing there is SQLAlchemy. To create the first database and start working with it was very easy. OOP simplified the tasks a lot. Difficulties began when three different databases should be registered in one application, here I learned about tablename and bind_key, it helped. Next, the connection one to many, at first tried to do it in two, already existing tables, but later learned that you can use a separate, third, table for links and the task became easier. On pages with customers, movies and rentals there are filters that return data on specified parameters from the database. There is an output of general statistics, how many movies, how many customers, how many cassettes were issued, if there was a cassette issue or return, then changes the number of cassettes available in stock.

* **Separate service** - data validation, before a specific action with the database: 

  - it is forbidden to give one movie twice to one client 
  - it is forbidden to delete a client or a movie if there is a rental for it
  - it is forbidden to issue a movie if there are no free tapes in stock
  - it is forbidden to give out a movie if there is no such movie (suddenly!) or if there is no client with such id
  - it is forbidden to release a movie rated R or higher to a customer who is under 18 years old.
  - it is forbidden to create the same movie as two different records in the database
  - it is forbidden to create the same client twice, the name must be unique
  - it is forbidden to create a client with a phone number of inappropriate format
  - it is forbidden to create a client who is under 14 years old
  - it is forbidden to create records if you do not specify data in all fields
  - when changing the data of a client or a movie, the new information is checked in the same way as when creating a record

* **Separate service** - filling a database of movies from the top 250 site [https://www.imdb.com/](https://www.imdb.com). When I made the database, I realized that it needs to fill it, but manually score movies did not want to. So I decided to take the top movies from the stringny site. 

  - Everything started well and seemed simple, load data from the site, find the necessary information in them (take the necessary tag from the HTML, which received). I.e. make a query, get the data, process it, write it to the database. And I did that, got from the top 250 movies, the first 25 movies, the first 25, not all of them. Here I learned what is dynamic pages (content is loaded not all at once, but sequentially, depending on user actions). How to solve it? The answer I found was through Selenium (parsing dynamic pages), I installed it and started to try it, it was difficult but normal, it was interesting until I needed to look at HTML in detail to solve parsing problems. And then, in the response from the site, in the tag \<script>, I saw a JSON file, which contained all 250 movies and all the information I needed on them. Even, links to movie posters, so I decided to add pictures to my database, initially I didnt plan to, but since there were links, there was no need to download and store pictures in the database. With the receipt of JSON, I finished the service quickly, and Selenium removed from the project. 
  - Further, it was necessary to realize the possibility of pressing a button on the site, which would start sevris. To do this, created and registered a separate router, where the function is called (the function is the service, it is imported from the appropriate module), with the help of JavaScript appears a window where you need to enter the quantity for each movie that will be added to the warehouse. Here is not realized displaying the loading process, hands did not reach.   

* **Backend** - writing code in python was not difficult, well, strictly speaking, there was nothing unusual here. Working with Flask itself was also fine, everything I needed was already in the framework. For example, the ability to upload the static folder to the server. In general, I liked working with the framework. 

  - I should also mention that for the first time I myself created the logic when we on the backend side get data (either just get it or get it by certain parameters) from the database and send it to the frontend. This is when I realized what a framework is and how it greatly simplifies development by linking different aspects. 
  - I would also like to mention - function calls to create a new rental or movie in the database. You can issue a movie in a separate tab, where you need the id of the client and the id of the movie, but you can also issue a movie while in a particular client. Here we had to implement the logic for each individual case. The possibility of deleting a movie in the database with rentals is also double, you can return the movie in a separate tab with all rentals, and you can return it in a specific client. Similar story with the creation of a new movie in the database, initially it was not planned to upload movies from the site and was realized only the possibility of manual addition, this possibility remained. But with the loading of movies from the site added and pictures of movies, which were not previously planned and stored these pictures by links, not themselves in the database. If with the movies from the site pictures - were already there, how to be with the poster when manually added? I added a default image, it will be displayed if the movie was added manually and it is also stored as a link in the database, and the image itself is in the static folder on the server, ie the link in this case, points to the path to this default picture, not a link in its usual sense.
  - **I made a mistake at the very beginning**, all routers were in one file and the architecture of the project looked very bad. As the project grew and new routers were added (for the client, for the movie, for the rental), as the service for getting movies was added, as HTML templates were added, the code looked huge and awkward to navigate. I got a feel for this when I had to change variable names throughout the project and add a bit of code in different places, within the same task (I even found extensions for VSCode, Bookmarks: Bookmarks and Todo Tree) so I could find the right places faster. 
  - The architecture problem was the most significant problem for me in this project. It took all day to get it into a usable form. Difficulties arose with import, or rather with cyclic import and the fact that some files needed access to the application instance. Then I got acquainted with \_\_init__.py and the problems were solved. The architecture itself, I divided it into separate components, depending on the purpose of the files. The whole application now lies in a separate folder, inside which there is a file with configuration and application instance, routers lie separately and are divided into 4 categories (client, movie, rent and homepage), and the file to run the application lies in the root of the project.

# How to run 

You can clone the repository, activate the virtual environment, install the dependencies `pip install -r requirements.txt` and run the `app.py` file. The project uses python 3.12.1

Or via docker, the image is available on the docker hub. To download, type `docker pull evanstrein/blockbuster_vhs`.