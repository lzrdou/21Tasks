# Day 03 - Databases boot camp

Introduction into a bit advanced Structured Query Language.

*Resume: Today you will continue to investigate the different possibilities , techniques and practices how to speak with  database on one SQL language*


## Contents

1. [Chapter I](#chapter-i) \
    1.1. [Preamble](#preamble)
2. [Chapter II](#chapter-ii) \
    2.1. [General Rules](#general-rules)
3. [Chapter III](#chapter-iii) \
    3.1. [Rules of the day](#rules-of-the-day)  
4. [Chapter IV](#chapter-iv) \
    4.1. [Exercise 00 - Drop a garbage](#exercise-00-drop-a-garbage)  
5. [Chapter V](#chapter-v) \
    5.1. [Exercise 01 - Let’s get hierarchy!](#exercise-01-lets-get-hierarchy)  
6. [Chapter VI](#chapter-vi) \
    6.1. [Exercise 02 - Random? Not, fibonacci!](#exercise-02-random-not-fibonacci) 
6. [Chapter VII](#chapter-vii) \
    6.1. [Exercise 03 - Hmmm, let’s innovate!](#exercise-03-hmmm-lets-innovate) 
7. [Chapter VIII](#chapter-viii) \
    7.1. [Exercise 04 - Alert! Sabotage is detected!](#exercise-04-alert-sabotage-is-detected)  
8. [Chapter IX](#chapter-ix) \
    8.1. [Exercise 05 - Return our data back](#exercise-05-return-our-data-back)  

<h2 id="chapter-i" >Chapter I</h2>
<h2 id="preamble" >Preamble</h2>

![D03_01](images/DB_Bootcamp_Day03_0.png)

**Michael Ralph Stonebraker** (born October 11, 1943) is a computer scientist specializing in database research.
Stonebraker's career can be broadly divided into two phases: his time at University of California, Berkeley when he focused on relational database management systems, and at Massachusetts Institute of Technology (MIT) where he developed more novel data management techniques such as C-Store, H-Store and SciDB. Stonebraker is currently a Professor Emeritus at UC Berkeley and an adjunct professor at MIT

Who is this person and why is he very famous in the database world and at last why did Day 01 Aliens turn to this person for help?
The answer is very simple - he is one of the evangelists to promote databases to the general public with the help of his own experience, skills and of course teaching.
In his past the future professor was impressed by the theory of relational databases from **Edgar Frank Codd** and in this regard, in 1973 he created the **Ingres** (Interactive Graphics and Retrieval System) database to demonstrate the possibility of creating a practical application of the theory of the relational model.
The key ideas from INGRES are still used in relational systems. Such as below.
-	B-tree indexes
-	primary-copy replication
-	Optimization Plans 
-	integrity constraints
-	database rules/triggers
-	performance with locking mechanism


But the most important thing is that the INGRES project has evolved into something more, namely in POST inGRES, or as we know this name in another way - **PostgreSQL**.  At the moment, there are a lot of engineers around the world who use and improve PostgreSQL from version to version, thereby creating a **PostgreSQL community**.

![D03_02](images/DB_Bootcamp_Day03_1.png)

Today, in the 20s of the 21st century, the **PostgreSQL** database is only gaining momentum and, in terms of usage trends, is included in the four world databases along with the paid and expensive Oracle and Microsoft SQL Server databases.


<h2 id="chapter-ii">Chapter II</h2>
<h2 id="general-rules" >General Rules</h2>

- Use this page as the only reference. Do not listen to any rumors and speculations on how to prepare your solution.
- Please make sure you are using the latest version of PostgreSQL.
- Please make sure you have installed and configured the latest version of Flyway by Redgate.
- Please use our [internal SQL Naming Convention rules](https://docs.google.com/document/d/1IxIOFUeb-8Z8fBOfkXiy4SkN-J1mPIXveJZUCNZFdp8/edit?usp=sharing)
- Please use our [Terms and Definitions](https://docs.google.com/document/d/1_ZTDpHcfYMASZ5BtnldurQLF0fJygGF1yuTwik0DOqk/edit?usp=sharing) document 
- That is completely OK if you are using IDE to write a source code (aka SQL script) and make a syntax check before migration at the final database solution by Flyway.
- Comments are also good in the SQL scripts. Anyway be careful with signs /*...*/ directly in SQL. These special symbols are used for Database Hints to improve SQL performance and these are not just comment marks :-).
- Pay attention to the permissions of your files and directories.
- To be assessed your solution must be in your GIT repository.
- Your solutions will be evaluated by your piscine mates.
- You should not leave in your directory any other file than those explicitly specified by the exercise instructions. It is recommended that you modify your .gitignore to avoid accidents.
- Do you have a question? Ask your neighbor on the right. Otherwise, try with your neighbor on the left.
- Your reference manual: mates / Internet / Google. 
- Read the examples carefully. They may require things that are not otherwise specified in the subject.
- And may the SQL-Force be with you!
- Absolutely everything can be presented in SQL! Let’s start and have fun!



<h2 id="chapter-iii">Chapter III</h2>
<h2 id="rules-of-the-day">Rules of the day</h2>


- Please make sure you have a separated database “data” on your PostgreSQL cluster. 
- Please make sure you have a database schema “data” in your “data” database.
- Please make sure you are working through database user “data” and password “data” with super admin permissions for your PostgreSQL cluster. 
- Each exercise of the day needs a Flyway tool for right versioning of the “data” database through user “data”.
- Please make changes in your “flyway.conf” file (located in “~/flyway-6.x.x/conf” directory) directly to configure a right and stable connection to your PostgreSQL database.

    - flyway.url = jdbc:postgresql://hostname:5432/data 
        - where hostname is DNS / IP address of PostgreSQL server 
        - where port is a port of PostgreSQL server , by default is 5432
        (jdbc:postgresql://localhost:5432/data OR  jdbc:postgresql://127.0.0.1:5432/data)
    - flyway.user = data
    - flyway.password = data
    - flyway.schemas = data
    - flyway.defaultSchema = data
- Please use the command line for Flyway to migrate changes into the database and get information about the current version from the database.
- Please don’t append additional parameters for “flyway” in a command line, all needed parameters should be changed in “flyway.conf” file
- All tasks contain a list of Allowed and Denied sections with listed database options, database types, SQL constructions etc. Please have a look at the section before you start.


<h2 id="chapter-iv">Chapter IV</h2>
<h2 id="exercise-00-drop-a-garbage">Exercise 00 - Drop a garbage</h2>

| Exercise 00: Drop a garbage |                                                                                                                          |
|-----------------------------|--------------------------------------------------------------------------------------------------------------------------|
| Turn-in directory           | ex00                                                                                                                     |
| Files to turn-in            | `V130__drop_after_refactoring.sql`                                                                                         |
| **Allowed**                     |                                                                                                                          |
| Operators                   | Standard DDL / DML operators to create / alter relations and to insert / update / delete / select data (CRUD operations) |
| SQL keywords                | <ul><li>DROP VIEW … CASCADE</li><li>DROP TABLE …</li><li>DROP VIEW</li></ul>

Aliens decided to carry out a full refactoring (means change / improvement of the database scheme based on business logic) on the 4th day of their system functioning.  They found the following objects that are completely unnecessary and can be removed from the circulation.
-	a database view `v_check_country_01_05_2020_covid`
-	a database view `v_dictionary_intersect`
-	a database view `v_dictionary_minus`
-	a database view `v_dictionary_symmetric_minus`
-	a database view `v_dictionary`

Moreover, the second Aliens team came with a request to delete their Heap Temporary table, since all checks and tests were passed.  Therefore, to the general list of deleted objects the following was added
-	a database table `country_north_america`




<h2 id="chapter-v">Chapter V</h2>
<h2 id="exercise-01-lets-get-hierarchy">Exercise 01 - Let’s get hierarchy!</h2>

| Exercise 01: Let’s get hierarchy! |                                                                                                                          |
|-----------------------------------|--------------------------------------------------------------------------------------------------------------------------|
| Turn-in directory                 | ex01                                                                                                                     |
| Files to turn-in                  | `V140__view_with_hierarchy.sql`                                                                                            |
| **Allowed**                           |                                                                                                                          |
| Operators                         | Standard DDL / DML operators to create / alter relations and to insert / update / delete / select data (CRUD operations) |
| SQL keywords                      | WITH RECURSIVE … AS ...                                                                                                  |
| SQL subqueries                    | INSERT INTO … VALUES ( (SELECT id FROM ...), ... )                                                                       |
| Database Naming Convention        | [internal SQL Naming Convention rules](https://docs.google.com/document/d/1IxIOFUeb-8Z8fBOfkXiy4SkN-J1mPIXveJZUCNZFdp8/edit?usp=sharing)<ul><li>pattern for view</li></ul>                                                                                              |
| **Denied**                            |                                                                                                                          |
| Static ID                         | Don’t use a static ID (red font) for SQL statements in exercise.

`INSERT INTO country (name, object_type_id)  VALUES (‘Gibraltar’, 2);`
`SELECT * FROM country WHERE par_id = 2;`
`SELECT setval(‘seq_country’, 2);`

In these cases,  2 is a static (hard-coded) value, use SQL subquery to get a dynamic ID value.                                                                              |

The Aliens realized that they did not take into account the land area of the *territories* in the system and the people who reside there. 
(The *territory* of a state is a part of the globe (including land and its subsoil, water and air), which is under the sovereignty of a certain state and within which its institutions exercise state power.  Territorial supremacy is part of the sovereignty of the state.) 

In order for the database to store information about territories, you must do the following
-	add a new value to the *land* dictionary called *territory* (specifying internal sorting,  `order_num` column value  = current `order_num` value for *land* dictionary + 1)
-	add the following *country* dependent territories with the condition of the country table hierarchy storage model and the condition of specifying the correct reference to the *land* dictionary.  The territories that need to be entered into the system are listed below
    -	Denmark -> Greenland
    -	United Kingdom -> Gibraltar
    -	United Kingdom -> Montserrat
    Don't forget the following 
    -	the ID field for the *country* table is generated automatically and don’t need set by explicitly in INSERT statement
    -	the values of the fields `time_start` and `time_end` are taken by default
    -	within the INSERT statement, use a subquery to retrieve the ID from the *dictionary* for the *territory value* and also to retrieve the parent country ID.

After all the preparation of the data, Aliens decided to write a hierarchical (recursive) SQL query that must be saved in the database view `v_hierarchy_europe` to get all the countries of the continent of Europe with the output of the following attributes
 `{id, name, type_of_land, level}`, where *level* is an indicator of the nesting level of the hierarchy relative to the initial Europe record.

The example of data output from the database view `v_hierarchy_europe`

| id  | name       | `type_of_land` | level |
|-----|------------|--------------|-------|
| 1   | Europe     | continent    | 1     |
| 14  | Belgium    | country      | 2     |
| 16  | Bulgaria   | country      | 2     |
| ... | ...        | ...          | ...   |
| 257 | Greenland  | territory    | 3     |
| 267 | Gibraltar  | territory    | 3     |
| 277 | Montserrat | territory    | 3     |


Please don’t forget comments for each database view and column! ;-)



<h2 id="chapter-vi">Chapter VI</h2>
<h2 id="exercise-02-random-not-fibonacci">Exercise 02 - Random? Not, fibonacci!</h2>

| Exercise 02: Random? Not, fibonacci! |                                                                                                                          |
|--------------------------------------|--------------------------------------------------------------------------------------------------------------------------|
| Turn-in directory                    | ex02                                                                                                                     |
| Files to turn-in                     | `V150__view_getting_fibonacci_countries.sql`                                                                               |
| **Allowed**                              |                                                                                                                          |
| Operators                            | Standard DDL / DML operators to create / alter relations and to insert / update / delete / select data (CRUD operations) |
| Database Naming Convention           | [internal SQL Naming Convention rules](https://docs.google.com/document/d/1IxIOFUeb-8Z8fBOfkXiy4SkN-J1mPIXveJZUCNZFdp8/edit?usp=sharing)<ul><li>pattern for view</li></ul>                                                                                              |
| SQL keywords                         |<ul><li>WITH RECURSIVE … AS …</li><li>UNION ALL</li><li>LEFT JOIN</li></ul>                                                                                                     |

Aliens decided to create a separate group for monitoring people, for this they had to select several arbitrary countries.  Aliens decided to abandon the random approach and to use a method of selecting a country by ID that matches the Fibonacci number.

Fibonacci sequence is a sequence of numbers {1, 1, 2, 3, 5, 8, 13,...} which are satisfied a mathematical rule: `Fn=F(n-1)+F(n-2)`

Help Aliens create the database view `v_fibonacci_country` with rules below.
-	use the Fibonacci number generation limit to 2000.
-	note that the output contains all Fibonacci numbers, if there is no corresponding country then the field is filled with null
-	sorting the result should be by the Fibonacci number itself
-	and one moment … please don’t forget comments for each database view and column! 

The example of the query result for the created `v_fibonacci_country` looks like this

| fibonacci | `country_id` | `country_name`  |
|-----------|------------|---------------|
| 1         | 1          | Europe        |
| 1         | 1          | Europe        |
| 2         | 2          | Africa        |
| 3         | 3          | Antarctica    |
| 5         | 5          | North America |
| 8         | 8          | Albania       |
| 13        | 13         | Belarus       |
| ...       | ...        | ...           |
| 144       | 144        | Liberia       |
| 233       | null       | null          |
| 377       | null       | null          |
| 610       | null       | null          |
| 987       | null       | null          |
| 1597      | null       | null          |





<h2 id="chapter-vii">Chapter VII</h2>
<h2 id="exercise-03-hmmm-lets-innovate">Exercise 03 - Hmmm, let’s innovate!</h2>

| Exercise 03: Hmmm, let’s innovate! |                                                                                                                          |
|------------------------------------|--------------------------------------------------------------------------------------------------------------------------|
| Turn-in directory                  | ex03                                                                                                                     |
| Files to turn-in                   | `V160__view_lateral_join.sql`                                                                                              |
| **Allowed**                            |                                                                                                                          |
| Operators                          | Standard DDL / DML operators to create / alter relations and to insert / update / delete / select data (CRUD operations) |
| Database Naming Convention         | [internal SQL Naming Convention](https://docs.google.com/document/d/1IxIOFUeb-8Z8fBOfkXiy4SkN-J1mPIXveJZUCNZFdp8/edit?usp=sharing) rules<ul><li>pattern for view</li></ul>                                                                                              |
| SQL keywords                       | <ul><li>WITH RECURSIVE … AS …</li><li>LEFT/RIGHT JOIN LATERAL … ON ...</li></ul>|

Aliens realized that the returned result from the database view `v_fibonacci_country` is not what they expect, since not specific countries, but for example entire continents, may be included in the selection.  Moreover, the SQL Performance Engineer said that it would be good to use a LATERAL JOIN instead of a regular LEFT / RIGHT JOIN.

Help Aliens create a new database view `v_fibonacci_country_lateral` with the following rules
-	Using the LATERAL JOIN, rewrite the original database SQL query `v_fibonacci_country` to select only countries based on the *land* dictionary.
-	I didn’t say before one important thing … please don’t forget comments for each database view and column! 

The example of the query result to the generated `v_fibonacci_country_lateral` looks like this.

| fibonacci | `country_id` | `country_name`   |
|-----------|------------|----------------|
| 1         | null       | null           |
| 1         | null       | null           |
| 2         | null       | null           |
| 3         | null       | null           |
| 5         | null       | null           |
| 8         | 8          | Albania        |
| 13        | 13         | Belarus        |
| 21        | 21         | Estonia        |
| 34        | 34         | Luxembourg     |
| 55        | 55         | United Kingdom |
| 89        | 89         | Jordan         |
| 144       | 144        | Liberia        |
| 233       | null       | null           |
| 377       | null       | null           |
| 610       | null       | null           |
| 987       | null       | null           |
| 1597      | null       | null           |





<h2 id="chapter-viii">Chapter VIII</h2>
<h2 id="exercise-04-alert-sabotage-is-detected">Exercise 04 - Alert! Sabotage is detected!</h2>

| Exercise 04: Alert! Sabotage is detected! |                                                                                                                          |
|-------------------------------------------|--------------------------------------------------------------------------------------------------------------------------|
| Turn-in directory                         | ex04                                                                                                                     |
| Files to turn-in                          | `V170__change_content.sql`                                                                                                 |
| **Allowed**                                   |                                                                                                                          |
| Operators                                 | Standard DDL / DML operators to create / alter relations and to insert / update / delete / select data (CRUD operations) |
| SQL statements                            | <ul><li>DELETE FROM …</li><li>UPDATE … SET ... FROM … WHERE</li></ul>                                                                                 |
| **Denied**                                    |                                                                                                                          |
| Static ID                                 | Don’t use a static ID (red font) for SQL statements in exercise.

`INSERT INTO country (name, object_type_id)  VALUES (‘Gibraltar’, 2);`
`SELECT * FROM country WHERE par_id = 2;`
`SELECT setval(‘seq_country’, 2);`

In these cases,  2 is a static (hard-coded) value,use SQL subquery to get a dynamic ID value.                                                                              |

Paul McAlien from Alien Support Team told you to make the following 2 changes regarding the `country_indicator` table.
-	Delete “Unemployment rate” data for all countries for January and November 2019 using one DELETE statement
-	Update data for all indicator values for the countries of the continent Asia using the formula: `current_value:= current_value + 100`, using one UPDATE statement with JOIN-like  query inside. (JOIN-like query is UPDATE t1 … FROM t2 WHERE … )



<h2 id="chapter-ix">Chapter IX</h2>
<h2 id="exercise-05-return-our-data-back">Exercise 05 - Return our data back</h2>

| Exercise 05: Return our data back |                                                                                                                          |
|-----------------------------------|--------------------------------------------------------------------------------------------------------------------------|
| Turn-in directory                 | ex05                                                                                                                     |
| Files to turn-in                  | `V180__regenerate_content.sql`                                                                                             |
| **Allowed**                           |                                                                                                                          |
| Operators                         | Standard DDL / DML operators to create / alter relations and to insert / update / delete / select data (CRUD operations) |
| SQL statements                    | TRUNCATE ...                                                                                                             |
| Functions                         | random()                                                                                                                 |
| **Denied**                            |                                                                                                                          |
| SQL statements                    | DELETE ...                                                                                                               |
| Functions                         | generate_series(integer, integer),
generate_series(date, date)                                                        | 
| Static ID                         | Don’t use a static ID (red font) for SQL statements in

`INSERT INTO country (name, object_type_id)  VALUES (‘Gibraltar’, 2);`
`SELECT * FROM country WHERE par_id = 2;`
`SELECT setval(‘seq_country’, 2);`

In these cases,  2 is a static (hard-coded) value, use SQL subquery to get a dynamic ID value.                                                                              |

The Aliens realized that Paul McAlien was not who he said he was, but the data in the database was corrupted.  For simplicity, the Aliens team decided to take the following steps:
-	full cleanup of `country_indicator` table
-	use another approach (please don't use `generate_series` functions) then SQL script from Day 00 / Day 01 to regenerate a data for each country and territory for each of 4 human indicators based on this limitations:
1.	“*Population of country*” for each 1st day of month 2019 year \
	(take a random value for this parameter from 0 to 1 000 000)
2.	“*Unemployment rate*” for each 1st day of month 2019 year \
(take a random value for this parameter from 0 to 100)
3.	“*Infected humans COVID-19*”  for each day beginning from 1st of May 2020 to 31st of August 2020 \
(take a random value for this parameter from 0 to 50)
4.	“*Area of the land*” for the 1st of May 2020 \
	(take a random value for this parameter from 10 000 to 10 000 000)
