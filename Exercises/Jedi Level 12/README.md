# Jedi Level 12

Link: https://www.udemy.com/course/learn-how-to-code/learn/lecture/11922408#overview

The videos and pdf of Section 27 call them  Ninja levels, I'm going with Jedi in my directories

## Hands-On 

Only 1 had a video and code, 2 and 3 are steps for github and cmd

**1** 
- Create a dog package. The dog package should have an exported func “Years” which takes human years and turns them into dog years (1 human year = 7 dog years). Document your code with comments. Use this code in func main.
    - run your program and make sure it works
    - run a local server with godoc and look at your documentation.
        - Ran these commands in the Level 12 dir:
            - **go install golang.org/x/tools/cmd/godoc**   // had to install godoc
            - **godoc -http=:6060**
        - Then went to chrome and looked up "http://localhost:6060/pkg/" and CTRL+F "dog"

**2** 
- Push the code to github. Get your documentation on godoc.org and take a screenshot. Delete
your code from github. Refresh godoc.org so that it no longer has your code.

**3** 
- Use godoc at the command line to look at the documentation for:
    - fmt
    - fmt Print
    - strings
    - strconv