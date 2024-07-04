## Hikers Programming Question

The Royal Gorge Bridge in Colorado is one of the sceneic bridges in the United 
States. Colorado Parks and Wildlife have been tracking the amount of time hikers
spend on the bridge to better handle traffic flow. 

Write a function that takes in information of all hikers from that day which includes their name, and their time spent on the bridge - that function should returns the most commonly spent amount of time on the bridge in minutes.

Example Data:
Name                Time Spent

Evan P                  10
Leslie K                12
David W                 9
Jeff M                  8  
Crouton L               10
Tin Z                   11

Returns: 10

### Corrupted Data

Something funky happened with the data! 

Write a function that will take in Example Data, and
Removal Data, and return the difference between the two. This function will remove all the hiker's information that is found in "Removal Data" from "Example Data".


Example Data:
Name                Time Spent

Evan P                  10
Leslie K                12
David W                 9
Jeff M                  8  
Crouton L               10
Tin Z                   11

Removal Data:

Name                Time Spent

Evan P                  10
Jeff M                  8  
Tin Z                   11

Returned Data:

Name                Time Spent

Leslie K                12
David W                 9
Crouton L               10