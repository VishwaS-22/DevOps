/*
Instructions

You are an avid bird watcher that keeps track of how many birds have visited your garden.Usually you use a tally in a notebook to count the birds, 
but to better work with your data, you've digitalized the bird counts for the past weeks.

1. Determine the total number of birds that you counted so far
Implement a function TotalBirdCount that accepts a slice of ints that contains the bird count per day. It should return the total number of birds that you counted.

2. Calculate the number of visiting birds in a specific week
Implement a function BirdsInWeek that accepts a slice of bird counts per day and a week number.
It returns the total number of birds that you counted in that specific week. You can assume weeks are always tracked completely.

3. Fix a counting mistake
Given this new information, write a function FixBirdCountLog that takes a slice of birds counted per day as an argument and returns the slice 
after correcting the counting mistake.
*/

package birdwatcher

// TotalBirdCount return the total bird count by summing the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    t:=0;
    for i:=0;i<len(birdsPerDay);i++{
          t=t+birdsPerDay[i];
    }
    return t;
}

// BirdsInWeek returns the total bird count by summingonly the items belonging to the given week.

func BirdsInWeek(birdsPerDay []int, week int) int {
    t:=0;
    for i:=(week-1)*7;i<(week*7);i++{
        t=t+birdsPerDay[i];
    }
    return t;
}

// FixBirdCountLog returns the bird counts after correcting the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    for i:=0;i<len(birdsPerDay);i+=2{
        birdsPerDay[i]++;
     }
     return birdsPerDay; 
}
