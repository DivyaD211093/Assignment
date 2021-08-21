# Assignment
# Problem statement
Build a sample Go app where you have a list of users, each with properties like Name, email, mobile, etc. 
Each user also has their preferred way of getting notification from app for. e.g. via email, SMS, call, etc.
 
Implement dummy methods for each of these notifications and demonstrate how they will be called for each user based on their choice of mode of notification.
 
You can leave out actual details about implementation of notifications via email, SMS, call.
Remember that focus is to be given on - 
 
1. selecting Go data type(s) to implement different notification types/pattern.
2. making use of Go-routines and appropriate communication mechanism to ensure no more than ‘10’ users are being processed at a time.
3. handling retry mechanism for any failure 

# solution
- I have created a consol based application for which I have taken a list of users from json in the hardcoded format. 
- As per the requirements, the design pattern that I have used is factory pattern.
- Have created a user array and did the unmarshalling on the json. Along with that have created two channels namely jobs and results.
- Have created 5 jobs, and created a workerpool. To control the events hapenning I have applied a rate limiter to 10.
- In the worker function I have programed the notification mode and have print the message in that function for notifications. 
- I have created an interface and supporting fuctions and methods for designing.  
