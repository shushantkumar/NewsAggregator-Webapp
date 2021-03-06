# News Aggregator

This is a Web Application developed using GoLang which collects daily news from The Washington Post website using their sitemap url.<br>

In this web app I have demonstrated how we can use <b>Go Routines</b> to speed up the loading speed of the site by reducing the idle time using <b>Concept of Concurrency</b><br>

Go routines are functions or methods that run concurrently with other functions or methods. Go routines can be thought of as light weight threads. The cost of creating a Go routine is tiny when compared to a thread. On Go we can run hundreds of thousands or millions of go routines.

<h2>How to use</h2>

### Requirements
Go must be installed to run the program.<br>
To install Go follow the official site instructions <a>https://golang.org/doc/install</a>

### Steps

1. Clone the repository


    ```
        https://github.com/shushantkumar/NewsAggregator-Webapp.git
    ```
2. Move into the directory  
    
    ```
        cd NewsAggregator-Webapp
    ```       
3. To run the version with Concurrency implemented 

    ```
        go run final_withconc.go
    ``` 
   and open http://localhost:8000 in your browser  
4. To run the version with Concurrency not implemented 

    ```
        go run before_conc.go
    ```
   and open http://localhost:8000 in your browser  
<br>

<h2> Working </h2>

* The file <b>before_conc.go</b> access sitemaps without any go routines. The idle time while loading was <b>11631.0 ms</b> (avg around 11500 ms for other tries).<br>
* Since the web app accesses hundreds of sitemaps and it sends request for each sitemap and waits for response it wastes a lot of time increasing the idle time.<br> 

![](before_concurrency.png)

* To reduce the loading time I used Go Routines.<br> 
* The file <b>final_withconc.go</b> implements the same function but uses Go Routines for accessing each internal sitemaps, so basically when one request is sent other requests are also sent concurrently without waiting for previous request to get back the response.<br>
* After implementing go routines for this web application the idle time went down to <b>1444.9 ms</b> (avg around 1450 ms for other tries).<br>

![](after_concurrency.jpg)

* The basic concept used here is concurrency. I also used <b>Channels</b> to provide concurrency. It allows goroutines to synchronize without explicit locks or condition variables.<br> 
* Synchronization needs to be there while using go routines. To implement it I imported "sync". Go provides another built in feature <b>defer</b> and functions under sync - Done(), Add() and Wait(). <br>

Sitemap link : <a>https://www.washingtonpost.com/news-sitemap-index.xml</a>

