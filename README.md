# CS3012 - Software Engineering

## Overall Thoughts
I had started the project previously in another repo based in python however I decided I wanted to challenge myself further and attempt this assignment in Go. I was trying to decide between Go and Rust however went for Go as it has wider community support and also AWS seems to offer Go as a language for building serverless infastructure, a topic I am interested in.

## What I consider to be the phases of the project

- ### Building my Code Pipeline
  I really want to implement a complete codepipeline for this project. To start focusing on a really simple Go setup with a basic function and a unit test for that function,then connect that to a continous integration platform such as Travis CI. After I complete the main part of this assignment, and if I get time I would also like to try and expand on this further so my assignment can become an API Gateway on AWS as serverless code. So then when I deploy to master, the code goes through testing then is deployed on AWS. It is short deadline and I worry implementing this could require a lot of "yak shaving".
  
  
- ### Building up User Acceptance Criteria
  Next my aim would be able to build up a list of functional criteria my solution must complete, so I can establish a definition of done for this project. E.g. Ensuring that if my user acceptance criteria defines a green bike, I don't deliver a blue car. If I complete my user acceptance criteria well enough, creating tests should be as simple as translating the english user acceptance criteria to code.
  
  
- ### Focus on MVP - Agile
  Whilst keeping in mind this project has a limited timescale, I would love to follow the methodologies of agile development by focusing on the MVP first and iterating, I added the idea of deployment to AWS in order to give another iteration if I am able to get the time to do so.
  
## Known Unknowns
- Could run into issues using Golang for the first time
- How much work will be given from other modules
- Difficulty in deploying serverless archiecture in Go
- Making APIs in Golang specifically

## MoSCoW - User Acceptance Criteria
### Must
- Be able to generate a tree of given nodes
- Be able to search tree for LCA of two children nodes
### Should
- Be able to handle all invalid input as nodes
- Be able to handle invalid operations
### Could
- Have all tests be run with travis ci upon each commit to master
- Have 100% code coverage
### Would
- Have an API version capable of input and output via JSON
- Be deployed from travis ci to AWS API Gateway
