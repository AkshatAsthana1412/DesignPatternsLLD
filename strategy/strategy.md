## Strategy Design Patttern
 Strategy design pattern allows to turn a set of behaviors into objects which can be used
interchangeably by the context object.

### Components of strategy pattern:

- **Strategy Interface**: 
  - It is an abstraction that defines the key contracts/method signatures for different concrete strategies/algorithms. 
  - By providing a common interface it allows seamless interchange of concrete strategies at runtime.
- **Concrete Strategies**:
  - These are the different algorithms that can be used to solve the specific problem.
  - They implement the strategy interface.
- **Context Object**:
  - This is the component that employes the various strategies.
  - It exposes methods to set a strategy at run time and carry it out. 

## Benefits of strategy pattern:

* It promotes separation of concerns.
* Makes it simple to extend, modify the strategies without modifying the context class.
* When modifying or adding new algorithms, one only needs to work with the concrete strategy classes without touching the context class.

