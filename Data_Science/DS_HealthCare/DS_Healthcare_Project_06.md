# Project 06 – Medical image segmentation

После выполнения заданий проекта просим заполнить [форму обартной связи](https://forms.gle/b8p1ywjE6YtP6QYBA)

Semantic Segmentation for Medical Images Using Machine Learning

Summary: This project is an introduction to machine learning application for Semantic Segmentation


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

Image segmentation is the process of searching and forming groups of pixels, which are formed taking into account the context of the original image (grouping pixels "by meaning").

In the sub-domain of “computer vision,” segmentation is the process of dividing a digital image into multiple segments (many pixels, also called superpixels).

The purpose of segmentation is to simplify and / or modify the representation of an image so that it is easier and easier to analyze. Image segmentation is commonly used to highlight objects and boundaries (lines, curves, etc.) in images. More specifically, image segmentation is the process of assigning labels to each pixel in an image such that pixels with the same labels have common visual characteristics.

Image segmentation is one of the stages in solving various problems of image analysis. For example, when analyzing medical images, segmentation is used to isolate tumors and other pathologies, determine the volume of tissues, when solving surgery problems using a computer, study the anatomical structure.






<h2 id="chapter-ii" >Chapter II </h2>
<h3 id="introduction">Introduction</h3>

The objects of interest of the researcher in medical images used in early diagnosis are usually small and low-contrast in comparison with the surrounding background. Visually detecting these objects — taking the first step towards diagnosing a medical image — can cause problems. On the one hand, they are due to the indicated features of the images themselves, on the other hand, to the limited characteristics of the researcher's visual system and distortions that arise in the images during their acquisition and display.

Fully convolutional networks are a special type of artificial neural networks, the result of which is a segmented image of the original, where the required elements have already been selected in the required way.

Fully convolutional neural networks are used for tasks in which it is necessary, for example, to determine the shape and location of the desired object or several objects. Such tasks are problematic to solve using simple convolutional neural networks. For a general understanding of why and when it is better to use fully convolutional networks instead of conventional convolutional networks, it is necessary to compare the indicated types of neural networks.

U-Net is one of the fully convolutional neural network architectures designed for image segmentation (originally for biomedical images). The network architecture is a sequence of convolution + pooling layers, which first reduce the spatial resolution of the image, and then increase it, having previously combined the images with the data and passed it through other convolution layers. Thus, the network acts as a kind of filter.

In this work, you have a dataset of medical images and their markup. The main task is to write a program for segmenting the snapshot. To do this, you need to analyze the available data, build a neural network and evaluate the effectiveness of its work. 

This problem is actually a semantic segmentation problem, so a wide variety of fully convolutional neural networks can be used to solve it.We invite you to try to build your own neural networks yourself, compare their efficiency and conclude which one is best for your task.  We hope you enjoy it.


<h2 id="chapter-iii" >Chapter III </h2>
<h3 id="goals" >Goals</h3>

The goal of this project is to give you an example of using a machine learning approach to solve a semantic segmentation problem. You will try different algorithms to segment medical images like an expert. In fact, you can use your program to create a support system for finding anatomical and pathological features in medical images. 


<h2 id="chapter-iv" >Chapter IV</h2>
<h3 id="general-instructions" >General Instructions</h3>

- This project will only be evaluated by humans. You are free to organize and name your files as you desire.
- Use Python as a programming language and any libraries and packages supported.
- Use Google Colab, Jupyter or PyCharm as a development environment.
- Write your program so that other people can understand it.
- Store the dataset in your Google Drive or locally to access it from your program.

<h2 id="chapter-v" >Chapter V. Mandatory Part</h2>
<h3 id="a-dataset" >a. Dataset</h3>

You will work with open dataset “Medical Segmentation Decathlon”:

<https://drive.google.com/file/d/1wEB2I6S6tQBVEPxir8cA5kFB8gTQadYY/view?usp=sharing>

The complete dataset consists of files with the nii extension. The original snapshot set is in the imagesTr folder, and the segmentation masks are in the labelsTr folder. Data was obtained from King's College London. Consists of 20 3D images of 20 patients. They have the Mono-modal MRI modality. The size of one slice is 320 X 320 pixels. The masks are markings where the left atrium is highlighted. You have to train and test your model on this data. 

<h3 id="b-task" >b. Task</h3>

1. Data Analysis
   1. Unpack the source archive. Read nii files. 
   1. Convert 3D volume to a series of flat images.
   1. Visualize images and masks of a single patient in a grid view .
1. Data Preparation
   1. Divide the sample into training and validation sets.
   1. Make data suitable for a neural network..
1. Pyramid Scene Parsing Network
   1. Build a basic PSPNet model.
   1. Use the Dice score to evaluate accuracy.
   1. Visualize dependence of accuracy and loss function on the epoch number.
   1. Save and load the trained models. 
   1. Calculate accuracy for test set.
   1. Segment the data for one patient. Visualize the result.
1. Feature Pyramid Network
   1. Build a basic FPN model.
   1. Use the Dice score to evaluate accuracy.
   1. Visualize dependence of accuracy and loss function on the epoch number.
   1. Save and load the trained models. 
   1. Calculate accuracy for test set.
   1. Segment the data for one patient. Visualize the result.
1. U-net
   1. Build a basic  U-net model.
   1. Use the Dice score to evaluate accuracy.
   1. Visualize dependence of accuracy and loss function on the epoch number.
   1. Save and load the trained models. 
   1. Calculate accuracy for test set.
   1. Segment the data for one patient. Visualize the result.

<h3 id="c-implementation" >c. Implementation</h3>

You can work in your private Git repository so a reviewer can access it.

You can use any library or any framework you want.

You should keep a research diary with all information about the used approaches and their metrics.

<h3 id="d-submission" >d. Submission</h3>
Share your program on your private Git repository with your reviewer to submit it. This repository should contain your working program and text explanation.


<h2 id="chapter-iii" >Chapter VI </h2>
<h3 id="bonus-part" >Bonus Part</h3>

1. Visualize results as animation for one patient .
2. Increase the sample using data augmentation.
3. Select model training hyperparameters to increase their training efficiency.
4. Calculate the training and execution time.
5. Make changes to the architecture of the Unet network, identify their impact on improving the accuracy of predictions.
6. Make final conclusions and recommendations for solving the problem of diagnosing diseases by symptoms: choose the best classifier, justify your choice, describe the main steps of the algorithm for solving the problem.





<h2 id="chapter-vii" >Chapter VII</h2>
<h3 id="submission-and-peer-correction" >Submission and Peer-Correction</h3>

Submit your private Git repository as usual. Only the content of your private git repository will be graded.

Here are the points that your peer-corrector will have to check:

- if all the approaches are tried (classifiers, neural networks),
- if the almost perfect accuracy achieved on the test dataset,
- if the research diary exists.
