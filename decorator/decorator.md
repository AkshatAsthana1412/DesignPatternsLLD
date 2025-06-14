## Decorator Pattern

- The decorator pattern is a structural pattern that allows adding new behavior to exisiting objects dynamically by placing them inside special wrappers called decorators.
- It enforces Open closed principle.

### Components in decorator pattern:
- **Component interface**: This defines the behavior of the base component.
- **Concrete component**: The concrete class/struct that implements the component interface which can later be decorated to augment its behavior.
- **Decorators**: Concrete decorators which can be added or removed on the concrete component to augment its behavior without modifying the base concrete component.
  
## Benefits of Decorator pattern

- Enforces Open closed principle which states that code should be open for extension and closed for modifications.
- Favors composition over inheritance.
- Reusable code
- Provides flexibility as you can add any number of new behaviors to the existing concrete component, as well as remove them at runtime.
