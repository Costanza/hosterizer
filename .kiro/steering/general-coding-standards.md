# General Language Independent Coding Standards

## General Principles
- Use Clean Architecture Principles at all times
- User a folder named "domain" for all domain related data like models
- Models should always live in a file per model
- Repositories and Services should implement from a consistent interface
- As much as possible prefer to name SQL statments NOT embedded in code
- Prefer handling errors and returning first in functions
- Limit nesting as much as possible
- Prefer composition over inheritence
- Prefer a test driven development mentaility ensuring everything comes with high test coverage