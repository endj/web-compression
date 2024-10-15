# Gzips stats

## Payload Stats

| Payload   | Original Size (bytes) | Compressed Size (bytes) | Compression Ratio |
|-----------|-----------------------|-------------------------|-------------------|
| Payload 1 | 104,182                | 1,295                   | 80.44              |
| Payload 2 | 1,041,642              | 88,982                  | 11.71              |
| Payload 3 | 10,417,662             | 889,883                 | 11.71              |

## Test Results

### Throttling: 10000kbit/s, Latency: 15ms

| Payload   | Compression  | Payload Size | Waiting for Server | Content Download  |
|-----------|--------------|--------------|--------------------|-------------------|
| Payload 1 | Not Compressed | 104 KB       | 17.84 ms           | 102.07 ms         |
|           | Compressed    | 1.5 KB       | 10.75 ms           | 1.84 ms           |
| Payload 2 | Not Compressed | 1.0 MB       | 16.4 ms            | 832.72 ms         |
|           | Compressed    | 89.5 KB      | 11.53 ms           | 71.06 ms          |
| Payload 3 | Not Compressed | 10.4 MB      | 35.42 ms           | 8.39 s            |
|           | Compressed    | 893 KB       | 26.28 ms           | 712.01 ms         |
