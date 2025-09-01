# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-09-01  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	124429040	         9.643 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2595951	       467.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11239028	       108.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   55976	     21158 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	641371134	         2.100 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	296164554	         4.052 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	296150318	         4.053 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	350022814	         3.434 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21726850	        58.69 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1093 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1292262	       922.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   70232	     17095 ns/op	   15858 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	275415078	         4.355 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	480869821	         2.491 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11277176	       106.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13715640	        88.08 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	481747038	         2.489 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11311420	       106.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	160720177	         7.469 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	 9629497	       124.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4854328	       246.1 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   73652	     16243 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	482180725	         2.493 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11275071	       105.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	428604705	         2.800 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	10980741	       109.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	183471264	         6.536 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8401027	       142.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	35983508	        33.94 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8669482	       138.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4237066	       283.6 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   72862	     16372 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	203094508	         5.908 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8318750	       143.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	52065735	        23.01 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10058972	       119.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4256030	       281.9 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	157910763	         7.596 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	12173174	        98.59 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9440196	       128.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   71517	     16662 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	321391891	         3.733 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8111726	       147.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	642626446	         1.867 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   76149	     15662 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20861491	        57.55 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2484487	       483.0 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  102369	     11799 ns/op	     146 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   71349	     16862 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	24688887	        48.47 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2687160	       447.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3570574	       336.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   70665	     16744 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 2.800 / 0 allocs | 109.1 / 0 allocs | **39.0x** | N/A | N/A | N/A | N/A |
| Enum | 4.355 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 58.69 / 0 allocs | 1093 / 89 B / 5 allocs | **18.6x** | 922.1 / 0 allocs | **15.7x** | 17095 / 15858 B / 76 allocs | **291.3x** |
| GTE | 2.489 / 0 allocs | 106.2 / 0 allocs | **42.7x** | N/A | N/A | N/A | N/A |
| MinLength | 23.01 / 0 allocs | 119.3 / 0 allocs | **5.2x** | 281.9 / 32 B / 2 allocs | **12.3x** | N/A | N/A |
| UUID | 48.47 / 0 allocs | 447.8 / 0 allocs | **9.2x** | 336.0 / 0 allocs | **6.9x** | 16744 / 15542 B / 76 allocs | **345.5x** |
| MaxItems | 6.536 / 0 allocs | 142.8 / 0 allocs | **21.8x** | N/A | N/A | N/A | N/A |
| MaxLength | 33.94 / 0 allocs | 138.6 / 0 allocs | **4.1x** | 283.6 / 32 B / 2 allocs | **8.4x** | 16372 / 15648 B / 81 allocs | **482.4x** |
| LT | 2.493 / 0 allocs | 105.3 / 0 allocs | **42.2x** | N/A | N/A | N/A | N/A |
| MinItems | 5.908 / 0 allocs | 143.7 / 0 allocs | **24.3x** | N/A | N/A | N/A | N/A |
| Alpha | 9.643 / 0 allocs | 467.6 / 0 allocs | **48.5x** | 108.0 / 0 allocs | **11.2x** | 21158 / 16937 B / 101 allocs | **2194.1x** |
| Required | 3.733 / 0 allocs | 147.9 / 0 allocs | **39.6x** | 1.867 / 0 allocs | **0.5x** | 15662 / 15488 B / 73 allocs | **4195.6x** |
| Length | 7.469 / 0 allocs | 124.6 / 0 allocs | **16.7x** | 246.1 / 32 B / 2 allocs | **32.9x** | 16243 / 15616 B / 79 allocs | **2174.7x** |
| URL | 57.55 / 0 allocs | 483.0 / 144 B / 1 allocs | **8.4x** | 11799 / 146 B / 1 allocs | **205.0x** | 16862 / 15681 B / 77 allocs | **293.0x** |
| Numeric | 7.596 / 0 allocs | 98.59 / 0 allocs | **13.0x** | 128.3 / 0 allocs | **16.9x** | 16662 / 15574 B / 78 allocs | **2193.5x** |
| GT | 2.491 / 0 allocs | 106.5 / 0 allocs | **42.8x** | 88.08 / 0 allocs | **35.4x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 2.100 | 0 allocs |
| CELMultipleExpressions | 4.052 | 0 allocs |
| CELBasic | 4.053 | 0 allocs |
| CELCrossField | 3.434 | 0 allocs |
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
