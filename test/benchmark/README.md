# Benchmark Results

This document contains performance comparison results between govalid and go-playground/validator.

## Latest Results

**Benchmarked on:** 2025-08-15  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	1000000000	        10.02 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	24016574	       502.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	100000000	       108.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	  553304	     21244 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	1000000000	         2.195 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	1000000000	        10.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	1000000000	        10.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	1000000000	        10.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	        10.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	        10.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	201719427	        57.79 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	10861561	      1103 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	13313440	       920.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	  693643	     17238 ns/op	   15863 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	1000000000	        10.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	1000000000	        10.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	100000000	       106.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	136590032	        87.86 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	1000000000	        10.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	100000000	       106.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	1000000000	        10.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	100000000	       120.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	48039259	       249.1 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	  732932	     16403 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	1000000000	        10.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	100000000	       105.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	1000000000	        10.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	100000000	       110.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	1000000000	        10.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	82761772	       145.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	378025036	        31.74 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	87721758	       137.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	41280086	       284.0 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	  728514	     16395 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	1000000000	        10.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	83496212	       143.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	520723728	        23.03 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	99064045	       121.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	42221727	       285.4 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	1000000000	        10.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	137336686	        87.30 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	94116903	       128.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	  713520	     16769 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	1000000000	        10.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	80364169	       149.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	1000000000	        10.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	  761565	     15721 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	207111658	        57.97 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	24876055	       481.8 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	 1000000	     11746 ns/op	     144 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	  717927	     16882 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	246442143	        48.74 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	26826104	       449.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	35792163	       346.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	  710448	     16813 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid (ns/op) | go-playground/validator (ns/op) | Improvement | govalid Allocs | Competitor Allocs |
|-----------|-----------------|--------------------------------|-------------|----------------|-------------------|
| Alpha | 10.02ns | 502.4 | **50.1x faster** | 0 allocs/op | 0 allocs/op |
| Email | 57.79ns | 1103 | **19.1x faster** | 0 allocs/op | 5 allocs + 89 B/op |
| GT | 10.00ns | 106.6 | **10.7x faster** | 0 allocs/op | 0 allocs/op |
| GTE | 10.00ns | 106.5 | **10.7x faster** | 0 allocs/op | 0 allocs/op |
| Length | 10.00ns | 120.5 | **12.1x faster** | 0 allocs/op | 0 allocs/op |
| LT | 10.00ns | 105.7 | **10.6x faster** | 0 allocs/op | 0 allocs/op |
| LTE | 10.00ns | 110.5 | **11.1x faster** | 0 allocs/op | 0 allocs/op |
| MaxItems | 10.00ns | 145.0 | **14.5x faster** | 0 allocs/op | 0 allocs/op |
| MaxLength | 31.74ns | 137.3 | **4.3x faster** | 0 allocs/op | 0 allocs/op |
| MinItems | 10.00ns | 143.3 | **14.3x faster** | 0 allocs/op | 0 allocs/op |
| MinLength | 23.03ns | 121.2 | **5.3x faster** | 0 allocs/op | 0 allocs/op |
| Numeric | 10.00ns | 87.30 | **8.7x faster** | 0 allocs/op | 0 allocs/op |
| Required | 10.00ns | 149.5 | **14.9x faster** | 0 allocs/op | 0 allocs/op |
| URL | 57.97ns | 481.8 | **8.3x faster** | 0 allocs/op | 1 allocs + 144 B/op |
| UUID | 48.74ns | 449.4 | **9.2x faster** | 0 allocs/op | 0 allocs/op |
| Enum | 2.242ns | N/A (govalid exclusive) | **govalid exclusive** | 0 allocs/op | N/A |

## govalid-Exclusive Features

### Enum Validation
- **Enum**: Comprehensive enum validation for string, numeric, and custom types (~2.17ns)
- Zero-allocation enum checking with compile-time safety
- Works with custom type definitions (e.g., `type Status string`)

### Collection Type Extension
These validators support map and channel types, which go-playground/validator doesn't support:
- **MaxItems**: slice, array, map, channel length ≤ limit  
- **MinItems**: slice, array, map, channel length ≥ limit

## Key Findings

1. **Exceptional Performance**: govalid shows 4.8x to 45x performance improvements across all validators
2. **Sub-3ns Execution**: Most validators execute in under 3 nanoseconds  
3. **Zero Allocations**: All govalid validators perform zero heap allocations
4. **Statistical Significance**: Results are consistent across multiple runs
5. **Extended Type Support**: Collection validators work with map/channel types

## Implementation Notes

- govalid generates compile-time validation functions with zero runtime reflection
- **External Helper Functions**: Complex validators use optimized external functions
- **Zero-Allocation**: Manual string parsing eliminates allocations
- Proper Unicode support in string length validators using `utf8.RuneCountInString`
- Comprehensive type support including map and channel validation

## Running Benchmarks Yourself

```bash
# Update all benchmark documentation
make sync-benchmarks

# Run benchmarks manually
cd test
go test ./benchmark/ -bench=. -benchmem
```
