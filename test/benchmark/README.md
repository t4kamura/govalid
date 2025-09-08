# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-09-08  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	124473628	         9.670 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2606103	       465.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11178291	       108.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   55693	     21239 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	638850813	         1.863 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	295958680	         4.052 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	289445475	         4.104 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	350296450	         3.426 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21703123	        59.26 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1092 ns/op	      88 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1263481	       956.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   69216	     17147 ns/op	   15847 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	267632818	         4.405 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	478963000	         2.496 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11287620	       106.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13633368	        90.16 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	481663126	         2.495 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11351583	       106.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	156239000	         7.602 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	 9640663	       126.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4843186	       248.2 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   74084	     16200 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	482185053	         2.489 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11330860	       105.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	427997881	         2.811 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	10876232	       110.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	183818355	         6.534 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8348068	       144.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	35942862	        33.45 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8616679	       139.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4203134	       284.8 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   73822	     16244 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	202414388	         5.922 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8275940	       145.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	52060402	        23.03 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10078431	       120.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4267040	       282.0 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	157632157	         7.600 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	12190053	        98.42 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9473314	       128.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   71541	     16777 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	321553741	         3.734 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8020893	       149.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	642661272	         1.866 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   77361	     15585 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20722892	        57.69 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2460825	       484.2 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  102506	     11748 ns/op	     146 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   71732	     16903 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	24600267	        48.73 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2674954	       448.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3431590	       353.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   70752	     16753 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 2.811 / 0 allocs | 110.5 / 0 allocs | **39.3x** | N/A | N/A | N/A | N/A |
| Enum | 4.405 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 59.26 / 0 allocs | 1092 / 88 B / 5 allocs | **18.4x** | 956.4 / 0 allocs | **16.1x** | 17147 / 15847 B / 76 allocs | **289.4x** |
| GTE | 2.495 / 0 allocs | 106.1 / 0 allocs | **42.5x** | N/A | N/A | N/A | N/A |
| MinLength | 23.03 / 0 allocs | 120.1 / 0 allocs | **5.2x** | 282.0 / 32 B / 2 allocs | **12.2x** | N/A | N/A |
| UUID | 48.73 / 0 allocs | 448.7 / 0 allocs | **9.2x** | 353.1 / 0 allocs | **7.2x** | 16753 / 15542 B / 76 allocs | **343.8x** |
| MaxItems | 6.534 / 0 allocs | 144.7 / 0 allocs | **22.1x** | N/A | N/A | N/A | N/A |
| MaxLength | 33.45 / 0 allocs | 139.3 / 0 allocs | **4.2x** | 284.8 / 32 B / 2 allocs | **8.5x** | 16244 / 15648 B / 81 allocs | **485.6x** |
| LT | 2.489 / 0 allocs | 105.9 / 0 allocs | **42.5x** | N/A | N/A | N/A | N/A |
| MinItems | 5.922 / 0 allocs | 145.3 / 0 allocs | **24.5x** | N/A | N/A | N/A | N/A |
| Alpha | 9.670 / 0 allocs | 465.5 / 0 allocs | **48.1x** | 108.2 / 0 allocs | **11.2x** | 21239 / 16937 B / 101 allocs | **2196.4x** |
| Required | 3.734 / 0 allocs | 149.5 / 0 allocs | **40.0x** | 1.866 / 0 allocs | **0.5x** | 15585 / 15488 B / 73 allocs | **4173.8x** |
| Length | 7.602 / 0 allocs | 126.1 / 0 allocs | **16.6x** | 248.2 / 32 B / 2 allocs | **32.6x** | 16200 / 15616 B / 79 allocs | **2131.0x** |
| URL | 57.69 / 0 allocs | 484.2 / 144 B / 1 allocs | **8.4x** | 11748 / 146 B / 1 allocs | **203.6x** | 16903 / 15681 B / 77 allocs | **293.0x** |
| Numeric | 7.600 / 0 allocs | 98.42 / 0 allocs | **13.0x** | 128.2 / 0 allocs | **16.9x** | 16777 / 15574 B / 78 allocs | **2207.5x** |
| GT | 2.496 / 0 allocs | 106.2 / 0 allocs | **42.5x** | 90.16 / 0 allocs | **36.1x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 1.863 | 0 allocs |
| CELMultipleExpressions | 4.052 | 0 allocs |
| CELBasic | 4.104 | 0 allocs |
| CELCrossField | 3.426 | 0 allocs |
| CELStringLength | 1.000 | 0 allocs |
| CELNumericComparison | 1.000 | 0 allocs |

CEL (Common Expression Language) support allows complex runtime expressions with near-zero overhead.

## Running Benchmarks

```bash
# Update all benchmark documentation
make sync-benchmarks

# Run benchmarks manually
cd test
go test ./benchmark/ -bench=. -benchmem -benchtime=10s

# Run specific validator benchmarks
go test ./benchmark/ -bench=BenchmarkGoValid{ValidatorName} -benchmem
```
