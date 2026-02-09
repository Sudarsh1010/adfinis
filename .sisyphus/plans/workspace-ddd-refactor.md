# Workspace DDD Refactoring

## TL;DR

> **Quick Summary**: Fix workspace DDD implementation by adding missing application layer, fixing file naming/typos, and implementing pragmatic DDD patterns with proper FX integration.
>
> **Deliverables**:
> - Fixed path typos (persistance → persistence, respositories → repositories)
> - Application layer (repository interface, service, DTOs, errors)
> - Files renamed to {domain}.{type}.go pattern
> - Unit tests for entity, repository, service
> - Updated FX module with correct providers
>
> **Estimated Effort**: Medium
> **Parallel Execution**: YES - 2 waves
> **Critical Path**: File renames → Application layer creation → Tests → FX module update

---

## Context

### Original Request
User wants to fix the DDD implementation of workspace module. Issues identified:
- File naming issues (creating duplicates)
- Uber FX integration problems
- Want pragmatic DDD approach (not pure DDD)

### Interview Summary

**Key Discussions**:
- Feature Scope: Backend Only - fix DDD structure + service layer, skip HTTP handlers
- Repository Pattern: Keep Interface for testability
- Naming Style: Use Recommended {domain}.{type}.go pattern (e.g., workspace.entity.go)
- Test Coverage: Yes, add unit tests for all layers

**Research Findings**:
- Pragmatic DDD focuses on solving real problems with less bureaucracy
- Vertical slice architecture recommended for Go + Fx
- Key patterns: Entities, Repository, Domain Services (DDD Lite)
- Naming convention: {domain}.{type}.go pattern
- Fx module pattern: fx.Provide all constructors
- Current codebase uses Bun ORM

**Code Inspection Findings**:
- Domain entity exists: `internal/domain/workspace/workspace.go` (pure, no infrastructure deps)
- Infrastructure exists: `internal/infrastructure/persistance/bun/models/workspace.go`, `internal/infrastructure/persistance/bun/respositories/workspace_repo.go`
- Path typos: "persistance" (should be "persistence"), "respositories" (should be "repositories")
- FX module references non-existent packages: `internal/application/workspace/`, workspace.NewWorkspaceRepository, workspace.NewWorkspaceService, workspace.NewWorkspaceHandler
- Missing: Repository interface, Service layer, Error definitions (ErrInvalidName referenced but not defined)
- No existing tests

### Metis Review

**Identified Gaps (addressed)**:
- Service API surface: Defined as CreateWorkspace, GetWorkspaceByID, ListWorkspaces, UpdateWorkspaceName, DeleteWorkspace
- Business rules: Validation for non-empty workspace name
- Error handling: Custom error types with wrapping
- DTO structure: CreateWorkspaceRequest, WorkspaceResponse
- Test coverage: Unit tests with table-driven patterns

**Guardrails Applied**:
- No HTTP handlers or transport layer (explicitly deferred)
- No database migrations or schema changes
- No integration tests with real database
- No other domain modules touched
- No events or message queues
- Zero regression: preserve all existing workspace entity behavior

---

## Work Objectives

### Core Objective
Refactor workspace module to follow pragmatic DDD principles with proper separation of concerns, fix file naming issues, and complete the missing application layer.

### Concrete Deliverables
- Fixed directory structure: `internal/infrastructure/persistence/bun/` (typos corrected)
- Application layer: `internal/application/workspace/` with repository interface, service, DTOs, errors
- Renamed files: `{domain}.{type}.go` pattern (workspace.entity.go, workspace.service.go, etc.)
- Unit tests: entity tests, repository tests, service tests
- Updated FX module: Correct providers for repository interface and service

### Definition of Done
- [x] All code compiles: `go build ./internal/workspace/...` and `go build ./cmd/server/...`
- [x] All tests pass: `go test ./internal/workspace/... -v -race`
- [x] Test coverage ≥ 70%: `go test ./internal/workspace/... -cover`
- [x] No path typos: `find internal/workspace -type d -name "*persistance*"` returns nothing
- [x] FX module compiles with correct imports
- [ ] Existing workspace entity behavior preserved

### Must Have
- Repository interface defined in application layer
- Service layer with business logic
- DTOs for data transfer (CreateWorkspaceRequest, WorkspaceResponse)
- Custom error types (ErrInvalidName, ErrWorkspaceNotFound)
- Unit tests for entity, repository, service
- Files renamed to {domain}.{type}.go pattern
- Path typos fixed (persistence, repositories)
- FX module updated with correct providers

### Must NOT Have (Guardrails)
- HTTP handlers for workspace
- API endpoints or routes
- Database migrations
- Integration tests with real database
- Other domain modules modified
- Event publishing or message queue integration
- Caching layers
- Complex error codes or internationalization

---

## Verification Strategy (MANDATORY)

> **UNIVERSAL RULE: ZERO HUMAN INTERVENTION**
>
> ALL tasks in this plan MUST be verifiable WITHOUT any human action.
> This is NOT conditional — it applies to EVERY task, regardless of test strategy.

### Test Decision
- **Infrastructure exists**: NO (no tests currently)
- **Automated tests**: YES (TDD) - Tests before implementation
- **Framework**: Go testing package (standard library)

### TDD Enabled

Each TODO follows RED-GREEN-REFACTOR:

**Task Structure:**
1. **RED**: Write failing test first
   - Test file: `{path}_test.go`
   - Test command: `go test {path} -run {TestName}`
   - Expected: FAIL (test exists, implementation doesn't)
2. **GREEN**: Implement minimum code to pass
   - Command: `go test {path} -run {TestName}`
   - Expected: PASS
3. **REFACTOR**: Clean up while keeping green
   - Command: `go test {path}`
   - Expected: PASS (all tests)

**Test Setup Task:**
- No setup needed - using standard Go testing package

### Agent-Executed QA Scenarios (MANDATORY — ALL tasks)

> Whether TDD is enabled or not, EVERY task MUST include Agent-Executed QA Scenarios.
> - **With TDD**: QA scenarios complement unit tests at integration level
> - **Without TDD**: QA scenarios are the PRIMARY verification method
>
> These describe how the executing agent DIRECTLY verifies the deliverable
> by running it — Go tests, compilation checks, file system verification.

**Verification Tool by Deliverable Type:**

| Type | Tool | How Agent Verifies |
|------|------|-------------------|
| **Go Code** | Bash (go build/test) | Compile, run tests, check coverage |
| **File Structure** | Bash (find/ls) | Verify files exist/renamed, old files deleted |
| **FX Module** | Bash (go build) | Verify module compiles with correct imports |

**Each Scenario MUST Follow This Format:**

```
Scenario: [Descriptive name — what is being verified]
  Tool: Bash (go build/test/find/ls)
  Preconditions: [What must be true before this scenario runs]
  Steps:
    1. [Exact command to run]
    2. [Next command with expected output]
    3. [Assertion with exact expected value]
  Expected Result: [Concrete, observable outcome]
  Failure Indicators: [What would indicate failure]
  Evidence: [Output capture / file path]
```

---

## Execution Strategy

### Parallel Execution Waves

```
Wave 1 (Start Immediately):
├── Task 1: Rename domain file (workspace.entity.go)
├── Task 2: Fix infrastructure path typos
├── Task 3: Create error definitions
└── Task 4: Create DTO structures

Wave 2 (After Wave 1):
├── Task 5: Create repository interface
├── Task 6: Implement repository tests (RED)
├── Task 7: Implement repository implementation (GREEN)
└── Task 8: Refactor repository implementation

Wave 3 (After Wave 2):
├── Task 9: Create service interface and implementation tests (RED)
├── Task 10: Implement service (GREEN)
└── Task 11: Refactor service implementation

Wave 4 (After Wave 3):
├── Task 12: Rename entity file in infrastructure
├── Task 13: Rename repository implementation file
├── Task 14: Update FX module
└── Task 15: Final verification and cleanup

Critical Path: Task 1 → Task 5 → Task 6 → Task 7 → Task 9 → Task 10 → Task 14
Parallel Speedup: ~35% faster than sequential
```

### Dependency Matrix

| Task | Depends On | Blocks | Can Parallelize With |
|------|------------|--------|---------------------|
| 1 | None | 5, 9 | 2, 3, 4 |
| 2 | None | 12, 13 | 1, 3, 4 |
| 3 | None | 5, 9 | 1, 2, 4 |
| 4 | None | 9, 10 | 1, 2, 3 |
| 5 | 1, 3 | 6, 9 | - |
| 6 | 5 | 7 | - |
| 7 | 6 | 8 | - |
| 8 | 7 | 12, 13 | - |
| 9 | 5 | 10 | - |
| 10 | 9, 4 | 11 | - |
| 11 | 10 | 12, 13 | - |
| 12 | 2, 8 | 14 | 13 |
| 13 | 2, 8 | 14 | 12 |
| 14 | 12, 13 | 15 | - |
| 15 | All | None | - |

### Agent Dispatch Summary

| Wave | Tasks | Recommended Agents |
|------|-------|-------------------|
| 1 | 1, 2, 3, 4 | task(category="unspecified-low", load_skills=[], run_in_background=false) |
| 2 | 5, 6, 7, 8 | task(category="unspecified-low", load_skills=[], run_in_background=false) |
| 3 | 9, 10, 11 | task(category="unspecified-low", load_skills=[], run_in_background=false) |
| 4 | 12, 13, 14, 15 | task(category="unspecified-low", load_skills=[], run_in_background=false) |

---

## TODOs

- [x] 1. Rename domain file to workspace.entity.go

  **What to do**:
  - Rename `internal/domain/workspace/workspace.go` to `internal/domain/workspace/workspace.entity.go`
  - Update package declaration if needed (should remain `package workspace`)
  - Ensure no other imports break from rename

  **Must NOT do**:
  - Change any domain entity behavior
  - Modify entity fields or methods
  - Touch infrastructure layer

  **Implementation Note**: Use standard Go file rename (`mv` or `git mv`), no content changes needed. Package name remains `package workspace` to match existing conventions.

  **Recommended Agent Profile**:
  - **Category**: `unspecified-low`
    - Reason: Simple file rename operation, Go code structure
  - **Skills**: None needed
  - **Skills Evaluated but Omitted**: git-master (not needed yet)

  **Parallelization**:
  - **Can Run In Parallel**: YES
  - **Parallel Group**: Wave 1 (with Tasks 2, 3, 4)
  - **Blocks**: Task 5 (repository interface references entity)
  - **Blocked By**: None (can start immediately)

  **References**:

  **Pattern References**:
  - `internal/domain/workspace/workspace.go:1-61` - Current domain entity structure

  **API/Type References**:
  - `internal/domain/workspace/workspace.go:Workspace` - Workspace entity type
  - `internal/domain/workspace/workspace.go:WorkspaceID` - Value object type
  - `internal/domain/workspace/workspace.go:NewWorkspace` - Factory function
  - `internal/domain/workspace/workspace.go:Rename` - Domain method

  **Test References**:
  - None exist yet - will create in Task 9

  **Documentation References**:
  - Research findings: Use {domain}.{type}.go naming convention

  **External References**:
  - Go file renaming: `go build ./internal/domain/workspace/...` to verify no import breakage

  **WHY Each Reference Matters**:
  - Current entity structure must be preserved during rename
  - Package name must remain `workspace` to avoid breaking existing code
  - Factory functions and domain methods are critical for service layer

  **Acceptance Criteria**:

  **If TDD (tests enabled):**
  - [ ] File renamed: internal/domain/workspace/workspace.entity.go exists
  - [ ] Old file deleted: internal/domain/workspace/workspace.go does not exist
  - [ ] go build ./internal/domain/workspace/... → PASS (no errors)

  **Agent-Executed QA Scenarios (MANDATORY — per-scenario, ultra-detailed):**

  Scenario: Domain file renamed correctly
    Tool: Bash (ls/find)
    Preconditions: internal/domain/workspace/workspace.go exists
    Steps:
      1. mv internal/domain/workspace/workspace.go internal/domain/workspace/workspace.entity.go
      2. ls internal/domain/workspace/workspace.entity.go
      3. ! ls internal/domain/workspace/workspace.go 2>/dev/null
    Expected Result: workspace.entity.go exists, old file does not exist
    Failure Indicators: Old file still exists, new file doesn't exist
    Evidence: File listing output

  Scenario: Domain package compiles after rename
    Tool: Bash (go build)
    Preconditions: workspace.entity.go exists
    Steps:
      1. go build ./internal/domain/workspace/...
      2. assert exit code is 0
    Expected Result: Successful compilation
    Failure Indicators: Compilation errors
    Evidence: Build output

  **Evidence to Capture**:
  - [ ] Directory listing showing workspace.entity.go
  - [ ] Build output showing successful compilation

  **Commit**: NO (group with Task 14)

- [x] 2. Fix infrastructure path typos

  **What to do**:
  - Rename directory `internal/infrastructure/persistance/` to `internal/infrastructure/persistence/`
  - Rename directory `.../respositories/` to `.../repositories/`
  - Update all import statements in affected files

  **Must NOT do**:
  - Modify any logic in files
  - Change database schema
  - Break existing repository functionality

  **Recommended Agent Profile**:
  - **Category**: `unspecified-low`
    - Reason: Directory rename and import updates, Go code structure
  - **Skills**: None needed

  **Parallelization**:
  - **Can Run In Parallel**: YES
  - **Parallel Group**: Wave 1 (with Tasks 1, 3, 4)
  - **Blocks**: Task 12, 13 (repository file renames)
  - **Blocked By**: None (can start immediately)

  **References**:

  **Pattern References**:
  - `internal/infrastructure/persistance/bun/` - Current directory structure with typos

  **API/Type References**:
  - `internal/infrastructure/persistance/bun/models/workspace.go` - Workspace model
  - `internal/infrastructure/persistance/bun/respositories/workspace_repo.go` - Repository implementation

  **Test References**:
  - None exist yet

  **Documentation References**:
  - Research findings: "persistance" → "persistence", "respositories" → "repositories"

  **External References**:
  - Directory renaming: `mv persistance persistence`

  **WHY Each Reference Matters**:
  - Correct spelling is critical for project maintainability
  - All imports must be updated to prevent compilation errors
  - Files in these directories depend on correct paths

  **Acceptance Criteria**:

  **If TDD (tests enabled):**
  - [ ] Directory renamed: internal/infrastructure/persistence/ exists
  - [ ] Old directory deleted: internal/infrastructure/persistance/ does not exist
  - [ ] Subdirectory renamed: persistence/bun/repositories/ exists
  - [ ] All files compile: go build ./internal/infrastructure/persistence/... → PASS

  **Agent-Executed QA Scenarios (MANDATORY — per-scenario, ultra-detailed):**

  Scenario: Infrastructure directory typos fixed
    Tool: Bash (find/mv)
    Preconditions: internal/infrastructure/persistance/ exists
    Steps:
      1. mv internal/infrastructure/persistance internal/infrastructure/persistence
      2. mv internal/infrastructure/persistence/bun/respositories internal/infrastructure/persistence/bun/repositories
      3. ls internal/infrastructure/persistence/bun/repositories/workspace_repo.go
      4. ! find internal/infrastructure -type d -name "*persistance*" 2>/dev/null
      5. ! find internal/infrastructure -type d -name "*respositories*" 2>/dev/null
    Expected Result: Correct directory names exist, old names don't
    Failure Indicators: Old directories still found, new directories missing
    Evidence: Directory listing output

  Scenario: Infrastructure package compiles after path fixes
    Tool: Bash (go build)
    Preconditions: persistence/ directory exists
    Steps:
      1. go build ./internal/infrastructure/persistence/...
      2. assert exit code is 0
    Expected Result: Successful compilation
    Failure Indicators: Compilation errors due to incorrect imports
    Evidence: Build output

  **Evidence to Capture**:
  - [ ] Directory listing showing correct structure
  - [ ] Build output showing successful compilation

  **Commit**: NO (group with Task 14)

- [x] 3. Create error definitions

  **What to do**:
  - Create `internal/application/workspace/workspace.errors.go`
  - Define `ErrInvalidName` error type
  - Define `ErrWorkspaceNotFound` error type
  - Implement error wrapping helpers if needed
  - Follow Go error best practices

  **Must NOT do**:
  - Create complex error codes
  - Add internationalization
  - Over-engineer error hierarchy

  **Recommended Agent Profile**:
  - **Category**: `unspecified-low`
    - Reason: Simple Go error type definitions
  - **Skills**: None needed

  **Parallelization**:
  - **Can Run In Parallel**: YES
  - **Parallel Group**: Wave 1 (with Tasks 1, 2, 4)
  - **Blocks**: Task 5, 9 (repository and service need errors)
  - **Blocked By**: None (can start immediately)

  **References**:

  **Pattern References**:
  - `internal/domain/workspace/workspace.go:56` - ErrInvalidName referenced

  **API/Type References**:
  - Go error package: `error` interface
  - `fmt.Errorf` - Error formatting

  **Test References**:
  - None exist yet

  **Documentation References**:
  - Research findings: Simple custom error types for pragmatic DDD
  - Go error handling best practices

  **External References**:
  - Go errors: https://go.dev/blog/error-handling-and-go

  **WHY Each Reference Matters**:
  - ErrInvalidName is already referenced but not defined - must provide
  - ErrWorkspaceNotFound needed for repository "not found" case
  - Simple errors are sufficient for pragmatic DDD

  **Acceptance Criteria**:

  **If TDD (tests enabled):**
  - [ ] File created: internal/application/workspace/workspace.errors.go
  - [ ] ErrInvalidName defined and exported
  - [ ] ErrWorkspaceNotFound defined and exported
  - [ ] go build ./internal/application/workspace/... → PASS

  **Agent-Executed QA Scenarios (MANDATORY — per-scenario, ultra-detailed):**

  Scenario: Error definitions created
    Tool: Bash (ls/go build)
    Preconditions: internal/application/workspace/ directory exists
    Steps:
      1. cat internal/application/workspace/workspace.errors.go
      2. grep -q "ErrInvalidName" internal/application/workspace/workspace.errors.go
      3. grep -q "ErrWorkspaceNotFound" internal/application/workspace/workspace.errors.go
      4. go build ./internal/application/workspace/...
    Expected Result: Both errors defined, file compiles
    Failure Indicators: Missing errors, compilation errors
    Evidence: File content and build output

  **Evidence to Capture**:
  - [ ] workspace.errors.go content
  - [ ] Build output

  **Commit**: NO (group with Task 14)

- [x] 4. Create DTO structures

  **What to do**:
  - Create `internal/application/workspace/workspace.dto.go`
  - Define `CreateWorkspaceRequest` struct (Name field)
  - Define `WorkspaceResponse` struct (ID, Name, IsDefault, CreatedAt fields)
  - Add validation if needed (non-empty name)

  **Must NOT do**:
  - Create excessive DTO types
  - Add view models
  - Implement API versioning

  **Recommended Agent Profile**:
  - **Category**: `unspecified-low`
    - Reason: Simple Go struct definitions
  - **Skills**: None needed

  **Parallelization**:
  - **Can Run In Parallel**: YES
  - **Parallel Group**: Wave 1 (with Tasks 1, 2, 3)
  - **Blocks**: Task 9, 10 (service uses DTOs)
  - **Blocked By**: None (can start immediately)

  **References**:

  **Pattern References**:
  - `internal/domain/workspace/workspace.go:17-22` - Workspace entity structure

  **API/Type References**:
  - `internal/domain/workspace/workspace.go:WorkspaceID` - Use for response ID field

  **Test References**:
  - None exist yet

  **Documentation References**:
  - Research findings: DTO pattern for data transfer, validation at service layer

  **External References**:
  - Go struct tags: https://go.dev/blog/json

  **WHY Each Reference Matters**:
  - CreateWorkspaceRequest needed for service input
  - WorkspaceResponse needed for service output
  - Must map entity fields to DTO fields

  **Acceptance Criteria**:

  **If TDD (tests enabled):**
  - [ ] File created: internal/application/workspace/workspace.dto.go
  - [ ] CreateWorkspaceRequest has Name field
  - [ ] WorkspaceResponse has ID, Name, IsDefault, CreatedAt fields
  - [ ] go build ./internal/application/workspace/... → PASS

  **Agent-Executed QA Scenarios (MANDATORY — per-scenario, ultra-detailed):**

  Scenario: DTO structures created
    Tool: Bash (ls/go build)
    Preconditions: internal/application/workspace/ directory exists
    Steps:
      1. cat internal/application/workspace/workspace.dto.go
      2. grep -q "CreateWorkspaceRequest" internal/application/workspace/workspace.dto.go
      3. grep -q "WorkspaceResponse" internal/application/workspace/workspace.dto.go
      4. go build ./internal/application/workspace/...
    Expected Result: Both DTOs defined, file compiles
    Failure Indicators: Missing DTOs, compilation errors
    Evidence: File content and build output

  **Evidence to Capture**:
  - [ ] workspace.dto.go content
  - [ ] Build output

  **Commit**: NO (group with Task 14)

- [x] 5. Create repository interface

  **What to do**:
  - Create `internal/application/workspace/workspace.repository.go`
  - Define `WorkspaceRepository` interface with methods:
    - Save(ctx, *Workspace) error
    - FindByID(ctx, WorkspaceID) (*Workspace, error)
    - FindAll(ctx) ([]*Workspace, error)
    - UpdateName(ctx, WorkspaceID, string) error
    - Delete(ctx, WorkspaceID) error
  - Reference domain entity `workspace.Workspace`

  **Must NOT do**:
  - Implement repository (implementation is separate)
  - Add infrastructure dependencies (Bun, SQL)
  - Create complex query methods

  **Recommended Agent Profile**:
  - **Category**: `unspecified-low`
    - Reason: Go interface definition, DDD patterns
  - **Skills**: None needed

  **Parallelization**:
  - **Can Run In Parallel**: NO
  - **Parallel Group**: Sequential
  - **Blocks**: Task 6 (repository tests reference interface)
  - **Blocked By**: Task 1 (entity renamed), Task 3 (errors defined)

  **References**:

  **Pattern References**:
  - `internal/domain/workspace/workspace.go:17-22` - Workspace entity structure
  - Research findings: Repository pattern for data access abstraction

  **API/Type References**:
  - `internal/domain/workspace/workspace.Workspace` - Entity type for interface methods
  - `internal/domain/workspace/workspace.WorkspaceID` - ID type for FindByID/Delete
  - `internal/application/workspace/workspace.errors.go` - Error types for interface

  **Test References**:
  - None exist yet - will create in Task 6

  **Documentation References**:
  - Research findings: Repository pattern in DDD, Go interfaces

  **External References**:
  - Go interfaces: https://go.dev/tour/methods-and-interfaces

  **WHY Each Reference Matters**:
  - Interface defines contract between application and infrastructure
  - Must reference domain entity to maintain purity
  - Methods match typical CRUD operations needed by service

  **Acceptance Criteria**:

  **If TDD (tests enabled):**
  - [ ] File created: internal/application/workspace/workspace.repository.go
  - [ ] WorkspaceRepository interface defined
  - [ ] Save, FindByID, FindAll, UpdateName, Delete methods defined
  - [ ] go build ./internal/application/workspace/... → PASS

  **Agent-Executed QA Scenarios (MANDATORY — per-scenario, ultra-detailed):**

  Scenario: Repository interface created
    Tool: Bash (ls/go build)
    Preconditions: workspace.entity.go, workspace.errors.go exist
    Steps:
      1. cat internal/application/workspace/workspace.repository.go
      2. grep -q "WorkspaceRepository" internal/application/workspace/workspace.repository.go
      3. grep -q "Save.*context.Context.*\*workspace.Workspace" internal/application/workspace/workspace.repository.go
      4. grep -q "FindByID.*context.Context.*workspace.WorkspaceID" internal/application/workspace/workspace.repository.go
      5. grep -q "FindAll.*context.Context" internal/application/workspace/workspace.repository.go
      6. grep -q "UpdateName.*context.Context.*workspace.WorkspaceID.*string" internal/application/workspace/workspace.repository.go
      7. grep -q "Delete.*context.Context.*workspace.WorkspaceID" internal/application/workspace/workspace.repository.go
      8. go build ./internal/application/workspace/...
    Expected Result: All methods defined, file compiles
    Failure Indicators: Missing methods, wrong signatures, compilation errors
    Evidence: File content and build output

  **Evidence to Capture**:
  - [ ] workspace.repository.go content
  - [ ] Build output

  **Commit**: NO (group with Task 14)

- [x] 6. Implement repository tests (RED)

  **What to do**:
  - Create `internal/infrastructure/persistence/bun/repositories/workspace.repository_impl_test.go`
  - Write tests for all repository methods:
    - TestSave_Success
    - TestFindByID_Success
    - TestFindByID_NotFound
    - TestFindAll_Success
    - TestUpdateName_Success
    - TestUpdateName_NotFound
    - TestDelete_Success
    - TestDelete_NotFound
  - Use table-driven test pattern where appropriate
  - Use simple in-memory slice for mock database (avoid real DB)

  **Must NOT do**:
  - Implement repository (this is test-only task)
  - Touch application layer
  - Use real database connection

  **Implementation Note - Mock Repository Pattern**:
  ```go
  // Create simple mock with in-memory storage:
  type mockWorkspaceRepository struct {
      mu sync.RWMutex
      workspaces []workspace.Workspace
  }

  // Implement WorkspaceRepository interface methods
  func (m *mockWorkspaceRepository) Save(ctx context.Context, ws *workspace.Workspace) error {
      m.mu.Lock()
      defer m.mu.Unlock()
      m.workspaces = append(m.workspaces, *ws)
      return nil
  }
  // ... other methods
  ```

  **Recommended Agent Profile**:
  - **Category**: `unspecified-low`
    - Reason: Go test writing, table-driven patterns
  - **Skills**: None needed

  **Parallelization**:
  - **Can Run In Parallel**: NO
  - **Parallel Group**: Sequential
  - **Blocks**: Task 7 (implementation needed to make tests pass)
  - **Blocked By**: Task 5 (repository interface defined)

  **References**:

  **Pattern References**:
  - Research findings: Table-driven test pattern in Go

  **API/Type References**:
  - `internal/application/workspace/workspace.repository.go` - Repository interface to test against
  - `internal/domain/workspace/workspace.Workspace` - Domain entity to use in tests
  - `internal/infrastructure/persistence/bun/models/workspace.Workspace` - Infrastructure model (avoid in tests)

  **Test References**:
  - None exist yet - this is first test file

  **Documentation References**:
  - Go testing package: https://go.dev/pkg/testing

  **External References**:
  - Table-driven tests: https://dave.cheney.net/2019/05/07/prefer-table-driven-tests

  **WHY Each Reference Matters**:
  - Interface defines contract to test against
  - Table-driven tests reduce code duplication
  - Tests cover both success and error paths

  **Acceptance Criteria**:

  **If TDD (tests enabled):**
  - [ ] Test file created: workspace.repository_impl_test.go
  - [ ] All repository methods have tests
  - [ ] Tests cover success paths (TestSave_Success, etc.)
  - [ ] Tests cover error paths (TestFindByID_NotFound, etc.)
  - [ ] go test ./internal/infrastructure/persistence/bun/repositories/ -run TestSave_Success → FAIL (RED)

  **Agent-Executed QA Scenarios (MANDATORY — per-scenario, ultra-detailed):**

  Scenario: Repository tests fail before implementation (RED)
    Tool: Bash (go test)
    Preconditions: workspace.repository.go interface exists, repository implementation exists
    Steps:
      1. go test ./internal/infrastructure/persistence/bun/repositories/ -v -run "TestSave|TestFindByID|TestFindAll|TestUpdateName|TestDelete"
      2. assert exit code is 1 (tests fail)
    Expected Result: Tests fail (RED)
    Failure Indicators: Tests pass unexpectedly
    Evidence: Test output

  **Evidence to Capture**:
  - [ ] Test output showing failures

  **Commit**: NO (group with Task 14)

- [x] 7. Implement repository implementation (GREEN)

  **What to do**:
  - Create `internal/infrastructure/persistence/bun/repositories/workspace.repository_impl.go`
  - Implement `WorkspaceRepository` interface:
    - Save: Insert new workspace with Bun model
    - FindByID: Query by ID, map model to domain entity
    - FindAll: Query all, map models to domain entities
    - UpdateName: Update workspace name in database
    - Delete: Delete workspace by ID
  - Map between Bun model and domain entity
  - Handle not found errors (return ErrWorkspaceNotFound)
  - Use context for all database operations

  **Must NOT do**:
  - Use domain entity directly in database operations
  - Skip error handling
  - Create complex queries

  **Implementation Note - Mapping Bun Model to Domain Entity**:
  ```go
  // For FindByID and FindAll:
  domainWorkspace := &workspace.Workspace{
      id:        workspace.WorkspaceID(model.ID),
      name:      model.Name,
      isDefault: model.IsDefault,
      createdAt: model.CreatedAt,
  }

  // For Save:
  model := &models.Workspace{
      ID:        string(entity.ID()),
      Name:      entity.Name(),
      IsDefault: entity.IsDefault(),
      CreatedAt: entity.CreatedAt(),
  }
  ```
  Use domain entity's getters (ID(), Name(), IsDefault(), CreatedAt()) and Bun model's public fields directly.

  **Recommended Agent Profile**:
  - **Category**: `unspecified-low`
    - Reason: Go implementation, Bun ORM integration
  - **Skills**: None needed

  **Parallelization**:
  - **Can Run In Parallel**: NO
  - **Parallel Group**: Sequential
  - **Blocks**: Task 8 (refactoring)
  - **Blocked By**: Task 6 (tests written)

  **References**:

  **Pattern References**:
  - `internal/infrastructure/persistence/bun/models/workspace.go:9-17` - Bun Workspace model
  - `internal/infrastructure/persistance/bun/respositories/workspace_repo.go:21-47` - Existing implementation pattern

  **API/Type References**:
  - `internal/application/workspace/workspace.repository.go` - Interface to implement
  - `internal/infrastructure/persistence/bun/models/workspace.Workspace` - Infrastructure model
  - `internal/domain/workspace/workspace.Workspace` - Domain entity to map to
  - `internal/application/workspace/workspace.errors.go:ErrWorkspaceNotFound` - Not found error

  **Test References**:
  - `internal/infrastructure/persistence/bun/repositories/workspace.repository_impl_test.go` - Tests to pass

  **Documentation References**:
  - Bun ORM docs: https://bun.uptrace.dev/guide
  - Research findings: Repository implementation pattern

  **External References**:
  - Bun Insert: https://bun.uptrace.dev/guide/insert
  - Bun Select: https://bun.uptrace.dev/guide/select

  **WHY Each Reference Matters**:
  - Interface defines contract to implement
  - Bun model is infrastructure-specific, must map to domain
  - Existing implementation shows patterns to follow

  **Acceptance Criteria**:

  **If TDD (tests enabled):**
  - [ ] File created/renamed: workspace.repository_impl.go
  - [ ] WorkspaceRepository interface implemented
  - [ ] All methods implemented
  - [ ] go test ./internal/infrastructure/persistence/bun/repositories/... -v → PASS (GREEN)

  **Agent-Executed QA Scenarios (MANDATORY — per-scenario, ultra-detailed):**

  Scenario: Repository tests pass after implementation (GREEN)
    Tool: Bash (go test)
    Preconditions: workspace.repository_impl.go exists
    Steps:
      1. go test ./internal/infrastructure/persistence/bun/repositories/... -v -race
      2. assert exit code is 0 (tests pass)
    Expected Result: All tests pass (GREEN)
    Failure Indicators: Test failures, race conditions
    Evidence: Test output

  Scenario: Repository compiles successfully
    Tool: Bash (go build)
    Preconditions: workspace.repository_impl.go exists
    Steps:
      1. go build ./internal/infrastructure/persistence/bun/repositories/...
      2. assert exit code is 0
    Expected Result: Successful compilation
    Failure Indicators: Compilation errors
    Evidence: Build output

  **Evidence to Capture**:
  - [ ] Test output showing all tests pass
  - [ ] Build output

  **Commit**: NO (group with Task 14)

- [x] 8. Refactor repository implementation

  **What to do**:
  - Clean up repository implementation code
  - Extract common patterns if needed
  - Add comments if unclear
  - Ensure code follows Go conventions
  - All tests must still pass

  **Must NOT do**:
  - Change public API
  - Modify interface
  - Break any tests

  **Recommended Agent Profile**:
  - **Category**: `unspecified-low`
    - Reason: Go refactoring, code cleanup
  - **Skills**: None needed

  **Parallelization**:
  - **Can Run In Parallel**: NO
  - **Parallel Group**: Sequential
  - **Blocks**: Task 12, 13 (file renames)
  - **Blocked By**: Task 7 (implementation complete)

  **References**:

  **Pattern References**:
  - `internal/infrastructure/persistence/bun/repositories/workspace.repository_impl.go` - Current implementation

  **Test References**:
  - `internal/infrastructure/persistence/bun/repositories/workspace.repository_impl_test.go` - Tests to verify refactoring

  **Documentation References**:
  - Go code review comments: https://go.dev/doc/effective_go#commentary

  **WHY Each Reference Matters**:
  - Tests ensure refactoring doesn't break behavior
  - Code conventions ensure maintainability

  **Acceptance Criteria**:

  **If TDD (tests enabled):**
  - [ ] go test ./internal/infrastructure/persistence/bun/repositories/... -v → PASS (still GREEN)

  **Agent-Executed QA Scenarios (MANDATORY — per-scenario, ultra-detailed):**

  Scenario: Repository refactoring preserves behavior
    Tool: Bash (go test)
    Preconditions: All tests passing
    Steps:
      1. go test ./internal/infrastructure/persistence/bun/repositories/... -v -race
      2. assert exit code is 0
    Expected Result: All tests still pass after refactoring
    Failure Indicators: Test failures, new bugs introduced
    Evidence: Test output

  **Evidence to Capture**:
  - [ ] Test output showing all tests pass

  **Commit**: NO (group with Task 14)

- [x] 9. Create service interface and implementation tests (RED)

  **What to do**:
  - Create `internal/application/workspace/workspace.service.go`
  - Define `WorkspaceService` interface with methods:
    - CreateWorkspace(ctx, CreateWorkspaceRequest) (*WorkspaceResponse, error)
    - GetWorkspaceByID(ctx, WorkspaceID) (*WorkspaceResponse, error)
    - ListWorkspaces(ctx) ([]*WorkspaceResponse, error)
    - UpdateWorkspaceName(ctx, WorkspaceID, string) (*WorkspaceResponse, error)
    - DeleteWorkspace(ctx, WorkspaceID) error
  - Create service struct with repository dependency
  - Write tests for all service methods:
    - TestCreateWorkspace_Success
    - TestCreateWorkspace_EmptyName
    - TestGetWorkspaceByID_Success
    - TestGetWorkspaceByID_NotFound
    - TestListWorkspaces_Success
    - TestUpdateWorkspaceName_Success
    - TestUpdateWorkspaceName_NotFound
    - TestDeleteWorkspace_Success
    - TestDeleteWorkspace_NotFound

  **Must NOT do**:
  - Implement service methods (this is test-only task)
  - Create HTTP handlers
  - Add business rules beyond validation

  **Recommended Agent Profile**:
  - **Category**: `unspecified-low`
    - Reason: Go test writing, service interface definition
  - **Skills**: None needed

  **Parallelization**:
  - **Can Run In Parallel**: NO
  - **Parallel Group**: Sequential
  - **Blocks**: Task 10 (implementation needed to make tests pass)
  - **Blocked By**: Task 5 (repository interface), Task 4 (DTOs defined)

  **References**:

  **Pattern References**:
  - Research findings: Service layer pattern in DDD

  **API/Type References**:
  - `internal/application/workspace/workspace.repository.go` - Repository to inject
  - `internal/application/workspace/workspace.dto.go:CreateWorkspaceRequest` - Input type
  - `internal/application/workspace/workspace.dto.go:WorkspaceResponse` - Output type
  - `internal/domain/workspace/workspace.Workspace` - Domain entity
  - `internal/application/workspace/workspace.errors.go:ErrInvalidName` - Validation error

  **Test References**:
  - None exist yet for service

  **Documentation References**:
  - Go testing package: https://go.dev/pkg/testing

  **External References**:
  - Table-driven tests: https://dave.cheney.net/2019/05/07/prefer-table-driven-tests

  **WHY Each Reference Matters**:
  - Service orchestrates repository calls + validation
  - DTOs define input/output contracts
  - Tests cover all service operations

  **Acceptance Criteria**:

  **If TDD (tests enabled):**
  - [ ] File created: workspace.service.go
  - [ ] WorkspaceService interface defined
  - [ ] All service methods have tests
  - [ ] go test ./internal/application/workspace/ -run TestCreateWorkspace_Success → FAIL (RED)

  **Agent-Executed QA Scenarios (MANDATORY — per-scenario, ultra-detailed):**

  Scenario: Service tests fail before implementation (RED)
    Tool: Bash (go test)
    Preconditions: workspace.service.go interface exists
    Steps:
      1. go test ./internal/application/workspace/ -v -run "TestCreateWorkspace|TestGetWorkspaceByID|TestListWorkspaces|TestUpdateWorkspaceName|TestDeleteWorkspace"
      2. assert exit code is 1 (tests fail)
    Expected Result: Tests fail (RED)
    Failure Indicators: Tests pass unexpectedly
    Evidence: Test output

  **Evidence to Capture**:
  - [ ] Test output showing failures

  **Commit**: NO (group with Task 14)

- [x] 10. Implement service (GREEN)

  **What to do**:
  - Implement `WorkspaceService` interface methods in workspace.service.go:
    - CreateWorkspace: Validate name, call repository.Save, map entity to response
    - GetWorkspaceByID: Call repository.FindByID, map entity to response
    - ListWorkspaces: Call repository.FindAll, map entities to responses
    - UpdateWorkspaceName: Call repository.UpdateName, map entity to response
    - DeleteWorkspace: Call repository.Delete
  - Add validation: name cannot be empty (return ErrInvalidName)
  - Handle not found errors from repository
  - Map between DTO, domain entity, and response
  - Use context for all operations

  **Must NOT do**:
  - Skip validation
  - Ignore repository errors
  - Add complex business rules beyond validation
  - Create domain events

  **Recommended Agent Profile**:
  - **Category**: `unspecified-low`
    - Reason: Go implementation, service layer pattern
  - **Skills**: None needed

  **Parallelization**:
  - **Can Run In Parallel**: NO
  - **Parallel Group**: Sequential
  - **Blocks**: Task 11 (refactoring)
  - **Blocked By**: Task 9 (tests written), Task 4 (DTOs defined), Task 7 (repository implementation)

  **References**:

  **Pattern References**:
  - Research findings: Service layer pattern with validation

  **API/Type References**:
  - `internal/application/workspace/workspace.repository.go` - Repository to use
  - `internal/application/workspace/workspace.dto.go:CreateWorkspaceRequest` - Input type
  - `internal/application/workspace/workspace.dto.go:WorkspaceResponse` - Output type
  - `internal/domain/workspace/workspace.NewWorkspace` - Domain entity factory
  - `internal/application/workspace/workspace.errors.go:ErrInvalidName` - Validation error
  - `internal/application/workspace/workspace.errors.go:ErrWorkspaceNotFound` - Not found error

  **Test References**:
  - `internal/application/workspace/workspace.service_test.go` - Tests to pass

  **Documentation References**:
  - Go context: https://go.dev/pkg/context

  **External References**:
  - Go error wrapping: https://go.dev/blog/error-handling-and-go

  **WHY Each Reference Matters**:
  - Repository is data access layer
  - DTOs define service API
  - Validation is service responsibility
  - Domain entity holds business rules

  **Acceptance Criteria**:

  **If TDD (tests enabled):**
  - [ ] All service methods implemented
  - [ ] Validation for empty name added
  - [ ] Error handling for not found added
  - [ ] go test ./internal/application/workspace/ -v → PASS (GREEN)

  **Agent-Executed QA Scenarios (MANDATORY — per-scenario, ultra-detailed):**

  Scenario: Service tests pass after implementation (GREEN)
    Tool: Bash (go test)
    Preconditions: workspace.service.go implementation exists
    Steps:
      1. go test ./internal/application/workspace/... -v -race
      2. assert exit code is 0 (tests pass)
    Expected Result: All tests pass (GREEN)
    Failure Indicators: Test failures, race conditions
    Evidence: Test output

  Scenario: Service compiles successfully
    Tool: Bash (go build)
    Preconditions: workspace.service.go exists
    Steps:
      1. go build ./internal/application/workspace/...
      2. assert exit code is 0
    Expected Result: Successful compilation
    Failure Indicators: Compilation errors
    Evidence: Build output

  **Evidence to Capture**:
  - [ ] Test output showing all tests pass
  - [ ] Build output

  **Commit**: NO (group with Task 14)

- [x] 11. Refactor service implementation

  **What to do**:
  - Clean up service implementation code
  - Extract DTO mapping functions if needed
  - Add comments if unclear
  - Ensure code follows Go conventions
  - All tests must still pass

  **Must NOT do**:
  - Change public API
  - Modify interface
  - Break any tests

  **Recommended Agent Profile**:
  - **Category**: `unspecified-low`
    - Reason: Go refactoring, code cleanup
  - **Skills**: None needed

  **Parallelization**:
  - **Can Run In Parallel**: NO
  - **Parallel Group**: Sequential
  - **Blocks**: Task 12, 13 (file renames)
  - **Blocked By**: Task 10 (implementation complete)

  **References**:

  **Pattern References**:
  - `internal/application/workspace/workspace.service.go` - Current implementation

  **Test References**:
  - `internal/application/workspace/workspace.service_test.go` - Tests to verify refactoring

  **Documentation References**:
  - Go code review comments: https://go.dev/doc/effective_go#commentary

  **WHY Each Reference Matters**:
  - Tests ensure refactoring doesn't break behavior
  - Code conventions ensure maintainability

  **Acceptance Criteria**:

  **If TDD (tests enabled):**
  - [ ] go test ./internal/application/workspace/... -v → PASS (still GREEN)

  **Agent-Executed QA Scenarios (MANDATORY — per-scenario, ultra-detailed):**

  Scenario: Service refactoring preserves behavior
    Tool: Bash (go test)
    Preconditions: All tests passing
    Steps:
      1. go test ./internal/application/workspace/... -v -race
      2. assert exit code is 0
    Expected Result: All tests still pass after refactoring
    Failure Indicators: Test failures, new bugs introduced
    Evidence: Test output

  **Evidence to Capture**:
  - [ ] Test output showing all tests pass

  **Commit**: NO (group with Task 14)

- [x] 12. Rename entity file in infrastructure

  **What to do**:
  - Rename `internal/infrastructure/persistence/bun/models/workspace.go` to `internal/infrastructure/persistence/bun/models/workspace.model.go`
  - Update package declaration if needed
  - Ensure no imports break from rename

  **Must NOT do**:
  - Modify entity structure or fields
  - Change database schema

  **Recommended Agent Profile**:
  - **Category**: `unspecified-low`
    - Reason: Simple file rename, Go code structure
  - **Skills**: None needed

  **Parallelization**:
  - **Can Run In Parallel**: YES
  - **Parallel Group**: Wave 4 (with Task 13)
  - **Blocks**: Task 14 (FX module update)
  - **Blocked By**: Task 2 (path typos fixed), Task 8 (repository refactored)

  **References**:

  **Pattern References**:
  - `internal/infrastructure/persistence/bun/models/workspace.go` - Current Bun model

  **API/Type References**:
  - `github.com/uptrace/bun` - Bun BaseModel

  **Test References**:
  - `internal/infrastructure/persistence/bun/repositories/workspace.repository_impl_test.go` - Tests that import model

  **Documentation References**:
  - Research findings: {domain}.{type}.go naming for infrastructure models

  **External References**:
  - Go file renaming: `go build ./...` to verify no import breakage

  **WHY Each Reference Matters**:
  - Renaming follows recommended convention
  - Imports must update to prevent breakage

  **Acceptance Criteria**:

  **If TDD (tests enabled):**
  - [ ] File renamed: internal/infrastructure/persistence/bun/models/workspace.model.go exists
  - [ ] Old file deleted: workspace.go does not exist
  - [ ] go build ./internal/infrastructure/persistence/bun/models/... → PASS
  - [ ] go test ./internal/infrastructure/persistence/bun/repositories/... → PASS (tests still work)

  **Agent-Executed QA Scenarios (MANDATORY — per-scenario, ultra-detailed):**

  Scenario: Infrastructure model renamed correctly
    Tool: Bash (ls/go build)
    Preconditions: workspace.go model exists
    Steps:
      1. mv internal/infrastructure/persistence/bun/models/workspace.go internal/infrastructure/persistence/bun/models/workspace.model.go
      2. ls internal/infrastructure/persistence/bun/models/workspace.model.go
      3. ! ls internal/infrastructure/persistence/bun/models/workspace.go 2>/dev/null
      4. go build ./internal/infrastructure/persistence/bun/models/...
    Expected Result: workspace.model.go exists, old file doesn't, compiles
    Failure Indicators: Old file exists, new file missing, compilation errors
    Evidence: File listing and build output

  **Evidence to Capture**:
  - [ ] Directory listing showing workspace.model.go
  - [ ] Build output

  **Commit**: NO (group with Task 14)

- [x] 13. Rename repository implementation file

  **What to do**:
  - Rename `workspace_repo.go` to `workspace.repository_impl.go`
  - Update package if needed
  - Ensure no imports break from rename

  **Must NOT do**:
  - Modify repository implementation logic
  - Change interface

  **Recommended Agent Profile**:
  - **Category**: `unspecified-low`
    - Reason: Simple file rename, Go code structure
  - **Skills**: None needed

  **Parallelization**:
  - **Can Run In Parallel**: YES
  - **Parallel Group**: Wave 4 (with Task 12)
  - **Blocks**: Task 14 (FX module update)
  - **Blocked By**: Task 2 (path typos fixed), Task 8 (repository refactored)

  **References**:

  **Pattern References**:
  - `internal/infrastructure/persistence/bun/repositories/workspace_repo.go` - Current file name

  **API/Type References**:
  - `internal/application/workspace/workspace.repository.go` - Interface being implemented

  **Test References**:
  - `internal/infrastructure/persistence/bun/repositories/workspace.repository_impl_test.go` - Test file (already renamed)

  **Documentation References**:
  - Research findings: {domain}.{type}.go naming for repository implementations

  **WHY Each Reference Matters**:
  - Renaming follows recommended convention
  - Test file already uses new name (created in Task 6)

  **Acceptance Criteria**:

  **If TDD (tests enabled):**
  - [ ] File renamed: internal/infrastructure/persistence/bun/repositories/workspace.repository_impl.go exists
  - [ ] Old file deleted: workspace_repo.go does not exist
  - [ ] go build ./internal/infrastructure/persistence/bun/repositories/... → PASS
  - [ ] go test ./internal/infrastructure/persistence/bun/repositories/... → PASS (tests still work)

  **Agent-Executed QA Scenarios (MANDATORY — per-scenario, ultra-detailed):**

  Scenario: Repository implementation renamed correctly
    Tool: Bash (ls/go build)
    Preconditions: workspace_repo.go exists
    Steps:
      1. mv internal/infrastructure/persistence/bun/repositories/workspace_repo.go internal/infrastructure/persistence/bun/repositories/workspace.repository_impl.go
      2. ls internal/infrastructure/persistence/bun/repositories/workspace.repository_impl.go
      3. ! ls internal/infrastructure/persistence/bun/repositories/workspace_repo.go 2>/dev/null
      4. go build ./internal/infrastructure/persistence/bun/repositories/...
    Expected Result: workspace.repository_impl.go exists, old file doesn't, compiles
    Failure Indicators: Old file exists, new file missing, compilation errors
    Evidence: File listing and build output

  **Evidence to Capture**:
  - [ ] Directory listing showing workspace.repository_impl.go
  - [ ] Build output

  **Commit**: NO (group with Task 14)

- [x] 14. Update FX module

  **What to do**:
  - Update `cmd/server/di/modules/workspace.go`:
    - Remove incorrect import `internal/application/workspace` (doesn't exist)
    - Remove `fx.Provide(workspace.NewWorkspaceRepository)` (interface, not implemented)
    - Remove `fx.Provide(workspace.NewWorkspaceService)` (doesn't exist)
    - Remove `fx.Provide(workspace.NewWorkspaceHandler)` (doesn't exist, out of scope)
    - Add correct import `internal/application/workspace`
    - Add `fx.Provide(workspace.NewWorkspaceRepository)` (interface)
    - Add `fx.Provide(repositories.NewWorkspaceRepositoryImpl)` (implementation) - needs constructor name
    - Add `fx.Provide(workspace.NewWorkspaceService)` (service) - needs constructor name
  - Ensure module compiles

  **Must NOT do**:
  - Add HTTP handler (out of scope)
  - Modify other modules
  - Create fx.Invoke patterns (not needed)

  **Recommended Agent Profile**:
  - **Category**: `unspecified-low`
    - Reason: Fx module configuration, Go dependency injection
  - **Skills**: None needed

  **Parallelization**:
  - **Can Run In Parallel**: NO
  - **Parallel Group**: Sequential
  - **Blocks**: Task 15 (final verification)
  - **Blocked By**: Task 12, 13 (files renamed)

  **References**:

  **Pattern References**:
  - `cmd/server/di/modules/workspace.go` - Current FX module
  - Research findings: Fx module pattern with fx.Provide

  **API/Type References**:
  - `internal/application/workspace/workspace.repository.go:WorkspaceRepository` - Interface
  - `internal/infrastructure/persistence/bun/repositories/workspace.repository_impl.go` - Implementation
  - `internal/application/workspace/workspace.service.go:WorkspaceService` - Service

  **Documentation References**:
  - Fx documentation: https://uber-go.github.io/fx/modules.html

  **External References**:
  - Fx provide: https://uber-go.github.io/fx/annotated-provide

  **WHY Each Reference Matters**:
  - FX module wires all dependencies together
  - Must reference correct constructors
  - Must have interface and implementation for repository

  **Acceptance Criteria**:

  **If TDD (tests enabled):**
  - [ ] fx.Provide calls added for repository interface and implementation
  - [ ] fx.Provide call added for service
  - [ ] No references to removed providers (workspace.NewWorkspaceRepository - wrong package)
  - [ ] go build ./cmd/server/... → PASS

  **Agent-Executed QA Scenarios (MANDATORY — per-scenario, ultra-detailed):**

  Scenario: FX module updated with correct providers
    Tool: Bash (go build/cat)
    Preconditions: All application layer files exist
    Steps:
      1. cat cmd/server/di/modules/workspace.go
      2. grep -q "fx.Provide.*workspace.NewWorkspaceRepository" cmd/server/di/modules/workspace.go
      3. grep -q "fx.Provide.*repositories.NewWorkspaceRepositoryImpl" cmd/server/di/modules/workspace.go
      4. grep -q "fx.Provide.*workspace.NewWorkspaceService" cmd/server/di/modules/workspace.go
      5. ! grep -q "NewWorkspaceHandler" cmd/server/di/modules/workspace.go
      6. go build ./cmd/server/di/...
    Expected Result: Correct providers present, handler removed, compiles
    Failure Indicators: Missing providers, handler still referenced, compilation errors
    Evidence: File content and build output

  **Evidence to Capture**:
  - [ ] workspace.go module content
  - [ ] Build output

  **Commit**: YES (groups with Tasks 1-4, 12-13)
  - Message: `refactor(workspace): fix DDD structure, add application layer, rename files`
  - Files: All modified files
  - Pre-commit: `go test ./internal/workspace/... && go test ./internal/infrastructure/persistence/bun/repositories/...`

- [x] 15. Final verification and cleanup

  **What to do**:
  - Run all tests to verify everything works
  - Check test coverage
  - Verify no path typos remain
  - Verify all code compiles
  - Delete any temporary files
  - Summarize changes

  **Must NOT do**:
  - Modify any implementation code
  - Add new features

  **Recommended Agent Profile**:
  - **Category**: `unspecified-low`
    - Reason: Verification, cleanup
  - **Skills**: None needed

  **Parallelization**:
  - **Can Run In Parallel**: NO
  - **Parallel Group**: Sequential
  - **Blocks**: None (final task)
  - **Blocked By**: All previous tasks (Task 14 complete)

  **References**:

  **Pattern References**:
  - All previous tasks

  **Test References**:
  - All test files created

  **Documentation References**:
  - Go test coverage: `go test -cover`

  **WHY Each Reference Matters**:
  - Final verification ensures all changes work together
  - Cleanup ensures clean state

  **Acceptance Criteria**:

  **If TDD (tests enabled):**
  - [ ] go test ./internal/workspace/... -v -race → PASS
  - [ ] go test ./internal/infrastructure/persistence/bun/repositories/... -v -race → PASS
  - [ ] go build ./internal/workspace/... → PASS
  - [ ] go build ./internal/infrastructure/persistence/... → PASS
  - [ ] go build ./cmd/server/... → PASS
  - [ ] go test ./internal/workspace/... -cover → Coverage ≥ 70%
  - [ ] ! find internal/workspace -type d -name "*persistance*" 2>/dev/null → No results
  - [ ] ! find internal/workspace -type d -name "*respositories*" 2>/dev/null → No results

  **Agent-Executed QA Scenarios (MANDATORY — per-scenario, ultra-detailed):**

  Scenario: All tests pass and coverage meets threshold
    Tool: Bash (go test)
    Preconditions: All implementation complete
    Steps:
      1. go test ./internal/workspace/... -v -race -cover
      2. go test ./internal/infrastructure/persistence/bun/repositories/... -v -race -cover
      3. assert exit code is 0 for both
      4. grep "coverage:" output to verify ≥ 70%
    Expected Result: All tests pass, coverage ≥ 70%, no race conditions
    Failure Indicators: Test failures, low coverage, race conditions
    Evidence: Test output with coverage

  Scenario: All code compiles successfully
    Tool: Bash (go build)
    Preconditions: All implementation complete
    Steps:
      1. go build ./internal/workspace/...
      2. go build ./internal/infrastructure/persistence/...
      3. go build ./cmd/server/...
      4. assert exit code is 0 for all
    Expected Result: Successful compilation across all packages
    Failure Indicators: Compilation errors
    Evidence: Build output

  Scenario: No path typos remain
    Tool: Bash (find)
    Preconditions: All paths fixed
    Steps:
      1. find internal/workspace -type d -name "*persistance*"
      2. find internal/workspace -type d -name "*respositories*"
      3. assert both return no results
    Expected Result: No typos found in directory structure
    Failure Indicators: Old typo directories still exist
    Evidence: find output

  **Evidence to Capture**:
  - [ ] Test output with coverage
  - [ ] Build output
  - [ ] Find output showing no typos

  **Commit**: NO (already committed in Task 14)

---

## Commit Strategy

| After Task | Message | Files | Verification |
|------------|---------|-------|--------------|
| 14 | `refactor(workspace): fix DDD structure, add application layer, rename files` | internal/domain/workspace/workspace.entity.go, internal/application/workspace/*.go, internal/infrastructure/persistence/bun/**/*.go, cmd/server/di/modules/workspace.go | `go test ./internal/workspace/... && go test ./internal/infrastructure/persistence/bun/repositories/...` |

---

## Success Criteria

### Verification Commands
```bash
# All tests pass
go test ./internal/workspace/... -v -race -cover
go test ./internal/infrastructure/persistence/bun/repositories/... -v -race -cover

# All code compiles
go build ./internal/workspace/...
go build ./internal/infrastructure/persistence/...
go build ./cmd/server/...

# No path typos
! find internal/workspace -type d -name "*persistance*"
! find internal/workspace -type d -name "*respositories*"

# Test coverage threshold
go test ./internal/workspace/... -cover | grep -E "coverage: [7-9][0-9]%"
```

### Final Checklist
- [x] All "Must Have" present (repository interface, service, DTOs, errors, tests, renamed files, fixed typos, updated FX module)
- [ ] All "Must NOT Have" absent (HTTP handlers, migrations, integration tests, other domains, events)
- [x] All tests pass (entity, repository, service)
- [x] Test coverage ≥ 70%
- [x] All code compiles
- [x] No path typos remain
- [ ] Existing workspace entity behavior preserved
