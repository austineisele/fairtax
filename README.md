# Initial Notes For System

## Project Description
This is a web application whose purpose is to dispalay specific data for citizens of villages and towns to see how their homes are assessed, and how their taxes determined. In the end it's relatively simple: it's a data display.

## The Framework
I'll be using golang and htmx, perhaps with templ. Or maybe svelte. Also MariaDb - because I want to learn these things.

## The Data
This should all be data that is publically available. As of yet, I don't know if there are apis or just static data.
### 1. Showing what? 
The basic thing I want to show is the current assessed value of a house, and what it's current taxes are, <i>in context</i>. This is really imporant. For every house in our village, or in any town/village of a certain zip code.
### 2. Initial thoughts on entities. 
Here are some of the entities I believe will be a part of this. 
#### a. Home: this is the basic entity. Data required:
##### i. Address
##### ii. Parcel #
##### iii. Year Build
##### iv. Square Foot
##### v. Current Assessment
##### vi. Current Taces
##### vii. Assessment History
##### viii. Sales History
#### b. Current Taxes: what is the current state of the home entity?
##### i. Village taxes
##### ii. School taxes
##### iii. Town taxes
##### iv. County taxes
#### c. Assessment History: how did the current assessment history come about? 
##### i. Date Assessed
##### ii. New value
##### iii. Delta value
#### d. Tax History: how did the current tax state come about?
##### i. Date Changed
##### ii. New value
##### iii. Delta value
#### e. Sales History
##### i. Date Sold
##### ii. Sale Price
