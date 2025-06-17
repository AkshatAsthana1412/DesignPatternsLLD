## Following are the 5 SOLID design principles:

- **Single Responsibility principle**: Each class should do 1 thing properly.
- **Open/closed principle**: Entities should be open for extension and closed for modification.
- **Liskov's substitution principle**: If a class S extends/implements T, then S should be valid wherever T is valid. If a class doesn't need a method, it shouldn't be forced to implement it.
- **Interface segregation principle**: Keep your interfaces focused, short with only the most necessary methods.
- **Dependancy inversion principle**: High level modules should not depend on low level modules and implementation details. E.g. EmailService class(High level) should not be tightly coupled to low level GmailClient class, because if later we have some other method to send emails, then the high level EmailService class will have to be modified.