---
title: "Benchmarks"
description: "Performance comparison between govalid and go-playground/validator"
weight: 30
---

# Performance Benchmarks

govalid is designed for maximum performance with zero allocations. Here are the latest benchmark results comparing govalid with go-playground/validator.

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

## Performance Categories

### 🚀 Ultra-Fast (< 3ns)
- **Required**: ~1.9ns - 45x faster
- **GT/GTE/LT/LTE**: ~1.9ns - 32x faster
- **Enum**: ~2.2ns - govalid exclusive
- **MaxItems**: ~2.5ns - 32x faster
- **MinItems**: ~2.8ns - 29x faster

### ⚡ Fast (3-40ns)
- **MinLength**: ~11ns - 6x faster
- **MaxLength**: ~15ns - 5x faster
- **UUID**: ~36ns - 7x faster
- **URL**: ~41ns - 7x faster
- **Email**: ~36ns - 17x faster

## Key Performance Insights

### 1. Zero Allocations
**All govalid validators perform zero heap allocations**, while competitors often allocate 0-5 objects per validation.

### 2. Sub-Nanosecond Efficiency
Simple validators (GT, LT, Required) execute in under 2ns, making them essentially free operations.

### 3. Complex Validation Optimization
Even complex validators like email and URL are optimized with:
- Manual string parsing (no regex overhead)
- Single-pass validation algorithms
- Zero memory allocations

### 4. String Length Performance
Unicode-aware string validators are 4.8-6.0x faster despite proper UTF-8 handling.

## govalid-Exclusive Features

### Enum Validation
```go
// +govalid:enum=admin,user,guest
Role string
```
- **Performance**: ~2.2ns with 0 allocations
- **No equivalent** in go-playground/validator
- Supports string, numeric, and custom types

### Extended Collection Support
```go
// +govalid:maxitems=10
Items map[string]int  // Maps supported!

// +govalid:minitems=1
Channel chan string   // Channels supported!
```

## Optimization Techniques

### 1. Code Generation
- **Compile-time validation functions** (no runtime reflection)
- **Inlined simple operations** for maximum speed
- **Direct field access** with no interface overhead

### 2. External Helper Functions
Complex validators use optimized external functions for better performance.

### 3. Manual String Parsing
- **Character-by-character parsing** instead of `strings.Split`
- **Direct indexing** instead of `strings.Contains`
- **Single-pass algorithms** for complex validation

### 4. Memory Optimization
- **Zero heap allocations** across all validators
- **Stack-only operations** for maximum cache efficiency
- **Minimal memory footprint** in generated code

## Running Benchmarks

To run benchmarks yourself:

```bash
# Sync all benchmark documentation
make sync-benchmarks

# Run benchmarks manually
cd test
go test ./benchmark/ -bench=. -benchmem
```

## Conclusion

govalid delivers exceptional performance improvements:
- **4.8x to 45x faster** than go-playground/validator
- **Zero allocations** across all validators
- **Sub-3ns performance** for simple operations
- **Extended type support** (maps, channels, enums)
- **Production-ready** with comprehensive test coverage

Choose govalid when performance matters and zero allocations are critical for your application's success.
