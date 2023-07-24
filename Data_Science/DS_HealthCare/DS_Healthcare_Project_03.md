# Project 03 – Sentiment Analysis in Drug Reviews

После выполнения проекта просим заполнить [форму обраной связи](https://docs.google.com/forms/d/e/1FAIpQLScVdZC6J0yZVhdDdAwgkeF1VZ6IoXad6WTDeEgTsFoZIHhAYg/viewform)

Rating Score Prediction Using Machine Learning

Summary: This project is an introduction to machine learning application for automating prediction of rating score


## Contents

1. [Chapter I](#chapter-i) \
    1.1. [Preamble](#preamble)
2. [Chapter II ](#chapter-ii) \
    2.1. [Introduction](#introduction)
3. [Chapter III](#chapter-iii) \
    3.1. [Goals](#goals)  
4. [Chapter IV](#chapter-iv) \
    4.1. [General instructions](#general-instructions)  
5. [Chapter V. Mandatory Part](#chapter-v-mandatory-part) \
    5.1. [a. Dataset](#a-dataset)  
    5.2. [b. Task](#b-task)  
    5.3. [c. Implementation](#c-implementation)  
    5.4. [d. Submission](#d-submission)
6. [Chapter VI](#chapter-vi) \
    6.1. [Bonus Part](#bonus-part)  
7. [Chapter VII](#chapter-vii) \
    7.1. [Submission and Peer-Correction](#submission-and-peer-correction)  

    

<h2 id="chapter-i" >Chapter I</h2>
<h3 id="preamble">Preamble</h3>

Sentiment analysis of a text is a complex process involving the extraction of useful subjective information from a text. A huge amount of user-generated content on the Internet appears every day. Every day millions of users express their opinions about products and services on blogs, social networks and other information resources. Providing reliable opinion extraction from unstructured text is essential for commercial organizations. With the help of the data provided, companies will be able to find out the opinion of buyers that is important to them, find imperfections invisible to their eyes and increase their sales. Sentiment analysis is used on many text documents that contain emotions and assessments of certain objects, for example, people, events, topics (for example, reviews of films, books, products). Sentiment analysis involves identifying sentiment in a document, and subsequently determining its positive / negative polarity.

In medical practice, the analysis of the sentiment of reviews can help to identify the success of a particular drug. And also to assess the psychological state of the person who wrote the review.

The accuracy and quality of the text sentiment analysis system is assessed by how well it agrees with the person's opinion regarding the emotional assessment of the text under study. According to various studies, experts usually agree on the sentiment of a particular text in 79% of cases. Consequently, a program that determines the tonality of a text with 70% accuracy does it almost as well as a human.





<h2 id="chapter-ii" >Chapter II </h2>
<h3 id="introduction">Introduction</h3>

Sentiment analysis is a class of methods of content analysis in computational linguistics, the main task of which is to classify text according to its mood. In simple cases, the task of sentiment analysis is reduced to a binary classification of texts into positive and negative. In some cases, add another class of neutral texts. More advanced approaches try to identify emotional states associated with a text, such as fear, anger, sadness, or happiness. In a number of approaches, texts are assigned values ​​of a predetermined scale: for example, from -2 for negative to 2 for positive; thus, the analysis is reduced to a regression problem. Aspect-based sentiment analysis is a subset of sentiment analysis, whose task is to determine the attitude towards a specific aspect of the main subject of discussion. All approaches to sentiment analysis can be divided into three groups.

The first is rule-based approaches. Most often, they use manually defined classification rules and emotionally marked vocabularies. These rules usually calculate the text class based on emotional keywords and their combination with other keywords. While they are excellently effective in subject matter, rule-based methods are poorly generalizable. In addition, they are extremely time consuming to create, especially when there is no access to a suitable sentiment dictionary.

The second group is machine learning approaches. They use automatic feature extraction from text and apply machine learning algorithms. The classical algorithms for classification of polarity are Naive Bayes Classifier, Decision Tree, Logistic Regression. In recent years, the attention of researchers has been attracted by deep learning methods, which are significantly superior to traditional methods in sentiment analysis.

In this work, you have a dataset of drug reviews and user ratings. The main task is to write a program to predict the review score. To do this, you need to analyze the available data, build a predictive model and evaluate the effectiveness of its work.

This problem is actually a classification and regression problem, so you can use a wide variety of classifiers to solve it, from logistic regression to neural networks. We invite you to try many of these forecasting models, compare their performance, and conclude which one is best for your task. We hope you enjoy it.


<h2 id="chapter-iii" >Chapter III </h2>
<h3 id="goals" >Goals</h3>

The goal of this project is to give you an example of using a machine learning approach to solve sentiment analysis problem. You will try different algorithms to predict a rating score by a review like an expert. You can actually use your program for the creation of a sentiment prediction support system. 


<h2 id="chapter-iv" >Chapter IV</h2>
<h3 id="general-instructions" >General Instructions</h3>

- This project will only be evaluated by humans. You are free to organize and name your files as you desire.
- Use Python as a programming language and any libraries and packages supported.
- Use Google Colab, Jupyter or PyCharm as a development environment.
- Write your program so that other people can understand it.
- Store the dataset in your Google Drive or locally to access it from your program.

<h2 id="chapter-v" >Chapter V. Mandatory Part</h2>
<h3 id="a-dataset" >a. Dataset</h3>

You will work with open dataset “WebMD Drug Reviews Dataset”:

<https://www.kaggle.com/rohanharode07/webmd-drug-reviews-dataset>

Complete dataset consists of one CSV files: “webmd.csv”. Data was acquired by scraping WebMD site. There are around 0.36 million rows of unique reviews. This CSV file has 12 columns: Drug (categorical): name of drug, DrugId (numerical): drug id, Condition (categorical): name of condition, Review (text): patient review, Side (text): side effects associated with drug (if any), EaseOfUse (numerical): 5 star rating, Effectiveness (numerical): 5 star rating, Satisfaction (numerical): 5 star rating, Date (date): date of review entry, UsefulCount (numerical): number of users who found review useful, Age (numerical): age group range of user, Sex (categorical): gender of user. You are required to train and test your model on data of reviews and satisfaction scores.

<h3 id="b-task" >b. Task</h3>

1. Data Analysis
   1. Read source data. Find out what they consist of, what data are needed to solve the problem.
   1. Clear data from gaps, extra spaces, corrupted data.
   1. Find out how many reviews there are with each rating.
   1. Find out the maximum, minimum, average length of reviews.
1. Data Preparation
   1. Tokenize reviews. Remove unnecessary characters.
   1. Lemmatize reviews.
   1. Determine how much the amount of data has changed.
   1. Visualize word clouds for each of the rating scores.
   1. Divide the sample into training and validation sets.
   1. Extract numeric features from obtained datasets using the CountVectorizer and TfidfVectorizer  methods.
1. Classification Models
   1. Build a predictive model for a naive Bayesian classifier.
   1. Use the Logistic Regression model to predict the score.
   1. Build a decision tree to predict the rating of reviews.
   1. Using the random forest method, build a predictive model.
1. Neural Networks
   1. Build and train a simple LSTM neural network.
   1. Visualize dependence of accuracy and loss function on the epoch number.
1. Model Evaluation
   1. For each of the above methods, evaluate the classification accuracy and MAE.
   1. Calculate Training time and Prediction time.
   1. Calculate for each review score Precision, Recall and F1-score
   1. Build and visualize the confusion matrix.

<h3 id="c-implementation" >c. Implementation</h3>

You can work in your private Git repository so a reviewer can access it.

You can use any library or any framework you want.

You should keep a research diary with all information about the used approaches and their metrics.

<h3 id="d-submission" >d. Submission</h3>
Share your program on your private Git repository with your reviewer to submit it. This repository should contain your working program and text explanation.


<h2 id="chapter-iii" >Chapter VI </h2>
<h3 id="bonus-part" >Bonus Part</h3>

1. Take into account the presence of smilies in the text of the review at the stage of data preparation.
2. Use word2vec methods for extracting numeric features besides CountVectorizer and TfidfVectorizer. Determine impact on the predictive accuracy of the estimate.
3. For classification models, select hyperparameters to improve prediction efficiency.
4. Make changes to the architecture of the LSTM network, identify their impact on improving the accuracy of predictions.
5. Change the original sample so that there are 2 classes of ratings (positive and negative). Build predictive models and a neural network for this data. Assess the results obtained.
6. Make final conclusions and recommendations for solving the problem of diagnosing diseases by symptoms: choose the best classifier, justify your choice, describe the main steps of the algorithm for solving the problem.




<h2 id="chapter-vii" >Chapter VII</h2>
<h3 id="submission-and-peer-correction" >Submission and Peer-Correction</h3>

Submit your private Git repository as usual. Only the content of your private git repository will be graded.

Here are the points that your peer-corrector will have to check:

- if all the approaches are tried (classifiers, neural networks),
- if the almost perfect accuracy achieved on the test dataset,
- if the research diary exists.
