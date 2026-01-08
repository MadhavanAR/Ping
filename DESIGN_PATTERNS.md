# Design Patterns & Best Practices

## Overview

This document outlines the design patterns and best practices used throughout the Ping codebase, following FAANG-level engineering standards.

## Design Patterns

### 1. Service Layer Pattern

**Purpose**: Encapsulate business logic in dedicated services

**Implementation**:
```go
type UserService struct {
    store    store.UserStore
    config   func() *model.Config
    license  func() *model.License
}

func (s *UserService) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
    // Business logic here
}
```

**Benefits**:
- Separation of concerns
- Testability
- Reusability
- Single Responsibility Principle

### 2. Repository Pattern

**Purpose**: Abstract data access layer

**Implementation**:
```go
type UserStore interface {
    Get(ctx context.Context, userID string) (*model.User, error)
    Save(ctx context.Context, user *model.User) (*model.User, error)
    Delete(ctx context.Context, userID string) error
}
```

**Benefits**:
- Database independence
- Easy to mock for testing
- Centralized data access logic

### 3. Dependency Injection

**Purpose**: Invert dependencies and improve testability

**Implementation**:
```go
type ServiceConfig struct {
    UserStore    store.UserStore
    ConfigFn     func() *model.Config
    LicenseFn    func() *model.License
}

func New(c ServiceConfig) (*Service, error) {
    // Validate config
    // Create service with dependencies
}
```

**Benefits**:
- Loose coupling
- Easy testing with mocks
- Flexible configuration

### 4. Factory Pattern

**Purpose**: Create objects without specifying exact classes

**Implementation**:
```go
func NewService(config ServiceConfig) (Service, error) {
    switch config.Type {
    case "trial":
        return NewTrialService(config)
    case "production":
        return NewProductionService(config)
    }
}
```

### 5. Strategy Pattern

**Purpose**: Define a family of algorithms and make them interchangeable

**Implementation**:
- Feature flags for different behaviors
- Multiple cache implementations (LRU, Redis)
- Different search engines (Elasticsearch, Bleve)

### 6. Observer Pattern

**Purpose**: Notify multiple objects about state changes

**Implementation**:
- Config listeners
- License listeners
- Cluster event handlers

### 7. Decorator Pattern

**Purpose**: Add behavior to objects dynamically

**Implementation**:
- Store layers (Timer → Retry → Cache → Search → SQL)
- Each layer wraps the next, adding functionality

### 8. Singleton Pattern

**Purpose**: Ensure only one instance exists

**Implementation**:
- Platform service
- Cache provider
- Search engine broker

## Architectural Patterns

### Layered Architecture

```
API Layer (HTTP handlers)
    ↓
Business Logic Layer (Services)
    ↓
Data Access Layer (Repositories/Stores)
    ↓
Infrastructure Layer (Database, Cache, etc.)
```

### Clean Architecture

- **Entities**: Core business objects (User, Team, Channel, Post)
- **Use Cases**: Business logic (CreateUser, SendMessage)
- **Interface Adapters**: Controllers, Presenters, Gateways
- **Frameworks**: HTTP, Database, Web Framework

### Domain-Driven Design (DDD)

- **Bounded Contexts**: Clear domain boundaries
- **Aggregates**: Root entities with invariants
- **Value Objects**: Immutable domain concepts
- **Domain Services**: Business logic outside entities

## Best Practices

### Error Handling

1. **Structured Errors**:
```go
type AppError struct {
    Id            string
    Message       string
    DetailedError string
    StatusCode    int
    Where         string
}
```

2. **Error Wrapping**:
```go
if err != nil {
    return nil, errors.Wrap(err, "failed to create user")
}
```

3. **Error Context**:
```go
return model.NewAppError(
    "CreateUser",
    "api.user.create_user.error",
    map[string]any{"user_id": userID},
    "",
    http.StatusBadRequest,
)
```

### Context Propagation

Always pass context through the call chain:
```go
func (s *Service) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
    // Use ctx for cancellation, timeouts, tracing
}
```

### Logging

1. **Structured Logging**:
```go
logger.Info("User created",
    mlog.String("user_id", userID),
    mlog.String("email", email),
    mlog.Int("duration_ms", duration),
)
```

2. **Log Levels**:
- DEBUG: Detailed information for debugging
- INFO: General informational messages
- WARN: Warning messages for potential issues
- ERROR: Error messages for failures

### Testing

1. **Unit Tests**: Test individual functions/methods
2. **Integration Tests**: Test component interactions
3. **E2E Tests**: Test complete workflows

### Code Organization

1. **Package Structure**:
```
app/
  ├── users/        # User domain
  ├── teams/        # Team domain
  ├── trial/        # Trial domain
  └── health/       # Health checks
```

2. **File Naming**:
- `service.go`: Main service implementation
- `service_test.go`: Tests
- `types.go`: Type definitions
- `errors.go`: Error definitions

### Performance

1. **Caching**: Multi-layer caching strategy
2. **Connection Pooling**: Database connection management
3. **Batch Operations**: Group similar operations
4. **Lazy Loading**: Load data only when needed

### Security

1. **Input Validation**: Validate all user inputs
2. **Output Encoding**: Encode outputs to prevent XSS
3. **Authentication**: Secure authentication mechanisms
4. **Authorization**: Role-based access control
5. **Rate Limiting**: Prevent abuse

## Code Quality

### SOLID Principles

1. **Single Responsibility**: Each class/function has one reason to change
2. **Open/Closed**: Open for extension, closed for modification
3. **Liskov Substitution**: Subtypes must be substitutable
4. **Interface Segregation**: Many specific interfaces > one general
5. **Dependency Inversion**: Depend on abstractions, not concretions

### DRY (Don't Repeat Yourself)

- Extract common logic into functions
- Use interfaces for shared behavior
- Create utility packages for common operations

### KISS (Keep It Simple, Stupid)

- Prefer simple solutions over complex ones
- Avoid premature optimization
- Write readable, maintainable code

## References

- [Clean Code](https://www.amazon.com/Clean-Code-Handbook-Software-Craftsmanship/dp/0132350882)
- [Design Patterns](https://www.amazon.com/Design-Patterns-Elements-Reusable-Object-Oriented/dp/0201633612)
- [Domain-Driven Design](https://www.amazon.com/Domain-Driven-Design-Tackling-Complexity-Software/dp/0321125215)

