# Gene Saver

## TODO / In Progress:

- finish deployment
- pull data from api for genetics overview and store in RDS via migration
- create first golang endpoint for genetics overview homepage

## Rough Decisions/Thoughts to track:

- ec2 instead of eks because it was free, but not recommended for managing k8s clusters
- should use terraform so changes are not overwritten/lost, but faster to spin up in aws cli
- could use an rds proxy for scalability or rds aurora + proxy to get reader and writer nodes (again for scalability or data quality concerns)
- security right now is not the concern, make things work, make it better. figure out vpc/subnets/security groups once connectivity is confirmed and implement ENV vars 
- gin for an http web framework (experience, speed)
- pgx driver for postgresql
- vuetify for mui alternative since mui was made primarily for react
- skip auth but would use something like auth0 or cognito for app auth
- skip api gateway for time, but would really prefer one for the following reasons:
  - security and access control
  - caching
  - transformations (gql to rest or vice versa) allowing for greater variety of architecture
  - rate limiting
  - reduce coupling between the api and the ui

## Completed Elements:

- stood up EC2 instances for my master and worker k8s nodes
- stood up RDS instance and connected it to my EC2 master instance (in the same VPC, but experiencing a inbound rule issue, tbd)
- dockerize the vue/golang app and local postgres via docker volume, once I get the vue and golang apps connected, I should be able to update the docker-compose with any missing elements so I can dev more quickly
- added Taskfiles for easier dev
- added pre-commit hooks for golang/vue code for consistency


## Project MVP:
************

<img width="1057" alt="Screenshot 2024-10-07 at 7 11 49 PM" src="https://github.com/user-attachments/assets/3e2be139-eeac-4ed9-a188-0d13947eb235">

