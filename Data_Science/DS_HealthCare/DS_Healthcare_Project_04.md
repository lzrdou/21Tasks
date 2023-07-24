# Project 04 – Named Entity Recognition according to the classification of word names and diseases in the text

Просим заполнить форму отзыва о задании [здесь](https://forms.gle/kGrN9Mxr75yEKQLM7)

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

Clinical studies often require detailed patients’ information documented in clinical narratives. Named Entity Recognition (NER) is a fundamental Natural Language Processing (NLP) task to extract entities of interest (e.g., disease names, medication names and lab tests) from clinical narratives, thus to support clinical and translational research. Clinical notes have been analyzed in greater detail to harness important information for clinical research and other healthcare operations, as they depict rich, detailed medical information.
In this challenge, Students are invited to extract all disease names from a given paragraphs/documents. 


<h2 id="chapter-ii" >Chapter II </h2>
<h3 id="introduction">Introduction</h3>

Before going ahead with deep learning and python based implementation, It is important to clearly understand the kind of problem NER is. Beginners may confound it with a sentence parsing problem or a classical classification problem. Essentially, unlike other sentence or document classification technique (as in this and this post), NER is a word classification problem where each word of the sentence has to be classified among the labelled tags.
An obvious question that arises is regarding the kind of classifier which can be used in such problem. A classification model which can model the words of sentence in sequence of states/nodes along with tagging each of these words with class tag. This will allow contextual learning of entities and classification of each word at the same time. In non deep learning models, Conditional Random field (CRF) had been an obvious and popular choice for modelling NER problems.
In addition to CRF, in this project you will have the opportunity to try to solve the problem using the neural network the LSTM architecture. The state of an LSTM network is represented through a state space vector. This method allows you to track the dependence of new observations on the past (even very distant ones).

We suggest that you try different word classification models, compare their effectiveness and conclude which one is best for your task. Hope you have some fun. 


<h2 id="chapter-iii" >Chapter III </h2>
<h3 id="goals" >Goals</h3>

The goal of this project is an introduction to Named-entity recognition (NER) (also known as entity extraction). This is essentially an introduction to natural language processing. In this project, you will try to classify individual sections of sentences (words) by tags.


<h2 id="chapter-iv" >Chapter IV</h2>
<h3 id="general-instructions" >General Instructions</h3>

- This project will only be evaluated by humans. You are free to organize and name your files as you desire.
- Use Python as a programming language and any libraries and packages supported.
- Use Google Colab, Jupyter or PyCharm as a development environment.
- Write your program so that other people can understand it.
- Store the dataset in your Google Drive or locally to access it from your program.

<h2 id="chapter-v-mandatory-part" >Chapter V. Mandatory Part</h2>
<h3 id="a-dataset" >a. Dataset</h3>

You will work with open dataset “hackathon_disease_extraction”.

We need files train.csv and test.csv
test (Test-set) : 20000 documents
train.csv (Train-set) : 30000 documents with labelled entities (diseases).


The train file has the following structure:

| Variable | Definition |
| ------ | ------ |
| id | Unique ID for a token/word |
| Doc_ID | Unique ID for a Document/Paragraph |
| Sent_ID | Unique ID for a Sentence |
| Word | Exact word/token |
| tag (Target) | Named Entity Tag |


The target ‘tag’ follows the Inside-outside-beginning (IOB) tagging format. The IOB format (short for inside, outside, beginning) is a common tagging format for tagging tokens in named entity recognition. The target ‘tag’ has three kinds of tags.
B-indications : Beginning tag indicates that the token is the beginning of a disease entity (disease name in this case).
I-indications : Inside tag indicates that the token is inside an entity.
O : Outside tag indicates that a token is outside a disease entity.

Therefore, any word which does not represent the disease name has to be classified as “O” tag. Similarly, the first word of disease name has to be classified as “B-Indication” and following words of disease name as “I-Indication”.


<h3 id="b-task" >b. Task</h3>

1. Data Analysis
   1. You only need to download train.csv
   1. Check the data for integrity. If there is an empty column in the data, then it must be removed.
   1. Visualize the contents of the first 10 columns
   1. Count the number of unique words and the number of documents that are presented in the sample
1. Data Preparation
_For LSTM:_
   1. Convert greek characters to ASCII characters
   1. Create your own dictionary of words and signs that are contained in the sample and output the number of words in the resulting dictionary
   1. Create your own dictionary from the tags contained in the selection and output the number of tags in the resulting dictionary
   1. Create dictionaries for word indices and for tag indices
   1. Split words into a list of sentences using Sent_ID. Calculate how many sentences are there.
   1. Find the maximal length of a sentence.
   1. Enumerate words in sentences assigning them indices.
   1. It is necessary to limit the number of words in the sentence to make all sentences have the same length. As an example, the suggestions are to make all sentences 180 words long. Thus, sentences longer than 180 words are truncated, and sentences less than 180 words in length are padded with some special word tagged as ‘O’ up to to 180 words length.
   1. Divide the sample into test and training parts
1. Data Preparation
_For CRF:_
   1. Fill null values in the dataset with some non-empty string, if necessary
   1. Define the part of speech for each word. For example, you can do this with the POS-tagger from Nltk. Add a part of speech column to your dataframe
   1. Split words into a list of sentences using Sent_ID. Calculate how many sentences are there.
   1. Get a set of features based on the resulting words:
        - convert to lower case
        - take the last 3 letters of the word
        - take the last 2 letters of the word
        - take a boolean value (flag) which would display whether the word contains only uppercase letters or not
        - take a boolean value (flag) that would display whether the word starts with a capital letter or not
        - take a boolean value (flag) that would display whether the word contains only numbers or not
        - take the word tag
        - take the first 2 letters of the word tag
        - take these features again for the previous word and for the next word in the sentence if there is one
   5. Divide the sample into test and training parts
1. LSTM
   1. Build OneHotEncoder to encode class names for training and test samples.
   1. Build a Bidirectional LSTM for training. Add an Embedding layer and a Dropout layer in front of it. Add Time Distributed Dense layer after it. The input dimension should be equal the number of words in the sentence, and the output dimension should be equal the number of classes (3).
   1. Train the neural network on the training sample.
   1. It is necessary to calculate the value of the following metrics: recall, accuracy, F1-score, precision.
1. CRF
   1. Build a CRF model with arbitrary parameters of your choice
   1. Train the CRF on the training sample
   1. It is necessary to calculate the value of the following metrics: recall, accuracy, F1-score, precision.


<h3 id="c-implementation" >c. Implementation</h3>

You can work in your private Git repository so a reviewer can access it.

You can use any library or any framework you want.

You should keep a research diary with all information about the used approaches and their metrics.

<h3 id="d-submission" >d. Submission</h3>
Share your program on your private Git repository with your reviewer to submit it. This repository should contain your working program and text explanation.


<h2 id="chapter-vi" >Chapter VI </h2>
<h3 id="bonus-part" >Bonus Part</h3>

1. Estimate the training time and work time of each model, formulate conclusions about their computational efficiency
2. Change the architecture of the recurrent neural network (change the number of units, the number of batches, the number of epochs, rate for the Dropout layer), and see how the learning rate and classification metrics depend on these parameters
3. Change the parameters of the CRF model (number of training iterations, coefficients, used algorithm for training and regularization) , and see how the learning rate and classification metrics depend on these parameters
4. Formulate the final conclusions and recommendations for solving the multiclass classification of word names and diseases in the text problem: choose the best classifier, justify your choice, describe the main stages of the algorithm for solving the problem.



<h2 id="chapter-vii" >Chapter VII</h2>
<h3 id="submission-and-peer-correction" >Submission and Peer-Correction</h3>

Submit your private Git repository as usual. Only the content of your private git repository will be graded.

Here are the points that your peer-corrector will have to check:

- if all the approaches are tried (classifiers, neural networks),
- if the almost perfect accuracy achieved on the test dataset,
- if the research diary exists.
