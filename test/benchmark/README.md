# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-08-26  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	120362011	         9.969 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2534731	       470.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11218630	       107.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   57214	     20969 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	546931515	         2.195 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	274889450	         4.365 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	275411840	         4.359 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	350410024	         3.425 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21595183	        57.28 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1083 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1309476	       916.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   71017	     17146 ns/op	   15843 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	296611316	         4.054 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	428030625	         2.803 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11273010	       106.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13711806	        87.88 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	385453700	         3.115 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11366865	       106.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	160388305	         7.478 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	 9845650	       119.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4866484	       246.3 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   74502	     16467 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	385871349	         3.113 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11391650	       105.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	478232884	         2.498 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	10798215	       110.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	202960171	         5.915 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8255736	       146.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	37684195	        31.72 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8808352	       136.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4242062	       282.5 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   73663	     16251 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	192609885	         6.229 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8373139	       143.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	52065604	        23.04 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	 9859564	       121.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4219867	       283.6 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	124294587	         9.654 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	13851303	        86.91 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9389650	       127.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   70880	     16700 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	296765882	         4.044 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8079484	       150.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	551009914	         2.180 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   76527	     15777 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20679566	        58.12 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2515224	       477.5 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  103478	     11721 ns/op	     146 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   70856	     16817 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	24585777	        48.67 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2689354	       447.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3548436	       349.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   70470	     16890 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 2.498 / 0 allocs | 110.9 / 0 allocs | **44.4x** | N/A | N/A | N/A | N/A |
| Enum | 4.054 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 57.28 / 0 allocs | 1083 / 89 B / 5 allocs | **18.9x** | 916.3 / 0 allocs | **16.0x** | 17146 / 15843 B / 76 allocs | **299.3x** |
| GTE | 3.115 / 0 allocs | 106.5 / 0 allocs | **34.2x** | N/A | N/A | N/A | N/A |
| MinLength | 23.04 / 0 allocs | 121.6 / 0 allocs | **5.3x** | 283.6 / 32 B / 2 allocs | **12.3x** | N/A | N/A |
| UUID | 48.67 / 0 allocs | 447.7 / 0 allocs | **9.2x** | 349.1 / 0 allocs | **7.2x** | 16890 / 15542 B / 76 allocs | **347.0x** |
| MaxItems | 5.915 / 0 allocs | 146.4 / 0 allocs | **24.8x** | N/A | N/A | N/A | N/A |
| MaxLength | 31.72 / 0 allocs | 136.5 / 0 allocs | **4.3x** | 282.5 / 32 B / 2 allocs | **8.9x** | 16251 / 15648 B / 81 allocs | **512.3x** |
| LT | 3.113 / 0 allocs | 105.2 / 0 allocs | **33.8x** | N/A | N/A | N/A | N/A |
| MinItems | 6.229 / 0 allocs | 143.1 / 0 allocs | **23.0x** | N/A | N/A | N/A | N/A |
| Alpha | 9.969 / 0 allocs | 470.9 / 0 allocs | **47.2x** | 107.7 / 0 allocs | **10.8x** | 20969 / 16937 B / 101 allocs | **2103.4x** |
| Required | 4.044 / 0 allocs | 150.3 / 0 allocs | **37.2x** | 2.180 / 0 allocs | **0.5x** | 15777 / 15488 B / 73 allocs | **3901.3x** |
| Length | 7.478 / 0 allocs | 119.9 / 0 allocs | **16.0x** | 246.3 / 32 B / 2 allocs | **32.9x** | 16467 / 15616 B / 79 allocs | **2202.1x** |
| URL | 58.12 / 0 allocs | 477.5 / 144 B / 1 allocs | **8.2x** | 11721 / 146 B / 1 allocs | **201.7x** | 16817 / 15681 B / 77 allocs | **289.3x** |
| Numeric | 9.654 / 0 allocs | 86.91 / 0 allocs | **9.0x** | 127.9 / 0 allocs | **13.2x** | 16700 / 15574 B / 78 allocs | **1729.9x** |
| GT | 2.803 / 0 allocs | 106.4 / 0 allocs | **38.0x** | 87.88 / 0 allocs | **31.4x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 2.195 | 0 allocs |
| CELMultipleExpressions | 4.365 | 0 allocs |
| CELBasic | 4.359 | 0 allocs |
| CELCrossField | 3.425 | 0 allocs |
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
