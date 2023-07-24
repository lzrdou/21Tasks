# Rush 01 - Database Boot camp

JSON or Relational Model?

*Resume: Today you are implementing a model inside another model! Sounds cool!*

## Contents
1. [Chapter I](#chapter-i) \
    1.1. [Preamble](#preamble)
2. [Chapter II](#chapter-ii) \
    2.1. [General Rules](#general-rules)
3. [Chapter III](#chapter-iii) \
    3.1. [Rules of the day](#rules-of-the-day)
4. [Chapter IV](#chapter-iv) \
    4.1. [Exercise 00 - Let’s try a new model](#exercise-00-lets-try-a-new-model)

<h2 id="chapter-i" >Chapter I</h2>
<h2 id="preamble" >Preamble</h2>

![R01_01](images/DB_Bootcamp_Team01_0.png)
source: https://db-engines.com/en/ranking_trend 

Physical relational databases - is it really just normalization, consistency, I,II,III, ... Normal Forms, functional dependency, multivalued dependency?
Hold on, but so far we have in the TOP-4 of the most common databases in use in the world - relational!

What is the secret here?  How are relational databases keeping up with such increased demand? Maybe the reason is in legacy code or in engineers who do not know or do not want to improve their skills and switch to such trendy / fast noSQL databases?

In my opinion, it is all about the constant development of relational databases that already contain data types (far from simple and that immediately break I Normal Form) or table types that are already adapting to fast data processing.  These data types and specialised tables include the following data types:
- XML
- json / jsonb
- edifact
- Nested Tables
- Index Organized Tables
- bloom filter indexes
- etc.

![R01_02](images/DB_Bootcamp_Team01_1.png)
source: https://brothers-smaller.ru/samyie-byistryie-zhivotnyie/

In other words, relational databases are now not those cute, mathematically correct "animals" that were described several decades ago.  These are now mutated into stronger and more enduring predators (of course, in a good sense of the word), solving the problems of search, data volume, scaling with already existing and proven mechanisms.

<h2 id="chapter-ii">Chapter II</h2>
<h2 id="general-rules" >General Rules</h2>

- Use this page as the only reference. Do not listen to any rumors and speculations on how to prepare your solution.
- Please make sure you are using the latest version of PostgreSQL.
- Please make sure you have installed and configured the latest version of Flyway by Redgate.
- Please use our [internal SQL Naming Convention rules](https://docs.google.com/document/d/1IxIOFUeb-8Z8fBOfkXiy4SkN-J1mPIXveJZUCNZFdp8/edit?usp=sharing)
- Please use our [Terms and Definitions](https://docs.google.com/document/d/1_ZTDpHcfYMASZ5BtnldurQLF0fJygGF1yuTwik0DOqk/edit?usp=sharing) document
- That is completely OK if you are using IDE to write a source code (aka SQL script) and make a syntax check before migration at the final database solution by Flyway.
- Comments are also good in the SQL scripts. Anyway be careful with signs /\*...\*/ directly in SQL. These special symbols are used for Database Hints to improve SQL performance and these are not just comment marks :-).
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

- Please make sure you have a separated database “rush” on your PostgreSQL cluster. 
- Please make sure you have a database schema “rush01” in your “rush” database.
- Please make sure you are working through database user “rush” and password “rush” with super admin permissions for your PostgreSQL cluster. 
- Please execute the next statement for the “rush01” schema before the task.
  `CREATE EXTENSION "uuid-ossp";`
    - check availability of extension by SQL 
      `select extname, extrelocatable from pg_catalog.pg_extension where extname = 'uuid-ossp'`

- Each exercise of the day needs a Flyway tool for right versioning of the “rush” database through user “rush”.
- Please make changes in your “flyway.conf” file (located in “~/flyway-6.x.x/conf” directory) directly to configure a right and stable connection to your PostgreSQL database.

    - flyway.url = jdbc:postgresql://hostname:5432/rush 
        - where hostname is DNS / IP address of PostgreSQL server 
        - where port is a port of PostgreSQL server , by default is 5432
        (jdbc:postgresql://localhost:5432/rush OR  jdbc:postgresql://127.0.0.1:5432/rush)
    - flyway.user = rush
    - flyway.password = rush
    - flyway.schemas = rush01
    - flyway.defaultSchema = rush01
- Please use the command line for Flyway to migrate changes into the database and get information about the current version from the database.
- Please don’t append additional parameters for “flyway” in a command line, all needed parameters should be changed in “flyway.conf” file
- All tasks contain a list of Allowed and Denied sections with listed database options, database types, SQL constructions etc. Please have a look at the section before you start.

<h2 id="chapter-iv">Chapter IV</h2>
<h2 id="exercise-00-lets-try-a-new-model">Exercise 00 - Let’s try a new model</h2>

|Exercise 00: Let’s try a new model||
---|---
Turn-in directory|ex00|
Files to turn-in| `V004__rush01_tables.sql` 
 `V005__rush01_view.sql` 
 screenshot `top_20_words.png`|
**Allowed**||
Database Tables| Original Heap Tables |
Database objects| User defined procedures / functions|
Operators | Standard DDL / DML operators to create / alter relations and to insert / update / delete / select data (CRUD operations) |
**Denied** ||
Database Types | Serial |
Database Objects | `triggers`
`trigger functions`
`database rules`
`PRIMARY KEY`
`generate_series(...)` |
Database Keywords | `LIKE` | 

Before beginning **please read** the section “Rules of the day” about steps to create a  new database and schema inside for Rush! ... and yes please don’t forget about comments for your tables / views / columns. Let’s make a world more clearly and understandable.

![R01_03](images/DB_Bootcamp_Team01_2.png)


Do you know Elon Musk? I guess, yes you are knowing this man and probably know his love to make tweets on twitter.com with other themes. So let’s try to use this [archive](https://drive.google.com/file/d/1poZcrhQVyBmeXKt7Ob-17ayWJ2RjA0no/view?usp=sharing) of tweets by Elon Musk for our Rush exercise. 

Each line from the file contains exactly one tweet and the corresponding structure of the tweet is presented below.


| Name | Type | Business Description |
| ------ | ------ | ------ |
| "Text" | TEXT | Content of tweet |
| "UserName" | VARCHAR | User Name (In our case, **elonmusk**) |
| "LinkToTweet" | VARCHAR | URL address of tweet |
| "TweetEmbedCode" | TEXT | Original HTML form of tweet |
| "CreatedAt" | VARCHAR | DateTime is presented like a string.
Example: December 03, 2017 at 10:24PM |

Let’s create a database table with format key-value storage. 
In our case key is generated an unique key based on UUID type (based on `uuid_generate_v4()`)
The value is **JSONB** type which will contain one tweet with all additional data features (other words one row in this table is equal to each line from the archive of tweets).

Please create a database table with name `archive_tweets_elonmusk` the next structure

| Name | Type |
| ------ | ------ |
| tweet_id | UUID
NOT NULL
DEFAULT uuid_generate_v4() |
| tweet_document | JSONB |

Use flyway file `V004__rush01_tables.sql` for this action.

Please migrate data from attached tweet archive into a table `archive_tweets_elonmusk`  into a database. You can use other  ETL methods which you are familiar with.

So, prepare a database view `v_get_tweet_about_tesla` which should return all tweets which contain the word “tesla” in any case (upper and lower). Please use flyway script `V005__rush01_view.sql` for that action.

The result should be looked like below

`select *`

`from v_get_tweet_about_tesla`


| tweet | 
| ------ | 
| `RT @KM2__S: Man, how beautiful. @Tesla Model 3  Here's Why the Tesla Model 3 Is the Coolest Car of 2017 @elonmusk #tesla #ElonMusk https://t.co/5ZVZO0qQql` |
| `@Jason @Tesla Sure` |
| `@neilsiegel @Tesla Coming very soon` |
| `RT @ElectrekCo: Watch a Tesla Model X all-electric SUV tow a semi truck stuck in the snow https://t.co/HR9WbuoBTh https://t.co/5kqnXhbHpz` |
| ... |

Based on returned data from `v_get_tweet_about_tesla` view please make a clearing stop-words from tweet. You can find out stop-words [here](https://drive.google.com/file/d/1X2ulCnAKmTfkt59qe3o7F4Z9o14YaFko/view?usp=sharing). If you want you can create a new table for your purposes or just clearing somewhere in memory of your computer or in the cloud. 
Then please create a TOP-20 more usable words  by Elon Musk with context about “Tesla” car or company. Please provide a screenshot `top_20_words.png` with this analysis. 

The result of this task can see like below

![R01_04](images/DB_Bootcamp_Team01_3.png)


So, the next task is to prepare a database view `v_get_distribution_tweet_by_month_year` which should return distribution tweets per pair year-month. Please make all parsing logic on database layer by creating a new Pl/PgSQL function `fnc_get_month_year(pcreatedate)`, this function should return a date in format “dd/mm/yyyy”, where is “dd” is always “01”. For example,  December 03, 2017 at 10:24PM will be presented like “01/12/2017”. 
Use flyway file `V006__rush01_view.sql` for this task. The result of this exercise please create a graphical presentation of tweet distribution for all years and months. The example looks like below.

![R01_05](images/DB_Bootcamp_Team01_4.png)
