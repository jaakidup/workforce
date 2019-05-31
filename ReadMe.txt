# Workforce

Simplifying the creation of go routines to execute a given function.
 
The workforce comprises of a Delegator, Workers and a Receiver.

The Delegator hands out the jobs to workers.  

The Workers handle the Task as handed to them by the Delegator.
Workers run in separate go-routines, making use of multi-core systems. Upon completion or failure, the job gets handed to the Receiver.

The Receiver takes all the jobs and reports back to the Workforce with the completed list. 