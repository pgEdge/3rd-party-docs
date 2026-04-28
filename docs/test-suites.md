# Test Suites

## Smoke Test (`smoke_test.sql`)

Quick sanity check that verifies:
- Extension loads successfully
- All 10 functions exist and return data
- Basic functionality works

**Run time:** ~1-2 seconds

```bash
make installcheck REGRESS=smoke_test
```

## Comprehensive Test Suite (`system_stats.sql`)

Full test coverage including:
- Data validation
- Type checking
- Value range verification
- Consistency checks
- Caching behavior
- Cross-platform compatibility

**Run time:** ~5-10 seconds

```bash
make installcheck REGRESS=system_stats
```

## Full Test Suite

Runs all tests in sequence:

```bash
make installcheck
```

