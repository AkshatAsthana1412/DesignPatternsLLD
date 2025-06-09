## When to use aggregation (has-a relationship):

* When an object can exist **independently** from the container
* When designing loosely couples systems
* When different objects need to be shared across multiple containers
* When following **SOLID** principles, especially Dependency inversion principle(DIP)
* For better maintainability, changes in one struct do not heavily impact the other, making the codebase easier to modify and extend

---
Examples:
school and teachers, employees and company naturally fit the aggregation model.