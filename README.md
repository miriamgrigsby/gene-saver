# gene-saver
Gene Saver

TODO:
- deploy

Thoughts to track:
- ec2 instead of eks because it was free, but not recommended 
- should use terraform so changes are not overwritten/lost
- could use a rds proxy for scalability or rds aurora + proxy to get reader and writer nodes (again for scalability or data concerns)
- security right now is not the concern, make things work, make it better. figure out vpc/subnets/security groups once connectivity is confirmed
