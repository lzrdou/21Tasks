# Project 05 – Multi-label Classification in Radiographic Imaging

Просим заполнить форму отзыва о задании [здесь](https://forms.gle/KZWYBFi77b4Vo9nW7)

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
Analysis and classification of the results of medical examinations is an important and urgent task in the modern world. Every year a person needs to remove ecg, do fluorography of the lungs, take a general blood test. Thousands of medical examinations are carried out every day, which must be processed by someone and there is not always enough human resources for this. At such moments, such as in a pandemic, the question of automating the process of processing or preprocessing medical data obtained as a result of examinations arises. One of the optimization methods is deep learning methods such as neural networks.

In addition, medical data can be processed using various methods of analysis and statistics, for example, it is possible to determine the disease by CT using histogram analysis or analysis of textural features.


<h2 id="chapter-ii">Chapter II</h2>
<h3 id="introduction">Introduction</h3>

In machine learning, multiclass or polynomial classification is the problem of classifying instances in one of three or more classes (classifying instances in one of the two classes is called binary classification). Multi-class classification of images is actively used in various fields, including medicine. You can perform multiclass classification using classical classifiers, neural networks, heuristic methods, feature analysis.

One of the types of features that can be obtained from an image is histogram features.Histogram features are statistical and describe the distribution of image brightness. This type of feature includes, for example, the mean, standard deviation, distortion, energy, entropy.

In addition, one of the peculiarities of solving classification problems is data imbalance. Data imbalance is when one class clearly prevails over another / others. In medicine, class imbalance occurs quite often, for example, with annual fluorography, many more diagnoses "without pathologies" are made than "with pathologies". To eliminate the imbalance of classes, various methods are used, one of which we will get acquainted with in this project.


<h2 id="chapter-iii">Chapter III</h2>
<h3 id="goals" >Goals</h3>

The goal of this project is to give you an example of using a machine learning approach to solve problems of multiclass classification of medical images on the example of CT of the lungs. You will try to classify images as using neural networks directly, and with the help of features that can be read from images, using classic classifiers

<h2 id="chapter-iv" >Chapter IV</h2>
<h3 id="general-instructions" >General Instructions</h3>

-   This project will only be evaluated by humans. You are free to organize and name your files as you desire.
-   Use Python as a programming language and any libraries and packages supported.
-    Use Google Colab, Jupyter or PyCharm as a development environment.
-    Write your program so that other people can understand it.
-   Store the dataset in your Google Drive or locally to access it from your program.

<h2 id="chapter-v" >Chapter V. Mandatory Part</h2>
<h3 id="a-dataset" >a. Dataset</h3>

You will work with open dataset “NIH Chest X-rays”:

<https://www.kaggle.com/nih-chest-xrays/data>

The dataset consists of 112,120 X-ray images with disease labels from 30,805 unique patients. In addition, the images are accompanied by the Dataentry2017 markup file. Dataentry2017.csv: 
Class labels and patient data for the entire dataset 
Image Index: File name
Finding Labels: Disease type (Class label)
Follow-up #
Patient ID
Patient Age
Patient gender
View Position: X-ray orientation
OriginalImageWidth
OriginalImageHeight
OriginalImagePixelSpacing_x
OriginalImagePixelSpacing_y

<h3 id="b-task" >b. Task</h3>

1. Data Analysis
    1. It is necessary to download images from the images_001 directory.
    1. Check the data for integrity. If there is an empty column in the data, then it must be removed.
    1. Visualize first 5 images.
    1. Visualize data on the frequency of occurrence of various classes of pathologies in a given dataset.
    1. Create dataset of 2500 images (file size with markup must contain data for 2500 images).
1. Data Preparation
    1.  Choose the 5 most common classes.
    1.  Visualize the numerical distribution of the resulting 5 classes.
    1.  Divide the dataset into training and control samples.
1. Convolutional neural network
    1.  Build OneHotEncoder to encode class names for training and test samples.
    1.  Create your own convolutional neural network architecture for image classification.
    1.  The network should have the following architecture: convolutional layers should alternate with pooling layers, while the size of the images from layer to layer decreases, and the numerical value of the filters increases. This is followed by Dense layers alternating with exclusion layers.
    1.  Train the neural network on the training sample.
    1.  It is necessary to calculate the value of the following metrics: recall, accuracy, F1-score, precision.
    1.  Visualize error matrix.
    1.  Construct a graph of the change in the loss function depending on the number of the training epoch.
1. MobilenetV2 
    1.  Build OneHotEncoder to encode class names for training and test samples.
    1.  Load basic neural network architecture.
    1.  The network must have the following architecture: load the architecture of the MobilenetV2 model from keras.applications with undefined weights and an alpha value of no more than 0.5.
    1.  Train the neural network on the training sample.
    1.  It is necessary to calculate the value of the following metrics: recall, accuracy, F1-score, precision.
    1.  Visualize error matrix.
    1.  Construct a graph of the change in the loss function depending on the number of the training epoch.
1. Histogram features for Classification Models
    1.  Select an area of ​​interest using an algorithm based on the Otsu method.
    1.  Calculate the following characteristics of the histogram for the selected area of ​​interest (lung) on ​​the original image: average gray level, gray level variance, median gray level, gray level fluctuation (difference between max and min), asymmetry coefficient, kurtosis coefficient.
    1.  Write the received characters to a file with the csv extension.
1. Classification Models
    1.  Build and evaluate a predictive model for a Naive Bayes Gaussian classifier.
    1.  Build and evaluate a predictive model for a k Nearest Neighbours classifier.
    1.  Build and evaluate a predictive model for a Random Forest classifier.
    1.  Build and evaluate a predictive model for a Gradient Boosting classifier.

<h3 id="c-implementation" >c. Implementation</h3>

You can work in your private Git repository so a reviewer can access it.

You can use any library or any framework you want.

You should keep a research diary with all information about the used approaches and their metrics.

<h3 id="d-submission" >d. Submission</h3>

Share your program on your private Git repository with your reviewer to submit it. This repository should contain your working program and text explanation.


<h2 id="chapter-iii" >Chapter VI </h2>
<h3 id="bonus-part" >Bonus Part</h3>

1. Estimate the training time and work time of each classifier, formulate conclusions about their computational efficiency
2. Modify the self-constructed convolutional neural network architecture (change the number of convolutional layers, the number of filters on the layer, the size of the kernel convolutions, the number of neurons on fully connected layers, activation functions, loss functions, learning rate), depend on learning rate and classification from these parameters
3. Balance classes:
    1. Find a class that prevails over the rest
    2.  Take fewer images from the prevailing class so that the classes have approximately the same number of images
    3.  Re-form the sample for training and for control
    4.  On a balanced sample, retrain the best classical classifier (for example, KNN) and the best self-constructed convolutional neural network
4. Formulate final conclusions and recommendations for solving the classification problem radiographs of the lungs: choose the best classifier, justify your choice, describe the main steps of the algorithm for solving the problem



<h2 id="chapter-vii" >Chapter VII</h2>
<h3 id="submission-and-peer-correction" >Submission and Peer-Correction</h3>

Submit your private Git repository as usual. Only the content of your private git repository will be graded.

Here are the points that your peer-corrector will have to check:

- if all the approaches are tried (classifiers, neural networks),
- if the almost perfect accuracy achieved on the test dataset,
- if the research diary exists.
