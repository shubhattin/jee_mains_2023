# How it Works

## Where it all started

I made a simple jupyter notebook which analysed the _HTML_ of _Response Sheet_ and _Answer key_ to calculated the result.([here](https://github.com/shubhattin/jee_mains_2023_score_calculator/blob/e60315a4a0004f2483df6dbf697f2d794d6982b5/tests/other/old_using_pandas.ipynb) [here](https://github.com/shubhattin/jee_mains_2023_score_calculator/tree/e60315a4a0004f2483df6dbf697f2d794d6982b5/tests/process_data))

I was also familiar to Web Development, so i thought that I thought to make a graphical interface to help people use it easily. I did not want to spend a lot of time of on frontend, i just wanted to get a decent looking UI and also I am _not very good at frontend_.

## How does it exactly work ??

A very genuine question would be that how can this _app_ calculate the result, how can i have the access to Official DataBase ?  
The Answer is that we have no such access to the any such database.

1. When you try to open the _Result_ for **first time** using your Application Number and Date of Birth, it will open a different page as your **result is currently not stored(cached) in our database**.
2. The Page on which you will be redirected will ask to you enter **HTML/URL** of the **your** Response Sheet(**URL/HTML**) and Answer Key(**HTML**). Then this data can be **used to extract relevant information** to calculate the result(score).(Source Code [here](./api/jee/process_data/get_data.go))
3. Then your Score will be displayed. This **result will be cached** for your convenience.
4. When you will be opening this **another time** you will be able to directly see your result.

## How it makes viewing Result easier

NTA only provides us with an Answer Key which has Question ID which their Correct OptionID, and moreover it is not in same order as that of the actual question paper making it harder to compare.  
Also the Response Sheet provided only has Question's image and ID and options' image and ID, without mentioning the correct answer anywhere.

This tool tries to simplify this task by automating things and **bringing Questins, their Answers', thier IDs' and thier actual Image Questions and Answers together**.

> **_Note :- This app currently might not work with the new released JEE2023 Session 2 Answer Key. Some changes obviously have to be made. It will take atleast 2 hours for me to Update this to work properly._**

## How exactly will my result look like

The App has a option to view **Sample Result**, this is the format how any particular result will be displayed.
