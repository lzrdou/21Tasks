# Project 02 – Predicting heart rate as a time series

После выполнения задания просим заполнить [форму отзывов](https://forms.gle/5FzDxHqEGvpuisnB9)

Forecasting Heart Rate Using Statistical Models and Machine Learning

Summary: This project is an introduction to statistical models and machine learning application for automating forecasting of heart rate


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
Time series is a relevant tool used in a variety of solutions, from predicting stock prices, weather forecasts, business planning, to resource allocation. Despite the fact that forecasting can be reduced to the construction of a controlled regression, there are features associated with the temporal nature of observations that must be taken into account using special tools.

A time series is a statistical material collected at different points in time about the value of any parameters (in the simplest case, one) of the process under study. Each unit of statistical material is called a dimension or count. The time series for each sample must indicate the measurement time or the number of measurements in order.

In medical practice, time series forecasting can be used to predict the course of various diseases. Based on the measurements collected over a period of time, future illnesses and complications can be predicted. So with the help of predicting the heart rate, it is possible to detect in advance dangerous dynamics, which may be the result of an exacerbation of a particular disease. 




<h2 id="chapter-ii" >Chapter II </h2>
<h3 id="introduction">Introduction</h3>

Like most other types of analysis, time series analysis assumes that the data contains a systematic component (usually with several components) and random noise (error) that makes it difficult to detect regular components. Most time series research methods include various noise filtering techniques to make the regular component more clearly visible. Most of the regular components of a time series belong to two classes: they are either a trend or a seasonal component. A trend is a common systematic linear or non-linear component that can change over time. The seasonal component is a recurring component. Both of these types of regular components are often present in a number at the same time. 

Among the popular ways to predict time series are statistical models. Such models are ARIMA and SARIMA. ARIMA - autoregressive integrated moving average - model and methodology for time series analysis. It is an extension of ARMA models for non-stationary time series, which can be made stationary by taking differences of some order from the original time series (the so-called integrated or difference-stationary time series). ARIMA's approach to time series is that the stationarity of the series is assessed first. Various tests reveal the presence of unit roots and the order of integration of the time series (usually limited to the first or second order). Further, if necessary (if the order of integration is greater than zero), the series is transformed by taking the difference of the corresponding order, and already for the transformed model, a certain ARMA-model is constructed, since it is assumed that the resulting process is stationary, in contrast to the original non-stationary process (difference-stationary or integrated process of order d). Seasonal Autoregressive Integrated Moving Average, SARIMA or Seasonal ARIMA, is an extension of ARIMA that explicitly supports univariate time series data with a seasonal component. It adds three new hyperparameters for specifying autoregressive (AR), difference (I), and moving average (MA) for the seasonal component of the series, as well as an additional parameter for the seasonality period (S).

Another prediction method is the use of recurrent neural networks. In particular, the use of the LSTM architecture. The state of an LSTM network is represented through a state space vector. This method allows you to track the dependence of new observations on the past (even very distant ones). 

In this work, you have a dataset of heart rate. The main task is to write a program to predict the future heart rate. To do this, you need to analyze the available data, build a forecasting model and evaluate the effectiveness of its work.

We invite you to try different time series forecasting models, compare their performance and conclude which one is best for your task. We hope you enjoy it. 


<h2 id="chapter-iii" >Chapter III </h2>
<h3 id="goals" >Goals</h3>

The goal of this project is to give you an example of using a machine learning approach to solve a heart rate prediction problem. You will try different algorithms to predict your heart rate. You can actually use your program for the creation of a heart rate prediction system. 


<h2 id="chapter-iv" >Chapter IV</h2>
<h3 id="general-instructions" >General Instructions</h3>

- This project will only be evaluated by humans. You are free to organize and name your files as you desire.
- Use Python as a programming language and any libraries and packages supported.
- Use Google Colab, Jupyter or PyCharm as a development environment.
- Write your program so that other people can understand it.
- Store the dataset in your Google Drive or locally to access it from your program.

<h2 id="chapter-v" >Chapter V. Mandatory Part</h2>
<h3 id="a-dataset" >a. Dataset</h3>

You will work with open dataset “Heart Rate Oscillations during Meditation”:

<https://physionet.org/content/meditation/1.0.0/>

Complete dataset consists Heart rate time series for 5 different groups of healthy subjects can be found in text form in the subdirectories of this one:

- chi: Chi meditation group. There are two time series for each of the eight subjects (C1, C2, ... C8), denoted by record names with the suffix pre for the pre-meditation period and med for the meditation period. Each series is about one hour in duration.
- yoga: Kundalini Yoga meditation group. As for the Chi group, there are pre and med series for each of the four subjects (Y1, Y2, Y3, Y4). Durations range from 17 to 47 minutes.
- normal: Spontaneous breathing group (N1, N2, ... N11). Volunteers were recorded while sleeping. Durations are 6 hours each, except for N3 (4.6 hours).
- metron: Metronomic breathing group (M1, M2, ... M14). Volunteers were recorded while supine and breathing at a fixed rate of 0.25 Hz for 10 minutes. 
- ironman: Elite athletes (I1, I2, ... I9). Subjects participated in the Ironman Triathlon; the recordings were obtained during sleeping hours before the event. Durations range from 1 to 1.7 hours.

Each set of three files has a common record name composed of a group identifier and a subject identifier (number). Although all beats are marked normal in these recordings, there may be small numbers of abnormal beats. The first column in each text file is the elapsed time since the beginning of the recording (in seconds), and the second column is the instantaneous heart rate (in beats/minute).

<h3 id="b-task" >b. Task</h3>

1. Data Analysis
   1. Select one of the files in the “normal” folder. Read source data. Find out what they consist of, what data are needed to solve the problem.
   1. Convert seconds to Datetime format.
   1. Visualize a time series.
1. Data Preparation
   1. Resampling data at minute intervals. 
   1. Split the resulting data into training and test samples.
   1. Visualize a resulting  time series.
1. Statistical Models
   1. Build a statistical model ARIMA.
   1. Find hyperparameters for best ARIMA model. 
   1. Build a statistical model SARIMA.
   1. Find hyperparameters for best SARIMA model.
1. Neural Networks
   1. Build and train a simple LSTM neural network.
   1. Visualize dependence of accuracy and loss function on the epoch number.
1. Model Evaluation
   1. For each of the above methods, evaluate the RMSE and MAE.
   1. Calculate Training time and Prediction time.
   1. Visualize a forecasting time series and ground truth time series.

<h3 id="c-implementation" >c. Implementation</h3>

You can work in your private Git repository so a reviewer can access it.

You can use any library or any framework you want.

You should keep a research diary with all information about the used approaches and their metrics.

<h3 id="d-submission" >d. Submission</h3>
Share your program on your private Git repository with your reviewer to submit it. This repository should contain your working program and text explanation.


<h2 id="chapter-iii" >Chapter VI </h2>
<h3 id="bonus-part" >Bonus Part</h3>

1. Select one subject of the files in the “chi” or “yoga” folder. Build forecasting models for pre and meditation time series. Evaluate the RMSE and MAE. Visualize a forecasting time series and ground truth time series.
2.  Make changes to the architecture of the LSTM network, identify their impact on improving the forecasting.
3. Make final conclusions and recommendations for solving the problem of heart rate predictions: choose the best model, justify your choice, describe the main steps of the algorithm for solving the problem.



<h2 id="chapter-vii" >Chapter VII</h2>
<h3 id="submission-and-peer-correction" >Submission and Peer-Correction</h3>

Submit your private Git repository as usual. Only the content of your private git repository will be graded.

Here are the points that your peer-corrector will have to check:

- if all the approaches are tried (classifiers, neural networks),
- if the almost perfect accuracy achieved on the test dataset,
- if the research diary exists.
