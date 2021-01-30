# Engineering Challenge for Golang Candidates

## Introduction

We are looking for engineers who can build web services which are **simple, stable and secure** and who have the user in mind. Good communication skills at Slack and video calls are also required due to the distributed nature of international team.

To test these requirements, we have a challenge which should not take more than 8 hours of your time. If it takes more then please stop at the 8 hour mark. The challenge is testing your communication skills and your coding skills. Please keep us updated in your private channel at our recruiting Slack workspace (icmrecruiting.slack.com) about your progress and feel free to ask questions there. Communication is key!

You should build something with which you are happy with and you think that it can be deployed to production.

## Challenge

### User Management API

Develop an API for user management with the following functionality:

1. create a user
2. delete a user via email
3. return all users sorted by `id` in descending order

The user object should have at least the properties `email`, `password` and `fullname`. The deletion route and the listing route do require authentication and are just allowed to users with an email which ends with "@test.com", see below to learn more.


### Authentication API

Additionally, develop a simple authentication API with the following functionality:

1. sign in a user via email and passsword. Return a Bearer token with a TTL of 24 hours and the logged in user
2. sign out

### Technical Requirements

During your development, always remind youself of developing **simple, stable and secure** software.

1. all functionality should be covered by unit and integration tests
2. errors should be handled including 401s, 404s and 500s
3. don’t forget code comments and logging
4. code as simple / easy to read and functional / modular as possible
5. you can use Golang, Postgres and Redis
6. Use Docker Compose for a local test setup of your solution
7. Publish your code on a public Github repo and use a generic, hard to guess project name
8. Push to Github all the time. We are especially interested in the Git history and not only the final commit

## Review

Please ping us on Slack when done. We will then review your work and invite you to a review meeting. We will discuss your project there and will also give you the opportunity to meet the team in person.
